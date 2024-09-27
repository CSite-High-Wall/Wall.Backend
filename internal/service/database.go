package service

import (
	"fmt"
	"wall-backend/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataBaseService struct {
	configService ConfigService
	DB            *gorm.DB
}

func NewDataBaseService(configService ConfigService) DataBaseService {
	return DataBaseService{
		configService: configService,
	}
}

func (service *DataBaseService) Connect() error {
	config := service.configService.GetDataBaseConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DataBaseName)

	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	service.DB = db
	return error
}

func (service *DataBaseService) InitializeDataTable() error {
	return service.DB.AutoMigrate(
		model.User{},
		model.Expression{},
	)
}
