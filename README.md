# routek - YAML-based HTTP Router

🛤️ Custom Konsultin YAML-based HTTP router for fasthttp with automatic handler binding.

## Installation

```bash
go get github.com/konsultin/routek
```

## Quick Start

```yaml
# api-route.yaml
users:
  route:
    - get: /v1/users
      handler: List
    - get: /v1/users/{id}
      handler: GetByID
    - post: /v1/users
      handler: Create
```

```go
import (
    "github.com/konsultin/routek"
    "github.com/valyala/fasthttp"
)

// Create handler struct
type UserHandler struct {
    service *UserService
}

func (h *UserHandler) List(ctx *fasthttp.RequestCtx) error {
    users, err := h.service.List()
    return err
}

// Initialize router
router, err := routek.NewRouter(routek.Config{
    RouteFile: "api-route.yaml",
    Handlers: map[string]any{
        "users": &UserHandler{service: userService},
    },
})

fasthttp.ListenAndServe(":8080", router.Handler)
```

## Features

- **YAML Configuration** - Define routes in external file
- **Automatic Handler Binding** - Map handlers by method name
- **fasthttp Integration** - Built on high-performance fasthttp
- **JSON Responder** - Built-in response helper
- **Multiple HTTP Methods** - GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS

## License

MIT License - see [LICENSE](LICENSE)
