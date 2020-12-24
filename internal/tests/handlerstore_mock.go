package tests

import "github.com/OpenSlides/openslides-permission-service/internal/perm"

// HandlerStoreMock implements the collection.HandlerStore interface.
type HandlerStoreMock struct {
	WriteHandler map[string]perm.WriteChecker
	ReadHandler  map[string]perm.ReadeChecker
}

// RegisterReadHandler registers a read handler.
func (m *HandlerStoreMock) RegisterReadHandler(name string, reader perm.ReadeChecker) {
	if m.ReadHandler == nil {
		m.ReadHandler = make(map[string]perm.ReadeChecker)
	}
	m.ReadHandler[name] = reader
}

// RegisterWriteHandler registers a write handler.
func (m *HandlerStoreMock) RegisterWriteHandler(name string, writer perm.WriteChecker) {
	if m.WriteHandler == nil {
		m.WriteHandler = make(map[string]perm.WriteChecker)
	}
	m.WriteHandler[name] = writer

}
