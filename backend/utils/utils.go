package utils

type Message struct {
	Status         int    `json:"status"`
	Err            string `json:"err"`
	Message        string `json:"message"`
	AdditionalInfo string `json:"addInfo"`
}

func NewMessage(status int, err error, message, addinfo string) *Message {
	return &Message{
		Status:         status,
		Err:            ErrorToString(err),
		Message:        message,
		AdditionalInfo: addinfo,
	}
}

func ErrorToString(err error) string {
	if err != nil {
		return err.Error()
	}
	return "null"
}
