package app

import (
	"KWeb/framework"
	"KWeb/framework/util"
	"errors"
	"flag"
	"github.com/google/uuid"
	"path/filepath"
)

// KApp 代表hade框架的App实现
type KApp struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
	appId      string              // 表示当前这个app的唯一id, 可以用于分布式锁等
}

// Version 实现版本
func (h KApp) Version() string {
	return "0.0.3"
}

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (h KApp) BaseFolder() string {
	if h.baseFolder != "" {
		return h.baseFolder
	}

	// 如果参数也没有，使用默认的当前路径
	return util.GetExecDirectory()
}

// ConfigFolder  表示配置文件地址
func (h KApp) ConfigFolder() string {
	return filepath.Join(h.BaseFolder(), "config")
}

// LogFolder 表示日志存放地址
func (h KApp) LogFolder() string {
	return filepath.Join(h.StorageFolder(), "log")
}

func (h KApp) HttpFolder() string {
	return filepath.Join(h.BaseFolder(), "http")
}

func (h KApp) ConsoleFolder() string {
	return filepath.Join(h.BaseFolder(), "console")
}

func (h KApp) StorageFolder() string {
	return filepath.Join(h.BaseFolder(), "storage")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (h KApp) ProviderFolder() string {
	return filepath.Join(h.BaseFolder(), "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (h KApp) MiddlewareFolder() string {
	return filepath.Join(h.HttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (h KApp) CommandFolder() string {
	return filepath.Join(h.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (h KApp) RuntimeFolder() string {
	return filepath.Join(h.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (h KApp) TestFolder() string {
	return filepath.Join(h.BaseFolder(), "test")
}

// NewKApp 初始化KApp
func NewKApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数，一个是容器，一个是baseFolder
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	// 如果没有设置，则使用参数
	if baseFolder == "" {
		flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数, 默认为当前路径")
		flag.Parse()
	}
	appId := uuid.New().String()
	return &KApp{baseFolder: baseFolder, container: container, appId: appId}, nil
}

// AppID 表示这个App的唯一ID
func (h KApp) AppID() string {
	return h.appId
}
