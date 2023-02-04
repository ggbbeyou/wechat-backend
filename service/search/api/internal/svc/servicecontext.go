package svc

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"time"
	"wechat-backend/service/search/api/internal/config"
	"wechat-backend/service/search/api/internal/middleware"
	"wechat-backend/service/search/api/internal/search"
)

type ServiceContext struct {
	Config                    config.Config
	VerifyAuthorityMiddleware rest.Middleware
	ElasticClient             *search.ElasticClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                    c,
		VerifyAuthorityMiddleware: middleware.NewVerifyAuthorityMiddleware(c).HandleVerifyAuth,

		ElasticClient: search.NewElasticClient(elasticsearch.Config{
			Addresses: c.Elastic.Addresses,
			Username:  c.Elastic.Username,
			Password:  c.Elastic.Password,
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   10,
				ResponseHeaderTimeout: time.Second,
				TLSClientConfig: &tls.Config{
					MinVersion:         tls.VersionTLS12,
					InsecureSkipVerify: true,
				},
			},
		}, c.Elastic.Index),
	}
}
