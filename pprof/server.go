package pprof

import (
	"net/http"
	_ "net/http/pprof"

	"go.uber.org/zap"
)

func init() {
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			zap.S().Errorf("start pprof server failed: %s", err.Error())
		}
	}()
}
