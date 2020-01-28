package process

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"net/url"

	"flamingo.me/flamingo-commerce/v3/cart/domain/cart"
	"flamingo.me/flamingo-commerce/v3/cart/domain/validation"
	"flamingo.me/flamingo/v3/framework/flamingo"

	"github.com/google/uuid"
)

type (
	// Provider for Processes
	Provider func() *Process

	// Process representing a place order process and has a current context with infos about result and current state
	Process struct {
		context     Context
		allStates   map[string]State
		failedState State
		logger      flamingo.Logger
	}

	// Factory use to get Process instance
	Factory struct {
		provider    Provider
		startState  State
		failedState State
	}

	// RollbackReference a reference that can be used to trigger a rollback
	RollbackReference struct {
		StateName string
		Data      RollbackData
	}

	// RollbackData needed for rollback of a state
	RollbackData interface{}

	// FailedReason gives a human readable reason for a state failure
	FailedReason interface {
		Reason() string
	}

	// ErrorOccurredReason is used for unspecified errors
	ErrorOccurredReason struct {
		Error string
	}

	// CanceledByCustomerReason is used when customer cancels order
	CanceledByCustomerReason struct{}

	// PaymentErrorOccurredReason is used for errors during payment
	PaymentErrorOccurredReason struct {
		Error string
	}

	// CartValidationErrorReason contains the ValidationResult
	CartValidationErrorReason struct {
		ValidationResult validation.Result
	}
)

func init() {
	gob.Register(ErrorOccurredReason{})
	gob.Register(PaymentErrorOccurredReason{})
	gob.Register(CartValidationErrorReason{})
	gob.Register(CanceledByCustomerReason{})
}

// Reason for the error occurred
func (e ErrorOccurredReason) Reason() string {
	return e.Error
}

// Reason for the error occurred
func (e PaymentErrorOccurredReason) Reason() string {
	return e.Error
}

// Reason for the error occurred
func (e CanceledByCustomerReason) Reason() string {
	return "Order place canceled by customer"
}

// Reason for failing
func (e CartValidationErrorReason) Reason() string {
	return "Cart invalid"
}

// Inject dependencies
func (f *Factory) Inject(
	provider Provider,
	dep *struct {
		StartState  State `inject:"startState"`
		FailedState State `inject:"failedState"`
	},
) {
	f.provider = provider

	if dep != nil {
		f.failedState = dep.FailedState
		f.startState = dep.StartState
	}
}

// New process with initial state
func (f *Factory) New(returnURL *url.URL, cart cart.Cart) (*Process, error) {
	if f.startState == nil {
		return nil, errors.New("no start state given")
	}
	p := f.provider()
	p.failedState = f.failedState
	p.context = Context{
		UUID:              uuid.New().String(),
		CurrrentStateName: f.startState.Name(),
		Cart:              cart,
		ReturnURL:         returnURL,
	}

	return p, nil
}

// NewFromProcessContext returns a new process with given Context
func (f *Factory) NewFromProcessContext(pctx Context) (*Process, error) {
	p := f.provider()
	p.failedState = f.failedState
	p.context = pctx

	return p, nil
}

// Inject dependencies
func (p *Process) Inject(
	allStates map[string]State,
	logger flamingo.Logger,
) *Process {
	p.allStates = allStates
	p.logger = logger.
		WithField(flamingo.LogKeyModule, "checkout").
		WithField(flamingo.LogKeyCategory, "process")

	return p
}

// Run triggers run on current state
func (p *Process) Run(ctx context.Context) {
	currentState, err := p.CurrentState()
	if err != nil {
		p.Failed(ctx, ErrorOccurredReason{Error: err.Error()})
		return
	}

	stateBeforeRun := p.Context().CurrrentStateName
	runResult := currentState.Run(ctx, p, p.context.CurrrentStateData)
	if runResult.RollbackData != nil {
		p.context.RollbackReferences = append(p.context.RollbackReferences, RollbackReference{
			StateName: currentState.Name(),
			Data:      runResult.RollbackData,
		})
	}
	if runResult.Failed != nil {
		p.Failed(ctx, runResult.Failed)
	}
	stateAfterRun := p.Context().CurrrentStateName

	//Continue Run until no state change happend
	// TODO - protect endless loops with a max counter
	if stateBeforeRun != stateAfterRun {
		p.logger.Info(fmt.Sprintf("State Changed: %v => %v  Trigger Run() again", stateBeforeRun, stateAfterRun))
		p.Run(ctx)
	}
}

// CurrentState of the process context
func (p *Process) CurrentState() (State, error) {
	state, found := p.allStates[p.Context().CurrrentStateName]
	if !found {
		return nil, fmt.Errorf("current process context state %q not found", p.Context().CurrrentStateName)
	}
	return state, nil
}

func (p *Process) rollback(ctx context.Context) error {
	for i := len(p.context.RollbackReferences) - 1; i >= 0; i-- {
		rollbackRef := p.context.RollbackReferences[i]
		state, ok := p.allStates[rollbackRef.StateName]
		if !ok {
			p.logger.Error(fmt.Errorf("state %q not found for rollback", rollbackRef.StateName))
			continue
		}

		// todo error types for fatal end and continue rollback chain
		_ = state.Rollback(ctx, rollbackRef.Data)
	}

	return nil
}

// Context to get current process context
func (p *Process) Context() Context {
	return p.context
}

// UpdateState updates the current state in the context and its related state data
func (p *Process) UpdateState(s string, stateData StateData) {
	p.context.CurrrentStateName = s
	p.context.CurrrentStateData = stateData
}

// Failed performs all collected rollbacks and switches to FailedState
func (p *Process) Failed(ctx context.Context, reason FailedReason) {
	err := p.rollback(ctx)
	if err != nil {
		p.logger.WithContext(ctx).Error("rollback failed: ", err)
	}

	p.context.FailedReason = reason
	p.UpdateState(p.failedState.Name(), nil)
}
