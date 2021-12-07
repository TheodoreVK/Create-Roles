package validation

import (
	"ksuser/domain/droles"
	"log"
)

type validation struct {
	log      *log.Logger
	roleRepo droles.RoleRepoInterface
}

func NewRoleValidation(log *log.Logger, roleRepo droles.RoleRepoInterface) droles.RoleValidationInterface {
	return &validation{
		log:      log,
		roleRepo: roleRepo,
	}
}
