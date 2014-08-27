package errors

func Unwrap(err error) error {
	switch specific := err.(type) {
	case *HereError:
		return specific.error
	case *CauseError:
		return specific.error
	case *TraceError:
		return specific.error
	default:
		return err
	}
}

func Equal(err1, err2 error) bool {
	return Unwrap(err1) == Unwrap(err2)
}
