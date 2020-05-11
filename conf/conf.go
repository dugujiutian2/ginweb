package conf

import (
	"flag"
	"fmt"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	env      string
	filePath string
	Conf     = &Config{}
)

type Config struct {
	ProjectName string
	Hystrix     *Hystrix
	HttpServer  *HttpServer
}

func init() {
	flag.StringVar(&env, "env", "", "env or prod")
	flag.StringVar(&filePath, "conf", "", "conf file path")
}

type Hystrix struct {
	Timeout                int // 超时时间
	MaxConcurrentRequests  int // 最大并发请求数
	RequestVolumeThreshold int // 开启熔断探测前的调用次数
	SleepWindow            int // 熔断发生后的等待恢复时间
	ErrorPercentThreshold  int // 错误百分比，请求数量大于等于RequestVolumeThreshold并且错误率到达这个百分比后就会启动熔断 默认值是50
}

type HttpServer struct {
	Port string
}

// 解析配置文件
func ParseConfig() error {
	flag.Parse()
	if filePath == "" {
		return errors.New("load conf path fail")
	}
	viper.SetConfigFile(filePath)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.UnmarshalKey(env, Conf)
	if err != nil {
		return err
	}
	fmt.Println(Conf.Hystrix)
	// 熔断配置
	hystrix.ConfigureCommand(Conf.ProjectName, hystrix.CommandConfig{
		Timeout:                Conf.Hystrix.Timeout,
		MaxConcurrentRequests:  Conf.Hystrix.MaxConcurrentRequests,
		RequestVolumeThreshold: Conf.Hystrix.RequestVolumeThreshold,
		SleepWindow:            Conf.Hystrix.SleepWindow,
		ErrorPercentThreshold:  Conf.Hystrix.ErrorPercentThreshold,
	})
	return nil
}
