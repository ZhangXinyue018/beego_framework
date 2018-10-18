package impl

import (
	"beego_framework/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type ExchangerRpc struct {
	BitstampUrl   string
	CoinMarketUrl string
	BitfinexUrl   string
}

type Exchanger string

const (
	BITSTAMP    Exchanger = "Bitstamp"
	COIN_MARKET Exchanger = "CoinMarketCap"
	BITFINEX    Exchanger = "Bitfinex"
)

var EXCHANGER_LIST = []Exchanger{COIN_MARKET}

func (rpc *ExchangerRpc) ListPrice(symbolNameList *[]string) (*[]domain.SimpleExchangerRate) {
	var resultList []domain.SimpleExchangerRate
ExchangerScan:
	for _, exchanger := range EXCHANGER_LIST {
		if exchanger == BITSTAMP {
			resultList = append(resultList, *rpc.getBitstampPrice(symbolNameList)...)
		} else if exchanger == COIN_MARKET {
			resultList = append(resultList, *rpc.getCoinMarketPrice(symbolNameList)...)
		} else if exchanger == BITFINEX {
			resultList = append(resultList, *rpc.getBitfinexPrice(symbolNameList)...)
		} else {
			continue ExchangerScan
		}
	}
	return &resultList
}

type BitstampObj struct {
	Last string `json:"last"`
}

func (rpc *ExchangerRpc) getBitstampPrice(symbolNameList *[]string) (*[]domain.SimpleExchangerRate) {
	var resultList []domain.SimpleExchangerRate
	for _, symbolName := range *symbolNameList {
		url := rpc.BitstampUrl + "/v2/ticker/" + symbolName + "usd"
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if response.StatusCode == 200 {
			r, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
				continue
			}

			priceResponse := &BitstampObj{}
			err = json.Unmarshal([]byte(string(r)), priceResponse)
			if err != nil {
				fmt.Println(err)
				continue
			}
			price, err := strconv.ParseFloat(priceResponse.Last, 64)
			if err != nil {
				fmt.Println(err)
				continue
			}
			result := domain.SimpleExchangerRate{
				Symbol:    symbolName,
				Exchanger: string(BITSTAMP),
				Price:     price,
			}
			resultList = append(resultList, result)
		} else {
			fmt.Println("error request bitstamp")
			continue
		}
	}
	return &resultList
}

type CoinMarketObj struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Symbol              string `json:"symbol"`
	PriceUsd            string `json:"price_usd"`
	PercentageChange24H string `json:"percent_change_24h"`
}

func (rpc *ExchangerRpc) getCoinMarketPrice(symbolNameList *[]string) (*[]domain.SimpleExchangerRate) {
	var resultList []domain.SimpleExchangerRate
	for _, symbolName := range *symbolNameList {
		result, err := rpc.getSingleCoinMarketPrice(symbolName)
		if err != nil || result == nil {
			continue
		}
		resultList = append(resultList, *result)
	}
	return &resultList
}

func (rpc *ExchangerRpc) getSingleCoinMarketPrice(symbolName string) (*domain.SimpleExchangerRate, error) {
	var err error
	defer func() {
		if x := recover(); x != nil {
			err = x.(error)
		}
	}()
	url := rpc.CoinMarketUrl + "/v1/ticker/" + symbolName
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if response.StatusCode == 200 {
		r, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		listPriceResponse := &([]CoinMarketObj{})
		err = json.Unmarshal([]byte(string(r)), listPriceResponse)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		price, err := strconv.ParseFloat((*listPriceResponse)[0].PriceUsd, 64)
		percentageChange24h, err := strconv.ParseFloat((*listPriceResponse)[0].PercentageChange24H, 64)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		result := domain.SimpleExchangerRate{
			Symbol:              symbolName,
			Exchanger:           string(COIN_MARKET),
			Price:               price,
			PercentageChange24H: percentageChange24h,
		}
		return &result, nil
	} else {
		fmt.Println("error request coinmarket")
		return nil, err
	}
	return nil, err
}

func (rpc *ExchangerRpc) getBitfinexPrice(symbolNameList *[]string) (*[]domain.SimpleExchangerRate) {
	var resultList []domain.SimpleExchangerRate
	var formatSymbolList []string
	formatSymbolMap := make(map[string]string)
	for _, symbolName := range *symbolNameList {
		formatSymbol := "t" + strings.ToUpper(symbolName) + "USD"
		formatSymbolList = append(formatSymbolList, formatSymbol)
		formatSymbolMap[formatSymbol] = symbolName
	}

	url := rpc.BitfinexUrl + "/v2/tickers?symbols=" + strings.Join(formatSymbolList, ",")
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return &([]domain.SimpleExchangerRate{})
	}
	if response.StatusCode == 200 {
		r, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			return &([]domain.SimpleExchangerRate{})
		}

		listPriceResponse := &([][]interface{}{})
		err = json.Unmarshal([]byte(string(r)), listPriceResponse)
		if err != nil {
			fmt.Println(err)
			return &([]domain.SimpleExchangerRate{})
		}
		for _, priceList := range *listPriceResponse {
			price := priceList[7].(float64)
			name := priceList[0].(string)
			resultList = append(resultList, domain.SimpleExchangerRate{
				Symbol:    formatSymbolMap[name],
				Exchanger: string(BITFINEX),
				Price:     price,
			})
		}
		return &resultList
	} else {
		fmt.Println("error request bitfinex")
		return &([]domain.SimpleExchangerRate{})
	}
}
