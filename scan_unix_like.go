package ctc

import (
	"bytes"
	"io"
)

var pre = []byte("\x1b[0")

func ScanUnixLike(r io.Reader) (Color, []byte, bool) {
	buf := [20]byte{}
	off := 0
	c := Color(0)
	for _, v := range pre {
		r.Read(buf[off : off+1])
		if buf[off] != v {
			return c, buf[:off+1], false
		}
		off++
	}

	fb := byte(0)
loop:
	for {
		if off >= 16 {
			return c, buf[:off], false
		}
		r.Read(buf[off : off+1])
		switch buf[off] {
		default:
			return c, buf[:off+1], false
		case 'm':
			return c, buf[:off+1], true
		case ';':
		}
		off++

		r.Read(buf[off : off+1])
		switch buf[off] {
		default:
			return c, buf[:off+1], false
		case '3', '4':
			fb = buf[off] - '3'
		case '1':
			if fb == 0 {
				c |= ForegroundBright
			} else {
				c |= BackgroundBright
			}
			continue loop
		}
		off++

		r.Read(buf[off : off+1])
		switch buf[off] {
		default:
			return c, buf[:off+1], false
		case '0', '1', '2', '3', '4', '5', '6', '7':
			if fb == 0 {
				c |= applyForeground
				c |= Color(buf[off] - '0')
			} else {
				c |= applyBackground
				c |= Color(buf[off]-'0') << 4
			}
		}
		off++
	}
}

func PipeUnixLike(w io.Writer) io.Writer {
	buf := newSyncBuffer()
	go func() {
		for {
			c, b, ok := ScanUnixLike(buf)
			if ok {
				c.Apply()
			} else {
				w.Write(b)
			}
		}
	}()
	return buf
}

type syncBuffer struct {
	buf bytes.Buffer
	s   chan struct{}
}

func newSyncBuffer() *syncBuffer {
	return &syncBuffer{
		s: make(chan struct{}),
	}
}

func (b *syncBuffer) Read(p []byte) (n int, err error) {
	n, err = b.buf.Read(p)
	if b.buf.Len() == 0 {
		<-b.s
	}
	return
}

func (b *syncBuffer) Write(p []byte) (n int, err error) {
	n, err = b.buf.Write(p)
	if b.buf.Len() != 0 {
		b.s <- struct{}{}
	}
	return
}
