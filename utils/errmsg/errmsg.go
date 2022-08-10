package errmsg

type ResCode int64

// 请求结果
const (
	SUCCESS ResCode = 200
	ERROR   ResCode = 500
)

// 用户相关错误
const (
	ErrorUsernameUsed ResCode = 1001 + iota
	ErrorPasswordWrong
	ErrorUserNotExist
	ErrorTokenExist
	ErrorTokenRuntime
	ErrorTokenWrong
	ErrorTokenTypeWrong
	ErrorUserNoRight
)

// ErrorArticleNotExist 文章相关错误
const ErrorArticleNotExist ResCode = 2001

// 分类相关错误
const (
	ErrorCategoryNameUsed ResCode = 3001 + iota
	ErrorCategoryNotExist
)

var codeMsg = map[ResCode]string{
	SUCCESS:             "OK",
	ERROR:               "服务器繁忙",
	ErrorUsernameUsed:   "用户名已存在",
	ErrorPasswordWrong:  "密码错误",
	ErrorUserNotExist:   "用户不存在",
	ErrorTokenExist:     "TOKEN不存在,请重新登录",
	ErrorTokenRuntime:   "TOKEN已过期,请重新登录",
	ErrorTokenWrong:     "TOKEN不正确,请重新登录",
	ErrorTokenTypeWrong: "TOKEN格式错误,请重新登录",
	ErrorUserNoRight:    "该用户无权限",

	ErrorArticleNotExist: "文章不存在",

	ErrorCategoryNameUsed: "该分类已存在",
	ErrorCategoryNotExist: "该分类不存在",
}

// GetMsg 根据code返回对应的信息
func (code ResCode) GetMsg() string {
	msg, ok := codeMsg[code]
	if !ok {
		msg = codeMsg[ERROR]
	}
	return msg
}
