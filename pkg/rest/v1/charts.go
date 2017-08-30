/*
Copyright 2017 caicloud authors. All rights reserved.
*/

package v1

import (
	"net/http"

	"github.com/mixj93/helm-registry/pkg/api/models"
)

// APIListCharts defines an api of listing charts
type APIListCharts struct {
	baseAPI
	// Space is the name of space
	Space string `kind:"path" name:"space"`
	// Start is the start index of list
	Start int `kind:"query" name:"start"`
	// Limit is the max length of list
	Limit int `kind:"query" name:"limit"`
}

// NewAPIListCharts creates an instance of APIListCharts
func NewAPIListCharts() *APIListCharts {
	api := &APIListCharts{}
	api.object = api
	api.method = http.MethodGet
	api.url = URLCharts
	api.result = &StringCollectionResult{}
	return api
}

// Convert converts result to *StringCollectionResult
func (api *APIListCharts) Convert(result interface{}, err error) (*StringCollectionResult, error) {
	if err != nil {
		return nil, err
	}
	return result.(*StringCollectionResult), nil
}

// APICreateChart defines an api of creating chart
type APICreateChart struct {
	baseAPI
	// Space is the name of space
	Space string `kind:"path" name:"space"`
	// Config is a json string of orchestration config
	Config string `kind:"body"`
}

// APICreateChart creates an instance of APICreateChart
func NewAPICreateChart() *APICreateChart {
	api := &APICreateChart{}
	api.object = api
	api.method = http.MethodPost
	api.url = URLCharts
	api.result = &models.ChartLink{}
	return api
}

// Convert converts result to *models.Link
func (api *APICreateChart) Convert(result interface{}, err error) (*models.ChartLink, error) {
	if err != nil {
		return nil, err
	}
	return result.(*models.ChartLink), nil
}

// APIUploadChart defines an api of uploading version
type APIUploadChart struct {
	baseAPI
	// Space is the name of space
	Space string `kind:"path" name:"space"`
	// ChartFile is a chart file
	ChartFile *File `kind:"file" name:"chartfile"`
}

// NewAPIUploadChart creates an instance of APIUploadChart
func NewAPIUploadChart() *APIUploadChart {
	api := &APIUploadChart{}
	api.object = api
	api.method = http.MethodPost
	api.url = URLCharts
	api.result = &models.ChartLink{}
	api.ChartFile = &File{}
	return api
}

// Convert converts result to *models.ChartLink
func (api *APIUploadChart) Convert(result interface{}, err error) (*models.ChartLink, error) {
	if err != nil {
		return nil, err
	}
	return result.(*models.ChartLink), nil
}

// APIDeleteChart defines an api of deleting chart
type APIDeleteChart struct {
	baseAPI
	// Space is the name of space
	Space string `kind:"path" name:"space"`
	// Chart is the name of Chart
	Chart string `kind:"path" name:"chart"`
}

// APICreateChart creates an instance of APICreateChart
func NewAPIDeleteChart() *APIDeleteChart {
	api := &APIDeleteChart{}
	api.object = api
	api.method = http.MethodDelete
	api.url = URLChart
	return api
}

// Convert converts result to *models.Link
func (api *APIDeleteChart) Convert(result interface{}, err error) error {
	return err
}
