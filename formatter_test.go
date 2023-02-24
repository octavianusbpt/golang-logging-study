package golangloggingstudy

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// ============================================================================================== //
// Logrus secara default punya 2 jenis formatter.
// Ada Text Formatter (default yg dipakai), sama JSON Formatter
// Dibawah ini contoh kl kita pake JSON Formatter
func TestFormatter(t *testing.T) {

	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello logging")
	logger.Warn("Hello logging")
	logger.Error("Hello logging")

}

// ============================================================================================== //
