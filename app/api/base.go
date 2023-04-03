package api

import (
	"github.com/gin-gonic/gin"
	"gva-lbx/app/model/request"
	"gva-lbx/app/service"
	"gva-lbx/global"
	"gva-lbx/response"
	"net/http"
)

var Base = new(base)

type base struct{}

// Generate
// @Tags    SystemCaptcha
// @Summary 获取验证码
// @accept  application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.Captcha} "验证码获取成功!"
// @Failure 400 {object} response.Response{data=any}              "Bad Request"
// @Router  /base/captcha [post]
func (a *base) Generate(c *gin.Context) response.Response {
	captcha, err := service.Captcha.Generate(c.Request.Context(), c.ClientIP())
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Message: "验证码获取失败!", Error: err}
	}
	return response.Response{Code: http.StatusOK, Data: captcha, Message: "验证码获取成功!"}
}

// Login
// @Tags    SystemCaptcha
// @Summary 登录
// @accept  application/json
// @Produce application/json
// @Param   data body     request.UserLogin                          true "请求参数"
// @Success 200  {object} response.Response{data=response.UserLogin} "登录成功!"
// @Failure 400  {object} response.Response{data=any}                "Bad Request"
// @Failure 500  {object} response.Response{data=any}                "登录失败!"
// @Router  /base/login [post]
func (a *base) Login(c *gin.Context) response.Response {
	var info request.UserLogin
	err := c.ShouldBindJSON(&info)
	if err != nil {
		return response.Response{Code: http.StatusBadRequest, Error: err}
	}
	ip := c.ClientIP()
	if !service.Captcha.ExplosionProof(c.Request.Context(), ip) || !global.CaptchaStore.Verify(info.CaptchaId, info.Captcha, true) {
		return response.Response{Code: http.StatusBadRequest, Message: "验证码错误!"}
	}
	data, err := service.User.Login(c.Request.Context(), info)
	if err != nil {
		service.Captcha.Explosion(c.Request.Context(), ip)
		return response.Response{Code: http.StatusInternalServerError, Error: err, Message: "用户名不存在或者密码错误!"}
	}
	if !data.Enable {
		return response.Response{Code: http.StatusOK, Message: "用户禁止登录!"}
	}
	service.Captcha.Explosion(c.Request.Context(), ip)
	return response.Response{Code: http.StatusOK, Data: data, Message: "登录成功!"}
}
