package tests

import "github.com/OpenSlides/openslides-permission-service/internal/perm"

// HandlerStoreMock implements the collection.HandlerStore interface.
type HandlerStoreMock struct {
	WriteHandler map[string]perm.WriteChecker
	ReadHandler  map[string]perm.ReadChecker
}

// RegisterReadHandler registers a read handler.
func (m *HandlerStoreMock) RegisterReadHandler(name string, reader perm.ReadChecker) {
	if m.ReadHandler == nil {
		m.ReadHandler = make(map[string]perm.ReadChecker)
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
