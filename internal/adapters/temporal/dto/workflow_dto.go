package dto

type TicketPurchaseRequest struct {
	UserID      string
	ConcertID   string
	SeatNumber  string
	PaymentInfo string
}
type RegistrationData struct {
	IDCard      string `json:"id_card"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}
