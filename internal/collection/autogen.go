package collection

//go:generate  sh -c "go run build_autogen_def/main.go > autogen_def.go && go fmt autogen_def.go"

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// Autogen adds routs for all simple permission cases.
type Autogen struct {
	dp dataprovider.DataProvider
}

// NewAutogen initializes an Autogen object.
func NewAutogen(dp dataprovider.DataProvider) *Autogen {
	return &Autogen{
		dp: dp,
	}
}

// Connect connects the simple routes.
func (a *Autogen) Connect(s perm.HandlerStore) {
	for route, perm := range autogenDef {
		parts := strings.Split(route, ".")
		if len(parts) != 2 {
			panic("Invalid autogen action: " + route)
		}
		s.RegisterWriteHandler(route, (writeChecker(a.dp, parts[1], perm)))
	}
}

func writeChecker(dp dataprovider.DataProvider, collName, permission string) perm.WriteChecker {
	return perm.WriteCheckerFunc(func(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
		// Find meetingID
		var meetingID int
		if err := json.Unmarshal(payload["meeting_id"], &meetingID); err != nil {
			var id int
			if err := json.Unmarshal(payload["id"], &id); err != nil {
				return nil, fmt.Errorf("no valid meeting_id or id in payload")
			}

			fqid := collName + "/" + strconv.Itoa(id)
			meetingID, err = dp.MeetingFromModel(ctx, fqid)
			if err != nil {
				return nil, fmt.Errorf("getting meeting id for %s: %w", fqid, err)
			}
		}

		if err := perm.EnsurePerms(ctx, dp, userID, meetingID, permission); err != nil {
			return nil, fmt.Errorf("ensuring permission: %w", err)
		}

		return nil, nil
	})
}
