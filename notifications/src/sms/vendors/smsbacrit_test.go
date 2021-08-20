package vendors

import (
	"context"
	"errors"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"

	"gopkg.in/h2non/gock.v1"
)

func Test_smsbacrit(tt *testing.T) {

	// // read toml file
	var cfg SMSVendrsCfg
	if _, err := toml.DecodeFile("../../../config/vendors.toml", &cfg); err != nil {
		panic(err.Error())
	}

	tt.Run("success response", func(t *testing.T) {
		defer gock.Off()

		gock.New("https://simpleservice.ahmadfajar.repl.co").
			Post("/send/sms/bacrit").
			Reply(200).
			JSON(map[string]interface{}{
				"recipients": map[string]interface{}{
					"items": []interface{}{
						map[string]interface{}{
							"recipient":    899,
							"status":       "sent",
							"statusReason": "sms is sent",
						},
					},
				},
			})

		cf := cfg.Configs["bacrit"]
		smsbarcit := BacritSMSVendor(&cf)
		_, err := smsbarcit.Send(context.Background(), SMSPayload{
			To:      "899",
			Message: "hi, folks",
		})

		assert.Equal(t, nil, err)
	})

	tt.Run("failed response", func(t *testing.T) {
		defer gock.Off()

		gock.New("https://simpleservice.ahmadfajar.repl.co").
			Post("/send/sms/bacrit").
			Reply(500).
			SetError(errors.New("error response ges"))

		cf := cfg.Configs["bacrit"]
		smsbarcit := BacritSMSVendor(&cf)
		_, err := smsbarcit.Send(context.Background(), SMSPayload{
			To:      "899",
			Message: "hi, folks",
		})

		assert.Equal(t, true, err != nil)
	})
}
