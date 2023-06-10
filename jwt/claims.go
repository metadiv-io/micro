package jwt

const (
	TYPE_ADMIN          = "admin"
	TYPE_USER           = "user"
	TYPE_WORKSPACE_user = "workspace_user"
)

type Claims struct {
	UUID      string `json:"uuid"`
	UserUUID  string `json:"user_uuid"`
	Username  string `json:"username"`
	Type      string `json:"type"`
	Email     string `json:"email"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Workspace string `json:"workspace"`
}
