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
	// Use flags: os.O_CREATE | os.O_WRONLY | os.O_APPEND
	// Use permissions: 0o644
	// Remember to Close the file.

	// TODO 2: encode the row with EncodeRow.

	// TODO 3: write all encoded bytes to the file.
	// Hint: file.Write(data) returns (n int, err error).
	// Validate n == len(data), otherwise return an error.

	_ = os.OpenFile
	_ = row
	return fmt.Errorf("not implemented")
}

// ScanRows reads every row from path in insertion order.
func ScanRows(path string) ([]Row, error) {
	// TODO 4: open path with os.Open.
	// Remember to Close the file.

	// TODO 5: create rows := []Row{}.

	// TODO 6: loop forever and read each row:
	//   a. make sizeBuf := make([]byte, 4)
	//   b. read exactly 4 bytes with io.ReadFull(file, sizeBuf)
	//   c. if err == io.EOF, return rows, nil
	//   d. if err != nil, return nil, err
	//   e. rowSize := binary.LittleEndian.Uint32(sizeBuf)
	//   f. fullRow := make([]byte, 4+rowSize)
	//   g. copy(fullRow[0:4], sizeBuf)
	//   h. read exactly rowSize bytes into fullRow[4:] with io.ReadFull
	//   i. decode with DecodeRow(fullRow)
	//   j. append decoded row to rows

	_ = path
	_ = binary.LittleEndian
	_ = io.ReadFull
	_ = os.Open
	return nil, fmt.Errorf("not implemented")
}
