package errors

import "fmt"

// Derive a map of detailed information about an error.
// For HereErrors, map includes a "location" key
// For CauseErrors, map includes one or more "cause" keys
// For TraceErrors, map includes a "trace" key

func Details(err error) map[string]string {
	dets := map[string]string{}

	switch specific := err.(type) {
	case *HereError:
		dets["error"] = specific.error.Error()
		dets["location"] = specific.FullLocation()
	case *CauseError:
		dets["error"] = specific.error.Error()
		dets["cause"] = specific.cause.Error()

		i := 2
		if cause, ok := specific.cause.(*CauseError); ok {
			for {
				dets[fmt.Sprintf("cause%d", i)] = cause.cause.Error()

				if sub, ok := cause.cause.(*CauseError); ok {
					cause = sub
				} else {
					break
				}
			}
		}
	case *TraceError:
		dets["error"] = specific.error.Error()
		dets["trace"] = specific.trace
	default:
		dets["error"] = err.Error()
	}

	return dets
}
