# Features
* Listen on permission changes
* Handle create and delete changes
* Speak with the other services instead of fake services
  * auth
  * message bus (redis) for perm changes and key changes
* Hash values so only new values are send to the client
* Logging, metrics and traces

# Internal
* Find a place for mocks and interfaces
* Run code checks in tests/Dockerfile
* .travis (by calling the tests/Dockerfile?)
