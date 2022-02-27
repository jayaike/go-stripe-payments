package payments

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

type InitialisePaymentRequest struct {
	ItemId string `json:"itemId" binding:"required"`
}

type InitialisePaymentResponse struct {
	ClientSecret string `json:"clientSecret"`
}
