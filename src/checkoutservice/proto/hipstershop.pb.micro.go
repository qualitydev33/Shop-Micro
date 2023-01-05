// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/hipstershop.proto

package hipstershop

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for CartService service

func NewCartServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CartService service

type CartService interface {
	AddItem(ctx context.Context, in *AddItemRequest, opts ...client.CallOption) (*Empty, error)
	GetCart(ctx context.Context, in *GetCartRequest, opts ...client.CallOption) (*Cart, error)
	EmptyCart(ctx context.Context, in *EmptyCartRequest, opts ...client.CallOption) (*Empty, error)
}

type cartService struct {
	c    client.Client
	name string
}

func NewCartService(name string, c client.Client) CartService {
	return &cartService{
		c:    c,
		name: name,
	}
}

func (c *cartService) AddItem(ctx context.Context, in *AddItemRequest, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "CartService.AddItem", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) GetCart(ctx context.Context, in *GetCartRequest, opts ...client.CallOption) (*Cart, error) {
	req := c.c.NewRequest(c.name, "CartService.GetCart", in)
	out := new(Cart)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartService) EmptyCart(ctx context.Context, in *EmptyCartRequest, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "CartService.EmptyCart", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CartService service

type CartServiceHandler interface {
	AddItem(context.Context, *AddItemRequest, *Empty) error
	GetCart(context.Context, *GetCartRequest, *Cart) error
	EmptyCart(context.Context, *EmptyCartRequest, *Empty) error
}

func RegisterCartServiceHandler(s server.Server, hdlr CartServiceHandler, opts ...server.HandlerOption) error {
	type cartService interface {
		AddItem(ctx context.Context, in *AddItemRequest, out *Empty) error
		GetCart(ctx context.Context, in *GetCartRequest, out *Cart) error
		EmptyCart(ctx context.Context, in *EmptyCartRequest, out *Empty) error
	}
	type CartService struct {
		cartService
	}
	h := &cartServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CartService{h}, opts...))
}

type cartServiceHandler struct {
	CartServiceHandler
}

func (h *cartServiceHandler) AddItem(ctx context.Context, in *AddItemRequest, out *Empty) error {
	return h.CartServiceHandler.AddItem(ctx, in, out)
}

func (h *cartServiceHandler) GetCart(ctx context.Context, in *GetCartRequest, out *Cart) error {
	return h.CartServiceHandler.GetCart(ctx, in, out)
}

func (h *cartServiceHandler) EmptyCart(ctx context.Context, in *EmptyCartRequest, out *Empty) error {
	return h.CartServiceHandler.EmptyCart(ctx, in, out)
}

// Api Endpoints for RecommendationService service

func NewRecommendationServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RecommendationService service

type RecommendationService interface {
	ListRecommendations(ctx context.Context, in *ListRecommendationsRequest, opts ...client.CallOption) (*ListRecommendationsResponse, error)
}

type recommendationService struct {
	c    client.Client
	name string
}

func NewRecommendationService(name string, c client.Client) RecommendationService {
	return &recommendationService{
		c:    c,
		name: name,
	}
}

func (c *recommendationService) ListRecommendations(ctx context.Context, in *ListRecommendationsRequest, opts ...client.CallOption) (*ListRecommendationsResponse, error) {
	req := c.c.NewRequest(c.name, "RecommendationService.ListRecommendations", in)
	out := new(ListRecommendationsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RecommendationService service

type RecommendationServiceHandler interface {
	ListRecommendations(context.Context, *ListRecommendationsRequest, *ListRecommendationsResponse) error
}

func RegisterRecommendationServiceHandler(s server.Server, hdlr RecommendationServiceHandler, opts ...server.HandlerOption) error {
	type recommendationService interface {
		ListRecommendations(ctx context.Context, in *ListRecommendationsRequest, out *ListRecommendationsResponse) error
	}
	type RecommendationService struct {
		recommendationService
	}
	h := &recommendationServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RecommendationService{h}, opts...))
}

type recommendationServiceHandler struct {
	RecommendationServiceHandler
}

func (h *recommendationServiceHandler) ListRecommendations(ctx context.Context, in *ListRecommendationsRequest, out *ListRecommendationsResponse) error {
	return h.RecommendationServiceHandler.ListRecommendations(ctx, in, out)
}

// Api Endpoints for ProductCatalogService service

func NewProductCatalogServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ProductCatalogService service

type ProductCatalogService interface {
	ListProducts(ctx context.Context, in *Empty, opts ...client.CallOption) (*ListProductsResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...client.CallOption) (*Product, error)
	SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...client.CallOption) (*SearchProductsResponse, error)
}

type productCatalogService struct {
	c    client.Client
	name string
}

func NewProductCatalogService(name string, c client.Client) ProductCatalogService {
	return &productCatalogService{
		c:    c,
		name: name,
	}
}

func (c *productCatalogService) ListProducts(ctx context.Context, in *Empty, opts ...client.CallOption) (*ListProductsResponse, error) {
	req := c.c.NewRequest(c.name, "ProductCatalogService.ListProducts", in)
	out := new(ListProductsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogService) GetProduct(ctx context.Context, in *GetProductRequest, opts ...client.CallOption) (*Product, error) {
	req := c.c.NewRequest(c.name, "ProductCatalogService.GetProduct", in)
	out := new(Product)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogService) SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...client.CallOption) (*SearchProductsResponse, error) {
	req := c.c.NewRequest(c.name, "ProductCatalogService.SearchProducts", in)
	out := new(SearchProductsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductCatalogService service

type ProductCatalogServiceHandler interface {
	ListProducts(context.Context, *Empty, *ListProductsResponse) error
	GetProduct(context.Context, *GetProductRequest, *Product) error
	SearchProducts(context.Context, *SearchProductsRequest, *SearchProductsResponse) error
}

func RegisterProductCatalogServiceHandler(s server.Server, hdlr ProductCatalogServiceHandler, opts ...server.HandlerOption) error {
	type productCatalogService interface {
		ListProducts(ctx context.Context, in *Empty, out *ListProductsResponse) error
		GetProduct(ctx context.Context, in *GetProductRequest, out *Product) error
		SearchProducts(ctx context.Context, in *SearchProductsRequest, out *SearchProductsResponse) error
	}
	type ProductCatalogService struct {
		productCatalogService
	}
	h := &productCatalogServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProductCatalogService{h}, opts...))
}

type productCatalogServiceHandler struct {
	ProductCatalogServiceHandler
}

func (h *productCatalogServiceHandler) ListProducts(ctx context.Context, in *Empty, out *ListProductsResponse) error {
	return h.ProductCatalogServiceHandler.ListProducts(ctx, in, out)
}

func (h *productCatalogServiceHandler) GetProduct(ctx context.Context, in *GetProductRequest, out *Product) error {
	return h.ProductCatalogServiceHandler.GetProduct(ctx, in, out)
}

func (h *productCatalogServiceHandler) SearchProducts(ctx context.Context, in *SearchProductsRequest, out *SearchProductsResponse) error {
	return h.ProductCatalogServiceHandler.SearchProducts(ctx, in, out)
}

// Api Endpoints for ShippingService service

func NewShippingServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ShippingService service

type ShippingService interface {
	GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...client.CallOption) (*GetQuoteResponse, error)
	ShipOrder(ctx context.Context, in *ShipOrderRequest, opts ...client.CallOption) (*ShipOrderResponse, error)
}

type shippingService struct {
	c    client.Client
	name string
}

func NewShippingService(name string, c client.Client) ShippingService {
	return &shippingService{
		c:    c,
		name: name,
	}
}

func (c *shippingService) GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...client.CallOption) (*GetQuoteResponse, error) {
	req := c.c.NewRequest(c.name, "ShippingService.GetQuote", in)
	out := new(GetQuoteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingService) ShipOrder(ctx context.Context, in *ShipOrderRequest, opts ...client.CallOption) (*ShipOrderResponse, error) {
	req := c.c.NewRequest(c.name, "ShippingService.ShipOrder", in)
	out := new(ShipOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	GetQuote(context.Context, *GetQuoteRequest, *GetQuoteResponse) error
	ShipOrder(context.Context, *ShipOrderRequest, *ShipOrderResponse) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) error {
	type shippingService interface {
		GetQuote(ctx context.Context, in *GetQuoteRequest, out *GetQuoteResponse) error
		ShipOrder(ctx context.Context, in *ShipOrderRequest, out *ShipOrderResponse) error
	}
	type ShippingService struct {
		shippingService
	}
	h := &shippingServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ShippingService{h}, opts...))
}

type shippingServiceHandler struct {
	ShippingServiceHandler
}

func (h *shippingServiceHandler) GetQuote(ctx context.Context, in *GetQuoteRequest, out *GetQuoteResponse) error {
	return h.ShippingServiceHandler.GetQuote(ctx, in, out)
}

func (h *shippingServiceHandler) ShipOrder(ctx context.Context, in *ShipOrderRequest, out *ShipOrderResponse) error {
	return h.ShippingServiceHandler.ShipOrder(ctx, in, out)
}

// Api Endpoints for CurrencyService service

func NewCurrencyServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CurrencyService service

type CurrencyService interface {
	GetSupportedCurrencies(ctx context.Context, in *Empty, opts ...client.CallOption) (*GetSupportedCurrenciesResponse, error)
	Convert(ctx context.Context, in *CurrencyConversionRequest, opts ...client.CallOption) (*Money, error)
}

type currencyService struct {
	c    client.Client
	name string
}

func NewCurrencyService(name string, c client.Client) CurrencyService {
	return &currencyService{
		c:    c,
		name: name,
	}
}

func (c *currencyService) GetSupportedCurrencies(ctx context.Context, in *Empty, opts ...client.CallOption) (*GetSupportedCurrenciesResponse, error) {
	req := c.c.NewRequest(c.name, "CurrencyService.GetSupportedCurrencies", in)
	out := new(GetSupportedCurrenciesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyService) Convert(ctx context.Context, in *CurrencyConversionRequest, opts ...client.CallOption) (*Money, error) {
	req := c.c.NewRequest(c.name, "CurrencyService.Convert", in)
	out := new(Money)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CurrencyService service

type CurrencyServiceHandler interface {
	GetSupportedCurrencies(context.Context, *Empty, *GetSupportedCurrenciesResponse) error
	Convert(context.Context, *CurrencyConversionRequest, *Money) error
}

func RegisterCurrencyServiceHandler(s server.Server, hdlr CurrencyServiceHandler, opts ...server.HandlerOption) error {
	type currencyService interface {
		GetSupportedCurrencies(ctx context.Context, in *Empty, out *GetSupportedCurrenciesResponse) error
		Convert(ctx context.Context, in *CurrencyConversionRequest, out *Money) error
	}
	type CurrencyService struct {
		currencyService
	}
	h := &currencyServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CurrencyService{h}, opts...))
}

type currencyServiceHandler struct {
	CurrencyServiceHandler
}

func (h *currencyServiceHandler) GetSupportedCurrencies(ctx context.Context, in *Empty, out *GetSupportedCurrenciesResponse) error {
	return h.CurrencyServiceHandler.GetSupportedCurrencies(ctx, in, out)
}

func (h *currencyServiceHandler) Convert(ctx context.Context, in *CurrencyConversionRequest, out *Money) error {
	return h.CurrencyServiceHandler.Convert(ctx, in, out)
}

// Api Endpoints for PaymentService service

func NewPaymentServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PaymentService service

type PaymentService interface {
	Charge(ctx context.Context, in *ChargeRequest, opts ...client.CallOption) (*ChargeResponse, error)
}

type paymentService struct {
	c    client.Client
	name string
}

func NewPaymentService(name string, c client.Client) PaymentService {
	return &paymentService{
		c:    c,
		name: name,
	}
}

func (c *paymentService) Charge(ctx context.Context, in *ChargeRequest, opts ...client.CallOption) (*ChargeResponse, error) {
	req := c.c.NewRequest(c.name, "PaymentService.Charge", in)
	out := new(ChargeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PaymentService service

type PaymentServiceHandler interface {
	Charge(context.Context, *ChargeRequest, *ChargeResponse) error
}

func RegisterPaymentServiceHandler(s server.Server, hdlr PaymentServiceHandler, opts ...server.HandlerOption) error {
	type paymentService interface {
		Charge(ctx context.Context, in *ChargeRequest, out *ChargeResponse) error
	}
	type PaymentService struct {
		paymentService
	}
	h := &paymentServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PaymentService{h}, opts...))
}

type paymentServiceHandler struct {
	PaymentServiceHandler
}

func (h *paymentServiceHandler) Charge(ctx context.Context, in *ChargeRequest, out *ChargeResponse) error {
	return h.PaymentServiceHandler.Charge(ctx, in, out)
}

// Api Endpoints for EmailService service

func NewEmailServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for EmailService service

type EmailService interface {
	SendOrderConfirmation(ctx context.Context, in *SendOrderConfirmationRequest, opts ...client.CallOption) (*Empty, error)
}

type emailService struct {
	c    client.Client
	name string
}

func NewEmailService(name string, c client.Client) EmailService {
	return &emailService{
		c:    c,
		name: name,
	}
}

func (c *emailService) SendOrderConfirmation(ctx context.Context, in *SendOrderConfirmationRequest, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "EmailService.SendOrderConfirmation", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EmailService service

type EmailServiceHandler interface {
	SendOrderConfirmation(context.Context, *SendOrderConfirmationRequest, *Empty) error
}

func RegisterEmailServiceHandler(s server.Server, hdlr EmailServiceHandler, opts ...server.HandlerOption) error {
	type emailService interface {
		SendOrderConfirmation(ctx context.Context, in *SendOrderConfirmationRequest, out *Empty) error
	}
	type EmailService struct {
		emailService
	}
	h := &emailServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&EmailService{h}, opts...))
}

type emailServiceHandler struct {
	EmailServiceHandler
}

func (h *emailServiceHandler) SendOrderConfirmation(ctx context.Context, in *SendOrderConfirmationRequest, out *Empty) error {
	return h.EmailServiceHandler.SendOrderConfirmation(ctx, in, out)
}

// Api Endpoints for CheckoutService service

func NewCheckoutServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CheckoutService service

type CheckoutService interface {
	PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...client.CallOption) (*PlaceOrderResponse, error)
}

type checkoutService struct {
	c    client.Client
	name string
}

func NewCheckoutService(name string, c client.Client) CheckoutService {
	return &checkoutService{
		c:    c,
		name: name,
	}
}

func (c *checkoutService) PlaceOrder(ctx context.Context, in *PlaceOrderRequest, opts ...client.CallOption) (*PlaceOrderResponse, error) {
	req := c.c.NewRequest(c.name, "CheckoutService.PlaceOrder", in)
	out := new(PlaceOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CheckoutService service

type CheckoutServiceHandler interface {
	PlaceOrder(context.Context, *PlaceOrderRequest, *PlaceOrderResponse) error
}

func RegisterCheckoutServiceHandler(s server.Server, hdlr CheckoutServiceHandler, opts ...server.HandlerOption) error {
	type checkoutService interface {
		PlaceOrder(ctx context.Context, in *PlaceOrderRequest, out *PlaceOrderResponse) error
	}
	type CheckoutService struct {
		checkoutService
	}
	h := &checkoutServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CheckoutService{h}, opts...))
}

type checkoutServiceHandler struct {
	CheckoutServiceHandler
}

func (h *checkoutServiceHandler) PlaceOrder(ctx context.Context, in *PlaceOrderRequest, out *PlaceOrderResponse) error {
	return h.CheckoutServiceHandler.PlaceOrder(ctx, in, out)
}

// Api Endpoints for AdService service

func NewAdServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AdService service

type AdService interface {
	GetAds(ctx context.Context, in *AdRequest, opts ...client.CallOption) (*AdResponse, error)
}

type adService struct {
	c    client.Client
	name string
}

func NewAdService(name string, c client.Client) AdService {
	return &adService{
		c:    c,
		name: name,
	}
}

func (c *adService) GetAds(ctx context.Context, in *AdRequest, opts ...client.CallOption) (*AdResponse, error) {
	req := c.c.NewRequest(c.name, "AdService.GetAds", in)
	out := new(AdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AdService service

type AdServiceHandler interface {
	GetAds(context.Context, *AdRequest, *AdResponse) error
}

func RegisterAdServiceHandler(s server.Server, hdlr AdServiceHandler, opts ...server.HandlerOption) error {
	type adService interface {
		GetAds(ctx context.Context, in *AdRequest, out *AdResponse) error
	}
	type AdService struct {
		adService
	}
	h := &adServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AdService{h}, opts...))
}

type adServiceHandler struct {
	AdServiceHandler
}

func (h *adServiceHandler) GetAds(ctx context.Context, in *AdRequest, out *AdResponse) error {
	return h.AdServiceHandler.GetAds(ctx, in, out)
}
