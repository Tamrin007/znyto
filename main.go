package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var rate float64

type btcJPY struct {
	Last   float64 `json:"last"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Vwap   float64 `json:"vwap"`
	Volume float64 `json:"volume"`
	Bid    float64 `json:"bid"`
	Ask    float64 `json:"ask"`
}

type znyBTC struct {
	Ticker struct {
		High       float64 `json:"high"`
		Low        float64 `json:"low"`
		Avg        float64 `json:"avg"`
		Lastbuy    float64 `json:"lastbuy"`
		Lastsell   float64 `json:"lastsell"`
		Buy        float64 `json:"buy"`
		Sell       float64 `json:"sell"`
		Lastprice  float64 `json:"lastprice"`
		Buysupport float64 `json:"buysupport"`
		Updated    int     `json:"updated"`
	} `json:"ticker"`
}

func getRate() {
	// fetch zny-btc
	resp, err := http.Get("https://c-cex.com/t/zny-btc.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	znyBTC := znyBTC{}
	err = json.Unmarshal(body, &znyBTC)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fetch btc_jpy
	resp, err = http.Get("https://api.zaif.jp/api/1/ticker/btc_jpy")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	btcJPY := btcJPY{}
	err = json.Unmarshal(body, &btcJPY)
	if err != nil {
		fmt.Println(err)
		return
	}

	rate = znyBTC.Ticker.Avg * btcJPY.Bid
}

func startPolling() {
	for {
		time.Sleep(1 * time.Second)
		go getRate()
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.WriteHeader(200)
	s := strconv.FormatFloat(rate, 'f', 10, 64)
	w.Write([]byte(s))
}

func main() {
	go startPolling()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
