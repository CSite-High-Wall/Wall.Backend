package service

import (
	"wall-backend/internal/model"

	"github.com/spf13/viper"
)

type ConfigService struct {
	config *viper.Viper
}

func NewConfigService() ConfigService {
	return ConfigService{}
}

func (service *ConfigService) Initialize() error {
	config := viper.New()
	service.config = config

	config.AddConfigPath("conf")
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.WatchConfig()

	return config.ReadInConfig()
}

func (service *ConfigService) GetDataBaseConfig() model.DataBaseConfig {
	return model.DataBaseConfig{
		User:         service.config.GetString("mysql.user"),
		Password:     service.config.GetString("mysql.password"),
		Host:         service.config.GetString("mysql.host"),
		Port:         service.config.GetInt("mysql.port"),
		DataBaseName: service.config.GetString("mysql.database_name"),
	}
}

func (service *ConfigService) GetStaticFileSystemConfig() (string, string) {
	return service.config.GetString("server.staticFs_schema"), service.config.GetString("server.staticFs_host")
}
