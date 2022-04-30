package data

type ActionError struct {
	text string
}

func CreateActionError(text string) *ActionError {
	return &ActionError{
		text,
	}
}

func (err *ActionError) Error() string {
	return err.text
}
