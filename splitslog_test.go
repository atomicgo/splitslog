package splitslog_test

import (
	"atomicgo.dev/splitslog"
	"bytes"
	"encoding/json"
	"log/slog"
	"os"
	"testing"
	"testing/slogtest"
)

func TestSplitHandler(t *testing.T) {
	var buf bytes.Buffer

	splitter := splitslog.Splitter{
		slog.LevelDebug: slog.NewJSONHandler(os.Stdout, nil),
		slog.LevelInfo:  slog.NewJSONHandler(os.Stdout, nil),
		slog.LevelWarn:  slog.NewJSONHandler(os.Stderr, nil),
		slog.LevelError: slog.NewJSONHandler(os.Stderr, nil),
	}

	h := splitslog.NewSplitHandler(splitter)

	results := func() []map[string]any {
		var ms []map[string]any

		for _, line := range bytes.Split(buf.Bytes(), []byte{'\n'}) {
			if len(line) == 0 {
				continue
			}

			var m map[string]any
			if err := json.Unmarshal(line, &m); err != nil {
				t.Fatal(err)
			}

			ms = append(ms, m)
		}

		return ms
	}
	err := slogtest.TestHandler(h, results)

	if err != nil {
		t.Fatal(err)
	}
}
