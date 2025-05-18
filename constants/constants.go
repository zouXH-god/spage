package constants

const (
	RoleAdmin       = "admin"        // 管理员
	RoleUser        = "user"         // 普通用户
	FlagSystemAdmin = "system_admin" // 系统管理员标志

	CaptchaTypeDisable   = "disable" // 禁用验证码
	CaptchaTypeTurnstile = "turnstile"
	CaptchaTypeReCaptcha = "recaptcha"
	CaptchaTypeHCaptcha  = "hcaptcha"
	ModeDev              = "dev"
	ModeProd             = "prod"
)
