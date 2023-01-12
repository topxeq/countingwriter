package countingwriter

import (
	"fmt"
	"io"
)

type CountingWriter struct {
	Count     int
	Writeback *string
	Callback  func(interface{}) interface{}
}

func NewWriter(writebackA *string, callbackA func(interface{}) interface{}) io.Writer {
	return &CountingWriter{Writeback: writebackA, Callback: callbackA, Count: 0}
}

func (pA *CountingWriter) Reset() {
	pA.Count = 0
}

func (pA *CountingWriter) SetCallback(funcA func(interface{}) interface{}) {
	pA.Callback = funcA
}

func (pA *CountingWriter) SetWriteback(writebackA *string) {
	pA.Writeback = writebackA
}

func (pA *CountingWriter) Write(p []byte) (n int, err error) {
	lenT := len(p)
	pA.Count += lenT

	infoT := fmt.Sprintf("%v", pA.Count)

	if pA.Callback != nil {
		pA.Callback(infoT)
	}

	if pA.Writeback != nil {
		*(pA.Writeback) = infoT
	}

	return lenT, nil
}
