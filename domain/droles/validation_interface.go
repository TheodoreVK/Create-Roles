package droles

import (
	"context"
	"ksuser/pb/roles"
)

type RoleValidationInterface interface {
	Create(ctx context.Context, id *roles.Role) error
}
