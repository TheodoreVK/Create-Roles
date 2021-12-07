package usecase

import (
	"context"
	"ksuser/domain/droles/validation"
	"ksuser/lib/helper"
	"ksuser/pb/roles"
)

func (u *roleservice) Create(ctx context.Context, in *roles.Role) (*roles.Role, error) {
	select {
	case <-ctx.Done():
		return nil, helper.ContextError(ctx)
	default:
	}

	ctx, err := helper.GetMetadata(ctx)
	if err != nil {
		return nil, err
	}

	dValidation := validation.NewRoleValidation(u.log, u.roleRepo)
	err = dValidation.Create(ctx, in)
	if err != nil {
		return nil, err
	}

	u.roleRepo.SetPb(in)

	err = u.roleRepo.Create(ctx)
	if err != nil {
		return nil, err
	}

	return u.roleRepo.GetPb(), nil
}
