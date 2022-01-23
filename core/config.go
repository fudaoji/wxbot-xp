package core

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

// RedisConfig Redis配置
var (
	RedisConfig redisConfig
	MySQLConfig mysqlConfig
)

// Redis配置
type redisConfig struct {
	Host     string // Redis主机
	Port     string // Redis端口
	Password string // Redis密码
	Db       int    // Redis库
}

// MySQL配置
type mysqlConfig struct {
	Host     string // 主机
	Port     string // 端口
	Username string // 用户名
	Password string // 密码
	DbName   string // 数据库名称
}

//InitConfig 读取配置文件
func InitConfig() {
	mode := flag.String("mode", "dev", "dev mode")
	flag.Parse()
	viper.SetConfigFile(fmt.Sprintf("./setting_%s.yaml", *mode))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

//GetVal 获取配置文件的字符串配置值
func GetVal(key string, defaultVal string) string {
	if viper.IsSet(key) {
		return viper.GetString(key)
	}
	return defaultVal
}

//GetIntVal 获取配置文件的整型配置值
func GetIntVal(key string, defaultVal int) int {
	if viper.IsSet(key) {
		return viper.GetInt(key)
	}
	return defaultVal
}
