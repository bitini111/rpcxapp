package conf

import "time"

var (
	ERR_PARAMS          = "请求参数错误！"
	ERR_NO_ROW_EFFECTED = "没有数据行被影响！"
)

//公共返回码
const (
	ECODE_RSP_FINISHED        int32 = 20000 //特殊状态码，返回了此状态吗表示某rpc服务已经直接回包了，不用重复回包；
	SUCCESS                   int32 = 0     //表示操作成功
	ECODE_OPERATE_DEFAULT_ERR int32 = -1    //默认的操作失败的错误码
	ECODE_OPERATE_FAIL        int32 = -1000
	ECODE_PARAM_ILLEGAL       int32 = -1001 //参数数值不合法

	ECODE_MYSQL_GET         int32 = -1002
	ECODE_MYSQL_SET         int32 = -1003
	ECODE_CONFIG_GET        int32 = -1004
	ECODE_CONFIG_WRONG      int32 = -1005
	ECODE_CONCURRENCY_WRONG int32 = -1006
	ECODE_NEED_LOGIN        int32 = -1007
	ECODE_REDIS_GET         int32 = -1008
	ECODE_REDIS_SET         int32 = -1009
)

//================game模块错误 1000+ ================
const (
	//登录子游戏
	ECODE_ILLEGAL_REQ              int32 = 1001 //表示不合理的请求，比如：提供给match用的gamesvr收到了普通玩家的入桌请求。
	ECODE_ILLEGAL_USER             int32 = 1002 //不合法的玩家信息，可能uid不合法；
	ECODE_ILLEGAL_ACSVID           int32 = 1003 //来自不合理的AC
	ECODE_ILLEGAL_TID              int32 = 1004 //传入的桌子id与实际桌子id不一致
	ECODE_CANT_FIND_TID            int32 = 1005 //找不到桌子对象
	ECODE_RECONNECT_CANT_FIND_USER int32 = 1006 //找不到桌子对象
	ECODE_GET_MONEY_ERR            int32 = 1007 //获取用户金币失败
	ECODE_NO_EMPTY_SEAT            int32 = 1008 //获取不到空闲的座位
	ECODE_HAS_SEATED               int32 = 1009 //已经坐下的玩家重复坐下
	ECODE_NON_LOOKON_SEAT          int32 = 1010 //非站着的人坐下
	ECODE_ILLEGAL_SEAT_NUM         int32 = 1011 //非法的座位号
	ECODE_CANT_FIND_USER           int32 = 1012 //围观阵容中找不到该玩家
	ECODE_CHARGE_CHIP_ERR          int32 = 1013 //兑换筹码失败
	ECODE_ILLEGAL_SYS_STATUS       int32 = 1014 //登录子游戏时，服务器状态异常
	ECODE_ILLEGAL_LEVEL            int32 = 1015 //Alloc找不到对应的场次信息
	ECODE_ILLEGAL_TABLE            int32 = 1016 //Alloc获取不到桌子对象，请检查桌子id
	ECODE_EMPTY_HEAP               int32 = 1017 //Alloc快速配桌时，桌子堆为空
	ECODE_NO_TABLE_FOUND           int32 = 1018 //Alloc快速配桌时，找不到合适的桌子，可能是金币数额不合适
	ECODE_ALLOC_PARAM_ILLEGAL      int32 = 1019 //指定桌子的登录的参数不对
	ECODE_IN_OTHER_TABLE           int32 = 1020 //玩家已经在其他桌子上了
	ECODE_NOT_ON_TABLE             int32 = 1021 //玩家不在桌子上，但是收到了玩家对桌子的操作。
	ECODE_FUTURE_FOLD_FORBID_IN    int32 = 1022 //提前离开的人不允许再进来，也不允许进别的桌子
	ECODE_RETRY_LATER              int32 = 1023 //玩家超时后，系统帮助其自动操作，操作结果广播前玩家消息又到了，就告知其稍后重试
	ECODE_RETIRE_ILLEGAL           int32 = 1024 //退休的请求数据异常
	ECODE_OUT_OF_SERVICE           int32 = 1025 //当前服务还未正常启动
	ECODE_GAME_RETIRE_LOGIN        int32 = 1026 //游戏退休时,不能登录
	ECODE_GAME_NOT_ENOUGH_MONEY    int32 = 1027 //玩家的金币筹码不足，不能快速开始。
	ECODE_DENIED_BY_ANTI_CHEATING  int32 = 1028 //进入房间失败，被防作弊机制限制

	ECODE_MATCH_START_TABLE_BUSY int32 = 1040 //match通知开启一局游戏时，桌子处于忙碌状态
	ECODE_ILLEGAL_PARAM          int32 = 1041 //非法的参数
	ECODE_ALLOC_ILLEGAL_STATUS   int32 = 1042 //ALLOC服务处于非法的状态下
	ECODE_ALLOC_MATCHING         int32 = 1043 //玩家处于比赛状态下，不允许通过alloc进入普通game中
	ECODE_MATCHID_DIFFERENT      int32 = 1044 //该桌子使用的matchId与本次请求使用的matchId不一致
	ECODE_ALLOC_CONTEXT_INVALID  int32 = 1045 //配桌服务取上下文中的信息失败
	ECODE_ALLOC_SOURCE_INVALID   int32 = 1047 //配桌服务请求来源无效
	ECODE_ALLOC_TABLE_NOT_EXISTS int32 = 1046 //桌子不存在

	//logout
	ECODE_USER_NOT_FOUND       int32 = 1050 //退出时用户未找到
	ECODE_NOT_IN_LOOKON_SEAT   int32 = 1051 //围观阵容和坐下阵容中找不到该玩家
	ECODE_ALL_IN_CAN_NOT_LEAVE int32 = 1052 //all-in用户不能主动离开

	//chipin/check/call
	ECODE_NON_SEAT_USER    int32 = 1060 //非坐下的玩家来执行下注操作
	ECODE_TABLE_NOT_BUSY   int32 = 1061 //玩家操作时，桌子已经处于空闲状态
	ECODE_ILLEGAL_OP_USER  int32 = 1062 //非法的操作人
	ECODE_ILLEGAL_OP_MONEY int32 = 1063 //操作的筹码数额不对，可能超过了持有筹码数
	ECODE_ILLEGAL_CHECK    int32 = 1064 //已下注数小于其他人，不能执行check操作
	ECODE_CALL_ZERO        int32 = 1065 //跟注的数额不能小于等于0
	ECODE_CALL_LESS        int32 = 1066 //玩家下注的数额小于最小下注数
	ECODE_ADD_LESS         int32 = 1067 //玩家下注的数额大于最小下注，小于最小加注，不合理
	ECODE_CALL_ILLEGAL     int32 = 1068 //金币下的比别人少，也没有allin
	ECODE_ILLEGAL_OPTYPE   int32 = 1069 //非法的操作类型
	ECODE_CHIP_IN_REPEATED int32 = 1070 //重复操作

	//stand sit
	ECODE_HAS_STAND         int32 = 1080 //站起时，已经站起或者离开的玩家
	ECODE_NOT_YOUR_TURN     int32 = 1081 //站起时，未轮到你操作
	ECODE_NO_SUITABLE_SEAT  int32 = 1082 //坐下时，系统找不到合适的座位
	ECODE_ILLEGAL_SEATUSER  int32 = 1083 //向seatUser中加入的数据有异常
	ECODE_SEATUSER_INIT_ERR int32 = 1084 //seatUser初始化异常
	ECODE_USER_ON_SEAT      int32 = 1085 //该座位上有人
	ECODE_USER_EXISTS       int32 = 1086 //玩家已经在另外的座位上了

	//设置自动买入
	ECODE_CHIP_MORE_THAN_MONEY int32 = 1090 //兑换筹码数量超过了金币数量
	ECODE_CHIP_MORE_THAN_SET   int32 = 1091 //兑换筹码数量超过了当前持有筹码数量
	ECODE_ILLEGAL_SET          int32 = 1092 //大于了设置的最大值或者小于了设置的最小值

	//设置玩家掉线 SetOffLine
	ECODE_NO_USER_FOUND int32 = 1100 //座位上和围观阵容中都找不到玩家

	//====================比赛模块
	ECODE_MATCH_NO_CONFIG_FOUND        int32 = 1200 //找不到比赛配置
	ECODE_MATCH_NO_SIGNED_FOUND        int32 = 1201 //找不到报名的玩家
	ECODE_MATCH_NO_TABLE_FOUND         int32 = 1202 //找不到合适的桌子
	ECODE_MATCH_NO_MATCH_FOUND         int32 = 1203 //延迟进入时，找不到比赛对象
	ECODE_MATCH_SIGN_ILLEGAL_TIME      int32 = 1204 //报名：未到报名时间或者已经超过了报名时间
	ECODE_MATCH_SIGN_FULL              int32 = 1205 //报名：人数已满
	ECODE_MATCH_SIGN_ILLEGAL_COIN      int32 = 1206 //报名：金币报名费不够
	ECODE_MATCH_SIGN_FAILED_COIN       int32 = 1207 //报名：扣除报名费失败
	ECODE_MATCH_SIGN_RESIGN            int32 = 1208 //报名：重复报名
	ECODE_MATCH_SIGN_FAILED            int32 = 1209 //报名：报名失败，可能是存redis的时候失败。
	ECODE_MATCH_UNSIGN_FAILED          int32 = 1210 //取消报名：取消报名失败。
	ECODE_MATCH_UNSIGN_FAILED_COIN     int32 = 1211 //取消报名：返还报名费失败
	ECODE_MATCH_UNSIGNED_UNSIGN        int32 = 1212 //没有报名的玩家来取消报名
	ECODE_MATCH_NO_TABLES_FOUND        int32 = 1213 //开赛时，找不到合适的桌子
	ECODE_MATCH_NO_ENOUGH_USERS        int32 = 1214 //开赛时，真实加入到比赛的玩家数量不够
	ECODE_MATCH_NO_ENOUGH_SIGNED       int32 = 1215 //开赛时检测到没有足够的报名人数。
	ECODE_TABLE_START_NOTIFY_ERR       int32 = 1216 //通知某桌子开始比赛的结果异常。
	ECODE_MATCH_DISP_USER_ERR          int32 = 1217 //把玩家分配到具体的某个桌子上出现异常。
	ECODE_MATCH_READY_NOT_SIGNED       int32 = 1218 //未报名的玩家不能进入备赛区
	ECODE_MATCH_READY_SIGN_FAILED      int32 = 1219 //未报名的玩家进入备赛区的时候失败
	ECODE_MATCH_WITHDRAW_FAILED        int32 = 1220 //解散比赛时出现异常
	ECODE_MATCH_RETURN_TABLE_FAILED    int32 = 1221 //解散比赛时,归还桌子失败
	ECODE_MATCH_TABLE_ERR              int32 = 1222 //比赛的玩家进桌子失败
	ECODE_MATCH_ILLEGAL_PARAM          int32 = 1223 //非法的比赛参数值
	ECODE_MATCH_RESIGN                 int32 = 1224 //玩家对某场比赛重复签到
	ECODE_MATCH_ILLEGAL_STATUS         int32 = 1225 //玩家签到时在状态不对.(比赛中,桌子上,B区域)
	ECODE_MATCH_ILLEGAL_MatchSign      int32 = 1226 //比赛中的玩家不允许报名
	ECODE_MATCH_MatchSign_WARNING      int32 = 1227 //警告玩家:很快会有一场MTT比赛,是否还要继续报名本比赛
	ECODE_MATCH_NOT_IN_THE_MATCH       int32 = 1228 //玩家不在此桌子上，不能进。
	ECODE_MATCH_SIGN_FAILED_UPSTREAM   int32 = 1229 //报名失败：上游服务返回失败
	ECODE_MATCH_UNSIGN_FAILED_UPSTREAM int32 = 1230 //取消报名失败：上游服务返回失败
	ECODE_MATCH_GEN_MATCH_ID_FAILED          = 1231 //生成比赛ID失败
	ECODE_MATCH_UNSIGN_FAILED_STARTED        = 1232 //取消报名失败：比赛已经开始

	//比赛分桌子
	ECODE_MATCHALLOC_NOT_ENOUGH    int32 = 1230 //申请的桌子数量不够
	ECODE_MATCHALLOC_ILLEGAL_PARAM int32 = 1231 //matchAlloc请求参数异常
	ECODE_MATCHALLOC_LESS_TABLES   int32 = 1232 //matchAlloc分配的桌子数小于请求的桌子数
	ECODE_MATCHALLOC_NOT_BORROWED  int32 = 1233 //归还桌子时，找不到该matchid的借桌子记录

	//登录大厅
	ECODE_HALL_LOGIN_ERR int32 = 1300
)

//各服务逻辑错误，5000+
const (
	//破产服务相关
	ECODE_BANKRUPT_MORE_MONEY      int32 = 5001 //未达到破产线
	ECODE_BANKRUPT_TIME_ILLEGAL    int32 = 5002 //破产时间未设置
	ECODE_BANKRUPT_CAN_NOT_SUBSIDY int32 = 5003 //当前无破产补助领取资格
	ECODE_BANKRUPT_DELIVERY_FAIL   int32 = 5004 //调用发货失败

	//商城服务相关
	ECODE_MALL_GOODS_ILLEGAL                  int32 = 5050 //商品不存在
	ECODE_MALL_PRICE_NOT_MATCH                int32 = 5051 //商品价格不匹配
	ECODE_MALL_PAY_MODE_ERROR                 int32 = 5052 //商品支付方式不匹配
	ECODE_MALL_DIAMOND_SAFE                   int32 = 5053 //钻石不能被加赠或是兑换
	ECODE_MALL_NOT_ENOUGH_MONEY               int32 = 5054 //钱不够
	ECODE_MALL_SUB_MONEY_ERR                  int32 = 5055 //扣费失败
	ECODE_MALL_DELIVERY_ERR                   int32 = 5056 //发货失败
	ECODE_MALL_GET_MONEY_ERR                  int32 = 5057 //获取金币信息失败
	ECODE_MALL_ITEM_NOT_EXISTS                int32 = 5058 //商品不存在
	ECODE_MALL_ORDER_FAILED                   int32 = 5059 //下单失败
	ECODE_MALL_ORDER_NOT_EXISTS               int32 = 5060 //订单不存在
	ECODE_MALL_ORDER_INVALID                  int32 = 5061 //订单状态异常
	ECODE_MALL_ORDER_DONE                     int32 = 5065 //订单已完成
	ECODE_MALL_ORDER_DELIVERING               int32 = 5063 //订单正在发货
	ECODE_MALL_ORDER_NOT_YOURS                int32 = 5064 //订单不是当前用户的
	ECODE_MALL_ORDER_UPDATE_DELIVERING_FAILED int32 = 5062 //订单更新状态失败
	ECODE_MALL_ORDER_VERIFY_FAILED            int32 = 5066 //订单验证失败
	ECODE_MALL_ORDER_UPDATE_DELIVERED_FAILED  int32 = 5067 //订单更新状态失败
	ECODE_MALL_ORDER_DELIVER_FAILED           int32 = 5068 //订单发货失败
	ECODE_MALL_HAS_NO_FIRST_BUY               int32 = 5069 //用户无首充优惠资格
	ECODE_MALL_VALID_FAILED                   int32 = 5070 //验证失败
	ECODE_MALL_HAS_NO_DISCOUNT                int32 = 5071 //用户无优惠购买资格
	ECODE_MALL_HAS_NO_RIGHT                   int32 = 5072 //用户领取资格
	ECODE_MALL_MONTH_CARD_OVER                int32 = 5073 //用户月卡购买上限
	ECODE_MALL_CREATE_ORDER_FORBIDDEN         int32 = 5074 //禁止特定地区的玩家下单
	ECODE_MALL_VALID_END                      int32 = 5080 //占位

	//登录相关
	ECODE_LOGIN_FAKE_INFO                 int32 = 5100 //非法构造请求
	ECODE_LOGIN_TOKEN_FAIL                int32 = 5101 //token校验未通过
	ECODE_LOGIN_THIRD_TOKEN_FAIL          int32 = 5102 //第三方token校验未通过
	ECODE_LOGIN_ACCOUNT_INFO_ERR          int32 = 5103 //账户信息不匹配
	ECODE_LOGIN_TYPE_ERR                  int32 = 5104 //账户类型信息有误
	ECODE_LOGIN_MONEY_GET                 int32 = 5105 //获取金币信息失败
	ECODE_LOGIN_REGITER_FAIL              int32 = 5106 //账号注册失败
	ECODE_LOGIN_ACCOUNT_NOT_EXIST         int32 = 5107 //账号不存在
	ECODE_LOGIN_ACCOUNT_EXIST             int32 = 5108 //账号存在
	ECODE_LOGIN_SET_ONLINE_FAILD          int32 = 5109 //设置在线信息失败
	ECODE_LOGIN_EMAIL_ERR                 int32 = 5110 //邮箱格式错误
	ECODE_LOGIN_SEND_MAIL_FAIL            int32 = 5111 //邮件发送失败
	ECODE_LOGIN_SEND_MAIL_TOO_FAST        int32 = 5112 //邮件发送过快
	ECODE_LOGIN_UN_SEND_VERIFICATION_CODE int32 = 5113 //未发送验证码
	ECODE_LOGIN_VERIFICATION_CODE_EXPIRED int32 = 5114 //验证码已过期
	ECODE_LOGIN_VERIFICATION_CODE_ERR     int32 = 5115 //验证码错误
	ECODE_LOGIN_PASSWORD_ERR              int32 = 5116 //密码格式错误或密码错误
	ECODE_LOGIN_VERIFICATION_CODE_USE     int32 = 5117 //验证码已校验
	ECODE_LOGIN_GUEST_NOT_EXIST           int32 = 5118 //游客账号不存在
	ECODE_LOGIN_PASSWORD_CANNOT_CHINESE   int32 = 5119 //密码不能为中文
	ECODE_LOGIN_IP_REGISTER_LIMIT         int32 = 5220 //注册ip上限

	//比赛领奖相关
	ECODE_MATCH_COLLECT_ONLY_COMMON int32 = 5150 //只能收藏自由场牌局记录
	ECODE_MATCH_AWARD_ALREADY       int32 = 5151 //领奖时，奖品已经领取
	ECODE_MATCH_AWARD_STATUS_WRONG  int32 = 5152 //领奖时，非可领取状态
	ECODE_MATCH_AWARD_GET_MONEY_ERR int32 = 5153 //获取金币信息失败
	ECODE_MATCH_AWARD_DELIVERY_ERR  int32 = 5154 //发货失败

	//签到相关
	ECODE_SIGNIN_AWARD_ALREADY      int32 = 5250 //奖品已经领取
	ECODE_SIGNIN_AWARD_STATUS_WRONG int32 = 5251 //领奖时，非可领取状态
	ECODE_SIGNIN_GET_MONEY_ERR      int32 = 5252 //获取金币信息失败
	ECODE_SIGNIN_DELIVERY_ERR       int32 = 5253 //发货失败
	ECODE_SIGNIN_REPAIR_CARD_ERR    int32 = 5254 //补签卡扣除失败
	ECODE_SIGNIN_REPAIR_MONEY_ERR   int32 = 5255 //补签扣除金币失败
	ECODE_SIGNIN_REPAIR_MONEY_LACK  int32 = 5256 //补签金币不足
	ECODE_SIGNIN_NEWER_EXPIRE       int32 = 5257 //注册日期超过活动期限

	//任务相关
	ECODE_TASK_AWARD_ALREADY      int32 = 5300 //奖品已经领取
	ECODE_TASK_AWARD_STATUS_WRONG int32 = 5301 //领奖时，非可领取状态
	ECODE_TASK_GET_MONEY_ERR      int32 = 5302 //获取用户金币信息失败
	ECODE_TASK_DELIVERY_ERR       int32 = 5302 //发货失败
	ECODE_TASK_NEWER_EXPIRE       int32 = 5303 //新手任务已经超过活动期限

	//背包相关
	ECODE_USERITME_NOT_OWN                  int32 = 5350 //用户不拥有此物品
	ECODE_USERITME_EXPIRED                  int32 = 5351 //物品已过期
	ECODE_USERITME_NOT_ENOUGH               int32 = 5352 //物品数量不够
	ECODE_USERITME_USE_GIFT_ITEM_ERR        int32 = 5353 //使用礼包物品出错
	ECODE_USERITME_CAN_NOT_SOLD             int32 = 5354 //该物品不支持出售
	ECODE_USERITEM_GET_MONEY_FAILED         int32 = 5355 //取用户金币信息出错
	ECODE_USERITEM_SOLD_FAILED              int32 = 5356 //出售物品时扣除物品失败
	ECODE_USERITEM_DELIVER_FAILED_WHEN_SOLD int32 = 5357 //出售物品时发货失败
	ECODE_USERITEM_NOT_EXIST                int32 = 5358 //对应物品不存在
	ECODE_USERITEM_IN_BLACKLIST             int32 = 5359 //在黑名单中
	ECODE_USERITEM_Anim_NOT_EXIST           int32 = 5360 //在黑名单中
	ECODE_USERITEM_ALREADY_HAD              int32 = 5555 //物品已经拥有

	//支付中心相关
	ECODE_PAYCENTER_PRICE_WRONG        int32 = 5400 //价格不匹配
	ECODE_PAYCENTER_ORDER_INFO_WRONG   int32 = 5401 //订单信息不正确
	ECODE_PAYCENTER_ORDER_EXPIRE       int32 = 5402 //订单信息不正确
	ECODE_PAYCENTER_CURRENCY_WRONG     int32 = 5403 //支付方式不正确
	ECODE_PAYCENTER_ORDER_STATUS_WRONG int32 = 5404 //订单状态非法
	ECODE_PAYCENTER_GET_MONEY_ERR      int32 = 5405 //获取金币信息出错
	ECODE_PAYCENTER_GET_GOODS_ERR      int32 = 5406 //获取商品信息出错
	ECODE_PAYCENTER_DELIVERY_ERR       int32 = 5407 //获取商品信息出错

	//money服务相关
	ECODE_MONEY_NOT_ENOUGH   int32 = 5450 //金币不足
	ECODE_MONEY_REGISTER_ERR int32 = 5452 //注册金币记录错误（aff==0）
	ECODE_MONEY_STOPPED      int32 = 5453 //金币服务已停止
	ECODE_MONEY_GET_ERROR    int32 = 5454 //金币服务取数据时错误

	//状态服相关
	ECODE_ONLINE_SET_ERR int32 = 5500 //设置状态时出错
	ECODE_ONLINE_GET_ERR int32 = 5501 //获取状态信息时出错
	ECODE_ONLINE_DEL_ERR int32 = 5502 //删除状态信息时出错
	ECODE_ONLINE_STOPPED int32 = 5503 //已停服

	//userInfo服务相关
	ECODE_USERINFO_CHANGE_OTHERS_UID int32 = 5550 //不能修改其他人的信息
	ECODE_USERINFO_GET_MONEY_ERR     int32 = 5551 //获取金币信息失败
	ECODE_USERINFO_UPDATE_NO_CHANGE  int32 = 5552 //修改信息无变化
	ECODE_USERINFO_NICK_EMPTY        int32 = 5553 //昵称为空
	ECODE_USERINFO_NOT_EXIST         int32 = 5554 //用户不存在

	//gameLog服务相关
	ECODE_GAMELOG_CARDS_INFO_ERR int32 = 5600 //牌数组格式有误
	ECODE_GAMELOG_INFO_ERR       int32 = 5601 //牌数组格式有误

	//图片上传服务
	ECODE_IMGUP_ENCODE_ERR     int32 = 5650 //图片编码有误
	ECODE_IMGUP_WRITE_FILE_ERR int32 = 5651 //写图片文件失败

	// 桌子服务
	ECODE_TABLE_INFO_GET_MONEY_ERR   int32 = 5700 //获取金币信息失败
	ECODE_TABLE_INFO_GET_TABLE_ERR   int32 = 5701 //获取桌子信息失败
	ECODE_TABLE_INFO_GET_UID         int32 = 5702 //获取桌子信息时未找到UID
	ECODE_TABLE_INFO_GET_MONEY       int32 = 5703 //获取桌子信息时查询金币信息失败
	ECODE_TABLE_LIST_ILLEGAL_PARAMS  int32 = 5704 //请求桌子列表时箱子号错误
	ECODE_TABLE_LIST_TABLE_NOT_FOUND int32 = 5705 //根据桌子id找不到桌子

	//ClienConfig服务
	ECODE_CLIENT_CFG_GAME_NO_DATA   int32 = 5750 //本大厅版本+子游戏，无对应资源包
	ECODE_CLIENT_CFG_COMM_NO_DATA   int32 = 5751 // 未找到对应公共包数据
	ECODE_CLIENT_CFG_GAME_LANG_ERR  int32 = 5752 //未找到合适的子游戏语言包数据
	ECODE_CLIENT_CFG_COMM_LANG_ERR  int32 = 5753 //未找到合适的公共语言包数据
	ECODE_CLIENT_CFG_DIFF_GAME      int32 = 5754 //未找到游戏差分信息
	ECODE_CLIENT_CFG_DIFF_COMM      int32 = 5755 //未找到公共包差分信息
	ECODE_CLIENT_CFG_DIFF_GAME_LANG int32 = 5756 //未找到游戏语言包差分数据
	ECODE_CLIENT_CFG_DIFF_COMM_LANG int32 = 5757 //未找到公共语言包差分数据

	//Avatar服务相关
	ECODE_AVATAR_HERO_NOT_MATCH  int32 = 5800 //角色不匹配
	ECODE_AVATAR_ITEM_EXPIRED    int32 = 5801 //物品已过期
	ECODE_AVATAR_ITEM_OWNED      int32 = 5802 //用户已经拥有此物品
	ECODE_AVATAR_ITEM_CANNOT_BUY int32 = 5803 //此物品不能被购买
	ECODE_AVATAR_ITEM_NOT_OWN    int32 = 5804 //此物品不能被购买
	ECODE_AVATAR_HERO_OWNED      int32 = 5805 //用户已经拥有此英雄
	ECODE_AVATAR_ILLEGAL         int32 = 5806 //用户形象不合法
	ECODE_AVATAR_REQUEST_LIMIT   int32 = 5807 //装扮索要上线

	//邮箱服务相关
	ECODE_MAIL_BATCH_FAIL                 int32 = 5900 //批量发送全部失败
	ECODE_MAIL_BATCH_PART_FAIL            int32 = 5901 //批量发送部分失败
	ECODE_MAIL_GIFT_GOT                   int32 = 5902 //物品已领取
	ECODE_MAIL_NOT_EXIST                  int32 = 5903 //邮件不存在
	ECODE_MAIL_GIFT_CANNOT_SEND           int32 = 5904 //邮件存在无法发放的物品
	ECODE_MAIL_DEL_WITH_GIFTS             int32 = 5905 //批量删除存在未领取物品
	ECODE_MAIL_GIFTS_STATUS_ERR           int32 = 5906 //礼包邮件非可领取状态
	ECODE_MAIL_GIFTS_EXPIRE               int32 = 5906 //礼包邮件已经过期
	ECODE_FRIEND_COUNT_MSG                int32 = 5907 //统计新读消息数量时发生错误
	ECODE_FRIEND_COUNT_REQUEST            int32 = 5908 //统计新好友请求数量时发生错误
	ECODE_FRIEND_UPDATE_MSG_STATUS_FAILED int32 = 5909 //更新消息状态失败

	//Badge服务相关
	ECODE_BADGE_ITEM_OWNED   int32 = 5950 //用户已经拥有此物品
	ECODE_BADGE_NOT_OWNED    int32 = 5951 //用户不拥有此物品
	ECODE_BADGE_ITEM_EXPIRED int32 = 5952 //物品已过期

)

//天梯系统的错误码
const (
	ECODE_LADDER_NO_UID                     int32 = 6000 //取用户ID失败
	ECODE_LADDER_LEVELS_CFG_GET_FAILD       int32 = 6001 //取配置数据失败
	ECODE_LADDER_SEASON_GET_FIALED          int32 = 6002 //取当前的赛季失败
	ECODE_LADDER_USER_SCORE_GET_FIALED      int32 = 6003 //取用户的大师分数据失败
	ECODE_LADDER_USER_PREV_SCORE_GET_FIALED int32 = 6004 //取用户前一期的大师分数据失败
	ECODE_LADDER_NO_REWARD                        = 6005 //没有奖励可领取
	ECODE_LADDER_DELIVER_FAILED                   = 6006 //奖励发放失败
	ECODE_LADDER_USER_HISTORY_GET_FAILED          = 6007 //历史战绩查询失败
	ECODE_LADDER_RANK_GET_FAILED                  = 6008 //排行榜查询失败
	ECODE_LADDER_GET_USER_RANK_FAILED             = 6009 //取用户排名失败
	ECODE_LADDER_CALC_REWARDS                     = 6010 //计算用户奖励失败
	ECODE_LADDER_GET_USER_BEST_DATA               = 6011 //取用户最佳赛季数据失败
)

//	邀请好友相关
const (
	ECODE_FRIEND_INVITE_INVALID_CODE    int32 = 6050 //	无效邀请码
	ECODE_FRIEND_INVITE_USER_BIND       int32 = 6051 //	用户已绑定
	ECODE_FRIEND_INVITE_TIME_OUT        int32 = 6052 //	邀请超时
	ECODE_FRIEND_INVITE_WAIT_BIND       int32 = 6053 //	待绑定
	ECODE_FRIEND_INVITE_RECEIVED        int32 = 6054 //	被邀请人已领取奖励
	ECODE_FRIEND_INVITE_RECORD_NOT_FIND int32 = 6055 //	邀请记录未找到
	ECODE_FRIEND_INVITE_NOT_REACHED     int32 = 6056 //	邀请未达标
	ECODE_FRIEND_INVITE_USER_RECEIVED   int32 = 6057 //	邀请人已领取
	ECODE_FRIEND_INVITE_STANDARD        int32 = 6058 //	被邀请人已达标
	ECODE_FRIEND_INVITE_TIMEOUT         int32 = 6059 //	不能填写比自己注册时间晚的玩家邀请码
)

//家园相关
const (
	ECODE_HOMELAND_HAVE_NO_PET           int32 = 6100 //玩家没有宠物
	ECODE_HOMELAND_ALREADY_INIT          int32 = 6101 //玩家已经初始化家园(领取过了宠物)
	ECODE_HOMELAND_DEDUCT_CARD_ERR       int32 = 6102 //扣除家园道具失败
	ECODE_HOMELAND_HAVE_HELPED           int32 = 6103 //已经帮助过该uid
	ECODE_HOMELAND_HELPED_TO_MAX         int32 = 6104 //被帮助达到了上限
	ECODE_HOMELAND_TREE_HAVE_BE_STOLENED int32 = 6105 //表示玩家摇钱树已经被偷过了(被偷成功到达上限了)
	ECODE_HOMELAND_TREE_I_HAVE_STEAL     int32 = 6106 //表示玩家A已经偷过玩家B的摇钱树了（不管偷成功还是偷失败了）
	ECODE_HOMELAND_TREE_NOT_RIPE         int32 = 6107 //表示玩家摇钱树还没成熟
	ECODE_HOMELAND_CHIPS_NOT_ENOUGH      int32 = 6108 //玩家偷摇钱树的时候筹码不足
	ECODE_HOMELAND_PET_PROTECT           int32 = 6109 //好友的摇钱树正在被保护中，不能被偷走果实
	ECODE_HOMELAND_HAVE_REWARDED         int32 = 6110 //表示玩家已经领取过奖励了
	ECODE_HOMELAND_NO_REWARDS            int32 = 6111 //表示没有奖励可以领取
	ECODE_HOMELAND_COIN_NOT_ENOUGH       int32 = 6112 //表示玩家的金币不足，兑换不到满足升级的经验
	ECODE_HOMELAND_I_STEAL_TO_MAX        int32 = 6113 //表示当天玩家偷摇钱树已经到达上限值了
)

//代理服务
const (
	ECODE_PROXY_INVALID_PARAM      int32 = 7000 //参数有误
	ECODE_PROXY_SERVER_UNSUPPORTED int32 = 7001 //不支持的服务
	ECODE_PROXY_FAILED                   = 7002 //转发失败
)

//好友服务相关
const (
	ECODE_FRIEND_ADD_FAILED_USER_NOT_EXISTS = 7101 //用户不存在
	ECODE_FRIEND_ADD_FAILED_OF_EXISTS       = 7102 //加好友失败：已经是好友了
	ECODE_FRIEND_ADD_FAILED_OF_PRIVACY      = 7103 //加好友失败：对方不允许添加好友
	ECODE_FRIEND_ADD_FAILED_OF_BLACKLIST    = 7104 //加好友失败：对方已将您添加到了黑名单中
	ECODE_FRIEND_ADD_FAILED_NUM_LIMITED     = 7105 //加好友失败：对方好友数量达到上限
	ECODE_FRIEND_ADD_FAILED_MY_NUM_LIMITED  = 7106 //加好友失败：自己好友数量达到上限
	ECODE_FRIEND_CHAT_FAILED_NOT_FRIEND     = 7107 //发聊天消息失败：不是好友
	ECODE_FRIEND_CHAT_FAILED_IN_BLACKLIST   = 7108 //发聊天消息失败：被对方加入了黑名单
	ECODE_FRIEND_OPERATION_TOO_FAST         = 7109 //操作过于频繁
	ECODE_FRIEND_INVITER_NOT_IN_GAME        = 7110 //邀请失败：当前用户不在房间内
	ECODE_FRIEND_INVITER_IS_PLAYING         = 7111 //邀请失败：当前用户不是旁观状态
	ECODE_FRIEND_INVITE_TARGET_NOT_ALLOWED  = 7112 //邀请失败：被邀请用户不允许被邀请
	ECODE_FRIEND_INVITE_TARGET_NOT_ONLINE   = 7113 //邀请失败：被邀请用户不在线
	ECODE_FRIEND_INVITE_TARGET_IS_PLAYING   = 7114 //邀请失败：被邀请用户正在玩牌
	ECODE_FRIEND_ADD_FAILED_OF_SELF         = 7115 //加好友失败：添加自己
	ECODE_FRIEND_PROCESS_FAILED_NOT_EXISTS  = 7116 //处理好友请求：申请记录不存在
	ECODE_FRIEND_PROCESS_FAILED_DONE        = 7117 //处理好友请求：申请已经处理过了
)

//黑名单
const (
	ECODE_BLACKLIST_NUMBER_LIMIT = 7201 //数量达到最大限制
)

//好友房
const (
	ECODE_CUSTOM_GAME_RECONNECT_INVALID     = 7301 //重连房间：当前不在指定的好友房
	ECODE_CUSTOM_GAME_ENTER_WRONG_PWD       = 7302 //进入房间：口令不正确
	ECODE_CUSTOM_GAME_CREATE_INVALID_CONFIG = 7303 //创建房间：房间配置无效
	ECODE_CUSTOM_GAME_CREATE_RETRY          = 7304 //创建房间：私人房达到最大数量限制
	ECODE_CUSTOM_GAME_CREATE_NO_TABLE       = 7305 //创建房间：当前服务无桌子可用
	ECODE_CUSTOM_GAME_TABLE_NOT_EXIST       = 7306 //创建房间/进入房间：找不到该桌子
	ECODE_CUSTOM_GAME_ENTER_TABLE_INVALID   = 7307 //进入房间：桌子已解散
	ECODE_CUSTOM_GAME_CREATE_MONEY_LIMIT    = 7308 //创建房间：筹码数量不足
	ECODE_CUSTOM_GAME_CREATE_NO_COINS       = 7309 //创建房间：金币数量不足
	ECODE_CUSTOM_GAME_EXP_LEVEL_LIMIT       = 7309 //创建或进入房间：级别不足
	ECODE_CUSTOM_GAME_CREATE_NOT_VIP        = 7310 //创建房间：不是白名单用户不能创建白名单房间
)

//百人场
const (
	ECODE_GAMBLE_CANT_LEAVE               = 7401 //游戏中不能离开
	ECODE_GAMBLE_DEALER_ANTES_INVALID     = 7402 //上庄金额无效
	ECODE_GAMBLE_ACTION_BAD_TIME          = 7403 //操作无效，当前非可操作时间
	ECODE_GAMBLE_ACTION_INVALID_OPTION    = 7404 //下注选项无效
	ECODE_GAMBLE_ACTION_INVALID_MONEY     = 7405 //下注金额无效
	ECODE_GAMBLE_ACTION_MONEY_LIMITED     = 7406 //下注金额被限
	ECODE_GAMBLE_ACTION_MONEY_NOT_ENOUGH  = 7406 //筹码不足
	ECODE_GAMBLE_DELAER_CAN_NOT_EXIT      = 7407 //庄家不能离开，需要先下庄
	ECODE_GAMBLE_NEXT_DEALER_CAN_NOT_EXIT = 7408 //下一任庄家不能离开
	ECODE_GAMBLE_DEALER_CAN_NOT_CANCEL    = 7410 //下一任庄家不能取消排庄
)

//活动服务
const (
	ECODE_ACTIVITY_NOT_IN            = 7400 //没有活动资格
	ECODE_ACTIVITY_NOT_RIGHT         = 7401 //没有领取资格
	ECODE_ACTIVITY_AWARD_ALREADY_GOT = 7402 //已经领取过了
	ECODE_ACTIVITY_AWARD_TIME_NOT_OK = 7403 //领取时间未到
)

//跑牌
const (
	ECODE_PAOPAI_ILLEGAL_CARDS        = 7501 //非法的操作
	ECODE_PAOPAI_USER_HAS_BEEN_READY  = 7502 //玩家已经准备过了，不需要重新准备；
	ECODE_PAOPAI_NOT_HAVE_THESE_CARDS = 7503 //玩家没有该牌（可能是已经是自动打出的牌再次被玩家打出）。
	ECODE_PAOPAI_USER_IS_GAMING       = 7504 //玩家正在玩牌中。（玩牌中的玩家不允许退出）
)

//大米
const (
	ECODE_DUMMY_ILLEGAL_CARDS        = 7601 //非法的操作
	ECODE_DUMMY_USER_HAS_BEEN_READY  = 7602 //玩家已经准备过了，不需要重新准备；
	ECODE_DUMMY_NOT_HAVE_THESE_CARDS = 7603 //玩家没有该牌（可能是已经是自动打出的牌再次被玩家打出）。
	ECODE_DUMMY_USER_IS_GAMING       = 7604 //玩家正在玩牌中。（玩牌中的玩家不允许退出）
)

//KING牌
const (
	ECODE_KING_ILLEGAL_CARDS        = 7701 //非法的操作
	ECODE_KING_USER_HAS_BEEN_READY  = 7702 //玩家已经准备过了，不需要重新准备；
	ECODE_KING_NOT_HAVE_THESE_CARDS = 7703 //玩家没有该牌（可能是已经是自动打出的牌再次被玩家打出）。
	ECODE_KING_USER_IS_GAMING       = 7704 //玩家正在玩牌中。（玩牌中的玩家不允许退出）
	ECODE_KING_ILLEGAL_OPCODE       = 7705 //非法的游戏操作码
)

const (
	ECODE_ILLEGAL_STAGE = 7701 //游戏阶段不正确
)

//模型服务
const (
	ECODE_MODEL_ERROR  int32 = 100000
	ECODE_MODEL_NO_ROW int32 = 100001
)

const (
	TIMEOUT_NORMAL    time.Duration = 10 * time.Second
	TIMEOUT_LONG      time.Duration = 30 * time.Second
	TIMEOUT_GAME_LONG time.Duration = 30 * time.Second
)
const (
	SYS_RESP_SUCCESS             = 0 //0表示成功
	SYS_RESP_CODE_ILLEGAL_REQ    = 1 //1表示请求异常(请求参数在acc层就解析不合法)
	SYS_RESP_CODE_TIME_OUT       = 2 //2执行超时
	SYS_RESP_CODE_INTERNAL_ERR   = 3 //服务内部错误比如说拿不到在线状态
	SYS_RESP_CODE_UNSAFE_REQ     = 4 //不安全的请求,eg:未经过合法登录就发过来的请求或者是黑名单中的请求
	SYS_RESP_CODE_CONCURRENT_REQ = 5 //并发的请求,eg:玩家连续的点击报名,并发的
)

//redis连接配置(对应阿波罗配置key)
const (
	REDIS_CONF_R1     = "redis.pool.r1"
	REDIS_CONF_R2     = "redis.pool.r2"
	REDIS_CONF_R3     = "redis.pool.r3"     //用于用户数据的数据库缓存
	REDIS_CONF_R4     = "redis.pool.r4"     //系统数据
	REDIS_CONF_R5     = "redis.pool.r5"     //在线信息专用
	REDIS_CONF_R6     = "redis.pool.r6"     //机器人相关信息专用
	REDIS_CONF_R7     = "redis.pool.r7"     //统计功能专用
	REDIS_CONF_R8     = "redis.pool.r8"     //限制类数据专用
	REDIS_CONF_R9     = "redis.pool.r9"     //社交相关数据专用
	REDIS_CONF_R9P1   = "redis.pool.r9p1"   //广播相关
	REDIS_CONF_R10    = "redis.pool.r10"    //游戏服务专用
	REDIS_CONF_R10P1  = "redis.pool.r10p1"  //AOF游戏专用
	REDIS_CONF_R10P2  = "redis.pool.r10p2"  //三公游戏专用
	REDIS_CONF_R10P3  = "redis.pool.r10p3"  //跑牌游戏专用
	REDIS_CONF_R10P4  = "redis.pool.r10p4"  //炸金花游戏专用
	REDIS_CONF_R10P5  = "redis.pool.r10p5"  //多米诺游戏专用
	REDIS_CONF_R10P6  = "redis.pool.r10p6"  //十三张游戏专用
	REDIS_CONF_R10P7  = "redis.pool.r10p7"  //二十一点游戏专用
	REDIS_CONF_R10P8  = "redis.pool.r10p8"  //大米游戏专用
	REDIS_CONF_R10P9  = "redis.pool.r10p9"  //斗地主不洗牌游戏专用
	REDIS_CONF_R10P10 = "redis.pool.r10p10" //斗地主叫地主游戏专用
	REDIS_CONF_R10P11 = "redis.pool.r10p11" //斗地主抢地主游戏专用
	REDIS_CONF_R10P12 = "redis.pool.r10p12" //斗地主癞子游戏专用
	REDIS_CONF_R10P13 = "redis.pool.r10p13" //King牌游戏专用
	REDIS_CONF_R11    = "redis.pool.r11"    //推送服务专用
	REDIS_CONF_R12    = "redis.pool.r12"    //普通场桌子数据专用
	REDIS_CONF_R12P1  = "redis.pool.r12p1"  //AOF桌子数据专用
	REDIS_CONF_R12P2  = "redis.pool.r12p2"  //三公桌子数据专用
	REDIS_CONF_R12P3  = "redis.pool.r12p3"  //跑牌桌子数据专用
	REDIS_CONF_R12P4  = "redis.pool.r12p4"  //炸金花桌子数据专用
	REDIS_CONF_R12P5  = "redis.pool.r12p5"  //多米诺99桌子数据专用
	REDIS_CONF_R12P6  = "redis.pool.r12p6"  //十三张桌子数据专用
	REDIS_CONF_R12P8  = "redis.pool.r12p8"  //大米桌子数据专用
	REDIS_CONF_R12P9  = "redis.pool.r12p9"  //斗地主不洗牌桌子数据专用
	REDIS_CONF_R12P10 = "redis.pool.r12p10" //斗地主叫地主桌子数据专用
	REDIS_CONF_R12P11 = "redis.pool.r12p11" //斗地主抢地主桌子数据专用
	REDIS_CONF_R12P12 = "redis.pool.r12p12" //斗地主癞子桌子数据专用
	REDIS_CONF_R12P13 = "redis.pool.r12p13" //King牌桌子数据专用
	REDIS_CONF_R13    = "redis.pool.r13"    //家园相关数据专用

	REDIS_CONF_R10SLV0   = "redis.pool.r10slv0"   //选场时找指定人数的桌子专用
	REDIS_CONF_R10SLV1   = "redis.pool.r10slv1"   //选场时找指定人数的桌子专用
	REDIS_CONF_R10SLV2   = "redis.pool.r10slv2"   //选场时找指定人数的桌子专用
	REDIS_CONF_R10SLV3   = "redis.pool.r10slv3"   //选场时找指定人数的桌子专用
	REDIS_CONF_R10SLV4   = "redis.pool.r10slv4"   //选场时找指定人数的桌子专用
	REDIS_CONF_R10SLV5   = "redis.pool.r10slv5"   //选场时找指定人数的桌子专用
	REDIS_CONF_R10SLV6   = "redis.pool.r10slv6"   //选场时找指定人数的桌子专用
	REDIS_CONF_R10SLV7   = "redis.pool.r10slv7"   //选场时找指定人数的桌子专用
	REDIS_CONF_R10SLV8   = "redis.pool.r10slv8"   //选场时找指定人数的桌子专用
	REDIS_CONF_R10SLV9   = "redis.pool.r10slv9"   //选场时找指定人数的桌子专用
	REDIS_CONF_R10P1SLV0 = "redis.pool.r10p1slv0" //选场时找指定人数的桌子专用
	REDIS_CONF_R10P1SLV1 = "redis.pool.r10p1slv1" //选场时找指定人数的桌子专用
	REDIS_CONF_R10P1SLV2 = "redis.pool.r10p1slv2" //选场时找指定人数的桌子专用
	REDIS_CONF_R10P1SLV3 = "redis.pool.r10p1slv3" //选场时找指定人数的桌子专用
	REDIS_CONF_R10P1SLV4 = "redis.pool.r10p1slv4" //选场时找指定人数的桌子专用
	REDIS_CONF_R10P1SLV5 = "redis.pool.r10p1slv5" //选场时找指定人数的桌子专用
	REDIS_CONF_R10P1SLV6 = "redis.pool.r10p1slv6" //选场时找指定人数的桌子专用
	REDIS_CONF_R10P1SLV7 = "redis.pool.r10p1slv7" //选场时找指定人数的桌子专用
	REDIS_CONF_R10P1SLV8 = "redis.pool.r10p1slv8" //选场时找指定人数的桌子专用
	REDIS_CONF_R10P1SLV9 = "redis.pool.r10p1slv9" //选场时找指定人数的桌子专用
	REDIS_CONF_R10P2SLV0 = "redis.pool.r10p2slv0" //【三公】选场时找指定人数的桌子专用
	REDIS_CONF_R10P2SLV1 = "redis.pool.r10p2slv1" //【三公】选场时找指定人数的桌子专用
	REDIS_CONF_R10P2SLV2 = "redis.pool.r10p2slv2" //【三公】选场时找指定人数的桌子专用
	REDIS_CONF_R10P2SLV3 = "redis.pool.r10p2slv3" //【三公】选场时找指定人数的桌子专用
	REDIS_CONF_R10P2SLV4 = "redis.pool.r10p2slv4" //【三公】选场时找指定人数的桌子专用
	REDIS_CONF_R10P2SLV5 = "redis.pool.r10p2slv5" //【三公】选场时找指定人数的桌子专用
	REDIS_CONF_R10P2SLV6 = "redis.pool.r10p2slv6" //【三公】选场时找指定人数的桌子专用
	REDIS_CONF_R10P2SLV7 = "redis.pool.r10p2slv7" //【三公】选场时找指定人数的桌子专用
	REDIS_CONF_R10P2SLV8 = "redis.pool.r10p2slv8" //【三公】选场时找指定人数的桌子专用
	REDIS_CONF_R10P2SLV9 = "redis.pool.r10p2slv9" //【三公】选场时找指定人数的桌子专用
	REDIS_CONF_R10P4SLV0 = "redis.pool.r10p4slv0" //【炸金花】选场时找指定人数的桌子专用
	REDIS_CONF_R10P4SLV1 = "redis.pool.r10p4slv1" //【炸金花】选场时找指定人数的桌子专用
	REDIS_CONF_R10P4SLV2 = "redis.pool.r10p4slv2" //【炸金花】选场时找指定人数的桌子专用
	REDIS_CONF_R10P4SLV3 = "redis.pool.r10p4slv3" //【炸金花】选场时找指定人数的桌子专用
	REDIS_CONF_R10P4SLV4 = "redis.pool.r10p4slv4" //【炸金花】选场时找指定人数的桌子专用
	REDIS_CONF_R10P4SLV5 = "redis.pool.r10p4slv5" //【炸金花】选场时找指定人数的桌子专用
	REDIS_CONF_R10P4SLV6 = "redis.pool.r10p4slv6" //【炸金花】选场时找指定人数的桌子专用
	REDIS_CONF_R10P4SLV7 = "redis.pool.r10p4slv7" //【炸金花】选场时找指定人数的桌子专用
	REDIS_CONF_R10P4SLV8 = "redis.pool.r10p4slv8" //【炸金花】选场时找指定人数的桌子专用
	REDIS_CONF_R10P4SLV9 = "redis.pool.r10p4slv9" //【炸金花】选场时找指定人数的桌子专用
	REDIS_CONF_R10P5SLV0 = "redis.pool.r10p5slv0" //【多米诺99】选场时找指定人数的桌子专用
	REDIS_CONF_R10P5SLV1 = "redis.pool.r10p5slv1" //【多米诺99】选场时找指定人数的桌子专用
	REDIS_CONF_R10P5SLV2 = "redis.pool.r10p5slv2" //【多米诺99】选场时找指定人数的桌子专用
	REDIS_CONF_R10P5SLV3 = "redis.pool.r10p5slv3" //【多米诺99】选场时找指定人数的桌子专用
	REDIS_CONF_R10P5SLV4 = "redis.pool.r10p5slv4" //【多米诺99】选场时找指定人数的桌子专用
	REDIS_CONF_R10P5SLV5 = "redis.pool.r10p5slv5" //【多米诺99】选场时找指定人数的桌子专用
	REDIS_CONF_R10P5SLV6 = "redis.pool.r10p5slv6" //【多米诺99】选场时找指定人数的桌子专用
	REDIS_CONF_R10P5SLV7 = "redis.pool.r10p5slv7" //【多米诺99】选场时找指定人数的桌子专用
	REDIS_CONF_R10P5SLV8 = "redis.pool.r10p5slv8" //【多米诺99】选场时找指定人数的桌子专用
	REDIS_CONF_R10P5SLV9 = "redis.pool.r10p5slv9" //【多米诺99】选场时找指定人数的桌子专用
	REDIS_CONF_R10P6SLV0 = "redis.pool.r10p6slv0" //【十三张】选场时找指定人数的桌子专用
	REDIS_CONF_R10P6SLV1 = "redis.pool.r10p6slv1" //【十三张】选场时找指定人数的桌子专用
	REDIS_CONF_R10P6SLV2 = "redis.pool.r10p6slv2" //【十三张】选场时找指定人数的桌子专用
	REDIS_CONF_R10P6SLV3 = "redis.pool.r10p6slv3" //【十三张】选场时找指定人数的桌子专用
	REDIS_CONF_R10P6SLV4 = "redis.pool.r10p6slv4" //【十三张】选场时找指定人数的桌子专用
	REDIS_CONF_R10P6SLV5 = "redis.pool.r10p6slv5" //【十三张】选场时找指定人数的桌子专用
	REDIS_CONF_R10P6SLV6 = "redis.pool.r10p6slv6" //【十三张】选场时找指定人数的桌子专用
	REDIS_CONF_R10P6SLV7 = "redis.pool.r10p6slv7" //【十三张】选场时找指定人数的桌子专用
	REDIS_CONF_R10P6SLV8 = "redis.pool.r10p6slv8" //【十三张】选场时找指定人数的桌子专用
	REDIS_CONF_R10P6SLV9 = "redis.pool.r10p6slv9" //【十三张】选场时找指定人数的桌子专用

	REDIS_CONF_R_TEST_YOGA = "redis.pool.r.test.yoga" // yoga test专用
)

//数据库连接配置(对应阿波罗配置key)
const (
	DB_CONF_USERDB       = "mysql.database.userdb"
	DB_CONF_USER_RECORDS = "mysql.database.user_records"

	DB_CONF_SUPER_ADMIN     = "mysql.database.superadmin"
	DB_CONF_SUPER_ADMIN_PRO = "mysql.database.superadmin_pro"
	DB_CONF_USER_ACTIVITY   = "mysql.database.user_activity"
	DB_CONF_SERVICE_DOG     = "mysql.database.service_dog"

	DB_CONF_CFR_PRFLOP  = "mysql.database.cfr_preflop"
	DB_CONF_CFR_FLOP    = "mysql.database.cfr_flop"
	DB_CONF_CFR_TURN    = "mysql.database.cfr_turn"
	DB_CONF_CFR_RIVER   = "mysql.database.cfr_river"
	DB_CONF_LOGS        = "mysql.database.logs"
	DB_CONF_SOCIAL_DATA = "mysql.database.social_data"
)

const (
	APOLLO_DATAPROXY_TEST = "address.dataProxy.test" //测试环境配置服务地址
	APOLLO_DATAPROXY_PRE  = "address.dataProxy.pre"  //预发布环境配置服务地址
	APOLLO_DATAPROXY_PRO  = "address.dataProxy.pro"  //正式环境配置服务地址
)

//不同服务使用的redis服务的对应关系
const (
	//非淘汰模式数据
	Redis_Online        = REDIS_CONF_R5  //玩家在线数据
	Redis_OnlineSlave   = REDIS_CONF_R5  //玩家在线数据复本(线上时需要设置为Redis_Online实例的从服)
	Redis_Ranking       = REDIS_CONF_R1  //排行榜数据
	Redis_RankMatchTime = REDIS_CONF_R1  //当前排位赛时间
	Redis_Robot         = REDIS_CONF_R6  //机器人相关信息
	Redis_Homeland      = REDIS_CONF_R13 //家园相关信息

	Redis_ResourcZip  = REDIS_CONF_R1 //提供给用户下载包的管理系统使用（用户请求会访问）
	Redis_AdminConfig = REDIS_CONF_R1 //AdminConfig默认使用此redis配置（需要单独配置再另行定义）

	Redis_Phone                   = REDIS_CONF_R1 //社交数据
	Redis_Mall                    = REDIS_CONF_R1
	Redis_TableInfo               = REDIS_CONF_R1 //场次配置
	Redis_Bankrupt                = REDIS_CONF_R1
	Redis_ThirdPay                = REDIS_CONF_R1
	Redis_Task                    = REDIS_CONF_R1 //任务配置信息
	Redis_ItemInfo                = REDIS_CONF_R1 //物品信息
	Redis_HeroInfo                = REDIS_CONF_R1 //角色信息
	Redis_SceneInfo               = REDIS_CONF_R1 //场景
	Redis_BadgeInfo               = REDIS_CONF_R1 //徽章
	Redis_GroupMail               = REDIS_CONF_R1 //群体邮件内容
	Redis_DressInfo               = REDIS_CONF_R1 //装扮信息
	Redis_SignIn                  = REDIS_CONF_R1 //签到
	Redis_TaskActiveRewards       = REDIS_CONF_R1 //任务活跃值奖励
	Redis_MatchConfig             = REDIS_CONF_R1 //比赛配置
	Redis_MatchPrizeGot           = REDIS_CONF_R1 //某场比赛各名次的领奖情况
	Redis_NoticConfig             = REDIS_CONF_R1 //公告/活动配置
	Redis_SamGongTableConfig      = REDIS_CONF_R1 //三公游戏配置
	Redis_ZJHTableConfig          = REDIS_CONF_R1 //炸金花游戏配置
	Redis_Domino99TableConfig     = REDIS_CONF_R1 //多米诺99游戏配置
	Redis_BlackJackTableConfig    = REDIS_CONF_R1 //二十一点游戏配置
	Redis_CapsaSusunTableConfig   = REDIS_CONF_R1 //十三张游戏配置
	Redis_LoginTokenKey           = REDIS_CONF_R1 //登录token生成时的密钥
	Redis_ItemTitleInfo           = REDIS_CONF_R1 //互动道具主题
	Redis_BroadCast               = REDIS_CONF_R1 //广播配置信息
	Redis_Game                    = REDIS_CONF_R1 //普通game的桌子状态
	Redis_AIConfig                = REDIS_CONF_R1 //AI配置
	Redis_Concurrent              = REDIS_CONF_R1 //存放防并发的redis key
	Redis_SupserAdmin             = REDIS_CONF_R2 //管理后台redis
	Redis_LadderLevels            = REDIS_CONF_R1 //天梯段位配置
	Redis_LadderEvents            = REDIS_CONF_R1 //天梯特殊事件配置
	Redis_LadderRank              = REDIS_CONF_R1 //天梯排行榜奖励配置
	Redis_LadderPeriod            = REDIS_CONF_R1 //天梯赛制周期配置
	Redis_LadderSeasons           = REDIS_CONF_R1 //天梯的赛季信息缓存
	Redis_LadderSeasonsV2         = REDIS_CONF_R1 //天梯的赛季信息缓存v2
	Redis_LadderLevelReward       = REDIS_CONF_R1 //天梯段位奖励
	Redis_TestCards               = REDIS_CONF_R1 //配牌数据
	Redis_ExpLevels               = REDIS_CONF_R1 //经验级别
	Redis_VipLevels               = REDIS_CONF_R1 //VIP经验级别
	Redis_PiggyBank               = REDIS_CONF_R1 //小猪储钱罐
	Redis_ScratchCard             = REDIS_CONF_R1 //刮刮卡
	Redis_MonthCard               = REDIS_CONF_R1 //月卡
	Redis_ReturnGift              = REDIS_CONF_R1 //回归礼包
	Redis_Hamster                 = REDIS_CONF_R1 //打地鼠
	Redis_SysNotices              = REDIS_CONF_R1 //系统公告
	Redis_RegisterRewards         = REDIS_CONF_R1 //新注发放
	Redis_HandBook                = REDIS_CONF_R1 //帮助信息
	Redis_Activity                = REDIS_CONF_R1 //帮助信息
	Redis_Activity_V2             = REDIS_CONF_R1 //新活动中心
	Redis_CustomizeConfig         = REDIS_CONF_R1 //自定义配置内容
	Redis_ClientLogCfg            = REDIS_CONF_R1 //用户日志上报配置
	Redis_BankruptThreshold       = REDIS_CONF_R1 //破产线配置
	Redis_UserDailyPlay           = REDIS_CONF_R1 //用户每日玩牌数据
	Redis_DailyFreeMoney          = REDIS_CONF_R7 //每日免费发放货币总量
	Redis_DailyPayMoney           = REDIS_CONF_R7 //每日付费发放货币总量
	Redis_DailyFee                = REDIS_CONF_R7 //每日回收货币总量(台费)
	Redis_DailyFieldTotalFee      = REDIS_CONF_R7 //每个场次回收货币总量 (0=>普通场 1=>MTT 2=>SNG)
	Redis_DailyGrantByField       = REDIS_CONF_R7 //发放系统细分，每个发放功能发放货币总量	包括钻石、金币和筹码
	Redis_FieldAll                = REDIS_CONF_R7 //所有发放功能 数据持久
	Redis_Field_People            = REDIS_CONF_R7 // 发放功能的领取人的集合
	Redis_SNG_Online              = REDIS_CONF_R1 //	SNG在线集合
	Redis_Ordinary_Online         = REDIS_CONF_R1 //	普通场在线集合(自由场)
	Redis_SNG_Play                = REDIS_CONF_R1 //	SNG在玩集合(自由场)
	Redis_Ordinary_Play           = REDIS_CONF_R1 //	普通场在玩集合(自由场)
	Redis_Ordinary_Table_Play     = REDIS_CONF_R1 //	普通场每个场次的在玩人数
	Redis_Agent_Uids              = REDIS_CONF_R7 //	需要记录分成的Uids
	Redis_SNG_Matches             = REDIS_CONF_R7 //	SNG开赛次数
	Redis_SNG_Total_Number        = REDIS_CONF_R7 //	SNG服务费总数量
	Redis_SNG_Mathch_People       = REDIS_CONF_R7 //	SNG比赛参赛集合
	Redis_SNG_Mathch_Count        = REDIS_CONF_R7 //	SNG比赛参赛次数集合
	Redis_Players_Play_Game_Total = REDIS_CONF_R7 //	玩家总玩牌数
	Redis_Average_Game_Time       = REDIS_CONF_R7 //	牌局平均时长
	Redis_Average_Game_Time_Field = REDIS_CONF_R1 //	分场次牌局平均时长
	Redis_Bottom_Pool_Distribute  = REDIS_CONF_R1 //	牌局底池分布
	Redis_Online_Users            = REDIS_CONF_R1 //	在线用户ID集合
	Redis_INVITE_ROBOT            = REDIS_CONF_R1 // 邀请机器人hash
	Redis_Random_Hero_Dress       = REDIS_CONF_R1 //	当前用户的随机角色装扮(hash) 有效期2小时
	Redis_DailyPayPeople          = REDIS_CONF_R7 //	每日付费人数集合
	Redis_DailyPayTotalMoney      = REDIS_CONF_R7 //	每日付费总金额集合
	Redis_StatHelperData          = REDIS_CONF_R7 //统计上报类的辅助数据
	Redis_Payment_Feedback_Status = REDIS_CONF_R3 //购买记录反馈状态
	Redis_Pay_Failed_Status       = REDIS_CONF_R3 //购买失败状态
	Redis_SocialData              = REDIS_CONF_R9 //社交数据

	//LRU淘汰模式数据
	Redis_UserInfo              = REDIS_CONF_R3 //用户数据缓存，LRU淘汰
	Redis_CFR                   = REDIS_CONF_R4 //CFR数据缓存，LRU淘汰
	Redis_SystemData            = REDIS_CONF_R4 //系统数据
	Redis_FriendCache           = REDIS_CONF_R3 //用户好友列表缓存
	Redis_GameTypeShowOnlineCnt = REDIS_CONF_R3 //场次列表中显示出来的在线人数

	//	发送邮件
	Redis_Send_Email = REDIS_CONF_R1 //	发送邮件(hash)

	Redis_Online_User_Ext_Data = REDIS_CONF_R7     //	用户在线扩展信息
	Redis_LimitData            = REDIS_CONF_R8     //限制数据
	Redis_GameData             = REDIS_CONF_R10    //游戏服务使用的各种数据
	Redis_AOFGameData          = REDIS_CONF_R10P1  //AOF游戏服务使用的各种数据
	Redis_SamgongGameData      = REDIS_CONF_R10P2  //三公游戏服务使用的各种数据
	Redis_PaoPaiGameData       = REDIS_CONF_R10P3  //跑牌游戏服务使用的各种数据
	Redis_ZJHGameData          = REDIS_CONF_R10P4  //炸金花游戏服务使用的各种数据
	Redis_Domino99GameData     = REDIS_CONF_R10P5  //多米诺99游戏服务使用的各种数据
	Redis_CapsaSusunGameData   = REDIS_CONF_R10P6  //十三张游戏服务使用的各种数据
	Redis_BlackJackGameData    = REDIS_CONF_R10P7  //二十一点游戏服务使用的各种数据
	Redis_DummyGameData        = REDIS_CONF_R10P8  //大米游戏服务使用的各种数据
	Redis_DdzbuGameData        = REDIS_CONF_R10P9  //斗地主-不洗牌游戏服务使用的各种数据
	Redis_DdzjiaoGameData      = REDIS_CONF_R10P10 //斗地主-叫地主游戏服务使用的各种数据
	Redis_DdzqiangGameData     = REDIS_CONF_R10P11 //斗地主-抢地主服务使用的各种数据
	Redis_DdzlaiGameData       = REDIS_CONF_R10P12 //斗地主-癞子游戏服务使用的各种数据
	Redis_KingGameData         = REDIS_CONF_R10P13 //king牌游戏服务使用的各种数据
	Redis_BroadcastData        = REDIS_CONF_R9P1   //广播数据
	Redis_PushData             = REDIS_CONF_R11    //推送服务数据
	Redis_StdTableData         = REDIS_CONF_R12    //普通场桌子数据
	Redis_AOFTableData         = REDIS_CONF_R12P1  //AOF桌子数据
	Redis_SamgongTableData     = REDIS_CONF_R12P2  //三公桌子数据
	Redis_PaoPaiTableData      = REDIS_CONF_R12P3  //跑牌桌子数据
	Redis_ZJHTableData         = REDIS_CONF_R12P4  //炸金花桌子数据
	Redis_Domino99TableData    = REDIS_CONF_R12P5  //多米诺99桌子数据
	Redis_CapsaSusunTableData  = REDIS_CONF_R12P6  //十三张张桌子数据
	Redis_DummyTableData       = REDIS_CONF_R12P8  //大米桌子数据
	Redis_DdzbuTableData       = REDIS_CONF_R12P9  //斗地主-不洗牌游戏桌子数据
	Redis_DdzjiaoTableData     = REDIS_CONF_R12P10 //斗地主-叫地主游戏桌子数据
	Redis_DdzqiangTableData    = REDIS_CONF_R12P11 //斗地主-抢地主游戏桌子数据
	Redis_DdzlaiTableData      = REDIS_CONF_R12P12 //斗地主-癞子游戏桌子数据
	Redis_KingTableData        = REDIS_CONF_R12P13 //king牌桌子数据

	//普通场的桌子人数有序集合
	Redis_TableInfoList0 = REDIS_CONF_R10SLV0 //存放桌子对应数值的数据（桌子列表展示使用）
	Redis_TableInfoList1 = REDIS_CONF_R10SLV1
	Redis_TableInfoList2 = REDIS_CONF_R10SLV2
	Redis_TableInfoList3 = REDIS_CONF_R10SLV3
	Redis_TableInfoList4 = REDIS_CONF_R10SLV4
	Redis_TableInfoList5 = REDIS_CONF_R10SLV5
	Redis_TableInfoList6 = REDIS_CONF_R10SLV6
	Redis_TableInfoList7 = REDIS_CONF_R10SLV7
	Redis_TableInfoList8 = REDIS_CONF_R10SLV8
	Redis_TableInfoList9 = REDIS_CONF_R10SLV9
	//AOF的桌子人数有序集合
	Redis_AOFTableInfoList0 = REDIS_CONF_R10P1SLV0 //存放桌子对应数值的数据（桌子列表展示使用）
	Redis_AOFTableInfoList1 = REDIS_CONF_R10P1SLV1
	Redis_AOFTableInfoList2 = REDIS_CONF_R10P1SLV2
	Redis_AOFTableInfoList3 = REDIS_CONF_R10P1SLV3
	Redis_AOFTableInfoList4 = REDIS_CONF_R10P1SLV4
	Redis_AOFTableInfoList5 = REDIS_CONF_R10P1SLV5
	Redis_AOFTableInfoList6 = REDIS_CONF_R10P1SLV6
	Redis_AOFTableInfoList7 = REDIS_CONF_R10P1SLV7
	Redis_AOFTableInfoList8 = REDIS_CONF_R10P1SLV8
	Redis_AOFTableInfoList9 = REDIS_CONF_R10P1SLV9
	//三公配桌数据
	Redis_SamgongGameDataSlave0 = REDIS_CONF_R10P2SLV0
	Redis_SamgongGameDataSlave1 = REDIS_CONF_R10P2SLV1
	Redis_SamgongGameDataSlave2 = REDIS_CONF_R10P2SLV2
	Redis_SamgongGameDataSlave3 = REDIS_CONF_R10P2SLV3
	Redis_SamgongGameDataSlave4 = REDIS_CONF_R10P2SLV4
	Redis_SamgongGameDataSlave5 = REDIS_CONF_R10P2SLV5
	Redis_SamgongGameDataSlave6 = REDIS_CONF_R10P2SLV6
	Redis_SamgongGameDataSlave7 = REDIS_CONF_R10P2SLV7
	Redis_SamgongGameDataSlave8 = REDIS_CONF_R10P2SLV8
	Redis_SamgongGameDataSlave9 = REDIS_CONF_R10P2SLV9
	//炸金花桌数据
	Redis_ZJHGameDataSlave0 = REDIS_CONF_R10P4SLV0
	Redis_ZJHGameDataSlave1 = REDIS_CONF_R10P4SLV1
	Redis_ZJHGameDataSlave2 = REDIS_CONF_R10P4SLV2
	Redis_ZJHGameDataSlave3 = REDIS_CONF_R10P4SLV3
	Redis_ZJHGameDataSlave4 = REDIS_CONF_R10P4SLV4
	Redis_ZJHGameDataSlave5 = REDIS_CONF_R10P4SLV5
	Redis_ZJHGameDataSlave6 = REDIS_CONF_R10P4SLV6
	Redis_ZJHGameDataSlave7 = REDIS_CONF_R10P4SLV7
	Redis_ZJHGameDataSlave8 = REDIS_CONF_R10P4SLV8
	Redis_ZJHGameDataSlave9 = REDIS_CONF_R10P4SLV9
	//多米诺99配桌数据
	Redis_DOMINO99GameDataSlave0 = REDIS_CONF_R10P5SLV0
	Redis_DOMINO99GameDataSlave1 = REDIS_CONF_R10P5SLV1
	Redis_DOMINO99GameDataSlave2 = REDIS_CONF_R10P5SLV2
	Redis_DOMINO99GameDataSlave3 = REDIS_CONF_R10P5SLV3
	Redis_DOMINO99GameDataSlave4 = REDIS_CONF_R10P5SLV4
	Redis_DOMINO99GameDataSlave5 = REDIS_CONF_R10P5SLV5
	Redis_DOMINO99GameDataSlave6 = REDIS_CONF_R10P5SLV6
	Redis_DOMINO99GameDataSlave7 = REDIS_CONF_R10P5SLV7
	Redis_DOMINO99GameDataSlave8 = REDIS_CONF_R10P5SLV8
	Redis_DOMINO99GameDataSlave9 = REDIS_CONF_R10P5SLV9
	//十三张配桌数据
	Redis_CapsaSusunGameDataSlave0 = REDIS_CONF_R10P6SLV0
	Redis_CapsaSusunGameDataSlave1 = REDIS_CONF_R10P6SLV1
	Redis_CapsaSusunGameDataSlave2 = REDIS_CONF_R10P6SLV2
	Redis_CapsaSusunGameDataSlave3 = REDIS_CONF_R10P6SLV3
	Redis_CapsaSusunGameDataSlave4 = REDIS_CONF_R10P6SLV4
	Redis_CapsaSusunGameDataSlave5 = REDIS_CONF_R10P6SLV5
	Redis_CapsaSusunGameDataSlave6 = REDIS_CONF_R10P6SLV6
	Redis_CapsaSusunGameDataSlave7 = REDIS_CONF_R10P6SLV7
	Redis_CapsaSusunGameDataSlave8 = REDIS_CONF_R10P6SLV8
	Redis_CapsaSusunGameDataSlave9 = REDIS_CONF_R10P6SLV9

	Redis_R_Test = REDIS_CONF_R_TEST_YOGA
)

//redis key的前缀设置- redis prefix
//缘由：统一设置redis的key前缀可以有效避免key冲突的情况。同时方便统一知晓系统redis使用情况
const (
	//状态服相关
	RP_onlineRecord    = "online_rec_%d"      //在线记录
	RP_onlineData      = "online_data_%d"     //在线记录额外数据
	RP_onlineUidSet    = "online_uids_p%d"    //在线UID列表(散列ONLINE_UID_SET_NUM个)
	RP_onlineSrvUidSet = "online_uids_s%dp%d" //某个Online服务的UID列表

	//比赛相关redis
	RP_SNGConfig     = "sng_config"       //eg: sng_config hash存放所有SNG比赛的配置
	RP_matchPriceGot = "match_prize_got_" //eg: match_price_got_{$tm} 存放某场比赛各名次的领奖情况

	//登录token密钥
	RP_loginTokenKey = "login_token_key" //存放生成token使用的密钥

	//自由桌相关redis
	RP_GameTableConfig  = "game_table_level_1001" //存放所有的场次配置信息
	RP_TexasProJackpot  = "texas_jackpot"         //专业场的jackpot数据
	RP_PrivateGameTable = "private_table"         //存放所有的场次配置信息
	RP_AOFGameTable     = "aof_table"             //存放AOF游戏所有的场次配置信息
	RP_gameLevelSvr     = "game_table_svr"        //hash 以level为key,存放svrID的数组
	RP_GambleConfig     = "gamble_config"         //存放百人场配置

	RP_cfrPreflopAllin = "cfr_pf_allin" ////CFR preflop allin选择信息

	//配置信息存放
	RP_itemTitleCfg         = "itemTitle_config"        //hash 存放itemTitle配置内容
	RP_noticCfg             = "notice_config"           //hash 存放notic配置内容
	RP_SamGongCfg           = "sg_table_config"         //hash 存放三公游戏场次配置内容
	RP_ZJHTableCfg          = "zjh_table_config"        //hash 存放炸金花游戏场次配置内容
	RP_Domino99TableCfg     = "domino99_table_config"   //hash 存储多米诺99游戏场次配置内容
	RP_BlackJackTableCfg    = "blackjack_table_config"  //hash 存储二十一点游戏场次配置内容
	RP_CapsaSusunTableCfg   = "capsasusun_table_config" //hash 存储十三张游戏场次配置内容
	RP_bankruptCfg          = "bankrupt_config"         //hash 存放破产配置内容
	RP_thirdPayCfg          = "thirdpay_config"         //hash 第三方支付配置内容
	RP_taskCfg              = "task_config"             //hash 存放任务配置内容
	RP_signInCfg            = "signin_rewards_info"     //存放签到配置内容
	RP_signInNewer          = "signin_newer"            //存放新人签到配置内容 (2021-6-3 3个月后可下线与此变量相关逻辑）
	RP_signInNewer2         = "signin_newer2"           //存放新人签到配置内容(广告版)
	RP_taskActiveRewardsCfg = "task_active_rewards"     //存放签到配置内容
	RP_itemInfoCfg          = "item_info"               //存放物品信息配置
	RP_heroInfoCfg          = "hero_info"               //存放角色信息配置
	RP_sceneInfoCfg         = "scene_info"              //存放场景信息配置
	RP_groupMailCfg         = "groupMail_info"          //存放场景信息配置
	RP_dressInfoCfg         = "dress_info"              //存放装扮信息配置
	RP_badgeInfoCfg         = "badge_info"              //存放徽章信息配置
	RP_mallGoodsCfg         = "mall_goods_"             //eg:mall_goods_{$currency_$OS} 不同货币，不同平台的商品配置信息
	RP_waresCfg             = "wares_"                  //eg:wares_{$Type} 不同的商品类型下的商品配置信息
	RP_broadcastCfg         = "broadcast_config"        //hash 存放广播的配置内容
	RP_AICfg                = "AI_config"               //hash AI配置 1-10级
	RP_AI_AOFCfg            = "AI_AOF"                  //AOF AI配置 1-10级
	RP_AI_SamgongCfg        = "AI_Samgong"              //Samgong AI配置 1-10级
	RP_AI_Domino99Cfg       = "AI_Domino99"             //Domino99 AI配置 1-10级
	RP_AI_TeenPattiCfg      = "AI_TeenPatti"            //TeenPatti AI配置 1-10级
	PR_hipoker_push_cfg     = "hipoker_push_cfg"        //hipoker推送配置
	PR_Paopai_table_cfg     = "paopai_table_cfg"        //跑牌桌子配置
	PR_Dummy_table_cfg      = "dummy_table_cfg"         //大米桌子配置
	PR_ddzbu_table_cfg      = "ddzbu_table_cfg"         //斗地主不洗牌桌子配置
	PR_ddzjiao_table_cfg    = "ddzjiao_table_cfg"       //斗地主叫地主桌子配置
	PR_ddzqiang_table_cfg   = "ddzqiang_table_cfg"      //斗地主抢地主桌子配置
	PR_ddzlai_table_cfg     = "ddzlai_table_cfg"        //斗地主癞子桌子配置
	PR_King_table_cfg       = "king_table_cfg"          //king牌桌子配置

	//包管理系统使用
	RP_PackageDiff_White          = "package_diff"    //APP升级包差量信息（白名单）
	RP_Package_Down_Switch        = "pkg_down_switch" //差量包下载模式开关
	RP_PackageDiff_All            = "package_diff_all"
	RP_Resource_app_latest_white  = "res_app_" //应用版本对应的最新资源包版本
	RP_Resource_app_latest_all    = "res_app_all"
	RP_Resource_lang_latest_white = "res_l_" // 子游戏/公共包下某个语言--> 最新语言版本号
	RP_Resource_lang_latest_all   = "res_l_all"
	RP_Resource_comm_latest_white = "res_c_" // 公共包下，某个子游戏特定版本 -->对应的公共包版本号
	RP_Resource_comm_latest_all   = "res_c_all"
	RP_Resource_diff              = "res_diff_"      //资源差分包信息
	RP_white_list                 = "res_white_list" //白名单

	//排行榜信息
	RP_Ranking       = "ranking_"       //排行榜前缀信息
	RP_RankMatchTime = "rankMatchTime_" //排位赛当前赛段信息

	//天梯功能相关
	RP_LadderLevelsCfg = "ladder_levels_cfg"  //天梯段位配置数据
	RP_LadderEventsCfg = "ladder_events_cfg"  //天梯特殊事件配置数据
	RP_LadderRankCfg   = "ladder_rank_cfg"    //天梯排行榜奖励配置数据
	RP_LadderPeriodCfg = "ladder_period_cfg"  //天梯周期配置
	RP_LadderSeasons   = "ladder_seasons_"    //天梯排名赛赛季数据
	RP_LadderSeasonsV2 = "ladder_seasons_v2_" //天梯排名赛赛季数据

	RP_ConfigVersion = "cfg_ver"         //配置数据版本号
	Rp_LoadingMsgCfg = "loading_msg_cfg" //加载提示语配置
	Rp_LoadingAdCfg  = "loading_ad_cfg"  //开屏广告

	//好友
	RP_FriendStatSet  = "friend_stat_set_"  //用户好友上报统计集合
	RP_FriendStatHash = "friend_stat_hash_" //用户好友上报统计HASH

	//配牌
	RP_TestCards           = "test_cards"            //配牌数据
	RP_SamgongTestCards    = "samgong_test_cards"    //三公配牌数据
	RP_PaopaiTestCards     = "paopai_test_cards"     //跑牌配牌数据
	RP_ZJHTestCards        = "zjh_test_cards"        //跑牌配牌数据
	RP_Domino99TestCards   = "domino99_test_cards"   //多米诺99配牌数据
	RP_BlackJackTestCards  = "blackjack_test_cards"  //二十一点配牌数据
	RP_CapsaSusunTestCards = "capsasusun_test_cards" //十三张配牌数据
	RP_DummyTestCards      = "dummy_test_cards"      //大米配牌数据
	RP_DdzbuTestCards      = "ddzbu_test_cards"      //斗地主不洗牌配牌数据
	RP_DdzjiaoTestCards    = "ddzjiao_test_cards"    //斗地主叫地主配牌数据
	RP_DdzqiangTestCards   = "ddzqiang_test_cards"   //斗地主抢地主配牌数据
	RP_DdzlaiTestCards     = "ddzlai_test_cards"     //斗地主癞子配牌数据
	RP_KingTestCards       = "king_test_cards"       //king牌配牌数据

	//经验级别
	RP_ExpLevels  = "exp_levels"  //经验级别配置
	RP_ExpConfig  = "exp_config"  //经验发放配置
	RP_ExpConfig2 = "exp_config2" //经验发放配置(按游戏区分）

	//VIP经验级别
	RP_VipLevels        = "vip_levels"         //vip等级配置
	RP_VipExp           = "vip_exp"            //vip经验发放
	RP_VipPrivilegeSort = "vip_privilege_sort" //vip特权排序

	//小猪储钱罐
	RP_PiggyBankCfg = "piggy_bank_cfg" //小猪储钱罐配置

	//刮刮卡配置
	RP_ScratchCardCfg = "scratch_card_cfg"

	//月卡配置
	RP_MonthCardCfg = "month_card_cfg"

	//回归礼包配置
	RP_ReturnGiftCfg = "return_gift_cfg"

	//打地鼠配置
	RP_HamsterCfg = "hamster_cfg"

	//排行榜徽章发放
	RP_RankBadgeCfg = "rank_badge_cfg"

	//系统公告
	RP_SysNoticesCfg = "sys_notices_cfg"

	//新注发放
	RP_RegisterRewards = "register_rewards"

	//客户端日志上传配置
	RP_ClientLogCfg = "client_log_cfg"

	//破产线配置
	RP_BankruptThreshold = "bankrupt_threshold"

	//帮助信息
	RP_HandBook = "handbook"

	//活动中心配置信息
	RP_Activity                = "activity"
	RP_Activity_V2             = "activity_cfg_v2"
	RP_User_Activity           = "user_activity_%d"          //用户活动
	RP_User_Activity_RedDot    = "user_activity_redot_%d_%d" //用户活动小红点
	RP_Discount_Gift_Countdown = "discount_gift_%d"          //用戶特惠礼包倒计时

	//自定义配置主key（与subKey共同组成最终的key)
	RP_CustomizeConfig = "com_cfg_"

	//AllocServer相关
	RP_AllocTableHash   = "alloc_hash_%d"    //tid
	RP_AllocTableZSet   = "alloc_zset_%d_%d" //smallBlind, seats
	RP_AllocTableSet    = "alloc_set_%d"     //gameServerID
	RP_AllocGameServers = "alloc_hash_srv"
	//SNGAlloc相关
	RP_matchAllocAvailable = "match_alloc_available" //sorted-set(serverId => freeTables)
	RP_matchAllocActive    = "match_alloc_active"    //hash(configId => serverId)
	//CustomAlloc相关
	RP_customAllocAvailable = "custom_alloc_available" //sorted-set(serverId => freeTables)
	RP_customPwdSeq         = "custom_alloc_pwd_seq"   //number
	RP_customPwdData        = "custom_alloc_pwd_data"  //hash(pwd => data)
	//GambleAlloc相关
	RP_gambleAllocAvailable = "gamble_alloc_available" //zset(playerCount => serverId)
	RP_gambleJackpot        = "gamble_jackpot"         //int
	RP_gambleRobots         = "gamble_robots"          //set
	RP_gambleTestCards      = "gamble_test_cards"      //string
	//PaoPaiAlloc相关
	RP_AllocWaitingList     = "alloc_waiting_list_%d" //difen 跑牌-正在配桌中的玩家队列。
	RP_AllocWaitingUser     = "alloc_waiting_user_%d" //uid 跑牌-存放玩家在哪个场次上配桌。
	RP_AllocWorkingServerId = "alloc_working_id"      //跑牌-存放正在工作中的allocServer的id

	//每日登录用户bitmap
	RP_LoginStat = "login_stat_%s"

	//用户每日玩牌数据
	RP_UserDailyPlay = "dailyplay_%d_%s" //uid,当前时间format(20060102)
	//每日免费发放货币总量
	RP_DailyFreeMoney = "daily_free_money_%s_%s_%d" //当前时间format(20060102)/区域/当前客户端版本1安卓2苹果
	//每日付费发放货币总量
	RP_DailyPayMoney = "daily_pay_money_%s_%s_%d" //当前时间format(20060102)/区域/当前客户端版本1安卓2苹果
	//	每日回收台费
	RP_DailyFee = "daily_fee_%s_%s_%d" //当前时间format(20060102)/区域/当前客户端版本1安卓2苹果
	// 每个场次货币回收总量	包括钻石、金币和筹码
	RP_DailyFieldTotalFee = "daily_field_total_fee_%s_%s_%d"   //当前时间format(20060102)/区域/当前客户端版本1安卓2苹果
	RP_DailyGrantByField  = "daily_grant_by_field_%d_%s_%s_%d" // 货币类型/当前时间format(20060102)/区域/当前客户端版本1安卓2苹果
	//
	RP_FieldAll     = "field_all"                // 发放功能列表
	RP_Field_People = "field_people_%d_%s_%s_%d" // 功能ID，当前时间format(20060102)/区域/当前客户端版本1安卓2苹果

	//	在线人数相关
	RP_SNG_Online      = "sng_online_%d"      //当前客户端版本1安卓2苹果	SNG在线人数
	RP_Ordinary_Online = "ordinary_online_%d" //当前客户端版本1安卓2苹果	普通场在线人数

	//	在玩人数相关
	RP_SNG_Play            = "sng_play_%d"         //当前客户端版本1安卓2苹果	SNG在玩集合(自由场)
	RP_Ordinary_Play       = "ordinary_play_%d"    //当前客户端版本1安卓2苹果	普通场在玩集合(自由场)
	RP_Ordinary_Table_Play = "ordinary_table_play" //	普通场次在玩人数(hash) 场次对应在玩人数

	//	SNG开赛次数
	RP_SNG_Matches = "sng_matches_%s" //时间format(20060102)
	//	SNG服务费总数量
	RP_SNG_Total_Number = "sng_total_number_%s_%s_%d" //时间format(20060102)/区域/当前客户端版本1安卓2苹果
	//	SNG参赛集合
	RP_SNG_Mathch_People = "sng_match_people_%s_%s_%d" //时间format(20060102)/区域/当前客户端版本1安卓2苹果
	//	SNG参赛人次
	RP_SNG_Mathch_Count = "sng_match_count_%s_%s_%d" //时间format(20060102)/区域/当前客户端版本1安卓2苹果

	//	玩家总玩牌数
	RP_Players_Play_Game_Total = "players_play_game_total_%s_%d" //时间format(20060102)/当前客户端版本1安卓2苹果

	//	牌局平均时长
	RP_Average_Game_Time = "averate_game_time_%s" //时间format(20060102)
	//	分场次牌局平均时长
	RP_Average_Game_Time_Field = "averate_game_time_field_%s" //时间format(20060102)
	//	底池分布
	RP_Bottom_Pool_Distribute = "bottom_pool_distribute_%s" //时间format(20060102)
	//	在线用户UID
	RP_Online_Users = "online_users_%d" // 当前客户端版本1安卓2苹果

	//	发送邮件相关
	RP_Send_Email = "send_email_%s" //	邮箱地址

	RP_Code_Phone = "code_phone_%s" //	手机验证码

	//	邀请相关
	RP_INVITE_ROBOT = "invite_robot" //	邀请机器人(hash类型)

	//	在线用户扩展信息
	RP_Online_User_Ext_Data   = "online_uinfo_%d"   // uid%10000 (把uid散列成1万个key)
	RP_Online_User_Ext_DataV3 = "online_uinfov3_%d" // uid%10000 (把uid散列成1万个key)

	//	当前用户的随机角色装扮(hash) 有效期2小时
	RP_Random_Hero_Dress = "random_hero_dress_%d" //	当前用户的随机角色装扮(hash) 有效期2小时

	//	每日付费人数
	RP_DailyPayPeople = "daily_pay_people_%s_%s_%d" // 时间format(20060102)/区域/当前客户端版本1安卓2苹果 每日付费人数
	//	每日付费总金额
	RP_DailyPayTotalMoney = "daily_pay_total_money_%s_%s_%d" // 时间format(20060102)/区域/当前客户端版本1安卓2苹果 每日付费总金额

	//限制类的KEY
	RP_LimitFriendRequest        = "fri_req_%d_to_%d" //发送者-接收者
	RP_LimitSlotMachinePlayCount = "slot_play_cnt_%d" //老虎机摇奖次数(最低金额)
	RP_LimitRobotTreeGrowFlag    = "robot_tree_grow"  //机器人每日摇钱树自然增长标志

	//广播列表
	RP_BroadcastHistory = "bc_history_list_%s" //地区编码
	//广播列表版本号
	RP_BroadcastHistoryVer = "bc_history_ver_%s" //地区编码

	//用户支付反馈状态
	RP_Pay_Feedback_Status = "pay_feedback_status_%d"
	//用户支付失败状态
	RP_Pay_Failed_Status = "pay_failed_status_%s_%d" // 时间format(20060102)/uid
	//FCM推送服务密钥配置数据
	RP_FCMAccessToken = "fcm_access_token"

	//桌子数据
	RP_TableStat = "sb%d_p%d" //游戏ID,小盲,座位
	RP_TableData = "t%d"      //游戏ID,桌子ID

	//子游戏配置
	RP_SubGame     = "sub_game"      //子游戏列表配置
	RP_SubGameSort = "sub_game_sort" //子游戏列表排序配置

	//场次配置中的模拟人数
	RP_GameTypeShowCnt = "game_type_cnt_%d" //场次列表中模拟人数

	//大R用户机器人彩蛋
	RP_RobotEgg        = "robot_egg"
	RP_RobotEggCache   = "robot_egg_cache"   //用于彩蛋计算时的缓存
	RP_RobotEggHistory = "robot_egg_history" //用于回流用户彩蛋计算
	RP_BigRChatCount   = "r_b_chat_count"    //大R用户聊天消息数量

	RP_RobotEmotion = "rb_emotion" //机器人互动道具配置

	//老虎机相关
	RP_SlotSchemeConfig   = "slot_scheme_cfg"       //老虎机图标配置方案
	RP_SlotMiscConfig     = "slot_misc_cfg"         //老虎机杂项配置
	RP_SlotSchemeID       = "slot_scheme_id"        //老虎机配置方案名
	RP_SlotPoolMoney      = "slot_pool_money"       //老虎机池子金额
	RP_SlotHourlyStatData = "slot_recycled_%d"      //老虎机每小时的统计数据
	RP_SlotAllUsers       = "slot_users_%d%d%d"     //老虎机每日全部人数统计
	RP_SlotAmountUsers    = "slot_amount_%d_%d%d%d" //老虎机每日各额度人数统计
	RP_SlotRewardUsers    = "slot_reward_%d_%d%d%d" //老虎机各选项的中奖人数统计

	//家园相关
	RP_HomelandPets          = "homeland_pets"         //宠物列表
	RP_HomelandPetFeatures   = "homeland_pet_features" //宠物技能列表
	RP_HomelandPetStages     = "homeland_pet_stages"   //宠物阶段列表
	RP_HomelandPetDropConfig = "homeland_pet_drop"     //宠物道具掉落配置
	RP_HomelandTreeLevels    = "homeland_tree_levels"  //摇钱树列表
	RP_HomelandMiscConfig    = "homeland_misc"         //家园杂项配置
	//家园业务数据
	RP_HomelandIInvite            = "homeland_I_invite_%s_%d"      //homeland_I_invite_$ymd_$uid 	按天记录我邀请了谁
	RP_HomelandBeInvited          = "homeland_be_invited_%s_%d"    //homeland_be_invited_$ymd_$uid	按天记录谁邀请了我
	RP_HomelandIHelp              = "homeland_I_help_%s_%d"        //homeland_I_help_$ymd_$uid	按天记录我助力过谁
	RP_HomelandFreeGift           = "homeland_free_gift_%s_%d"     //homeland_free_gift_$ymd_$uid	按天记录我领取过几次免费的食物、免费的肥料
	RP_HomelandPetFreeFeed        = "homeland_pet_freefeed_%s_%d"  //homeland_pet_freefeed_$ymd_$uid	按天记录宠物被好友喂食记录
	RP_HomelandPetBenefit         = "homeland_pet_benefit_%d"      //homeland_pet_benefit_$uid	玩家的宠物宝箱信息
	RP_HomelandTreeFreeFeed       = "homeland_tree_freefeed_%d_%d" //homeland_tree_freefeed_$segTime_$uid	按摇钱树周期记录摇钱树被好友免费施肥的记录
	RP_HomelandTreeBenefit        = "homeland_tree_benefit_%d_%d"  //homeland_tree_benefit_$segTime_$uid	按摇钱树周期记录玩家的附加收益
	RP_HomelandEvents             = "homeland_events_%d"           //homeland_events_$uid	记录玩家家园动态。
	RP_HomelandEventWatchTime     = "homeland_event_watch_time_%d" //homeland_event_watch_time_$uid	记录玩家最近一次查看动态信息的时间。
	RP_HomelandTreeBeStolenRecord = "homeland_tree_stolen_%d_%d"   //homeland_tree_stolen_$segTime_$uid	按周期记录摇钱树被谁偷过
	RP_HomelandTreeIStealRecord   = "homeland_tree_I_steal_%s_%d"  //homeland_tree_stolen_$ymd_$uid	按天记录我偷过那些人
	RP_HomelandTaskReward         = "homeland_task_%s_%d"          //homeland_task_$ymd_$uid	按天记录我领取过哪些任务的奖励
	//家园推送相关数据
	RP_HomelandUidToRipe   = "homeland_tree_info_%d" //homeland_tree_ripe_$uid	玩家更换宠物的时候、自己喂养宠物的时候，好友来喂养宠物的时候会触发修改这个数值
	RP_HomelandPushProject = "homeland_push_%d"      //homeland_push_$time

	RP_AgentUid = "agent_uids" //需要记录台费分成的uid集合
)

//redis key的前缀设置- redis prefix (R3库)
//缘由：统一设置redis的key前缀可以有效避免key冲突的情况。同时方便统一知晓系统redis使用情况
const (
	//LRU缓存类数据
	RP_userMoney = "m_" //vmoney缓存数据，存储key为 m_{$uid}
	RP_userInfo  = "u_" //userinfo缓存数据，存储key为 u_{$uid}

	RP_userVistors = "uv" //eg:uv_{$uid} 存放此用户的访客信息
)

//redis key的前缀设置- redis prefix (R4库)
//缘由：统一设置redis的key前缀可以有效避免key冲突的情况。同时方便统一知晓系统redis使用情况
const (
	//LRU缓存类数据
	RP_CFRTreeData = ""   //CFR Tree的缓存数据，存储key为 {$step}{$dbName}{$location}
	RP_CFRTreeRoot = "r_" //CFR Tree的缓存数据，存储key为 {$step}{$dbName}{$location}
)

//不同数据库表使用的连接配置
const (
	DB_vmoney        = DB_CONF_USERDB //用户金币表
	DB_user_info     = DB_CONF_USERDB
	DB_robot_info    = DB_CONF_USERDB //机器人信息表
	DB_user_ext      = DB_CONF_USERDB //扩展信息表
	DB_user_feedback = DB_CONF_USERDB //用户反馈
	DB_account       = DB_CONF_USERDB
	DB_logs          = DB_CONF_LOGS //	用户日志
	DB_daily_money   = DB_CONF_LOGS //	每日货币统计表

	DB_match_rank = DB_CONF_USERDB //比赛名次信息

	DB_config_value = DB_CONF_SUPER_ADMIN //管理后台配置信息落地存放表
	DB_admin_config = DB_CONF_USERDB      //管理后台配置项落地在mysql中供RPC服务加载使用

	DB_payment    = DB_CONF_USERDB //支付订单
	DB_month_card = DB_CONF_USERDB //月卡

	DB_user_game_data = DB_CONF_USERDB //用户牌局指标数据记录

	DB_invite_code    = DB_CONF_USERDB // 邀请码表
	DB_invited_record = DB_CONF_USERDB // 邀请记录表

	DB_game_log = DB_CONF_SUPER_ADMIN //牌局记录

	//全局牌局记录表
	DB_game_log_lobby      = "mysql.database.game_log_lobby"
	DB_game_log_private    = "mysql.database.game_log_private"
	DB_game_log_sng        = "mysql.database.game_log_sng"
	DB_game_log_aof        = "mysql.database.game_log_aof"
	DB_game_log_gamble     = "mysql.database.game_log_gamble"
	DB_game_log_samgong    = "mysql.database.game_log_samgong"
	DB_game_log_paopai     = "mysql.database.game_log_paopai"
	DB_game_log_zjh        = "mysql.database.game_log_zjh"
	DB_game_log_domino99   = "mysql.database.game_log_domino99"
	DB_game_log_blackjack  = "mysql.database.game_log_blackjack"
	DB_game_log_capsasusun = "mysql.database.game_log_capsasusun"
	DB_game_log_dummy      = "mysql.database.game_log_dummy"
	DB_game_log_ddzbu      = "mysql.database.game_log_ddzbu"
	DB_game_log_ddzjiao    = "mysql.database.game_log_ddzjiao"
	DB_game_log_ddzqiang   = "mysql.database.game_log_ddzqiang"
	DB_game_log_ddzlai     = "mysql.database.game_log_ddzlai"
	DB_game_log_king       = "mysql.database.game_log_king"

	DB_admin_user = DB_CONF_SUPER_ADMIN //存放管理后台用户信息

	DB_cfr_preflop         = DB_CONF_CFR_PRFLOP  //翻牌前AI 的CFR数据
	DB_cfr_flop            = DB_CONF_CFR_FLOP    //翻牌AI 的CFR数据
	DB_cfr_turn            = DB_CONF_CFR_TURN    //转牌AI 的CFR数据
	DB_cfr_river           = DB_CONF_CFR_RIVER   //河牌AI 的CFR数据
	DB_master_match_season = DB_CONF_SUPER_ADMIN //大师赛赛程
	DB_nick_map            = DB_CONF_USERDB      //用户昵称映射
	DB_sng_matches         = DB_CONF_USERDB      //SNG比赛记录
	DB_robot               = DB_CONF_USERDB      //机器人

	//多库多表记录
	DB_money_records     = DB_CONF_USER_RECORDS //用户金币流水记录表
	DB_user_item         = DB_CONF_USER_RECORDS //用户背包物品
	DB_user_badge        = DB_CONF_USER_RECORDS //用户徽章物品
	DB_user_hero         = DB_CONF_USER_RECORDS //用户角色
	DB_user_dress        = DB_CONF_USER_RECORDS //用户角色
	DB_exchange_record   = DB_CONF_USER_RECORDS //兑换记录
	DB_delivery_record   = DB_CONF_USER_RECORDS //发货记录
	DB_user_title        = DB_CONF_USER_RECORDS //用户互动道具主题信息
	DB_user_task         = DB_CONF_USER_RECORDS //用户任务信息
	DB_user_game_log     = DB_CONF_USER_RECORDS //用户牌局记录
	DB_user_mail         = DB_CONF_USER_RECORDS //用户邮箱
	DB_user_match        = DB_CONF_USER_RECORDS //用户比赛获奖信息
	DB_user_masterscores = DB_CONF_USER_RECORDS //用户大师分
	DB_user_data         = DB_CONF_USER_RECORDS //用户通用数据
	DB_user_activity     = DB_CONF_USER_RECORDS //用户活动信息
	DB_user_activity_v2  = DB_CONF_USER_RECORDS //用户活动信息v2
)

const BANKRUPT_THRESHOLD = 20000
const ITEM_SOLD_RETURN_CHIP = 1000 //用户出售物品之后获得的筹码数量

//money server中moneyType
const (
	MONEY_TYPE_DIAMOND  = 1
	MONEY_TYPE_CHIP     = 2
	MONEY_TYPE_GOLDCOIN = 3
)

//在线状态-OnlineStatus
const (
	//在比赛中的状态单独存储在online服中，比赛中的状态不管是否掉线都不能修改。
	//“比赛中”的玩家，可能属于“玩牌中”，也可能不属于玩牌中；
	//ONLINE_NONE		= 1	//表示不在线
	ONLINE_IN_HALL           = 2 //在大厅（可理解为“其他”）
	ONLINE_IN_GAME           = 3 //在游戏中
	ONLINE_IN_OFFLINE        = 4 //玩家掉线（当从掉线中恢复的时候，查看桌子信息存在则转到3状态，不存在就转到2状态）
	ONLINE_IN_GAME_HALL      = 5 //在普通场桌子列表处
	ONLINE_IN_MATCH_HALL     = 6 //在比赛场列表处
	ONLINE_IN_MATCH_SNG_HALL = 7 //在SNG比赛场大厅处
	ONLINE_IN_MATCH_MTT_HALL = 8 //在MTT比赛场大厅处
)

const (
	//全局状态：客户端拿到这些状态后需要做后续动作。
	G_STATUS_DEFAULT           = 0 //默认状态
	G_STATUS_MATCH_MTT_SIGN_IN = 1 //比赛MTT的已签到的状态
	G_STATUS_MATCH_SNG_SIGN_UP = 2 //比赛SNG已报名
	G_STATUS_MATCHING          = 3 //玩家在比赛中(备注：比赛玩牌中的玩家，状态定义为比赛中)
	G_STATUS_IN_TABLE          = 4 //玩家在普通桌子上玩牌中
)

//定义比赛相关状态
const (
	MATCH_STATUS_NO_SIGN_UP  = 0 //未报名任何比赛
	MATCH_STATUS_SIGN_UP     = 1 //已报名某场比赛
	MATCH_STATUS_SIGN_IN     = 2 //已报名且签到（MTT）
	MATCH_STATUS_MATCHING    = 3 //比赛中
	MATCH_STATUS_SNG_SIGN_IN = 4 //已报名（SNG）
)

//定义金流及物品改变时的“动作类型”
const (
	ACT_TYPE_BEGIN                       = 0
	ACT_TYPE_REGIST                      = 1   //1 注册
	ACT_TYPE_BUY                         = 2   //2 充值
	ACT_TYPE_EXCHANGE_LOSE               = 3   //3 兑换失去
	ACT_TYPE_EXCHANGE_GET                = 4   //4 兑换获得
	ACT_TYPE_ROLLBACK                    = 5   //5 回滚（退钱）
	ACT_TYPE_BUYGIVE                     = 6   //6 购买加赠
	ACT_TYPE_SIGNIN                      = 7   //7 签到礼包
	ACT_TYPE_BUY_IN_COIN                 = 8   //8 买入筹码(已作废)
	ACT_TYPE_REEXCHANGE                  = 9   //9 筹码换回金币(已作废)
	ACT_TYPE_BANKRUPT_SUBSIDY            = 10  //10 破产补助
	ACT_TYPE_ADMIN_MANEGEMENT            = 11  //11 管理后台操作
	ACT_TYPE_GIFTS_EXCHANGE              = 12  //12 礼包兑换
	ACT_TYPE_MATCH_SIGN                  = 13  //13 报名比赛，扣除报名金币
	ACT_TYPE_MATCH_ROLLBACK              = 14  //14 取消报名比赛，返还之前扣除的金币
	ACT_TYPE_TASK_PRIZE                  = 15  //15 任务获奖
	ACT_TYPE_MATCH_AWARD                 = 16  //16 比赛发奖
	ACT_TYPE_FIRST_BUY_GIVE              = 17  //17 首充加赠
	ACT_TYPE_ROBOT_ADD                   = 18  //18 机器人添加金币
	ACT_TYPE_SIGN_REPAIR                 = 19  //19 补签消耗金币
	ACT_TYPE_MAIL_ITEMS                  = 20  //20 邮件发放
	ACT_TYPE_ACTIVITY_ITEMS              = 21  //21 活动发放
	ACT_TYPE_LADDER_SEASON_REWARD        = 22  //22 天梯段位奖励
	ACT_TYPE_LADDER_RANK_REWARD          = 23  //23 天梯排行榜奖励
	ACT_TYPE_ITEM_SOLD                   = 24  //24 道具出售
	ACT_TYPE_RANK                        = 25  //25 排行榜发放
	ACT_TYPE_INVITE                      = 26  //26 邀请好友本人领取奖励
	ACT_TYPE_INVITE_BIND                 = 27  // 27 邀请好友邀请人领取奖励
	ACT_TYPE_CUSTOM_GAME_LOST            = 28  // 28 好友房失去(非打牌)
	ACT_TYPE_CUSTOM_GAME_GOT             = 29  // 29 好友房获得(非打牌)
	ACT_TYPE_CUSTOM_CREATE               = 30  // 30 创建好友房消耗
	ACT_TYPE_CUSTOM_ROLLBACK             = 31  // 31 创建好友房回滚
	ACT_TYPE_BROADCAST_SEND              = 32  // 32 发送广播消耗
	ACT_TYPE_BROADCAST_ROLLBACK          = 33  // 33 发送广播回滚
	ACT_TYPE_NEWER_SIGN_IN               = 34  // 34 新人签到
	ACT_TYPE_TASK_PRIZE_NEWER            = 35  // 35 新人任务奖励
	ACT_TYPE_NEWER_TASK_LEVLE_PRIZE      = 36  // 36 新人阶段任务奖励
	ACT_TYPE_NEWER_WARE                  = 37  // 37 新人特惠商品
	ACT_TYPE_NEWER_WARE_GIFT             = 38  // 38 新人特惠商品礼包
	ACT_TYPE_TASK_PRIZE_DAILY            = 39  // 每日任务奖励
	ACT_TYPE_TASK_PRIZE_LIFE             = 40  // 生涯任务奖励
	ACT_TYPE_TASK_PRIZE_GROW             = 41  // 等级任务奖励
	ACT_TYPE_TASK_PRIZE_ACTIVE           = 42  //任务活跃值奖励
	ACT_TYPE_EXCHANGE_LOSE_MALL          = 43  //商城兑换
	ACT_TYPE_QUICK_BUY_SIGN_IN           = 44  //快速购买金币（签到）
	ACT_TYPE_QUICK_BUY_AVATAR            = 45  //快速购买金币（装扮）
	ACT_TYPE_QUICK_BUY_PRIVATE_TABLE     = 46  //快速购买金币（私人桌）
	ACT_TYPE_QUICK_BUY_BANKRUPT          = 47  //快速购买筹码（破产）
	ACT_TYPE_QUICK_BUY_LOBBY_TABLE       = 48  //快速购买筹码（进场）
	ACT_TYPE_QUICK_BUY_SNG               = 49  //快速购买筹码（SNG）
	ACT_TYPE_QUICK_BUY_MTT               = 50  //快速购买筹码（MTT）
	ACT_TYPE_PARTNER_TRANSFER            = 51  //合作伙伴工具
	ACT_TYPE_TASK_PRIZE_ACHIEVE          = 52  //成就奖励
	ACT_TYPE_SET_NICKNAME                = 53  //修改用户昵称
	ACT_TYPE_ACTIVITY_CENTER             = 54  //活动中心
	ACT_TYPE_DISCOUNT_NEWER              = 55  //新人特惠礼包 (扣钻石)
	ACT_TYPE_DISCOUNT_GIFT               = 56  //首充特惠礼包 (扣钻石)
	ACT_TYPE_FREE_CHIPS                  = 57  //广告免费领取
	ACT_TYPE_LOGIN_PRIZE_CHIPS           = 58  //登录广告奖励
	ACT_TYPE_VIP_BUY_CHIP_BONUS          = 59  //vip购买筹码加赠
	ACT_TYPE_VIP_SUBSIDY                 = 60  //vip破产补助加赠
	ACT_TYPE_SUPER_VIP_BADGE             = 61  //vip10获得至尊VIP（徽章）
	ACT_TYPE_BANKRUPT_GOODS              = 62  //破产商品
	ACT_TYPE_PIGGY_BANK_EXCHANGE         = 63  //小猪储钱罐砸猪扣钻石
	ACT_TYPE_PIGGY_BANK_BADGE            = 64  //小猪储钱罐徽章
	ACT_TYPE_ROBOT_REDUCE                = 65  //机器人削减资产
	ACT_TYPE_SCRATCH_CARD                = 66  //刮刮卡扣钻石
	ACT_TYPE_SCRATCH_CARD_RECEIVE        = 68  //刮刮卡领取筹码
	ACT_TYPE_BUY_CAPSA_SUSUN_TOOL        = 69  //购买十三张智能理牌卡
	ACT_TYPE_UNIQUE_DISCOUNT             = 70  //购买破产专属特惠
	ACT_TYPE_TEXAS_JACKPOT_WON           = 71  //Jackpot中奖
	ACT_TYPE_TEXAS_JACKPOT_VIP_EXTRA     = 72  //Jackpot的VIP加成
	ACT_TYPE_MONTH_CARD                  = 73  //购买月卡
	ACT_TYPE_RETURN_GIFT_LOGIN           = 74  //回归登录奖励
	ACT_TYPE_RETURN_GIFT                 = 75  //回归礼包奖励
	ACT_TYPE_DRESS_GIVE                  = 76  //装扮赠送(主动赠送)
	ACT_TYPE_DRESS_REQUEST_GIVE          = 77  //装扮赠送(索要)
	ACT_TYPE_BUY_ALL_DRESS               = 78  //购买全套装扮
	ACT_TYPE_SLOT_MACHINE_CONSUME        = 79  //老虎机消耗
	ACT_TYPE_SLOT_MACHINE_REWARD         = 80  //老虎机奖励
	ACT_TYPE_HAMSTER_BUY                 = 81  //打地鼠购买关卡
	ACT_TYPE_HAMSTER_REWARD              = 82  //打地鼠奖励
	ACT_TYPE_BONUS_LOTTERY               = 84  //商城购买钻石赠送彩票
	ACT_TYPE_LADDER_EXT_REWARD           = 83  //天梯段位额外奖励
	ACT_TYPE_LADDER_SALARY               = 85  //天梯段位奖励
	ACT_TYPE_ANNIVERSARY_GIFT            = 86  //周年庆礼包
	ACT_TYPE_HOMELAND_TREE_STEAL_SUCCESS = 87  //【家园】摇钱树偷成功获得筹码
	ACT_TYPE_HOMELAND_TREE_STEAL_FAILED  = 88  //【家园】摇钱树偷失败失去筹码
	ACT_TYPE_HOMELAND_TREE_REWARD        = 89  //【家园】领取摇钱树的奖励
	ACT_TYPE_HOMELAND_TREE_UPGRADE       = 90  //【家园】摇钱树升级消耗金币
	ACT_TYPE_HOMELAND_FREE_FOOD          = 91  //【家园】领取免费的食物、肥料
	ACT_TYPE_HOMELAND_PET_REWARD         = 92  //【家园】领取宠物的奖励
	ACT_TYPE_HOMELAND_PET_UPGRADE        = 93  //【家园】宠物升级消耗金币
	ACT_TYPE_HOMELAND_BUY_PET            = 94  //【家园】购买宠物
	ACT_TYPE_ADMIN_LOTTERY               = 95  //后台赠送彩票
	ACT_TYPE_CHRISTMAS_GIFT              = 96  //圣诞季活动礼盒奖励
	ACT_TYPE_CHRISTMAS_RANK              = 97  //圣诞季活动排名奖励
	ACT_TYPE_PARNTER_MONTH_BILL          = 98  //代理每月月结分成筹码发放
	ACT_TYPE_PARTNER_WITHDRAWAL          = 99  //合伙人提现
	ACT_TYPE_INVITE_MANUAL_BIND          = 100 //邀请手动绑定

	//往前面添加
	ACT_TYPE_GAME_WIN  = 1000 //玩牌赢得
	ACT_TYPE_GAME_LOSE = 1001 //玩牌输去
	ACT_TYPE_END       = 1002 //不要往后面添加
)

//物品类型
const (
	ITEM_TYPE_BEGIN      = 0
	ITEM_TYPE_DIAMOND    = 1  //钻石
	ITEM_TYPE_GOLDCOIN2  = 2  //金币
	ITEM_TYPE_CHIP       = 3  //筹码
	ITEM_TYPE_COMMON     = 4  //一般道具
	ITEM_TYPE_GIFTS      = 5  //道具礼包
	ITEM_TYPE_CONTEST    = 6  // 赛事道具
	ITEM_TYPE_ITEMTITLE  = 7  // 互动道具主题
	ITEM_TYPE_HERO       = 8  // 角色
	ITEM_TYPE_DRESS      = 9  // 装扮物品
	ITEM_TYPE_SCENE      = 10 // 场景
	ITEM_TYPE_BADGE      = 11 //徽章
	ITEM_TYPE_ACTIVITY   = 12 //活动类物品	支付场景上报的时候，刮刮卡、月卡类会定义的该值。
	ITEM_TYPE_HOMELAND   = 13 //家园相关物品：食物、肥料、100经验卡、N经验卡
	ITEM_TYPE_LOTTERY    = 14 //彩票
	ITEM_TYPE_MONTH_CARD = 15 //周卡月卡
	ITEM_TYPE_END        = 20
)

//商品类型
const (
	GOODS_TYPE_BEGIN          = 0 //开始，用于判断合法性
	GOODS_TYPE_DIAMON_ANDRIOD = 1
	GOODS_TYPE_DIAMOND_IOS    = 2
	GOODS_TYPE_GOLDCOIN       = 3
	GOODS_TYPE_ITEMGOODS      = 4  //道具商品
	GOODS_TYPE_ITEMTITLE      = 5  //交互道具主题
	GOODS_TYPE_HERO           = 6  //角色
	GOODS_TYPE_DRESS          = 7  //装扮物品
	GOODS_TYPE_CHIP           = 8  //筹码
	GOODS_TYPE_SCENE          = 9  //场景
	GOODS_TYPE_BANKRUPT_CHIP  = 10 //破产筹码
	GOODS_TYPE_BANKRUPT_CHIP2 = 11 //破产筹码
	GOODS_TYPE_SPECIAL_CHIP   = 12 //特殊筹码（不进商城）
	GOODS_TYPE_NEWER_GIFT     = 13 //新人特惠礼包
	GOODS_TYPE_DISCOUNT_GIFT  = 14 //特惠礼包
	GOODS_TYPE_BANKRUPT_GOODS = 15 //破产商品
	GOODS_TYPE_UNIQUE_GOODS   = 16 //专属特惠
	GOODS_TYPE_HOME_GOODS     = 17 //家园商品
	GOODS_TYPE_END            = 18 //结束，用于判断合法性
)

//发货类型
const (
	DELIVERY_TYPE_DIAMOND    = 1
	DELIVERY_TYPE_GOLDCOIN   = 2
	DELIVERY_TYPE_CHIP       = 3
	DELIVERY_TYPE_ITEM       = 4  //道具
	DELIVERY_TYPE_ITEMTITLE  = 5  //交互道具主题
	DELIVERY_TYPE_HERO       = 6  //角色
	DELIVERY_TYPE_DRESS      = 7  //装扮
	DELIVERY_TYPE_SCENE      = 8  //场景
	DELIVERY_TYPE_BADGE      = 9  //徽章
	DELIVERY_TYPE_LOTTERY    = 10 //彩票
	DELIVERY_TYPE_MONTH_CARD = 11 //周卡月卡
	DELIVERY_TYPE_HOMELAND   = 13 //家园相关道具卡
)

//NSQ topic类型
const (
	NSQ_TOPIC_DELIVERY_RECORD   = "DeliveryRecord"
	NSQ_TOPIC_EXPIRE_DEALER     = "ExpireRecordsDealer"
	NSQ_TOPIC_MONEY_RECORD      = "MoneyRecord"
	NSQ_TOPIC_ROBOTCFR_RECORD   = "robotCFRRecord"
	NSQ_TOPIC_TASK_MSG          = "TaskMsg"
	NSQ_TOPIC_RANKING_MSG       = "RankingMsg"
	NSQ_TOPIC_MATCH_RANK_MSG    = "MatchRankMsg" //一场比赛的最终排名信息
	NSQ_TOPIC_BROADCAST_MSG     = "BroadcastMsg"
	NSQ_TOPIC_ONE_GAME          = "oneGame"         //普通场一局游戏结束上报
	NSQ_TOPIC_LOGIN_SUCC        = "LoginSucc"       //用户登录成功
	NSQ_TOPIC_ONLINE_CHANGED    = "OnlineChanged"   //用户在线状态变化(即将废弃)
	NSQ_TOPIC_USER_INFO_UPDATED = "UserInfoUpdated" //用户个人信息变化s
	//NSQ_TOPIC_LOG_ERROR		= "LogError"			//错误日志管理
	NSQ_TOPIC_PUSH_SUBMIT              = "PushMsg"                //推送信息上报
	NSQ_TOPIC_STAT_SUBMIT              = "StatMsg"                //数据上报
	NSQ_TOPIC_STAT_EVENT_SUBMIT        = "StatEventMsg"           //统计相关事件上报
	NSQ_TOPIC_ICON_DOWNLOAD            = "IconDownload"           //从第三方平台下载头像到本地
	NSQ_TOPIC_LOGIN_LOG                = "LoginLog"               //	登录日志
	NSQ_TOPIC_PAY_DONE                 = "PayDone"                //	用户支付完成
	NSQ_TOPIC_LOGIN_LOG_UPDATE_USER    = "LoginLogUpdateUser"     //	用户支付完成
	NSQ_TOPIC_PAY_GOOD                 = "PayGood"                //	购买商品后写进日志
	NSQ_TOPIC_ONLINE_EVENT             = "OnlineEvent"            //在线状态变化事件(V2)
	NSQ_TOPIC_SEND_MAIL                = "SendMail"               //发送邮件
	NSQ_TOPIC_INVITE_BIND              = "InviteBind"             //邀请绑定
	NSQ_TOPIC_BANKRUPT_RECORD          = "BankruptRecord"         //	破产记录
	NSQ_TOPIC_RECEIVER_BANKRUPT_RECORD = "ReceiverBankruptRecord" //	领取破产奖励记录
	NSQ_TOPIC_ACTIVITY_DATA            = "ActivityData"           //	活动数据
	NSQ_TOPIC_ACTIVITY_CENTER          = "ActivityCenter"         //活动中心使用到的数据
	NSQ_TOPIC_GAMBLE_100P_LOG          = "Gamble100pLog"          //百人场牌局日志
	NSQ_TOPIC_CUSTOM_GAME_USER_STAT    = "CustomGameUserStat"     //私人房用户玩牌统计
	NSQ_TOPIC_SAMGONG_LOG              = "SamgongLog"             //三公牌局日志
	NSQ_TOPIC_GAME_PAOPAI_LOG          = "GamePaoPaiLog"          //跑牌牌局日志
	NSQ_TOPIC_GAME_ZJH_LOG             = "ZJHLog"                 //炸金花牌局日志
	NSQ_TOPIC_GAME_DOMINO99_LOG        = "Domino99Log"            //多米诺99牌局日志
	NSQ_TOPIC_GAME_BLACKJACK_LOG       = "BlackJackLog"           //二十一点牌局日志
	NSQ_TOPIC_GAME_CAPSA_SUSUN_LOG     = "CapsaSusunLog"          //十三张牌局日志
	NSQ_TOPIC_GAME_DUMMY_LOG           = "DummyLog"               //大米牌局日志
	NSQ_TOPIC_GAME_DDZBU_LOG           = "DdzbuLog"               //斗地主不洗牌牌局日志
	NSQ_TOPIC_GAME_DDZJIAO_LOG         = "DdzjiaoLog"             //斗地主叫地主牌局日志
	NSQ_TOPIC_GAME_DDZQIANG_LOG        = "DdzqiangLog"            //斗地主抢地主牌局日志
	NSQ_TOPIC_GAME_DDZLAI_LOG          = "DdzlaiLog"              //斗地主癞子牌局日志
	NSQ_TOPIC_GAME_KING_LOG            = "KingLog"                //king牌牌局日志
	NSQ_TOPIC_TABLE_USER_CHANGE        = "TableUserChange"        //桌子用户数据变化
	NSQ_TOPIC_SNG_TABLE_EVENT          = "SNGTableEvent"          //SNG桌子事件
	NSQ_TOPIC_GAME_SERVER_EVENT        = "GameServerEvent"        //游戏服务事件
	NSQ_TOPIC_HOMELAND_ACTION          = "HomelandAction"         //家园关联行为
)

const (
	OS_SYSTEM_TYPE_ANDRIOD = 1 //Andriod系统
	OS_SYSTEM_TYPE_IOS     = 2 //IOS系统
)

//商品的支付类型
const PAY_MODE_CASH = 1     //现金
const PAY_MODE_DIAMOND = 2  //钻石
const PAY_MODE_GOLDCOIN = 3 //金币

//rpc协议中,透传map中的自定义字段名称.
const (
	//rpc透传时,map中的field名称.(放到const中主要是为了避免滥用冲突)
	ReqMetaDataKey_FIELD_srcName   = "srcName"   //传递rpc请求源的名称
	ReqMetaDataKey_FIELD_srcSvid   = "srcSvid"   //传递rpc请求源的svrid
	ReqMetaDataKey_FIELD_clientSeq = "clientSeq" //传递客户端请求序列号,rpc服务在必要时可以快速响应该seq的请求.(比如 game服快速响应操作)
	ReqMetaDataKey_FIELD_clientUid = "clientUid" //rpc调用时,传递客户端的这个socket连接对应的uid,便于业务做安全性检测
	ReqMetaDataKey_FIELD_dstSvid   = "dstSvid"   //rpcClient中记录下目标svrid,以便于选择合适的服务端
	ReqMetaDataKey_FIELD_remoteIp  = "remoteIp"  //请求方的ip地址；access转入登录的时候会用到。
)

//Game服务中，
const (
	//玩家的几种操作类型
	GAME_CHIPIN_FOLD              int32 = 1  //弃牌
	GAME_CHIPIN_CHECK             int32 = 2  //看牌
	GAME_CHIPIN_CALL              int32 = 3  //下注
	GAME_CHIPIN_SMALL_BLIND       int32 = 4  //小盲（重连有需要）
	GAME_CHIPIN_BIG_BLIND         int32 = 5  //大盲（重连有需要）
	GAME_CHIPIN_STANDFOLD         int32 = 6  //站起弃牌
	GAME_CHIPIN_TABLE_FEE         int32 = 7  //台费（重连有需要）
	GAME_CHIPIN_ANTE_FEE          int32 = 8  //前注费用（重连有需要）
	GAME_CHIPIN_CALL_CALL         int32 = 9  //下注-跟注
	GAME_CHIPIN_CALL_RAISE        int32 = 10 //下注-加注(每一轮的开始时，玩家的下注动作相比于0也是用加注)
	GAME_CHIPIN_CALL_ALLIN        int32 = 11 //下注-ALLIN
	GAME_CHIPIN_CALL_ALLIN_FOLLOW int32 = 12 //下注-ALLIN(跟注型ALLIN==被动型下注) 牌局记录自己算
	GAME_CHIPIN_CALL_ALLIN_ADD    int32 = 13 //下注-ALLIN（加注型ALLIN==主动型下注）牌局记录自己算
)

//比赛中的涨盲规则
const (
	BLIND_UP_TYPE_NORMAL int16 = 1 //普通的涨盲方式，以一定倍数增长
	BLIND_UP_TYPE_STEP   int16 = 2 //按精细的方式涨盲。每次固定盲注数值
)

const (
	OPERATOR_SYS  byte = 1 //GAME服务中，表示操作者为系统
	OPERATOR_USER byte = 2 //GAME服务中，表示操作者为用户
)

const (
	HERO_DEFAULT_MAX_ID = 100 //默读角色的最大ID
)

const (
	ORDER_STATUS_UNPAY      = 0 //订单状态：未支付
	ORDER_STATUS_DELIVERING = 1 //订单状态：正在发货
	ORDER_STATUS_DONE       = 2 //订单状态：已发货
	ORDER_STATUS_CANCELED   = 3 //订单状态：用户取消
	ORDER_STATUS_FAILED     = 9 //订单状态：发货失败(无效订单)
)

const (
	ONLINE_FIELD_STATUS          int32 = 3  //在线状态
	ONLINE_FIELD_IP              int32 = 4  //用户IP
	ONLINE_FIELD_ACCESS_SRV_ID   int32 = 5  //AccessServerID
	ONLINE_FIELD_ACCESS_SRV_ADDR int32 = 6  //AccessServer地址
	ONLINE_FIELD_GAME_ID         int32 = 7  //游戏ID
	ONLINE_FIELD_GAME_SRV_ID     int32 = 8  //游戏服务ID
	ONLINE_FIELD_GAME_SRV_ADDR   int32 = 9  //游戏服务地址
	ONLINE_FIELD_TID             int32 = 10 //桌子ID
	ONLINE_FIELD_TOKEN           int32 = 11 //玩家的token信息
	ONLINE_FIELD_COUNTRY         int32 = 12 //国家
	ONLINE_FIELD_LANGUAGE        int32 = 13 //语言
	ONLINE_FIELD_PLATFORM        int32 = 14 //客户端类型
	ONLINE_FIELD_VERSION         int32 = 15 //客户端版本
	ONLINE_FIELD_CHANNEL         int32 = 16 //用户来源渠道
)

const (
	AVATAR_BEGIN int32 = iota
	AVATAR_CLOTH       //1
	AVATAR_HAIR        //头发
	AVATAR_GLASSES
	AVATAR_BEARD       //胡子
	AVATAR_HAND_LEFT_1 //5
	AVATAR_HAND_LEFT_2
	AVATAR_HAND_LEFT_3
	AVATAR_HAND_LEFT_4
	AVATAR_HAND_LEFT_5
	AVATAR_HAND_RIGHT_1 // 10
	AVATAR_HAND_RIGHT_2
	AVATAR_HAND_RIGHT_3
	AVATAR_HAND_RIGHT_4
	AVATAR_HAND_RIGHT_5
	AVATAR_LUCK_ITEM   //15 吉祥物
	AVATAR_BADGE       //徽章
	AVATAR_SCENE       //场景
	AVATAR_WRIST_LEFT  //左手腕
	AVATAR_WRIST_RIGHT //右手腕
	AVATAR_END
)

const (
	ONLINE_EVENT_LOGIN     = 1 //登录
	ONLINE_EVENT_LOGOUT    = 2 //注销登录
	ONLINE_EVENT_ENTER     = 3 //进入房间
	ONLINE_EVENT_EXIT      = 4 //离开房间
	ONLINE_EVENT_SIT       = 5 //坐下
	ONLINE_EVENT_STAND     = 6 //站起
	ONLINE_EVENT_PLAY      = 7 //开始玩牌
	ONLINE_EVENT_OFFLINE   = 8 //断线
	ONLINE_EVENT_RECONNECT = 9 //重连
)

const (
	ONLINE_STATUS_LOGOUT  = -1 //登出
	ONLINE_STATUS_OFFLINE = 0  //离线
	ONLINE_STATUS_ONLINE  = 1  //在线
)

const (
	PLAY_STATUS_NONE      = 0 //无
	PLAY_STATUS_PREPARING = 1 //准备中(可能未进入房间)
	PLAY_STATUS_WATCHING  = 2 //旁观中
	PLAY_STATUS_PLAYING   = 3 //玩牌中
)

const (
	UPSTREAM_TYPE_REGULAR_GAME = 1  //普通场游戏
	UPSTREAM_TYPE_SNG_MATCH    = 2  //SNG比赛
	UPSTREAM_TYPE_MTT_MATCH    = 3  //MTT比赛
	UPSTREAM_TYPE_CUSTOM_GAME  = 4  //私人房
	UPSTREAM_TYPE_AOF_GAME     = 5  //AOF场
	UPSTREAM_TYPE_GAMBLE_100P  = 6  //押大小百人场
	UPSTREAM_TYPE_PAOPAI       = 7  //跑牌
	UPSTREAM_TYPE_SAMGONG      = 8  //三公
	UPSTREAM_TYPE_ZHAJINHUA    = 9  //炸金花
	UPSTREAM_TYPE_DOMINO99     = 10 //多米诺99
	UPSTREAM_TYPE_BLACK_JACK   = 11 //21点
	UPSTREAM_TYPE_LAMI         = 12 //印度拉米
	UPSTREAM_TYPE_CAPSA_SUSUN  = 13 //十三张
	UPSTREAM_TYPE_KING         = 14 //King牌
	UPSTREAM_TYPE_DUMMY        = 15 //泰国大米
	UPSTREAM_TYPE_DDZBU        = 16 //斗地主不洗牌
	UPSTREAM_TYPE_DDZJIAO      = 17 //斗地主叫地主
	UPSTREAM_TYPE_DDZQIANG     = 18 //斗地主抢地主
	UPSTREAM_TYPE_DDZLAI       = 19 //斗地主癞子
)

const (
	SPECIAL_ONLINE_SERVER_TABLE_ID = 65536 //特殊的服务ID，跑牌类的配桌玩家，配桌时需要将状态上报到online服，此时需要设置一个SERVERID，设置该值表示玩家在配桌中！
)

const (
	UPSTREAM_SEAT_NONE = -1 //空座位号
)

const (
	ONLINE_UID_SET_NUM = 100 //在线UID集合数量
)

const (
	SEND_EMAIL_BIND_ACCOUNT    = 1 //绑定账号发送邮件
	SEND_EMAIL_CHANGE_PASSWORD = 2 //找回密码发送邮件
)

//	设备类型
const (
	DEVICE_TYPE_ANDROID = 1 //	设备类型安卓
	DEVICE_TYPE_IOS     = 2 //	设备类型IOS
	DEVICE_TYPE_STEAM   = 3 //	设备类型STEAM
)

//	网络类型
const (
	NETWORK_TYPE_NotReachable = 0
	NETWORK_TYPE_WWAN         = 1 //	运营商移动数据
	NETWORK_TYPE_WIFI         = 2 //	WIFI
)

//	被邀请用户奖励领取状态状态
const (
	INVITE_STATUS_NOT_BIND     = 1 // 未绑定
	INVITE_STATUS_NOT_RECEIVED = 2 // 未领取
	INVITE_STATUS_TIMEOUT      = 3 // 已过期
	INVITE_STATUS_RECEIVED     = 4 // 以领取
)

// 被邀请人奖励的货币类型
const (
	INVITE_REWARD_TYPE_DIAMOND  = 1 // 钻石
	INVITE_REWARD_TYPE_CHIP     = 2 // 筹码
	INVITE_REWARD_TYPE_GOLDCOIN = 3 // 金币
)

//	邀请状态
const (
	INVITE_UN_RECEIVED = 0 // 未领取
	INVITE_RECEIVED    = 1 // 已领取
)

//用户通用数据字段定义
const (
	USERDATA_INVITE_REWARD_TIME         = 10001
	USERDATA_INVITE_ROBOT_GENERATE_TIME = 10002 //	生成邀请机器人的时间
	USERDATA_INVITE_INCRBY_TIME         = 10003 //	随机增加邀请人的时间
	USERDATA_PUSH_TOKEN                 = 10004 //推送设备标识
	USERDATA_PUSH_SERVICE               = 10005 //推送服务类型
	USERDATA_PAY_AMOUNT                 = 10006 //用户支付额度(美金,单位是千分之一元)
	USERDATA_SOCIAL_SETTING             = 10007 //社交相关设置
	USERDATA_BLACKLIST_VER              = 10008 //黑名单列表的数据版本号
	USERDATA_FRIEND_LIST_VER            = 10009 //好友列表的数据版本号
	USERDATA_FRIEND_REQUEST_LIST_VER    = 10010 //好友请求消息列表的数据版本号
	USERDATA_LOGIN_PROMOTE_SOURCE       = 10011 //	登录来源设置
	USERDATA_RECENT_PLAYERS_VER         = 10012 //最近玩牌玩家列表的数据版本号
	USERDATA_FAKE_COUNTRY               = 10013 //	登录玩家手机设置的国家地区
	USERDATA_USER_APP_VERSION           = 10014 //应用版本号(提审的版本号)
	USERDATA_FIRST_SET_NICKNAME         = 10015 //用户是否首次修改昵称
	USERDATA_USER_SUB_GAME_SELECT       = 10016 //用户选择放置在首页的子游戏及顺序
	USERDATA_LOTTERY_MAIL               = 10017 //用户中奖邮箱（幸运大转盘）
	USERDATA_VIP_EXP                    = 10018 //VIP成长点
	USERDATA_FREE_CHIPS                 = 10019 //用户看广告免费筹码每天领取数量
	USERDATA_TASK_AD                    = 10020 //用户登录有奖及活跃看广告相关的信息
	USERDATA_VIP_SET_NICKNAME_TIME      = 10021 //vip上次免费修改名字时间
	USERDATA_PAY_MAX_AMOUNT             = 10022 //用户支付最大单笔支付额度
	USERDATA_CAPSA_SUSUN_SMART_TOOL     = 10023 //用户十三张智能理牌工具购买时间
	USERDATA_SCRATCH_CARD_FREE          = 10024 //刮刮卡付费用户赠送刮刮卡
	USERDATA_UNIQUE_DISCOUNT            = 10025 //专属特惠发货时间
	USERDATA_SLOT_MACHINE_LOG           = 10026 //用户最近一次老虎机中奖记录
	USERDATA_LAST_LOGIN_TIME            = 10027 //登录用户上次登录时间
	USERDATA_SCRATCH_CARD_LOTTERY       = 10028 //刮刮卡抽奖次数
	USERDATA_DISCOUNT_GIFT              = 10029 //特惠礼包发货时间
	USERDATA_LADDER_FIRST_SEASON        = 10030 //用户第一参加钻石赛季（新赛制）
	USERDATA_DAILY_FIRST_LOGIN_TIME     = 10031 //玩家某天首次登录时间（玩家最近登录的那天的第一次登录时间）
	USERDATA_LOGIN_DAYS                 = 10032 //玩家总的登录天数
	USERDATA_DAILY_FIRST_PLAY_TIME      = 10033 //玩家某天首次玩牌时间（玩家最近玩牌的那天的第一次玩牌时间）
	USERDATA_PLAY_DAYS                  = 10034 //玩家总的玩牌天数
	USERDATA_MAIL_RED_NUM               = 10035 //邮件红点数
	USERDATA_TREE_AWARD_NOTIFY          = 10036 //家园摇钱树摇钱开关
	USERDATA_PARTNER_COUNTDOWN          = 10037 //成为合伙人倒计时
	USERDATA_PARTNER_FRIEND_DATA_VER    = 10038 //成为合伙人倒计时
)

//社交相关设置的比特位定义
const (
	SOCIAL_SETTING_BIT_ADD_FRIEND  = 0
	SOCIAL_SETTING_BIT_INVITE_PLAY = 1
)

//社交好友状态的游戏状态
const (
	SOCIAL_GAME_STATUS_DISABLE = 0 //不可以邀请和追踪
	SOCIAL_GAME_STATUS_ALL     = 1 //可以邀请和追踪
	SOCIAL_GAME_STATUS_INVITE  = 2 //可以邀请,不可追踪
	SOCIAL_GAME_STATUS_TRACK   = 3 //不可以邀请,可追踪
)

//用户离开房间原因定义
const (
	KICK_BY_USER                  = 1  //被用户踢
	KICK_BY_ADMIN                 = 2  //被管理员踢
	KICK_FOR_NO_CUSTOM_GAME       = 3  //在规定时间内没有开始好友房牌局踢出
	KICK_FOR_CUSTOM_GAME_TIMEOUT  = 4  //好友房时间达期踢出
	KICK_FOR_CUSTOM_NO_USER       = 5  //好友房规定时间内无人坐下解散踢出
	KICK_FOR_SERVER_ERROR         = 6  //因服务器内部错误踢出
	KICK_FOR_OFFLINE              = 7  //掉线
	KICK_FOR_MONEY_NOT_ENOUGH     = 8  //资金不足
	KICK_FOR_MAX_IDLE_TIMES       = 9  //太久没有参与游戏
	KICK_FOR_USER_NOT_CLICK_READY = 10 //因为一局牌结束之后，玩家一直不点击“准备”按钮，系统将其踢出。【跑牌使用】
	KICK_FOR_SINGLE_LONG_WAIT     = 11 //因为独自一人准备太久无人匹配，客户端收到这个原因后重新帮助该玩家点击配桌。【跑牌使用】
	KICK_FOR_NOT_FULL_TABLE       = 12 //因为桌子没有满，需要解散重新配桌。【跑牌使用】
)

//用户状态
const (
	USER_STATUS_NORMAL   = 0
	USER_STATUS_DISABLED = 1
)

//用户VIP特权
const (
	VIP_PRIVILEGE_BUY_CHIP       int32 = 1 //筹码购买加成，加筹码
	VIP_PRIVILEGE_SUBSIDY        int32 = 2 //破产补助加成，加筹码
	VIP_PRIVILEGE_DRESS_DISCOUNT int32 = 3 //装扮折扣，金币折扣
	VIP_PRIVILEGE_LADDER         int32 = 4 //天梯分加成
	VIP_PRIVILEGE_FRIEND         int32 = 5 //好友数量上限
	VIP_PRIVILEGE_IDENTITY       int32 = 6 //身份特效
	VIP_PRIVILEGE_FREE_SET_NAME  int32 = 7 //每周一次改名字
	VIP_PRIVILEGE_HIPOKER        int32 = 8 //专属客服
	VIP_PRIVILEGE_NOAD           int32 = 9 //免广告
)

//子游戏名称
const (
	GAME_KEY_TEXAS   = "texas"
	GAME_KEY_AOF     = "aof"
	GAME_KEY_SAMGONG = "samgong"
	GAME_KEY_PAOPAI  = "paopai"
)

//桌子事件
const (
	TABLE_EVENT_ADD_PLAYER    = 1
	TABLE_EVENT_DEL_PLAYER    = 2
	TABLE_EVENT_UPDATE_PLAYER = 3
)

//用户个人信息变化类型
const (
	SET_INFO_TYPE_NICKNAME  = 1 //修改昵称
	SET_INFO_TYPE_LABELING  = 2 //Labeling
	SET_INFO_TYPE_ICON      = 3 //icon
	SET_INFO_TYPE_AVATAR    = 4 //形象
	SET_INFO_TYPE_COUNTRY   = 5 //国家
	SET_INFO_TYPE_LANGUAGE  = 6 //语言
	SET_INFO_TYPE_PLATFORM  = 7 //平台
	SET_INFO_TYPE_VIP_LEVEL = 8 //vip升级
	SET_INFO_TYPE_VIP_HIDE  = 9 //vip显示隐藏
)

//房间内聊天消息类型
const (
	USER_BROADCAST_MSG_SIMPLE  = 1 //快捷聊天
	USER_BROADCAST_MSG_FACES   = 2 //发表情
	USER_BROADCAST_MSG_TEXT    = 3 //文本聊天
	USER_BROADCAST_MSG_DYNAMIC = 4 //动态表情
)

//服务事件类型
const (
	SERVER_EVENT_START       = 1 //服务启动
	SERVER_EVENT_STOP        = 2 //服务停止
	SERVER_EVENT_QUIT        = 3 //服务退出
	SERVER_EVENT_UNREACHABLE = 4 //服务不可达
)

//家园NSQ消息类型
const (
	HOMELAND_ACTION_EXP_CHANGED          = 1 //经验值变化
	HOMELAND_ACTION_LADDER_SCORE_CHANGED = 2 //天梯分变化
	HOMELAND_ACTION_GAME_MONEY_CHANGED   = 3 //游戏结算(输或赢) - 废弃
)

//宠物技能定义
const (
	PET_FEATURE_ADD_EXP          = 1 //经验值加成
	PET_FEATURE_ADD_LADDER_SCORE = 2 //天梯分加成
	//PET_FEATURE_REDUCE_LOSE_MONEY = 3 //输钱补偿
	//PET_FEATURE_ADD_WON_MONEY     = 4 //赢钱加成
	PET_FEATURE_ADD_TREE_MONEY = 5 //摇钱树收益加成
	PET_FEATURE_DROP_ITEM      = 6 //摇钱树掉落道具
	PET_FEATURE_STOP_STEAL     = 7 //摇钱树防偷
	PET_FEATURE_SHOW_IN_GAME   = 8 //上桌
	PET_FEATURE_TREE_SPEED_UP  = 9 //摇钱树加速成长
)

//用于控制灰度下载的开关
const (
	PACKAGE_SWITCH_ALL   = 0
	PACKAGE_SWITCH_WHITE = 1
	PACKAGE_SWITCH_GRAY  = 2
)

//包管理系统资源类型定义
const (
	RES_TYPE_LANG_RES         = "lang_res"
	RES_TYPE_GAME_COMM        = "game_comm"
	RES_TYPE_GAME_COMM_LANG   = "comm_lang"
	RES_TYPE_GAME_RES         = "game_res"  //子游戏包
	RES_TYPE_GAME_LANG_RES    = "game_lang" //子游戏语言包
	RES_TYPE_LAZY_FOR_COUNTRY = "lazy4C"    //延迟加载资源(By Country)
	RES_TYPE_LAZY             = "lazy"      //延迟加载资源(Not By Country)
)
