package repositories

import (
	"database/sql"
	"ksuser/domain/droles"
	"ksuser/pb/roles"
	"log"
)

type repo struct {
	db  *sql.DB
	log *log.Logger
	pb  roles.Role
}

func NewRoleRepo(db *sql.DB, log *log.Logger) droles.RoleRepoInterface {
	return &repo{
		db:  db,
		log: log,
	}
}

func (u *repo) GetPb() *roles.Role {
	return &u.pb
}

func (u *repo) SetPb(in *roles.Role) {
	if len(in.Id) > 0 {
		u.pb.Id = in.Id
	}
	if len(in.Name) > 0 {
		u.pb.Name = in.Name
	}
	if len(in.UpdatedBy) > 0 {
		u.pb.UpdatedBy = in.UpdatedBy
	}
	if len(in.CreatedAt) > 0 {
		u.pb.CreatedAt = in.CreatedAt
	}
	if len(in.UpdatedAt) > 0 {
		u.pb.UpdatedAt = in.UpdatedAt
	}
}
