package def

type RawResponse struct {
	Code int
	Message string
	Data interface{}
}

type Response struct {
	Code int
	Message string
	Data map[string]interface{}
}