package models

type Event struct {
	CallbackUrl string       `json:"callback_url"`
	Message     string       `json:"message"`
	Status      string       `json:"status"`
	ExtraFields []ExtraField `json:"extra_fields"`
}

type ExtraField struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
