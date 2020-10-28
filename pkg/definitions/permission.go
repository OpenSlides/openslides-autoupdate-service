package definitions

// Docs: See https://github.com/FinnStutzenstein/OpenSlides/blob/permissionService/docs/interfaces/permission-service.txt
type Permission interface {
	IsAllowed(name string, userId int, data FqfieldData) (bool, map[string]interface{}, error)
	RestrictFqIds(fqids map[Fqid]bool, userId int) (map[Fqid]bool, error)
	RestrictFqfields(fqfields map[Fqfield]bool, userId int) (map[Fqfield]bool, error)
	AdditionalUpdate(updated FqfieldData) ([]Id, error)
}
