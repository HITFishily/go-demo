package models

import (
"time"
)

type ActRewardProducts202012 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Type      string `xorm:"not null default 0 comment('类型')  TINYINT(1)"`
	Title     string `xorm:"not null default '' comment('名称') VARCHAR(128)"`
	Image     string `xorm:"not null default '' comment('图片') VARCHAR(512)"`
	Inventory string `xorm:"not null default 0 comment('库存')  INT(10)"`
	WinRate   string `xorm:"not null default 0 comment('中签值')  INT(10)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type ActRewardRecords202012 struct {
	Id           string `xorm:"not null pk autoincr  INT(10)"`
	Uid          string `xorm:"not null default 0 comment('用户微博ID') index  BIGINT(20)"`
	ProductId    string `xorm:"not null default 0 comment('奖品ID') index  INT(10)"`
	ProductTitle string `xorm:"not null default '' comment('奖品名称') VARCHAR(128)"`
	Lucky        string `xorm:"not null default 0 comment('中奖结果') index  TINYINT(1)"`
	CreatedAt    string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt    string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel        int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type ActShares202012 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Uid       string `xorm:"not null default 0 comment('用户微博ID') index  BIGINT(20)"`
	Type      int    `xorm:"not null default 0 comment('类型') index(idx_type_rid) TINYINT(1)"`
	Rid       string `xorm:"not null default 0 comment('对象ID') index(idx_type_rid)  BIGINT(18)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type ActVotes202012 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Uid       string `xorm:"not null default 0 comment('用户微博ID') index  BIGINT(20)"`
	Sn        string `xorm:"not null default 0 comment('投诉号') index  BIGINT(18)"`
	Vote      string `xorm:"not null default 0 comment('状态')  TINYINT(1)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type ActivityVotes struct {
	Id         string `xorm:"not null pk autoincr comment('自增主键')  INT(10)"`
	Uid        string `xorm:"not null default 0 comment('用户微博ID') index  BIGINT(20)"`
	ActivityId string `xorm:"not null default 0 comment('活动ID')  INT(10)"`
	Vote       string `xorm:"not null default 0 comment('票数')  INT(10)"`
	Src        string `xorm:"not null default 0 comment('投票来源')  TINYINT(3)"`
	CreatedAt  string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	IsDel      string `xorm:"not null default 0 comment('是否被删除')  TINYINT(3)"`
}

type ActivityVotes201812 struct {
	Id         string `xorm:"not null pk autoincr comment('自增主键')  INT(10)"`
	Uid        string `xorm:"not null default 0 comment('用户微博ID') index  BIGINT(20)"`
	ActivityId string `xorm:"not null default 0 comment('活动ID')  INT(10)"`
	Vote       string `xorm:"not null default 0 comment('票数')  INT(10)"`
	StarUid    string `xorm:"not null default 0 comment('明星微博ID') index  BIGINT(20)"`
	Src        string `xorm:"not null default 0 comment('投票来源：1.h5 2.黑猫app 3.众测app')  TINYINT(3)"`
	DeviceId   string `xorm:"not null default '' comment('投票设备id') VARCHAR(128)"`
	CreatedAt  string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	IsDel      string `xorm:"not null default 0 comment('是否被删除')  TINYINT(3)"`
}

type ActivityVotes201812InvitationUsers struct {
	Id          string `xorm:"not null pk autoincr comment('自增主键')  INT(10)"`
	Uid         string `xorm:"not null default 0 comment('邀请人UID') index  BIGINT(20)"`
	InviteCount string `xorm:"not null default 0 comment('邀请人数')  INT(10)"`
	ActivityId  int    `xorm:"not null default 0 comment('活动id') INT(10)"`
	CreatedAt   string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt   string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel       string `xorm:"not null default 0 comment('是否被删除')  TINYINT(3)"`
}

type ActivityVotes201812Invitations struct {
	Id         string `xorm:"not null pk autoincr comment('自增主键')  INT(10)"`
	InviterUid string `xorm:"not null default 0 comment('邀请人UID') index  BIGINT(20)"`
	InviteeUid string `xorm:"not null default 0 comment('被邀请人UID') unique  BIGINT(20)"`
	Status     string `xorm:"not null default 0 comment('邀请状态0，邀请中；1，邀请成功')  TINYINT(2)"`
	ActivityId int    `xorm:"not null default 0 comment('活动id') INT(11)"`
	CreatedAt  string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt  string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel      string `xorm:"not null default 0 comment('是否被删除')  TINYINT(3)"`
}

type ActivityVotes201812Stars struct {
	Id        string `xorm:"not null pk autoincr comment('自增主键')  INT(10)"`
	Uid       string `xorm:"not null default 0 comment('明星微博ID') index  BIGINT(20)"`
	Sname     string `xorm:"not null default '' comment('显示名称') VARCHAR(100)"`
	Name      string `xorm:"not null default '' comment('真实名称') VARCHAR(100)"`
	Py        string `xorm:"not null default '' comment('拼音') VARCHAR(100)"`
	Vote      string `xorm:"not null default 0 comment('基础票数')  INT(10)"`
	Sort      string `xorm:"not null default 0 comment('排序')  INT(10)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	CreatedBy string `xorm:"not null default '' comment('创建者') VARCHAR(32)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(11)"`
	UpdatedBy string `xorm:"not null default '' comment('更新者') VARCHAR(32)"`
	IsDel     string `xorm:"not null default 0 comment('是否被删除')  TINYINT(3)"`
}

type AppVersion struct {
	Id          string `xorm:"not null pk autoincr  INT(10)"`
	DeviceId    string `xorm:"not null default '' comment('设备ID') unique VARCHAR(64)"`
	DeviceModel string `xorm:"not null default '' comment('设备型号') VARCHAR(64)"`
	Uid         string `xorm:"not null default 0 comment('用户微博ID') index  BIGINT(20)"`
	Os          string `xorm:"not null default 0 comment('系统信息') index  TINYINT(1)"`
	Version     string `xorm:"not null default '' comment('版本') index VARCHAR(32)"`
	CreatedAt   string `xorm:"not null default 0 comment('创建时间') index  INT(10)"`
	UpdatedAt   string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel       int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type Articles struct {
	Id               string `xorm:"not null pk autoincr comment('自增主键 ID')  INT(11)"`
	Title            string `xorm:"not null default '' comment('标题') VARCHAR(200)"`
	Summary          string `xorm:"not null default '' comment('摘要') VARCHAR(2000)"`
	Content          string `xorm:"not null comment('正文') MEDIUMTEXT"`
	Cover            string `xorm:"not null default '' comment('封面图片') VARCHAR(1000)"`
	UpvoteAmount     string `xorm:"not null default 0 comment('点赞数量')  INT(11)"`
	ShareAmount      string `xorm:"not null default 0 comment('分享数量')  INT(11)"`
	Type             int    `xorm:"not null default 1 comment('类型：1课堂，2奇葩说，3报告，4法规') TINYINT(3)"`
	Status           string `xorm:"not null default 0 comment('条文的状态，默认为草稿')  TINYINT(3)"`
	CommentId        string `xorm:"not null default '' comment('搜索内容的评论 ID') VARCHAR(32)"`
	CreatedAt        string `xorm:"not null default 0 comment('创建时间')  INT(11)"`
	CreatedBy        string `xorm:"not null default '' comment('创建者') VARCHAR(32)"`
	UpdatedAt        string `xorm:"not null default 0 comment('更新时间')  INT(11)"`
	UpdatedBy        string `xorm:"not null default '' comment('更新者') VARCHAR(32)"`
	FirstPublishedAt string `xorm:"not null default 0 comment('条目首次发布时间')  INT(11)"`
	IsDel            int    `xorm:"not null default 0 comment('是否被删除') TINYINT(3)"`
	IsHot            int    `xorm:"not null default 0 comment('是否热门') TINYINT(1)"`
	Media            string `xorm:"not null default '' comment('媒体') VARCHAR(128)"`
	IsPub            int    `xorm:"not null default 0 comment('是否公告0.否 1.是') TINYINT(1)"`
}

type AuthAssignment struct {
	Id         string    `xorm:"not null pk autoincr  INT(10)"`
	UserId     int       `xorm:"not null comment('用户表id') INT(10)"`
	ItemId     int       `xorm:"not null comment('角色 权限或者路由的id') INT(10)"`
	CreateTime time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
}

type AuthItem struct {
	Id         string    `xorm:"not null pk autoincr  INT(10)"`
	Name       string    `xorm:"not null comment('名称') VARCHAR(255)"`
	Type       int       `xorm:"not null default 1 comment('1.路由 2.权限 3.角色') TINYINT(255)"`
	Data       string    `xorm:"not null comment('备注') TEXT"`
	CreateTime time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
	Status     int       `xorm:"not null default 1 comment('状态 1.正常') TINYINT(1)"`
	IsDel      int       `xorm:"not null default 0 comment('是否删除 0 否 1 是') TINYINT(1)"`
	Location   string    `xorm:"not null default '' comment('地区') VARCHAR(255)"`
}

type AuthItemChild struct {
	Id         string    `xorm:"not null pk autoincr  INT(10)"`
	Parent     int       `xorm:"not null default 0 comment('父级id') INT(10)"`
	Child      int       `xorm:"not null default 0 comment('子id') INT(10)"`
	CreateTime time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
}

type AuthMenu struct {
	Id         string    `xorm:"not null pk autoincr  INT(10)"`
	Name       string    `xorm:"not null comment('菜单名称') VARCHAR(255)"`
	ParentId   int       `xorm:"not null default 0 comment('父级菜单id') INT(11)"`
	Router     string    `xorm:"not null comment('菜单对应路由') VARCHAR(255)"`
	Status     int       `xorm:"not null default 1 comment('状态 1.正常') TINYINT(1)"`
	CreateTime time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
	Data       string    `xorm:"not null comment('菜单样式') VARCHAR(500)"`
	IsDel      int       `xorm:"not null default 0 comment('是否删除 0否 1.是') TINYINT(1)"`
	Sort       int       `xorm:"not null default 100 comment('排序') INT(10)"`
}

type AuthUser struct {
	Id         string    `xorm:"not null pk autoincr  INT(10)"`
	Username   string    `xorm:"not null comment('用户名 一般为邮箱前缀') unique VARCHAR(255)"`
	Name       string    `xorm:"not null comment('姓名') VARCHAR(255)"`
	Phone      string    `xorm:"not null comment('联系方式') VARCHAR(30)"`
	IsDel      int       `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
	Remark     string    `xorm:"not null comment('备注') VARCHAR(255)"`
	CreateTime time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
}

type AutoBrands struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	ZhName    string `xorm:"not null default '' comment('中文名称') VARCHAR(128)"`
	EnName    string `xorm:"not null default '' comment('英文名称') VARCHAR(128)"`
	PyName    string `xorm:"not null default '' comment('拼音名称') VARCHAR(128)"`
	Logo      string `xorm:"not null default '' comment('logo') VARCHAR(256)"`
	AutoId    string `xorm:"not null default 0 comment('汽车数据库ID') unique  INT(10)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除') TINYINT(1)"`
}

type AutoCorps struct {
	Id          string `xorm:"not null pk autoincr  INT(10)"`
	Title       string `xorm:"not null default '' comment('名称') VARCHAR(128)"`
	BrandAutoid string `xorm:"not null default 0 comment('品牌ID') index  INT(10)"`
	AutoId      string `xorm:"not null default 0 comment('汽车数据库ID') unique  INT(10)"`
	CreatedAt   string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt   string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel       int    `xorm:"not null default 0 comment('是否删除') TINYINT(1)"`
}

type AutoModels struct {
	Id           string `xorm:"not null pk autoincr  INT(10)"`
	Title        string `xorm:"not null default '' comment('名称') VARCHAR(128)"`
	SerialAutoid string `xorm:"not null default 0 comment('车系ID') index  INT(10)"`
	AutoId       string `xorm:"not null default 0 comment('汽车数据库ID') unique  INT(10)"`
	CreatedAt    string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt    string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel        int    `xorm:"not null default 0 comment('是否删除') TINYINT(1)"`
}

type AutoSerials struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	Title      string `xorm:"not null default '' comment('名称') VARCHAR(128)"`
	CorpAutoid string `xorm:"not null default 0 comment('厂商ID') index  INT(10)"`
	AutoId     string `xorm:"not null default 0 comment('汽车数据库ID') unique  INT(10)"`
	CreatedAt  string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt  string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel      int    `xorm:"not null default 0 comment('是否删除') TINYINT(1)"`
}

type AutoreviewRecs struct {
	Id           string `xorm:"not null pk autoincr  INT(10)"`
	Sn           string `xorm:"not null default 0 comment('投诉单号') unique  BIGINT(18)"`
	ReviewStatus int    `xorm:"not null default 0 comment('审核状态 1通过 2驳回') TINYINT(1)"`
	VerifyStatus int    `xorm:"not null default 0 comment('复查状态') index(idx_ct_status_verifyby) TINYINT(1)"`
	VerifyAt     string `xorm:"not null default 0 comment('复核时间')  INT(10)"`
	VerifyBy     string `xorm:"not null default '' comment('复查人') index(idx_ct_status_verifyby) VARCHAR(32)"`
	RuleId       string `xorm:"not null default 0 comment('策略ID')  INT(10)"`
	CreatedAt    string `xorm:"not null default 0 comment('创建时间') index(idx_ct_status_verifyby)  INT(10)"`
	UpdatedAt    string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel        int    `xorm:"not null default 0 comment('是否删除') TINYINT(1)"`
}

type AutoreviewRules struct {
	Id              string `xorm:"not null pk autoincr  INT(10)"`
	Title           string `xorm:"not null default '' comment('名称') VARCHAR(256)"`
	ReviewOp        int    `xorm:"not null default 0 comment('审核类型 1通过 2驳回') TINYINT(1)"`
	Field1Id        string `xorm:"not null default 0 comment('一级领域id')  INT(10)"`
	Field2Id        string `xorm:"not null default 0 comment('二级领域id')  INT(10)"`
	CompanyType     int    `xorm:"not null default 0 comment('商家属性 1正常 2虚拟') TINYINT(1)"`
	Extralist       string `xorm:"not null default '' comment('商家uid') VARCHAR(1024)"`
	Exceplist       string `xorm:"not null default '' comment('商家uid') VARCHAR(1024)"`
	ReqImg          int    `xorm:"not null default 0 comment('上传图片') TINYINT(1)"`
	CheckCompRepeat int    `xorm:"not null default 0 comment('重复投诉') TINYINT(1)"`
	CreatedAt       string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	CreatedBy       string `xorm:"not null default '' comment('创建人') VARCHAR(32)"`
	UpdatedAt       string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	UpdatedBy       string `xorm:"not null default '' comment('更新人') VARCHAR(32)"`
	IsDel           int    `xorm:"not null default 0 comment('是否删除') TINYINT(1)"`
}

type Coappeals struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	Couid      string `xorm:"not null default 0 comment('商家微博ID') index(idx_couid_sn) index(idx_couid_status)  BIGINT(20)"`
	Sn         string `xorm:"not null default 0 comment('投诉单号') index(idx_couid_sn)  BIGINT(18)"`
	Title      string `xorm:"not null default '' comment('投诉标题') VARCHAR(128)"`
	Cotitle    string `xorm:"not null default '' comment('商家名称') VARCHAR(128)"`
	Uid        string `xorm:"not null default 0 comment('用户微博ID')  BIGINT(20)"`
	Type       int    `xorm:"not null default 0 comment('商家申诉类型：1重复投诉 2非本商户投诉') TINYINT(1)"`
	Status     int    `xorm:"not null default 0 comment('商家申诉状态：0未发起申诉 1申诉中 2申诉驳回 3申诉通过') index(idx_couid_status) TINYINT(1)"`
	InitNum    int    `xorm:"not null default 0 comment('申诉次数') TINYINT(1)"`
	CreateTime string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	CheckTime  string `xorm:"not null default 0 comment('审核时间')  INT(10)"`
	IsDel      int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type Companies struct {
	Id              int 	`xorm:"not null pk autoincr comment('自增主键')  INT(10)"`
	Level           int    	`xorm:"not null default 1 comment('公司账号级别，1是主账号，2是子账号，0不参与排行榜') TINYINT(4)"`
	Field           int    	`xorm:"not null default 0 comment('给商家添加的标签') TINYINT(4)"`
	Uid             string 	`xorm:"not null default 0 comment('对应的微博蓝 v 号的 uid') unique  BIGINT(20)"`
	PUid            string 	`xorm:"not null default 0 comment('父级uid') index  BIGINT(20)"`
	PTitle          string 	`xorm:"not null default '' comment('父级公司名') VARCHAR(128)"`
	Title           string 	`xorm:"not null default '' comment('公司名') VARCHAR(128)"`
	ShowTitle       string 	`xorm:"not null default '' comment('显示的公司名') VARCHAR(128)"`
	ValidAmount     int 	`xorm:"not null default 0 comment('有效投诉数量')  INT(10)"`
	RepliedAmount   int 	`xorm:"not null default 0 comment('已回复投诉数量')  INT(10)"`
	CompletedAmount int		`xorm:"not null default 0 comment('已完成投诉数量')  INT(10)"`
	RedSort         int 	`xorm:"not null default 0 comment('红榜排序') index(red_black_sort)  INT(10)"`
	BlackSort       int 	`xorm:"not null default 0 comment('黑榜排序') index(red_black_sort)  INT(10)"`
	Status          int    	`xorm:"not null default 0 comment('当前公司的入驻状态，默认为待审核状态') TINYINT(4)"`
	CreatedAt       int 	`xorm:"not null default 0 comment('创建时间，即公司申请入驻的时间或用户前台创建的时间')  INT(10)"`
	CheckedAt       int 	`xorm:"not null default 0 comment('管理员的审核时间，如果审核通过即为商家的入驻时间')  INT(10)"`
	CheckedBy       string `xorm:"not null default '' comment('审核人员') VARCHAR(32)"`
	Remark          string `xorm:"not null default '' comment('备注内容，用来放管理员审核不通过的原因等内容') VARCHAR(10000)"`
	IsDel           int    `xorm:"not null default 0 comment('是否被删除') TINYINT(4)"`
	Suggested       int    `xorm:"not null default 1 comment('是否下架：0下架，1正常') TINYINT(4)"`
	Email           string `xorm:"not null default '' comment('商家邮箱') VARCHAR(500)"`
	Bkrpt           int    `xorm:"not null default 0 comment('是否破产') TINYINT(4)"`
	BkrptInfo       string `xorm:"not null default '' comment('破产信息') VARCHAR(5000)"`
	EvalAmt         int    `xorm:"not null default 0 comment('用户评价数量')  INT(10)"`
	EvalAtt         int    `xorm:"not null default 0 comment('用户评价服务态度')  INT(10)"`
	EvalPro         int    `xorm:"not null default 0 comment('用户评价处理速度')  INT(10)"`
	EvalSat         int    `xorm:"not null default 0 comment('用户评价满意度')  INT(10)"`
	Phone           string `xorm:"not null default '' comment('电话') VARCHAR(500)"`
	IsPzxfmember    int    `xorm:"not null default 0 comment('是否为品质消费领导者组织成员：0否 1是') TINYINT(1)"`
	HasFld          int    `xorm:"not null default 0 comment('商家是否已分类，默认0否，1是') TINYINT(1)"`
	IsVirtual       int    `xorm:"not null default 0 comment('虚拟商家标识：0否，1是') INT(1)"`
	IsAnchor        int    `xorm:"not null default 0 comment('主播商家标识，0:否，1:是')  TINYINT(2)"`
	HmEmail         string `xorm:"not null default '' comment('黑猫设置邮箱') VARCHAR(500)"`
	AgreeSign       int    `xorm:"not null default 0 comment('签署协议状态') TINYINT(1)"`
	AgreeSignAt     int 	`xorm:"not null default 0 comment('签署协议时间')  INT(10)"`
	Valid30d        int 	`xorm:"not null default 0 comment('30天有效量')  INT(10)"`
	Reply30d        int 	`xorm:"not null default 0 comment('30天回复量')  INT(10)"`
	Complete30d     int 	`xorm:"not null default 0 comment('30天完成量')  INT(10)"`
}

type CompaniesPro struct {
	Id              string `xorm:"not null pk autoincr comment('自增主键')  INT(10)"`
	Level           int    `xorm:"not null default 1 comment('公司账号级别，1是主账号，2是子账号，0不参与排行榜') TINYINT(4)"`
	Field           int    `xorm:"not null default 0 comment('给商家添加的标签') TINYINT(4)"`
	Uid             string `xorm:"not null default 0 comment('对应的微博蓝 v 号的 uid') unique  BIGINT(20)"`
	PUid            string `xorm:"not null default 0 comment('父级uid') index  BIGINT(20)"`
	PTitle          string `xorm:"not null default '' comment('父级公司名') VARCHAR(128)"`
	Title           string `xorm:"not null default '' comment('公司名') VARCHAR(128)"`
	ShowTitle       string `xorm:"not null default '' comment('显示的公司名') VARCHAR(128)"`
	ValidAmount     string `xorm:"not null default 0 comment('有效投诉数量')  INT(10)"`
	RepliedAmount   string `xorm:"not null default 0 comment('已回复投诉数量')  INT(10)"`
	CompletedAmount string `xorm:"not null default 0 comment('已完成投诉数量')  INT(10)"`
	RedSort         string `xorm:"not null default 0 comment('红榜排序') index(red_black_sort)  INT(10)"`
	BlackSort       string `xorm:"not null default 0 comment('黑榜排序') index(red_black_sort)  INT(10)"`
	Status          int    `xorm:"not null default 0 comment('当前公司的入驻状态，默认为待审核状态') TINYINT(4)"`
	CreatedAt       string `xorm:"not null default 0 comment('创建时间，即公司申请入驻的时间或用户前台创建的时间')  INT(10)"`
	CheckedAt       string `xorm:"not null default 0 comment('管理员的审核时间，如果审核通过即为商家的入驻时间')  INT(10)"`
	CheckedBy       string `xorm:"not null default '' comment('审核人员') VARCHAR(32)"`
	Remark          string `xorm:"not null default '' comment('备注内容，用来放管理员审核不通过的原因等内容') VARCHAR(10000)"`
	IsDel           int    `xorm:"not null default 0 comment('是否被删除') TINYINT(4)"`
	Suggested       int    `xorm:"not null default 1 comment('是否下架：0下架，1正常') TINYINT(4)"`
	Email           string `xorm:"not null default '' comment('商家邮箱') VARCHAR(500)"`
	Bkrpt           int    `xorm:"not null default 0 comment('是否破产') TINYINT(4)"`
	BkrptInfo       string `xorm:"not null default '' comment('破产信息') VARCHAR(5000)"`
	EvalAmt         string `xorm:"not null default 0 comment('用户评价数量')  INT(10)"`
	EvalAtt         string `xorm:"not null default 0 comment('用户评价服务态度')  INT(10)"`
	EvalPro         string `xorm:"not null default 0 comment('用户评价处理速度')  INT(10)"`
	EvalSat         string `xorm:"not null default 0 comment('用户评价满意度')  INT(10)"`
	Phone           string `xorm:"not null default '' comment('电话') VARCHAR(500)"`
	IsPzxfmember    int    `xorm:"not null default 0 comment('是否为品质消费领导者组织成员：0否 1是') TINYINT(1)"`
	HasFld          int    `xorm:"not null default 0 comment('商家是否已分类，默认0否，1是') TINYINT(1)"`
	IsVirtual       int    `xorm:"not null default 0 comment('虚拟商家标识：0否，1是') INT(1)"`
	IsAnchor        string `xorm:"not null default 0 comment('主播商家标识，0:否，1:是')  TINYINT(2)"`
	HmEmail         string `xorm:"not null default '' comment('黑猫设置邮箱') VARCHAR(500)"`
}

type CompanyApplySettle struct {
	Id          string `xorm:"not null pk autoincr  INT(10)"`
	Uid         string `xorm:"not null default 0 comment('申请人微博ID')  BIGINT(20)"`
	Title       string `xorm:"not null default '' comment('申请人名称') VARCHAR(128)"`
	Vid         string `xorm:"not null default 0 comment('虚拟ID')  BIGINT(20)"`
	Vtitle      string `xorm:"not null default '' comment('虚拟商家名称') VARCHAR(128)"`
	Status      int    `xorm:"not null default 0 comment('状态：0待审核，1成功，2失败') TINYINT(1)"`
	ContactInfo string `xorm:"not null default '' comment('商家联系方式') VARCHAR(512)"`
	CreateTime  string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime  string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel       int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type CompanyField struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Level     int    `xorm:"not null default 1 comment('级数') TINYINT(4)"`
	Parent    string `xorm:"not null default 0 comment('父级id')  INT(10)"`
	Title     string `xorm:"not null default '' comment('名称') VARCHAR(128)"`
	Appeal    string `xorm:"not null default '' comment('诉求') VARCHAR(1024)"`
	Issue     string `xorm:"not null default '' comment('问题') VARCHAR(1024)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type CompanyFieldMonthStast struct {
	Id          string `xorm:"not null pk autoincr  INT(10)"`
	FieldId     string `xorm:"not null default 0 comment('行业ID') index(idx_field_date)  INT(10)"`
	FieldTitle  string `xorm:"not null default '' comment('行业名称') VARCHAR(128)"`
	Level       int    `xorm:"not null default 1 comment('行业等级') TINYINT(1)"`
	Date        string `xorm:"not null default 0 comment('日期') index(idx_field_date)  INT(10)"`
	Valid       string `xorm:"not null default 0 comment('有效量')  INT(10)"`
	Reply       string `xorm:"not null default 0 comment('回复量')  INT(10)"`
	Complete    string `xorm:"not null default 0 comment('完成量')  INT(10)"`
	ReplyPer    string `xorm:"not null default '' comment('回复率') VARCHAR(16)"`
	CompletePer string `xorm:"not null default '' comment('完成率') VARCHAR(16)"`
	Replytime   string `xorm:"not null default '' comment('响应时间	') VARCHAR(16)"`
	Eval        string `xorm:"not null default '' comment('评价均值') VARCHAR(16)"`
	AvgValid    string `xorm:"not null default 0 comment('有效量均值')  INT(10)"`
	AvgReply    string `xorm:"not null default 0 comment('回复量均值')  INT(10)"`
	AvgComplete string `xorm:"not null default 0 comment('完成量均值')  INT(10)"`
	Companies   string `xorm:"not null default '' comment('商家数据分布') VARCHAR(1024)"`
	CreateTime  string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime  string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel       int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type CompanyFieldMonthUserportrait struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	FieldId    string `xorm:"not null default 0 comment('行业ID') index(idx_field_date)  INT(10)"`
	FieldTitle string `xorm:"not null default '' comment('行业名称') VARCHAR(128)"`
	Date       string `xorm:"not null default 0 comment('日期') index(idx_field_date)  INT(10)"`
	XztaskId   string `xorm:"not null default '' comment('寻辙分析任务ID') VARCHAR(64)"`
	XzStatus   string `xorm:"not null default 0 comment('寻辙任务状态：0处理中，1已完成')  INT(1)"`
	Result     string `xorm:"not null default '' comment('寻辙返回数据') VARCHAR(4096)"`
	Age        string `xorm:"not null default '' comment('年龄分布') VARCHAR(1024)"`
	Gender     string `xorm:"not null default '' comment('性别分布') VARCHAR(1024)"`
	Province   string `xorm:"not null default '' comment('省份分布') VARCHAR(1024)"`
	CreateTime string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel      int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type CompanyFieldPro struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Level     int    `xorm:"not null default 1 comment('级数') TINYINT(4)"`
	Parent    string `xorm:"not null default 0 comment('父级id')  INT(10)"`
	Title     string `xorm:"not null default '' comment('名称') VARCHAR(128)"`
	Appeal    string `xorm:"not null default '' comment('诉求') VARCHAR(1024)"`
	Issue     string `xorm:"not null default '' comment('问题') VARCHAR(1024)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type CompanyFieldRel struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	CompanyUid string `xorm:"not null default 0 comment('商家uid') index unique(idx_uid_f1_f2)  BIGINT(20)"`
	Field1Id   string `xorm:"not null default 0 comment('一级领域id') index(idx_fld) unique(idx_uid_f1_f2)  INT(10)"`
	Field2Id   string `xorm:"not null default 0 comment('二级领域id') index(idx_fld) unique(idx_uid_f1_f2)  INT(10)"`
	CreatedAt  string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type CompanyFieldRelPro struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	CompanyUid string `xorm:"not null default 0 comment('商家uid') index  BIGINT(20)"`
	Field1Id   string `xorm:"not null default 0 comment('一级领域id') index(idx_fld)  INT(10)"`
	Field2Id   string `xorm:"not null default 0 comment('二级领域id') index(idx_fld)  INT(10)"`
	CreatedAt  string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type CompanyProducts struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Uid       string `xorm:"not null default 0 comment('微博UID') unique(idx_uid_type_pid_upt)  BIGINT(20)"`
	Type      int    `xorm:"not null default 0 comment('类型') unique(idx_uid_type_pid_upt) TINYINT(1)"`
	ProductId string `xorm:"not null default 0 comment('商品ID') unique(idx_uid_type_pid_upt)  INT(10)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	CreatedBy string `xorm:"not null default '0' comment('创建人') VARCHAR(32)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间') unique(idx_uid_type_pid_upt)  INT(10)"`
	UpdatedBy string `xorm:"not null default '0' comment('更新人') VARCHAR(32)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除') TINYINT(1)"`
}

type CompanySuggestedLogs struct {
	Id        string `xorm:"not null pk autoincr comment('主键')  INT(11)"`
	Uid       string `xorm:"not null default 0 comment('uid') index  BIGINT(20)"`
	Type      string `xorm:"not null default 0 comment('操作类型，0是下架操作，1是上架操作')  TINYINT(2)"`
	Reason    string `xorm:"not null default '' comment('操作原因') VARCHAR(255)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(11)"`
	CreatedBy string `xorm:"not null default '' comment('创建人') VARCHAR(32)"`
	IsDel     string `xorm:"not null default 0 comment('是否删除')  TINYINT(2)"`
}

type CompanyVirtualReplace struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	Vid        string `xorm:"not null default 0 comment('虚拟ID')  BIGINT(20)"`
	Vtitle     string `xorm:"not null default '' comment('虚拟名称') VARCHAR(128)"`
	Uid        string `xorm:"not null default 0 comment('微博ID')  BIGINT(20)"`
	Title      string `xorm:"not null default '' comment('名称') VARCHAR(128)"`
	Status     int    `xorm:"not null default 0 comment('状态：0进行中，1成功，2失败') TINYINT(1)"`
	OpBy       string `xorm:"not null default '' comment('执行人') VARCHAR(32)"`
	CreateTime string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel      int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type CompetitorDayStast struct {
	Id              string `xorm:"not null pk autoincr  INT(10)"`
	Couid           string `xorm:"not null default 0 comment('商家微博ID') index(idx_couid_date)  BIGINT(18)"`
	Date            string `xorm:"not null default 0 comment('日期') index(idx_couid_date)  INT(10)"`
	DValid          string `xorm:"not null default 0 comment('当日有效量')  INT(10)"`
	Valid           string `xorm:"not null default 0 comment('月累计有效量')  INT(10)"`
	Reply           string `xorm:"not null default 0 comment('月累计回复量')  INT(10)"`
	Complete        string `xorm:"not null default 0 comment('月累计完成量')  INT(10)"`
	ReplyPer        string `xorm:"not null default '' comment('月累计回复率') VARCHAR(16)"`
	CompletePer     string `xorm:"not null default '' comment('月累计完成率') VARCHAR(16)"`
	Replytime       string `xorm:"not null default '' comment('月累计响应时间	') VARCHAR(16)"`
	Eval            string `xorm:"not null default '' comment('月累计评价均值') VARCHAR(16)"`
	UserComplete    string `xorm:"not null default 0 comment('月累计用户完成量')  INT(10)"`
	UserCompletePer string `xorm:"not null default '' comment('月累计用户完成率') VARCHAR(16)"`
	CreateTime      string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime      string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel           int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type CompetitorFieldDayStast struct {
	Id          string `xorm:"not null pk autoincr  INT(10)"`
	FieldId     string `xorm:"not null default 0 comment('行业ID') index(idx_field_date)  INT(10)"`
	FieldTitle  string `xorm:"not null default '' comment('行业名称') VARCHAR(128)"`
	Date        string `xorm:"not null default 0 comment('日期') index(idx_field_date)  INT(10)"`
	DValid      string `xorm:"not null default 0 comment('当日有效量均值')  INT(10)"`
	Valid       string `xorm:"not null default 0 comment('月累计有效量')  INT(10)"`
	Reply       string `xorm:"not null default 0 comment('月累计回复量')  INT(10)"`
	Complete    string `xorm:"not null default 0 comment('月累计完成量')  INT(10)"`
	ReplyPer    string `xorm:"not null default '' comment('月累计回复率') VARCHAR(16)"`
	CompletePer string `xorm:"not null default '' comment('月累计完成率') VARCHAR(16)"`
	Replytime   string `xorm:"not null default '' comment('月累计响应时间	') VARCHAR(16)"`
	Eval        string `xorm:"not null default '' comment('月累计评价均值') VARCHAR(16)"`
	CreateTime  string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime  string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel       int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type CompetitorHistoryStast struct {
	Id              string `xorm:"not null pk autoincr  INT(10)"`
	Couid           string `xorm:"not null default 0 comment('商家微博ID') index  BIGINT(18)"`
	Valid           string `xorm:"not null default 0 comment('历史有效量')  INT(10)"`
	Reply           string `xorm:"not null default 0 comment('历史回复量')  INT(10)"`
	Complete        string `xorm:"not null default 0 comment('历史完成量')  INT(10)"`
	ReplyPer        string `xorm:"not null default '' comment('历史回复率') VARCHAR(16)"`
	CompletePer     string `xorm:"not null default '' comment('历史完成率') VARCHAR(16)"`
	Replytime       string `xorm:"not null default '' comment('历史响应时间	') VARCHAR(16)"`
	Eval            string `xorm:"not null default '' comment('历史评价均值') VARCHAR(16)"`
	UserComplete    string `xorm:"not null default 0 comment('月累计用户完成量')  INT(10)"`
	UserCompletePer string `xorm:"not null default '' comment('月累计用户完成率') VARCHAR(16)"`
	CreateTime      string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime      string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel           int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type ComplaintAuto struct {
	Id           string    `xorm:"not null pk autoincr  INT(10)"`
	Sn           string    `xorm:"not null default 0 comment('投诉单号') unique  BIGINT(18)"`
	Brand        string    `xorm:"not null default '' comment('品牌') VARCHAR(128)"`
	BrandAutoid  string    `xorm:"not null default 0 comment('品牌ID')  INT(10)"`
	Corp         string    `xorm:"not null default '' comment('厂商') VARCHAR(128)"`
	CorpAutoid   string    `xorm:"not null default 0 comment('厂商ID')  INT(10)"`
	Serial       string    `xorm:"not null default '' comment('车系') VARCHAR(128)"`
	SerialAutoid string    `xorm:"not null default 0 comment('车系ID')  INT(10)"`
	Model        string    `xorm:"not null default '' comment('车型') VARCHAR(128)"`
	ModelAutoid  string    `xorm:"not null default 0 comment('车型ID')  INT(10)"`
	Vin          string    `xorm:"not null default '' comment('车辆识别代码') VARCHAR(32)"`
	CarNumber    string    `xorm:"not null default '' comment('车牌号') VARCHAR(16)"`
	PurchaseTime time.Time `xorm:"not null default '1970-01-01 00:00:00' comment('购买时间') DATETIME"`
	Province     string    `xorm:"not null default '' comment('省份') VARCHAR(64)"`
	City         string    `xorm:"not null default '' comment('城市') VARCHAR(64)"`
	Dealer       string    `xorm:"not null default '' comment('经销商') VARCHAR(128)"`
	ServiceIssue string    `xorm:"not null default '' comment('服务问题') VARCHAR(512)"`
	QualityIssue string    `xorm:"not null default '' comment('质量问题') VARCHAR(512)"`
	CreatedAt    string    `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt    string    `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel        int       `xorm:"not null default 0 comment('是否删除') TINYINT(1)"`
}

type ComplaintEvals struct {
	Id           string `xorm:"not null pk autoincr  INT(10)"`
	Sn           string `xorm:"not null default 0 comment('投诉单号')  BIGINT(18)"`
	Uid          string `xorm:"not null default 0 comment('用户微博ID')  BIGINT(20)"`
	Couid        string `xorm:"not null default 0 comment('商家微博ID')  BIGINT(20)"`
	Content      string `xorm:"not null default '' comment('内容') VARCHAR(4096)"`
	Attitude     int    `xorm:"not null default 0 comment('服务态度') index TINYINT(1)"`
	Process      int    `xorm:"not null default 0 comment('处理速度') index TINYINT(1)"`
	Satisfaction int    `xorm:"not null default 0 comment('满意度') index TINYINT(1)"`
	AvgScore     int    `xorm:"not null default 0 comment('平均分') index TINYINT(1)"`
	Suggested    int    `xorm:"not null default 0 comment('是否推荐') index TINYINT(3)"`
	EvalTime     string `xorm:"not null default 0 comment('评价时间') index  INT(10)"`
	CreateTime   string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime   string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel        int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type Complaints struct {
	Id               int	 `xorm:"not null pk autoincr comment('自增主键')"`
	Sn               int	 `xorm:"not null comment('投诉单号 ') unique"`
	Title            string `xorm:"not null comment('投诉内容的标题') VARCHAR(128)"`
	Uid              int `xorm:"not null default 0 comment('用户的 id') index(idx_uid_co_id)"`
	Couid            int `xorm:"not null default 0 comment('对应的公司的 uid, 若为0即表示投诉对象没有入住') index(idx_assigntime_couid) index(idx_co_st_del_test_ct) index(idx_co_st_exp_del_test) index(idx_couid_cocompletet_cocompletes) index(idx_ct_couid) index(idx_evt_couid) index(idx_mt_couid) index(idx_uid_co_id)"`
	Cotitle          string `xorm:"not null default '' comment('对应公司的名称') VARCHAR(128)"`
	Phone            string `xorm:"comment('电话') VARCHAR(200)"`
	Appeal           string `xorm:"not null default '' comment('诉求') VARCHAR(256)"`
	Content          string `xorm:"not null default '' comment('投诉的 id') VARCHAR(2048)"`
	Issue            string `xorm:"not null default '' comment('问题') VARCHAR(1024)"`
	CommentId        string `xorm:"not null default '' comment('搜索内容的评论 ID') index VARCHAR(32)"`
	AppId            string `xorm:"not null default '' comment('此单推到 cms 的 app.id') VARCHAR(32)"`
	UpvoteAmount     int `xorm:"not null default 0 comment('点赞的数量')  INT(10)"`
	Anonymous        int `xorm:"not null default 0 comment('是否匿名投诉')  TINYINT(3)"`
	Attach           string `xorm:"comment('附件') MEDIUMTEXT"`
	HideAttach       int `xorm:"not null default 0 comment('是否隐藏投诉附件 默认不隐藏')"`
	Privacy          string `xorm:"not null default '' comment('隐私内容') VARCHAR(2048)"`
	Cost             string `xorm:"not null default '' comment(' 涉诉金额') VARCHAR(32)"`
	EvaluateU        string `xorm:"not null default '' comment('投诉者评价') VARCHAR(2048)"`
	EvaluateC        string `xorm:"not null comment('投诉者评价') VARCHAR(128)"`
	Status           int    `xorm:"not null default 0 index(idx_co_st_del_test_ct) index(idx_co_st_exp_del_test) index(idx_expat_status) index(idx_field_expt_status) index(idx_prov_expt_status) index index(idx_status_exp_test_del) index(idx_status_hassupp_replyat) index(idx_status_msgb_assignat) index(idx_status_test_ct) TINYINT(4)"`
	Field            int `xorm:"not null default 0 comment('领域') index index(idx_field_expt_status)"`
	Source           int `xorm:"not null default 0 comment('投诉来源')"`
	AssignedAt       int `xorm:"not null default 0 comment('分配时间') index(idx_assigntime_couid) index(idx_status_msgb_assignat)"`
	RepliedAt        int `xorm:"not null default 0 comment('回复时间') index(idx_status_hassupp_replyat)"`
	CreatedAt        int `xorm:"not null default 0 comment('创建时间') index(idx_co_st_del_test_ct) index(idx_ct_couid) index(idx_status_test_ct)"`
	UpdatedAt        int `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	CheckedAt        int `xorm:"not null default 0 comment('审核时间') index  INT(10)"`
	CheckedBy        string `xorm:"not null default '' comment('审核者') VARCHAR(32)"`
	EditedAt         int `xorm:"not null default 0 comment('审核通过前修改投诉对象时间')  INT(10)"`
	EditedBy         string `xorm:"not null default '' comment('审核通过前修改投诉对象的管理员') VARCHAR(32)"`
	LastModifiedAt   int `xorm:"not null default 0 comment('用户最后修改投诉时间') index(idx_mt_couid)  INT(10)"`
	IsTest           int `xorm:"not null default 0 comment('是否为测试内容') index(idx_co_st_del_test_ct) index(idx_co_st_exp_del_test) index(idx_status_exp_test_del) index(idx_status_test_ct)  TINYINT(3)"`
	IsDel            int `xorm:"not null default 0 index(idx_co_st_del_test_ct) index(idx_co_st_exp_del_test) index(idx_status_exp_test_del)  TINYINT(3)"`
	MsgBReply        int `xorm:"not null default 0 comment('通知商家回复消息的发送次数') index(idx_status_msgb_assignat)  TINYINT(3)"`
	MsgCAction       int `xorm:"not null default 0 comment('商家回复后通知用户操作消息的发送次数')  TINYINT(3)"`
	Note             string `xorm:"not null default '' comment('运营添加的注释信息') VARCHAR(256)"`
	ShowPhone        int    `xorm:"not null default 0 comment('是否在商家端显示手机号') TINYINT(1)"`
	Province         string `xorm:"not null default '' comment('省份') index(idx_prov_expt_status) VARCHAR(20)"`
	City             string `xorm:"not null default '' comment('城市') VARCHAR(20)"`
	LocalStation     string `xorm:"not null default '' comment('地方站') index VARCHAR(20)"`
	OrigCotitle      string `xorm:"not null default '' comment('原投诉对象名称') VARCHAR(128)"`
	CmsId            string `xorm:"not null default '' comment('CMS返回的ID') VARCHAR(32)"`
	ExtSrc           int `xorm:"not null default 0 comment('外部投诉来源')  TINYINT(3)"`
	Exposed          int `xorm:"not null default 1 comment('曝光状态：0未曝光，1已曝光') index(idx_co_st_exp_del_test) index(idx_status_exp_test_del)  TINYINT(1)"`
	ExposedAt        int `xorm:"not null default 0 comment('曝光时间') index(idx_expat_status) index(idx_field_expt_status) index(idx_prov_expt_status)  INT(10)"`
	HasSupp          int `xorm:"not null default 0 comment('商家回复后是否有新的补充投诉') index(idx_status_hassupp_replyat)  TINYINT(1)"`
	Gsn              int `xorm:"not null default 0 comment('集体投诉编号') index  BIGINT(18)"`
	Danger           int `xorm:"not null default 0 comment('高危投诉：0非高危 1低危 2中危 3高危')  TINYINT(3)"`
	Pltdgwords       string `xorm:"not null default '' comment('匹配的平台高危敏感词') VARCHAR(1024)"`
	Codgwords        string `xorm:"not null default '' comment('匹配的商家高危敏感词') VARCHAR(1024)"`
	MdaTag           string `xorm:"not null default '' comment('商家标签') VARCHAR(512)"`
	MdaNote          string `xorm:"not null default '' comment('商家备注') VARCHAR(512)"`
	Replyuid         int `xorm:"not null default 0 comment('最近一次回复投诉的商家/客服微博ID')  BIGINT(20)"`
	CoappealStatus   int    `xorm:"not null default 0 comment('商家申诉状态：0未发起申诉 1申诉中 2申诉驳回 3申诉通过') TINYINT(1)"`
	CoappealAt       int `xorm:"not null default 0 comment('商家申诉时间')  INT(10)"`
	Autoreview       int    `xorm:"not null default 0 comment('自动审核：1驳回') index TINYINT(1)"`
	CocompleteStatus int    `xorm:"not null default 0 comment('商家申诉结案：0未发起结案 1结案中 2用户拒绝结案 3用户同意结案 4平台自动结案') index index(idx_couid_cocompletet_cocompletes) TINYINT(1)"`
	CocompleteReason int    `xorm:"not null default 0 comment('商家结案申请原因：0未发起结案 1与用户达成一致 2联系不上用户 3最终解决方案') TINYINT(1)"`
	CocompleteNum    int    `xorm:"not null default 0 comment('商家申诉结案次数') TINYINT(1)"`
	Autoreply        int    `xorm:"not null default 0 comment('自动回复状态：0未自动回复 1自动回复 2自动回复后人工回复') TINYINT(1)"`
	CocompleteAt     int `xorm:"not null default 0 comment('结案申请时间') index(idx_couid_cocompletet_cocompletes)  INT(10)"`
	CompleteAt       int `xorm:"not null default 0 comment('完成时间')  INT(10)"`
	SuppAt           int `xorm:"not null default 0 comment('补充投诉时间')  INT(10)"`
	SyncMedia        int    `xorm:"not null default 0 comment('是否可同步本人信息到媒体0:否，1:同意') TINYINT(2)"`
	SyncWb           int    `xorm:"not null default 0 comment('是否可同步投诉内容到微博 0:否，1:同意') TINYINT(2)"`
	EvalAt           int `xorm:"not null default 0 comment('评价投诉时间') index(idx_evt_couid)  INT(10)"`
	WaterarmyType    int `xorm:"not null default 0 comment('疑似水军类型，0不是水军，1是一般水军，2是重度水军')  TINYINT(2)"`
	ShareImg         string `xorm:"not null default '' comment('分享图片地址') index VARCHAR(128)"`
	VirtualCompany   int `xorm:"not null default 0 comment('虚拟商家')  TINYINT(1)"`
	TopicId          int `xorm:"not null default 0 comment('话题ID')  INT(10)"`
	CocompleteIll    int    `xorm:"not null default 0 comment('违规结案原因') TINYINT(1)"`
	Tpl              int    `xorm:"not null default 0 comment('投诉模版类型') TINYINT(1)"`
	Reviewfail       int    `xorm:"not null default 0 comment('驳回理由') TINYINT(1)"`
}

type DataBlackCatRecords struct {
	Id             string `xorm:"not null pk comment('主键') VARCHAR(64)"`
	ComplainTarget string `xorm:"comment('投诉对象') VARCHAR(255)"`
	ComplainTimes  int    `xorm:"default 0 comment('投诉次数') INT(6)"`
	LoanFlag       string `xorm:"comment('该软件是否有投诉明确不放款
非明确不放款：0
明确不放款：1') VARCHAR(8)"`
	ConfirmLoanFlag string `xorm:"comment('即投诉内容中不包含催收和利率
不包含催收和利率：1
包含催收或利率：0') VARCHAR(8)"`
	RecentComplainContent string    `xorm:"comment('最近投诉前5条内容') TEXT"`
	Memo                  string    `xorm:"comment('备注') TEXT"`
	CreateTime            time.Time `xorm:"comment('创建时间') DATETIME"`
	KetechFlag            string    `xorm:"comment('判定催收和利率（0-不包含，1-包含）') VARCHAR(8)"`
	OnlineLoanFlag        string    `xorm:"comment('网贷不放款 1 网贷不放款 0 非网贷不放款') VARCHAR(8)"`
}

type ExtArticles struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	ExtId      string `xorm:"not null default '' comment('外部ID') index VARCHAR(256)"`
	Title      string `xorm:"not null default '' comment('标题') VARCHAR(1024)"`
	Content    string `xorm:"not null comment('正文') MEDIUMTEXT"`
	Cover      string `xorm:"not null default '' comment('封面图片') VARCHAR(1024)"`
	Ctime      string `xorm:"not null default 0 comment('文章创建时间')  INT(10)"`
	Media      string `xorm:"not null default '' comment('媒体') VARCHAR(128)"`
	Source     string `xorm:"not null default 0 comment('文章来源')  TINYINT(3)"`
	ArtId      string `xorm:"not null default 0 comment('黑猫报道ID')  INT(10)"`
	ImportedAt string `xorm:"not null default 0 comment('导入时间')  INT(10)"`
	ImportedBy string `xorm:"not null default '' comment('操作人') VARCHAR(32)"`
	CreatedAt  string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	CreatedBy  string `xorm:"not null default '' comment('创建者') VARCHAR(32)"`
	UpdatedAt  string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	UpdatedBy  string `xorm:"not null default '' comment('更新者') VARCHAR(32)"`
	IsDel      int    `xorm:"not null default 0 comment('是否删除') TINYINT(1)"`
}

type ExtComplaints struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	ExtSn      string `xorm:"not null default '' comment('外部编号') VARCHAR(64)"`
	Title      string `xorm:"not null default '' comment('投诉标题') VARCHAR(128)"`
	Uid        string `xorm:"not null default 0 comment('用户微博id')  BIGINT(20)"`
	Couid      string `xorm:"not null default 0 comment('商家微博id')  BIGINT(20)"`
	Cotitle    string `xorm:"not null default '' comment('商家名称') VARCHAR(128)"`
	Content    string `xorm:"not null default '' comment('投诉内容') VARCHAR(2048)"`
	Procedures string `xorm:"not null comment('投诉流程') MEDIUMTEXT"`
	Attach     string `xorm:"not null comment('附件') MEDIUMTEXT"`
	EvaluateU  string `xorm:"not null default '' comment('投诉人评价') VARCHAR(2048)"`
	Province   string `xorm:"not null default '' comment('省份') VARCHAR(20)"`
	City       string `xorm:"not null default '' comment('城市') VARCHAR(20)"`
	Status     int    `xorm:"not null default 0 comment('投诉状态') TINYINT(4)"`
	Source     string `xorm:"not null default 0 comment('投诉来源')  TINYINT(3)"`
	Sn         string `xorm:"not null default 0 comment('平台编号')  BIGINT(18)"`
	CreatedAt  string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt  string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel      string `xorm:"not null default 0 comment('是否被删除')  TINYINT(3)"`
	OpCode     string `xorm:"not null default 0 comment('操作类型：1创建 2更新 3删除')  TINYINT(3)"`
	Synced     string `xorm:"not null default 0 comment('是否已同步')  TINYINT(3)"`
}

type Feedback struct {
	Id              string `xorm:"not null pk autoincr  INT(10)"`
	Phone           string `xorm:"not null default '' comment('电话') VARCHAR(128)"`
	Type            int    `xorm:"not null default 0 comment('类型') index TINYINT(1)"`
	Content         string `xorm:"not null default '' comment('内容') VARCHAR(1024)"`
	Attach          string `xorm:"not null default '' comment('附件') VARCHAR(1024)"`
	UserDeviceModel string `xorm:"not null default '' comment('用户设备型号') VARCHAR(64)"`
	Version         string `xorm:"not null default '' comment('版本号') VARCHAR(32)"`
	Channel         int    `xorm:"not null default 0 comment('反馈渠道') index TINYINT(1)"`
	Sn              string `xorm:"not null default 0 comment('投诉单号 ')  BIGINT(18)"`
	DeviceId        string `xorm:"not null default '' comment('设备ID') VARCHAR(64)"`
	DeviceModel     string `xorm:"not null default '' comment('设备型号') VARCHAR(64)"`
	Province        string `xorm:"not null default '' comment('省份') VARCHAR(16)"`
	City            string `xorm:"not null default '' comment('城市') VARCHAR(16)"`
	From            int    `xorm:"not null default 0 comment('反馈入口') index TINYINT(1)"`
	CreatedAt       string `xorm:"not null default 0 comment('创建时间') index  INT(10)"`
	UpdatedAt       string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel           int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type GrpComps struct {
	Id           string `xorm:"not null pk autoincr  INT(10)"`
	Gsn          string `xorm:"not null default 0 comment('集体投诉编号') unique  BIGINT(18)"`
	Status       int    `xorm:"not null default 0 comment('状态') TINYINT(4)"`
	Uid          string `xorm:"not null default 0 comment('投诉人微博ID')  BIGINT(20)"`
	Couid        string `xorm:"not null default 0 comment('商家微博ID') index  BIGINT(20)"`
	Title        string `xorm:"not null default '' comment('标题') VARCHAR(256)"`
	Cotitle      string `xorm:"not null default '' comment('商家名称') VARCHAR(128)"`
	Content      string `xorm:"not null default '' comment('详情') VARCHAR(2048)"`
	Appeal       string `xorm:"not null default '' comment('诉求') VARCHAR(256)"`
	Issue        string `xorm:"not null default '' comment('问题') VARCHAR(1024)"`
	Cost         string `xorm:"not null default '' comment(' 涉诉金额') VARCHAR(32)"`
	Source       string `xorm:"not null default 0 comment('投诉来源')  TINYINT(3)"`
	Province     string `xorm:"not null default '' comment('省份') VARCHAR(20)"`
	City         string `xorm:"not null default '' comment('城市') VARCHAR(20)"`
	LocalStation string `xorm:"not null default '' comment('地方站') VARCHAR(20)"`
	Anonymous    string `xorm:"not null default 0 comment('是否匿名')  TINYINT(3)"`
	CmntId       string `xorm:"not null default '' comment('评论ID') VARCHAR(32)"`
	CompAmt      string `xorm:"not null default 0 comment('参与投诉数量') index  INT(10)"`
	ReplyAmt     string `xorm:"not null default 0 comment('已回复投诉数量')  INT(10)"`
	CompleteAmt  string `xorm:"not null default 0 comment('已完成投诉数量')  INT(10)"`
	UpvoteAmount string `xorm:"not null default 0 comment('点赞数量')  INT(10)"`
	ShareAmt     string `xorm:"not null default 0 comment('分享数量')  INT(10)"`
	ShareImg     string `xorm:"not null default '' comment('分享图片地址') VARCHAR(1024)"`
	HasSupp      string `xorm:"not null default 0 comment('商家回复后是否有新的补充投诉')  TINYINT(1)"`
	Replied      string `xorm:"not null default 0 comment('商家是否已回复')  TINYINT(1)"`
	Completed    string `xorm:"not null default 0 comment('是否完成')  TINYINT(1)"`
	CreatedAt    string `xorm:"not null default 0 comment('创建时间') index  INT(10)"`
	UpdatedAt    string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	CheckedAt    string `xorm:"not null default 0 comment('审核时间')  INT(10)"`
	CheckedBy    string `xorm:"not null default '' comment('审核人') VARCHAR(64)"`
	AssignAt     string `xorm:"not null default 0 comment('分配时间')  INT(10)"`
	PublishAt    string `xorm:"not null default 0 comment('发布时间')  INT(10)"`
	IsTest       string `xorm:"not null default 0 comment('是否为测试')  TINYINT(3)"`
	IsDel        string `xorm:"not null default 0 comment('是否删除')  TINYINT(3)"`
	ReplyPer     string `xorm:"not null default 0 comment('回复率') index  INT(10)"`
	CompletePer  string `xorm:"not null default 0 comment('完成率') index  INT(10)"`
	TopicId      string `xorm:"not null default 0 comment('话题ID')  INT(10)"`
}

type GrpProcs struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Gsn       string `xorm:"not null default 0 comment('集体投诉编号') index  BIGINT(18)"`
	Type      string `xorm:"not null default 0 comment('流程类型')  INT(10)"`
	Relation  string `xorm:"not null default '' comment('流程内容') VARCHAR(20000)"`
	VdiProg   string `xorm:"not null default 0 comment('音视频转码状态：0 默认 1 未完成') index  TINYINT(3)"`
	Uid       string `xorm:"not null default 0 comment('用户微博ID')  BIGINT(20)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel     string `xorm:"not null default 0 comment('是否删除')  TINYINT(3)"`
}

type HmRankCfg struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	FieldId    string `xorm:"not null default 0 comment('分类ID')  INT(10)"`
	FieldTitle string `xorm:"not null default '' comment('分类名称') VARCHAR(64)"`
	Title      string `xorm:"not null default '' comment('分组名称') VARCHAR(128)"`
	Sort       int    `xorm:"not null default 0 comment('显示排序') TINYINT(1)"`
	Show       int    `xorm:"not null default 0 comment('是否显示') TINYINT(1)"`
	ShowType   string `xorm:"not null default '' comment('榜单类型：周/月/季') VARCHAR(32)"`
	Companies  string `xorm:"not null comment('商家信息') TEXT"`
	CreatedAt  string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	CreatedBy  string `xorm:"not null default '' comment('创建人') VARCHAR(32)"`
	UpdatedAt  string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	UpdatedBy  string `xorm:"not null default '' comment('更新人') VARCHAR(32)"`
	IsDel      int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
	DelBy      string `xorm:"not null default '' comment('删除人') VARCHAR(32)"`
}

type HmRankData struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	Type       int    `xorm:"not null default 0 comment('榜单类型') index TINYINT(1)"`
	Date       string `xorm:"not null default 0 comment('日期') index  INT(10)"`
	Span       int    `xorm:"not null default 0 comment('榜单类型：周/月/季') index TINYINT(1)"`
	FieldId    string `xorm:"not null default 0 comment('分类ID') index  INT(10)"`
	FieldTitle string `xorm:"not null default '' comment('分类名称') VARCHAR(64)"`
	Data       string `xorm:"not null comment('数据') LONGTEXT"`
	CreatedAt  string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt  string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel      int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type HotComment struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Type      int    `xorm:"not null default 0 comment('类型') index TINYINT(1)"`
	Uid       string `xorm:"not null comment('用户UID') index  BIGINT(20)"`
	Sn        string `xorm:"not null comment('投诉单号') index  BIGINT(20)"`
	Content   string `xorm:"not null default '' comment('评论内容') VARCHAR(1024)"`
	Sort      int    `xorm:"not null default 0 comment('排序') TINYINT(1)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	CreatedBy string `xorm:"not null default '' comment('创建人') VARCHAR(32)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	UpdatedBy string `xorm:"not null default '' comment('更新人') VARCHAR(32)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type Instructions struct {
	Id        string `xorm:"not null pk autoincr comment('自增主键 ID')  INT(11)"`
	Title     string `xorm:"not null default '' comment('标题') VARCHAR(200)"`
	Summary   string `xorm:"not null default '' comment('摘要') VARCHAR(2000)"`
	Content   string `xorm:"not null comment('正文') MEDIUMTEXT"`
	Type      string `xorm:"not null default '' comment('类型') VARCHAR(30)"`
	Status    string `xorm:"not null default 0 comment('条文的状态，默认为草稿')  TINYINT(3)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(11)"`
	CreatedBy string `xorm:"not null default '' comment('创建者') VARCHAR(32)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(11)"`
	UpdatedBy string `xorm:"not null default '' comment('更新者') VARCHAR(32)"`
	IsDel     int    `xorm:"not null default 0 comment('是否被删除') TINYINT(3)"`
}

type Jury struct {
	Id           string `xorm:"not null pk autoincr comment('自增主键ID')  INT(10)"`
	Type         int    `xorm:"not null default 0 comment('类型') TINYINT(4)"`
	Uid          string `xorm:"not null default 0 comment('微博ID')  BIGINT(20)"`
	ShowName     string `xorm:"not null default '' comment('显示名称') VARCHAR(128)"`
	Intro        string `xorm:"not null default '' comment('简介') VARCHAR(2000)"`
	InviteAmount string `xorm:"not null default 0 comment('被邀请的数量')  INT(10)"`
	AcceptAmount string `xorm:"not null default 0 comment('参与案例的数量')  INT(10)"`
	Attach       string `xorm:"not null default '' comment('附件') VARCHAR(2048)"`
	Status       int    `xorm:"not null default 0 comment('入驻状态：0待审核，1审核通过，2审核失败') TINYINT(4)"`
	CheckedAt    string `xorm:"not null default 0 comment('审核时间')  INT(11)"`
	CheckedBy    string `xorm:"not null default '' comment('审核者') VARCHAR(32)"`
	Remark       string `xorm:"not null default '' comment('备注内容') VARCHAR(2000)"`
	CreatedAt    string `xorm:"not null default 0 comment('创建时间')  INT(11)"`
	CreatedBy    string `xorm:"not null default '' comment('创建者') VARCHAR(32)"`
	UpdatedAt    string `xorm:"not null default 0 comment('更新时间')  INT(11)"`
	UpdatedBy    string `xorm:"not null default '' comment('更新者') VARCHAR(32)"`
	IsDel        int    `xorm:"not null default 0 comment('是否被删除') TINYINT(3)"`
	Suggested    int    `xorm:"not null default 0 comment('是否推荐') TINYINT(3)"`
	VoteAmt      string `xorm:"not null default 0 comment('点赞数量')  INT(10)"`
	ReplyAmt     string `xorm:"not null default 0 comment('回复数量')  INT(10)"`
	Sort         string `xorm:"not null default 0 comment('排序')  INT(10)"`
}

type JuryFieldRel struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	Jid        string `xorm:"not null default 0 comment('评审团ID') index  INT(10)"`
	FieldId    string `xorm:"not null default 0 comment('一级领域id') index  INT(10)"`
	CreateTime string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type JuryInvites struct {
	Id          string `xorm:"not null pk autoincr comment('自增主键ID')  INT(10)"`
	ComplaintSn string `xorm:"not null comment('投诉单序列号') index(idx_sn_status)  BIGINT(20)"`
	JuryUid     string `xorm:"not null default 0 comment('评审团微博ID') index(idx_jury_status)  BIGINT(20)"`
	InviterUid  string `xorm:"not null default 0 comment('发起人微博ID')  BIGINT(20)"`
	Reason      string `xorm:"not null default '' comment('邀请理由') VARCHAR(2048)"`
	Status      int    `xorm:"not null default 0 comment('邀请状态：0待处理，1接受，2拒绝') index(idx_jury_status) index(idx_sn_status) TINYINT(4)"`
	CreatedAt   string `xorm:"not null default 0 comment('创建时间')  INT(11)"`
	UpdatedAt   string `xorm:"not null default 0 comment('更新时间')  INT(11)"`
	IsDel       int    `xorm:"not null default 0 comment('是否被删除') TINYINT(3)"`
	VoteNoticed int    `xorm:"not null default 0 comment('是否已发送点赞通知') TINYINT(1)"`
	Replytime   string `xorm:"not null default 0 comment('回复时间')  INT(10)"`
	Handletime  string `xorm:"not null default 0 comment('处理时间')  INT(10)"`
	Vote        int    `xorm:"not null default 0 comment('赞状态') TINYINT(1)"`
	Votetime    string `xorm:"not null default 0 comment('赞/取消时间')  INT(10)"`
}

type JuryVotes struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Juid      string `xorm:"not null default 0 comment('评审团微博ID') index(idx_jid_vote_uid)  BIGINT(20)"`
	Uid       string `xorm:"not null default 0 comment('用户微博ID') index(idx_jid_vote_uid)  BIGINT(20)"`
	Vote      int    `xorm:"not null default 0 comment('状态') index(idx_jid_vote_uid) TINYINT(1)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
}

type LaunchScreen struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Title     string `xorm:"not null default '' comment('广告名称') VARCHAR(128)"`
	ImgKey    string `xorm:"not null default '' comment('图片地址') VARCHAR(256)"`
	Link      string `xorm:"not null default '' comment('跳转链接') VARCHAR(256)"`
	JumpType  int    `xorm:"not null default 0 comment('跳转类型') TINYINT(1)"`
	Stay      int    `xorm:"not null default 0 comment('停留时长（秒）') TINYINT(1)"`
	Starttime string `xorm:"not null default 0 comment('开始投放时间') index  INT(10)"`
	Endtime   string `xorm:"not null default 0 comment('结束投放时间') index  INT(10)"`
	Status    int    `xorm:"not null default 0 comment('状态') index TINYINT(1)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	CreatedBy string `xorm:"not null default '' comment('创建人') VARCHAR(32)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	UpdatedBy string `xorm:"not null default '' comment('更新人') VARCHAR(32)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type Lawfirm struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Title     string `xorm:"not null default '' comment('标题') index VARCHAR(128)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	CreatedBy string `xorm:"not null default '' comment('创建人') VARCHAR(32)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	UpdatedBy string `xorm:"not null default '' comment('更新人') VARCHAR(32)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type LawfirmComplaint struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	LawfirmId string `xorm:"not null default 0 comment('律所ID') unique(idx_lawfirm_sn)  INT(10)"`
	Sn        string `xorm:"not null comment('投诉单号') unique(idx_lawfirm_sn) index  BIGINT(20)"`
	Status    int    `xorm:"not null default 0 comment('状态') index(idx_createtime_status) TINYINT(1)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间') index(idx_createtime_status)  INT(10)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type Laws struct {
	Id           string `xorm:"not null pk autoincr comment('自增主键内容')  INT(11)"`
	Title        string `xorm:"not null default '' comment('标题') VARCHAR(128)"`
	Summary      string `xorm:"not null default '' VARCHAR(2048)"`
	Content      string `xorm:"not null comment('正文内容') MEDIUMTEXT"`
	AppId        string `xorm:"not null default '' comment('推给 CMS 的时候的 app.id') VARCHAR(32)"`
	UpvoteAmount string `xorm:"not null default 0 comment('点赞的数量')  INT(11)"`
	ShareAmount  string `xorm:"not null default 0 comment('分享的数量')  INT(11)"`
	Status       string `xorm:"not null default 0 comment('条文的状态，默认为草稿')  TINYINT(3)"`
	CreatedAt    string `xorm:"not null default 0 comment('创建时间')  INT(11)"`
	CreatedBy    string `xorm:"not null default '' comment('创建者') VARCHAR(32)"`
	UpdatedAt    string `xorm:"not null default 0 comment('更新时间')  INT(11)"`
	UpdatedBy    string `xorm:"not null default '' comment('更新者') VARCHAR(32)"`
	IsDel        int    `xorm:"not null default 0 comment('是否被删除') TINYINT(3)"`
}

type Leadership struct {
	Id           string `xorm:"not null pk autoincr comment('自增主键 ID')  INT(10)"`
	Name         string `xorm:"not null default '' comment('名称') VARCHAR(255)"`
	Introduction string `xorm:"not null default '' comment('简介') VARCHAR(1000)"`
	ImgUrl       string `xorm:"not null default '' comment('图片链接') VARCHAR(1000)"`
	ShowOrder    string `xorm:"not null default 0 comment('显示顺序')  INT(10)"`
	Type         string `xorm:"not null default 0 comment('类型')  TINYINT(2)"`
	CreatedAt    string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt    string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel        string `xorm:"not null default 0 comment('是否被删除')  TINYINT(2)"`
}

type Logs202010 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202011 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202012 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202101 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202102 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202103 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202104 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202105 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202106 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202107 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202108 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202109 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202110 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202111 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202112 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202201 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type Logs202202 struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Pk        string `xorm:"not null default 0 comment('pk') index(idx_pk_path)  BIGINT(12)"`
	User      string `xorm:"not null default '' comment('username') VARCHAR(64)"`
	Ip        string `xorm:"not null default 0 comment('client_ip')  INT(10)"`
	Ua        string `xorm:"not null default '' comment('User-Agent') VARCHAR(128)"`
	Host      string `xorm:"not null default '' comment('操作 host') VARCHAR(32)"`
	Path      string `xorm:"not null default '' comment('server.path') index(idx_pk_path) VARCHAR(32)"`
	Request   string `xorm:"not null default '' comment('query string') VARCHAR(10000)"`
	Detail    string `xorm:"not null comment('detail') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type MdaAuthAssign struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	UserId     int    `xorm:"not null default 0 comment('用户表id') index INT(10)"`
	ItemId     int    `xorm:"not null default 0 comment('角色、权限或者路由的id') INT(10)"`
	CreateTime string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	Rolename   string `xorm:"not null default '' comment('角色名称') VARCHAR(255)"`
}

type MdaAuthItem struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	Name       string `xorm:"not null default '' comment('名称') VARCHAR(255)"`
	Type       int    `xorm:"not null default 1 comment('1.路由 2.权限 3.角色') TINYINT(1)"`
	Status     int    `xorm:"not null default 1 comment('状态 1.正常') TINYINT(1)"`
	CreateTime string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel      int    `xorm:"not null default 0 comment('是否删除 0 否 1 是') TINYINT(1)"`
}

type MdaAuthItemRel struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	Parent     int    `xorm:"not null default 0 comment('父级id') INT(10)"`
	Child      int    `xorm:"not null default 0 comment('子id') INT(10)"`
	CreateTime string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
}

type MdaAuthMenu struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	Name       string `xorm:"not null default '0' comment('菜单名称') VARCHAR(255)"`
	ParentId   int    `xorm:"not null default 0 comment('父级菜单id') INT(11)"`
	Route      string `xorm:"not null default '' comment('菜单对应路由') VARCHAR(255)"`
	Status     int    `xorm:"not null default 1 comment('状态 1.正常') TINYINT(1)"`
	Icon       string `xorm:"not null default '' comment('菜单样式') VARCHAR(500)"`
	Sort       int    `xorm:"not null default 100 comment('排序') INT(10)"`
	CreateTime string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel      int    `xorm:"not null default 0 comment('是否删除 0否 1.是') TINYINT(1)"`
}

type MdaAuthUser struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	Uid        string `xorm:"not null default 0 comment('微博ID') index  BIGINT(20)"`
	Couid      string `xorm:"not null default 0 comment('商家微博ID') index  BIGINT(20)"`
	Name       string `xorm:"not null default '' comment('微博昵称') VARCHAR(255)"`
	Phone      string `xorm:"not null default '' comment('手机号码') VARCHAR(30)"`
	CreateTime string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel      int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type MdaCompanyMonthStast struct {
	Id          string `xorm:"not null pk autoincr  INT(10)"`
	Couid       string `xorm:"not null default 0 comment('商家微博ID') index  BIGINT(18)"`
	Date        string `xorm:"not null default 0 comment('日期')  INT(10)"`
	Valid       string `xorm:"not null default 0 comment('有效量')  INT(10)"`
	Reply       string `xorm:"not null default 0 comment('回复量')  INT(10)"`
	Complete    string `xorm:"not null default 0 comment('完成量')  INT(10)"`
	ReplyPer    string `xorm:"not null default '' comment('月累计回复率') VARCHAR(16)"`
	CompletePer string `xorm:"not null default '' comment('月累计完成率') VARCHAR(16)"`
	Replytime   string `xorm:"not null default '' comment('月累计回复时间') VARCHAR(16)"`
	Eval        string `xorm:"not null default '' comment('评价') VARCHAR(1024)"`
	CreateTime  string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime  string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel       int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type MdaContracts struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	Couid      string `xorm:"not null default 0 comment('商家UID') index  BIGINT(20)"`
	Cotitle    string `xorm:"not null default '' comment('商家名称') index VARCHAR(128)"`
	Code       string `xorm:"not null default '' comment('合同编号') VARCHAR(128)"`
	Type       int    `xorm:"not null default 0 comment('合同类型') index TINYINT(1)"`
	Price      string `xorm:"not null default 0 comment('合同金额')  INT(10)"`
	Startime   string `xorm:"not null default 0 comment('开始日期')  INT(10)"`
	Endtime    string `xorm:"not null default 0 comment('截止日期')  INT(10)"`
	YmExptime  string `xorm:"not null default 0 comment('铀媒账号有效期')  INT(10)"`
	Status     int    `xorm:"not null default 0 comment('履约状态') index TINYINT(1)"`
	CreatedBy  string `xorm:"not null default '' comment('创建人') VARCHAR(32)"`
	UpdatedBy  string `xorm:"not null default '' comment('更新人') VARCHAR(32)"`
	CreatedAt  string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt  string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel      int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
	SafeSign   int    `xorm:"not null default 0 comment('签署安全协议状态') TINYINT(1)"`
	SafeSignAt string `xorm:"not null default 0 comment('签署安全协议时间')  INT(10)"`
}

type MdaCos struct {
	Id                    string `xorm:"not null pk autoincr  INT(10)"`
	Uid                   string `xorm:"not null default 0 comment('微博ID') unique  BIGINT(20)"`
	Title                 string `xorm:"not null default '' comment('公司名') VARCHAR(128)"`
	Status                int    `xorm:"not null default 1 comment('入驻状态：1正常') TINYINT(4)"`
	MsgCfg                string `xorm:"not null default '' comment('消息配置') VARCHAR(1024)"`
	FltCfg                string `xorm:"not null default '' comment('删选配置') VARCHAR(1024)"`
	TagCfg                string `xorm:"not null default '' comment('标签配置') VARCHAR(2048)"`
	ColCfg                string `xorm:"not null default '' comment('表头配置') VARCHAR(1024)"`
	Alertword             string `xorm:"not null default '' comment('敏感词配置') VARCHAR(2048)"`
	CreateTime            string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime            string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	CreateBy              string `xorm:"not null default '' comment('创建人') VARCHAR(64)"`
	IsDel                 int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
	SmsLimit              string `xorm:"not null default 0 comment('短信发送上限')  INT(10)"`
	SmsSent               string `xorm:"not null default 0 comment('已发送短信量')  INT(10)"`
	UmReg                 int    `xorm:"not null default 0 comment('是否注册铀媒账号') TINYINT(1)"`
	UmEndtime             string `xorm:"not null default 0 comment('铀媒账号到期时间')  INT(10)"`
	UmEndtimeChg          int    `xorm:"not null default 0 comment('是否修改铀媒账号到期时间') TINYINT(1)"`
	YqtKeyword            string `xorm:"not null default '' comment('舆情通关键词') VARCHAR(512)"`
	Autoreply             int    `xorm:"not null default 0 comment('自动回复：0关闭 1开启') TINYINT(1)"`
	Autoreplytime         string `xorm:"not null comment('自动回复触发时间') TEXT"`
	Autoreplytpl          string `xorm:"not null comment('自动回复模版') TEXT"`
	AutoreplyCfgtime      string `xorm:"not null default 0 comment('自动回复开关配置时间')  INT(10)"`
	RolegrpId             int    `xorm:"not null default 0 comment('角色组id') INT(10)"`
	Competitors           string `xorm:"not null default '' comment('竟品商家') VARCHAR(256)"`
	CompetitorAccounttype int    `xorm:"not null default 1 comment('竞品账号类型：1主账号 2子账号 3主子账号') TINYINT(1)"`
	ContractValid         int    `xorm:"not null default 0 comment('履约状态') TINYINT(1)"`
	ContractEndtime       string `xorm:"not null default 0 comment('履约截止日期')  INT(10)"`
}

type MdaPunishRecords struct {
	Id          string `xorm:"not null pk autoincr  INT(10)"`
	Type        string `xorm:"not null default 0 comment('类型:1结案')  TINYINT(1)"`
	Uid         string `xorm:"not null default 0 comment('推送对象微博ID') index  BIGINT(20)"`
	Cotitle     string `xorm:"not null default '' comment('推送对象名称') VARCHAR(128)"`
	DisableDays string `xorm:"not null default 0 comment('功能停用时长	')  INT(10)"`
	Startime    string `xorm:"not null default 0 comment('开始时间')  INT(10)"`
	Endtime     string `xorm:"not null default 0 comment('结束时间')  INT(10)"`
	CreatedAt   string `xorm:"not null default 0 comment('创建时间') index  INT(10)"`
	UpdatedAt   string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel       int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
	PushNum     string `xorm:"not null default '' comment('推送次数') VARCHAR(8)"`
}

type MdaPunishRules struct {
	Id               string `xorm:"not null pk autoincr  INT(10)"`
	AssessmentMonths string `xorm:"not null default 0 comment('考核周期')  TINYINT(1)"`
	Type             string `xorm:"not null default 0 comment('类型:1结案')  TINYINT(1)"`
	Startime         string `xorm:"not null default 0 comment('开始时间')  INT(10)"`
	Endtime          string `xorm:"not null default 0 comment('结束时间')  INT(10)"`
	ViolationNum     string `xorm:"not null default 0 comment('违规达标次数')  INT(10)"`
	DisableDays      string `xorm:"not null default 0 comment('功能停用时长	')  INT(10)"`
	CreatedAt        string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedBy        string `xorm:"not null default '' comment('创建人') VARCHAR(32)"`
	UpdatedAt        string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel            int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type MdaSaleClues struct {
	Id           string `xorm:"not null pk autoincr  INT(10)"`
	Cotitle      string `xorm:"not null default '' comment('企业名称') index VARCHAR(64)"`
	ContactName  string `xorm:"not null default '' comment('联系人姓名') VARCHAR(32)"`
	ContactPhone string `xorm:"not null default '' comment('联系人电话') VARCHAR(128)"`
	ContactJob   string `xorm:"not null default '' comment('岗位') VARCHAR(32)"`
	From         int    `xorm:"not null default 0 comment('来源') TINYINT(1)"`
	Md5Phone     string `xorm:"not null default '' comment('电话索引') index VARCHAR(32)"`
	CreatedAt    string `xorm:"not null default 0 comment('创建时间') index  INT(10)"`
	UpdatedAt    string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel        int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type MdaSms struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	Couid      string `xorm:"not null default 0 comment('商家微博ID') index(idx_couid_time)  BIGINT(20)"`
	Cotitle    string `xorm:"not null default '' comment('商家名称') VARCHAR(128)"`
	Phone      string `xorm:"not null default '' comment('目标手机号码') VARCHAR(32)"`
	Content    string `xorm:"not null default '' comment('内容') VARCHAR(2048)"`
	Sn         string `xorm:"not null default 0 comment('投诉单号')  BIGINT(18)"`
	Type       string `xorm:"not null default 0 comment('短信类型')  INT(5)"`
	Status     int    `xorm:"not null default 0 comment('发送状态：1成功 2失败') TINYINT(1)"`
	Result     string `xorm:"not null default '' comment('企信通发送响应') VARCHAR(1024)"`
	CreateTime string `xorm:"not null default 0 comment('创建时间') index(idx_couid_time)  INT(10)"`
	Limit      string `xorm:"not null default 0 comment('短信上限数量')  INT(10)"`
	Left       string `xorm:"not null default 0 comment('短信剩余数量')  INT(10)"`
	IsDel      int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type MdaStastDay struct {
	Id            string `xorm:"not null pk autoincr  INT(10)"`
	Couid         string `xorm:"not null default 0 comment('商家微博ID') index  BIGINT(18)"`
	Date          string `xorm:"not null default 0 comment('日期')  INT(10)"`
	Valid         string `xorm:"not null default 0 comment('有效量')  INT(10)"`
	Reply         string `xorm:"not null default 0 comment('回复量')  INT(10)"`
	Complete      string `xorm:"not null default 0 comment('完成量')  INT(10)"`
	ReplyPer      string `xorm:"not null default '' comment('月累计回复率') VARCHAR(16)"`
	CompletePer   string `xorm:"not null default '' comment('月累计完成率') VARCHAR(16)"`
	Replytime     string `xorm:"not null default '' comment('月累计回复时间') VARCHAR(16)"`
	Noreply       string `xorm:"not null default 0 comment('待回复量')  INT(10)"`
	Eval          string `xorm:"not null default '' comment('评价') VARCHAR(1024)"`
	Appeal        string `xorm:"not null default '' comment('申诉数据') VARCHAR(1024)"`
	Autoreply     string `xorm:"not null default 0 comment('自动回复量')  INT(10)"`
	AutoreplyPer  string `xorm:"not null default '' comment('月累计自动回复率') VARCHAR(16)"`
	Cocomplete    string `xorm:"not null default 0 comment('自主结案完成量')  INT(10)"`
	CocompletePer string `xorm:"not null default '' comment('月累计自主结案完成率') VARCHAR(16)"`
	DgValid       string `xorm:"not null default 0 comment('高危投诉有效量')  INT(10)"`
	DgReply       string `xorm:"not null default 0 comment('高危投诉回复量')  INT(10)"`
	DgComplete    string `xorm:"not null default 0 comment('高危投诉完成量')  INT(10)"`
	DgReplyPer    string `xorm:"not null default '' comment('高危投诉月累计回复率') VARCHAR(16)"`
	DgCompletePer string `xorm:"not null default '' comment('高危投诉月累计完成率') VARCHAR(16)"`
	DgReplytime   string `xorm:"not null default '' comment('高危投诉月累计回复时间') VARCHAR(16)"`
	DgNoreply     string `xorm:"not null default 0 comment('高危投诉待回复量')  INT(10)"`
	DgEval        string `xorm:"not null default '' comment('高危投诉评价') VARCHAR(1024)"`
	DgPltValid    string `xorm:"not null default 0 comment('包含平台敏感词的投诉量')  INT(10)"`
	DgCoValid     string `xorm:"not null default 0 comment('包含商家敏感词的投诉量')  INT(10)"`
	CreateTime    string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime    string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel         int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type MdaSysmsgRecords struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	TaskId     string `xorm:"not null default 0 comment('任务ID') index  INT(10)"`
	Reason     string `xorm:"not null default 0 comment('推送原因') index  TINYINT(1)"`
	TplType    string `xorm:"not null default 0 comment('模板类别') index  TINYINT(1)"`
	Title      string `xorm:"not null default '' comment('标题') index VARCHAR(256)"`
	Content    string `xorm:"not null default '' comment('内容') VARCHAR(4096)"`
	Couid      string `xorm:"not null default 0 comment('推送商家微博ID') index  BIGINT(20)"`
	Cotitle    string `xorm:"not null default '' comment('推送商家名称') VARCHAR(128)"`
	Uid        string `xorm:"not null default 0 comment('推送商家成员微博ID') index  BIGINT(20)"`
	Read       int    `xorm:"not null default 0 comment('已读0.否 1.是') TINYINT(1)"`
	CanceledBy string `xorm:"not null default '' comment('撤回人') VARCHAR(32)"`
	CanceledAt string `xorm:"not null default 0 comment('撤回时间') index  INT(10)"`
	CreatedBy  string `xorm:"not null default '' comment('创建人') VARCHAR(32)"`
	CreatedAt  string `xorm:"not null default 0 comment('创建时间') index  INT(10)"`
	UpdatedAt  string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel      int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type MdaSysmsgTasks struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Reason    string `xorm:"not null default 0 comment('推送原因') index  TINYINT(1)"`
	Title     string `xorm:"not null default '' comment('标题') index VARCHAR(256)"`
	Content   string `xorm:"not null default '' comment('内容') VARCHAR(4096)"`
	Uids      string `xorm:"not null default '' comment('推送对象微博ID') VARCHAR(1024)"`
	WhlUids   string `xorm:"not null default '' comment('白名单对象微博ID') VARCHAR(1024)"`
	Status    string `xorm:"not null default 0 comment('状态')  TINYINT(1)"`
	CreatedBy string `xorm:"not null default '' comment('创建人') VARCHAR(32)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间') index  INT(10)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type MdaSysmsgTpls struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Type      string `xorm:"not null default 0 comment('类型')  TINYINT(1)"`
	Title     string `xorm:"not null default '' comment('标题') VARCHAR(256)"`
	Content   string `xorm:"not null default '' comment('内容') VARCHAR(4096)"`
	CreatedBy string `xorm:"not null default '' comment('创建人') VARCHAR(32)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedBy string `xorm:"not null default '' comment('更新人') VARCHAR(32)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type MdaYqtReport struct {
	Id              string `xorm:"not null pk autoincr  INT(10)"`
	Couid           string `xorm:"not null default 0 comment('商家微博ID') index(idx_couid_st)  BIGINT(20)"`
	Title           string `xorm:"not null default '' comment('标题') VARCHAR(512)"`
	Keyword         string `xorm:"not null default '' comment('关键词') VARCHAR(512)"`
	Filterkeyword   string `xorm:"not null default '' comment('过滤关键词') VARCHAR(512)"`
	Stime           string `xorm:"not null default 0 comment('统计开始时间') index(idx_couid_st)  INT(10)"`
	Etime           string `xorm:"not null default 0 comment('统计结束时间')  INT(10)"`
	TaskId          string `xorm:"not null default 0 comment('任务ID')  INT(10)"`
	Ticket          string `xorm:"not null default '' comment('ticket') VARCHAR(64)"`
	YqtStatus       int    `xorm:"not null default 0 comment('任务状态') TINYINT(4)"`
	Spreadtrend     string `xorm:"comment('传播走势图') TEXT"`
	Negativetrend   string `xorm:"comment('负面走势图') TEXT"`
	Sensitivestast  string `xorm:"comment('情感分析占比') TEXT"`
	Positiveword    string `xorm:"comment('正面高频词汇') TEXT"`
	Negativeword    string `xorm:"comment('负面高频词汇') TEXT"`
	Wordcloud       string `xorm:"comment('关键词云') TEXT"`
	Areastast       string `xorm:"comment('信息地域分布') TEXT"`
	Sourcetype      string `xorm:"comment('来源类型对比') TEXT"`
	Sourcemedia     string `xorm:"comment('媒体活跃度对比') TEXT"`
	Mediamonitoring string `xorm:"comment('重点媒体监测') TEXT"`
	Status          int    `xorm:"not null default 0 comment('状态:0处理中 1成功 2失败') index TINYINT(4)"`
	CreateTime      string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime      string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel           int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type Messages struct {
	Id        string `xorm:"not null pk autoincr comment('自增主键 ID')  INT(10)"`
	Type      string `xorm:"not null default 0 comment('站内信模版类型')  INT(5)"`
	Uid       string `xorm:"not null default 0 comment(' 接收方 uid') index(idx_uid_st_del)  BIGINT(20)"`
	Relation  string `xorm:"not null default '' comment('文本内容') VARCHAR(2048)"`
	Status    string `xorm:"not null default 0 comment('消息状态') index(idx_uid_st_del)  TINYINT(3)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	IsDel     string `xorm:"not null default 0 comment('是否被删除') index(idx_uid_st_del)  TINYINT(3)"`
	Danger    string `xorm:"not null default 0 comment('高危投诉')  TINYINT(1)"`
	Msgtype   string `xorm:"not null default 1 comment('通知类型：1投诉通知2系统通知')  TINYINT(1)"`
}

type OpWorkConfiguration struct {
	Id        string `xorm:"not null pk autoincr comment('主键')  INT(10)"`
	Day       string `xorm:"not null default 0 comment('日期')  INT(10)"`
	TplType   string `xorm:"not null default 0 comment('模板类型，1：平日，2假日')  TINYINT(2)"`
	Config    string `xorm:"not null comment('具体配置信息') TEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	CreatedBy string `xorm:"not null default '' comment('创建人') VARCHAR(32)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	UpdatedBy string `xorm:"not null default '' comment('更新人') VARCHAR(32)"`
}

type Procedures struct {
	Id        string `xorm:"not null pk autoincr comment('自增主键')  INT(10)"`
	Sn        string `xorm:"not null default 0 comment('投诉单号') index  BIGINT(12)"`
	Type      string `xorm:"default 0 comment('流程类型') index(idx_couid_tp_del_ct) index(idx_ctime_type) index  INT(10)"`
	Relation  string `xorm:"not null default '' comment('处理内容') VARCHAR(20000)"`
	CreatedAt int    `xorm:"not null default 0 comment('创建时间') index(idx_couid_tp_del_ct) index(idx_ctime_type) INT(11)"`
	IsDel     string `xorm:"not null default 0 comment(' 是否被删除') index(idx_couid_tp_del_ct)  TINYINT(3)"`
	Status    string `xorm:"not null default 0 comment('0 默认 1音视频未处理完')  TINYINT(3)"`
	Couid     string `xorm:"not null default 0 comment('商家微博ID') index(idx_couid_tp_del_ct)  BIGINT(20)"`
}

type ProceduresTest struct {
	Id        string `xorm:"not null pk autoincr comment('自增主键')  INT(10)"`
	Sn        string `xorm:"not null default 0 comment('投诉单号') index  BIGINT(12)"`
	Type      string `xorm:"default 0 comment('流程类型') index(idx_couid_tp_del_ct) index  INT(10)"`
	Relation  string `xorm:"not null default '' comment('处理内容') VARCHAR(20000)"`
	CreatedAt int    `xorm:"not null default 0 comment('创建时间') index(idx_couid_tp_del_ct) INT(11)"`
	IsDel     string `xorm:"not null default 0 comment(' 是否被删除') index(idx_couid_tp_del_ct)  TINYINT(3)"`
	Status    string `xorm:"not null default 0 comment('0 默认 1音视频未处理完')  TINYINT(3)"`
	Couid     string `xorm:"not null default 0 comment('商家微博ID') index(idx_couid_tp_del_ct)  BIGINT(20)"`
}

type RewardProducts struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Code      string `xorm:"not null default '' comment('代码') index VARCHAR(256)"`
	CodeIos   string `xorm:"not null default '' comment('ios代码') VARCHAR(128)"`
	Price     string `xorm:"not null default 0 comment('金额')  INT(10)"`
	Title     string `xorm:"not null default '' comment('名称') VARCHAR(128)"`
	Image     string `xorm:"not null default '' comment('图片') VARCHAR(512)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type RewardTrades struct {
	Id          string `xorm:"not null pk autoincr  INT(10)"`
	OrderId     string `xorm:"not null default '' comment('订单号') unique VARCHAR(128)"`
	Sn          string `xorm:"not null default 0 comment('投诉号') index  BIGINT(18)"`
	ProductCode string `xorm:"not null default '' comment('商品代码') index VARCHAR(256)"`
	ProductNum  string `xorm:"not null default 0 comment('商品数量')  INT(10)"`
	Uid         string `xorm:"not null default 0 comment('用户微博ID') index  BIGINT(20)"`
	PayType     string `xorm:"not null default 0 comment('支付方式') index  TINYINT(1)"`
	PayAt       string `xorm:"not null default 0 comment('支付时间')  INT(10)"`
	PayId       string `xorm:"not null default '' comment('支付ID') VARCHAR(256)"`
	PayPrice    string `xorm:"not null default 0 comment('支付金额')  INT(10)"`
	Status      string `xorm:"not null default 0 comment('订单状态') index  TINYINT(1)"`
	Source      string `xorm:"not null default 0 comment('来源端口') index  TINYINT(1)"`
	ShareImg    string `xorm:"not null default '' comment('分享图片地址') VARCHAR(1024)"`
	CreatedAt   string `xorm:"not null default 0 comment('创建时间') index  INT(10)"`
	UpdatedAt   string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel       int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type SlideImages struct {
	Id           string `xorm:"not null pk autoincr comment('自增主键 ID')  INT(10)"`
	Title        string `xorm:"not null default '' comment('标题') VARCHAR(500)"`
	ImgUrl       string `xorm:"not null default '' comment('图片链接') VARCHAR(1000)"`
	Link         string `xorm:"not null default '' comment('图片超链接') VARCHAR(1000)"`
	ShowOrder    string `xorm:"not null default 0 comment('显示顺序')  TINYINT(3)"`
	Type         string `xorm:"not null default 0 comment('类型：1为pc,2为wap')  TINYINT(1)"`
	CreatedAt    string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt    string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel        string `xorm:"not null default 0 comment('是否被删除')  TINYINT(3)"`
	LocalStation string `xorm:"not null default '' comment('地方站') VARCHAR(20)"`
}

type Sms struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Uid       string `xorm:"not null default 0 comment('发送此短信的用户 UID')  BIGINT(20)"`
	Phone     string `xorm:"not null default '' comment('目标手机号码') CHAR(11)"`
	Content   string `xorm:"not null default '' comment('短信内容') VARCHAR(2048)"`
	Response  string `xorm:"not null default '' comment('企信通发送响应') VARCHAR(1024)"`
	Status    string `xorm:"not null default 0  TINYINT(1)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
}

type SpecialCompanies struct {
	Id         string `xorm:"not null pk autoincr  INT(10)"`
	Couid      string `xorm:"not null default 0 comment('商家微博ID')  BIGINT(20)"`
	Cotitle    string `xorm:"not null default '' comment('商家名称') VARCHAR(128)"`
	Tag        string `xorm:"not null default '' comment('标签') index VARCHAR(1024)"`
	CreateTime string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdateTime string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel      int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type SuggestWords struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Word      string `xorm:"not null default '' comment('suggest 关键词') VARCHAR(255)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	CreatedBy string `xorm:"not null default '' comment('更新者') VARCHAR(255)"`
	IsDel     string `xorm:"not null default 0  TINYINT(3)"`
}

type SyncOrg struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Title     string `xorm:"not null default '' comment('标题') index VARCHAR(128)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	CreatedBy string `xorm:"not null default '' comment('创建人') VARCHAR(32)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	UpdatedBy string `xorm:"not null default '' comment('更新人') VARCHAR(32)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type SyncOrgComplaint struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	OrgId     string `xorm:"not null default 0 comment('机构ID') unique(idx_org_sn)  INT(10)"`
	Sn        string `xorm:"not null comment('投诉单号') unique(idx_org_sn) index  BIGINT(20)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间') index  INT(10)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type Topic struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Title     string `xorm:"not null default '' comment('标题') index VARCHAR(128)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	CreatedBy string `xorm:"not null default '' comment('创建人') VARCHAR(32)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	UpdatedBy string `xorm:"not null default '' comment('更新人') VARCHAR(32)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type TopicCompany struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	Uid       string `xorm:"not null comment('商家UID') index  BIGINT(20)"`
	Remark    string `xorm:"not null default '' comment('备注') VARCHAR(128)"`
	Sort      int    `xorm:"not null default 0 comment('排序') TINYINT(1)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	CreatedBy string `xorm:"not null default '' comment('创建人') VARCHAR(32)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	UpdatedBy string `xorm:"not null default '' comment('更新人') VARCHAR(32)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type TopicComplaintRel struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	TopicId   string `xorm:"not null default 0 comment('话题ID') index unique(idx_topic_sn)  INT(10)"`
	Sn        string `xorm:"not null comment('投诉单号') index unique(idx_topic_sn)  BIGINT(20)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间') index  INT(10)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type TopicStat struct {
	Id        string `xorm:"not null pk autoincr  INT(10)"`
	TopicId   string `xorm:"not null default 0 comment('话题ID') index  INT(10)"`
	Type      int    `xorm:"not null default 0 comment('类型') index TINYINT(1)"`
	Date      string `xorm:"not null default 0 comment('日期') index  INT(10)"`
	Data      string `xorm:"not null comment('数据') LONGTEXT"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	UpdatedAt string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel     int    `xorm:"not null default 0 comment('是否删除0.否 1.是') TINYINT(1)"`
}

type UserInfo struct {
	Id              string `xorm:"not null pk autoincr comment('自增主键')  INT(10)"`
	Uid             string `xorm:"not null default 0 comment('用户uid') unique  BIGINT(20)"`
	Phone           string `xorm:"not null default '' comment('用户手机号') VARCHAR(255)"`
	Md5Phone        string `xorm:"not null default '' comment('手机号的md5') index VARCHAR(255)"`
	TotalComplaints int    `xorm:"not null default 0 comment('创建过的投诉单总数') INT(10)"`
	CreatedAt       string `xorm:"not null default 0 comment('用户创建时间')  INT(10)"`
	UpdatedAt       string `xorm:"not null default 0 comment('更新时间')  INT(10)"`
	IsDel           string `xorm:"not null default 0 comment('是否删除')  TINYINT(2)"`
}

type Votes struct {
	Id        string `xorm:"not null pk autoincr comment('自增主键')  INT(10)"`
	Uid       string `xorm:"not null default 0 comment('uid') index(idx_rid_type_uid) index(idx_uid_rid)  BIGINT(20)"`
	Type      string `xorm:"not null default 0 comment('关联主键类型 ') index(idx_rid_type) index(idx_rid_type_uid)  TINYINT(3)"`
	Rid       string `xorm:"not null default 0 comment('关联主键 ID') index index(idx_rid_type) index(idx_rid_type_uid) index(idx_uid_rid)  BIGINT(12)"`
	Vote      string `xorm:"not null default 0 comment('点赞与否')  TINYINT(3)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	IsDel     string `xorm:"not null default 0 comment('是否被删除')  TINYINT(3)"`
}

type VotesBak struct {
	Id        string `xorm:"not null pk autoincr comment('自增主键')  INT(10)"`
	Uid       string `xorm:"not null default 0 comment('uid')  BIGINT(20)"`
	Type      string `xorm:"not null default 0 comment('关联主键类型 ')  TINYINT(3)"`
	Rid       string `xorm:"not null default 0 comment('关联主键 ID') index  BIGINT(12)"`
	Vote      string `xorm:"not null default 0 comment('点赞与否')  TINYINT(3)"`
	CreatedAt string `xorm:"not null default 0 comment('创建时间')  INT(10)"`
	IsDel     string `xorm:"not null default 0 comment('是否被删除')  TINYINT(3)"`
}

type WeiboMsg struct {
	Id        string `xorm:"not null pk autoincr comment('自增主键 ID')  INT(10)"`
	Uid       string `xorm:"not null default 0 comment(' 接收方 uid')  BIGINT(20)"`
	Content   string `xorm:"not null default '' comment('消息内容') VARCHAR(2048)"`
	Response  string `xorm:"not null default '' comment('微博发送响应') VARCHAR(2048)"`
	Status    string `xorm:"not null default 0 comment('发送状态')  TINYINT(3)"`
	CreatedAt string `xorm:"not null default 0 comment('发送时间')  INT(10)"`
	IsDel     string `xorm:"not null default 0 comment('是否被删除')  TINYINT(3)"`
}

