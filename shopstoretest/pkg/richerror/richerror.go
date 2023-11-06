package richerror

type Kind int

const (
	KindInvalid = iota + 1
	KindForbidden
	KindNotFound
	KindUnexpected
	KindDontHaveCredit
	KindNotUnique
	KindUnauthorized
)

type RichError struct {
	operation    string
	wrappedError error
	message      string
	kind         Kind
	meta         map[string]interface{}
}

func (r RichError) Error() string {

	return r.message
}

func New(op string) RichError {
	newRichError := RichError{
		operation: op,
	}

	return newRichError
}

func (r RichError) WithKind(kind Kind) RichError {
	r.kind = kind

	return r
}

func (r RichError) WithMessage(message string) RichError {
	r.message = message

	return r
}

func (r RichError) WithError(err error) RichError {
	r.wrappedError = err

	return r
}

func (r RichError) WithMeta(meta map[string]interface{}) RichError {
	r.meta = meta

	return r
}

func (r RichError) Kind() Kind {

	return r.kind
}

func (r RichError) Message() string {

	return r.message
}
