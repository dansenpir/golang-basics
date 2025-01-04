package main

import (
	"Basics/utils"
	"fmt"
)

// PaymentMethod interface
type PaymentMethod interface {
	ProcessPayment(amount float64) bool
}

// CreditCard struct
type CreditCard struct {
	Number string
	CVV    string
}

// ProcessPayment method for CreditCard
func (c CreditCard) ProcessPayment(amount float64) bool {
	fmt.Printf("Processing credit card payment of $%.2f\n", amount)
	// Simulating payment processing
	return true
}

// PayPal struct
type PayPal struct {
	Email string
}

// ProcessPayment method for PayPal
func (p PayPal) ProcessPayment(amount float64) bool {
	fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
	// Simulating payment processing
	return true
}

// BankTransfer struct
type BankTransfer struct {
	AccountNumber string
	BankCode      string
}

// ProcessPayment method for BankTransfer
func (b BankTransfer) ProcessPayment(amount float64) bool {
	fmt.Printf("Processing bank transfer of $%.2f\n", amount)
	// Simulating payment processing
	return true
}

// ProcessOrder function that accepts any PaymentMethod
func ProcessOrder(pm PaymentMethod, amount float64) bool {\
	return pm.ProcessPayment(amount)
}

func main() {
	utils.PrintSection("Payment Processing System")

	// Create different payment methods
	creditCard := CreditCard{Number: "1234-5678-9012-3456", CVV: "123"}
	payPal := PayPal{Email: "user@example.com"}
	bankTransfer := BankTransfer{AccountNumber: "987654321", BankCode: "ABCBANK"}

	// Process orders with different payment methods
	order1Amount := 100.50
	order2Amount := 75.25
	order3Amount := 200.00

	fmt.Println("Order 1:")
	ProcessOrder(creditCard, order1Amount)

	fmt.Println("\nOrder 2:")
	ProcessOrder(payPal, order2Amount)

	fmt.Println("\nOrder 3:")
	ProcessOrder(bankTransfer, order3Amount)

	utils.PrintSection("Type Assertions")

	// Demonstrate type assertions
	var paymentMethod PaymentMethod = creditCard

	if cc, ok := paymentMethod.(CreditCard); ok {
		fmt.Printf("This is a credit card payment with number %s\n", cc.Number)
	}

	if _, ok := paymentMethod.(PayPal); ok {
		fmt.Println("This is a PayPal payment")
	} else {
		fmt.Println("This is not a PayPal payment")
	}

	utils.PrintSection("Type Switches")

	// Demonstrate type switches
	identifyPaymentMethod(creditCard)
	identifyPaymentMethod(payPal)
	identifyPaymentMethod(bankTransfer)
}

func identifyPaymentMethod(pm PaymentMethod) {
	switch v := pm.(type) {
	case CreditCard:
		fmt.Printf("Credit Card payment with number %s\n", v.Number)
	case PayPal:
		fmt.Printf("PayPal payment with email %s\n", v.Email)
	case BankTransfer:
		fmt.Printf("Bank Transfer payment with account number %s\n", v.AccountNumber)
	default:
		fmt.Println("Unknown payment method")
	}
}
