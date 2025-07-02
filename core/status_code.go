package core

const (
	Success = iota

	CommandRunFailed
	WebhookSendFailed
	FileNotFoundError
	TomlDecodeError
	JsonDecodeError
	DescriptionMaxLenError
	TitleMaxLenError
	FooterMaxLenError
	WriteJsonFailed
)
