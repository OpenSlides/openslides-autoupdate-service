package assignment_candidate

import "github.com/OpenSlides/openslides-permission-service/internal/allowed"

// TODO: through model...
// TODO: assignments.can_nominate_self and assignments.can_nominate_other
var Create = allowed.BuildCreate([]string{
	"assignment_id",
	"user_id",
}, "assignments.can_manage")

// TODO: through model...
// needs assignments.can_manage from the meeting of the assignment
var Sort = allowed.BuildModify([]string{"assignment_id",
	"candidate_ids"}, "assignment", "assignments.can_manage")

// TODO: assignments.can_nominate_self and assignments.can_nominate_other
var Delete = allowed.BuildModify([]string{"id"}, "assignment", "assignments.can_manage")
