package auth

import (
	"context"

	"github.com/atom-apps/auth/modules/auth/service"
	v1 "github.com/atom-apps/auth/proto/v1"
	"github.com/atom-providers/log"
	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/rogeecn/atom/contracts"
	"go-micro.dev/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ v1.UserServiceHandler = (*UserService)(nil)

// @provider
type UserService struct {
	*Common
	Name string `inject:"false"`

	micro          contracts.MicroService
	door           *casdoorsdk.Client
	userMappingSvc *service.UserMappingService
	tenantSvc      *service.TenantService
	tenantUserSvc  *service.TenantUserService
}

func (svc *UserService) Prepare() error {
	log.Info("register: UserService")
	return v1.RegisterUserServiceHandler(svc.micro.GetEngine().(micro.Service).Server(), svc)
}

// GetByName implements v1.UserServiceHandler.
func (svc *UserService) GetMapping(ctx context.Context, in *v1.GetMappingRequest, out *v1.GetMappingResponse) error {
	claim, err := svc.Claim(ctx, svc.door)
	if err != nil {
		return status.Error(codes.PermissionDenied, "invalid token: "+err.Error())
	}

	if !claim.IsAdmin {
		in.Name = claim.Name
	}

	if in.Name == "" {
		return status.Error(codes.PermissionDenied, "empty name is not allowed")
	}

	userModel, err := svc.userMappingSvc.GetByName(ctx, in.Name)
	if err != nil {
		return err
	}

	tenantUser, err := svc.tenantUserSvc.GetByUserID(ctx, userModel.ID)
	if err != nil {
		return err
	}

	tenant, err := svc.tenantSvc.GetByID(ctx, tenantUser.TenantID)
	if err != nil {
		return err
	}

	out.Id = userModel.ID
	out.Name = userModel.Name
	out.TenantId = tenant.ID
	out.TenantName = tenant.Name
	out.IsTenantAdmin = tenantUser.IsAdmin

	return nil
}
