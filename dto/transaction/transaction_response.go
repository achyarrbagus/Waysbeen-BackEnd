package transactiondto

type transactionResponse struct {
	Name    string              `json:"name"`
	Email   string              `json:"email"`
	Phone   int                 `json:"phone"`
	Gender  string              `json:"gender"`
	Address string              `json:"address"`
	Product []CreateRequestCart `json:"product"`
}
