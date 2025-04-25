# XCoin Client

XCoin Client is a Go library for interacting with XCoin Bot API. The library provides a convenient interface for managing trading pairs and executing trading operations.

## Installation

```bash
go get github.com/wrkit/xcoin_client
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "time"
    "github.com/wrkit/experiment/xcoin_client"
)

func main() {
    // Create a new client
    cfg := xcoin_client.NewConfig("http://your-bot-url:port")
    client := xcoin_client.NewClient(cfg)

    // Create a context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
    defer cancel()

    // Get bot version
    version, err := client.GetVersion(ctx)
    if err != nil {
        fmt.Printf("Error getting version: %v\n", err)
        return
    }
    fmt.Printf("Version: %+v\n", version)
}
```

## Main Features

### Client Configuration

```go
// Basic configuration
cfg := xcoin_client.NewConfig("http://bot-url:port")

// Timeout configuration
cfg.WithTimeout(60 * time.Second)

// Using custom HTTP client
customClient := &http.Client{}
cfg.WithHTTPClient(customClient)
```

### Working with Trading Pairs

```go
// Get list of trading pairs
pairs, err := client.GetPairList2(ctx)

// Add new trading pair
err = client.AddPair(ctx, "BTC", "USD", "binance", "mode")

// Get pair settings
settings, err := client.GetPairCurrentSettings(ctx, pairID)

// Get pair orders
orders, err := client.GetPairOrders(ctx, pairID)

// Update pair settings
err = client.UpdatePairSettings(ctx, pairID, newSettings)
```

## Architecture

The library is built on the following principles:

1. Configurable client with support for bot URL and HTTP client settings
2. Non-blocking requests with context support
3. Typed data structures for all API responses

## Error Handling

All methods return an error as a second value. It is recommended to always check for errors:

```go
result, err := client.GetVersion(ctx)
if err != nil {
    // Handle error
    return err
}
```

## License

MIT