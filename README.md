# bitstamp trading stats

## what

* A small Go script
* that listens to [Bitstamp's Websocket v2 API](https://www.bitstamp.net/websocket/v2/),
* counts `order_created` and `trade` events for all currency pairs
* and pushes those counts into a [Statsd+Graphite stack](https://hub.docker.com/r/graphiteapp/docker-graphite-statsd)
* all tied together with Docker Compose.

## how

See `Makefile`, it's quite simple.

## resources

* https://www.digitalocean.com/community/tutorials/how-to-configure-statsd-to-collect-arbitrary-stats-for-graphite-on-ubuntu-14-04
* https://www.digitalocean.com/community/tutorials/how-to-install-and-use-graphite-on-an-ubuntu-14-04-server