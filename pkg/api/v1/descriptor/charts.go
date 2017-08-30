/*
Copyright 2017 caicloud authors. All rights reserved.
*/

package descriptor

import (
	"net/http"

	"github.com/mixj93/helm-registry/pkg/api/definition"
	"github.com/mixj93/helm-registry/pkg/api/models"
	"github.com/mixj93/helm-registry/pkg/api/v1/handlers"
	"github.com/mixj93/helm-registry/pkg/common"
)

func init() {
	registerDescriptors(charts)
}

// charts descriptors
var charts = []definition.Descriptor{
	{
		Path: "/spaces/{space}/charts",
		Handlers: []definition.Handler{
			{
				HTTPMethod: http.MethodGet,
				Handler:    definition.NewHandlerDecoration(definition.VerbList, handlers.ListCharts).Handle,
				Doc:        "List all charts in space",
				PathParams: []definition.Param{
					{
						Name:     "space",
						Type:     "string",
						Doc:      "space name",
						Required: true,
					},
				},
				QueryParams: []definition.Param{
					{
						Name:     "start",
						Type:     "number",
						Doc:      "Query start index",
						Required: false,
						Default:  0,
					},
					{
						Name:     "limit",
						Type:     "number",
						Doc:      "Specify the number of records to return",
						Required: false,
						Default:  common.DefaultPagingLimit,
					},
				},
				StatusCode: []definition.StatusCode{
					definition.StatusCode{Code: http.StatusOK, Message: "Success and respond with an array of chart names",
						Sample: &models.ListResponse{
							Metadata: models.Metadata{
								Total:       10,
								ItemsLength: 1,
							},
							Items: []string{
								"chartName",
							},
						}},
				},
			},
			{
				HTTPMethod: http.MethodPost,
				Handler:    definition.NewHandlerDecoration(definition.VerbCreate, handlers.CreateOrUploadChart).Handle,
				Doc:        "Create a chart by config or Upload a chart",
				Note: `
If ContentType is 'multipart/form-data', the request is handled as uploading a chart. Otherwise it is
handled by creating chart and request body should be an orchestration config. The config is a json string;
below is a sample:
{
    "save":{                            // key, required
        "chart":"chart name",           // string, required
        "version":"1.0.0",              // string, required
        "description":"description"     // string, optional
    },
    "configs":{                         // key, required
        "package":{                     // key, required
            "independent":true,         // boolean, required
            "space":"space name",       // string, required
            "chart":"chart name",       // string, required
            "version":"version number"  // string, required
        },
        "_config": {                    // key, required
            // root chart config
        },
        "chartB": {
            "package":{
                "independent":true,
                "space":"space name",
                "chart":"chart name",
                "version":"version number"
            },
            "_config": {
                // chartB config
            },
            "chartD":{
                "package":{
                    "independent":false,
                    "space":"space name",
                    "chart":"chart name",
                    "version":"version number"
                },
                "_config": {
                    // chartD config
                }
            }
        },
        "chartC": {
            "package":{
                "independent":false,
                "space":"space name",
                "chart":"chart name",
                "version":"version number"
            },
            "_config": {
                // chartC config
            }
        }
    }
}
`,
				QueryParams: []definition.Param{
					{
						Name:     "chartfile",
						Type:     "multipart/form-data",
						Doc:      "An archive file of chart. Only valid when upload a chart",
						Required: true,
					},
				},
				StatusCode: []definition.StatusCode{
					definition.StatusCode{Code: http.StatusCreated, Message: "Create successfully",
						Sample: &models.ChartLink{
							Space:   "spaceName",
							Chart:   "chartName",
							Version: "1.0.0",
							Link:    "/spaces/spaceName/charts/chartName/versions/1.0.0",
						}},
				},
			},
		},
	},
	{
		Path: "/spaces/{space}/charts/{chart}",
		Handlers: []definition.Handler{
			{
				HTTPMethod: http.MethodDelete,
				Handler:    definition.NewHandlerDecoration(definition.VerbDelete, handlers.DeleteChart).Handle,
				Doc:        "Delete a chart and its all versions",
				PathParams: []definition.Param{
					{
						Name:     "space",
						Type:     "string",
						Doc:      "space name",
						Required: true,
					},
					{
						Name:     "chart",
						Type:     "string",
						Doc:      "chart name",
						Required: true,
					},
				},
				StatusCode: []definition.StatusCode{
					definition.StatusCode{Code: http.StatusNoContent, Message: "Delete successfully"},
				},
			},
		},
	},
}
