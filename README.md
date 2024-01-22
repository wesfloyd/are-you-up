<img src="assets/kramer-up.jpeg" alt="kramer-up" width="300"/>

# Are You Up?
A minimal AVS to determine probabilistically whether an endpoint is up/reachable and confirm on-chain.

## Motivations
- Build an AVS that has a greater than 0% chance to be valuable for reuse by other AVSs.
- Educate & demonstrate the bare minimum capabilities needed to participate as an AVS in the EigenLayer ecosystem. Inspired by [Scaffold-ETH](https://github.com/scaffold-eth/scaffold-eth-2).

## Vision
- Build this AVS to detect and confirm with on chain whether a given host is up (reachable).
- Build incremental services as simple capabilities towards useful decentralized web services:
    - "Can You Redirect?" (CYR) a service to keep track of a global set of servers that can act as minimium viable proxies to forward requests to other servers confirmed via "Are You Up". A mini load balancer.
        - "Can They Redirect?" a global load balancer of CYRs.
        - "Are You Fast?"
        - "Where Are You?"
    - "Can You Serve?" (CYS) a service to keep track of whether a server is currently serving a simple Lambda style function via "Can You Redirect".
    - "Show Me What You Got" a simple log daemon that writes log files to IPFS.
    

## Architecture

### Offchain: 
Minimal Docker image running a minimal go binary that pings (via traceroute) a server, URL, or IP to confirm it is "up" (alive, available).
?Does an AVS design require some "consensus" among the operators about a given fact, or state, or proof?

### Onchain 
Record of the endpoint and its liveliness (is it up)
Slash Operators that do not report their results for a given request.
Pay Operators that report any results for a given request.

### Submitting New Jobs
?Could an AVS end user submit their request to check server liveliness via a dApp? I'd like to avoid centralized components in the arch if possible.
?Does AVS design have any hard requirements on whether the user job flow is initiated on chain or off chain?

### Roadmap
Enhanced cryptoeconomic incentives to make cheating non-profitable.
More robust logging (traceroute) of each request to make cheating more difficult.


# Usage

```
go run main.go https://example.com

#or
go build
./are-you-up http://google.com
```

# todos
Add https://github.com/urfave/cli/
Add dockerfile - FROM golang:latest

# Open Questions


