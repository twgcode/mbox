/**
@Author: wei-g
@Email: 17600113577@163.com
@Date: 2022/10/11 17:21
@Description: ioc 容器
*/

package app

import (
	"fmt"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ginApps = map[string]GinApp{}
)

// GinApp Http服务的实例
type GinApp interface {
	Config() error
	Registry(gin.IRouter)
	Name() string
	Version() string
}

func genKey(name, version string) string {
	return fmt.Sprintf("%s/%s", name, version)

}

// RegistryGinApp 服务实例注册
func RegistryGinApp(app GinApp) {
	// 已经注册的服务禁止再次注册
	key := genKey(app.Name(), app.Version())
	_, ok := ginApps[key]
	if ok {
		panic(fmt.Sprintf("http app %s has registed", key))
	}

	ginApps[key] = app
}

// LoadedGinApp 查询加载成功的服务
func LoadedGinApp() (apps []string) {
	for k := range ginApps {
		apps = append(apps, k)
	}
	return
}

func GetGinApp(name, version string) (GinApp, error) {
	key := genKey(name, version)
	app, ok := ginApps[key]
	if !ok {
		return nil, fmt.Errorf("http app %s not registed", key)
	}
	return app, nil
}

// LoadGinApp 装载所有的gin app
func LoadGinApp(pathPrefix string, root gin.IRouter) {
	for _, api := range ginApps {
		if pathPrefix != "" && !strings.HasPrefix(pathPrefix, "/") {
			pathPrefix = "/" + pathPrefix
		}
		api.Registry(root.Group(path.Join(pathPrefix, api.Version(), api.Name())))
	}
}
