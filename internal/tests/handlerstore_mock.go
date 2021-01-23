package tests

import "github.com/OpenSlides/openslides-permission-service/internal/perm"

// HandlerStoreMock implements the collection.HandlerStore interface.
type HandlerStoreMock struct {
	WriteHandler map[string]perm.ActionChecker
	ReadHandler  map[string]perm.RestricterChecker
}

// RegisterRestricter registers a read handler.
func (m *HandlerStoreMock) RegisterRestricter(name string, reader perm.RestricterChecker) {
	if m.ReadHandler == nil {
		m.ReadHandler = make(map[string]perm.RestricterChecker)
	}
	m.ReadHandler[name] = reader
}

// RegisterAction registers a write handler.
func (m *HandlerStoreMock) RegisterAction(name string, writer perm.ActionChecker) {
	if m.WriteHandler == nil {
		m.WriteHandler = make(map[string]perm.ActionChecker)
	}
	m.WriteHandler[name] = writer

}
