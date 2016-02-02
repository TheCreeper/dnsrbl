// Package dnsbl performs rbl lookups to either the builtin or specified
// servers, This package attempts to be as lightweight as possible and provide
// basic capabilities to query dnsrbl servers.
package dnsbl

import (
	"fmt"
	"net"
)

// Map of built in dns block lists that we can query if no dnsbl is specified.
var Servers = map[string]bool{
	"0spam.fusionzero.com": true,
	"zen.spamhaus.org":     true,
}

type Result struct {
	// Listed indicates that the host or IP is on the list.
	Listed bool

	// A dnsrbl will sometimes return a TXT record containing
	// information about why the host or IP is on the list.
	Text string
}

// Query will query a single dnsbl server about a host and return a struct
// containing information returned by the dnsbl server.
func Query(rbl, host string) (r Result, err error) {
	// Format the host address in the format of google.com.bad.example.com or
	// 99.2.0.192.bad.example.com
	name := fmt.Sprintf("%s.%s", host, rbl)

	addrs, err := net.LookupHost(name)
	if err != nil {
		return
	}

	// Check if the dnsbl server replied with an A record that points to
	// a local address such as 127.0.0.0/8 or [::1].
	if len(addrs) == 0 {
		return
	}

	// Check if the address returned by the dnsbl is an address that
	// has 127 as the first octet. If this is true then we can safely
	// assume that the host or IP is on the dnsrbl.
	if net.ParseIP(addrs[0]).IsLoopback() {
		r.Listed = true
	}

	// A dnsbl may return a txt record explaining why the host or IP
	// is on the list.
	txts, err := net.LookupTXT(name)
	if err != nil {
		return
	}

	// Usually only the first txt first is populated so we don't need to
	// loop through them.
	if len(txts) != 0 {
		r.Text = txts[0]
	}
	return
}

// QueryBultin will query all dnsbk servers that are specified in the
// Servers map.
func QueryBultin(host string) (results []Result, err error) {
	for rbl, _ := range Servers {
		r, err := Query(rbl, host)
		if err != nil {
			return nil, err
		}
		results = append(results, r)
	}
	return
}
