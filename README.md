# Businnect Go SDK

A professional Go client for the businnect API, featuring high-performance support for **Microposts**.

## 📖 API Documentation

For detailed information on the underlying REST API, endpoints, and authentication schemas, please visit the official documentation:
* **API Reference**: [http://api.businnect.com](https://api.businnect.com/docs/client/)

## Features

* **Micropost System**: Create text posts or file-based updates with community-specific targeting.
* **Context Support**: Full `context.Context` support for timeouts and cancellations in every request.

## Installation

```bash
go get github.com/BusinnectCode/businnect_go

```

## Quick Start

### Initialize the Client

The client can be initialized with an explicit API key or automatically look for environment variables.

```go
import "github.com/BusinnectCode/businnect_go"

func main() {
    // Pass the API key and Base URL
    client := businnect.NewClient("your-api-token", "https://api.businnect.com")
}

```

### Create a Micropost

Easily create posts with optional titles and bodies.

```go
title := "New Update"
req := resources.CreateMicropostRequest{
    Title: &title,
    Body:  "The Go SDK is now live!",
}

resp, err := client.Micropost.Create(context.Background(), req)
if err != nil {
    panic(err)
}
fmt.Printf("Post created with ID: %s\n", resp.PublicID)

```

## Environment Variables

The SDK can utilize the following environment variables for testing or default configuration:

* `BUSINNECT_API_TOKEN`: Used by integration tests to authenticate against the live API.

## Testing

Run the test suite using the standard Go toolchain:

```powershell
# In PowerShell
$env:BUSINNECT_API_TOKEN="your-token"; go test -v ./...

```

```bash
# In Bash
BUSINNECT_API_TOKEN="your-token" go test -v ./...

```