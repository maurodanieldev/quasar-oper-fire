package responses


type HTTPError struct {
Code     int         `json:"code"`
Message  interface{} `json:"message"`
Internal string       `json:"error"`
}
