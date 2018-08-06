package application

import (
	"fmt"

	cartDomain "flamingo.me/flamingo-commerce/cart/domain/cart"
	productDomain "flamingo.me/flamingo-commerce/product/domain"
	"flamingo.me/flamingo/framework/flamingo"
	"flamingo.me/flamingo/framework/web"
	"github.com/pkg/errors"
)

// CartService application struct
type (
	//CartService provides methods to modify the cart
	CartService struct {
		CartReceiverService *CartReceiverService         `inject:""`
		ProductService      productDomain.ProductService `inject:""`
		Logger              flamingo.Logger              `inject:""`
		CartValidator       cartDomain.CartValidator     `inject:",optional"`

		ItemValidator  cartDomain.ItemValidator `inject:",optional"`
		EventPublisher EventPublisher           `inject:""`

		DefaultDeliveryCode string `inject:"config:cart.defaultDeliveryCode,optional"`

		DeliveryInfoBuilder cartDomain.DeliveryInfoBuilder `inject:""`

		CartCache CartCache `inject:",optional"`
	}
)

// ValidateCart validates a carts content
func (cs CartService) ValidateCart(ctx web.Context, decoratedCart *cartDomain.DecoratedCart) cartDomain.CartValidationResult {

	if cs.CartValidator != nil {
		result := cs.CartValidator.Validate(ctx, decoratedCart)
		return result
	}

	return cartDomain.CartValidationResult{}
}

// ValidateCurrentCart validates the current active cart
func (cs CartService) ValidateCurrentCart(ctx web.Context) (cartDomain.CartValidationResult, error) {
	decoratedCart, err := cs.CartReceiverService.ViewDecoratedCart(ctx)
	if err != nil {
		return cartDomain.CartValidationResult{}, err
	}

	return cs.ValidateCart(ctx, decoratedCart), nil
}

func (cs CartService) UpdateBillingAddress(ctx web.Context, billingAddress *cartDomain.Address) error {
	cart, behaviour, err := cs.CartReceiverService.GetCart(ctx)
	if err != nil {
		return err
	}
	cart, err = behaviour.UpdateBillingAddress(ctx, cart, billingAddress)
	if err != nil {
		cs.handleCartNotFound(ctx, err)
		cs.Logger.WithField("category", "cartService").WithField("subCategory", "UpdateItemQty").Error(err)
		return err
	}
	cs.updateCartInCache(ctx, cart)
	return nil
}

func (cs CartService) UpdateDeliveryInfo(ctx web.Context, deliveryCode string, deliveryInfo cartDomain.DeliveryInfo) error {
	cart, behaviour, err := cs.CartReceiverService.GetCart(ctx)
	if err != nil {
		return err
	}
	if deliveryCode == "" {
		deliveryCode = cs.DefaultDeliveryCode
	}
	cart, err = behaviour.UpdateDeliveryInfo(ctx, cart, deliveryCode, deliveryInfo)
	if err != nil {
		cs.handleCartNotFound(ctx, err)
		cs.Logger.WithField("category", "cartService").WithField("subCategory", "UpdateItemQty").Error(err)
		return err
	}
	cs.updateCartInCache(ctx, cart)
	return nil
}

func (cs CartService) UpdatePurchaser(ctx web.Context, purchaser *cartDomain.Person, additionalData map[string]string) error {
	cart, behaviour, err := cs.CartReceiverService.GetCart(ctx)
	if err != nil {
		return err
	}
	cart, err = behaviour.UpdatePurchaser(ctx, cart, purchaser, additionalData)
	if err != nil {
		cs.handleCartNotFound(ctx, err)
		cs.Logger.WithField("category", "cartService").WithField("subCategory", "UpdateItemQty").Error(err)
		return err
	}
	cs.updateCartInCache(ctx, cart)
	return nil
}

// UpdateItemQty
func (cs CartService) UpdateItemQty(ctx web.Context, itemId string, deliveryCode string, qty int) error {
	cart, behaviour, err := cs.CartReceiverService.GetCart(ctx)
	if err != nil {
		return err
	}
	if deliveryCode == "" {
		deliveryCode = cs.DefaultDeliveryCode
	}
	item, err := cart.GetByItemId(itemId, deliveryCode)
	if err != nil {
		cs.Logger.WithField("category", "cartService").WithField("subCategory", "UpdateItemQty").Error(err)
		return err
	}
	qtyBefore := item.Qty
	if qty < 1 {
		return cs.DeleteItem(ctx, itemId, deliveryCode)
	}

	cs.EventPublisher.PublishChangedQtyInCartEvent(ctx, item, qtyBefore, qty, cart.ID)
	itemUpdate := cartDomain.ItemUpdateCommand{
		Qty: &qty,
	}
	cart, err = behaviour.UpdateItem(ctx, cart, itemId, deliveryCode, itemUpdate)
	if err != nil {
		cs.handleCartNotFound(ctx, err)
		cs.Logger.WithField("category", "cartService").WithField("subCategory", "UpdateItemQty").Error(err)
		return err
	}
	cs.updateCartInCache(ctx, cart)
	return nil
}

func (cs CartService) UpdateItemSourceId(ctx web.Context, itemId string, deliveryCode string, sourceId string) error {
	cart, behaviour, err := cs.CartReceiverService.GetCart(ctx)
	if err != nil {
		return err
	}
	if deliveryCode == "" {
		deliveryCode = cs.DefaultDeliveryCode
	}
	_, err = cart.GetByItemId(itemId, deliveryCode)
	if err != nil {
		cs.Logger.WithField("category", "cartService").WithField("subCategory", "UpdateItemSourceId").Error(err)
		return err
	}
	itemUpdate := cartDomain.ItemUpdateCommand{
		SourceId: &sourceId,
	}
	cart, err = behaviour.UpdateItem(ctx, cart, itemId, deliveryCode, itemUpdate)
	if err != nil {
		cs.handleCartNotFound(ctx, err)
		cs.Logger.WithField("category", "cartService").WithField("subCategory", "UpdateItemSourceId").Error(err)
		return err
	}
	cs.updateCartInCache(ctx, cart)
	return nil
}

// DeleteItem in current cart
func (cs CartService) DeleteItem(ctx web.Context, itemId string, deliveryCode string) error {
	cart, behaviour, err := cs.CartReceiverService.GetCart(ctx)
	if err != nil {
		return err
	}
	if deliveryCode == "" {
		deliveryCode = cs.DefaultDeliveryCode
	}
	item, err := cart.GetByItemId(itemId, deliveryCode)
	if err != nil {
		cs.Logger.WithField("category", "cartService").WithField("subCategory", "DeleteItem").Error(err)
		return err
	}
	qtyBefore := item.Qty
	cs.EventPublisher.PublishChangedQtyInCartEvent(ctx, item, qtyBefore, 0, cart.ID)
	cart, err = behaviour.DeleteItem(ctx, cart, itemId, deliveryCode)
	if err != nil {
		cs.handleCartNotFound(ctx, err)
		cs.Logger.WithField("category", "cartService").WithField("subCategory", "DeleteItem").Error(err)
		return err
	}
	cs.updateCartInCache(ctx, cart)
	return nil
}

// DeleteAllItems in current cart
func (cs CartService) DeleteAllItems(ctx web.Context) error {
	cart, behaviour, err := cs.CartReceiverService.GetCart(ctx)
	if err != nil {
		return err
	}
	for _, delivery := range cart.Deliveries {
		for _, item := range delivery.Cartitems {
			qtyBefore := item.Qty
			cs.EventPublisher.PublishChangedQtyInCartEvent(ctx, &item, qtyBefore, 0, cart.ID)
			cart, err = behaviour.DeleteItem(ctx, cart, item.ID, delivery.DeliveryInfo.Code)
			if err != nil {
				cs.handleCartNotFound(ctx, err)
				cs.Logger.WithField("category", "cartService").WithField("subCategory", "DeleteAllItems").Error(err)
				return err
			}
		}
	}

	cs.updateCartInCache(ctx, cart)
	return nil
}

// PlaceOrder
func (cs *CartService) PlaceOrder(ctx web.Context, payment *cartDomain.CartPayment) (string, error) {
	cart, behaviour, err := cs.CartReceiverService.GetCart(ctx)
	if err != nil {
		return "", err
	}

	orderNumber, err := behaviour.PlaceOrder(ctx, cart, payment)
	if err != nil {
		cs.handleCartNotFound(ctx, err)
		cs.Logger.WithField("category", "cartService").WithField("subCategory", "PlaceOrder").Error(err)
		return "", err
	}
	cs.EventPublisher.PublishOrderPlacedEvent(ctx, cart, orderNumber)
	cs.DeleteSavedSessionGuestCartId(ctx)
	cs.deleteCartInCache(ctx, cart)
	return orderNumber, err
}

// BuildAddRequest Helper to build
func (cs *CartService) BuildAddRequest(ctx web.Context, marketplaceCode string, variantMarketplaceCode string, qty int) cartDomain.AddRequest {
	if qty < 0 {
		qty = 0
	}

	return cartDomain.AddRequest{
		MarketplaceCode: marketplaceCode,
		Qty:             qty,
		VariantMarketplaceCode: variantMarketplaceCode,
	}
}

// AddProduct Add a product
func (cs *CartService) AddProduct(ctx web.Context, deliveryCode string, addRequest cartDomain.AddRequest) (error, productDomain.BasicProduct) {
	if deliveryCode == "" {
		deliveryCode = cs.DefaultDeliveryCode
	}

	addRequest, product, err := cs.checkProductForAddRequest(ctx, deliveryCode, addRequest)
	if err != nil {
		cs.Logger.WithField("category", "cartService").WithField(flamingo.LogKeySubCategory, "AddProduct").Error(err)
		return err, nil
	}

	cs.Logger.WithField(flamingo.LogKeyCategory, "cartService").WithField(flamingo.LogKeySubCategory, "AddProduct").Debug("AddRequest received %#v  / %v", addRequest, deliveryCode)

	cart, behaviour, err := cs.CartReceiverService.GetCart(ctx)
	cart, err = cs.CreateInitialDeliveryIfNotPresent(ctx, deliveryCode)

	if err != nil {
		cs.Logger.WithField("category", "cartService").WithField(flamingo.LogKeySubCategory, "AddProduct").Error(err)
		return err, nil
	}

	if err != nil {
		cs.Logger.WithField("category", "cartService").WithField("subCategory", "AddProduct").Error(err)
		return err, nil
	}

	cart, err = behaviour.AddToCart(ctx, cart, deliveryCode, addRequest)
	if err == cartDomain.DeliveryCodeNotFound {
		//old Magento adapter will never return this
		//Edge case...
		// For later - todo:
		// call initialCreateDelivery
		// retry AddToCart
	}

	if err != nil {
		cs.handleCartNotFound(ctx, err)
		cs.Logger.WithField("category", "cartService").WithField("subCategory", "AddProduct").Error(err)
		return err, nil
	}
	cs.publishAddtoCartEvent(ctx, *cart, addRequest)
	cs.updateCartInCache(ctx, cart)
	return nil, product
}

func (cs *CartService) CreateInitialDeliveryIfNotPresent(ctx web.Context, deliveryCode string) (*cartDomain.Cart, error) {
	cart, behaviour, err := cs.CartReceiverService.GetCart(ctx)

	if cart.HasDeliveryForCode(deliveryCode) {
		return cart, nil
	}
	delInfo, err := cs.DeliveryInfoBuilder.BuildByDeliveryCode(deliveryCode)
	if err != nil {
		return nil, err
	}
	return behaviour.UpdateDeliveryInfo(ctx, cart, deliveryCode, delInfo)
}

func (cs *CartService) ApplyVoucher(ctx web.Context, couponCode string) (*cartDomain.Cart, error) {

	oldCart, behaviour, err := cs.CartReceiverService.GetCart(ctx)
	if err != nil {
		cs.Logger.WithField("category", "cartService").WithField("subCategory", "ApplyVoucher").Error(err)
		return nil, err
	}

	cart, err := behaviour.ApplyVoucher(ctx, oldCart, couponCode)
	cs.updateCartInCache(ctx, cart)
	return cart, err
}

func (cs *CartService) handleCartNotFound(ctx web.Context, err error) {
	if err == cartDomain.CartNotFoundError {
		cs.DeleteSavedSessionGuestCartId(ctx)
	}
}

// checkProductForAddRequest existence and validate with productService
func (cs *CartService) checkProductForAddRequest(ctx web.Context, deliveryCode string, addRequest cartDomain.AddRequest) (cartDomain.AddRequest, productDomain.BasicProduct, error) {
	product, err := cs.ProductService.Get(ctx, addRequest.MarketplaceCode)
	if err != nil {
		return addRequest, nil, fmt.Errorf("cart.application.cartservice - AddProduct Error: %v", err)
	}

	if product.Type() == productDomain.TYPECONFIGURABLE {
		if addRequest.VariantMarketplaceCode == "" {
			return addRequest, nil, errors.New("cart.application.cartservice - AddProduct:No Variant given for configurable product")
		}

		configurableProduct := product.(productDomain.ConfigurableProduct)
		if !configurableProduct.HasVariant(addRequest.VariantMarketplaceCode) {
			return addRequest, nil, errors.New("cart.application.cartservice - AddProduct:Product has not the given variant")
		}
	}

	//Now Validate the Item with the optional registered ItemValidator
	if cs.ItemValidator != nil {
		return addRequest, product, cs.ItemValidator.Validate(ctx, deliveryCode, addRequest, product)
	}

	return addRequest, product, nil
}

func (cs *CartService) publishAddtoCartEvent(ctx web.Context, currentCart cartDomain.Cart, addRequest cartDomain.AddRequest) {
	if cs.EventPublisher != nil {
		cs.EventPublisher.PublishAddToCartEvent(ctx, addRequest.MarketplaceCode, addRequest.VariantMarketplaceCode, addRequest.Qty)
	}
}

func (cs *CartService) updateCartInCache(ctx web.Context, cart *cartDomain.Cart) {
	if cs.CartCache != nil {
		id, err := BuildIdentifierFromCart(cart)
		if err != nil {
			return
		}
		err = cs.CartCache.CacheCart(ctx, *id, cart)
		if err != nil {
			cs.Logger.Error("Error while caching cart: %v", err)
		}
	}
}

func (cs *CartService) deleteCartInCache(ctx web.Context, cart *cartDomain.Cart) {
	if cs.CartCache != nil {
		id, err := BuildIdentifierFromCart(cart)
		if err != nil {
			return
		}
		err = cs.CartCache.Delete(ctx, *id)
		if err != nil {
			cs.Logger.Error("Error while deleting cart in cache: %v", err)
		}
	}
}
