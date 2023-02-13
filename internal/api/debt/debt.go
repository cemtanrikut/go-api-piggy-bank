package api

type DebtModel struct {
	ID          string
	To          string
	Title       string
	Description string
	Total       int
	NewTotal    int
	Payments    []paymentModel
	IsActive    bool
	Completed   bool
}

type paymentModel struct {
	ID     string
	DebtID string
	Date   string
}

// Add debt
func Add(debt DebtModel) {

}

// Get debt by debtID
func Get(debtID string) {

}

// Update debt
func Update(debt DebtModel) {

}

// Delete debt
func Delete(debtID string) {

}

// Add payment
func AddPayment(payment paymentModel) {

}

// Get payment by payment ID
func GetPayment(paymentID string) {

}

// Get Debt's payments
func GetPaymentByDebt(debtID string) {

}

// Delete payment
func DeletePayment(debtID, paymentID string) {

}
