package main

import "fmt"

type payment struct{}

func (p payment) makePayment(amount float32) {
	rezorpayPaymentGW := razorpay{}
	rezorpayPaymentGW.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("amount", amount)
}

func main() {

	newPayment := payment{}
	newPayment.makePayment(100)

}
