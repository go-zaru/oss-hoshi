package metaheaders

import (
	"github.com/gin-gonic/gin"
	"github.com/go-zaru/oss-hoshi/packages/nioh/pkgs/config"
	"github.com/google/uuid"
)

// MetaHeadersConfig ...
type metaheadersConfig struct {
	Revision            string `mapstructure:""`
	RevisionHeaderName  string `mapstructure:""`
	EnableRequestID     bool   `mapstructure:""`
	RequestIDHeaderName string `mapstructure:""`
}

var mhcf = metaheadersConfig{
	RevisionHeaderName:  "x-revision",
	EnableRequestID:     true,
	RequestIDHeaderName: "x-request-id",
}

const key = "metaheaders"

// EnableMetaHeaders ...
func EnableMetaHeaders(provider *config.CfgProvider) gin.HandlerFunc {
	provider.UnmarshalKey(key, &mhcf)

	return func(c *gin.Context) {
		if mhcf.Revision != "" {
			c.Writer.Header().Set(mhcf.RevisionHeaderName, mhcf.Revision)
		}

		if mhcf.EnableRequestID {
			c.Writer.Header().Set(mhcf.RequestIDHeaderName, uuid.New().String())
		}
	}
}
