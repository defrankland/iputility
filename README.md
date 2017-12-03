# iputility
IP utilities for go - including type for representing and comparing ip/range/CIDR/FQDN

# Type Ip API
Used for identifying whether an endpoint (represented as a string) is:
  - IP address such as `192.111.65.3`
  - IP address range such as `192.168.0.0-192.200.255.255`
  - CIDR such as `192.168.55.33/8`
  - FQDN such as `somefqdn.com`

```go 
type Ip struct {
	Type     IpId
	Endpoint string
	loUint   uint64
	hiUint   uint64
}
```

`GetType(endpoint string) Ip` returns an object of type `Ip` that designates which of the above types the input is. If none of the above, then the type is set to undefined.  

`(t *Ip) GetFirst() string` returns the first endpoint in the range/CIDR. If t is an IP address, it returns the IP address. 

`(t *Ip) In(t1 Ip) bool` returns `true` when the passed t1 Ip falls inside of t. If they are exactly equal, then it returns `false`. It works for any possible combination of IP, IP range, and CIDR. *Note:* Clearly the FQDN type doesn't seem to fit here, but I was going to tie this into an `nslookup` to resolve the FQDN. So, I'm only temporarily violating the single responsibility principle :). 

`(t *Ip) Equals(t1 Ip) bool` returns `true` when t and t1 are exactly equal and works for any possible combination of IP, IP range, and CIDR.


# Other Stuff
The commands.go file has `ping()` and `nslookup()` functions that seem to work on both Mac and Linux. The ping is timed to 3 seconds and will take that long to complete every time. 
