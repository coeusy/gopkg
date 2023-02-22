package persistence

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/coeusy/gopkg/cfg"
)

const urlFmt = "%s:%s@tcp(%s:%d)/%s?charset=%s"

func NewConnection(cfg cfg.RDSConf) *gorm.DB {
	connCfg := cfg.Connection
	url := fmt.Sprintf(urlFmt, connCfg.User, connCfg.Password, connCfg.Host, connCfg.Port, cfg.Database, connCfg.Charset)
	dbConn, err := gorm.Open(cfg.Dialect, url)
	if err != nil {
		panic(err)
	}
	if cfg.Debug {
		return dbConn.Debug()
	}
	return dbConn
}
