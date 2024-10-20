package v1

type Request interface {
	GetBody() interface{}
	GetHeader(key string) string
}

type Response interface {
	SetHeader(key string, value string)
	SetStatusCode(code int)
	WriteBody(body interface{})
}
