package trading

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	tradingdb "github.com/zhs007/tradingdb/trading"
)

var strQueryCandles = `query GetCandles($code: String!, $name: String!, $startTime: Timestamp!, $endTime: Timestamp!) {
	candleChunks(code: $code, name: $name, startTime: $startTime, endTime: $endTime) {
	  startTime,
	  endTime,
	  candles {
		curTime,
		open,
		close,
		high,
		low,
		volume
	  }
	}
  }`

// GetCandles - get candles from tradingdb
//			starttime is like '2006-01-02'
//			endtime is like '2006-01-02'
//			range is [starttime, endtime]
//			tz is timezone
func GetCandles(ctx context.Context, code string, name string, starttime string, endtime string, loc *time.Location) (*tradingdb.ResultCandleChunks, error) {
	checkInitMod()

	// loc, err := time.LoadLocation(tz)
	// if err != nil {
	// 	return nil, err
	// }

	stm, err := time.ParseInLocation("2006-01-02", starttime, loc)
	if err != nil {
		return nil, err
	}

	etm, err := time.ParseInLocation("2006-01-02", endtime, loc)
	if err != nil {
		return nil, err
	}

	cc := make(map[string]interface{})

	cc["code"] = code
	cc["name"] = name
	cc["startTime"] = stm.Unix()
	cc["endTime"] = etm.Unix() + 60*60*24

	buf, err := json.Marshal(cc)
	if err != nil {
		return nil, err
	}

	queryReply, err := singleClient.Query(ctx, strQueryCandles, string(buf))
	if err != nil {
		return nil, err
	}

	fmt.Print(queryReply.Result)

	rcc := tradingdb.ResultCandleChunks{}
	err = json.Unmarshal([]byte(queryReply.Result), &rcc)
	if err != nil {
		return nil, err
	}

	return &rcc, nil

	// var lst []([]interface{})

	// for _, v := range rcc.Data.CandleChunks.Candles {
	// 	var cd []interface{}

	// 	cd = append(cd, time.Unix(v.CurTime, 0).In(loc).Format("2006-01-02 15:04:05"))
	// 	cd = append(cd, v.Open/100.0)
	// 	cd = append(cd, v.Close/100.0)
	// 	cd = append(cd, v.Low/100.0)
	// 	cd = append(cd, v.High/100.0)
	// 	cd = append(cd, v.Volume/100.0)

	// 	lst = append(lst, cd)
	// }

	// retstr, err := json.Marshal(lst)
	// if err != nil {
	// 	return "", nil
	// }

	// return string(retstr), nil
}

// FormatCandles2Arr - format to [[time, open, close, low, high, volume], ...]
func FormatCandles2Arr(loc *time.Location, rcc *tradingdb.ResultCandleChunks) (string, error) {
	var lst []([]interface{})

	for _, v := range rcc.Data.CandleChunks.Candles {
		var cd []interface{}

		cd = append(cd, time.Unix(v.CurTime, 0).In(loc).Format("2006-01-02 15:04:05"))
		cd = append(cd, v.Open/100.0)
		cd = append(cd, v.Close/100.0)
		cd = append(cd, v.Low/100.0)
		cd = append(cd, v.High/100.0)
		cd = append(cd, v.Volume/100.0)

		lst = append(lst, cd)
	}

	retstr, err := json.Marshal(lst)
	if err != nil {
		return "", err
	}

	return string(retstr), nil
}
