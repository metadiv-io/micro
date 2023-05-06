package jwt

type Claims struct {
	UUID       string   `json:"uuid"`
	Name       string   `json:"name"`
	IP         string   `json:"ip"`
	TokenType  string   `json:"token_type"`
	Workspaces []string `json:"workspaces"`
}

// NewClaims creates a new Claims.
func NewClaims(uuid, name, ip, tokenType string, workspaces []string) *Claims {
	return &Claims{
		UUID:       uuid,
		Name:       name,
		IP:         ip,
		TokenType:  tokenType,
		Workspaces: workspaces,
	}
}

func (c *Claims) HasWorkspace(workspace string) bool {
	for _, ws := range c.Workspaces {
		if ws == workspace {
			return true
		}
	}
	return false
}
