# Test Cases

This folder contains test cases in yaml format. No need to lern golang to write
tests.

Each file ending on `.yml` or `.yaml` in this folder and all subfolders is used
for testing.

Each file is a test-case object.

Each object can have the folloing fields:

* name: Name of the test case. It is shown if the test case fails.

* db: Key values of the database. A key can either be a FQField, a FQID or a
  name of a collection. The value has to be an according value. For each key, the
  id Field (Collection/X/id) is automaticly created.

* user_id: ID of the user that is used for the check. If the user does not
  exist, it is created. The value 0 means anonymous user. Default is 1337.

* meeting_id: The test user is put into the meeting with this id. Default is 1.

* permission: String of the permission the user has. A Group with ID 1337 is
  created that has only the test user and this permission. If the test user is
  in other groups, he could also have other permissions. This field is ignored
  for anonymous user. The default is no permission.

* action: Write-Action, that is tested.

* payload: User-Payload that is used for the action.

* fqfields: Fields that are checks if the test-user can see them.

* is_allowed: true or false. If a test-case has this field, the IsAllowed is
  callend and the result is comaired to this field.

* can_see: list of fqfields. If a test-case has this field, then
  RestrictFQFields is callend and the result is comaired to this field.

* cases: A list of sub-test-cases. Each sub-test-case can have the same fields.
  If a sub-test-case does not have a field, the parent field is used. If the
  parent- and the sub-test both have a db-field, then both dbs are merged. The
  db of the sub-test has a higher priority. Only `is_allowed` and `restricted`
  are not passed on.