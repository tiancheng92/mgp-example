package user

import (
	"fmt"
	"mgp_example/config"
	"mgp_example/pkg/ecode"

	ump_sdk "github.com/Yostardev/ump-sdk"
	"github.com/tiancheng92/mgp"
	"github.com/tiancheng92/mgp/errors"
)

func GetInfo(ctx *mgp.Context, applicationID int) (*ump_sdk.UserInfo, error) {
	if userInfo, ok := ctx.Get(fmt.Sprintf("user_info_%d", applicationID)); ok {
		if d, ok1 := userInfo.(*ump_sdk.UserInfo); ok1 {
			return d, nil
		}
	}

	token := ctx.GetHeader("Authorization")

	if token == "" {
		token = ctx.Value("token").(string)
	}

	user, err := ump_sdk.NewClient(config.GetConf().Ump.Url, applicationID, token).GetUserInfo()
	if err != nil {
		return nil, err
	}

	ctx.Set(fmt.Sprintf("user_info_%d", applicationID), user)
	if applicationID != 0 {
		ctx.Set("user_info_0", user)
	}
	return user, errors.WithCode(ecode.ErrServerGet, err)
}
