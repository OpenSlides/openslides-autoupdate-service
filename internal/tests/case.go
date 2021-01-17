package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/pkg/permission"
	"gopkg.in/yaml.v3"
)

//go:generate  sh -c "go run gen_fields/main.go > fields.go && go fmt fields.go"

// Case object for testing.
type Case struct {
	Name     string
	DB       map[string]interface{}
	FQFields []string
	FQIDs    []string

	UserID     *int `yaml:"user_id"`
	userID     int
	MeetingID  int `yaml:"meeting_id"`
	Permission string

	Payload map[string]interface{}
	Action  string

	IsAllowed *bool    `yaml:"is_allowed"`
	CanSee    []string `yaml:"can_see"` // TODO: fix nil != undefined
	CanNotSee []string `yaml:"can_not_see"`

	Cases []*Case
}

func (c *Case) walk(f func(*Case)) {
	f(c)
	for _, s := range c.Cases {
		s.walk(f)
	}
}

func (c *Case) test(t *testing.T) {
	if onlyTest := os.Getenv("TEST_CASE"); onlyTest != "" {
		onlyTest = strings.TrimPrefix(onlyTest, "TestCases/")
		if !strings.HasPrefix(c.Name, onlyTest) {
			return
		}
	}
	if c.IsAllowed != nil {
		c.testWrite(t)
	}
	if c.CanSee != nil || c.CanNotSee != nil {
		c.testRead(t)
	}
}

func (c *Case) loadDB() (map[string]json.RawMessage, error) {
	data := make(map[string]json.RawMessage)
	for dbKey, dbValue := range c.DB {
		parts := strings.Split(dbKey, "/")
		switch len(parts) {
		case 1:
			map1, ok := dbValue.(map[interface{}]interface{})
			if !ok {
				return nil, fmt.Errorf("invalid type in db key %s: %T", dbKey, dbValue)
			}
			for rawID, rawObject := range map1 {
				id, ok := rawID.(int)
				if !ok {
					return nil, fmt.Errorf("invalid id type: got %T expected int", rawID)
				}
				field, ok := rawObject.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("invalid object type: got %T, expected map[string]interface{}", rawObject)
				}

				for fieldName, fieldValue := range field {
					fqfield := fmt.Sprintf("%s/%d/%s", dbKey, id, fieldName)
					bs, err := json.Marshal(fieldValue)
					if err != nil {
						return nil, fmt.Errorf("creating test db. Key %s: %w", fqfield, err)
					}
					data[fqfield] = bs
				}

				idField := fmt.Sprintf("%s/%d/id", dbKey, id)
				data[idField] = json.RawMessage(strconv.Itoa(id))
			}

		case 2:
			field, ok := dbValue.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("invalid object type: got %T, expected map[string]interface{}", dbValue)
			}

			for fieldName, fieldValue := range field {
				fqfield := fmt.Sprintf("%s/%s/%s", parts[0], parts[1], fieldName)
				bs, err := json.Marshal(fieldValue)
				if err != nil {
					return nil, fmt.Errorf("creating test db. Key %s: %w", fqfield, err)
				}
				data[fqfield] = bs
			}

			idField := fmt.Sprintf("%s/%s/id", parts[0], parts[1])
			data[idField] = []byte(parts[1])

		case 3:
			bs, err := json.Marshal(dbValue)
			if err != nil {
				return nil, fmt.Errorf("creating test db. Key %s: %w", dbKey, err)
			}
			data[dbKey] = bs

			idField := fmt.Sprintf("%s/%s/id", parts[0], parts[1])
			data[idField] = []byte(parts[1])
		default:
			return nil, fmt.Errorf("invalid db key %s", dbKey)
		}

	}

	return data, nil
}

func (c *Case) service() (*permission.Permission, error) {
	data, err := c.loadDB()
	if err != nil {
		return nil, fmt.Errorf("loading database: %w", err)
	}

	if c.userID != 0 {
		// Make sure the user does exists.
		userFQID := fmt.Sprintf("user/%d", c.userID)
		if data[userFQID+"/id"] == nil {
			data[userFQID+"/id"] = []byte(strconv.Itoa(c.userID))
		}

		// Make sure, the user is in the meeting.
		meetingFQID := fmt.Sprintf("meeting/%d", c.MeetingID)
		data[meetingFQID+"/user_ids"] = jsonAddInt(data[meetingFQID+"/user_ids"], c.userID)

		// Create group with the user and the given permissions.
		data["group/1337/id"] = []byte("1337")
		data[meetingFQID+"/group_ids"] = []byte("[1337]")
		data["group/1337/user_ids"] = []byte(fmt.Sprintf("[%d]", c.userID))
		f := fmt.Sprintf("user/%d/group_$%d_ids", c.userID, c.MeetingID)
		data[f] = jsonAddInt(data[f], 1337)
		data["group/1337/meeting_id"] = []byte(strconv.Itoa(c.MeetingID))
		if c.Permission != "" {
			data["group/1337/permissions"] = []byte(fmt.Sprintf(`["%s"]`, c.Permission))
		}
	}

	return permission.New(&dataProvider{data}), nil
}

func (c *Case) testWrite(t *testing.T) {
	p, err := c.service()
	if err != nil {
		t.Fatalf("Can not create permission service: %v", err)
	}

	payload := make(map[string]json.RawMessage, len(c.Payload))
	for k, v := range c.Payload {
		bs, err := json.Marshal(v)
		if err != nil {
			t.Fatalf("Invalid Payload: %v", err)
		}
		payload[k] = bs

	}
	dataList := []map[string]json.RawMessage{payload}

	got, err := p.IsAllowed(context.Background(), c.Action, c.userID, dataList)
	if err != nil {
		t.Fatalf("IsAllowed retuend unexpected error: %v", err)
	}

	if got != *c.IsAllowed {
		t.Errorf("Got %t, expected %t", got, *c.IsAllowed)
	}
}

func (c *Case) testRead(t *testing.T) {
	p, err := c.service()
	if err != nil {
		t.Fatalf("Can not create permission service: %v", err)
	}

	for _, fqid := range c.FQIDs {
		c.FQFields = append(c.FQFields, expandFQID(fqid)...)
	}

	got, err := p.RestrictFQFields(context.Background(), c.userID, c.FQFields)
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}

	if c.CanSee != nil {
		canSee := expandFQIDList(c.CanSee)
		canNotSee := sliceSub(c.FQFields, canSee)
		c.readTestResult(t, got, canSee, canNotSee)
	}

	if c.CanNotSee != nil {
		canNotSee := expandFQIDList(c.CanNotSee)
		canSee := sliceSub(c.FQFields, canNotSee)
		c.readTestResult(t, got, canSee, canNotSee)
	}
}

func (c *Case) readTestResult(t *testing.T, got map[string]bool, canSee, canNotSee []string) {
	if len(got) != len(canSee) {
		t.Errorf("Got %d fields, expected %d", len(got), len(canSee))
	}

	for _, f := range canSee {
		if !got[f] {
			t.Errorf("Got field %s", f)
		}
	}

	gotnot := sliceSubSet(c.FQFields, got)
	set := make(map[string]bool, len(gotnot))
	for _, v := range gotnot {
		set[v] = true
	}

	for _, f := range canNotSee {
		if !set[f] {
			t.Errorf("Did not get field %s", f)
		}
	}
}

func (c *Case) initSub() {
	for i, s := range c.Cases {
		name := s.Name
		if name == "" {
			name = fmt.Sprintf("case_%d", i)
		}
		name = strings.ReplaceAll(name, " ", "_")
		s.Name = c.Name + ":" + name

		db := make(map[string]interface{})
		for k, v := range c.DB {
			db[k] = v
		}
		for k, v := range s.DB {
			db[k] = v
		}
		s.DB = db

		if s.FQFields == nil {
			s.FQFields = c.FQFields
		}
		if s.FQIDs == nil {
			s.FQIDs = c.FQIDs
		}

		s.userID = c.userID
		if s.UserID != nil {
			s.userID = *s.UserID
		}
		if s.MeetingID == 0 {
			s.MeetingID = c.MeetingID
		}
		if s.Permission == "" {
			s.Permission = c.Permission
		}
		if s.Payload == nil {
			s.Payload = c.Payload
		}
		if s.Action == "" {
			s.Action = c.Action
		}

		s.initSub()
	}
}

func walk(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() == false && (strings.HasSuffix(info.Name(), ".yaml") || strings.HasSuffix(info.Name(), ".yml")) {
			files = append(files, path)
		}
		return nil

	})
	if err != nil {
		return nil, fmt.Errorf("walking %s: %w", path, err)
	}
	return files, nil
}

func loadFile(path string) (*Case, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	var c Case
	if err := yaml.NewDecoder(f).Decode(&c); err != nil {
		return nil, err
	}

	name := strings.TrimPrefix(path, "../../tests/")
	if c.Name != "" {
		name += ":" + c.Name
	}
	c.Name = name

	if c.MeetingID == 0 {
		c.MeetingID = 1
	}
	c.userID = 1337
	if c.UserID != nil {
		c.userID = *c.UserID
	}

	c.initSub()

	return &c, nil
}

// defaultInt returns returns the given value or d, if value == 0
func defaultInt(value int, d int) int {
	if value == 0 {
		return d
	}
	return value
}

// jsonAddInt adds the given int to the encoded json list.
//
// If the value exists in the list, the list is returned unchanged.
func jsonAddInt(list json.RawMessage, value int) json.RawMessage {
	var decoded []int
	if list != nil {
		json.Unmarshal(list, &decoded)
	}

	for _, i := range decoded {
		if i == value {
			return list
		}
	}

	decoded = append(decoded, value)
	list, _ = json.Marshal(decoded)
	return list
}

func sliceSub(slice, sub []string) []string {
	set := make(map[string]bool, len(sub))
	for _, v := range sub {
		set[v] = true
	}
	return sliceSubSet(slice, set)
}

func sliceSubSet(slice []string, sub map[string]bool) []string {
	var reduced []string
	for _, v := range slice {
		if !sub[v] {
			reduced = append(reduced, v)
		}
	}
	return reduced
}

// expandFQID returns all fqfields for an fqid.
func expandFQID(fqid string) []string {
	var fqfields []string
	parts := strings.Split(fqid, "/")
	for _, field := range collectionFields[parts[0]] {
		fqfields = append(fqfields, fmt.Sprintf("%s/%s/%s", parts[0], parts[1], field))
	}
	return fqfields
}

// expandFQIDList calls expandFQID on every value in the list. Values that are
// not an fqid are added to the output as it.
func expandFQIDList(values []string) []string {
	var expanded []string
	for _, value := range values {
		if strings.Count(value, "/") == 1 {
			expanded = append(expanded, expandFQID(value)...)
			continue
		}
		expanded = append(expanded, value)
	}
	return expanded
}
