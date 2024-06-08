package errors

type UserError struct {
	Code       string   `json:"code"`
	Message    string   `json:"message"`
	Trace      []string `json:"trace"`
	codeStatus int
}

func NewUserError(code string, codeStatus int) UserError {
	return UserError{
		Code:       code,
		Message:    code,
		codeStatus: codeStatus,
		Trace:      []string{},
	}
}

func NewUserErrorWithError(code string, codeStatus int, err error) UserError {
	return UserError{
		Code:       code,
		Message:    code,
		codeStatus: codeStatus,
		Trace: []string{
			err.Error(),
		},
	}
}

func (error UserError) GetStatusServer() int {
	return error.codeStatus
}

func (error UserError) Error() string {
	return error.Message
}
