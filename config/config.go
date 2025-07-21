package config
package config

import "github.com/spf13/viper"

type Config struct {
    Server struct {
        Port int `mapstructure:"port"`
    } `mapstructure:"server"`
    Database struct {
        Host     string `mapstructure:"host"`
        Port     int    `mapstructure:"port"`
        User     string `mapstructure:"user"`
        Password string `mapstructure:"password"`
        Dbname   string `mapstructure:"dbname"`
        Sslmode  string `mapstructure:"sslmode"`
    } `mapstructure:"database"`
    Wechat struct {
        Appid string `mapstructure:"appid"`
        Mchid string `mapstructure:"mchid"`
    } `mapstructure:"wechat"`
}

var Cfg *Config

func LoadConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./config") // 指定配置文件路径

    if err := viper.ReadInConfig(); err != nil {
        panic("读取配置文件失败: " + err.Error())
    }

    if err := viper.Unmarshal(&Cfg); err != nil {
        panic("解析配置文件失败: " + err.Error())
    }
}