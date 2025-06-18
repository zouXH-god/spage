package constants

const (
	GlobalRoleAdmin   = "admin"        // 管理员 Admin
	GlobalRoleUser    = "user"         // 普通用户 User
	OrgRoleOwner      = "owner"        // 组织所有者 Organization Owner
	OrgRoleMember     = "member"       // 组织成员 Organization Member
	ProjectRoleOwner  = "owner"        // 项目所有者 Project Owner
	ProjectRoleMember = "member"       // 项目成员 Project Member
	FlagSystemAdmin   = "system_admin" // 系统管理员标志 SystemAdmin

	// DB

	PreloadFieldMembers = "Members"
	PreloadFieldOwners  = "Owners"
	PreloadFieldProject = "Project"
	PreloadFieldFile    = "File"

	CaptchaTypeDisable   = "disable"     // 禁用验证码 Captcha
	CaptchaTypeTurnstile = "turnstile"   // 云flare turnstile
	CaptchaTypeReCaptcha = "recaptcha"   // Google reCAPTCHA
	CaptchaTypeHCaptcha  = "hcaptcha"    // HCaptcha
	CaptchaDevPasscode   = "dev-captcha" // 开发者验证码 Developer Captcha

	ModeDev  = "dev"  // 开发者模式 Developer Mode
	ModeProd = "prod" // 生产模式 Production Mode

	OwnerTypeUser = "user"         // 个人用户 Personal user
	OwnerTypeOrg  = "organization" // 组织用户 Organization user

	FileDriverLocal      = "local"
	FileDriverWebdav     = "webdav"
	FileDriverS3         = "s3"
	WebDavPolicyProxy    = "proxy"
	WebDavPolicyRedirect = "redirect"
)
