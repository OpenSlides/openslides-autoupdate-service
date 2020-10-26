# Features
* Listen on permission changes
* Logging, metrics and traces
* Max keys per request

* Keysbuilder: Create keys at the first call of Keys() and not on create! Test this!


# Think about
* Do not request keys in datastore, when the client is not allowed to see them.