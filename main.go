package main

import (
	"log"
	"os"
	"time"

	"github.com/ayesha54/dns-over-tls/config"
	"github.com/ayesha54/dns-over-tls/handler"
	"github.com/miekg/dns"
	"github.com/urfave/cli/v2"
)

var (
	App = &cli.App{
		Name:  "DNS-over-TLS Proxy over Cloudflare upstream",
		Usage: "A DNS-over-TLS proxy server using Cloudflare as an upstream resolver",
	}

	// Default configuration values
	defaultConfig = config.Config{
		UpStreamResolverIp:   "1.1.1.1",
		UpStreamResolverPort: "853",
		TCPPort:              ":53",
		UPDPort:              ":53",
		UpstreamTimeout:      time.Millisecond * 3000,
		Net:                  "udp", // Specify default network type
		Port:                 ":53", // Specify default port
	}
)

func main() {
	// Create a new instance of the App
	App.Commands = []*cli.Command{
		{
			Name:  "udp",
			Usage: "Run the UDP/53 server",
			Action: func(c *cli.Context) error {
				conf := defaultConfig
				conf.Net = "udp"
				return startServer(conf)
			},
		},
		{
			Name:  "tcp",
			Usage: "Run the TCP/53 server",
			Action: func(c *cli.Context) error {
				conf := defaultConfig
				conf.Net = "tcp"
				return startServer(conf)
			},
		},
	}

	// Run the CLI application
	if err := App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func startServer(conf config.Config) error {
	// Create a new DNS server with the provided configuration
	server := dns.Server{Addr: conf.Port, Net: conf.Net}

	// Register DNS handler
	dns.HandleFunc(".", handler.DNSHandler(conf))

	// Start DNS server
	log.Printf("DNS server is running on %s://%s", conf.Net, server.Addr)
	if conf.Net == "tcp" {
		log.Printf("Try in CLI: dig +short +tcp google.com @localhost")
	} else {
		log.Printf("Try in CLI: dig +short google.com @localhost")
	}

	return server.ListenAndServe()
}
