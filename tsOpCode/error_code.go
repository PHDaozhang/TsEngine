package tsOpCode

// 错误代码，每个模块占用 100个。并给予说明（什么用途）
const (
	// 基础操作结果
	OPERATION_SUCCESS = 200

	// 基本错误信息 10000~10099
	OPERATION_REQUEST_FAILED   = 10000
	OPERATION_DB_FAILED        = 10001
	DATA_NOT_EXIST             = 10002
	TIME_OUT                   = 10003
	REPEATED_SUBMIT            = 10004
	EXIST_CHILD_DEPENDENCY     = 10005 // 存在子项依赖
	OPERATION_DELETE_FAILED    = 10006 // 删除失败
	PARENT_NOT_AVAILABLE       = 10007 // 父项不存在
	USER_NOT_LOGIN             = 10008
	ACCOUNT_DENIDE             = 10009
	NO_PERMISSION              = 10010
	IP_BLOCKED                 = 10011
	RETURN_RATE_SET_HIGH_ERROR = 10012
	PASSWORD_INCONSISTENT      = 10013
	SERVER_ERROR               = 10014
	OPERATION_NOT_EXIST        = 10015
	GET_PAGES_ERROR            = 10016
	DATA_EXIST                 = 10017
	ACCOUNT_ERROR              = 10018
	NO_PERMISSION_UPDATE_CHILD = 10019
	MAKE_DIR_FAILED            = 10020
	JSON_UNMARSHAL_FAILED      = 10021
	ACCOUNT_ALREADY_LOGIN      = 10022 // 账户已登陆
	CONTACT_INF_EXISITS        = 10023 // 联系方式重复
	CLOSE_ACCOUNT_FIRST        = 10024 // 请先关闭提现
	DISABLED_ACCOUNT_FIRST     = 10025 // 请先禁用账户

	// cloud
	USER_NAME_EXIST                = 10100
	TEL_NUMBER_EXISTED             = 10101
	PASSWORD_NOT_ALLOW_USER_UPDATE = 10102
	PASSWORD_ERROR                 = 10103
	CONFIG_FILE_READ_FAILED        = 10104
	GET_UPLOAD_FILE_ERROR          = 10105
	SAVE_UPLOAD_FILE_ERROR         = 10106
	SPEED_TASK_ADD_ERROR           = 10107
	WITHDRAW_ORDER_CLOSED          = 10108 //提现订单已关闭
	PASSWORD_CONFIRM_FAILED        = 10109 //确认密码错误
	NOT_SELF_ROLES                 = 10110 //不可修改非自建角色
	ORDER_SATATUS_ERROR            = 10111 //订单状态错误
	CONNIT_REPEAT_ORDER            = 10112 //重复提交订单
	AGENT_NOTICE_OPENED            = 10113 //代理存在已开启的公告
	NO_HAVE_OPEN_CHANNEL           = 10114 //没有开启的支付渠道
	CHANNEL_LIMITED                = 10115 //所有渠道提现额度到上限
	REQUEST_APPROVALED             = 10116 //所批复申请已完成
	APPROVAL_AMOUNT_ERROR          = 10117 //批复金额错误
	RATE_CAN_NOT_HEIGH_THEN_SELF   = 10118 //比率不可高于自身
	DONT_COVER_YOUR_PROMISS        = 10119 //不可越权
	PASSWORD_CANT_EMPTY            = 10120 //密码不可为空
	PASSWORD_SHORTLY               = 10121 //密码过短
	PASSWORD_LONG                  = 10122 //密码过长
	MOBILE_ERROR                   = 10123 //手机号错误
	CREATE_ROLE_CHILD_DENIDE       = 10124 //不可创建该角色员工
	ACCOUNT_SCORE_NOT_ENOUGH       = 10125 //账户余额不足
	GAME_SERVER_DISCONNECTED       = 10126 //游戏服务器未连接，操作失败
	DIACTIVITY_FIRST_BEFORE_EDIT   = 10127 //账户未停用不可编辑
	COLUD_COIN_NOT_ENOUGH          = 10128 //云币不足
	PLAYER_CARD_NOT_FOUND          = 10129 //获取玩家提现绑定信息失败
	NO_WITHDRAW_CHANNEL            = 10130 //没有可用的提现渠道
	AGENT_SCORE_NOT_ENOUGH         = 10131 //身上余额不足
	RECORED_STATUS_ERROR           = 10132 //记录状态异常
	ACCOUNT_STATUS_ERROR           = 10133 //账号已拒绝登录
	ROLES_INVALID_ERROR            = 10134 //请检查角色
	USERNAME_IS_EXISTS             = 10135 //账号已存在
	INVALID_USERNAME               = 10136 //账号长度异常或有非法字符
	INVALID_PASSWORD               = 10137 //密码长度异常或有非法字符
	INVALID_STATUS                 = 10138 //账号状态设置错误
	INVALID_NAME                   = 10139 //不可用的账号名称
	PALYER_NOT_EXITS               = 10140 //玩家不存在
	INVALID_IP                     = 10141 //IP不合法
	UUID_GENRATE_FAILED            = 10142 //UUID生成失败
	KEYS_GENRATE_FAILED            = 10143 //密钥生成失败
	CONFIG_FILE_GENRATE_FAILED     = 10144 //配置文件生成错误
	WAS_SERVER_REQ_ERROR           = 10145 //WAS服务器请求出错
	WITHDRAW_ACCOUNT_ALL_OFFLINE   = 10146 //提现账户全部离线
	CHANNEL_IOS_SIGN_CLOSED        = 10147 //渠道超级签开关
	IOS_SIGN_TIMES_NOT_ENOUGH      = 10148 //渠道超级签次数不足

	// domain
	DOMAIN_ERROR_INTERFACE             = 10230 // 可能是由于被限制次数导致的
	DOMAIN_ERROR_REGISTERED            = 10231 // 已被注册
	DOMAIN_ERROR_NETWORK               = 10232 // 网络出错
	DOMAIN_ERROR_TLD                   = 10233 // 不支持的域名后缀
	DOMAIN_ERROR_SCHEMA                = 10234 // 可能是注册模板有误
	DOMAIN_ERROR_PAY                   = 10235 // 支付失败(无支付信息，或余额不足)
	DOMAIN_ERROR_UNKNOWN               = 10236 // 未知错误
	DOMAIN_REG_SUC_BUT_INSERT_DB_ERROR = 10237
	SEARCH_BEFORE_BUY                  = 10238 //购买前请先搜索
	DOMAIN_SOLD                        = 10239 //域名已售

	// db error
	DB_CREATE_ERROR = 10300
	DB_DEL_ERROR    = 10301
	DB_INSERT_ERROR = 10302
	DB_UPDATE_ERROR = 10303
	DB_SELECT_ERROR = 10304

	// sts错误信息 11000-11099
	APPNAME_EXISTED = 11000
	TOKEN_ERROR     = 11001
	APP_TOKEN_ERROR = 11002

	// 微服务错误code 11100-11199
	NO_AUTHORIZATION   = 11100
	LOGIN_STS_ERROR    = 11101
	LOGIN_STS_FAILED   = 11102
	CHECK_TOKEN_ERROR  = 11103 // 失败
	CHECK_TOKEN_FAILED = 11104 // token有误
	SIGN_FAILED        = 11105 // 签名校验失败
	EXPED_REQUEST      = 11106 // 无效(过期)请求
	REPEAT_ORDER       = 11107 // 重复的订单号
	NO_SUIT_CHANNEL    = 11108 // 没有合适的充值通道

	// 提现 11200-11299
	ORDER_ERROR    = 11200 // 订单处理出错，无法重试。可能需要新建订单
	PLATFORM_ERROR = 11201 // 转账平台出错，可能是账号出了问题
	ORDER_PAYED    = 11202 // 转账平台已经处理过该订单
	TRANS_FAILED   = 11203 // 转账失败，可能是收款人信息有误
	TRANS_ERROR    = 11204 // 转账失败
	CLIENT_BUSY    = 11205 // 客户端正忙

	// 充值渠道管理 11300-11399
	RECHARGE_RANGE_ERROR           = 11300 // 充值区间有误
	RECHARGE_CHANNEL_VIP_CONFIG    = 11301 // 优质通道配置有误
	RECHARGE_CHANNEL_NORMAL_CONFIG = 11302 // 普通通道配置有误

	// 提现连接服务器 11400-11500
	UUID_CLIENT_DISCONNECTED = 11400 //此UUID的客户端已下线
	CHANNEL_NOT_EXISIT       = 11401 //渠道不存在
	CLIENT_IS_EXIST          = 11402 //客户端已存在

	//云平台管理系统 12000-12999
	DATA_MASHAL_FAILED  = 12000 //结构转换失败
	INVLID_CONNECT_INFO = 12001 //联系方式错误

	//客服平台13000-13999
	HAVE_NO_CUSTOMERSERVICE     = 13000 //CustomerService表没有该客服
	HAVE_NO_CUSTOMERSERVICETIME = 13001 //CommuterTime表没有该客服
	NO_CORRESPONDING_FIELD      = 13002 //更新失败，没有相应字段
	NO_SUCH_ORDER               = 13003 //没查到该订单账号相关附件

	//文件
	FILE_LARGE         = 11400 //文件过大
	COMPRESSION_FAILED = 11401 //图片压缩失败

)
