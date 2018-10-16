package trading

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"time"

	tradingdb "github.com/zhs007/tradingdb/trading"
)

var strQueryTradingData = `query GetTradingData($name: ID!) {
	tradingData(name: $name) {
	  keyID
	  orders {
		orderID
		orderType
		orderSide
		price
		volume
		startTime
		avgPrice
		doneVolume
		doneTime
	  }
	  trades {
		tradeID
		orderID
		curTime
		price
		volume
	  }
	}
  }`

// GetTradingData - get tradingdata from tradingdb
//			name is tradingdata name
//			tz is timezone
func GetTradingData(ctx context.Context, name string) (*tradingdb.ResultTradingData, error) {
	checkInitMod()

	// loc, err := time.LoadLocation(tz)
	// if err != nil {
	// 	return nil, err
	// }

	cc := make(map[string]interface{})

	cc["name"] = name

	buf, err := json.Marshal(cc)
	if err != nil {
		return nil, err
	}

	queryReply, err := singleClient.Query(ctx, strQueryTradingData, string(buf))
	if err != nil {
		return nil, err
	}

	fmt.Print(queryReply.Result)

	rtd := tradingdb.ResultTradingData{}
	err = json.Unmarshal([]byte(queryReply.Result), &rtd)
	if err != nil {
		return nil, err
	}

	return &rtd, nil

	// var lst []([]interface{})

	// for _, v := range rcc.Data.TradingData.Orders {
	// 	var cd []interface{}

	// 	cd = append(cd, v.OrderID)
	// 	cd = append(cd, v.OrderType)
	// 	cd = append(cd, v.OrderSide)
	// 	cd = append(cd, time.Unix(v.StartTime, 0).In(loc).Format("2006-01-02 15:04:05"))
	// 	cd = append(cd, v.Price/100.0)
	// 	cd = append(cd, v.Volume/100.0)
	// 	cd = append(cd, time.Unix(v.DoneTime, 0).In(loc).Format("2006-01-02 15:04:05"))
	// 	cd = append(cd, v.AvgPrice/100.0)
	// 	cd = append(cd, v.DoneVolume/100.0)

	// 	lst = append(lst, cd)
	// }

	// retstr, err := json.Marshal(lst)
	// if err != nil {
	// 	return "", nil
	// }

	// return string(retstr), nil
}

// FormatTradingData2Arr - format to [[id, type, side, newtime, price, volume, tradetime, avgprice, donevolume], ...]
func FormatTradingData2Arr(loc *time.Location, rtd *tradingdb.ResultTradingData) (string, error) {
	var lst []([]interface{})

	for _, v := range rtd.Data.TradingData.Orders {
		var cd []interface{}

		cd = append(cd, v.OrderID)
		cd = append(cd, v.OrderType)
		cd = append(cd, v.OrderSide)
		cd = append(cd, time.Unix(v.StartTime, 0).In(loc).Format("2006-01-02 15:04:05"))
		cd = append(cd, v.Price/100.0)
		cd = append(cd, v.Volume/100.0)
		cd = append(cd, time.Unix(v.DoneTime, 0).In(loc).Format("2006-01-02 15:04:05"))
		cd = append(cd, v.AvgPrice/100.0)
		cd = append(cd, v.DoneVolume/100.0)

		lst = append(lst, cd)
	}

	retstr, err := json.Marshal(lst)
	if err != nil {
		return "", err
	}

	return string(retstr), nil
}

func countAvgPrice(cm float64, sv float64, sp float64, cv float64, cp float64) (curmoney float64, curvolume float64, avgprice float64) {
	if sv == 0 {
		return cm, cv, cp
	}

	if cv == 0 {
		return cm, sv, sp
	}

	if sv > 0 {
		if cv > 0 {
			tt := sv*sp + cv*cp

			return cm, sv + cv, tt / (sv + cv)
		}

		if sv > math.Abs(cv) {
			cw := math.Abs(cv) * (cp - sp)
			return cm + cw, sv - math.Abs(cv), sp
		} else if sv < math.Abs(cv) {
			cw := sv * (cp - sp)
			return cm + cw, -(math.Abs(cv) - sv), cp
		}

		cw := sv * (cp - sp)
		return cm + cw, 0, 0
	}

	if cv < 0 {
		tt := math.Abs(sv)*sp + math.Abs(cv)*cp

		return cm, sv + cv, tt / math.Abs(sv+cv)
	}

	if math.Abs(sv) > cv {
		cw := cv * (sp - cp)
		return cm + cw, -(math.Abs(sv) - cv), sp
	} else if math.Abs(sv) < cv {
		cw := math.Abs(sv) * (sp - cp)
		return cm + cw, (cv - math.Abs(sv)), cp
	}

	cw := math.Abs(sv) * (sp - cp)
	return cm + cw, 0, 0
}

// FormatTradingData2PNLChart - format to [donetime, ...] & [value, ...]
func FormatTradingData2PNLChart(loc *time.Location, rtd *tradingdb.ResultTradingData) ([]string, []float64) {
	var lsttime []string
	var lstval []float64

	curval := 0.0
	curvol := 0.0
	avgprice := 0.0

	for _, v := range rtd.Data.TradingData.Orders {
		lsttime = append(lsttime, time.Unix(v.DoneTime, 0).In(loc).Format("2006-01-02 15:04:05"))

		if v.DoneVolume > 0 {
			cv := float64(v.DoneVolume) / 100.0
			cp := float64(v.AvgPrice) / 100.0

			if v.OrderSide == "SELL" {
				cv = -cv
			}

			curval, curvol, avgprice = countAvgPrice(curval, curvol, avgprice, cv, cp)

			lstval = append(lstval, curval)
		}
	}

	return lsttime, lstval
}
