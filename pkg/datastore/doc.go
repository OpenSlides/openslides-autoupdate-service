// Package datastore fetches the data from postgres or other sources.
//
// The datastore object uses a cache to only request keys once. If a key in the
// cache gets an update via the keychanger, the cache gets updated.
//
// All components of the datastore are independent and can be put together.
//
// The data flows in both directions. With the Get() method, keys can be
// requested. With the Update() method, values get pushed in the other
// direction. All components implement the two methods and thereby the flow.Flow
// interface.
//
// The cache2 caches most of the values. The cache1 only caches value, that are
// only needed by the projector.
package datastore
