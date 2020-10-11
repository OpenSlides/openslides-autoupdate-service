# Features
* Listen on permission changes
* Listen on logout messages
* Speak with the auth service
    * Validate token
* Logging, metrics and traces
* Max keys per request


# Think about
* Do not request keys in datastore, when the client is not allowed to see them.