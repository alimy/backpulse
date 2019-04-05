package serve

import (
	"github.com/backpulse/core/cmd/core"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"time"
)

var (
	address string
)

func init() {
	frontpulseCmd := &cobra.Command{
		Use:   "frontpulse",
		Short: "Start frontpulse ui service",
		Long:  "Start frontpulse ui service",
		Run:   serveRun,
	}

	// Parse flags for serveCmd
	frontpulseCmd.Flags().StringVarP(&address, "addr", "a", ":3001", "service listen address")

	// Register serveCmd as sub-command
	core.Register(frontpulseCmd)
}

func serveRun(cmd *cobra.Command, args []string) {
	// Setup http.Server
	server := &http.Server{
		Handler: newAssets(),
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("server listening in %s", address)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
