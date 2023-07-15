// Package datastore implements components to fetch data from postgres or other
// sources.
//
// The package contains a cache component, to fetch a value only once.
//
// All components of the datastore are independent and can be put together.
//
// The data flows in both directions. With the Get() method, keys can be
// requested. With the Update() method, values get pushed in the other
// direction. All components implement the two methods and thereby the flow.Flow
// interface.
package datastore
