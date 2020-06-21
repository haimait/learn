package initialize

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"haimait/learn/gin01/global"
	"haimait/learn/gin01/models"
	"os"
)

type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
	TablePrefix  string `mapstructure:"table_prefix" json:"table_prefix" yaml:"table_prefix"`
}

func Newmysql() *Mysql {
	return &Mysql{
		Username:     "root",
		Password:     "123456",
		Path:         "127.0.0.1:3306",
		Dbname:       "gin01",
		Config:       "charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdleConns: 10,
		MaxOpenConns: 10,
		LogMode:      true,
		TablePrefix:  "hm_",
	}
}

//连接数据库
func InitMySQL() {
	var mysqlconf = Newmysql()
	db, err := gorm.Open("mysql", mysqlconf.Username+":"+mysqlconf.Password+"@("+mysqlconf.Path+")/"+mysqlconf.Dbname+"?"+mysqlconf.Config)
	if err != nil {
		//global.GVA_LOG.Error("MySQL启动异常", err)
		fmt.Println("MySQL启动异常", err)
		os.Exit(0)
		global.DB.Close()
		return
	}
	global.DB = db
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return mysqlconf.TablePrefix + defaultTableName
	}

	global.DB.DB().SetMaxIdleConns(mysqlconf.MaxIdleConns)
	global.DB.DB().SetMaxOpenConns(mysqlconf.MaxOpenConns)
	global.DB.LogMode(mysqlconf.LogMode)

	// 全局禁用表名复数
	global.DB.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	DBTables()
}

// 初使化创建数据表
func DBTables() {
	db := global.DB
	db.AutoMigrate(
		models.Todo{},
	)
	fmt.Println("register table success")
}
