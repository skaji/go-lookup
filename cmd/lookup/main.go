package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	dnsServer := flag.String("dns-server", "", "dns server")
	retry := flag.Int("retry", 3, "how many times retry dns query")
	domain := flag.String("domain", "", "query domain")
	showAll := flag.Bool("all", false, "show all DNS results")
	ipv4only := flag.Bool("ipv4-only", true, "show IPv4 only")
	timeout := flag.Int("timeout", 500, "dns query timeout (msec)")
	flag.Parse()

	if *dnsServer == "" {
		fmt.Fprintln(os.Stderr, "need -dns-server option")
		os.Exit(1)
	}
	if *domain == "" {
		fmt.Fprintln(os.Stderr, "need -domain option")
		os.Exit(1)
	}

	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "tcp", *dnsServer+":53")
		},
	}

	ips := []string{}
	lastErr := errors.New("unknown error")
	for i := 0; i < *retry; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Millisecond)
		addrs, err := r.LookupIPAddr(ctx, *domain)
		cancel()
		if err != nil {
			lastErr = err
			continue
		}
		if len(addrs) > 0 {
			for _, addr := range addrs {
				if *ipv4only && len(addr.IP) != net.IPv4len {
					continue
				}
				ips = append(ips, addr.IP.String())
			}
			break
		}
	}

	if len(ips) == 0 {
		fmt.Fprintf(os.Stderr, "ERROR: dnsServer: %s:53, domain: %s, details: %+v\n", *dnsServer, *domain, lastErr)
		os.Exit(1)
	}

	if *showAll {
		for _, ip := range ips {
			fmt.Println(ip)
		}
	} else {
		fmt.Println(ips[0])
	}
}
