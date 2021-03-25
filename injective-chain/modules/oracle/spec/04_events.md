# Events

The stdreference module emits the following events:

## Types

```go
const (
    // Event types
    EventTypeSetRef = "set_ref"
    EventTypeGetReferenceData = "get_reference_data"

    // Attribute types
    // Msg/Relay
    AttributeKeyRelayer = "relayer"
    AttributeKeySymbol = "symbol"
    AttributeKeyRate = "rate"
    AttributeKeyResolveTime = "resolve_time"
    AttributeKeyRequestID = "request_id"
    // function/GetReferenceData
    AttributeKeyBaseQuoteSymbol = "base_quote_symbol"
    AttributeKeyBaseQuoteRate = "base_quote_rate"
    AttributeKeyLastUpdateBase = "last_update_base"
    AttributeKeyLastUpdateQuote = "last_update_quote"
)
```

## Service Messages

### Msg/Relay

| Attribute Key           | Attribute Value |
| ----------------------- | --------------- |
| AttributeKeyRelayer     | string          |
| AttributeKeySymbol      | string          |
| AttributeKeyRate        | string          |
| AttributeKeyResolveTime | string          |
| AttributeKeyRequestID   | string          |

### function/GetReferenceData

| Attribute Key               | Attribute Value |
| --------------------------- | --------------- |
| AttributeKeyBaseQuoteSymbol | string          |
| AttributeKeyBaseQuoteRate   | string          |
| AttributeKeyLastUpdateBase  | string          |
| AttributeKeyLastUpdateQuote | string          |
