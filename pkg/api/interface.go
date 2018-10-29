package api

import "net/http"

// Server defines how to create requests for various functionalities.
type Server interface {
	NewChartUploadRequest(baseURL, contextPath, chartPackagePath string, force bool) (*http.Request, error)
}
