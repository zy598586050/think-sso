package application

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
	"think-sso/internal/dao"
	"think-sso/internal/model"
	"think-sso/internal/service"
	"think-sso/utility"
)

type sApplication struct{}

func init() {
	service.RegisterApplication(New())
}

func New() *sApplication {
	return &sApplication{}
}

// HasApp 其他系统登录查询应用信息
func (s *sApplication) HasApp(ctx context.Context, AppId string, AppSecret string) (app *model.Application, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		count, err := dao.Application.Ctx(ctx).Where(dao.Application.Columns().AppId, AppId).Count()
		utility.ErrIsNil(ctx, err, "获取数据失败")
		if count <= 0 {
			utility.ErrIsNil(ctx, gerror.New("AppId不存在"))
		}
		count, err = dao.Application.Ctx(ctx).Where(dao.Application.Columns().AppId, AppId).Where(dao.Application.Columns().AppSecret, AppSecret).Count()
		if count <= 0 {
			utility.ErrIsNil(ctx, gerror.New("AppSecret不正确"))
		} else {
			dao.Application.Ctx(ctx).Where(dao.Application.Columns().AppId, AppId).Where(dao.Application.Columns().AppSecret, AppSecret).Scan(&app)
		}
	})
	return
}

// AppListByIds 通过ID列表查询应用
func (s *sApplication) AppListByIds(ctx context.Context, ids string) (res []*model.Application, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		appIds := strings.Split(ids, ",")
		if len(appIds) > 0 {
			err = dao.Application.Ctx(ctx).Where("id In(?)", appIds).Scan(&res)
		}
	})
	return
}
