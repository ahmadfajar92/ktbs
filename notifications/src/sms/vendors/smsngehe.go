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
	ngehe struct {
		token     string
		enable    bool
		uri       string
		asdefault bool
		client    *http.Client
	}

	ngeheResponse struct {
		ID              string    `json:"id"`
		Link            string    `json:"href"`
		CreatedDatetime time.Time `json:"createdDatetime"`
		Recipient       int64     `json:"recipient"`
		Status          string    `json:"status"`
		StatusDatetime  time.Time `json:"statusDatetime"`
		Reason          string    `json:"reason"`
	}
)

var ngStatus = map[string]string{
	"terkirim": smsStatus.Sent,
	"diterima": smsStatus.Delivered,
	"tertunda": smsStatus.Pending,
}

func NgeheSMSVendor(cfg *SMSVendrCfg) shared.Vendor {
	ng := new(ngehe)

	// setup http client
	ng.httpClientSetup()

	ng.token = cfg.Token
	ng.uri = cfg.Uri
	ng.enable = cfg.Enable == "true"
	ng.asdefault = cfg.IsDefault == "true"

	return ng
}

func (ng *ngehe) httpClientSetup() {
	ng.client = &http.Client{}
}

func (ng *ngehe) Send(ctx context.Context, p interface{}) (r interface{}, err error) {
	payload, ok := p.(SMSPayload)
	if !ok {
		err = errors.New("failed to construct payload into SMSPayload struct")
		return
	}

	j, err := json.Marshal(payload)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", ng.uri, bytes.NewBuffer(j))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if ng.token != "" {
		req.Header.Set("Authorization", ng.token)
	}

	res, err := ng.client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	var d ngeheResponse
	err = json.NewDecoder(res.Body).Decode(&d)
	if err != nil {
		return
	}

	r = ng.respSerialize(d)

	return
}

func (ng *ngehe) respSerialize(d ngeheResponse) SMSResponse {
	return SMSResponse{
		To:      strconv.Itoa(int(d.Recipient)),
		Status:  ngStatus[d.Status],
		Message: d.Reason,
	}
}

func (ng *ngehe) IsOn() bool {
	return ng.enable
}

func (ng *ngehe) IsDefault() bool {
	return ng.asdefault
}

func (ng *ngehe) OnOff(ctx context.Context, s bool) (err error) {
	ng.enable = s
	return
}
