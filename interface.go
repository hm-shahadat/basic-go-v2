package main

import "fmt"

type payment struct {
	gateway razorpay
}

func (p payment) makePayment(amount float32) {
	// rezorpayPaymentGW := razorpay{}
	// rezorpayPaymentGW.pay(amount)

	// newStrip := stripe{}
	// newStrip.pay(amount)

	p.gateway.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	fmt.Println("amount razorpay: ", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {

	fmt.Print("amount strip: ", amount)
}

func main() {
	newRazorpay := razorpay{}
	newPayment := payment{
		gateway: newRazorpay,
	}
	newPayment.makePayment(100)

}
