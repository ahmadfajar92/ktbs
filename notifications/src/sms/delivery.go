package sms

import (
	"context"
	"net/http"
	"notifications/config"
	"notifications/src/shared"

	"github.com/labstack/echo"
)

type httpDelivery struct {
	cfg     config.Config
	usecase SMSUsecase
}

func HTTPDelivery(cfg config.Config, ucs shared.Usecases) *httpDelivery {
	dlvry := new(httpDelivery)
	dlvry.cfg = cfg
	dlvry.usecase = ucs.Call("sms").(SMSUsecase)
	return dlvry
}

func (dlvry *httpDelivery) Send(c echo.Context) error {
	ctx := context.Background()
	var r SMSNotification

	err := c.Bind(&r)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			shared.FailedResult("error bind request into SMSNotification model", err),
		)
	}

	res := dlvry.usecase.Send(ctx, r)
	return c.JSON(res.Code, res)

}

func (dlvry *httpDelivery) Toggle(c echo.Context) error {
	ctx := context.Background()
	var r ToggleSMSNotification

	err := c.Bind(&r)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			shared.FailedResult("error bind request into ToggleSMSNotification model", err),
		)
	}

	res := dlvry.usecase.OnOff(ctx, r)
	return c.JSON(res.Code, res)
}

func (dlvry *httpDelivery) Vendors(c echo.Context) error {
	ctx := context.Background()
	res := dlvry.usecase.Vendors(ctx)
	return c.JSON(res.Code, res)

}
