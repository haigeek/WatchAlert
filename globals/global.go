package globals

import (
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"go.uber.org/zap"
	"prometheus-manager/models"
	"prometheus-manager/pkg/cache"
)

var (
	Config    models.App
	Logger    *zap.Logger
	FeiShuCli *lark.Client
	CacheCli  *cache.InMemoryCache
)

var (
	Layout     = "2006-01-02T15:04:05.000Z"
	AlertType  string
	DataSource string
	RespBody   []byte
)
