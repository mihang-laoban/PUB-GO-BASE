package def

type RawResponse struct {
	Code int
	Message string
	data interface{}
}

type Response struct {
	Code int
	Message string
	data map[string]interface{}
}