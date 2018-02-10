package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"melon_micro/proto/currency"
	"net/http"

	"errors"

	"github.com/micro/go-micro"
)

// Currency 貨幣
type Currency struct{}

// Ping 測試連線
func (c *Currency) Ping(ctx context.Context, req *currency.PingRequest, res *currency.PongResponse) (err error) {
	res.ServiceName = "currency"
	return
}

// List 貨幣列表
func (c *Currency) List(ctx context.Context, req *currency.CurrencyListRequest, rsp *currency.CurrencyListResponse) error {
	rsp.List = map[string]string{
		"USD": "美金",
		"TWD": "台幣",
		"JPY": "日幣",
	}
	return nil
}

// Exchange 匯率
func (c *Currency) Exchange(ctx context.Context, req *currency.CurrencyRequest, rsp *currency.CurrencyResponse) error {
	currencyConvtKey := req.From + req.To

	url := "https://tw.rter.info/capi.php"

	curl, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(curl)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	type ExchangeInfo struct {
		Exchange float64 `json:"Exrate"`
		Update   string  `json:"UTC"`
	}
	exchangeInfo := map[string]ExchangeInfo{}
	json.Unmarshal(body, &exchangeInfo)
	exchangeData, isOK := exchangeInfo[currencyConvtKey]
	if isOK {
		rsp.Exchange = exchangeData.Exchange
		rsp.UpdateTime = exchangeData.Update
	} else {
		return errors.New("No Data")
	}

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("currency"),
		micro.Version("latest"),
		micro.BeforeStart(func() error {
			log.Println("🐳  Currency Service Start 🐳")
			return nil
		}),
		micro.AfterStop(func() error {
			log.Println("🔥  Currency Service Stop 🔥")
			return nil
		}),
	)

	service.Init()
	currency.RegisterCurrencyHandler(service.Server(), new(Currency))

	if err := service.Run(); err != nil {
		log.Printf("🎃  Currency Service return an error : %v 🎃", err)
	}
}
