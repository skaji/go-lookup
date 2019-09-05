# lookup

A simple CLI tool to lookup IPs

# Usage

## Example

```
❯ lookup -dns-server 8.8.8.8 -domain google.com
172.217.26.46
```

## Help

```
❯ lookup -h
Usage of lookup:
  -all
    	show all DNS results
  -dns-server string
    	dns server
  -domain string
    	query domain
  -ipv4-only
    	show IPv4 only (default true)
  -retry int
    	how many times retry dns query (default 3)
  -timeout int
    	dns query timeout (msec) (default 500)
```

# Author

Shoichi Kaji

# License

MIT
