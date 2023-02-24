package golangloggingstudy

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

// ============================================================================================== //
// Hook merupakan sebuah struct yg bisa ditambahkan ke logger sbg callback yg di execute saat
// event log tertentu.
// Pakai function AddHook()

type SampleHook struct {
}

// Levels menentukan pada level mana aja si hook ini di execute
func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel}
}

// Fire merupakan function yg bakal di execute saat hook ter-trigger
func (s *SampleHook) Fire(entry *logrus.Entry) error {
	// Bisa apa aja, send email, sms, chat, dll
	fmt.Println("Hook triggered", entry.Level, entry.Message)
	// Bisa return error-nya kl emg ad error
	return nil
}

func TestHook(t *testing.T) {

	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	logger.Info("Infooo")
	logger.Warn("Warninggg")

}

// ============================================================================================== //
