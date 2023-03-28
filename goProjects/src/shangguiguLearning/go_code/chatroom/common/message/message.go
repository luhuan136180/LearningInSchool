package message

//用常量，确定一些信息类型
const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

//这里我们定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的类型
}

//定义具体的消息
type LoginMes struct {
	UserId   int    `json:"user_id"`   //用户id
	UserPwd  string `json:"user_pwd"`  //用户密码
	UserName string `json:"user_name"` //用户名字
}

//登录返回的消息结构体
type LoginResMes struct {
	Code    int    `json:"code"` // 返回状态码 500 表示该用户未注册 200表示登录成功
	UsersId []int  //增加字段，保存用户id的切片
	Error   string `json:"error"` // 返回错误信息
}
type RegisterMes struct {
	User User `json:"user"`
}
type RegisterResMes struct {
	Code  int    `json:"code"`  // 返回状态码 400 表示该用户已经占用 200表示注册成功
	Error string `json:"error"` // 返回错误信息
}

type NotifyUserStatusMes struct {
	UserId int `json:"user_id"` //用户的id
	Status int `json:"status"`  //用户的状态
}

//增加一个SmsMes //发送的消息
type SmsMes struct {
	Content string `json:"content"` //内容
	User           //匿名结构体
}
