package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"github.com/zhuxbin01/pkg/mysql/gormotel"
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	dbIns *gorm.DB
)

func GetDBIns(opts ...Option) (*gorm.DB, error) {
	cfg := _defaultOptions
	for _,o := range opts {
		o.apply(&cfg)
	}

	var err error

	if err != nil {
		return nil, err
	}

	once.Do(func() {
		dbIns, err = New(cfg)
	})

	return dbIns, nil
}

func New(opts Options) (*gorm.DB, error) {
	dns := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Database,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		return nil, err
	}

	if opts.enableTrace {
		// todo
		db.Use(gormotel.NewPlugin(opts.gormOtelOptions...))
	}

	sqlDB,err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// maxLifeTime必须要比mysql服务器设置的wait_timeout小，否则会导致golang侧连接池依然保留已被mysql服务器关闭了的连接
	// 可通过show variables like 'wait_timeout'查看
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	sqlDB.SetConnMaxIdleTime(opts.MaxConnMaxIdleTime)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}

func ReplaceDefualtOptions(opts Options)  {
	 _defaultOptions = opts
}