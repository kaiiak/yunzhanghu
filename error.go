package yunzhanghu

type Error struct {
	Code               StatusCode
	Message, RequestId string
}

func (*Error) Error() string {
	return ""
}
