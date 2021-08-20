package vendors

import (
	"context"
	"errors"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func Test_smsngehe(tt *testing.T) {

	// // read toml file
	var cfg SMSVendrsCfg
	if _, err := toml.DecodeFile("../../../config/vendors.toml", &cfg); err != nil {
		panic(err.Error())
	}

	tt.Run("success response", func(t *testing.T) {
		defer gock.Off()

		gock.New("https://simpleservice.ahmadfajar.repl.co").
			Post("/send/sms/ngehe").
			Reply(200).
			JSON(map[string]interface{}{
				"recipients": 899,
				"status":     "terkirim",
				"reason":     "sms terkirim",
			})

		cf := cfg.Configs["ngehe"]
		smsbarcit := NgeheSMSVendor(&cf)
		_, err := smsbarcit.Send(context.Background(), SMSPayload{
			To:      "899",
			Message: "hi, folks",
		})

		assert.Equal(t, nil, err)
	})

	tt.Run("failed response", func(t *testing.T) {
		defer gock.Off()

		gock.New("https://simpleservice.ahmadfajar.repl.co").
			Post("/send/sms/ngehe").
			Reply(500).
			SetError(errors.New("error cuy"))

		cf := cfg.Configs["ngehe"]
		smsbarcit := NgeheSMSVendor(&cf)
		_, err := smsbarcit.Send(context.Background(), SMSPayload{
			To:      "899",
			Message: "hi, folks",
		})

		assert.Equal(t, true, err != nil)
	})
}
