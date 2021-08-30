package main

import (
	"os"

	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/njhale/combo/internal/cmd"
)

func main() {
	log := zap.New()
	if err := cmd.Execute(log); err != nil {
		log.Error(err, "command failed")
		os.Exit(1)
	}
}
