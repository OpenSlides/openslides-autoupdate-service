package datastore

import (
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// NewRecorder from datastore.NewRecorder
var NewRecorder = datastore.NewRecorder

// DoesNotExistError is a type alias from datastore.DoesNotExistError
type DoesNotExistError = dsfetch.DoesNotExistError

// Key is a type alias from dskey.Key
type Key = dskey.Key

// KeyFromString from package dskey.
var KeyFromString = dskey.FromString
