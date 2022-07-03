package integration

import (
	"app1/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type vendorsIntegration struct {
}

func NewVendorsIntegration() *vendorsIntegration {
	return &vendorsIntegration{}
}

func (vi *vendorsIntegration) GetByID(ID string) (*domain.Vendor, error) {
	response, err := http.Get(fmt.Sprintf("localhost:8080/vendors/%v", ID))
	if err != nil {
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	retrievedVendor := &domain.Vendor{}
	err = json.Unmarshal(responseBody, retrievedVendor)
	if err != nil {
		return nil, err
	}

	return retrievedVendor, nil
}
