package datastore

import (
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// NewRecorder from datastore.NewRecorder
var NewRecorder = datastore.NewRecorder

// DoesNotExistError is a type alias from datastore.DoesNotExistError
type DoesNotExistError = dsfetch.DoesNotExistError

// Key is a type alias from datastore.Key
type Key = datastore.Key

// KeyFromString from package datastore.
var KeyFromString = datastore.KeyFromString
