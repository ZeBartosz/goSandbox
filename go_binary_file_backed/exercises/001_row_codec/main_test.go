package rowcodec

import (
	"encoding/hex"
	"testing"
)

func TestEncodeRow(t *testing.T) {
	got := EncodeRow(Row{ID: 1, Name: "Luna"})

	// row_size=12, id=1, name_len=4, name="Luna"
	wantHex := "0c00000001000000040000004c756e61"
	if hex.EncodeToString(got) != wantHex {
		t.Fatalf("encoded bytes mismatch\n got: %s\nwant: %s", hex.EncodeToString(got), wantHex)
	}
}

func TestDecodeRow(t *testing.T) {
	bytes, _ := hex.DecodeString("0c00000001000000040000004c756e61")
	got, err := DecodeRow(bytes)
	if err != nil {
		t.Fatalf("DecodeRow returned error: %v", err)
	}
	if got.ID != 1 || got.Name != "Luna" {
		t.Fatalf("got %+v, want ID=1 Name=Luna", got)
	}
}

func TestRoundTrip(t *testing.T) {
	rows := []Row{
		{ID: 7, Name: "Al"},
		{ID: 42, Name: "Bartosz"},
		{ID: 99, Name: ""},
	}

	for _, row := range rows {
		got, err := DecodeRow(EncodeRow(row))
		if err != nil {
			t.Fatalf("round trip error for %+v: %v", row, err)
		}
		if got != row {
			t.Fatalf("round trip got %+v, want %+v", got, row)
		}
	}
}

func TestDecodeRejectsBadSizes(t *testing.T) {
	bad := []byte{1, 0, 0, 0, 99, 88, 77, 66}
	if _, err := DecodeRow(bad); err == nil {
		t.Fatalf("expected error for malformed row")
	}
}
