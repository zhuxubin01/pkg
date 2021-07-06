// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"github.com/zhuxbin01/pkg/mysql/gormotel"
	"time"

	//"github.com/marmotedu/iam/pkg/db"
)

// Options defines optsions for mysql database.
var _defaultOptions = Options{
	Host:               "127.0.0.1:3306",
	Username:           "root",
	Password:           "root",
	Database:           "test",
	MaxIdleConnections: 100,
	// 链接池中链接最大持续空闲时间,主要看服务的qps情况,一般不做设置,链接的关闭可以通过MaxConnectionLifeTime实现
	MaxConnMaxIdleTime:    0 * time.Second,
	MaxConnectionLifeTime: 1200 * time.Second,
	LogLevel:              0,
}

type Options struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	MaxConnMaxIdleTime    time.Duration
	LogLevel              int
	enableTrace           bool
	gormOtelOptions       []gormotel.Option
}

type Option interface {
	apply(c *Options)
}

type optionFunc func(server *Options)

func (f optionFunc) apply(o *Options) {
	f(o)
}

func WithHost(host string) Option {
	return optionFunc(
		func(o *Options) {
			o.Host = host
		})
}

func WithPassword(password string) Option {
	return optionFunc(
		func(o *Options) {
			o.Password = password
		})
}

func WithUsername(userName string) Option {
	return optionFunc(
		func(o *Options) {
			o.Username = userName
		})
}

func WithDatabase(db string) Option {
	return optionFunc(
		func(o *Options) {
			o.Database = db
		})
}

func WithEnableTrace(enableTrace bool) Option {
	return optionFunc(
		func(o *Options) {
			o.enableTrace = enableTrace
		})
}
