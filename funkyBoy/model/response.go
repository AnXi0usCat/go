package model

// response models
type GoodResponse struct {
	Code uint16 `json:"code"`
	Value map[string]string `json:"value"`
	ScaledValue string `json:"scaled_value"`
}

type BadRequestResponse struct {
	Code uint16 `json:"code"`
	Message string `json:"message"`
}


// response model constructors
func NewGoodResponse(vl map[string]string, sv string) GoodResponse {
	return GoodResponse{Code: 200, Value: vl, ScaledValue: sv}
}

func NewBadRequestResponse(m string) BadRequestResponse {
	return BadRequestResponse{Code:400, Message:m}
}