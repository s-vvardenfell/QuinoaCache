package utility

import (
	"io"

	"github.com/sirupsen/logrus"
)

func BytesFromReader(r io.Reader) []byte {
	byteValue, err := io.ReadAll(r)
	if err != nil {
		logrus.Fatalf("error reading in BytesFromReader(), %v", err)
	}
	return byteValue
}
