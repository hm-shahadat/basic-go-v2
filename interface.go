package main

import "fmt"

type paymentr interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymentr
}

func (p payment) makePayment(amount float32, account string) {
	// rezorpayPaymentGW := razorpay{}
	// rezorpayPaymentGW.pay(amount)

	// newStrip := stripe{}
	// newStrip.pay(amount)

	p.gateway.pay(amount)
	p.gateway.refund(amount, account)

}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	fmt.Println("amount razorpay: ", amount)
}

func (r razorpay) refund(amount float32, account string) {
	fmt.Println("amount razorpay: ", amount, account)
}

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway for testing purpose")
}
func (f fakepayment) refund(amount float32, account string) {
	fmt.Println("making payment using fakepayment 2:", amount, account)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string) {
	fmt.Printf("making payment using paypal 2: %v and account number: %v \n", amount, account)
}

// type stripe struct{}

// func (s stripe) pay(amount float32) {

// 	fmt.Print("amount strip: ", amount)
// }

func main() {
	raz := razorpay{}
	fake := fakepayment{}
	pay := paypal{}

	paypalPayment := payment{

		gateway: pay,
	}
	paypalPayment.makePayment(100, "Shahadat4090")

	fakePayment := payment{

		gateway: fake,
	}
	fakePayment.makePayment(200, "Hossain 4040")

	newRaz := payment{

		gateway: raz,
	}
	newRaz.makePayment(300, "Gazi 2030")
}
