package golangloggingstudy

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// ============================================================================================== //
// Biasanya kita ingin menambahkan informasi ke dlm logger, kl manual bisa masukkin dlm message-nya
// Kl pake logger, bisa menggunakan function WithField() => Key value pair
// Valuenya itu interface, jadi bisa dimasukkin tipe data apa aja
func TestField(t *testing.T) {

	logger := logrus.New()

	logger.WithField("username", "octavianus").Info("Hello world")
	logger.WithField("username", "octavianus").
		WithField("Name", "Octa").
		WithField("LastLogin", "Friday").
		Info("Multiple fields")

}

// ============================================================================================== //

// ============================================================================================== //
// Kalo mau multiple fields sekaligus bisa pake WithFields(), bntknya map
func TestFields(t *testing.T) {

	logger := logrus.New()

	logger.WithFields(logrus.Fields{
		"username": "octavianus",
		"Name":     "Octa",
		"Age":      22,
	}).Info("Multiple fields")

}

// ============================================================================================== //
