// Code generated by goctl. DO NOT EDIT.
package types

type RegisterRequest struct {
	UserName string `json:"userName" validate:"notblank"`
	Password string `json:"password" validate:"notblank"`
	UserType string `json:"userType" validate:"oneof=normal vip"`
}

type LoginRequest struct {
	UserName string `json:"userName" validate:"notblank"`
	Password string `json:"password" validate:"notblank"`
}

type UpgradeToVipRequest struct {
	UserId string `json:"userId" validate:"notblank"`
}

type BaseResponse struct {
	RetCode int    `json:"code"`
	Message string `json:"message"`
}

type User struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	UserType string `json:"userType" validate:"oneof=normal vip"` // 用户类型 normal, vip……
}

type CreateOrderRequest struct {
	OrderType string `json:"orderType" validator:"notblank oneof=cafe vip charge"`
	ProductId string `json:"productId" validator:"notblank"`
	UserId    string `json:"userId" validator:"notblank"`
	Quantity  int    `json:"quantity" validator:"gt=0"`
}

type CreateOrderResponse struct {
	BaseResponse
	TrxNo string `json:"trxNo"`
}

type OrderDoneReq struct {
	OrderId string `json:"orderId" validator="notblank"`
	Status  int    `json:"status" validator="gt=0"`
}

type PayReq struct {
	OrderId string  `json:"orderId" validator="notblank"`
	Amount  float64 `json:"amount" validator="gt=0"`
}

type WalletDescribeResp struct {
	BaseResponse
	Balance float64 `json:"balance"`
}

type ChargeReq struct {
	Amount float64 `json:"amount" validator:"gt=0"`
}