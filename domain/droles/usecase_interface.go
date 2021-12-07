package droles

import (
	"context"
	"ksuser/pb/roles"
)

type RoleUsecaseInterface interface {
	Create(ctx context.Context, in *roles.Role) (*roles.Role, error)
}
