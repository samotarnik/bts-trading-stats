{
  "graphiteHost": "127.0.0.1",
  "graphitePort": 2003,
  "port": 8125,
  "flushInterval": 1000,  // overwrite to 1 second (instead of 10)
  "servers": [
    { server: "./servers/udp", address: "0.0.0.0", port: 8125 }
  ]
}
