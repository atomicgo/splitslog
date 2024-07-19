package splitslog_test

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"testing"
	"testing/slogtest"

	"atomicgo.dev/splitslog"
)

func TestSplitHandler(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	splitter := splitslog.Splitter{
		slog.LevelDebug: slog.NewJSONHandler(&buf, nil),
		slog.LevelInfo:  slog.NewJSONHandler(&buf, nil),
		slog.LevelWarn:  slog.NewJSONHandler(&buf, nil),
		slog.LevelError: slog.NewJSONHandler(&buf, nil),
	}

	handler := splitslog.NewSplitHandler(splitter)

	results := func() []map[string]any {
		var resultMap []map[string]any

		for _, line := range bytes.Split(buf.Bytes(), []byte{'\n'}) {
			if len(line) == 0 {
				continue
			}

			var m map[string]any
			if err := json.Unmarshal(line, &m); err != nil {
				t.Fatal(err)
			}

			resultMap = append(resultMap, m)
		}

		return resultMap
	}

	err := slogtest.TestHandler(handler, results)
	if err != nil {
		t.Fatal(err)
	}
}
