package serve

import (
	"github.com/backpulse/core/cmd/core"
	"github.com/backpulse/core/database"
	"github.com/backpulse/core/routes"
	"github.com/backpulse/core/utils"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
)

func init() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start backpulse service",
		Long:  "Start backpulse service",
		Run:   serveRun,
	}

	// Register serveCmd as sub-command
	core.Register(serveCmd)
}

func serveRun(cmd *cobra.Command, args []string) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config := utils.GetConfig()

	database.Connect(config.URI, config.Database)
	utils.InitStripe()

	r := routes.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Access-Control-Allow-Origin", "origin", "X-Requested-With", "Authorization", "Content-Type", "Language"},
		AllowedMethods: []string{"DELETE", "POST", "GET", "PUT"},
	})

	handler := c.Handler(r)

	var port string
	if os.Getenv("PORT") == "" {
		port = ":8000"
	} else {
		port = ":" + os.Getenv("PORT")
	}

	err := http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatal(err)
	}
}
