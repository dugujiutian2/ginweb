package conf

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/pkg/errors"
)

var (
	env      string
	filePath string
	// Conf config
	Conf *Config
)

type Config struct {
	ProjectName string
	Hystrix     *Hystrix
	DB          *DB
	HttpServer  *HttpServer
}

// Default new a config with specified defualt value.
func Default() *Config {
	return &Config{}
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

type DB struct {
	UserName string
	Passwd   string
	Host     string
	Schema   string
	Idle     int
	ShowSQL  bool
	LogFile  string
	LogLevel int
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
	Conf = Default()
	_, err := toml.DecodeFile(filePath, &Conf)
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
