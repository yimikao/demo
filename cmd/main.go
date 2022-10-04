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
	// clear flags initialized in static dependencies
	if flag.CommandLine.Lookup("log_dir") != nil {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	}
	klog.InitFlags(nil)
	var loggingFormat string
	flag.StringVar(&loggingFormat, "loggingFormat", logging.TextFormat, "This flag sets the logging mode for klog")

	// if err := flag.Set("v", "2"); err != nil {
	// 	// fmt.Printf("failed to set log level: %s", err.Error())
	// 	if errors.Is(err, fmt.Errorf("no such flag -v")) {
	// 		fmt.Println("yo")
	// 	} else {
	// 		fmt.Println("yayyyyyyy")
	// 	}

	// }
	flag.Parse()
	// return
	/*

		// if loggingFormat == "text" {
		log.SetLogger(klog.NewKlogr())
		if loggingFormat == "json" {
			// zapLog, _ := zap.NewProduction()
			zl := zerolog.New(os.Stderr)
			zl = zl.With().Caller().Timestamp().Logger()
			klog.SetLogger(zerologr.New(&zl))
			log.SetLogger(klogr.New())
		}
		setupLog.Info("custom message")
		sa.Info("hey")
	*/
	if err := logging.Setup(loggingFormat); err != nil {
		fmt.Println("failed to setup logger", err)
		os.Exit(1)
	}

	setup.Info("code in main.go")
	config.Config("added some config")

}
