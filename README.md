# dnsbl

dnsbl is a library written to help you check if an address is listed on with
a dnsrbl. You can either specify a dnsrbl to query or use the builtin list.
This package attempts to be as lightweight as possible and provide basic
capabilities to query dnsrbl servers.


## Examples

Query a single dnsrbl server.
```Go
r, err := dnsrbl.Query("zen.spamhaus.org", "google.com")
if err != nil {
	log.Fatal(err)
}

if r.Listed {
	// Do something.
}
```

Query the builtin list of dnsrbl servers.
```Go
results, err := dnsrbl.QueryBuiltin("google.com")
if err != nil {
	log.Fatal(err)
}

for _, result := range results {
	if result.Listed {
		// Do something.
	}
}
```
