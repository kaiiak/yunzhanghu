package economy

type RegisterNotice struct {
	Data      string `json:"data"`
	Mess      string `json:"mess"`
	TimeStamp string `json:"time_stamp"`
	Sign      string `json:"sign"`
	SignType  string `json:"sign_type"`
}

type RegisterNoticeData struct {
	OpenId              string `json:"open_id"`
	DealerUserId        string `json:"dealer_user_id"`
	SubmitAt            string `json:"submit_at"`
	RegistedAt          string `json:"registed_at"`
	Status              int    `json:"status"`
	StatusMessage       string `json:"status_message"`
	StatusDetail        int    `json:"status_detail"`
	StatusDetailMessage string `json:"status_detail_message"`
	DealerId            string `json:"dealer_id"`
	BrokerId            string `json:"broker_id"`
	Uscc                string `json:"uscc"`
	IdCard              string `json:"id_card"`
	RealName            string `json:"real_name"`
	Type                int    `json:"type"`
}
