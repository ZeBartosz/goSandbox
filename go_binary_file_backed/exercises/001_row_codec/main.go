package rowcodec

import (
	"encoding/binary"
	"fmt"
)

// Row is deliberately tiny: enough to learn binary row storage before LunaSQL
// grows pages, indexes, or a real SQL type system.
type Row struct {
	ID   uint32
	Name string
}

// EncodeRow should produce this binary layout, using little endian integers:
//
//	row_size uint32  // number of bytes after this field
//	id       uint32
//	name_len uint32
//	name     []byte
func EncodeRow(r Row) []byte {
	nameBytes := []byte(r.Name)
	rowSize := uint32(4 + 4 + len(nameBytes)) // id + name_len + name

	buf := make([]byte, 4+rowSize)

	// TODO 1: write rowSize into buf[0:4].
	binary.LittleEndian.PutUint32(buf[0:4], rowSize)
	// TODO 2: write r.ID into buf[4:8].
	binary.LittleEndian.PutUint32(buf[4:8], r.ID)
	// TODO 3: write len(nameBytes) into buf[8:12].
	binary.LittleEndian.PutUint32(buf[8:12], uint32(len(nameBytes)))
	// TODO 4: copy nameBytes into buf[12:].
	copy(buf[12:], nameBytes)

	return buf
}

// DecodeRow reverses EncodeRow.
func DecodeRow(buf []byte) (Row, error) {
	if len(buf) < 12 {
		return Row{}, fmt.Errorf("buffer too small: %d", len(buf))
	}

	// TODO 5: read row_size from buf[0:4]. Validate that int(row_size)+4 == len(buf).
	rowSize := binary.LittleEndian.Uint32(buf[0:4])
	if rowSize+4 != uint32(len(buf)) {
		return Row{}, fmt.Errorf("Bad row size")
	}

	// TODO 6: read id from buf[4:8].
	id := binary.LittleEndian.Uint32(buf[4:8])

	// TODO 7: read name_len from buf[8:12]. Validate that 12+name_len == len(buf).
	nameSize := binary.LittleEndian.Uint32(buf[8:12])
	if 12+nameSize != uint32(len(buf)) {
		return Row{}, fmt.Errorf("Bad name size")
	}

	// TODO 8: convert buf[12:] into a string and return Row{ID: id, Name: name}.
	name := string(buf[12:])

	return Row{
		ID:   id,
		Name: name,
	}, nil
}
