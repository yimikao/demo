package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yimikao/browse/config"
	"github.com/yimikao/browse/logging"
	"k8s.io/klog/v2"
)

var (
	setup = logging.WithName("setup")
)

func main() {
	if flag.CommandLine.Lookup("log_dir") != nil {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	}
	klog.InitFlags(nil)
	var loggingFormat string
	flag.StringVar(&loggingFormat, "loggingFormat", logging.TextFormat, "This flag sets the logging mode for klog")

	flag.Parse()

	if err := logging.Setup(loggingFormat); err != nil {
		fmt.Println("failed to setup logger", err)
		os.Exit(1)
	}

	setup.Info("code in main.go")
	config.Config("added some config")

}
