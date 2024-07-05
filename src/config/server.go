package config

import (
	"os"

	"github.com/adiet95/book-store/src/routers"
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

		mainRoute.Run(addrs)
		return nil
	} else {
		return err
	}
}
