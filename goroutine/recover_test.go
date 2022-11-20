package goroutine

import (
	"testing"

	"github.com/coeusy/gopkg/logger"
)

func TestGoRecover(t *testing.T) {
	logger.InitZap("test")
	defer GoRecover()
	panic("test")
}
