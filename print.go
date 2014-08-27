package errors

import (
	"io"
	"os"
)

func Print(err error, out io.Writer) {
	switch specific := err.(type) {
	case *HereError:
		out.Write([]byte(" from: " + specific.FullLocation() + "\n"))
		Print(specific.error, out)
	case *CauseError:
		Print(specific.error, out)
		out.Write([]byte("cause: " + specific.cause.Error() + "\n"))

		if cause, ok := specific.cause.(*CauseError); ok {
			for {
				out.Write([]byte("cause: " + cause.cause.Error() + "\n"))

				if sub, ok := cause.cause.(*CauseError); ok {
					cause = sub
				} else {
					break
				}
			}
		}
	case *TraceError:
		Print(specific.error, out)
		out.Write([]byte("trace:\n" + specific.trace))

	default:
		out.Write([]byte("error: " + err.Error() + "\n"))
	}
}

func Show(err error) {
	Print(err, os.Stderr)
}
