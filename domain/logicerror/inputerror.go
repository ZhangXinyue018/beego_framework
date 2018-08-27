package logicerror

type InputError error

type InputInvalidError struct {
	InputError
}
