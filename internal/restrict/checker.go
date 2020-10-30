package restrict

//go:generate  sh -c "go run gendef/main.go > def.go && go fmt def.go"
import (
	"encoding/json"
	"fmt"
	"strings"
)

// RelationChecker creates a map of checkers from a map of relation-lists to
// there to-model.
func RelationChecker(relationLists map[string]string, permer Permissioner) map[string]Checker {
	checkers := make(map[string]Checker)
	for k, v := range relationLists {
		// Generic relation list.
		var checker Checker = &relationList{
			permer: permer,
			model:  v,
		}
		if v == "*" {
			checker = &genericRelationList{
				permer: permer,
			}
		}

		// Structured fields.
		if strings.Contains(k, "$") {
			checkers[k] = &templateField{permer: permer}
			k = k[:strings.IndexByte(k, '$')]
		}

		checkers[k] = checker
	}
	return checkers
}

type relationList struct {
	permer Permissioner
	model  string
}

func (r *relationList) Check(uid int, key string, value json.RawMessage) (json.RawMessage, error) {
	var ids []int
	if err := json.Unmarshal(value, &ids); err != nil {
		return nil, fmt.Errorf("decoding %s=%s: %w", key, value, err)
	}

	keys := make([]string, len(ids))
	keyToID := make(map[string]int)
	for i, id := range ids {
		keys[i] = fmt.Sprintf("%s/%d", r.model, id)
		keyToID[keys[i]] = id
	}

	allowed, err := r.permer.RestrictFQIDs(uid, keys)
	if err != nil {
		return nil, fmt.Errorf("check fqids: %w", err)
	}

	allowedIDs := make([]int, 0, len(ids))
	for key, a := range allowed {
		if a {
			allowedIDs = append(allowedIDs, keyToID[key])
		}
	}

	v, err := json.Marshal(allowedIDs)
	if err != nil {
		return nil, fmt.Errorf("encoding restricted ids: %w", err)
	}
	return v, nil
}

type genericRelationList struct {
	permer Permissioner
}

func (g *genericRelationList) Check(uid int, key string, value json.RawMessage) (json.RawMessage, error) {
	var fqids []string
	if err := json.Unmarshal(value, &fqids); err != nil {
		return nil, fmt.Errorf("decoding %s=%s: %w", key, value, err)
	}

	keys := make([]string, len(fqids))
	for i, fqid := range fqids {
		keys[i] = fqid
	}

	allowed, err := g.permer.RestrictFQIDs(uid, keys)
	if err != nil {
		return nil, fmt.Errorf("check fqids: %w", err)
	}

	allowedFQIDs := make([]string, 0, len(fqids))
	for fqid, a := range allowed {
		if a {
			allowedFQIDs = append(allowedFQIDs, fqid)
		}
	}

	v, err := json.Marshal(allowedFQIDs)
	if err != nil {
		return nil, fmt.Errorf("encoding restricted fqids: %w", err)
	}
	return v, nil
}

type templateField struct {
	permer Permissioner
}

func (s *templateField) Check(uid int, key string, value json.RawMessage) (json.RawMessage, error) {
	var replacments []string
	if err := json.Unmarshal(value, &replacments); err != nil {
		return nil, fmt.Errorf("decoding key %s=%s: %w", key, value, err)
	}

	keys := make([]string, len(replacments))
	keyToReplacement := make(map[string]string, len(replacments))
	for i, r := range replacments {
		keys[i] = strings.Replace(key, "$", "$"+r, 1)
		keyToReplacement[keys[i]] = r
	}

	allowed, err := s.permer.RestrictFQFields(uid, keys)
	if err != nil {
		return nil, fmt.Errorf("check generated structured fields: %w", err)
	}

	allowedReplacements := make([]string, 0, len(allowed))
	for key := range allowed {
		if !allowed[key] {
			continue
		}
		allowedReplacements = append(allowedReplacements, keyToReplacement[key])
	}

	v, err := json.Marshal(allowedReplacements)
	if err != nil {
		return nil, fmt.Errorf("encoding restricted template field: %w", err)
	}

	return v, nil
}
