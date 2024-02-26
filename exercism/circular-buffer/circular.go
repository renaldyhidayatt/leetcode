package circularbuffer

import "errors"

type Buffer struct {
	buff              []byte
	start, end, count int
}

func NewBuffer(size int) *Buffer {
	if size <= 0 {
		return nil
	}

	buff := &Buffer{
		start: 0,
		end:   0,
		buff:  make([]byte, size),
	}

	return buff
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0 {
		return 0, errors.New("empty buffer")
	}

	c := b.buff[b.start]

	b.start++

	b.start %= len(b.buff)

	b.count--

	return c, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.count == len(b.buff) {
		return errors.New("full buffer")
	}

	b.buff[b.end] = c

	b.end++
	b.end %= len(b.buff)

	b.count++

	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.count != len(b.buff) {
		b.WriteByte(c)
	} else {
		b.buff[b.start] = c
		b.start++
		b.start %= len(b.buff)
	}
}

func (b *Buffer) Reset() {
	b.start, b.end, b.count = 0, 0, 0
}
