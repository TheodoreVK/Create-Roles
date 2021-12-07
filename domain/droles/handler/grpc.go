package handler

import (
	"context"
	"ksuser/domain/droles"
	"ksuser/pb/roles"
)

type RoleHandler struct {
	usecase droles.RoleUsecaseInterface
}

func NewRoleHandler(usecase droles.RoleUsecaseInterface) *RoleHandler {
	handler := new(RoleHandler)
	handler.usecase = usecase
	return handler
}

func (u *RoleHandler) Create(ctx context.Context, in *roles.Role) (*roles.Role, error) {
	return u.usecase.Create(ctx, in)
}
