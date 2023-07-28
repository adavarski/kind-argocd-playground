package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var options struct {
	addr     string
	interval time.Duration
}

var rootCmd = &cobra.Command{
	Use:   "dummy-metrics",
	Short: "run dummy-metrics server",
	Long:  `run dummy-metrics server.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s := newServer(args[0], options.interval)
		http.HandleFunc("/metrics", s.metrics)

		return http.ListenAndServe(options.addr, nil)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var debug bool

	fs := rootCmd.Flags()
	fs.StringVar(&options.addr, "addr", ":8080", "The address to bind to")
	fs.DurationVar(&options.interval, "interval", 30*time.Second, "Metrics update interval")
	fs.BoolVar(&debug, "debug", false, "Enable debug logging")

	cobra.OnInitialize(func() {
		var logger *zap.Logger
		var err error
		if debug {
			logger, err = zap.NewDevelopment()
		} else {
			logger, err = zap.NewProduction()
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to initialize logger %v\n", err)
		}
		zap.ReplaceGlobals(logger)
	})
}
