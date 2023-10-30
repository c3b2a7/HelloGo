package main

import (
	"context"
	"fmt"
	"github.com/miekg/dns"
	"net"
	"os"
	"testing"
)

var (
	domain     = "www.ebay.com.hk"
	nameserver = "114.114.114.114"
)

func Test_DNSQuery(t *testing.T) {
	os.Setenv("GODEBUG", "netdns=2")
	resolver := net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.Dial("udp", net.JoinHostPort(nameserver, "53"))
		},
	}
	host, err := resolver.LookupHost(context.Background(), domain)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(host)
}

func TestDNS(t *testing.T) {
	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	var err error
	msg, err = dns.Exchange(msg, net.JoinHostPort(nameserver, "53"))
	if err != nil {
		t.Error(err)
	}
	println(msg.String())
}
