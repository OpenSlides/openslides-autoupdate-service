# ExampleDatastore

Development Datastore server that serves the example data.

This tool can be used as dropin replacement for the
openslides-datastore-service. The autoupdate-service can connect to it, like to
the real datastore-service. It uses the example data.

It listens to stdin for changed data. Changed data has to be given as a
json-object. For example:

```
{"topic/10/meeting_id": 1, "topic/10/title": "ZZ", "topic/10/text": ""}
```

Changed data is sent to the autoupdate-service via redis.

To connect the autoupdate-service to this service, use

```
DATASTORE=service MESSAGING=redis ./autoupdate
```