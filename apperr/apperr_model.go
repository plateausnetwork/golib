package apperr

type AppErr struct {
	HTTPCode int
	Err      error
	Key      string
	Message  string
}

type Options struct {
	HTTPCode int
	Err      error
	Key      string
	Message  string
}
