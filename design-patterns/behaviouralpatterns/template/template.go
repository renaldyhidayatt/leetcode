package template

import (
	"behaviouralpatterns/template/models"
	"fmt"
)

func Tempalte() {
	smsOTP := &models.Sms{}
	o := models.Otp{
		Iotp: smsOTP,
	}
	o.GenAndSendOTP(4)
	fmt.Println("")
	emailOTP := &models.Email{}
	o = models.Otp{
		Iotp: emailOTP,
	}
	o.GenAndSendOTP(4)
}
