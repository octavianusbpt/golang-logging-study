package golangloggingstudy

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

// ============================================================================================== //
// Secara default, logger bakal tampilin informasi logger-nya ke dlm console.
// Kalo kita mau ubah destination diplay information-nya, kita bisa pake logger.output
// Misal kita mau display informasi-nya di file/database aja.
func TestOutput(t *testing.T) {

	logger := logrus.New()

	// Buat file baru utk tampung info logger
	dummyLogFile, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(dummyLogFile)

	logger.Info("Hello logging")
	logger.Warn("Hello logging")
	logger.Error("Hello logging")
}

// ============================================================================================== //
