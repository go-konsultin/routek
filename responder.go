package routek

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

type Responder struct {
	debug bool
}

// NewResponder creates a responder; debug=true will include error details in responses.
func NewResponder(debug bool) *Responder {
	return &Responder{debug: debug}
}

// Success sends a successful Response with the given status, code, message, and payload data.
func (r *Responder) Success(ctx *fasthttp.RequestCtx, status int, code Code, message string, data any) {
	resp := Response[any]{
		Message:   message,
		Code:      code,
		Data:      data,
		Timestamp: time.Now().UTC().UnixMilli(),
	}
	r.write(ctx, status, resp)
}

// Error standardizes error responses.
func (r *Responder) Error(ctx *fasthttp.RequestCtx, status int, code Code, message string, err error) {
	var data any

	if err != nil && r.debug {
		data = map[string]any{"error": err.Error()}
	}

	if status == 0 {
		status = fasthttp.StatusInternalServerError
	}
	if code == "" {
		code = CodeInternalError
	}
	if message == "" {
		message = "internal server error"
	}

	resp := Response[any]{
		Message:   message,
		Code:      code,
		Data:      data,
		Timestamp: time.Now().UTC().UnixMilli(),
	}
	r.write(ctx, status, resp)
}

// write marshals the payload and writes it to the response, with a resilient fallback when marshaling fails.
func (r *Responder) write(ctx *fasthttp.RequestCtx, status int, payload any) {
	body, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal response: %v", err)
		fallback := Response[any]{
			Message:   "internal server error",
			Code:      CodeInternalError,
			Data:      nil,
			Timestamp: time.Now().UTC().UnixMilli(),
		}
		fallbackBody, fallbackErr := json.Marshal(fallback)
		if fallbackErr != nil {
			log.Printf("failed to marshal fallback response: %v", fallbackErr)
			ctx.Response.Header.Set("Content-Type", "application/json")
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.SetBodyString(
				fmt.Sprintf(
					`{"message":"internal server error","code":"INTERNAL_ERROR","data":null,"timestamp":%d}`,
					time.Now().UTC().UnixMilli(),
				),
			)
			return
		}
		ctx.Response.Header.Set("Content-Type", "application/json")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody(fallbackBody)
		return
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(status)
	ctx.SetBody(body)
}
