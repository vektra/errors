package errors

// If err is a Here, Cause, or Trace wrapper, return the inner error
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

// Check 2 errors are equal by removing any context wrappers
func Equal(err1, err2 error) bool {
	return Unwrap(err1) == Unwrap(err2)
}
