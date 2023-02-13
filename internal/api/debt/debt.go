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
