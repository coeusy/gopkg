package pprof

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"go.uber.org/zap"
)

var port int

func InitCMD() {
	flag.IntVar(&port, "pprof.port", 8080, "int: port for pprof")
}

func InitPprof() {
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			zap.S().Errorf("start pprof server failed: %s", err.Error())
		}
	}()
}
