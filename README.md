# Are You Up?
A minimal AVS to determine probabilistically whether an endpoint is up/reachable and confirm on-chain.

<img src="assets/kramer-up.jpeg" alt="kramer-up" width="300"/>

## Motivations
- Build an AVS that can be reusable by other AVSs.
- Educate & demonstrate the bare minimum capabilities needed to participate as an AVS in the EigenLayer ecosystem. Inspired by [Scaffold-ETH](https://github.com/scaffold-eth/scaffold-eth-2).

## Example Use Case
- An AVS service is current operating an "active set" of node operators. The service would like to keep a record on chain of which of those operators are truly responsive to web requests (alive and healthy) vs non-responsive.


## Architecture

### Process flow

- User submits request for an endpoint to be tested for liveliness.
- The network of AYU operators all test the liveliness of that endpoint and report their results.
- Record of the endpoint and its liveliness (via ping, traceroute, etc.).
- Incentives:
  - Slash AYU operators that do not report their liveliness results for a given request.
  - Pay AYU operators that report results for a given request.


### Operator Image
Minimal Docker image running a minimal go binary that pings (via traceroute) a server, URL, or IP to confirm it is "up" (alive, available).

### Roadmap
Enhanced cryptoeconomic incentives to make cheating non-profitable.
More robust logging (traceroute) of each request to make cheating more difficult.


# Usage

```
go run main.go <ip-address-or-domain>

#or
go build
./are-you-up http://google.com
```

# Todos
- Add AVS registration smart contracts
- Determine the best initial invocation process: off chain vs on chain?
- Determine the best place for users to write results: on chain (L1, L2), EigenDA, IPFS?
- Add https://github.com/urfave/cli/
- Add dockerfile - FROM golang:latest


