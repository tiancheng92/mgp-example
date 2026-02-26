package service

import (
	"mgp_example/config"
	"mgp_example/pkg/ecode"
	"mgp_example/pkg/user"

	ump_sdk "github.com/Yostardev/ump-sdk"
	"github.com/tiancheng92/mgp"
	"github.com/tiancheng92/mgp/errors"
)

type userService struct{}

func NewUserService() UserServiceInterface {
	return new(userService)
}

func (u *userService) GetUserInfo(ctx *mgp.Context) (*ump_sdk.UserInfo, error) {
	return user.GetInfo(ctx, 0)
}

func (u *userService) GetAuthList(ctx *mgp.Context, authType string) ([]string, error) {
	umpConfig := config.GetConf().Ump
	var applicationID int
	switch authType {
	case "menu":
		applicationID = umpConfig.MenuAppID
	case "functional":
		applicationID = umpConfig.FunctionalAppID
	default:
		return nil, errors.WithCode(ecode.ErrServerGet, "auth_type not found")
	}

	info, err := user.GetInfo(ctx, applicationID)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0, len(info.Authorities))
	for i := range info.Authorities {
		res = append(res, info.Authorities[i].Obj)
	}

	return res, nil
}
