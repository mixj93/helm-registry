/*
Copyright 2017 caicloud authors. All rights reserved.
*/

package cmd

import (
	"time"

	"github.com/mixj93/helm-registry/pkg/api"
	"github.com/mixj93/helm-registry/pkg/common"
	"github.com/mixj93/helm-registry/pkg/log"
	"github.com/emicklei/go-restful"
	"github.com/go-openapi/spec"
	"github.com/spf13/cobra"
	"gopkg.in/tylerb/graceful.v1"

	restfulspec "github.com/emicklei/go-restful-openapi"
)

// major version of helm-registry
const version = "v1"

// config path
var configPath = ""

// serveCmd starts a http server for managing charts
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts a http server for managing a charts repository",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// read config
		config, err := newConfig(configPath)
		if err != nil {
			log.Fatal(err)
		}

		// init SpaceManager
		common.Set(common.ContextNameSpaceManager, config.Manager.Name)
		common.Set(common.ContextNameSpaceParameters, config.Manager.Parameters)
		common.MustGetSpaceManager()

		// start server
		api.Initialize()

		// install openapi path
		restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(
			restfulspec.Config{
				WebServices:    restful.RegisteredWebServices(),
				WebServicesURL: config.Listen,
				APIPath:        "/apidocs.json",
				PostBuildSwaggerObjectHandler: enrichSwaggerObject,
			},
		))

		log.Infof("Listening address %s", config.Listen)
		graceful.Run(config.Listen, 5*time.Minute, restful.DefaultContainer)
		log.Error("Server stopped")
	},
}

// enrichSwaggerObject adds more documentation to swagger API
func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Helm registry",
			Description: "Helm Registry stores helm charts in a hierarchy storage structure and provides a function to orchestrate charts from existing ones.",
			Contact: &spec.ContactInfo{
				Name:  "Guowei",
				Email: "guowei@caicloud.io",
			},
			Version: version,
		},
	}
}

func init() {
	// bind variable configPath with flag --config or -c
	serveCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "path of config.yaml")
	rootCmd.AddCommand(serveCmd)
}
