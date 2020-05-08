package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hero1s/ginweb/i18n"
	"github.com/hero1s/ginweb/log"
	"github.com/hero1s/ginweb/pkg/validation"
	"net/http"
	"strconv"
)

type BaseUserInfo struct {
	Uid      uint64
	DeviceId string
}

// 返回给前端的数据格式
type Response struct {
	Code  error       `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Token string      `json:"token,omitempty"`
	Count int64       `json:"count,omitempty"`
}

type response struct {
	Code  string      `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Token string      `json:"token,omitempty"`
	Count int64       `json:"count,omitempty"`
}

// 自定义json marshal,因为error类型要转换为string返回给前端
func (r Response) MarshalJSON() ([]byte, error) {
	m := response{
		Code:  r.Code.Error(),
		Msg:   r.Msg,
		Data:  r.Data,
		Token: r.Token,
		Count: r.Count,
	}
	return json.Marshal(m)
}

type Gin struct {
	C *gin.Context
	BaseUserInfo
}

func (g *Gin) Response(httpCode int, err error, data interface{}) {
	g.C.JSON(httpCode, response{
		Code: err.Error(),
		Msg:  i18n.GetErrorMsg(err, 1),
		Data: data,
	})
	return
}

func (g *Gin) ResponseOk(err error, data interface{}) {
	g.C.JSON(http.StatusOK, response{
		Code: err.Error(),
		Msg:  i18n.GetErrorMsg(err, 1),
		Data: data,
	})
	return
}

func (g *Gin) ResponseSuccess(data interface{}) {
	g.C.JSON(http.StatusOK, response{
		Code: i18n.Success.Error(),
		Msg:  i18n.GetErrorMsg(i18n.Success, 1),
		Data: data,
	})
	return
}

// err must bee no-nil
func (g *Gin) ResponseError(err error) {
	errCode, errParse := strconv.ParseInt(err.Error(), 10, 64)
	if errParse == nil {
		if errCode >= i18n.ClientErrorBegin && errCode <= i18n.ClientErrorEnd {
			g.Response(http.StatusBadRequest, err, nil)
		} else if errCode >= i18n.SystemErrorBegin && errCode <= i18n.SystemErrorEnd {
			g.Response(http.StatusInternalServerError, err, nil)
		}
	}
	if errors.Is(err, i18n.DatabaseError) {
		g.ResponseOk(i18n.DatabaseError, nil)
		log.Error(err.Error())
		return
	} else if errors.Is(err, i18n.SystemError) {
		log.Error(err.Error())
		g.ResponseOk(i18n.SystemError, nil)
		return
	} else if _, ok := i18n.ErrorCode[err]; !ok {
		// err doesn't exist in ErrorCode, so print the unknown error
		g.Response(http.StatusInternalServerError, err, nil)
		log.Error("%v自定义错误:ip:%v:%v", g.C.Request.RequestURI, g.C.ClientIP(), err.Error())
		return
	} else {
		if err.Error() != i18n.TokenExpired.Error() && err.Error() != i18n.Unauthorized.Error() { //token过期不打印
			log.Error("%v未识别错误:ip:%v:%v", g.C.Request.RequestURI, g.C.ClientIP(), i18n.GetErrorMsg(err, 1))
		}
		g.ResponseOk(err, nil)
		return
	}
}

// 校验参数, p is a struct
func (g *Gin) ValidParams(p interface{}) bool {
	var Valid validation.Validation
	_, err := Valid.RecursiveValid(p)
	if err != nil {
		log.Error("%v 校验参数错误:%v", g.C.Request.RequestURI, err.Error())
		g.ResponseError(i18n.ParamsError)
		return false
	}
	if Valid.HasErrors() {
		for _, err := range Valid.Errors {
			log.Error("%v 参数不符合要求error:%v,key:%v,message:%v", g.C.Request.RequestURI, err.Error(), err.Key, err.Message)
			g.ResponseError(errors.New(fmt.Sprintf("%v", err.Message)))
			//g.ResponseError(i18n.ParamsNotFit)
			return false
		}
	}
	return true
}



