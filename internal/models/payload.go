package models

type Payload struct {
	WebhookCallbackUrl string `json:"webhook_callback_url"`
	Data               Data   `json:"data"`
}

type Data struct {
	Message     string       `json:"message"`
	Status      string       `json:"status"`
	ExtraFields []ExtraField `json:"extra_fields"`
}
