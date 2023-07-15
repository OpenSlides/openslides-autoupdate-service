package dsmock

// // MockDatastore implements the autoupdate.Datastore interface.
// type MockDatastore struct {
// 	*datastore.Datastore
// 	flow    *Flow
// 	counter *Counter
// 	err     error
// }

// // NewMockDatastore create a MockDatastore with data.
// //
// // It is a wrapper around the datastore.Datastore object.
// func NewMockDatastore(data map[dskey.Key][]byte) (*MockDatastore, func(context.Context, func(error))) {
// 	source := NewFlow(data, NewCounter)
// 	rawDS, bg, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(source))
// 	if err != nil {
// 		panic(err)
// 	}

// 	ds := &MockDatastore{
// 		flow:      source,
// 		Datastore: rawDS,
// 	}

// 	ds.counter = source.Middlewares()[0].(*Counter)

// 	return ds, bg
// }

// // Get calls the Get() method of the datastore.
// func (d *MockDatastore) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
// 	if d.err != nil {
// 		return nil, d.err
// 	}

// 	return d.Datastore.Get(ctx, keys...)
// }

// // InjectError lets the next calls to Get() return the injected error.
// func (d *MockDatastore) InjectError(err error) {
// 	d.err = err
// }

// // Requests returns a list of all requested keys.
// func (d *MockDatastore) Requests() [][]dskey.Key {
// 	return d.counter.Requests()
// }

// // ResetRequests resets the list returned by Requests().
// func (d *MockDatastore) ResetRequests() {
// 	d.counter.Reset()
// }

// // KeysRequested returns true, if all given keys where requested.
// func (d *MockDatastore) KeysRequested(keys ...dskey.Key) bool {
// 	requestedKeys := make(map[dskey.Key]bool)
// 	for _, l := range d.Requests() {
// 		for _, k := range l {
// 			requestedKeys[k] = true
// 		}
// 	}

// 	for _, k := range keys {
// 		if !requestedKeys[k] {
// 			return false
// 		}
// 	}
// 	return true
// }

// // Send updates the data.
// //
// // This method is unblocking. If you want to fetch data afterwards, make sure to
// // block until data is processed. For example with RegisterChanceListener.
// func (d *MockDatastore) Send(data map[dskey.Key][]byte) {
// 	d.flow.Send(data)
// }

// // Update implements the datastore.Updater interface.
// func (d *MockDatastore) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
// 	d.flow.Update(ctx, updateFn)
// }
