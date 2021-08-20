package sms

import (
	"context"
	"errors"
	"notifications/src/shared"
	mockintf "notifications/src/shared/mock_shared"
	"notifications/src/sms/vendors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_send_usecase(tt *testing.T) {
	ctrl := gomock.NewController(tt)
	mckVendors := mockintf.NewMockVendors(ctrl)
	mckVendor := mockintf.NewMockVendor(ctrl)

	smspayload := vendors.SMSPayload{
		To:      "876923",
		Message: "hi, folks",
	}
	smsresponse := vendors.SMSResponse{
		To:      "876923",
		Message: "hi, folks",
	}

	tt.Run("positive send usecase", func(t *testing.T) {

		mckVendors.EXPECT().ActiveVendor().Return(mckVendor)
		mckVendor.EXPECT().Send(context.Background(), smspayload).Return(smsresponse, nil)

		ucs := Usecase(mckVendors)
		resp := ucs.Send(context.Background(), SMSNotification{
			To:      876923,
			Message: "hi, folks",
		})

		assert.Equal(t, true, resp.Error == nil)
	})

	tt.Run("cannot bind request send usecase", func(t *testing.T) {

		mckVendors.EXPECT().ActiveVendor().Return(mckVendor)
		mckVendor.EXPECT().Send(context.Background(), smspayload).Return(smsresponse, nil)

		ucs := Usecase(mckVendors)
		resp := ucs.Send(context.Background(), nil)
		assert.Equal(t, true, resp.Error != nil)
	})

	tt.Run("vendor return error send usecase", func(t *testing.T) {

		mckVendors.EXPECT().ActiveVendor().Return(mckVendor)
		mckVendor.EXPECT().Send(context.Background(), smspayload).Return(smsresponse, errors.New("vendor return error"))

		ucs := Usecase(mckVendors)
		resp := ucs.Send(context.Background(), nil)
		assert.Equal(t, true, resp.Error != nil)
	})
}

func Test_onoff_usecase(tt *testing.T) {
	ctrl := gomock.NewController(tt)
	mckVendors := mockintf.NewMockVendors(ctrl)
	mckVendor := mockintf.NewMockVendor(ctrl)

	tt.Run("positive toggle usecase", func(t *testing.T) {
		name := "ngehe"
		status := true
		mckVendors.EXPECT().Get(name).Return(mckVendor)
		mckVendor.EXPECT().IsDefault().Return(false)
		mckVendor.EXPECT().OnOff(context.Background(), status)
		mckVendor.EXPECT().IsOn().Return(status)

		ucs := Usecase(mckVendors)
		resp := ucs.OnOff(context.Background(), ToggleSMSNotification{
			Name:   name,
			Status: status,
		})

		assert.Equal(t, true, resp.Error == nil)
	})

	tt.Run("negative toggle usecase if vendor is default", func(t *testing.T) {
		name := "ngehe"
		status := false
		mckVendors.EXPECT().Get(name).Return(mckVendor)
		mckVendor.EXPECT().IsDefault().Return(true)
		mckVendor.EXPECT().OnOff(context.Background(), status)
		mckVendor.EXPECT().IsOn().Return(status)

		ucs := Usecase(mckVendors)
		resp := ucs.OnOff(context.Background(), ToggleSMSNotification{
			Name:   name,
			Status: status,
		})

		assert.Equal(t, true, resp.Error != nil)
	})

	tt.Run("negative toggle usecase if vendor doesn't exists", func(t *testing.T) {
		name := "ngehe"
		status := false
		mckVendors.EXPECT().Get(name).Return(nil)
		mckVendor.EXPECT().IsDefault().Return(false)
		mckVendor.EXPECT().OnOff(context.Background(), status)
		mckVendor.EXPECT().IsOn().Return(status)

		ucs := Usecase(mckVendors)
		resp := ucs.OnOff(context.Background(), ToggleSMSNotification{
			Name:   name,
			Status: status,
		})

		assert.Equal(t, true, resp.Error != nil)
	})
}

func Test_list_vendors_usecase(tt *testing.T) {
	ctrl := gomock.NewController(tt)
	mckVendors := mockintf.NewMockVendors(ctrl)
	mckVendor := mockintf.NewMockVendor(ctrl)

	tt.Run("positive get list vendors usecase", func(t *testing.T) {

		mckVendor.EXPECT().IsOn().Return(true)
		mckVendor.EXPECT().IsDefault().Return(true)
		mckVendors.EXPECT().All().Return(map[string]shared.Vendor{
			"ngehe": mckVendor,
		})

		ucs := Usecase(mckVendors)
		resp := ucs.Vendors(context.Background())
		assert.Equal(t, true, resp.Error == nil)
	})
}
