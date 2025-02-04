package utils

import (
	"encoding/binary"
	"time"
)

func TimeToByte(input time.Time) ([]byte, error) {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(input.UnixNano()))
	return []byte{}, nil
}
