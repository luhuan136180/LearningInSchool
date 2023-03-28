package message

//定义一个结构体

type User struct {
	//确定字段信息
	//为了序列化和反序列化成功，我们必须保证
	//用户信息的json字符串的key 和 结构体的字段对应的 tag 名字一致!!!
	UserId int`json:"user_id"`
	UserPwd string`json:"user_pwd"`
	UserName string`json:"user_name"`
	UserStatus int `json:"userStatus"` //用户状态..
}