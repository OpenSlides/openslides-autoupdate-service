# Permission package

The permission package was its own git repository. It was merged into this
repository because it was the only place it was used and this makes changed easier.

The current code is only temporary. The code should be merged into the restict
package to get one api.

There are test cases as yaml files. To run only this test, you can go into this
repository and run:

```
sh test.sh
```

The format of the yaml files is as follows:

Each file ending on `.yml` or `.yaml` in this directory and all subdirectories
is used for testing.

Each file is a test-case object.

Reading FQFields: To test whether the user can see FQFields or not.

  * `fqfields`: A list of fields to be tested.

  * `fqids`: List of fqids. Adds all fields of the objects to the fqfields
    attribute.

  * `can_see`: A list of fields that the user is expected to see. Omit fields
    that the user can not see.

  * `can_not_see`: A shortcut for saying `can_see` everything from `fqfields`
    expect thes once.

Both types can be combined.

Each test case object can have the following additional keywords to define the
test parameters:

* `name`: Name of the test case. It is shown if the test fails.

* `db`: Content of the database. A key can either be a FQField, a FQID or a
  Collection. The value has to be an according value. For each key, the `id`
  field (Collection/X/id) is automaticly created.

  Example:

```yaml
  db:
    meeting/1/name: Test Meeting
    user/1:
      first_name: Hugo
      last_name: Boss
    tag:
      1:
        name: Green
      2:
        name: Yellow
```

* `user_id`: ID of the user that is used for the test. If the user does not
  exist, it is created. The value 0 means anonymous user. Default is 1337.

* `meeting_id`: The test user is put into the meeting with this id. Default is
  `1`.

* `permission`: String of the permission the user has. A Group with ID 1337 is
  created that has only the test user and this permission. If the test user is
  in other groups, he could also have other permissions. This field is ignored
  for anonymous user. The default is no permission.

* `cases`: A list of sub test cases. Each sub test case can have the mentioned
  keywords.

  If a sub test case does not have a keyword, the parent field is used. If the
  parent and the sub test both have a `db`, then both will be merged. The `db`
  of the sub test has a higher priority. Only `is_allowed` and `can_see` are not
  passed on.
