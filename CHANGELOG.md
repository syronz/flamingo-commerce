# Changelog

## v3

- general cleanups and linting fixes - that includes several renames of packagenames and types.
- price object introduced:
    - cart and product model don't use float64 anymore but a Price type
    - use commercePriceFormat templatefunc instead (core) priceFormat where you want to render a price object. This will automatically render a "Payable" price.
- cart module:
    - Has a new secondary port: PlaceOrderService
    - The meaning of DeliveryInfo.Method has changed! The former meaning is now represented in the property DeliveryInfo.Workflow. See Readme of cart ackage for details
    - The complete pricefields are changed! Check readme for details on the new price fields and methods
- checkout: 
    - removed depricated viewdata (CartTotals)
- products:
    - product category breadcrumb is not filled in controller - if you want a breadcrum you can use category data functions
    - product category fields are changed to use a categoryTeaser
- category:
    - Tree object uses a Tree Entity now which contains NOT all category properties. You have to fetch the category details separate on demand:
        - search for usages of the data funcs - they may need changes in rendering the data: `data('category´´..`

## v3.0.1
- Update dingo and form dependency to latest version
        
## v3.1.0
**tests**
* Added GraphQL integration tests for new Place Order Process, run manually with `make integrationtest`
* To run the GraphQL Demo project use `make run-integrationtest-demo-project`
* To regenerate the GraphQL files used by the integration tests / demo project use  `make generate-integrationtest-graphql`

**cart**
* Add `additionalData` to the `AddRequest` used during add to cart
    * Breaking: Update helper/builder function `BuildAddRequest`
* Breaking: Change to `EventPublisher` interface, `PublishChangedQtyInCartEvent` and `PublishAddToCartEvent` now
include a cart as a parameter
* Breaking: Change to behaviour of `AddToCartEvent` and `ChangedQtyInCartEvent`, they are now thrown after
the cart has been adjusted and written back to cache
* Events deferred from `ModifyBehaviour` are dispatched before `AddToCartEvent` and `ChangedQtyInCartEvent`
* The `AddToCartEvent` includes the current cart (with added product)
* The `ChangedQtyInCartEvent` includes the current cart (with updated quantities)

* Mark `CartReceiverService.RestoreCart()` as deprecated, use `CartService.RestoreCart()` instead,
  the cart adapter therefore needs to implement the `CompleteBehaviour` interface.
* Add `CartReceiverService.ModifyBehaviour()` to easily receive the current behaviour (guest/customer)

* Add `CompleteBehaviour` interface which ensures that the cart adapter offers Complete / Restore functionality
* Add `CartService.CompleteCurrentCart()` and `CartService.RestoreCart()` which rely on the new `CompleteBehaviour` interface
* **Breaking**: Update `CartService.CancelOrder()` to use `CartService.RestoreCart()` instead of `CartReceiverService.RestoreCart()`,
  if your cart supports completing/restoring please implement `CompleteBehaviour` interface
* Add `CartService.CancelOrderWithoutRestore()` to allow order cancellation without restoring the cart

* Mark `GuestCartService.RestoreCart` as deprecated, will be replaced by `CompleteBehaviour`
* Mark `CustomerCartService.RestoreCart` as deprecated, will be replaced by `CompleteBehaviour`

* Add mocks for all behaviours, you can use a specific one e.g. `&mocks.CompleteBehaviour{}` or the all in one `&mocks.AllBehaviour{}`

* Update `InMemoryBehaviour` to fulfill the `CompleteBehaviour` interface (adds `Complete()`/`Restore()`)
* Update `InMemoryCartStorage`, add Mutex to be thread safe

* Update `SimplePaymentFormService` to allow gift cards in the `PaymentSelection`, please use the
  config `commerce.cart.simplePaymentForm.giftCardPaymentMethod`to specify the default payment method for gift cards

* Add missing `product` module dependency to cart module

**checkout**
* Move config to commerce namespace, from `checkout` to `commerce.checkout`
* Add legacy config mapping so old mappings can still be used
* Add cue based config to have config validation in place

* Add `OrderService.CancelOrderWithoutRestore()` which uses the new `CartService` function
* Add `OrderService.CartPlaceOrder()` to place a provided cart instead of fetching it from the `CartService`

* Add new GraphQL Place Order process which relies on a new state machine please referer to the module readme for more details
    * Transition all actions of the checkout controller to separate states
    * Add new `ContextStore` port to provide a storage for the place order process
        * Provide InMemory and Redis Adapter
    * Add new `TryLocker` port to provide an easy way to sync multiple order processes across different nodes
        * Provide InMemory and Redis Adapter
    * Breaking: Add new GraphQL mutations / queries to start / stop / refresh the place order process
    
**payment**
* Add `PaymentService` to easily work with bound PaymentGateway's
    * `PaymentService.AvailablePaymentGateways()` returns all bound gateways
    * `PaymentService.PaymentGateway()` gets the payment gateway by gateway code
    * `PaymentService.PaymentGatewayByCart()` gets the payment gateway of the cart payment selection

* Extend the `FlowStatus` struct with more standardized `FlowActionData`
* Add standardized Flow Actions `PaymentFlowActionShowIframe`, `PaymentFlowActionShowHTML`, `PaymentFlowActionRedirect`,
  `PaymentFlowActionPostRedirect` please use these in your payment adapter since the standard place order relies on them.
 
**search**
* Extend `Suggestion` struct with `Type` and `AdditionalAttributes` to be able to distinguish between product/category suggestions

## v3.1.1 [upcoming]
**w3cdatalayer**
* Fixed a bug that causes the datalayer to panic if it failed to build an absolute url