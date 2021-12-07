package droles

import (
	"context"
	"ksuser/pb/roles"
)

type RoleRepoInterface interface {
	Create(ctx context.Context) error
	GetPb() *roles.Role
	SetPb(*roles.Role)
}
