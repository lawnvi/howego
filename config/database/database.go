package database

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

var DB *gorm.DB

func Init() error {
	//读取当前环境数据库配置
	db, err := loadInfos()
	if err != nil {
		db = DBInfo{
			Name:     "howego_debug",
			DbType:   "mysql",
			Host:     "150.158.161.214",
			Port:     "3306",
			User:     "root",
			Password: "3.14",
			Config: DBConfig{
				MaxIdle: 2,
				MaxOpen: 4,
			},
		}
	}
	//获取数据库连接
	DB, err = db.connect()
	return err
}

//数据库连接信息
type DBInfo struct {
	Name     string `yaml:"name"`
	DbType   string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Config   DBConfig
}

//连接池配置
type DBConfig struct {
	MaxIdle int `yaml:"max_idle"`
	MaxOpen int `yaml:"max_open"`
}

//对应database.yml
type DataBase struct {
	Debug   DBInfo
	Release DBInfo
}

//读取当前运行环境的数据库配置
func loadInfos() (info DBInfo, err error) {
	fileData, err := ioutil.ReadFile("./config/database/database.yml")
	if err != nil {
		return
	}
	var infos = DataBase{}
	err = yaml.Unmarshal(fileData, &infos)
	fmt.Println(infos)
	mode := gin.Mode()
	if mode == gin.DebugMode || mode == gin.TestMode {
		info = infos.Debug
	} else {
		info = infos.Release
	}
	return
}

//连接数据库
func (db *DBInfo) connect() (*gorm.DB, error) {
	gormDB, err := gorm.Open(db.DbType, db.User+":"+db.Password+"@tcp("+db.Host+")/"+db.Name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//全局禁用表复数
	//DB.SingularTable(true)
	gormDB.DB().SetMaxIdleConns(2)    //最大空闲连接数
	gormDB.DB().SetMaxOpenConns(4)    //最大连接数
	gormDB.DB().SetConnMaxLifetime(time.Second + 300) //设置连接空闲超时
	return gormDB, nil
}

func HasTable(table interface{}) bool {
	has := DB.HasTable(table)
	if !has {
		DB.CreateTable(table)
	}
	return has
}
