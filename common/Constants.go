package common

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	DEFAULTTIMELAYOUT = "2006-01-02 15:04:05"

	SESSIONKEY  = "logininfo"
	SESSIONUSER = "userinfo"
	SESSIONMAX  = 36000

	SCHEDULER_SUCCESS = "success"

	CLUESTS_DROP  = "F类:放入公海"
	CLUESTS_ORDER = "V类:已生成订单"

	BUS_ORDER_STS       = "新订单,已分配制单,制单完成,已分配外勤,执照已拿取,订单完成待交接,已交接,已转为代账客户"
	CHANGE_ORDER_STS    = "新订单,已分配制单,制单完成,已分配外勤,执照已拿取,订单完成待交接,已交接"
	CANCELBUS_ORDER_STS = "新订单,已分配制单,公示公告,已分配外勤,注销单已取,已分配银行,订单完成待交接,已交接"
	CANCELBUS_TAX_STS   = "新订单,已分配外勤,已完成"
)

func GetNextCacncelBusOrderSts(orderstep string) (nextstep string) {
	step := strings.Split(CANCELBUS_ORDER_STS, ",")
	for i := 0; i < len(step); i++ {
		if step[i] == orderstep {
			if i+1 <= len(step) {
				nextstep = step[i+1]
				return
			}
		}
	}
	return
}

func GetNextChangeOrderSts(orderstep string) (nextstep string) {
	step := strings.Split(CHANGE_ORDER_STS, ",")
	for i := 0; i < len(step); i++ {
		if step[i] == orderstep {
			if i+1 <= len(step) {
				nextstep = step[i+1]
				return
			}
		}
	}
	return
}

func GetNextBussOrderSts(orderstep string) (nextstep string) {
	step := strings.Split(BUS_ORDER_STS, ",")
	for i := 0; i < len(step); i++ {
		if step[i] == orderstep {
			if i+1 <= len(step) {
				nextstep = step[i+1]
				return
			}
		}
	}
	return
}

func GetCANCELBUSTAXSTS(orderstep string) (nextstep string) {
	step := strings.Split(CANCELBUS_TAX_STS, ",")
	for i := 0; i < len(step); i++ {
		if step[i] == orderstep {
			if i+1 <= len(step) {
				nextstep = step[i+1]
				return
			}
		}
	}
	return
}

func Float2Float(num float32) float64 {
	float_num, _ := strconv.ParseFloat(fmt.Sprintf("%.8f", num), 32)
	return float_num
}
