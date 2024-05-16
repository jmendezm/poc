package dto

type ResponsePayload struct {
	Header ResponseHeader
	Data   interface{}
	Page   ResponsePage
}

type ResponseHeader struct {
	Code     int
	Method   string
	Resource string
	Message  string
	Errors   []string
}

type ResponsePage struct {
	Offset int
	Limit  int
	Sort   string
	Order  string
}
