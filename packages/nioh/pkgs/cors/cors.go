package cors

import (
	"time"

	"github.com/go-zaru/oss-hoshi/packages/nioh/pkgs/config"

	gors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/*
  Microsoft IE :: Will not be supported from 17th August 2021
  So you can (and will) use lowercase headers
*/

var gConf = gors.Config{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"*"},
	AllowHeaders:     []string{"access-control-allow-headers", "content-type", "content-length", "accept-encoding", "x-csrf-token", "authorization", "accept", "origin", "cache-control", "x-requested-with"},
	ExposeHeaders:    []string{"*"},
	AllowCredentials: true,
	MaxAge:           12 * time.Hour,
}

const key = "cors"

//  Essentially doing...
//  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//  c.Writer.Header().Set("Access-Control-Max-Age", "86400")
//  c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, UPDATE, OPTIONS")
//  c.Writer.Header().Set("access-control-allow-headers", "content-type", "content-length", "accept-encoding", "x-csrf-token", "authorization", "accept", "origin", "cache-control", "x-requested-with")
//  c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
//  c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

// EnableCors enables cors for server
func EnableCors(provider *config.CfgProvider) gin.HandlerFunc {
	provider.UnmarshalKey(key, &gConf)
	instance := gors.New(gConf)
	return instance
}
