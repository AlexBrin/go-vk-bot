package config

import (
	"strings"

	"github.com/AlexBrin/go-vkbot/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	Path = "ConfigPath"

	LogLevel = "Log.Level"

	BotGoroutine = "Bot.Goroutine"

	VkToken         = "Vk.Token"
	VkGroupId       = "Vk.GroupId"
	VkCommandPrefix = "Vk.Command.Prefix"
)

type Configuration struct {
	Conf *viper.Viper
}

func NewConf(configPath string) (*Configuration, error) {
	viperConfig := viper.New()

	viperConfig.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	_ = viper.BindPFlag(Path, pflag.Lookup("config-path"))

	if configPath == "" {
		configPath = viper.GetString(Path)
	}

	if configPath != "" {
		logger.Message{Message: "Чтение конфигурации"}.Info()
		viperConfig.SetConfigFile(configPath)

		err := viperConfig.ReadInConfig()
		if err != nil {
			logger.Message{Message: "Ошибка чтения конфигурации", Err: err}.AddField("config-path", configPath).Error()
			return nil, err
		}
	}

	setDefaultInt(viperConfig, LogLevel, 4)
	logger.SetLevel(logrus.Level(viperConfig.GetInt(LogLevel)))

	setDefaultStringSlice(viperConfig, VkCommandPrefix, []string{"!", "/", "."})

	return &Configuration{
		Conf: viperConfig,
	}, nil
}

func setDefaultStringSlice(viper *viper.Viper, param string, value []string) {
	if len(viper.GetStringSlice(param)) == 0 {
		viper.SetDefault(param, value)
	}
}

func setDefaultString(viper *viper.Viper, param string, value string) {
	if viper.GetString(param) == "" {
		viper.SetDefault(param, value)
	}
}

func setDefaultInt(viper *viper.Viper, param string, value int) {
	if viper.GetInt(param) == 0 {
		viper.SetDefault(param, value)
	}
}

func setDefaultBool(viper *viper.Viper, param string, value bool) {
	if !viper.GetBool(param) {
		viper.SetDefault(param, value)
	}
}
