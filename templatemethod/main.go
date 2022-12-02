package main

func main() {
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	emailOTP := &Email{}
	o = Otp{iOtp: emailOTP}
	o.genAndSendOTP(4)
}
