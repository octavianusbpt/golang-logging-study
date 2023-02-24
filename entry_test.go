package golangloggingstudy

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// ============================================================================================== //
// Saat kita buat logger.info dll itu sbnrnya dy buat entry baru, yg isinya ada level, info, dll
// Kita bisa buat entry sndr, walaupun jrg kepake
func TestEntry(t *testing.T) {

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.WithField("name", "octa")
	entry.Info("Hello log")

}

// ============================================================================================== //
