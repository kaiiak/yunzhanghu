package yunzhanghu

import "fmt"

type Error struct {
	Code      StatusCode
	Message   string
	RequestId string
	ApiName   string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s:%s:%s(request_id:%s)", e.ApiName, e.Code, e.Message, e.RequestId)
}
