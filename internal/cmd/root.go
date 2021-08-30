package cmd

import (
	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	rootLog = zap.New()
	rootCmd = &cobra.Command{
		Use:   "combo",
		Short: "Create combinations of kubernetes manifests",
	}
)

// Execute executes the root command.
func Execute(log logr.Logger) error {
	if log != nil {
		rootLog = log
	}

	return rootCmd.Execute()
}
