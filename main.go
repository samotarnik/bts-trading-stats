package main

import (
	"log"

	"github.com/samotarnik/bitstamp-go"
	"github.com/smira/go-statsd"
)

const statsdPrefix = "bts." // important: used in storage-schemas.conf

var channelsToStats = map[string]string{
	"live_orders_bchbtc": "orders.bchbtc",
	"live_orders_bcheur": "orders.bcheur",
	"live_orders_bchusd": "orders.bchusd",
	"live_orders_btceur": "orders.btceur",
	"live_orders_btcusd": "orders.btcusd",
	"live_orders_ethbtc": "orders.ethbtc",
	"live_orders_etheur": "orders.etheur",
	"live_orders_ethusd": "orders.ethusd",
	"live_orders_eurusd": "orders.eurusd",
	"live_orders_ltcbtc": "orders.ltcbtc",
	"live_orders_ltceur": "orders.ltceur",
	"live_orders_ltcusd": "orders.ltcusd",
	"live_orders_xrpbtc": "orders.xrpbtc",
	"live_orders_xrpeur": "orders.xrpeur",
	"live_orders_xrpusd": "orders.xrpusd",
	"live_trades_bchbtc": "trades.bchbtc",
	"live_trades_bcheur": "trades.bcheur",
	"live_trades_bchusd": "trades.bchusd",
	"live_trades_btceur": "trades.btceur",
	"live_trades_btcusd": "trades.btcusd",
	"live_trades_ethbtc": "trades.ethbtc",
	"live_trades_etheur": "trades.etheur",
	"live_trades_ethusd": "trades.ethusd",
	"live_trades_eurusd": "trades.eurusd",
	"live_trades_ltcbtc": "trades.ltcbtc",
	"live_trades_ltceur": "trades.ltceur",
	"live_trades_ltcusd": "trades.ltcusd",
	"live_trades_xrpbtc": "trades.xrpbtc",
	"live_trades_xrpeur": "trades.xrpeur",
	"live_trades_xrpusd": "trades.xrpusd",
}

func main() {
	// all channels
	wsChannels := make([]string, 0)
	for channel := range channelsToStats {
		wsChannels = append(wsChannels, channel)
	}

	// init statsd client
	sd := statsd.NewClient("127.0.0.1:8125", statsd.MetricPrefix(statsdPrefix))

	// init ws client
	ws, err := bitstamp.NewWsClient()
	if err != nil {
		log.Panic(err)
	}

	ws.Subscribe(wsChannels...)

	for {
		select {
		case e := <-ws.Stream:
			if e.Event == "order_created" {
				sd.Incr(channelsToStats[e.Channel], 1)
				sd.Incr("orders.all", 1)
			} else if e.Event == "trade" {
				sd.Incr(channelsToStats[e.Channel], 1)
				sd.Incr("trades.all", 1)
			}
		case err := <-ws.Errors:
			// crash in case of an error and let runit restart the program
			log.Fatal(err)
		}
	}
}
