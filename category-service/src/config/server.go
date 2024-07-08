package config

import (
	"net/http"
	"os"

	"github.com/adiet95/book-store/category-service/src/routers"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start apllication",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		var addrs string = os.Getenv("PORT")

		server := &http.Server{
			Addr:    addrs,
			Handler: mainRoute,
		}
		err = server.ListenAndServe()
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}
