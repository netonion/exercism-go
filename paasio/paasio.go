package paasio

import (
	"io"
	"sync"
)

type counter struct {
	sync.Mutex
	n    int64
	nops int
}

func (c *counter) add(n int) {
	c.Lock()
	defer c.Unlock()
	c.n += int64(n)
	c.nops++
}

func (c *counter) count() (int64, int) {
	return c.n, c.nops
}

type readCounter struct {
	reader io.Reader
	counter
}

type writeCounter struct {
	writer io.Writer
	counter
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
	r.add(n)
	return
}

func (r *readCounter) ReadCount() (int64, int) {
	return r.count()
}

func (w *writeCounter) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)
	if err != nil {
		return
	}
	w.add(n)
	return
}

func (w *writeCounter) WriteCount() (int64, int) {
	return w.count()
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
