package distributed

import (
	"KWeb/framework"
	"KWeb/framework/contract"
	"errors"
	"github.com/gofrs/flock"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// LocalDistributedService 代表框架的App实现
type LocalDistributedService struct {
	container framework.Container // 服务容器
}

// NewLocalDistributedService 初始化本地分布式服务
func NewLocalDistributedService(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	return &LocalDistributedService{container: container}, nil
}

// Select 为分布式选择器
func (s LocalDistributedService) Select(serviceName string, appID string, holdTime time.Duration) (selectAppID string, err error) {
	appService := s.container.MustMake(contract.AppKey).(contract.App)
	runtimeFolder := appService.RuntimeFolder()
	lockFile := filepath.Join(runtimeFolder, "distribute_"+serviceName)

	// 创建文件锁
	fileLock := flock.New(lockFile)

	// 尝试获取文件锁，设置非阻塞模式
	locked, err := fileLock.TryLock()
	if err != nil {
		return "", err
	}

	// 如果未能获取锁
	if !locked {
		// 读取被选择的appid
		selectAppIDByt, err := os.ReadFile(lockFile)
		if err != nil {
			return "", err
		}
		return string(selectAppIDByt), nil
	}

	// 在一段时间内，选举有效，其他节点在这段时间不能再进行抢占
	go func() {
		defer func() {
			// 释放文件锁
			fileLock.Unlock()
			// 删除文件锁对应的文件
			os.Remove(lockFile)
		}()
		// 创建选举结果有效的计时器
		timer := time.NewTimer(holdTime)
		// 等待计时器结束
		<-timer.C
	}()

	// 这里已经是抢占到了，将抢占到的appID写入文件
	if err := ioutil.WriteFile(lockFile, []byte(appID), 0666); err != nil {
		return "", err
	}
	return appID, nil
}
