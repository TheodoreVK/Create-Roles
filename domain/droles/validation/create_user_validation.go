package validation

import (
	"context"
	"ksuser/lib/helper"
	"ksuser/pb/roles"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *validation) Create(ctx context.Context, in *roles.Role) error {
	select {
	case <-ctx.Done():
		return helper.ContextError(ctx)
	default:
	}

	if len(in.Name) == 0 {
		u.log.Println("please supply valid name")
		return status.Error(codes.InvalidArgument, "please supply valid name")
	}

	return nil
}
