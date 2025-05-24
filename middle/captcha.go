package middle

import (
	"context"
	"github.com/LiteyukiStudio/spage/config"
	"github.com/LiteyukiStudio/spage/resps"
	"github.com/LiteyukiStudio/spage/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type captchaType struct{}

var Captcha = captchaType{}

type CaptchaReq struct {
	CaptchaToken string `json:"captcha_token"`
}

func (captchaType) UseCaptcha() app.HandlerFunc {
	captchaConfig := &utils.CaptchaConfig{
		Type:        config.CaptchaType,
		SiteSecrete: config.CaptchaSiteKey,
		SecretKey:   config.CaptchaSecretKey,
	}
	restyClient := resty.New()
	return func(ctx context.Context, c *app.RequestContext) {
		var req CaptchaReq
		if err := c.BindAndValidate(&req); err != nil {
			resps.BadRequest(c, resps.ParameterError)
			c.Abort()
			return
		}
		ok, err := utils.Captcha.VerifyCaptcha(restyClient, captchaConfig, req.CaptchaToken)
		if err != nil {
			logrus.Error("Captcha verification error:", err)
			resps.InternalServerError(c, "Captcha verification failed")
			c.Abort()
			return
		}
		if !ok {
			logrus.Warn("Captcha verification failed for token:", req.CaptchaToken)
			resps.Forbidden(c, "Captcha verification failed")
			c.Abort()
			return
		}
		c.Next(ctx) // Proceed to the next handler if captcha verification is successful
		return
	}
}
