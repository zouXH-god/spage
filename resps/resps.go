package resps

import "github.com/cloudwego/hertz/pkg/app"

// Custom 自定义响应
func Custom(c *app.RequestContext, code int, message string, data ...map[string]any) {
	if len(data) == 0 {
		data = append(data, map[string]any{})
	}
	data[0]["message"] = message
	c.JSON(code, data[0])
}

// 2xx

func Ok(c *app.RequestContext, message string, data ...map[string]any) {
	if len(data) == 0 {
		data = append(data, map[string]any{})
	}
	data[0]["message"] = message
	c.JSON(200, data[0])
}

// 4xx

func BadRequest(c *app.RequestContext, message string) {
	c.JSON(400, map[string]string{"message": message})
}

func Unauthorized(c *app.RequestContext, message string) {
	c.JSON(401, map[string]string{"message": message})
}

func Forbidden(c *app.RequestContext, message string) {
	c.JSON(403, map[string]string{"message": message})
}

func NotFound(c *app.RequestContext, message string) {
	c.JSON(404, map[string]string{"message": message})
}

// 5xx

func InternalServerError(c *app.RequestContext, message string) {
	c.JSON(500, map[string]string{"message": message})
}

func ServiceUnavailable(c *app.RequestContext, message string) {
	c.JSON(503, map[string]string{"message": message})
}

func RespMessageWithError(message string, err error) string {
	return message + ": " + err.Error()
}
