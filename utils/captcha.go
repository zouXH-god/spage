package utils

import (
	"fmt"
	"github.com/LiteyukiStudio/spage/constants"
)

type captchaType struct{}

var Captcha = captchaType{}

type CaptchaConfig struct {
	Type        string
	SiteSecrete string
	SecretKey   string
}

func (captchaType) VerifyCaptcha(captchaConfig *CaptchaConfig, captchaToken string) (bool, error) {
	switch captchaConfig.Type {
	case constants.CaptchaTypeDisable:
		return true, nil
	case constants.CaptchaTypeTurnstile:
		return true, nil
	default:
		return false, nil
	}
}

// GenerateCaptchaWidget 生成Captcha组件HTML
func (captchaType) GenerateCaptchaWidget(captchaConfig *CaptchaConfig) (string, error) {
	switch captchaConfig.Type {
	case constants.CaptchaTypeDisable:
		return "", nil
	case constants.CaptchaTypeTurnstile:
		return fmt.Sprintf(`
<div class="checkbox mb-3">
      <!-- The following line controls and configures the Turnstile widget. -->
      <div class="cf-turnstile" data-sitekey="%s" data-theme="light"></div>
      <!-- end. -->
</div>
<script src="https://challenges.cloudflare.com/turnstile/v0/api.js" async defer></script>
`, captchaConfig.SiteSecrete), nil
	case constants.CaptchaTypeReCaptcha:
		return fmt.Sprintf(`
<div class="g-recaptcha" data-sitekey="%s"></div>
<script src="https://www.google.com/recaptcha/api.js" async defer></script>
`, captchaConfig.SiteSecrete), nil
	case constants.CaptchaTypeHCaptcha:
		return fmt.Sprintf(`
<div class="h-captcha" data-sitekey="%s"></div>
<script src="https://js.hcaptcha.com/1/api.js" async defer></script>
`, captchaConfig.SiteSecrete), nil
	default:
		return "", nil
	}
}
