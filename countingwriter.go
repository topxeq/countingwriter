package countingwriter

import (
	"fmt"
	"io"
	"sync"
)

type CountingWriter struct {
	Count      int
	Total      int
	IfPercent  bool
	WritebackI *int
	WritebackS *string
	WritebackA *interface{}
	Callback   func(interface{}) interface{}
	Lock       *sync.Mutex
}

func IfSwitchExistsWhole(argsA []string, switchStrA string) bool {
	if argsA == nil {
		return false
	}

	if len(argsA) < 1 {
		return false
	}

	for _, argT := range argsA {
		if argT == switchStrA {
			return true
		}
	}

	return false
}

func NewCountingWriter(argsA ...interface{}) io.Writer {
	vT := &CountingWriter{}
	argsT := make([]string, 0, len(argsA))
	for _, v := range argsA {
		if v == nil {
			continue
		}

		if nv, ok := v.(string); ok {
			argsT = append(argsT, nv)
			continue
		}

		if nv, ok := v.(int); ok {
			vT.Total = nv
			continue
		}

		if nv, ok := v.(int64); ok {
			vT.Total = int(nv)
			continue
		}

		if nv, ok := v.(*int); ok {
			vT.WritebackI = nv
			continue
		}

		if nv, ok := v.(*string); ok {
			vT.WritebackS = nv
			continue
		}

		if nv, ok := v.(*interface{}); ok {
			vT.WritebackA = nv
			continue
		}

		if nv, ok := v.(func(interface{}) interface{}); ok {
			vT.Callback = nv
			continue
		}

		if nv, ok := v.(*sync.Mutex); ok {
			vT.Lock = nv
			continue
		}
	}

	if IfSwitchExistsWhole(argsT, "-percent") {
		vT.IfPercent = true
	}

	return vT // &CountingWriter{Lock: lockA, Writeback: writebackA, WritebackI: writebackIA, Callback: callbackA, Count: 0}
}

func (pA *CountingWriter) Reset() {
	if pA.Lock != nil {
		pA.Lock.Lock()
	}

	pA.Count = 0

	if pA.Lock != nil {
		pA.Lock.Unlock()
	}
}

// func (pA *CountingWriter) SetCallback(funcA func(interface{}) interface{}) {
// 	pA.Callback = funcA
// }

// func (pA *CountingWriter) SetWriteback(writebackA *string) {
// 	pA.Writeback = writebackA
// }

func (pA *CountingWriter) Write(p []byte) (n int, err error) {
	lenT := len(p)

	if pA.Lock != nil {
		pA.Lock.Lock()
	}

	pA.Count += lenT

	var infoT string

	if pA.IfPercent {
		if pA.Total >= 0 {
			infoT = fmt.Sprintf("%v%%", pA.Count*100/pA.Total)
		} else {
			infoT = "-%"
		}
	} else {
		infoT = fmt.Sprintf("%v", pA.Count)
	}

	if pA.Callback != nil {
		pA.Callback(infoT)
	}

	if pA.WritebackS != nil {
		*(pA.WritebackS) = infoT
	}

	if pA.WritebackI != nil {
		if pA.IfPercent {
			if pA.Total >= 0 {
				*(pA.WritebackI) = pA.Count * 100 / pA.Total
			} else {
				*(pA.WritebackI) = 0
			}
		} else {
			*(pA.WritebackI) = pA.Count
		}
	}

	if pA.WritebackA != nil {
		*(pA.WritebackA) = infoT
	}

	if pA.Lock != nil {
		pA.Lock.Unlock()
	}

	return lenT, nil
}

// type CountingWriter struct {
// 	Count     int
// 	Writeback *string
// 	Callback  func(interface{}) interface{}
// }

// func NewWriter(writebackA *string, callbackA func(interface{}) interface{}) io.Writer {
// 	return &CountingWriter{Writeback: writebackA, Callback: callbackA, Count: 0}
// }

// func (pA *CountingWriter) Reset() {
// 	pA.Count = 0
// }

// func (pA *CountingWriter) SetCallback(funcA func(interface{}) interface{}) {
// 	pA.Callback = funcA
// }

// func (pA *CountingWriter) SetWriteback(writebackA *string) {
// 	pA.Writeback = writebackA
// }

// func (pA *CountingWriter) Write(p []byte) (n int, err error) {
// 	lenT := len(p)
// 	pA.Count += lenT

// 	infoT := fmt.Sprintf("%v", pA.Count)

// 	if pA.Callback != nil {
// 		pA.Callback(infoT)
// 	}

// 	if pA.Writeback != nil {
// 		*(pA.Writeback) = infoT
// 	}

// 	return lenT, nil
// }
