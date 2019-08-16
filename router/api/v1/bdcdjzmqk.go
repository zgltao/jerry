package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func InitLocal() *xorm.Engine {
	return openDB(viper.GetString("oracle_config"))
}
func openDB(oracle_config string) *xorm.Engine {
	engine, err := xorm.NewEngine("oci8", oracle_config)
	if err != nil {
		log.Fatalf("%s, Database connection failed. Database name: %s \n", err, oracle_config)
	}
	err = engine.Ping()
	if err != nil {
		log.Fatalf("%s, Database is killed. Database name: %s \n", err, oracle_config)
	}
	setupDB(engine)
	return engine
}
func setupDB(db *xorm.Engine) {
	db.SetLogLevel(core.LOG_DEBUG)
}

func GetBdcdjzmqk(c *gin.Context) {
	var err error
	engine := InitLocal()
	if err != nil {
		c.Error(err)
	}
	tabs, err := engine.DBMetas()
	if err != nil {
		c.Error(err)
	}
	println(len(tabs))

	sql := "select * from DEPT"
	results, err := engine.Query(sql)
	for _, one := range results {
		jsonStr, err := json.Marshal(one)
		println(jsonStr)
		if err != nil {
			c.Error(err)
		}
	}
	c.JSON(http.StatusOK, results)
	engine.Clone()
}
