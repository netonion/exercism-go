package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	sync.Mutex
	reader io.Reader
	n      int64
	nops   int
}

type writeCounter struct {
	sync.Mutex
	writer io.Writer
	n      int64
	nops   int
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

func (r *readCounter) Read(p []byte) (n int, err error) {
	n, err = r.reader.Read(p)
	if err != nil {
		return
	}
	r.Lock()
	defer r.Unlock()
	r.n += int64(n)
	r.nops++
	return
}

func (r *readCounter) ReadCount() (n int64, nops int) {
	return r.n, r.nops
}

func (w *writeCounter) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)
	if err != nil {
		return
	}
	w.Lock()
	defer w.Unlock()
	w.n += int64(n)
	w.nops++
	return
}

func (w *writeCounter) WriteCount() (n int64, nops int) {
	return w.n, w.nops
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return readWriteCounter{NewReadCounter(rw), NewWriteCounter(rw)}
}
