package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/helm/chartmuseum/pkg/repo"
	"k8s.io/helm/pkg/chartutil"
)

// AzureContainerRegistry defines APIs for Azure Container Registry
type AzureContainerRegistry struct{}

// NewChartUploadRequest creates a new request for uploading a chart to ACR
func (AzureContainerRegistry) NewChartUploadRequest(baseURL, _, chartPackagePath string, force bool) (*http.Request, error) {
	file, err := ioutil.ReadFile(chartPackagePath)
	if err != nil {
		return nil, err
	}

	chart, err := chartutil.LoadArchive(bytes.NewBuffer(file))
	if err != nil {
		return nil, err
	}

	filename := repo.ChartPackageFilenameFromNameVersion(chart.Metadata.Name, chart.Metadata.Version)
	url := path.Join(baseURL, "_blobs", filename)
	method := "PUT"
	if force {
		method = "PATCH"
	}

	return http.NewRequest(method, url, bytes.NewBuffer(file))
}
