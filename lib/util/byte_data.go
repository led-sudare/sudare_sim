package util

import (
	"sync"
)

type EnumByteDataCallback func(i int, c byte)

type ImmutableByteData interface {
	GetAt(i int) byte
	Copy() ByteData
	Length() int

	ConcurrentForEach(callback EnumByteDataCallback)
	ConcurrentForEachAll(callback EnumByteDataCallback)
}

type ByteData interface {
	ImmutableByteData
	SetAt(i int, c byte)
	Clear()
	GetBytes() []byte
	EditSafe(editor func(data ByteData) error) error
}

type ByteDataImpl struct {
	data  []byte
	mutex *sync.Mutex
}

func NewByteData(length int) ByteData {
	data := ByteDataImpl{
		make([]byte, length),
		new(sync.Mutex)}

	return &data
}

func (l *ByteDataImpl) SetAt(i int, c byte) {
	l.data[i] = c
}

func (l *ByteDataImpl) GetAt(i int) byte {
	return l.data[i]
}

func (l *ByteDataImpl) Length() int {
	return len(l.data)
}

func (l *ByteDataImpl) Copy() ByteData {
	cp := NewByteData(len(l.data))
	l.ConcurrentForEachAll(func(i int, c byte) {
		cp.SetAt(i, l.GetAt(i))
	})
	return cp
}

func (l *ByteDataImpl) EditSafe(editableBlock func(editable ByteData) error) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return editableBlock(l)
}

func (l *ByteDataImpl) Clear() {
	ConcurrentEnum(0, l.Length(), func(i int) {
		l.SetAt(i, 0)
	})
}

func (l *ByteDataImpl) GetBytes() []byte {
	return l.data
}

func (l *ByteDataImpl) ConcurrentForEach(callback EnumByteDataCallback) {
	ConcurrentEnum(0, l.Length(), func(i int) {
		c := l.GetAt(i)
		if c != 0 {
			callback(i, c)
		}
	})
}

func (l *ByteDataImpl) ConcurrentForEachAll(callback EnumByteDataCallback) {
	ConcurrentEnum(0, l.Length(), func(i int) {
		c := l.GetAt(i)
		callback(i, c)
	})
}
