package sms

import (
	"context"
	"errors"
	"strconv"

	"notifications/src/shared"
	"notifications/src/sms/vendors"
)

// smsUsecase struct
type smsUsecase struct {
	smsvndrs shared.Vendors
}

// Usecase interface
type SMSUsecase interface {
	Send(ctx context.Context, d interface{}) (r *shared.Result)
	OnOff(ctx context.Context, d interface{}) (r *shared.Result)
	Vendors(ctx context.Context) (r *shared.Result)
}

func Usecase(smsvndrs shared.Vendors) SMSUsecase {
	return &smsUsecase{
		smsvndrs: smsvndrs,
	}
}

func (notif *smsUsecase) Send(ctx context.Context, d interface{}) (r *shared.Result) {
	request, ok := d.(SMSNotification)
	if !ok {
		msg := "failed bind request payload to SMSNotification struct"
		r = shared.FailedResult(msg, errors.New(msg))
		return
	}

	res, err := notif.smsvndrs.ActiveVendor().Send(ctx, vendors.SMSPayload{
		To:      strconv.Itoa(request.To),
		Message: request.Message,
	})

	if err != nil {
		r = shared.FailedResult(err.Error(), err)
		return
	}

	r = shared.SuccessResult(res)
	return
}

func (notif *smsUsecase) OnOff(ctx context.Context, d interface{}) (r *shared.Result) {
	request, ok := d.(ToggleSMSNotification)
	if !ok {
		msg := "failed bind request payload to ToggleSMSNotification struct"
		r = shared.FailedResult(msg, errors.New(msg))
		return
	}

	vndr := notif.smsvndrs.Get(request.Name)
	if vndr == nil {
		msg := "vendor not founded"
		r = shared.FailedResult(msg, errors.New(msg))
		return
	}
	if vndr.IsDefault() && !request.Status {
		msg := "cannot deactivate default vendor"
		r = shared.FailedResult(msg, errors.New(msg))
		return
	}

	vndr.OnOff(ctx, request.Status)
	return shared.SuccessResult(map[string]interface{}{
		"name":   request.Name,
		"status": vndr.IsOn(),
	})
}

func (notif *smsUsecase) Vendors(ctx context.Context) (r *shared.Result) {
	vendors := notif.smsvndrs.All()
	vndrs := []interface{}{}
	for k, vndr := range vendors {
		vndrs = append(vndrs, map[string]interface{}{
			"name":    k,
			"active":  vndr.IsOn(),
			"default": vndr.IsDefault(),
		})
	}

	return shared.SuccessResult(vndrs)
}
