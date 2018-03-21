package users

import (
	"time"
)

type Users struct {
	Id                  int       `orm:"column(id);pk"`                 // 用户id
	Code                string    `orm:"column(code)"`                  // 用户编码
	Account             string    `orm:"column(account)"`               // 帐号
	LoginPassword       string    `orm:"column(login_password)"`        // 登录密码
	State               int       `orm:"column(state)"`                 // 用户状态
	CreateTime          time.Time `orm:"column(create_time)"`           // 注册时间
	CreateDate          time.Time `orm:"column(create_date)"`           // 注册日期
	MobileType          string    `orm:"column(mobile_type)"`           // 手机类型
	MobileVersion       string    `orm:"column(mobile_version)"`        // 手机版本
	LoginTime           time.Time `orm:"column(login_time)"`            // 登录时间
	HisDevicecode       string    `orm:"column(his_devicecode)"`        // 手机设备标识
	MobileTypeRecent    string    `orm:"column(mobile_type_recent)"`    // 最近操作的结果，手机型号
	OperationTime       time.Time `orm:"column(operation_time)"`        // 最近操作的时间（1.2.1 版本加入的获取用户接口）
	App                 int       `orm:"column(app)"`                   // 1 ios,2 android,3 wx,4pc5 wd 6 wap,7:落地页
	Token               string    `orm:"column(token)"`                 // 服务器保存一致信息
	PkgType             int       `orm:"column(pkg_type)"`              // 包的标识
	Source              string    `orm:"column(source)"`                // 渠道
	OutPutSource        string    `orm:"column(out_put_source)"`        // 外放渠道
	MobileVersionRecent string    `orm:"column(mobile_version_recent)"` // 最近操作的结果，手机app 版本号
	AppVersion          string    `orm:"column(app_version)"`           // app版本
	AppVersionRecent    string    `orm:"column(app_version_recent)"`    // 最近操作app版本
	UserImg             string    `orm:"column(user_img)"`              // 用户头像
	Isemulator          int       `orm:"column(isemulator)"`            // 是否是模拟器
	Jpushid             string    `orm:"column(jpushid)"`               // 手机推送标识
	LoanControl         int       `orm:"column(loan_control)"`          // 借款控制: 0进行中有一笔不能申请第二笔,1:不做限制
	Ip                  string    `orm:"column(ip)"`                    // IP地址
	HisDevicecodeRecent string    `orm:"column(his_devicecode_recent)"` // 最近手机设备标识
	Location            string    `orm:"column(location)"`              // 根据开启的定位，获取到的市
	Address             string    `orm:"column(address)"`               // 根据开启的定位，获取到的详细地址
	IpLocation          string    `orm:"column(ip_location)"`           // 根据IP地址，获取到的市
	IpAddress           string    `orm:"column(ip_address)"`            // 根据IP地址，获取到的详细地址
	IsIndexPopup        int       `orm:"column(is_index_popup)"`        // 0:弹窗1：不弹窗
	IsAuthPopup         int       `orm:"column(is_auth_popup)"`         // 0:弹窗1：不弹窗
	IsSignPrompt        int       `orm:"column(is_sign_prompt)"`        // 是否开启签到提醒   0：不提醒   1：提醒
	LastSignDate        time.Time `orm:"column(last_sign_date)"`        // 最近签到时间
	HasZmxy             int       `orm:"column(has_zmxy)"`              // 0:没有芝麻信用1：有芝麻信用
	ModifyTime          time.Time `orm:"column(modify_time)"`           // 修改时间
	ActiveTime          time.Time `orm:"column(active_time)"`           // 激活时间
	RegisterSource      int       `orm:"column(register_source)"`       // 1:主包，>1：分包
	WxToken             string    `orm:"column(wx_token)"`              // 微融生成的微信公众号登录token
	ActiveDate          time.Time `orm:"column(active_date)"`
	RecentApp           int       `orm:"column(recent_app)"`   // 1 ios,2 android,3 wx,4pc5 wd 6 wap,7:落地页
	HezuoSource         string    `orm:"column(hezuo_source)"` // 微融合作平台source
}
