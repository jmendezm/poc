package dto

type RequestPayload struct {
	Header RequestHeader
	Data   interface{}
	Page   RequestPage
}

type RequestHeader struct {
	Method   string
	Resource string
}

type RequestPage struct {
	Offset int
	Limit  int
	Sort   string
	Order  string
}
