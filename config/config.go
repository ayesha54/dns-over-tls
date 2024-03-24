package config

import "time"

type Config struct {
	UpStreamResolverIp   string
	UpStreamResolverPort string
	TCPPort              string
	UPDPort              string
	Net                  string // Add Net field to specify network type (udp/tcp)
	Port                 string // Add Port field to specify the port number
	UpstreamTimeout      time.Duration
}
