package orders

import (
	"grpc-microservice/services/common/genproto/orders"
	"grpc-microservice/services/common/util"
	"grpc-microservice/services/orders/types"
	"net/http"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersService(ordersService types.OrderService) *OrdersHttpHandler {
	return &OrdersHttpHandler{
		ordersService: ordersService,
	}
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)

	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	util.WriteJSON(w, http.StatusOK, res)
}
