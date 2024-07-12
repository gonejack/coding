package main

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
	"golang.org/x/net/proxy"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println(check())
}

func check() (err error) {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1086", nil, nil)
	if err != nil {
		return fmt.Errorf("解析出错: %s", err)
	}

	//client := dns.Client{
	//	Net: "tcp",
	//	Dialer: &net.Dialer{
	//		Resolver: &net.Resolver{Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
	//			return dialer.Dial(network, address)
	//		}},
	//	},
	//}

	conn, err := dialer.Dial("tcp", "1.0.0.1:53")
	if err != nil {
		return fmt.Errorf("conn failed: %s", err)
	}

	var msg dns.Msg
	{
		msg.SetQuestion(dns.Fqdn("inoreader.com"), dns.TypeA)
	}

	var client = dns.Client{
		Net:     "udp",
		UDPSize: 4196,
	}
	answer, _, err := client.ExchangeWithConn(&msg, &dns.Conn{
		Conn:         conn,
		UDPSize:      client.UDPSize,
		TsigSecret:   client.TsigSecret,
		TsigProvider: client.TsigProvider,
	})
	//answer, _, err := client.Exchange(&msg, "1.0.0.1:53")
	if err != nil {
		return fmt.Errorf("回答出错: %s", err)
	}
	for _, ans := range answer.Answer {
		fmt.Println(ans.String())
		log.Println(ans.(*dns.A).A)
	}
	return
}
