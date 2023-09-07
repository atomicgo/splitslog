/*
Package splitslog provides a handler that splits log records to different handlers based on their level.

The most common use case is to split logs to stdout and stderr based on their level.

	func main() {
		splitter := splitslog.Splitter{
			// Debug and info messages are printed to stdout.
			slog.LevelDebug: slog.NewJSONHandler(os.Stdout, nil),
			slog.LevelInfo:  slog.NewJSONHandler(os.Stdout, nil),
			// Warn and error messages are printed to stderr.
			slog.LevelWarn:  slog.NewJSONHandler(os.Stderr, nil),
			slog.LevelError: slog.NewJSONHandler(os.Stderr, nil),
		}

		handler := splitslog.NewSplitHandler(splitter)
		logger := slog.New(handler)

		logger.Info("info message prints to stdout")
		logger.Error("error message prints to stderr")

		// Output:
		// stdout: {"time":"2023-09-07T16:56:22.563817+02:00","level":"INFO","msg":"info message prints to stdout"}
		// stderr: {"time":"2023-09-07T16:56:22.564103+02:00","level":"ERROR","msg":"error message prints to stderr"}
	}
*/
package splitslog
