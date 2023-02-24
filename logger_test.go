package golangloggingstudy

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

// ============================================================================================== //
// Logger & fmt.println outputnya mirip2, tp kl logger ad infomasi tambahan spt waktu dan level
func TestLogger(t *testing.T) {

	logger := logrus.New()

	logger.Println("Hello logger")
	fmt.Println("Hello world")
}

// ============================================================================================== //
