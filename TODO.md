# Features
* Listen on permission changes
* Listen on logout messages
* Handle delete events
* Speak with the other services instead of fake services
  * auth
  * message bus (redis) for perm changes
* Hash values so only new values are send to the client
* Logging, metrics and traces

# Internal
* Find a place for mocks and interfaces
* Run code checks in tests/Dockerfile

# Discuss
* What should Restricter return when an error happens in an open connection (after http 200 was send)