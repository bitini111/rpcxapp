package conf

const (
	LOGIN_TYPE_START    int32 = 1 //合法登录类型的起点值
	LOGIN_TYPE_GUEST    int32 = 1
	LOGIN_TYPE_FACEBOOK int32 = 2 //FACEBOOK 登录
	LOGIN_TYPE_PHONE    int32 = 3 //手机登录
	LOGIN_TYPE_OPPO     int32 = 4
	LOGIN_TYPE_VIVO     int32 = 5
	LOGIN_TYPE_HUAWEI   int32 = 6
	LOGIN_TYPE_STEAM	int32 = 7
	LOGIN_TYPE_EMAIL    int32 = 8
	LOGIN_TYPE_END      int32 = 8 //合法登录类型的终点值（新增LOGIN_TYPE时，该值也需要更新）
	LOGIN_TYPE_HIDDEN   int32 = 100 //特殊的类型，真人改为机器人时用了，fb用户删除时也用了
)

//登录信息中的param参数的见容
type LoginParam struct {
	Token       string `json:"token"`
	AccessToken string `json:"access_token"`
}



//FB登录校验返回参数
//{
//"data": {
//"app_id": "669856953432696",
//"type": "USER",
//"application": "test",
//"data_access_expires_at": 1564305624,
//"expires_at": 1562912584,
//"is_valid": true,
//"issued_at": 1557728584,
//"metadata": {
//"auth_type": "rerequest",
//"sso": "ios"
//},
//"scopes": [
//"email",
//"public_profile"
//],
//"user_id": "115807719618574"
//}
//}
type FBCheckLoginRsp_Data_MetaData struct {
	AuthType string `json:"auth_type"`
	SSO      string `json:"sso"`
}
type FBCheckLoginRsp_Data struct {
	AppID               string                        `json:"app_id"`
	Type                string                        `json:"type"`
	Application         string                        `json:"application"`
	DataAccessExpiresAt int64                         `json:"data_access_expires_at"`
	ExpiresAt           int64                         `json:"expires_at"`
	IsValid             bool                          `json:"is_valid"`
	IssuedAt            int64                         `json:"issued_at"`
	MetaData            FBCheckLoginRsp_Data_MetaData `json:"metadata"`
	Scopes              []string                      `json:"scopes"`
	UserID              string                        `json:"user_id"`
}
type FBCheckLoginRsp struct {
	Data FBCheckLoginRsp_Data `json:"data"`
}

//const (
//	salt1 = "q2@1"
//	salt2 = "*6%x"
//)
//
////Token生成算法
//func GenToken(LoginType int32, account string) string {
//	t := strconv.FormatInt(time.Now().Unix(),10)
//	h := md5.New()
//	io.WriteString(h, strconv.Itoa(int(LoginType)))
//	io.WriteString(h, account)
//	io.WriteString(h, salt1)
//	io.WriteString(h, t)
//	io.WriteString(h, salt2)
//	return fmt.Sprintf("%x", h.Sum(nil))
//}

//redis中存放的token加密密钥及过期时间
type TokenKey struct {
	KeyA     string `redis:"key_a"`
	KeyB     string `redis:"key_b"`
	ExpireAt int64  `redis:"expire_at"`
}

type FBUserInfoPictureData struct {
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
	Url    string `json:"url"`
}

type FBPictureData struct {
	Data struct {
		Width        int32  `json:"width"`
		IsSilhouette bool   `json:"is_silhouette"`
		Height       int32  `json:"height"`
		Url          string `json:"url"`
	}
}
type FBUserInfoPicture struct {
	Data FBUserInfoPictureData `json:"data"`
}

type FBUserInfo struct {
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Picture FBUserInfoPicture `json:"picture"`
}

//OPPO登录返回
type OPPOUserInfo struct {
	Name string `json:"name"` //玩家昵称
}

type EmailVerificationCode struct {
	Email      string //	邮箱地址
	Code       int    //	验证码
	ValidTime  int64  //	有效期
	SendTime   int64  //	发送时间
	ExpireTime int64  //	有效时间(秒)
	IsCheck    bool   //	是否校验通过false未通过true通过
}


type PhoneVerificationCode struct {
	Code       int    //	验证码
	ValidTime  int64  //	有效期
}



type SendEmailMsg struct {
	Email    string `json:"email"`
	Language string `json:"language"`
}

//	邮件内容模板
type EmailHtml struct {
	Context string
}

//Steam返回用户信息
type SteamUserInfo struct {
	Name string `json:"name"` //昵称
	Avatar string `json:"avatar"` //头像
}