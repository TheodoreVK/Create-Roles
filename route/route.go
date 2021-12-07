package route

import (
	"database/sql"
	"ksuser/pb/roles"
	"ksuser/pb/users"
	"log"

	userHandler "ksuser/domain/dusers/handler"
	userRepo "ksuser/domain/dusers/repositories"
	userUsecase "ksuser/domain/dusers/usecase"

	roleHandler "ksuser/domain/droles/handler"
	roleRepo "ksuser/domain/droles/repositories"
	roleUsecase "ksuser/domain/droles/usecase"

	"google.golang.org/grpc"
)

func GrpcRoute(grpcServer *grpc.Server, log *log.Logger, db *sql.DB) {
	userServer := userHandler.NewUserHandler(
		userUsecase.NewUserService(log, userRepo.NewUserRepo(db, log)),
	)

	users.RegisterUsersServiceServer(grpcServer, userServer)

	roleServer := roleHandler.NewRoleHandler(
		roleUsecase.NewRoleService(log, roleRepo.NewRoleRepo(db, log)),
	)

	roles.RegisterRolesServiceServer(grpcServer, roleServer)
}
