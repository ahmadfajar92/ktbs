package sms

// SMSNotification struct
type SMSNotification struct {
	To      int    `json:"to"`
	Message string `json:"message"`
}

type ToggleSMSNotification struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
