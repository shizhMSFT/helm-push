package helm

type (
	// ServerInfo describes the server information
	ServerInfo struct {
		ContextPath string `json:"contextPath"`
		APISpecType string `json:"ApiSpecType"`
	}
)
