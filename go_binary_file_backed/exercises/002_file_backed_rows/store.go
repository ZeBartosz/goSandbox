package filestore

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// AppendRow appends one encoded row to path. If the file does not exist, create it.
func AppendRow(path string, row Row) error {
	// TODO 1: open path with os.OpenFile.
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	// TODO 2: encode the row with EncodeRow.
	rowEncoded := EncodeRow(row)

	// TODO 3: write all encoded bytes to the file.
	byteSize, err := file.Write(rowEncoded)
	if err != nil {
		return fmt.Errorf("write row: %w", err)
	}
	if byteSize != len(rowEncoded) {
		return io.ErrShortWrite
	}

	return nil
}

// ScanRows reads every row from path in insertion order.
func ScanRows(path string) ([]Row, error) {
	// TODO 4: open path with os.Open.
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	// TODO 5: create rows := []Row{}.
	rows := []Row{}

	// TODO 6: loop forever and read each row:
	for {
		sizeBuf := make([]byte, 4)
		_, err := io.ReadFull(file, sizeBuf)
		if err == io.EOF {
			return rows, nil
		}
		if err != nil {
			return nil, err
		}

		rowSize := binary.LittleEndian.Uint32(sizeBuf)
		fullRow := make([]byte, 4+rowSize)
		copy(fullRow[0:4], sizeBuf)

		_, err = io.ReadFull(file, fullRow[4:])
		if err != nil {
			return nil, err
		}

		rowDecoded, err := DecodeRow(fullRow)

		if err != nil {
			return nil, fmt.Errorf("decode row: %w", err)
		}

		rows = append(rows, rowDecoded)
	}

}
