# State

The `x/stdreference` module keeps state of one object, the Ref. Ref is a raw data from BandChain, which has a requestID inside it. The requestID will allow anyone to check that the price data actually exists in the BandChain. The state is basically a mapping from string(symbol) to a struct Ref.

```go
type Ref struct {
    Rate        uint64
    ResolveTime uint64 // unix timestamp
    RequestID   uint64
}
```

- Ref: prefix + symbol -> Ref struct
