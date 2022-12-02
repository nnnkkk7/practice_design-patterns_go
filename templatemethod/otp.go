package main

type IOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct {
	iOtp IOtp
}

func (o *Otp) genAndSendOTP(optLength int) error {
	otp := o.iOtp.genRandomOTP(optLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	if err := o.iOtp.sendNotification(message); err != nil {
		return err
	}
	return nil
}
