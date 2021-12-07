package usecase

import (
	"ksuser/domain/droles"
	"log"
)

type roleservice struct {
	log      *log.Logger
	roleRepo droles.RoleRepoInterface
}

func NewRoleService(log *log.Logger, roleRepo droles.RoleRepoInterface) droles.RoleUsecaseInterface {
	return &roleservice{
		log:      log,
		roleRepo: roleRepo,
	}
}
