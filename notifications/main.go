package main

import (
	"fmt"
	"notifications/config"
	"notifications/src/shared"
	"notifications/src/sms"
	smsvendrs "notifications/src/sms/vendors"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Load env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	// Load config
	cfg := config.Load()
	// Load vendor factory
	smsvndrs := smsvendrs.LoadFromFile("./config/vendors.toml")
	// Load use case factory
	usecases := shared.UsecaseFactory(cfg)

	// Add sms notification usecase
	usecases.Add("sms", sms.Usecase(smsvndrs))

	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	// sms http delivery
	smsHTTPDlvry := sms.HTTPDelivery(cfg, usecases)
	e.POST("/sms/send", smsHTTPDlvry.Send)
	e.POST("/sms/vendor/toggle", smsHTTPDlvry.Toggle)
	e.GET("/sms", smsHTTPDlvry.Vendors)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Get("PORT").(string))))
}
