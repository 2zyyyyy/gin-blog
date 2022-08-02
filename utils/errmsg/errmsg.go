package errmsg

// 请求结果
const (
	SUCCESS = 200
	ERROR   = 500
)

// 用户相关错误
const (
	ErrorUsernameUsed = 1001 + iota
	ErrorPasswordWrong
	ErrorUserNotExist
	ErrorTokenExist
	ErrorTokenRuntime
	ErrorTokenWrong
	ErrorTokenTypeWrong
	ErrorUserNoRight
)

// ErrorArtNotExist 文章相关错误
const ErrorArtNotExist = 2001

// 分类相关错误
const (
	ErrorCateNameUsed = 3001 + iota
	ErrorCateNotExist
)

var codeMsg = map[int]string{
	SUCCESS:             "OK",
	ERROR:               "FAIL",
	ErrorUsernameUsed:   "用户名已存在！",
	ErrorPasswordWrong:  "密码错误",
	ErrorUserNotExist:   "用户不存在",
	ErrorTokenExist:     "TOKEN不存在,请重新登录",
	ErrorTokenRuntime:   "TOKEN已过期,请重新登录",
	ErrorTokenWrong:     "TOKEN不正确,请重新登录",
	ErrorTokenTypeWrong: "TOKEN格式错误,请重新登录",
	ErrorUserNoRight:    "该用户无权限",

	ErrorArtNotExist: "文章不存在",

	ErrorCateNameUsed: "该分类已存在",
	ErrorCateNotExist: "该分类不存在",
}

// GetMsg 根据code返回对应的信息
func GetMsg(code int) string {
	return codeMsg[code]
}
