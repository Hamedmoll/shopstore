package entity

type AccessControl struct {
	ID           uint      `json:"id"`
	ActorID      uint      `json:"actor_id"`
	ActorType    ActorType `json:"actor_type"`
	PermissionID uint      `json:"permission_id"`
}

type ActorType string

const (
	AdminActorType = "admin"
	UserActorType  = "user"
)
