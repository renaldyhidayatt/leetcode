package models

import "behaviouralpatterns/template/interfaces"

type Otp struct {
	Iotp interfaces.Iotp
}

func (o *Otp) GenAndSendOTP(otpLength int) error {
	otp := o.Iotp.GenRandomOTP(otpLength)
	o.Iotp.SaveOTPCache(otp)
	message := o.Iotp.GetMessage(otp)
	err := o.Iotp.SendNotification(message)
	if err != nil {
		return err
	}
	o.Iotp.PublishMetric()
	return nil
}
