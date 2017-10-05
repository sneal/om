package api

import (
	"io/ioutil"
	"net/http"
)

type CertificatesService struct {
	client httpClient
}

func NewCertificatesService(client httpClient) CertificatesService {
	return CertificatesService{
		client: client,
	}
}

func (c CertificatesService) Generate(string) (string, error) {
	req, err := http.NewRequest("POST", "/api/v0/certificates/generate", nil)
	if err != nil {
		return "", err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if err = ValidateStatusOK(resp); err != nil {
		return "", err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}
