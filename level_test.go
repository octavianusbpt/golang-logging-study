package golangloggingstudy

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// ============================================================================================== //
// Di logger ada level2-nya dari yg terendah sampe yg paling tinggi, urutan dr paling rendah:
// 1. Trace
// 2. Debug
// 3. Info (Yg paling sering dipake)
// 4. Warn
// 5. Error
// 6. Fatal => Yg ini langsung panggil os.exit(1) stlh logging
// 7. Panic => Yg ini panggil panic stlh logging
func TestLevel(t *testing.T) {

	logger := logrus.New()

	logger.Trace("Level Trace") // Untuk level trace dan debug gk bakal keluar di console
	logger.Debug("Level Debug") // Cuma level info ke atas yg bakal muncul
	logger.Info("Level Info")
	logger.Warn("Level Warn")
	logger.Error("Level Error")
	// logger.Fatal("Level Fatal")
	// logger.Panic("Level Panic")

}

// ============================================================================================== //

// ============================================================================================== //
// Kalo kita mau set tampilin error message di console-nya mulai dr level mana ke atas
// bisa pake setLevel
func TestLoggingLevel(t *testing.T) {

	logger := logrus.New()

	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("Level Trace")
	logger.Debug("Level Debug")
	logger.Info("Level Info")
	logger.Warn("Level Warn")
	logger.Error("Level Error")

}

// ============================================================================================== //
