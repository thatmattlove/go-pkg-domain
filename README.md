This repository contains a Go [Cloudflare Worker](https://developers.cloudflare.com/workers/) (using the [`workers`](https://github.com/syumai/workers) package by @syumai) that handles HTTP requests for the `go.mdl.wtf` domain for Go packaging. [See here](https://go.dev/ref/mod#goproxy-protocol) for details.

In a nutshell, the worker in this repository enables this to work:

```go
import (
    // This ↓
    "go.mdl.wtf/go-macaddr"
    // Instead of ↓
    "github.com/thatmattlove/go-macaddr"
)
```

![GitHub License](https://img.shields.io/github/license/thatmattlove/go-pkg-domain)
