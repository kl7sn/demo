package dto

type Forwarder struct {
	WebHook Webhook `json:"webhook"`
}

type Webhook struct {
	Url     string                 `json:"url"`
	Headers map[string]string      `json:"headers"`
	Body    map[string]interface{} `json:"body"`
}
