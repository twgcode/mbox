/**
@Author: wei-g
@Email: 17600113577@163.com
@Date: 2022/10/11 17:21
@Description: ioc 容器
*/

package app

import (
	"fmt"

	"google.golang.org/grpc"
)

var (
	grpcApps = map[string]GRPCApp{}
)

// GRPCApp GRPC服务的实例
type GRPCApp interface {
	Config() error
	Registry(*grpc.Server)
	Name() string
}

// RegistryGrpcApp 服务实例注册
func RegistryGrpcApp(app GRPCApp) {
	// 已经注册的服务禁止再次注册
	_, ok := grpcApps[app.Name()]
	if ok {
		panic(fmt.Sprintf("grpc app %s has registed", app.Name()))
	}
	grpcApps[app.Name()] = app
}

// LoadedGrpcApp 查询加载成功的服务
func LoadedGrpcApp() (apps []string) {
	for k := range grpcApps {
		apps = append(apps, k)
	}
	return
}

func GetGrpcApp(name string) (GRPCApp, error) {
	app, ok := grpcApps[name]
	if !ok {
		return nil, fmt.Errorf("grpc app %s not registed", name)
	}
	return app, nil
}

// LoadGrpcApp 加载所有的Grpc app
func LoadGrpcApp(server *grpc.Server) {
	for _, app := range grpcApps {
		app.Registry(server)
	}
}
