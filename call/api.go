package call

import (
	"github.com/gin-gonic/gin"
)

// GET make a GET request with micro trace standard
func GET[T any](ctx *gin.Context, url string, params map[string]string, headers map[string]string) (*Response[T], error) {
	return get[T](ctx, url, params, headers)
}

// POST make a POST request with micro trace standard
func POST[T any](ctx *gin.Context, url string, body interface{}, headers map[string]string) (*Response[T], error) {
	return nonGet[T](ctx, url, "POST", body, headers)
}

// PUT make a PUT request with micro trace standard
func PUT[T any](ctx *gin.Context, url string, body interface{}, headers map[string]string) (*Response[T], error) {
	return nonGet[T](ctx, url, "PUT", body, headers)
}

// DELETE make a DELETE request with micro trace standard
func DELETE[T any](ctx *gin.Context, url string, body interface{}, headers map[string]string) (*Response[T], error) {
	return nonGet[T](ctx, url, "DELETE", body, headers)
}
