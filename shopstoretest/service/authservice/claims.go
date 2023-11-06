package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"shopstoretest/entity"
)

type Claim struct {
	ID           uint        `json:"id"`
	Role         entity.Role `json:"role"`
	TokenSubject string      `json:"token_subject"`
	jwt.RegisteredClaims
}
