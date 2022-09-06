package conf

const (
	//Access 玩家异地登录
	CMD_ACCESS_ANO_LOGIN int32 = 0x1001 //玩家异地登录
	//通用广播命令字
	CMD_COMMON_BROADCAST int32 = 0x1002
	//牌局成就达成通知
	CMD_ACHIEVE_FINISHED int32 = 0x1003 //玩家达成成就通知
	//玩家等级变化通知
	CMD_EXP_LEVEL_CHANGED int32 = 0x1004 //玩家等级变化通知
	//用户天梯星级变化通知
	CMD_LADDER_CHANGED int32 = 0x1005 //用户天梯星级变化通知
	//添加好友请求消息通知
	CMD_FRIEND_REQUEST int32 = 0x1006 //添加好友请求消息通知
	//未读好友消息数量变化通知
	CMD_NEW_CHAT_MSG int32 = 0x1007 //未读好友消息数量变化通知
	//邀请玩牌的通知
	CMD_INVITE_PLAY int32 = 0x1008 //邀请玩牌的通知
	//好友请求被处理的通知
	CMD_FRIEND_REQUEST_RESULT int32 = 0x1009 //好友请求被处理的通知
	//被好友被删除的通知
	CMD_DELETE_BY_FRIEND int32 = 0x1010 //被好友被删除的通知
	//数据有更新通知
	CMD_COMMON_DATA_CHANGE_NOTIFY int32 = 0x1011 //数据有更新通知
	//支付成功通知
	CMD_PAY_SUCCESS_NOTIFY int32 = 0x1012 //支付成功之后通知客户端（该命令字区别于0X4100，收到0X1012时客户端会展示出得到的物品，0X4100只会修改数值）
	//玩牌活动轮次更新通知
	CMD_ACTIVTIY_ROUND_UPDATE_NOTIFY = 0x1013
	//用户VIP等级变化
	CMD_VIP_LEVEL_CHANGED int32 = 0x1014 //用户VIP等级变化
	//用户生成回归礼包
	CMD_GEN_RETURN_GIFT int32 = 0x1015 //用户生成回归礼包
	//好友状态变更
	CMD_FRIEND_STATUS_UPDATE int32 = 0x1016 //好友状态变更
	//天梯段位榜奖励
	CMD_LADDER_LEVEL_REWARD int32 = 0x1017 //天梯段位榜奖励，可以领取(每日玩牌可领取)
	//天梯积分变化推送
	CMD_LADDER_SCORE_CHANGE int32 = 0x1018 //天梯积分变化推送
	//支付成功通知,推送新协议，支持通用的弹窗
	CMD_PAY_SUCCESS_NOTIFY2 int32 = 0x1019 //支付成功之后通知客户端（该命令字区别于0X4100，收到0X1012时客户端会展示出得到的物品，0X4100只会修改数值）
	//用户段位榜排名升级
	CMD_RANKING_UPGRADE int32 = 0x1020
	//支付成功后推送
	CMD_PAY_SUCC_PUSH int32 = 0x1021
	//game用命令字
	CMD_GAME_USER_SIT               int32 = 0x2001 //玩家A坐下时，广播消息给桌子上其他人（广播通知给桌子上其他人）
	CMD_GAME_TABLE_ENV              int32 = 0x2002 //玩家A坐下时，将桌子的环境信息广播给A
	CMD_GAME_STAGE                  int32 = 0x2003 //广播桌子阶段值
	CMD_GAME_HAND_CARD              int32 = 0x2004 //广播玩家手牌
	CMD_GAME_BLIND                  int32 = 0x2005 //广播大小盲注
	CMD_GAME_INVITE_CHIP            int32 = 0x2006 //通知玩家来下注
	CMD_GAME_CHIP_STATUS            int32 = 0x2007 //通知下注人，本次操作成功or失败
	CMD_GAME_CHIP_SUCCESS_TO_OTHERS int32 = 0x2008 //将下注人成功或者失败的信息，广播给其他人！(排除自己)
	CMD_GAME_POTS                   int32 = 0x2009 //广播奖池信息的命令
	CMD_GAME_DEALCARD               int32 = 0x200A //广播公牌的命令
	CMD_GAME_SETTLE_ACCOUNT         int32 = 0x200B //广播结算信息
	CMD_GAME_DEALER                 int32 = 0x200C //广播庄家的座位号
	CMD_GAME_STAND                  int32 = 0x200D //广播用户站起
	CMD_GAME_RECONNECT              int32 = 0x200E //广播重连的信息
	CMD_GAME_EXIT                   int32 = 0x200F //广播退出的信息
	CMD_GAME_MONEY_CHIP             int32 = 0x2010 //广播玩家金币和筹码信息（桌子内玩家）
	CMD_GAME_TABLE_ONLINE           int32 = 0x2011 //广播桌子在线人数情况（只返回有人的桌子）
	CMD_GAME_TABLE_FEE              int32 = 0x2012 //广播扣除台费的消息
	//CMD_GAME_RETIRE_LOGIN	int32 = 0x2013		//退休状态下,不允许进入桌子
	CMD_GAME_ANTE_FEE               int32 = 0x2014 //广播扣除前注的消息
	CMD_GAME_SHOW_CARDS             int32 = 0x2015 //广播亮手牌(all-in之后的动画)
	CMD_GAME_USER_BROADCAST         int32 = 0x2016 //用户房间内广播(聊天、表情)
	CMD_GAME_STANDING_NUMBER_CHANGE int32 = 0x2017 //广播房间内旁观人数变化
	CMD_GAME_SHOW_USER_CARDS        int32 = 0x2018 //广播结算阶段弃牌用户亮手牌

	//广播通知玩家进行游戏内重连
	CMD_GAME_INVITE_RECONNECT       int32 = 0x2020 //广播通知玩家，邀请其重连回桌子
	CMD_GAME_INVITE_MATCH_RECONNECT int32 = 0x2021 //广播通知玩家，邀请其重连回比赛

	//德州JACKPOT相关
	CMD_GAME_TEXAS_JACKPOT_UPDATE int32 = 0x2022 //通知用户Jackpot金额变化
	CMD_GAME_TEXAS_JACKPOT_WON    int32 = 0x2023 //通用用户Jackpot开奖

	//百人场命令字
	CMD_GAME_GAMBLE_START              = 0x2100 //百人场牌局开始
	CMD_GAME_GAMBLE_STAGE_UPDATE       = 0x2101 //百人场牌局状态变化
	CMD_GAME_GAMBLE_OPTIONS_UPDATE     = 0x2102 //百人场下注选项数据更新
	CMD_GAME_GAMBLE_RESULT             = 0x2103 //百人场结算
	CMD_GAME_GAMBLE_USER_BROADCAST     = 0x2104 //百人场用户广播(表情、聊天)
	CMD_GAME_GAMBLE_TAKE_DEALER_RESULT = 0x2105 //百人场通知用户上庄结果
	GMD_GAME_GAMBLE_USER_EXIT          = 0x2106 //百人场用户离开通知

	//三公命令字
	CMD_GAME_SAMGONG_START             = 0x2200 //三公牌局开始
	CMD_GAME_SAMGONG_STAGE_UPDATE      = 0x2201 //三公牌局状态变化
	CMD_GAME_SAMGONG_ASK_ACTION        = 0x2202 //三公通知玩家进行操作
	CMD_GAME_SAMGONG_ACTION_RESULT     = 0x2203 //三公用户操作结果
	CMD_GAME_SAMGONG_RESULT            = 0x2204 //三公牌局结算
	CMD_GAME_SAMGONG_USER_SIT          = 0x2205 //三公用户坐下
	CMD_GAME_SAMGONG_USER_STAND        = 0x2206 //三公用户站起
	CMD_GAME_SAMGONG_USER_EXIT         = 0x2207 //三公通知用户离开
	CMD_GAME_SAMGONG_STANDING_UPDATE   = 0x2208 //三公旁观人数变化
	CMD_GAME_SAMGONG_USER_MONEY_UPDATE = 0x2209 //三公用户筹码数量变化
	CMD_GAME_SAMGONG_USER_BROADCAST    = 0x2210 //三公用户房间内广播(快捷聊天、自定义聊天、表情)

	//跑牌命令字
	CMD_GAME_PAOPAI_ENTER_SUCCESS      int32 = 0x2300 //跑牌用户进桌子成功的消息！（进桌子的玩家才会收到此消息）
	CMD_GAME_PAOPAI_USER_SIT_AND_READY int32 = 0x2301 //跑牌用户坐下且准备了
	CMD_GAME_PAOPAI_USER_EXIT          int32 = 0x2302 //跑牌通知用户离开
	CMD_GAME_PAOPAI_COUNT_DOWN_SECONDS int32 = 0x2303 //跑牌-游戏开局倒计时（单位：秒）
	CMD_GAME_PAOPAI_START              int32 = 0x2304 //跑牌牌局开始（整个桌子的快照信息）
	CMD_GAME_PAOPAI_STAGE_UPDATE       int32 = 0x2305 //跑牌牌局状态变化（1表示开始 99表示结算了 100表示结束了）
	CMD_GAME_PAOPAI_TABLE_FEE          int32 = 0x2306 //跑牌广播扣除台费的消息
	CMD_GAME_PAOPAI_HAND_CARD          int32 = 0x2307 //跑牌广播玩家手牌
	CMD_GAME_PAOPAI_INVITE_ACTION      int32 = 0x2308 //跑牌通知玩家进行操作（邀请玩家出牌）
	CMD_GAME_PAOPAI_ACTION_RESULT      int32 = 0x2309 //跑牌用户操作结果（广播玩家的出牌行为）
	CMD_GAME_PAOPAI_ACTION_ZHANPAI     int32 = 0x230A //跑牌斩牌动作发生（广播斩牌发生）
	CMD_GAME_PAOPAI_RESULT             int32 = 0x230B //跑牌牌局结算
	CMD_GAME_PAOPAI_USER_BROADCAST     int32 = 0x230C //跑牌用户房间内广播(快捷聊天、自定义聊天、表情)
	CMD_GAME_PAOPAI_USER_READY         int32 = 0x230D //跑牌用户准备
	CMD_GAME_PAOPAI_USER_MONEY_UPDATE  int32 = 0x230E //跑牌用户筹码数量变化
	CMD_GAME_PAOPAI_ROUND_FINISH       int32 = 0x2310 //跑牌一轮结束（客户端收到这个消息，进行一些桌面清理工作）

	//炸金花命令字
	CMD_GAME_ZJH_START             = 0x2400 //炸金花牌局开始
	CMD_GAME_ZJH_STAGE_UPDATE      = 0x2401 //炸金花牌局状态变化
	CMD_GAME_ZJH_ASK_ACTION        = 0x2402 //炸金花通知玩家进行操作
	CMD_GAME_ZJH_ACTION_RESULT     = 0x2403 //炸金花用户操作结果
	CMD_GAME_ZJH_RESULT            = 0x2404 //炸金花牌局结算
	CMD_GAME_ZJH_USER_SIT          = 0x2405 //炸金花用户坐下
	CMD_GAME_ZJH_USER_STAND        = 0x2406 //炸金花用户站起
	CMD_GAME_ZJH_USER_EXIT         = 0x2407 //炸金花通知用户离开
	CMD_GAME_ZJH_STANDING_UPDATE   = 0x2408 //炸金花旁观人数变化
	CMD_GAME_ZJH_USER_MONEY_UPDATE = 0x2409 //炸金花用户筹码数量变化
	CMD_GAME_ZJH_USER_BROADCAST    = 0x2410 //炸金花用户房间内广播(快捷聊天、自定义聊天、表情)
	CMD_GAME_ZJH_SHOW_CARDS        = 0x2411 //炸金花用户亮牌

	//多米诺99命令字
	CMD_GAME_DOMINO99_START             = 0x2500 //多米诺99牌局开始
	CMD_GAME_DOMINO99_STAGE_UPDATE      = 0x2501 //多米诺99牌局状态变化
	CMD_GAME_DOMINO99_ASK_ACTION        = 0x2502 //多米诺99要求用户进行下注操作
	CMD_GAME_DOMINO99_ACTION_RESULT     = 0x2503 //多米诺99转发用户下注操作结果
	CMD_GAME_DOMINO99_ASK_CONFIRM       = 0x2504 //多米诺99要求所有玩家确认牌型
	CMD_GAME_DOMINO99_SUBMIT_CONFIRM    = 0x2505 //多米诺99转发用户确认牌型状态
	CMD_GAME_DOMINO99_RESULT            = 0x2506 //多米诺99牌局结算
	CMD_GAME_DOMINO99_USER_SIT          = 0x2507 //多米诺99用户坐下
	CMD_GAME_DOMINO99_USER_STAND        = 0x2508 //多米诺99用户站起
	CMD_GAME_DOMINO99_USER_EXIT         = 0x2509 //多米诺99通知用户离开
	CMD_GAME_DOMINO99_STANDING_UPDATE   = 0x2510 //多米诺99旁观人数变化
	CMD_GAME_DOMINO99_USER_MONEY_UPDATE = 0x2511 //多米诺99用户筹码数量变化
	CMD_GAME_DOMINO99_USER_BROADCAST    = 0x2512 //多米诺99用户房间内广播(快捷聊天、自定义聊天、表情)

	//十三张命令字
	CMD_GAME_CAPSASUSUN_START             = 0x2600 //十三张-牌局开始
	CMD_GAME_CAPSASUSUN_STAGE_UPDATE      = 0x2601 //十三张-牌局状态变化
	CMD_GAME_CAPSASUSUN_ACTION_RESULT     = 0x2602 //十三张-用户提交牌型
	CMD_GAME_CAPSASUSUN_SHOW_CARDS        = 0x2603 //十三张-比牌动画
	CMD_GAME_CAPSASUSUN_RESULT            = 0x2604 //十三张-牌局结算
	CMD_GAME_CAPSASUSUN_USER_SIT          = 0x2605 //十三张-用户坐下
	CMD_GAME_CAPSASUSUN_USER_STAND        = 0x2606 //十三张-用户站起
	CMD_GAME_CAPSASUSUN_USER_EXIT         = 0x2607 //十三张-通知用户离开
	CMD_GAME_CAPSASUSUN_STANDING_UPDATE   = 0x2608 //十三张-旁观人数变化
	CMD_GAME_CAPSASUSUN_USER_MONEY_UPDATE = 0x2609 //十三张-用户筹码数量变化
	CMD_GAME_CAPSASUSUN_USER_BROADCAST    = 0x2610 //十三张-用户房间内广播(快捷聊天、自定义聊天、表情)

	//大米命令字
	CMD_GAME_DUMMY_ENTER_SUCCESS      int32 = 0x2700 //大米用户进桌子成功的消息！（进桌子的玩家才会收到此消息）
	CMD_GAME_DUMMY_USER_SIT_AND_READY int32 = 0x2701 //大米用户坐下且准备了
	CMD_GAME_DUMMY_USER_EXIT          int32 = 0x2702 //大米通知用户离开
	CMD_GAME_DUMMY_COUNT_DOWN_SECONDS int32 = 0x2703 //大米-游戏开局倒计时（单位：秒）
	CMD_GAME_DUMMY_START              int32 = 0x2704 //大米牌局开始（整个桌子的快照信息）
	CMD_GAME_DUMMY_STAGE_UPDATE       int32 = 0x2705 //大米牌局状态变化（1表示开始 99表示结算了 100表示结束了）
	CMD_GAME_DUMMY_TABLE_FEE          int32 = 0x2706 //大米广播扣除台费的消息
	CMD_GAME_DUMMY_HAND_CARD          int32 = 0x2707 //大米广播玩家手牌
	CMD_GAME_DUMMY_INVITE_ACTION      int32 = 0x2708 //大米通知玩家进行操作（邀请玩家出牌）
	CMD_GAME_DUMMY_ACTION_RESULT      int32 = 0x2709 //大米用户操作结果（广播玩家的出牌行为）
	CMD_GAME_DUMMY_ACTION_ZHANPAI     int32 = 0x270A //大米斩牌动作发生（广播斩牌发生）
	CMD_GAME_DUMMY_RESULT             int32 = 0x270B //大米牌局结算
	CMD_GAME_DUMMY_USER_BROADCAST     int32 = 0x270C //大米用户房间内广播(快捷聊天、自定义聊天、表情)
	CMD_GAME_DUMMY_USER_READY         int32 = 0x270D //大米用户准备
	CMD_GAME_DUMMY_USER_MONEY_UPDATE  int32 = 0x270E //大米用户筹码数量变化
	CMD_GAME_DUMMY_ROUND_FINISH       int32 = 0x2710 //大米一轮结束（客户端收到这个消息，进行一些桌面清理工作）
	CMD_GAME_DUMMY_TUOGUAN_STATUS     int32 = 0x2711 //大米-玩家托管状态

	//斗地主命令字
	CMD_GAME_DDZ_ENTER_SUCCESS      int32 = 0x2800 //斗地主用户进桌子成功的消息！（进桌子的玩家才会收到此消息）
	CMD_GAME_DDZ_USER_SIT_AND_READY int32 = 0x2801 //斗地主用户坐下且准备了
	CMD_GAME_DDZ_USER_EXIT          int32 = 0x2802 //斗地主通知用户离开
	CMD_GAME_DDZ_COUNT_DOWN_SECONDS int32 = 0x2803 //斗地主-游戏开局倒计时（单位：秒）
	CMD_GAME_DDZ_START              int32 = 0x2804 //斗地主牌局开始（整个桌子的快照信息）
	CMD_GAME_DDZ_STAGE_UPDATE       int32 = 0x2805 //斗地主牌局状态变化（1表示开始 99表示结算了 100表示结束了）
	CMD_GAME_DDZ_TABLE_FEE          int32 = 0x2806 //斗地主广播扣除台费的消息
	CMD_GAME_DDZ_HAND_CARD          int32 = 0x2807 //斗地主广播玩家手牌
	CMD_GAME_DDZ_INVITE_CHUPAI      int32 = 0x2808 //斗地主通知玩家进行操作（邀请玩家出牌）
	CMD_GAME_DDZ_INVITE_QIANG       int32 = 0x2809 //斗地主邀请玩家来抢（邀请玩家抢地主）
	CMD_GAME_DDZ_INVITE_JIABEI      int32 = 0x280A //斗地主邀请玩家来加倍（邀请玩家加倍）
	CMD_GAME_DDZ_BEISHU_INFO        int32 = 0x280B //斗地主桌子上倍数信息
	CMD_GAME_DDZ_DIPAI              int32 = 0x280C //斗地主-广播底牌信息
	CMD_GAME_DDZ_ACTION_MING        int32 = 0x280D //斗地主用户明牌操作结果（广播玩家的出牌行为）
	CMD_GAME_DDZ_ACTION_QIANG       int32 = 0x280E //斗地主用户抢地主操作结果（广播玩家的出牌行为）
	CMD_GAME_DDZ_ACTION_JIABEI      int32 = 0x280F //斗地主用户加倍操作结果（广播玩家的出牌行为）
	CMD_GAME_DDZ_ACTION_CHUPAI      int32 = 0x2810 //斗地主用户出牌操作结果（广播玩家的出牌行为）
	CMD_GAME_DDZ_LASTONE_TUOGUAN    int32 = 0x2811 //斗地主-广播玩家最后一张手牌托管了。
	CMD_GAME_DDZ_TIMEOUT_TUOGUAN    int32 = 0x2812 //斗地主-玩家操作超时的托管
	CMD_GAME_DDZ_JIPAIQI            int32 = 0x2813 //斗地主-玩家的记牌器信息
	CMD_GAME_DDZ_RESULT             int32 = 0x2814 //斗地主牌局结算
	CMD_GAME_DDZ_USER_BROADCAST     int32 = 0x2815 //斗地主用户房间内广播(快捷聊天、自定义聊天、表情)
	CMD_GAME_DDZ_USER_READY         int32 = 0x2816 //斗地主用户准备
	CMD_GAME_DDZ_USER_MONEY_UPDATE  int32 = 0x2817 //斗地主用户筹码数量变化
	CMD_GAME_DDZ_ROUND_FINISH       int32 = 0x2818 //斗地主一轮结束（客户端收到这个消息，进行一些桌面清理工作）

	//KING牌命令字
	CMD_GAME_KING_ENTER_SUCCESS      int32 = 0x2900 //King牌用户进桌子成功的消息！（进桌子的玩家才会收到此消息）
	CMD_GAME_KING_USER_SIT_AND_READY int32 = 0x2901 //King牌用户坐下且准备了
	CMD_GAME_KING_USER_EXIT          int32 = 0x2902 //King牌通知用户离开
	CMD_GAME_KING_COUNT_DOWN_SECONDS int32 = 0x2903 //King牌-游戏开局倒计时（单位：秒）
	CMD_GAME_KING_START              int32 = 0x2904 //King牌牌局开始（整个桌子的快照信息）
	CMD_GAME_KING_STAGE_UPDATE       int32 = 0x2905 //King牌牌局状态变化（1表示开始 99表示结算了 100表示结束了）
	CMD_GAME_KING_TABLE_FEE          int32 = 0x2906 //King牌广播扣除台费的消息
	CMD_GAME_KING_HAND_CARD          int32 = 0x2907 //King牌广播玩家手牌
	CMD_GAME_KING_INVITE_ACTION      int32 = 0x2908 //King牌通知玩家进行操作（邀请玩家出牌）
	CMD_GAME_KING_ACTION_RESULT      int32 = 0x2909 //King牌用户操作结果（广播玩家的出牌行为）
	CMD_GAME_KING_ACTION_ZHANPAI     int32 = 0x290A //King牌斩牌动作发生（广播斩牌发生）
	CMD_GAME_KING_RESULT             int32 = 0x290B //King牌牌局结算
	CMD_GAME_KING_USER_BROADCAST     int32 = 0x290C //King牌用户房间内广播(快捷聊天、自定义聊天、表情)
	CMD_GAME_KING_USER_READY         int32 = 0x290D //King牌用户准备
	CMD_GAME_KING_USER_MONEY_UPDATE  int32 = 0x290E //King牌用户筹码数量变化
	CMD_GAME_KING_ROUND_FINISH       int32 = 0x2910 //King牌一轮结束（客户端收到这个消息，进行一些桌面清理工作）
	CMD_GAME_KING_TUOGUAN_STATUS     int32 = 0x2911 //King牌-玩家托管状态

	//SNG比赛用命令字
	CMD_MATCH_ANNOUNCE_IN    int32 = 0x3001 //比赛服通知玩家进桌
	CMD_MATCH_WEED_OUT       int32 = 0x3002 //比赛服通知玩家被淘汰
	CMD_MATCH_REBALANCE      int32 = 0x3003 //比赛服通知玩家拆桌，重新进行调度
	CMD_MATCH_LEFT_TIME      int32 = 0x3004 //比赛服通知玩家比赛倒计时
	CMD_MATCH_WITHDRAW       int32 = 0x3005 //比赛服通知玩家比赛解散
	CMD_MATCH_RANKS          int32 = 0x3006 //通知比赛排行榜
	CMD_MATCH_SIGNEDB_NUM    int32 = 0x3007 //通知比赛已签到的人数
	CMD_MATCH_BLIND_INFO     int32 = 0x3008 //开赛时广播盲注和涨盲计划
	CMD_MATCH_SIGNED_NUM     int32 = 0x3009 //定时广播各比赛当前的报名人数
	CMD_MATCH_ANNOUNCE_START int32 = 0x3010 //SNG比赛开始通知

	//二十一点命令字
	CMD_GAME_BLACKJACK_START                = 0x3100 //二十一点牌局开始
	CMD_GAME_BLACKJACK_DEAL_CARDS           = 0x3101 //二十一点发手牌
	CMD_GAME_BLACKJACK_STAGE_UPDATE         = 0x3102 //二十一点牌局状态变化
	CMD_GAME_BLACKJACK_ASK_ACTION           = 0x3103 //二十一点要求用户进行操作
	CMD_GAME_BLACKJACK_ACTION_RESULT        = 0x3104 //二十一点转发用户操作结果
	CMD_GAME_BLACKJACK_DEALER_ACTION_RESULT = 0x3105 //二十一点转发庄家的操作结果
	CMD_GAME_BLACKJACK_RESULT               = 0x3106 //二十一点牌局结算
	CMD_GAME_BLACKJACK_USER_SIT             = 0x3107 //二十一点用户坐下
	CMD_GAME_BLACKJACK_USER_STAND           = 0x3108 //二十一点用户站起
	CMD_GAME_BLACKJACK_USER_EXIT            = 0x3109 //二十一点通知用户离开
	CMD_GAME_BLACKJACK_STANDING_UPDATE      = 0x3110 //二十一点旁观人数变化
	CMD_GAME_BLACKJACK_USER_MONEY_UPDATE    = 0x3111 //二十一点用户筹码数量变化
	CMD_GAME_BLACKJACK_USER_BROADCAST       = 0x3112 //二十一点用户房间内广播(快捷聊天、自定义聊天、表情)

	//通用消息：
	CMD_VMONEY_CHANGE_INFO int32 = 0x4100 //玩家金币变动的消息通知
)
