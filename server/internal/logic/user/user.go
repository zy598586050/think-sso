package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "think-sso/api/v1"
	"think-sso/internal/consts"
	"think-sso/internal/dao"
	"think-sso/internal/model"
	"think-sso/internal/model/entity"
	"think-sso/internal/service"
	"think-sso/utility"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) GetUserByEmailPassword(ctx context.Context, req *v1.EmailLoginReq) (res *entity.User, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		userCount, _ := dao.User.Ctx(ctx).Where(dao.User.Columns().Email, req.Email).Count()
		if userCount <= 0 {
			utility.ErrIsNil(ctx, gerror.New("用户不存在"))
		}
		err = dao.User.Ctx(ctx).Where(dao.User.Columns().Email, req.Email).Where(dao.User.Columns().Password, req.Password).Scan(&res)
	})
	return
}

func (s *sUser) GetUserList(ctx context.Context, req *v1.UserListReq) (total int, userList []*model.User, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.User.Ctx(ctx)
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		total, err = m.Count()
		utility.ErrIsNil(ctx, err, "获取用户数据失败")
		err = m.Page(req.PageNum, req.PageSize).Scan(&userList)
	})
	return
}
