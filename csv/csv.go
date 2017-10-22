package csv

import (
	"bytes"
	"io"
	"reflect"
	"unsafe"
)

type Encoder struct {
	// contains writer and columns number
	w   io.Writer
	e   error
	le  int
	del string
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer, l int, delimiter string) *Encoder {
	return &Encoder{w: w, le: l, del: delimiter}
}

func (enc *Encoder) Encode(v interface{}) error {

	val := reflect.Indirect(reflect.ValueOf(v))
	var buffer bytes.Buffer

	for i := 0; i < enc.le; i++ {
		field := val.Type().Field(i).Name
		v := val.FieldByName(field)
		p := v.InterfaceData()
		b := printPtrString(p[1])
		// Get the value from the pointer add
		if i == 0 {
			_, err := buffer.WriteString("\n")
			if err != nil {
				return err
			}
		}

		if i > 0 {
			_, err := buffer.WriteString(enc.del)
			if err != nil {
				return err
			}
		}

		_, err := buffer.WriteString(b)
		if err != nil {
			return err
		}
	}

	if _, err := enc.w.Write(buffer.Bytes()); err != nil {
		enc.e = err
	}
	return nil

}

// Return string from pointer
func printPtrString(ptr uintptr) string {
	p := unsafe.Pointer(ptr)
	// Return empty string for nil pointer
	if p == nil {
		return ""
	} else {
		return *(*string)(p)
	}

}
