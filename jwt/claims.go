package jwt

type Claims struct {
	UUID      string              `json:"uuid"`
	Name      string              `json:"name"`
	IP        string              `json:"ip"`
	TokenType string              `json:"token_type"`
	AuthMap   map[string][]string `json:"auth_map"` // map[workspace_uuid][][auth_group_uuid...]
}

// NewClaims creates a new Claims.
func NewClaims(uuid, name, ip, tokenType string, authMap map[string][]string) *Claims {
	return &Claims{
		UUID:      uuid,
		Name:      name,
		IP:        ip,
		TokenType: tokenType,
		AuthMap:   authMap,
	}
}
