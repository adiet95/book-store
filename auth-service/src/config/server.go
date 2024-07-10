package config

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"

	"github.com/adiet95/book-store/auth-service/src/routers"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start apllication http",
	RunE:  Server,
}

// @title           Swagger Book Store
// @version         1.0
// @description     Book Store Services.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath

// @securityDefinitions.bearer  BearerAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func Server(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		var addrs string = os.Getenv("PORT")

		server := &http.Server{
			Addr:    addrs,
			Handler: mainRoute,
		}
		mainRoute.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		err = server.ListenAndServe()
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}
