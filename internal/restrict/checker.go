package restrict

//go:generate  sh -c "go run models/main.go > def.go && go fmt def.go"
import (
	"encoding/json"
	"fmt"
	"strings"
)

// OpenSlidesChecker returns the restricter checkers for the openslides models.
func OpenSlidesChecker(perm Permission) map[string]Checker {
	checkers := make(map[string]Checker)
	for k, v := range relationLists {
		// TODO structured fields.
		if strings.Contains(k, "$") {
			continue
		}

		// Generic relation list.
		if v == "*" {
			checkers[k] = &genericRelationList{
				perm: perm,
			}
			continue
		}

		checkers[k] = &relationList{
			perm:  perm,
			model: v,
		}
	}
	return checkers
}

type relationList struct {
	perm  Permission
	model string
}

func (r *relationList) Check(uid int, key string, value json.RawMessage) (json.RawMessage, error) {
	var ids []int
	if err := json.Unmarshal(value, &ids); err != nil {
		return nil, fmt.Errorf("decoding %s: %w", key, err)
	}

	keys := make([]string, len(ids))
	keyToID := make(map[string]int)
	for i, id := range ids {
		keys[i] = fmt.Sprintf("%s/%d", r.model, id)
		keyToID[keys[i]] = id
	}

	allowed, err := r.perm.CheckFQIDs(uid, keys)
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
	perm Permission
}

func (g *genericRelationList) Check(uid int, key string, value json.RawMessage) (json.RawMessage, error) {
	var fqids []string
	if err := json.Unmarshal(value, &fqids); err != nil {
		return nil, fmt.Errorf("decoding %s: %w", key, err)
	}

	keys := make([]string, len(fqids))
	for i, fqid := range fqids {
		keys[i] = fqid
	}

	allowed, err := g.perm.CheckFQIDs(uid, keys)
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
