package core

const (
	Success = iota

	WebhookSendFailed

	FileNotFoundError
	TomlDecodeError
	JsonDecodeError
	DescriptionMaxLenError
	TitleMaxLenError
	FooterMaxLenError
)
