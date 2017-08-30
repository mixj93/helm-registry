/*
Copyright 2017 caicloud authors. All rights reserved.
*/

package v1

import (
	"net/http"

	"github.com/mixj93/helm-registry/pkg/storage"
)

// APIFetchChartMetadata defines an api of fetching chart metadata
type APIFetchChartMetadata struct {
	baseAPI
	// Space is the name of space
	Space string `kind:"path" name:"space"`
	// Chart is the name of Chart
	Chart string `kind:"path" name:"chart"`
	// Start is the start index of list
	Start int `kind:"query" name:"start"`
	// Limit is the max length of list
	Limit int `kind:"query" name:"limit"`
}

// NewAPIFetchChartMetadata creates an instance of APIFetchChartMetadata
func NewAPIFetchChartMetadata() *APIFetchChartMetadata {
	api := &APIFetchChartMetadata{}
	api.object = api
	api.method = http.MethodGet
	api.url = URLChartMetadata
	api.result = &MetadataCollectionResult{}
	return api
}

// Convert converts result to *models.ChartLink
func (api *APIFetchChartMetadata) Convert(result interface{}, err error) (*MetadataCollectionResult, error) {
	if err != nil {
		return nil, err
	}
	return result.(*MetadataCollectionResult), nil
}

// APIFetchVersionMetadata defines an api of fetching version metadata
type APIFetchVersionMetadata struct {
	baseAPI
	// Space is the name of space
	Space string `kind:"path" name:"space"`
	// Chart is the name of chart
	Chart string `kind:"path" name:"chart"`
	// Version is the name of Version
	Version string `kind:"path" name:"version"`
}

// NewAPIFetchVersionMetadata creates an instance of APIFetchVersionMetadata
func NewAPIFetchVersionMetadata() *APIFetchVersionMetadata {
	api := &APIFetchVersionMetadata{}
	api.object = api
	api.method = http.MethodGet
	api.url = URLVersionMetadata
	api.result = &storage.Metadata{}
	return api
}

// Convert converts result to *chart.Metadata
func (api *APIFetchVersionMetadata) Convert(result interface{}, err error) (*storage.Metadata, error) {
	if err != nil {
		return nil, err
	}
	return result.(*storage.Metadata), nil
}

// APIFetchVersionValues defines an api of fetching version values
type APIFetchVersionValues APIFetchVersionMetadata

// NewAPIFetchVersionValues creates an instance of APIFetchVersionValues
func NewAPIFetchVersionValues() *APIFetchVersionValues {
	api := &APIFetchVersionValues{}
	api.object = api
	api.method = http.MethodGet
	api.url = URLVersionValues
	api.result = []byte{}
	return api
}

// Convert converts result to []byte
func (api *APIFetchVersionValues) Convert(result interface{}, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	return result.([]byte), nil
}
