package tests

import "github.com/OpenSlides/openslides-permission-service/internal/collection"

// HandlerStoreMock implements the collection.HandlerStore interface.
type HandlerStoreMock struct {
	WriteHandler map[string]collection.WriteChecker
	ReadHandler  map[string]collection.ReadeChecker
}

// RegisterReadHandler registers a read handler.
func (m *HandlerStoreMock) RegisterReadHandler(name string, reader collection.ReadeChecker) {
	if m.ReadHandler == nil {
		m.ReadHandler = make(map[string]collection.ReadeChecker)
	}
	m.ReadHandler[name] = reader
}

// RegisterWriteHandler registers a write handler.
func (m *HandlerStoreMock) RegisterWriteHandler(name string, writer collection.WriteChecker) {
	if m.WriteHandler == nil {
		m.WriteHandler = make(map[string]collection.WriteChecker)
	}
	m.WriteHandler[name] = writer

}
