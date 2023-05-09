package auth

import (
	"github.com/metadiv-io/env"
	"github.com/metadiv-io/err_map"
)

var (
	AUTH_SERVICE_URL string
)

func init() {
	AUTH_SERVICE_URL = env.String("AUTH_SERVICE_URL", "")
	if AUTH_SERVICE_URL == "" {
		panic("AUTH_SERVICE_URL is required")
	}
}

// These are the types of jwt tokens.
const (
	JWT_TYPE_ADMIN = "admin"
	JWT_TYPE_USER  = "user"
	JWT_TYPE_API   = "api"
)

// This is the jwt public key.
var (
	JWT_PUBLIC_PEM string
)

// These are the errors will be used in all the services.
const (
	ERR_CODE_UNAUTHORIZED          = "b97cf20d-42b6-470e-9e08-b4bb852c3811"
	ERR_CODE_FORBIDDEN             = "7792176d-0196-4a57-a959-93062c2b9b41"
	ERR_CODE_INTERNAL_SERVER_ERROR = "b6a82bc6-5884-41e1-8b6f-1a013b7da835"
)

func init() {
	err_map.Register(ERR_CODE_UNAUTHORIZED, "Unauthorized")
	err_map.Register(ERR_CODE_FORBIDDEN, "Forbidden")
	err_map.Register(ERR_CODE_INTERNAL_SERVER_ERROR, "Internal Server Error")
}
