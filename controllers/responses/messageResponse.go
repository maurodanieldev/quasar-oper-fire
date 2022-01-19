package responses

type MessageResponse struct {
	Position Position  `json:"position"`
	Message  string `json:"message"`
}

type Position struct {
	X float64  `json:"x"`
	Y float64  `json:"y"`
}