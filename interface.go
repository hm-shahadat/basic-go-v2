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

//--------------------------------------------------------------------
type try struct{}

func (t try) tryMe(amount float32) {
	fortry := try2{}
	fortry := tryMe(amount)
}

type try2 struct{}

func (t try2) tryMe2(amount float32) {
	fmt.Print("amount2: ", amount)
}

func main() {

	newPayment := payment{}
	newPayment.makePayment(100)
	//-----------------------------------------------------------

	newTry := try{}
	newTry.tryMe(100)
}
