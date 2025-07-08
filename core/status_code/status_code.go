package status_code

const (
	Success = iota

	CommandRunFailed
	WebhookSendFailed
	FileNotFoundError
	TomlDecodeError
	JsonDecodeError
	MaxLengthEmbedError
	WriteJsonFailed
	WriteFileFailed
	CreateFileFailed
	FlushFileFailed
	RunProgramFailed
	FieldNotProvided
)
