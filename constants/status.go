package constants

// 订单状态
const (
	STATUS_ORDER_INIT = iota
	STATUS_ORDER_PAYED
	STATUS_ORDER_NOT_PAY
	STATUS_ORDER_DONE
)

// 异步队列相关常量
const (
	ORDER_EXPIRED_MINUTES = 10
	ORDER_Q_NAME          = "order"
	Q_ORDER_EXPIRED       = "asynq:order:expired"
)

// 返回错误码
const (
	ERR_DB = iota + 10086
	ERR_DB_CREATE

	ERR_ORDER_CREATE

	ERR_PAY_NOT_ALLOWED
	ERR_PAY_NOT_PAY
	ERR_PAY_PAYED
	ERR_PAY_BALANCE_Insufficient
)
