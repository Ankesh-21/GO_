package main
import (
	"fmt"
)

type payment struct{}
// function related to the payment struct
func (p payment) pay (amount float32){
	razorpayInstance := RazorPay{}
	razorpayInstance.makePayment(amount)
}

type RazorPay struct{}
func (r RazorPay) makePayment(amount float32){
	fmt.Println("making payment",amount)
}

func main(){
	
}