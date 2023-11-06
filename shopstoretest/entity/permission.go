package entity

type Permission struct {
	ID    uint            `json:"id"`
	Title PermissionTitle `json:"title"`
}

type PermissionTitle string

const (
	AddCategoryPermission = PermissionTitle("add_category")
	AddProductPermission  = PermissionTitle("add_product")
)
