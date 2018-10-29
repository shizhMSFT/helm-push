package api

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

// ChartMuseum defines APIs for ChartMuseum
type ChartMuseum struct{}

// NewChartUploadRequest creates a new request for uploading a chart to ChartMuseum
func (ChartMuseum) NewChartUploadRequest(baseURL, contextPath, chartPackagePath string, force bool) (*http.Request, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(contextPath, "api", strings.TrimPrefix(u.Path, contextPath), "charts")
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}

	// Add ?force to request querystring to force an upload if chart version already exists
	if force {
		req.URL.RawQuery = "force"
	}

	err = setUploadChartPackageRequestBody(req, chartPackagePath)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func setUploadChartPackageRequestBody(req *http.Request, chartPackagePath string) error {
	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	defer w.Close()
	fw, err := w.CreateFormFile("chart", chartPackagePath)
	if err != nil {
		return err
	}
	w.FormDataContentType()
	fd, err := os.Open(chartPackagePath)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = io.Copy(fw, fd)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Body = ioutil.NopCloser(&body)
	return nil
}
