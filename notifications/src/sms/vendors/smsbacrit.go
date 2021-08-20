package vendors

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"notifications/src/shared"
	"strconv"
	"time"
)

type (
	bacrit struct {
		token     string
		enable    bool
		uri       string
		asdefault bool
		client    *http.Client
	}

	bacritResponse struct {
		ID                string      `json:"id"`
		Href              string      `json:"href"`
		Direction         string      `json:"direction"`
		Type              string      `json:"type"`
		Originator        string      `json:"originator"`
		Body              string      `json:"body"`
		Reference         interface{} `json:"reference"`
		Validity          interface{} `json:"validity"`
		Gateway           interface{} `json:"gateway"`
		TypeDetails       TypeDetails `json:"typeDetails"`
		Datacoding        string      `json:"datacoding"`
		Mclass            int         `json:"mclass"`
		ScheduledDatetime interface{} `json:"scheduledDatetime"`
		CreatedDatetime   time.Time   `json:"createdDatetime"`
		Recipients        Recipients  `json:"recipients"`
	}

	TypeDetails struct{}

	Items struct {
		Recipient      int64     `json:"recipient"`
		Status         string    `json:"status"`
		StatusDatetime time.Time `json:"statusDatetime"`
		StatusReason   string    `json:"statusReason"`
	}

	Recipients struct {
		TotalCount               int     `json:"totalCount"`
		TotalSentCount           int     `json:"totalSentCount"`
		TotalDeliveredCount      int     `json:"totalDeliveredCount"`
		TotalDeliveryFailedCount int     `json:"totalDeliveryFailedCount"`
		Items                    []Items `json:"items"`
	}
)

var bcStatus = map[string]string{
	"sent":      smsStatus.Sent,
	"delivered": smsStatus.Delivered,
	"pending":   smsStatus.Pending,
}

func BacritSMSVendor(cfg *SMSVendrCfg) shared.Vendor {
	bc := new(bacrit)

	// setup http client
	bc.httpClientSetup()

	bc.token = cfg.Token
	bc.uri = cfg.Uri
	bc.enable = cfg.Enable == "true"
	bc.asdefault = cfg.IsDefault == "true"

	return bc
}

func (bc *bacrit) httpClientSetup() {
	bc.client = &http.Client{}
}

func (bc *bacrit) Send(ctx context.Context, p interface{}) (r interface{}, err error) {
	payload, ok := p.(SMSPayload)
	if !ok {
		err = errors.New("failed to construct payload into SMSPayload struct")
		return
	}

	j, err := json.Marshal(payload)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", bc.uri, bytes.NewBuffer(j))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	if bc.token != "" {
		req.Header.Set("Authorization", bc.token)
	}

	res, err := bc.client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	var d bacritResponse
	err = json.NewDecoder(res.Body).Decode(&d)
	if err != nil {
		return
	}

	r = bc.respSerialize(d)

	return
}

func (bc *bacrit) respSerialize(d bacritResponse) SMSResponse {
	return SMSResponse{
		To:      strconv.Itoa(int(d.Recipients.Items[0].Recipient)),
		Status:  bcStatus[d.Recipients.Items[0].Status],
		Message: d.Recipients.Items[0].StatusReason,
	}
}

func (bc *bacrit) IsOn() bool {
	return bc.enable
}

func (bc *bacrit) IsDefault() bool {
	return bc.asdefault
}

func (bc *bacrit) OnOff(ctx context.Context, s bool) (err error) {
	bc.enable = s
	return
}
