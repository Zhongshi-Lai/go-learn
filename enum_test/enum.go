package main

// 门店购订单状态
type SGOrderStatus int

const (
	_                        SGOrderStatus = iota
	SGOrderStatusWaitPay                   // 待支付
	SGOrderStatusWaitSend                  // 待发货
	SGOrderStatusWaitReceive               // 待收货
	SGOrderStatusCancel                    // 已取消
	SGOrderStatusDone                      // 交易完成
)

func main() {
	// a := -1

	// var b SGOrderStatus
	// fmt.Println(b)
	// b =a

}
