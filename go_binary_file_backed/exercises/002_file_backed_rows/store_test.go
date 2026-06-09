package filestore

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestAppendAndScanRows(t *testing.T) {
	path := filepath.Join(t.TempDir(), "users.table")

	want := []Row{
		{ID: 1, Name: "Luna"},
		{ID: 2, Name: "Bartosz"},
		{ID: 3, Name: "Al"},
	}

	for _, row := range want {
		if err := AppendRow(path, row); err != nil {
			t.Fatalf("AppendRow(%+v) error: %v", row, err)
		}
	}

	got, err := ScanRows(path)
	if err != nil {
		t.Fatalf("ScanRows error: %v", err)
	}

	if len(got) != len(want) {
		t.Fatalf("got %d rows, want %d: %+v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("row %d got %+v, want %+v", i, got[i], want[i])
		}
	}
}

func TestScanEmptyFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "empty.table")
	if err := os.WriteFile(path, nil, 0o644); err != nil {
		t.Fatal(err)
	}

	rows, err := ScanRows(path)
	if err != nil {
		t.Fatalf("ScanRows empty file error: %v", err)
	}
	if len(rows) != 0 {
		t.Fatalf("got rows %+v, want empty", rows)
	}
}

func TestScanRejectsTruncatedRow(t *testing.T) {
	path := filepath.Join(t.TempDir(), "broken.table")
	good := EncodeRow(Row{ID: 1, Name: "Luna"})
	truncated := good[:len(good)-2]
	if err := os.WriteFile(path, truncated, 0o644); err != nil {
		t.Fatal(err)
	}

	_, err := ScanRows(path)
	if err == nil {
		t.Fatalf("expected an error for truncated row")
	}
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Fatalf("got %v, want io.ErrUnexpectedEOF", err)
	}
}

func TestAppendActuallyAppends(t *testing.T) {
	path := filepath.Join(t.TempDir(), "append.table")

	if err := AppendRow(path, Row{ID: 10, Name: "First"}); err != nil {
		t.Fatal(err)
	}
	firstSize := fileSize(t, path)

	if err := AppendRow(path, Row{ID: 11, Name: "Second"}); err != nil {
		t.Fatal(err)
	}
	secondSize := fileSize(t, path)

	if secondSize <= firstSize {
		t.Fatalf("file did not grow: first=%d second=%d", firstSize, secondSize)
	}
}

func fileSize(t *testing.T, path string) int64 {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	return info.Size()
}
