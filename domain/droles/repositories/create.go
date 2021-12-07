package repositories

import (
	"context"
	"ksuser/lib/app"
	"ksuser/lib/helper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *repo) Create(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return helper.ContextError(ctx)
	default:
	}

	query := `
        INSERT INTO roles (
					company_id, name, created_by, updated_by)
        VALUES ($1, $2, $3 ,$4)
    `

	_, err := u.db.ExecContext(ctx, query,
		ctx.Value(app.Ctx("companyID")).(string), u.pb.Name, ctx.Value(app.Ctx("userID")).(string), ctx.Value(app.Ctx("userID")).(string))

	if err != nil {
		u.log.Println(err.Error())
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
