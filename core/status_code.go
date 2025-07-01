package core

const (
	Success = iota

	FileNotFoundError
	TomlDecodeError
	JsonDecodeError
	DescriptionMaxLenError
	TitleMaxLenError
	FooterMaxLenError
)
