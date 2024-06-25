/**
@Author: wei-g
@Email: 17600113577@163.com
@Date: 2022/10/11 17:21
@Description: ioc 容器
*/

package app

func InitAllApp() error {
	// 优先初始化内部app
	for _, api := range internalApps {
		if err := api.Config(); err != nil {
			return err
		}
	}

	for _, api := range grpcApps {
		if err := api.Config(); err != nil {
			return err
		}
	}

	for _, api := range ginApps {
		if err := api.Config(); err != nil {
			return err
		}
	}

	return nil
}
