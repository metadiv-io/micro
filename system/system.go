package system

import "github.com/metadiv-io/env"

var (
	SYSTEM_UUID string
	SYSTEM_NAME string
)

func init() {
	SYSTEM_UUID = env.String("SYSTEM_UUID", "")
	if SYSTEM_UUID == "" {
		panic("SYSTEM_UUID is required")
	}
	SYSTEM_NAME = env.String("SYSTEM_NAME", "")
	if SYSTEM_NAME == "" {
		panic("SYSTEM_NAME is required")
	}
}
