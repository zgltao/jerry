package v1

import (
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

func B2S(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

func SqlSelectFormat(results []map[string][]uint8) []map[string]string {
	var resultsstr []map[string]string
	//fmt.Printf("resultsstr type:%T\n", resultsstr)
	for _, one := range results {
		//fmt.Printf("one type:%T\n", one)
		onestring := make(map[string]string)
		for keeee, oneee := range one {
			//fmt.Printf("oneee type:%T\n", oneee)
			var oneeestr = B2S(oneee)
			//fmt.Printf("oneee type:%T\n", oneeestr)
			onestring[keeee] = oneeestr
		}
		resultsstr = append(resultsstr, onestring)
		println(resultsstr)
	}
	return resultsstr
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
	if err != nil {
		c.Error(err)
	}
	resultsstr := SqlSelectFormat(results)
	c.JSON(http.StatusOK, resultsstr)
	engine.Clone()
}
