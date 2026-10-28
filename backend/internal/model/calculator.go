package model


type AddRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type AddResponse struct {
	Result float64 `json:"result"`
}
type SubtractRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type SubtractResponse struct {
	Result float64 `json:"result"`
}
type MultiplyRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type MultiplyResponse struct {
	Result float64 `json:"result"`
}
type DivideRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type DivideResponse struct {
	Result float64 `json:"result"`
}