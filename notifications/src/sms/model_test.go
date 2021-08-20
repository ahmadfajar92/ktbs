package sms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_model(tt *testing.T) {

	tt.Run("success create model", func(t *testing.T) {
		msg := "hi, Folks"
		to := 899

		smsnotif := SMSNotification{
			To:      to,
			Message: msg,
		}

		assert.Equal(t, to, smsnotif.To)
		assert.Equal(t, msg, smsnotif.Message)
	})

}
