package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type CheckOut struct {
	notify *Notify
}

func NewCheckOut() *CheckOut {
	return &CheckOut{
		notify: NewNotify(),
	}
}

func (ch *CheckOut) Checkout(ctx echo.Context) error {
	msg := "TemplateCode: EDUCATION-E011\n" +
		"ColorCode: C001\n" +
		"StyleCode: M001\n" +
		"GoogleDrive: xxx.com\n" +
		"OtherService: เปลี่ยนสี/ใส่ข้อมูล/psd\n" +
		"\nCustomer:\nName: Nattapon\nEmail: xxx@gmail.com\nLine: @xxx\nPhone: 0930945168\n" +
		"total: 300.0 baht"

	notifyResponse := ch.notify.Send(msg,"C:\\Users\\Nattapon\\Downloads\\qfxl0r5t7zWsuc4XCm2A-o.jpg")
	return ctx.JSON(http.StatusOK, notifyResponse)
}
