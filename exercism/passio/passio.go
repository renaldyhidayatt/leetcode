package passio

import (
	"io"
	"sync"
)

type ReadCounter interface {
	io.Reader
	// ReadCount returns the total number of bytes successfully read along
	// with the total number of calls to Read().
	ReadCount() (n int64, nops int)
}

// WriteCounter is an interface describing objects that can be written to,
// and that can count the number of times they have been written to.
//
// If multiple goroutines concurrently call Write, implementations are not
// required to provide any guarantees about interleaving of the Write calls.
// However, implementations MUST guarantee that calls to WriteCount always return
// correct results even in the presence of concurrent Write calls.
type WriteCounter interface {
	io.Writer
	// WriteCount returns the total number of bytes successfully written along
	// with the total number of calls to Write().
	WriteCount() (n int64, nops int)
}

// ReadWriteCounter is the union of ReadCounter and WriteCounter.
//
// All guarantees that apply to either of ReadCounter or WriteCounter
// also apply to ReadWriteCounter.
type ReadWriteCounter interface {
	ReadCounter
	WriteCounter
}

type counter struct {
	bytes uint64
	ops   uint32
	sync.Mutex
}

func (c *counter) increment(n int) {
	c.Lock()

	defer c.Unlock()

	c.bytes += uint64(n)

	c.ops++
}

func (c *counter) count() (int64, int) {
	c.Lock()

	defer c.Unlock()

	return int64(c.bytes), int(c.ops)
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
	writeCounter
	readCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {

	return &writeCounter{writer: writer}

}

func NewReadCounter(reader io.Reader) ReadCounter {

	return &readCounter{reader: reader}

}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {

	return &readWriteCounter{NewWriteCounter(readwriter), NewReadCounter(readwriter)}

}

func (rc *readCounter) Read(p []byte) (int, error) {
	n, err := rc.reader.Read(p)
	rc.increment(n)
	return n, err

}

func (rc *readCounter) ReadCount() (int64, int) {
	return rc.count()
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.writer.Write(p)
	wc.increment(n)
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	return wc.count()
}
