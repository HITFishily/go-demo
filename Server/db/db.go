package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"
	"github.com/pkg/errors"
)

const DefaultDriver = "mysql"

var defaultDbEngine *xorm.Engine   // 读写
var defaultDBEngineRO *xorm.Engine // 只读

// 初始化读写库
func InitDefaultDbEngine(dataSource string, driverName ...string) error {
	if defaultDbEngine == nil {
		engine, err := InitDBEngine(dataSource, driverName...)
		if err != nil {
			return err
		}
		defaultDbEngine = engine
	}
	return nil
}

// 初始化只读库
func InitDefaultDbEngineRO(dataSource string, driverName ...string) error {
	if defaultDBEngineRO == nil {
		engine, err := InitDBEngine(dataSource, driverName...)
		if err != nil {
			return err
		}
		defaultDBEngineRO = engine
	}
	return nil
}

func InitDBEngine(dataSource string, driverName ...string) (engine *xorm.Engine, err error) {
	driver := DefaultDriver
	if len(driverName) > 0 {
		if tmp := driverName[0]; tmp != "" {
			driver = tmp
		}
	}

	if x, e := xorm.NewEngine(driver, dataSource); e != nil {
		err = errors.Wrap(e, "xorm.NewEngine")
	} else {
		engine = x
	}
	return
}

// 获取读写库的句柄
func GetDefaultDbEngine() *xorm.Engine {
	if defaultDbEngine == nil {
		panic(fmt.Errorf("default defaultDbEngine is nil, call InitDefaultDbEngine first"))
	}
	return defaultDbEngine
}

// 获取只读库句柄
func GetDefaultDbEngineRO() *xorm.Engine {
	if defaultDBEngineRO == nil {
		panic(fmt.Errorf("default defaultDbEngineRO is nil, call InitDefaultDbEngineRO first"))
	}
	return defaultDBEngineRO
}
