package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Con = new(Config)
)

type Config struct {
	Ser *Server `mapstructure:"Server"`
	Mys *Mysql  `mapstructure:"Mysql"`
}

type Server struct {
	HttpPort     int    `mapstructure:"HttpPort"`
	RunMode      string `mapstructure:"RunMode"`
	ReadTimeOut  int    `mapstructure:"ReadTimeOut"`
	WriteTimeOut int    `mapstructure:"WriteTimeOut"`
}

type Mysql struct {
	Name         string `mapstructure:"Name"`
	Host         string `mapstructure:"Host"`
	Port         string `mapstructure:"Port"`
	PassWord     string `mapstructure:"PassWord"`
	DBName       string `mapstructure:"DBName"`
	Charset      string `mapstructure:"Charset"`
	MaxIdleConns int    `mapstructure:"MaxIdleConns"`
	MaxOpenConns int    `mapstructure:"MaxOpenConns"`
}

type Setting struct {
	vp *viper.Viper
}

func InitSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")


	//vp.SetConfigFile("config.yaml")
	vp.AddConfigPath("./config/")
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := vp.Unmarshal(Con); err != nil {

	}
	vp.WatchConfig()
	vp.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置已经被修改")
	})


	return &Setting{
		vp: vp,
	}, nil

}
