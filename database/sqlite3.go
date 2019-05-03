package database

import (
	"fmt"
	"os"
	"time"

	"github.com/rxc-wujianhua/go-utils/logger"

	"github.com/go-log/log"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	// lite 驱动
	_ "github.com/mattn/go-sqlite3"
)

var (
	// engine sqlite3引擎
	engine *xorm.Engine
	// sqlite3Server
	sqlite3SQLServer string

	sqlite3MAXIdleConns          = 1000
	sqlite3MaxOpenConns          = 10000
	sqlite3ConnMaxLifeTime int64 = 120
	reportDelaySec               = 30
)

// con map结构
type con struct {
	k string
	v interface{}
}

// OpenSqlite3Conn 创建一个新的xorm引擎，用于连接sqlite3
func OpenSqlite3Conn() *xorm.Engine {
	defer logger.ExceptionDump()

	s := time.Now()
	// 开启一个新的引擎
	engine, err := xorm.NewEngine("sqlite3", "./blog.db")
	if err != nil {
		panic(fmt.Sprintf("failed to connect sqlite3:%v", err))
	}

	// 设置基础引擎参数
	engine.DB().SetMaxIdleConns(sqlite3MAXIdleConns)
	engine.DB().SetMaxOpenConns(sqlite3MaxOpenConns)
	engine.DB().SetConnMaxLifetime(time.Duration(120 * time.Second))

	// 名称映射规则
	tbMapper := core.NewPrefixMapper(core.GonicMapper{}, "bl_")
	engine.SetMapper(tbMapper)

	// 开启调试模式
	if os.Getenv("DEBUG_SQL") == "true" {
		engine.ShowSQL(true)
		engine.Logger().SetLevel(core.LOG_DEBUG)
		log.Logf("Entering DEBUG_SQL %v", os.Getenv("DEBUG_SQL"))
	}

	// 返回引擎
	log.Logf("connected to sqlite3 took: %v to %v", time.Since(s), sqlite3SQLServer)
	return engine
}

// StartXormSqlite3 启动xorm的sqlite3引擎，保持连接
func StartXormSqlite3(initTable func(sqlEngine *xorm.Engine)) {
	defer logger.ExceptionDump()

	sqlEngine := OpenSqlite3Conn()

	initTable(sqlEngine)
}

// countOpenConnections 返回当前引擎的状态
func countOpenConnections() int {
	return engine.DB().Stats().OpenConnections
}

// CopySqlite3Conn 复制引擎,如果当前引擎未被销毁，则直接使用，反之则重新创建
func CopySqlite3Conn() *xorm.Engine {
	if err := engine.DB().Ping(); err != nil {
		engine = OpenSqlite3Conn()
		return engine
	}

	return engine
}
