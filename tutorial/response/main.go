package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Header struct
type Header struct {
	Key, Value string
}

// Status struct
type Status struct {
	Code   int
	Reason string
}

type errWriter struct {
	io.Writer
	err error
}

func (e *errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	_, e.err = e.Writer.Write(buf)
	return 0, nil
}

// WriteResponse func
func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &errWriter{Writer: w}
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)

	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}
	fmt.Fprint(ew, "\r\n")

	io.Copy(ew, body)
	return ew.err
}
func main() {
	var buf bytes.Buffer
	st := Status{Code: 555, Reason: "account not found."}
	headers := []Header{
		{"Content-Type", "application/json"},
	}
	body := strings.NewReader("this is a body")

	err := WriteResponse(&buf, st, headers, body)

	fmt.Printf("buf: \n%s\n", buf.String())
	// fmt.Printf("buf: \n%v\n", buf)
	fmt.Println("error:", err)
}
