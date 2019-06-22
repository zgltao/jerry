package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zgltao/jerry/libs/erro"
	"github.com/zgltao/jerry/libs/logging"
	lv "github.com/zgltao/jerry/libs/validator"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strings"
	"unicode"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	// 取最后一个Error为返回的Error
	length := len(c.Errors)
	if length > 0 {
		e := c.Errors[length-1]
		logging.Warn("全局日志:", e)
		switch e1 := e.Err.(type) {
		case *erro.HttpErr:
			writeHttpError(c, *e1)
		case erro.HttpErr:
			writeHttpError(c, e1)
		case validator.ValidationErrors:
			writeParamError(c, e1)
		case *validator.ValidationErrors:
			writeParamError(c, *e1)
		default:
			writeError(c, e.Err.Error())
		}
	}
}

func writeError(ctx *gin.Context, errString interface{}) {
	status := http.StatusBadRequest
	if ctx.Writer.Status() != http.StatusOK {
		status = ctx.Writer.Status()
	}
	s := erro.UnKnown.SetMsg(errString).SetUrl(ctx.Request.URL.String())
	ctx.JSON(status, s)
}

func writeParamError(ctx *gin.Context, e1 validator.ValidationErrors) {
	mapErrors := make(map[string]string)
	var (
		finalStr string
		s        *erro.HttpErr
	)
	for _, err := range e1 {
		param := err.Param()
		finalStr = err.Translate(lv.Trans)
		runes := []rune(err.StructField())
		runes[0] = unicode.ToLower(runes[0])
		fieldName := string(runes)
		finalStr = replaceParam(param, finalStr)
		mapErrors[fieldName] = finalStr
	}
	status := http.StatusBadRequest
	if ctx.Writer.Status() != http.StatusOK {
		status = ctx.Writer.Status()
	}
	if len(mapErrors) > 1 {
		s = erro.ParamsErr.SetMsg(mapErrors).SetUrl(ctx.Request.URL.String())
	} else {
		s = erro.ParamsErr.SetMsg(finalStr).SetUrl(ctx.Request.URL.String())
	}
	ctx.JSON(status, s)
}

func replaceParam(param, str string) string {
	switch param {
	case "Password":
		var replacer = "输入密码"
		res := strings.Replace(str, param, replacer, 1)
		return res
	default:
		return str
	}
}

func writeHttpError(ctx *gin.Context, httpErr erro.HttpErr) {
	s := httpErr.SetUrl(ctx.Request.URL.String())
	ctx.JSON(httpErr.HttpCode, s)
}
