package filestore

import (
	"encoding/binary"
	"fmt"
)

type Row struct {
	ID   uint32
	Name string
}

func EncodeRow(r Row) []byte {
	nameBytes := []byte(r.Name)
	rowSize := uint32(4 + 4 + len(nameBytes))

	buf := make([]byte, 4+rowSize)
	binary.LittleEndian.PutUint32(buf[0:4], rowSize)
	binary.LittleEndian.PutUint32(buf[4:8], r.ID)
	binary.LittleEndian.PutUint32(buf[8:12], uint32(len(nameBytes)))
	copy(buf[12:], nameBytes)

	return buf
}

func DecodeRow(buf []byte) (Row, error) {
	if len(buf) < 12 {
		return Row{}, fmt.Errorf("buffer too small: %d", len(buf))
	}

	rowSize := binary.LittleEndian.Uint32(buf[0:4])
	if int(rowSize)+4 != len(buf) {
		return Row{}, fmt.Errorf("bad row size: row_size=%d total=%d", rowSize, len(buf))
	}

	id := binary.LittleEndian.Uint32(buf[4:8])
	nameLen := binary.LittleEndian.Uint32(buf[8:12])
	if 12+int(nameLen) != len(buf) {
		return Row{}, fmt.Errorf("bad name length: name_len=%d total=%d", nameLen, len(buf))
	}

	return Row{ID: id, Name: string(buf[12:])}, nil
}
