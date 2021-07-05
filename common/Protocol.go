package common

import (
	"encoding/json"
	"net"
	"time"

	"github.com/segmentio/ksuid"
)

type Response struct {
	Errno int         `json:"errno"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

// {"phoneno":"18963601153","fee":20000,"contractddl1":"20181111","contracturl":"aaaa/bbbb/cccc","contractddl2":"20181111","contractddl5":"20181111","contractddl3":"20181111","contractddl4":"20181111","fee1":20000,"fee2":20000,"fee3":20000,"fee4":20000,"fee5":20000,"birthday":"19880521","entrydate":"20180521","bankname":"招商银行","cardno":"622123231233221","testfee":10000,"phoneno1":"18963601153","phoneno2":"18963601153","username":"李乃东","workid":"0536","certno":"320902198805216072","password":"wssglnd","usertype":"管理员","userdept":"技术","userinfo":"开发人员"}

//员工信息表
type UserInfo struct {
	PhoneNo       string  `json:"phoneNo" bson:"phoneNo"`             //手机号码
	PhoneNo1      string  `json:"phoneNo1" bson:"phoneNo1"`           //紧急手机号码
	PhoneNo2      string  `json:"phoneNo2" bson:"phoneNo2"`           //紧急手机号码
	UserName      string  `json:"userName" bson:"userName"`           //用户名
	WorkId        string  `json:"workId" bson:"workId"`               //员工号
	CertNo        string  `json:"certNo" bson:"certNo"`               //身份证号码
	PassWord      string  `json:"passWord" bson:"passWord"`           //密码
	PassWordSts   string  `json:"passWordSts" bson:"passWordSts"`     //密码
	UserRoles     string  `json:"userRoles" bson:"userRoles"`         //用户岗位
	UserType      string  `json:"userType" bson:"userType"`           //用户岗位
	UserDept      string  `json:"userDept" bson:"userDept"`           //用户部门
	UserStatus    string  `json:"userStatus" bson:"userStatus"`       //用户状态
	UserInfo      string  `json:"userInfo" bson:"userInfo"`           //user信息
	Birthday      string  `json:"birthday" bson:"birthday"`           //生日
	BankName      string  `json:"bankName" bson:"bankName"`           //入职日期
	TestFee       float64 `json:"testFee" bson:"testFee"`             //实习工资
	Fee           float64 `json:"fee" bson:"fee"`                     //正式工资
	CardNo        string  `json:"cardNo" bson:"cardNo"`               //银行卡号
	EntryDate     string  `json:"entryDate" bson:"entryDate"`         //入职日期
	Fee1          float64 `json:"fee1" bson:"fee1"`                   //薪资调整1
	Fee2          float64 `json:"fee2" bson:"fee2"`                   //薪资调整2
	Fee3          float64 `json:"fee3" bson:"fee3"`                   //薪资调整3
	Fee4          float64 `json:"fee4" bson:"fee4"`                   //薪资调整4
	Fee5          float64 `json:"fee5" bson:"fee5"`                   //薪资调整5
	ContractDdl1  string  `json:"contractDdl1" bson:"contractDdl1"`   //合同到期日1
	ContractDdl2  string  `json:"contractDdl2" bson:"contractDdl2"`   //合同到期日2
	ContractDdl3  string  `json:"contractDdl3" bson:"contractDdl3"`   //合同到期日3
	ContractDdl4  string  `json:"contractDdl4" bson:"contractDdl4"`   //合同到期日4
	ContractDdl5  string  `json:"contractDdl5" bson:"contractDdl5"`   //合同到期日5
	ContractTimes string  `json:"contractTimes" bson:"contractTimes"` //合同修订次数
	ContractUrl   string  `json:"contractUrl" bson:"contractUrl"`     //合同图片Url
	CreateTime    string  `json:"createTime" bson:"createTime"`       //user信息
}

type SysConfig struct {
	ConfigName string `json:"configName" bson:"configName"` //配置名称
	ConfigInfo string `json:"configInfo" bson:"configInfo"` //配置内容
	CreateTime string `json:"createTime" bson:"createTime"` //时间
}

type SaleChannel struct {
	ChannelName string `json:"channelName" bson:"channelName"` //渠道名称
	ChannelInfo string `json:"channelInfo" bson:"channelInfo"` //渠道说明
	CreateTime  string `json:"createTime" bson:"createTime"`   //user信息
}

type Department struct {
	DepartName string `json:"departName" bson:"departName"` //部门名称
	DepartInfo string `json:"departInfo" bson:"departInfo"` //部门说明
	CreateTime string `json:"createTime" bson:"createTime"` //user信息
}

type Role struct {
	RoleName   string `json:"roleName" bson:"roleName"`     //角色名称
	RoleInfo   string `json:"roleInfo" bson:"roleInfo"`     //角色说明
	CreateTime string `json:"createTime" bson:"createTime"` //user信息
}

type RoleSaleChannel struct {
	RoleName     string `json:"roleName" bson:"roleName"`
	SaleChannels string `json:"saleChannels" bson:"saleChannels"`
	CreateTime   string `json:"createTime" bson:"createTime"`
}

// 线索编号、线索状态、渠道分类、营销渠道、渠道备注、业务类型、联系人或公司名称、客户备注、客户经理、部门、线索类别、废弃原因、
// 下次追踪日期、最近一次工作记录、最近一次追踪时间、录入时间、录入人员、电话、追踪总次数、首次分配人、分配后追踪次数、分配时间、成单日期、线索成本、咨询备注、联系方式备注、订单编号；
//线索表
type Clue struct {
	ClueNo         string `uname:"线索编号" json:"clueNo" bson:"clueNo" xlsx:"clueNo"`    //线索编号 新增时不需要输入，系统自动生成
	ClueStatus     string `uname:"线索状态" json:"clueStatus" bson:"clueStatus"`          //线索状态
	ClueType       string `uname:"线索类别" json:"clueType" bson:"clueType"`              //线索类别
	ChannelType    string `uname:"渠道分类" json:"channelType" bson:"channelType"`        //渠道分类
	SaleChannel    string `uname:"营销渠道" json:"saleChannel" bson:"saleChannel"`        //营销渠道
	ChannelInfo    string `uname:"渠道备注" json:"channelInfo" bson:"channelInfo"`        //渠道备注
	BussType       string `uname:"业务类型" json:"bussType" bson:"bussType"`              //业务类型
	ContractName   string `uname:"联系人或公司名称" json:"contractName" bson:"contractName"`  //联系人或公司名称
	CustomInfo     string `uname:"客户备注" json:"customInfo" bson:"customInfo"`          //客户备注
	CustomMgr      string `uname:"客户经理" json:"customMgr" bson:"customMgr"`            //客户经理
	Dept           string `uname:"部门" json:"dept" bson:"dept"`                        //部门
	DropInfo       string `uname:"废弃原因" json:"dropInfo" bson:"dropInfo"`              //废弃原因
	DropReason     string `uname:"废弃备注" json:"dropReason" bson:"dropReason"`          //废弃备注
	TraceDate      string `uname:"最近一次追踪时间" json:"traceDate" bson:"traceDate"`        //最近一次追踪时间
	TraceInfo      string `uname:"最近一次工作记录" json:"traceInfo" bson:"traceInfo"`        //最近一次工作记录
	NextTraceDate  string `uname:"下次追踪日期" json:"nextTraceDate" bson:"nextTraceDate"`  //下次追踪日期
	CreateTime     string `uname:"录入时间" json:"createTime" bson:"createTime"`          //录入时间 新增时不需要输入，系统自动生成
	CreateUser     string `uname:"录入人员" json:"createUser" bson:"createUser"`          //录入人员
	CluePhone      string `uname:"线索电话" json:"cluePhone" bson:"cluePhone"`            //线索电话
	FirstCustomMgr string `uname:"首次分配人" json:"firstCustomMgr" bson:"firstCustomMgr"` //首次分配人
	AsssignedDate  string `uname:"分配时间" json:"asssignedDate" bson:"asssignedDate"`    //分配时间
	DealDate       string `uname:"成单日期" json:"dealDate" bson:"dealDate"`              //成单日期
	ClueCost       string `uname:"线索成本" json:"clueCost" bson:"clueCost"`              //线索成本
	AdviceInfo     string `uname:"咨询备注" json:"adviceInfo" bson:"adviceInfo"`          //咨询备注
	ContractInfo   string `uname:"联系方式备注" json:"contractInfo" bson:"contractInfo"`    //联系方式备注
	OrderNo        string `uname:"订单编号" json:"orderNo" bson:"orderNo"`                //订单编号
}

//企业线索表
type EntClue struct {
	CompanyName       string  `json:"companyName" bson:"companyName"`             //公司名称
	BusinessId        string  `json:"businessId" bson:"businessId"`               //营业执照
	MobilePhone       string  `json:"mobilePhone" bson:"mobilePhone"`             //手机
	PhoneStatus       string  `json:"phoneStatus" bson:"phoneStatus"`             //手机
	TelePhone         string  `json:"telePhone" bson:"telePhone"`                 //电话
	BuildDate         string  `json:"buildDate" bson:"buildDate"`                 //注册时间
	LegalPerson       string  `json:"legalPerson" bson:"legalPerson"`             //法人
	BuildCity         string  `json:"buildCity" bson:"buildCity"`                 //注册城市
	BuildRegion       string  `json:"buildRegion" bson:"buildRegion"`             //注册区
	BuildAddress      string  `json:"buildAddress" bson:"buildAddress"`           //公司地址
	RegistCapital     float64 `json:"registCapital" bson:"registCapital"`         //注册资本
	DomForFunds       string  `json:"domForFunds" bson:"domForFunds"`             //内外资
	CreateDate        string  `json:"createDate" bson:"createDate"`               //录入时间
	InfoFrom          string  `json:"infoFrom" bson:"infoFrom"`                   //信息来源
	InfoFlag          string  `json:"infoFlag" bson:"infoFlag"`                   //企业备注
	CompanyInfoStatus string  `json:"companyInfoStatus" bson:"companyInfoStatus"` //领取状态:是不是已经被领取，或者被丢弃
	WantAccount       string  `json:"wantAccount" bson:"wantAccount"`
	DrawDate          string  `json:"drawDate" bson:"drawDate"`             //领取时间
	CustomMgr         string  `json:"customMgr" bson:"customMgr"`           //领取客户经理
	CustomMgrPhone    string  `json:"customMgrPhone" bson:"customMgrPhone"` //领取客户经理
	Dept              string  `json:"dept" bson:"dept"`                     //领取部门
	NextTraceDate     string  `json:"nextTraceDate" bson:"nextTraceDate"`   //下次追踪日期
	TraceDate         string  `json:"traceDate" bson:"traceDate"`           //最近一次追踪时间
	TraceInfo         string  `json:"traceInfo" bson:"traceInfo"`           //最近一次工作记录
	DropInfo          string  `json:"dropInfo" bson:"dropInfo"`             //废弃原因
	DropReason        string  `json:"dropReason" bson:"dropReason"`         //废弃备注
}

//企业线索追踪
type EntTraceClue struct {
	CompanyName       string `json:"companyName" bson:"companyName"`             //公司名称
	BusinessId        string `json:"businessId" bson:"businessId"`               //营业执照
	PhoneStatus       string `json:"phoneStatus" bson:"phoneStatus"`             //手机
	CompanyInfoStatus string `json:"CompanyInfoStatus" bson:"CompanyInfoStatus"` //领取状态:空、未被领取、被领取、已转线索
	CustomMgr         string `json:"customMgr" bson:"customMgr"`                 //领取客户经理
	Dept              string `json:"dept" bson:"dept"`                           //领取部门
	NextTraceDate     string `json:"nextTraceDate" bson:"nextTraceDate"`         //下次追踪日期
	TraceDate         string `json:"traceDate" bson:"traceDate"`                 //最近一次追踪时间
	TraceInfo         string `json:"traceInfo" bson:"traceInfo"`                 //最近一次工作记录
	DropReason        string `json:"dropReason" bson:"dropReason"`               //废弃备注
}

type TraceUpdClue struct {
	ClueNo        string `json:"clueNo" bson:"clueNo"`
	ClueStatus    string `json:"clueStatus" bson:"clueStatus"`       //线索状态
	ClueType      string `json:"clueType" bson:"clueType"`           //线索类别
	NextTraceDate string `json:"nextTraceDate" bson:"nextTraceDate"` //下次追踪日期
	DropInfo      string `json:"dropInfo" bson:"dropInfo"`           //废弃原因
	TraceDate     string `json:"traceDate" bson:"traceDate"`         //最近一次追踪时间
	TraceInfo     string `json:"traceInfo" bson:"traceInfo"`         //最近一次工作记录
	DropReason    string `json:"dropReason" bson:"dropReason"`       //废弃备注
	CustomMgr     string `json:"customMgr" bson:"customMgr"`         //领取客户经理
	Dept          string `json:"dept" bson:"dept"`                   //领取部门
}

//线索追踪
type ClueTrace struct {
	ID            string `json:"id" bson:"id"`         //唯一id
	ClueNo        string `json:"clueNo" bson:"clueNo"` //线索编号
	CustomMgr     string `json:"customMgr" bson:"customMgr"`
	Dept          string `json:"dept" bson:"dept"`
	ClueStatus    string `json:"clueStatus" bson:"clueStatus"`       //线索状态
	ClueType      string `json:"clueType" bson:"clueType"`           //线索类别
	TraceInfo     string `json:"traceInfo" bson:"traceInfo"`         //线索编号
	DropInfo      string `json:"dropInfo" bson:"dropInfo"`           //线索编号
	DropReason    string `json:"dropReason" bson:"dropReason"`       //废弃备注
	NextTraceDate string `json:"nextTraceDate" bson:"nextTraceDate"` //下次跟踪时间
	CreateTime    string `json:"createTime" bson:"createTime"`       //录入时间
}

type RoleLimits struct {
	RoleName   string `json:"roleName" bson:"roleName"`
	PartName   string `json:"partName" bson:"partName"` //模块名称：比如 线索
	PartOpt    string `json:"partOpt" bson:"partOpt"`   //模块下属功能：比如 线索管理-新增
	CreateTime string `json:"createTime" bson:"createTime"`
}

//订单表
type Order struct {
	OrderNo           string  `uname:"order编号" json:"orderNo" bson:"orderNo"`                    //order编号
	OrderType         string  `uname:"order类型" json:"orderType" bson:"orderType"`                //order类型  注册 迁入 变更 注销 其他 商标 建筑资质
	OrderStatus       string  `uname:"order状态" json:"orderStatus" bson:"orderStatus"`            //order状态
	OrderDate         string  `uname:"order日期" json:"orderDate" bson:"orderDate"`                //order日期
	OrderInfo         string  `uname:"order信息" json:"orderInfo" bson:"orderInfo"`                //order信息
	OrderNotes        string  `uname:"order备注" json:"orderNotes" bson:"orderNotes"`              //order备注
	ClueNo            string  `uname:"order的线索编号" json:"clueNo" bson:"clueNo"`                   //order的线索编号
	ChannelType       string  `uname:"线索渠道分类" json:"channelType" bson:"channelType"`             //线索渠道分类
	SaleChannel       string  `uname:"线索营销渠道" json:"saleChannel" bson:"saleChannel"`             //线索营销渠道
	CustomMgr         string  `uname:"领取客户经理" json:"customMgr" bson:"customMgr"`                 //领取客户经理
	CustomMgrPhone    string  `uname:"领取客户经理" json:"customMgrPhone" bson:"customMgrPhone"`       //领取客户经理
	Dept              string  `uname:"领取部门" json:"dept" bson:"dept"`                             //领取部门
	CompanyName       string  `uname:"公司名称" json:"companyName" bson:"companyName"`               //公司名称
	BusinessScope     string  `uname:"公司经营范围" json:"businessScope" bson:"businessScope"`         //公司经营范围
	CompanyNotes      string  `uname:"公司的备注事项" json:"companyNotes" bson:"companyNotes"`          //公司的备注事项
	ExternalChannelNo string  `uname:"外部渠道编号" json:"externalChannelNo" bson:"externalChannelNo"` //外部渠道编号
	SocialCreditCode  string  `uname:"统一信用代码" json:"socialCreditCode" bson:"socialCreditCode"`   //统一信用代码
	Clientele         string  `uname:"委托人" json:"clientele" bson:"clientele"`                    //委托人
	ClientelePhone1   string  `uname:"委托人电话1" json:"clientelePhone1" bson:"clientelePhone1"`     //委托人电话1
	ClientelePhone2   string  `uname:"委托人电话2" json:"clientelePhone2" bson:"clientelePhone2"`     //委托人电话2
	ClienteleWechat   string  `uname:"委托人微信号" json:"clienteleWechat" bson:"clienteleWechat"`     //委托人微信号
	LegalPerson       string  `uname:"法人" json:"legalPerson" bson:"legalPerson"`                 //法人
	LegalPhone        string  `uname:"法人电话" json:"legalPhone" bson:"legalPhone"`                 //法人电话
	LegalCertid       string  `uname:"法人身份证号码" json:"legalCertid" bson:"legalCertid"`            //法人身份证号码
	Supervisor        string  `uname:"监事" json:"supervisor" bson:"supervisor"`                   //监事
	OwnershipRatio    string  `uname:"股权比例" json:"ownershipRatio" bson:"ownershipRatio"`         //股权比例
	RegistCapital     float64 `uname:"注册资本" json:"registCapital" bson:"registCapital"`           //注册资本
	PublicAnnDate     string  `uname:"公示公告日期" json:"publicAnnDate" bson:"publicAnnDate"`         //公示公告日期
	TaxStatus         string  `uname:"税务状态" json:"taxStatus" bson:"taxStatus"`                   //税务状态 /会计去税务局注销税务状态，会计外勤
	TaxMethod         string  `uname:"税务核算方式" json:"taxMethod" bson:"taxMethod"`                 //税务核算方式
	RelProject        string  `uname:"关联项目申请" json:"relProject" bson:"relProject"`               //关联项目申请  食品经营许可证 进出口权代理 道路运输许可证；可多选
	RegAddress        string  `uname:"注册地址" json:"regAddress" bson:"regAddress"`                 //注册地址
	RegAddressType    string  `uname:"注册地址是否是实际地址" json:"regAddressType" bson:"regAddressType"`  //注册地址是否是实际地址
	BuildRegion       string  `uname:"注册区" json:"buildRegion" bson:"buildRegion"`                //注册区 亭湖区、盐南高新区、开发区、盐都区、其他
	BuildRegionNote   string  `uname:"注册区备注" json:"buildRegionNote" bson:"buildRegionNote"`      //注册区备注
	AgentAccount      string  `uname:"是否代账" json:"agentAccount" bson:"agentAccount"`             //是否代账 是 否
	NilReturn         string  `uname:"是否零申报" json:"nilReturn" bson:"nilReturn"`                  //是否零申报
	RegFee            float64 `uname:"注册金额" json:"regFee" bson:"regFee"`                         //注册金额
	RegFrontFee       float64 `uname:"注册定金" json:"regFrontFee" bson:"regFrontFee"`               //注册定金
	AgentFee          float64 `uname:"代账费" json:"agentFee" bson:"agentFee"`                      //代账费
	AgentFrontFee     float64 `uname:"代账费定金" json:"agentFrontFee" bson:"agentFrontFee"`          //代账费定金
	AgentBeginDate    string  `uname:"代账开始日期" json:"agentBeginDate" bson:"agentBeginDate"`       //代账开始日期
	AgentFeeEndDate   string  `uname:"代账费截止日期" json:"agentFeeEndDate" bson:"agentFeeEndDate"`    //代账费截止日期
	AgentFeeType      string  `uname:"代账费用形式" json:"agentFeeType" bson:"agentFeeType"`           //代账费用形式 选项有：年付、月付、季付、半年付、两年付、三年付、四年付、其他
	IntelligenceType  string  `uname:"资质类型" json:"intelligenceType" bson:"intelligenceType"`     //资质类型
	ApplyDate         string  `uname:"申请日期" json:"applyDate" bson:"applyDate"`                   //申请日期
	BuildDate         string  `uname:"成立日期" json:"buildDate" bson:"buildDate"`                   //成立日期
	FeeDone           string  `uname:"费用是否已经补齐" json:"feeDone" bson:"feeDone"`                   //费用是否已经补齐
	AgentDone         string  `uname:"是否已经转为代账客户" json:"agentDone" bson:"agentDone"`             //是否已经转为代账客户
	OtherNote         string  `uname:"其他事项" json:"otherNote" bson:"otherNote"`                   //其他事项
	AgentFile         string  `uname:"委托表图片" json:"agentFile" bson:"agentFile"`                  //委托表图片
	ContractAddress   string  `uname:"签约地址" json:"contractAddress" bson:"contractAddress"`       //签约地址  公司门店 客户门店 网签 其他
	RelCustom         string  `uname:"关联客户" json:"relCustom" bson:"relCustom"`                   //关联客户
	ToAgent           string  `uname:"是否给中介" json:"toAgent" bson:"toAgent"`                      //是否给中介
	AgentInfo         string  `uname:"中介信息" json:"agentInfo" bson:"agentInfo"`                   //中介信息
	Coordinator       string  `uname:"统筹人员" json:"coordinator" bson:"coordinator"`               //统筹人员
	DocStaff          string  `uname:"制单员" json:"docStaff" bson:"docStaff"`                      //制单员
	BusOutWorker      string  `uname:"工商外勤" json:"busOutWorker" bson:"busOutWorker"`             //工商外勤
	TaxOutWorker      string  `uname:"税务外勤" json:"taxOutWorker" bson:"taxOutWorker"`             //税务外勤
	BankOutWorker     string  `uname:"银行外勤" json:"bankOutWorker" bson:"bankOutWorker"`           //银行外勤
	SocialOutWorker   string  `uname:"社保外勤" json:"socialOutWorker" bson:"socialOutWorker"`       //社保外勤
	FundOutWorker     string  `uname:"公积金外勤" json:"fundOutWorker" bson:"fundOutWorker"`          //公积金外勤
	OutWorker         string  `uname:"外勤" json:"outWorker" bson:"outWorker"`                     //外勤
	CreateTime        string  `uname:"创建时间" json:"createTime" bson:"createTime"`                 //创建时间
}

//订单跟进
type OrderEvent struct {
	OrderNo        string `json:"orderNo" bson:"orderNo"`               //order编号
	OrderEventType string `json:"orderEventType" bson:"orderEventType"` //orderevent类型
	EventNo        string `json:"eventNo" bson:"eventNo"`               //Event序号
	CompanyName    string `json:"companyName" bson:"companyName"`       //公司名称
	EventDate      string `json:"eventDate" bson:"eventDate"`           //办理日期
	EventDateZone  string `json:"eventDateZone" bson:"eventDateZone"`   //上午 or 下午
	OutStaff       string `json:"outStaff" bson:"outStaff"`             //办理人员
	Promoter       string `json:"promoter" bson:"promoter"`             //发起人
	PromotTime     string `json:"promotTime" bson:"promotTime"`         //发起时间
	EventType      string `json:"eventType" bson:"eventType"`           //eventType类型 工商制单、工商外勤、税务外勤、银行外勤、社保公积金外勤、其他外勤
	EventType2     string `json:"eventType2" bson:"eventType2"`         //子类型
	EventNote      string `json:"eventNote" bson:"eventNote"`           //事项备注
	AcceptDoc      string `json:"acceptDoc" bson:"acceptDoc"`           //接受资料
	WorkArea       string `json:"workArea" bson:"workArea"`             //办理区域
	WorkAreaNote   string `json:"workAreaNote" bson:"workAreaNote"`     //办理区域备注
	WorkSite       string `json:"workSite" bson:"workSite"`             //办理地点
	EventStatus    string `json:"eventStatus" bson:"eventStatus"`       //办理进度 已完成、废弃
	Reason         string `json:"reason" bson:"reason"`                 //未完成原因
	LastWorkTime   string `json:"lastWorkTime" bson:"lastWorkTime"`     //最近工作时间
	CreateTime     string `json:"createTime" bson:"createTime"`
}

//代账客户表
type AgentCustom struct {
	CustomNo              string  `uname:"客户编号" json:"customNo" bson:"customNo"`                               //客户编号
	CompanyName           string  `uname:"公司名称" json:"companyName" bson:"companyName"`                         //公司名称
	CustomLevel           string  `uname:"客户等级" json:"customLevel" bson:"customLevel"`                         //客户等级
	AgentStatus           string  `uname:"代账状态" json:"agentStatus" bson:"agentStatus"`                         //代账状态
	AgentStatusChangeTime string  `uname:"代账状态变更时间" json:"agentStatusChangeTime" bson:"agentStatusChangeTime"` //代账状态变更时间
	TaxMethod             string  `uname:"税务核算方式" json:"taxMethod" bson:"taxMethod"`                           //税务核算方式
	NilReturn             string  `uname:"是否零申报" json:"nilReturn" bson:"nilReturn"`                            //是否零申报
	InType                string  `uname:"转化来源" json:"inType" bson:"inType"`                                   //转化来源
	CustomDate            string  `uname:"生成代账客户时间" json:"customDate" bson:"customDate"`                       //生成代账客户时间
	Clientele             string  `uname:"委托人" json:"clientele" bson:"clientele"`                              //委托人
	ClientelePhone1       string  `uname:"委托人电话1" json:"clientelePhone1" bson:"clientelePhone1"`               //委托人电话1
	ClientelePhone2       string  `uname:"委托人电话2" json:"clientelePhone2" bson:"clientelePhone2"`               //委托人电话2
	ClienteleWechat       string  `uname:"委托人微信号" json:"clienteleWechat" bson:"clienteleWechat"`               //委托人微信号
	LegalPerson           string  `uname:"法人" json:"legalPerson" bson:"legalPerson"`                           //法人
	LegalPhone            string  `uname:"法人电话" json:"legalPhone" bson:"legalPhone"`                           //法人电话
	LegalCertid           string  `uname:"法人身份证号码" json:"legalCertid" bson:"legalCertid"`                      //法人身份证号码
	OrderNo               string  `uname:"订单编号" json:"orderNo" bson:"orderNo"`                                 //订单编号
	AgentFeeType          string  `uname:"代账费用形式" json:"agentFeeType" bson:"agentFeeType"`                     //代账费用形式
	AgentFeeNote          string  `uname:"代账费支付周期备注" json:"agentFeeNote" bson:"agentFeeNote"`                  //代账费支付周期备注
	AgentFee              float64 `uname:"代账费" json:"agentFee" bson:"agentFee"`                                //代账费
	CustomMgr             string  `uname:"领取客户经理" json:"customMgr" bson:"customMgr"`                           //领取客户经理
	CustomMgrPhone        string  `uname:"领取客户经理电话" json:"customMgrPhone" bson:"customMgrPhone"`               //领取客户经理电话
	ServiceSite           string  `uname:"服务地址" json:"serviceSite" bson:"serviceSite"`                         //服务地址
	Accountant            string  `uname:"主管会计" json:"accountant" bson:"accountant"`                           //主管会计
	AccountantChangeCount string  `uname:"主管会计更换次数" json:"accountantChangeCount" bson:"accountantChangeCount"` //主管会计更换次数
	AccountantChangeTime  string  `uname:"最近主管会计更换时间" json:"accountantChangeTime" bson:"accountantChangeTime"` //最近主管会计更换时间
	EntryAccountant       string  `uname:"记账会计" json:"entryAccountant" bson:"entryAccountant"`                 //记账会计
	BillingAccountant     string  `uname:"开票会计" json:"billingAccountant" bson:"billingAccountant"`             //开票会计
	AgentBeginDate        string  `uname:"代账开始日期" json:"agentBeginDate" bson:"agentBeginDate"`                 //代账开始日期
	AgentFeeEndDate       string  `uname:"代账费截止日期" json:"agentFeeEndDate" bson:"agentFeeEndDate"`              //代账费截止日期
	TaxControl            string  `uname:"税盘控" json:"taxControl" bson:"taxControl"`                            //税盘控
	TaxPlateType          string  `uname:"税盘种类" json:"taxPlateType" bson:"taxPlateType"`                       //税盘种类
	TaxControlSite        string  `uname:"税控地址" json:"taxControlSite" bson:"taxControlSite"`                   //税控地址
	RelCustom             string  `uname:"与哪家企业是一家的" json:"relCustom" bson:"relCustom"`                        //与哪家企业是一家的
	AgentAgree            string  `uname:"代账合同是否已入库" json:"agentAgree" bson:"agentAgree"`                      //代账合同是否已入库
	BussYearReport        string  `uname:"工商年报是否已完成" json:"bussYearReport" bson:"bussYearReport"`              //工商年报是否已完成
	BussYearReportMgr     string  `uname:"工商年报人员" json:"bussYearReportMgr" bson:"bussYearReportMgr"`           //工商年报人员
	BussYearReportDate    string  `uname:"工商年报完成日期" json:"bussYearReportDate" bson:"bussYearReportDate"`       //工商年报完成日期
	SocialCreditCode      string  `uname:"税号" json:"socialCreditCode" bson:"socialCreditCode"`                 //税号
	TaxOffice             string  `uname:"税收地址" json:"taxOffice" bson:"taxOffice"`                             //税收地址
	TaxOfficeNote         string  `uname:"主管税局备注" json:"taxOfficeNote" bson:"taxOfficeNote"`                   //主管税局备注
	SocialInsurance       string  `uname:"是否有社保" json:"socialInsurance" bson:"socialInsurance"`                //是否有社保
	BuildDate             string  `uname:"工商成立日期" json:"buildDate" bson:"buildDate"`                           //工商成立日期
	RegisteNo             string  `uname:"工商注册号" json:"registeNo" bson:"registeNo"`                            //工商注册号
	RegAddress            string  `uname:"注册地址" json:"regAddress" bson:"regAddress"`                           //注册地址
	RegAddressType        string  `uname:"注册地址是否是实际地址" json:"regAddressType" bson:"regAddressType"`            //注册地址是否是实际地址
	BuildRegion           string  `uname:"注册区" json:"buildRegion" bson:"buildRegion"`                          //注册区
	BuildRegionNote       string  `uname:"注册区备注" json:"buildRegionNote" bson:"buildRegionNote"`                //注册区备注
	RegistCapital         string  `uname:"注册资本" json:"registCapital" bson:"registCapital"`                     //注册资本
	CheckYear             string  `uname:"年检年份" json:"checkYear" bson:"checkYear"`                             //年检年份
	CheckStatus           string  `uname:"年检状态" json:"checkStatus" bson:"checkStatus"`                         //年检状态
	CheckCause            string  `uname:"年检未完成原因" json:"checkCause" bson:"checkCause"`                        //年检未完成原因
	CheckDate             string  `uname:"年检日期" json:"checkDate" bson:"checkDate"`                             //年检日期
	Checker               string  `uname:"年检专员" json:"checker" bson:"checker"`                                 //年检专员
	CreateTime            string  `uname:"录入时间" json:"createTime" bson:"createTime"`                           //录入时间
	ServiceMon            string  `uname:"服务月份" json:"serviceMon" bson:"serviceMon"`                           //服务月份
	TaxReturnSts          string  `uname:"抄报税状态" json:"taxReturnSts" bson:"taxReturnSts"`                      //抄报税状态
	ClearCard             string  `uname:"清卡状态" json:"clearCard" bson:"clearCard"`                             //清卡状态
	AccountantSts         string  `uname:"是否需要做账" json:"accountantSts" bson:"accountantSts"`                   //是否需要做账
	PaySheet              string  `uname:"工资表" json:"paySheet" bson:"paySheet"`                                //工资表
	Statreceipt           string  `uname:"对账单回单" json:"statreceipt" bson:"statreceipt"`                        //对账单回单
	IncomeSts             string  `uname:"进项是否已确认" json:"incomeSts" bson:"incomeSts"`                          //进项是否已确认
	EstimatedCost         string  `uname:"暂估成本" json:"estimatedCost" bson:"estimatedCost"`                     //暂估成本
	FrontIncomeTax        string  `uname:"拟缴纳企业所得税金额" json:"frontIncomeTax" bson:"frontIncomeTax"`             //拟缴纳企业所得税金额
	AccountingInfor       string  `uname:"做账资料是否已全" json:"accountingInfor" bson:"accountingInfor"`             //做账资料是否已全
	AccountingSts         string  `uname:"记账状态" json:"accountingSts" bson:"accountingSts"`                     //记账状态
	ReportTax             string  `uname:"报税状态" json:"reportTax" bson:"reportTax"`                             //报税状态
	CertiLoadSts          string  `uname:"凭证装订状态" json:"certiLoadSts" bson:"certiLoadSts"`                     //凭证装订状态
	OperatingIncome       float64 `uname:"营业收入" json:"operatingIncome" bson:"operatingIncome"`                 //营业收入
	ValueAddedTax         float64 `uname:"增值税" json:"valueAddedTax" bson:"valueAddedTax"`                      //增值税
	BusinessIncomeTax     float64 `uname:"企业所得税" json:"businessIncomeTax" bson:"businessIncomeTax"`            //企业所得税
	AdditionalTax         float64 `uname:"附加税" json:"additionalTax" bson:"additionalTax"`                      //附加税
	StampTax              float64 `uname:"印花税" json:"stampTax" bson:"stampTax"`                                //印花税
	PersonalIncomeTax     float64 `uname:"个税" json:"personalIncomeTax" bson:"personalIncomeTax"`               //个税
	TaxAll                float64 `uname:"总税款" json:"taxAll" bson:"taxAll"`                                    //总税款
	CertiSts              string  `uname:"凭证状态" json:"certiSts" bson:"certiSts"`                               //凭证状态
	CertiCount            int     `uname:"凭证数量" json:"certiCount" bson:"certiCount"`                           //凭证数量
	CertiRegister         string  `uname:"凭证登记人员" json:"certiRegister" bson:"certiRegister"`                   //凭证登记人员
	CertiRegTime          string  `uname:"凭证登记时间" json:"certiRegTime" bson:"certiRegTime"`                     //凭证登记时间
	CertiInPersonnel      string  `uname:"凭证入库人员" json:"certiInPersonnel" bson:"certiInPersonnel"`             //凭证入库人员
	CertiInTime           string  `uname:"凭证入库时间" json:"certiInTime" bson:"certiInTime"`                       //凭证入库时间
	CertiTemporary        string  `uname:"凭证临时领用人" json:"certiTemporary" bson:"certiTemporary"`                //凭证临时领用人
	CertiOutPersonnel     string  `uname:"凭证出库人员" json:"certiOutPersonnel" bson:"certiOutPersonnel"`           //凭证出库人员
	CertiOutTime          string  `uname:"凭证出库时间" json:"certiOutTime" bson:"certiOutTime"`                     //凭证出库时间
	CertiHandover         string  `uname:"凭证交接人" json:"certiHandover" bson:"certiHandover"`                    //凭证交接人
	CertiHandoverNote     string  `uname:"凭证交接备注" json:"certiHandoverNote" bson:"certiHandoverNote"`           //凭证交接备注
	CertiHandoverTime     string  `uname:"凭证交接时间" json:"certiHandoverTime" bson:"certiHandoverTime"`           //凭证交接时间

}

type CertiInfo struct {
	ServiceMon        string `json:"serviceMon" bson:"serviceMon"`               //服务月份 年月格式，每月1日至31日，因报税所属期为上个月，所以此处是上个月月份，比方说，现在是2020年11月，但是11月报的是10月的税，采集10月数据，所以此处显示：2020-10
	CustomNo          string `json:"customNo" bson:"customNo"`                   //从代账客户表获取 custom编号
	CompanyName       string `json:"companyName" bson:"companyName"`             //从代账客户表获取 公司名称
	CertiSts          string `json:"certiSts" bson:"certiSts"`                   //凭证状态 选项有：凭证已登记、凭证已入库、凭证待出库、凭证已交接
	CertiCount        int    `json:"certiCount" bson:"certiCount"`               //凭证数量 数字输入框，整数格式；
	CertiRegister     string `json:"certiRegister" bson:"certiRegister"`         //凭证登记人员
	CertiRegTime      string `json:"certiRegTime" bson:"certiRegTime"`           //凭证登记时间
	CertiInPersonnel  string `json:"certiInPersonnel" bson:"certiInPersonnel"`   //凭证入库人员
	CertiInTime       string `json:"certiInTime" bson:"certiInTime"`             //凭证入库时间
	CertiTemporary    string `json:"certiTemporary" bson:"certiTemporary"`       //凭证临时领用人
	CertiOutPersonnel string `json:"certiOutPersonnel" bson:"certiOutPersonnel"` //凭证出库人员
	CertiOutTime      string `json:"certiOutTime" bson:"certiOutTime"`           //凭证出库时间
	CertiHandover     string `json:"certiHandover" bson:"certiHandover"`         //凭证交接人
	CertiHandoverNote string `json:"certiHandoverNote" bson:"certiHandoverNote"` //凭证交接备注
	CertiHandoverTime string `json:"certiHandoverTime" bson:"certiHandoverTime"` //凭证交接时间
}

//月度报税
type TaxRept struct {
	ServiceMon        string  `json:"serviceMon" bson:"serviceMon"`               //服务月份 年月格式，每月1日至31日，因报税所属期为上个月，所以此处是上个月月份，比方说，现在是2020年11月，但是11月报的是10月的税，采集10月数据，所以此处显示：2020-10
	CustomNo          string  `json:"customNo" bson:"customNo"`                   //从代账客户表获取 custom编号
	CompanyName       string  `json:"companyName" bson:"companyName"`             //从代账客户表获取 公司名称
	SocialCreditCode  string  `json:"socialCreditCode" bson:"socialCreditCode"`   //从代账客户表获取 税号
	AgentStatus       string  `json:"agentStatus" bson:"agentStatus"`             //从代账客户表获取 代账状态分为：正常、本司注销、已转走、停报，新客户默认‘正常’
	TaxMethod         string  `json:"taxMethod" bson:"taxMethod"`                 //从代账客户表获取 税务核算方式 一般纳税人为A、小规模纳税人为B、个体为C
	Accountant        string  `json:"accountant" bson:"accountant"`               //从代账客户表获取 主管会计
	EntryAccountant   string  `json:"entryAccountant" bson:"entryAccountant"`     //从代账客户表获取 记账会计
	TaxControl        string  `json:"taxControl" bson:"taxControl"`               //从代账客户表获取 税盘控选项：是、否；生成代账客户时选择或者后期修改
	TaxControlSite    string  `json:"taxControlSite" bson:"taxControlSite"`       //从代账客户表获取 选项：法治广场、新龙广场、客户
	TaxReturnSts      string  `json:"taxReturnSts" bson:"taxReturnSts"`           //抄报税状态  月度报税编辑页面下拉选择，选项有：抄报税已完成、抄报税异常、无需抄报税；非必填项。
	ClearCard         string  `json:"clearCard" bson:"clearCard"`                 //清卡状态  月度报税编辑页面下拉选择，选项有：清卡已完成、清卡异常、无需清卡；非必填项。
	AccountantSts     string  `json:"accountantSts" bson:"accountantSts"`         //是否需要做账  月度报税编辑页面下拉选择，选项有：需要做账、无需做账；非必填项。
	PaySheet          string  `json:"paySheet" bson:"paySheet"`                   //工资表  月度报税编辑页面下拉选择，选项有：和上个月一样、已发钉钉、无需做工资；非必选项。
	Statreceipt       string  `json:"statreceipt" bson:"statreceipt"`             //对账单回单 月度报税编辑页面下拉选择，选项有：有银行对账单、无银行对账单；非必选项。
	IncomeSts         string  `json:"incomeSts" bson:"incomeSts"`                 //进项是否已确认 月度报税编辑页面下拉选择，选项有：进项已确认、无需确认；非必选项。
	EstimatedCost     string  `json:"estimatedCost" bson:"estimatedCost"`         //暂估成本  编辑页面文本格式，非必填项
	FrontIncomeTax    string  `json:"frontIncomeTax" bson:"frontIncomeTax"`       //拟缴纳企业所得税金额  编辑页面文本格式，非必填项
	AccountingInfor   string  `json:"accountingInfor" bson:"accountingInfor"`     //做账资料是否已全 月度报税编辑页面下拉选择，选项有：做账资料已全、无需记账；非必选项。
	AccountingSts     string  `json:"accountingSts" bson:"accountingSts"`         //记账状态 月度报税编辑页面下拉选择，选项有：记账已完成、无需记账；非必选项。
	ReportTax         string  `json:"reportTax" bson:"reportTax"`                 //报税状态 月度报税编辑页面下拉选择，选项有：报税已完成、不需要报税；非必选项。
	CertiLoadSts      string  `json:"certiLoadSts" bson:"certiLoadSts"`           //凭证装订状态  不需要报税、报税已完成
	OperatingIncome   float64 `json:"operatingIncome" bson:"operatingIncome"`     //营业收入  编辑页面输入，数字，有小数点，非必填项
	ValueAddedTax     float64 `json:"valueAddedTax" bson:"valueAddedTax"`         //增值税 编辑页面输入，数字，有小数点，非必填项
	BusinessIncomeTax float64 `json:"businessIncomeTax" bson:"businessIncomeTax"` //企业所得税 编辑页面输入，数字，有小数点，非必填项
	AdditionalTax     float64 `json:"additionalTax" bson:"additionalTax"`         //附加税 编辑页面输入，数字，有小数点，非必填项
	StampTax          float64 `json:"stampTax" bson:"stampTax"`                   //印花税 编辑页面输入，数字，有小数点，非必填项
	PersonalIncomeTax float64 `json:"personalIncomeTax" bson:"personalIncomeTax"` //个税 编辑页面输入，数字，有小数点，非必填项
	TaxAll            float64 `json:"taxAll" bson:"taxAll"`                       //总税款 编辑页面输入，数字，有小数点，非必填项
	CreateTime        string  `json:"createTime" bson:"createTime"`               //录入时间
}

//收费表
type Collection struct {
	CollId          string  `uname:"序号" json:"collId" bson:"collId"`                        //序号 新增时系统自动生成 更新时代入
	CollFromType    string  `uname:"收款来源类型" json:"collFromType" bson:"collFromType"`        //收款来源类型           订单、客户（用来区分是订单收款还是客户收款）
	CollNo          string  `uname:"收款编号" json:"collNo" bson:"collNo"`                      //收款编号  订单编号或者代账客户管理里面的客户编号
	CollFrom        string  `uname:"收款来源" json:"collFrom" bson:"collFrom"`                  //收款来源   根据订单类别分别显示为：代账客户、注册代账订单、迁入订单、变更订单、注销订单、商标订单、建筑资质订单、其他订单
	CollTime        string  `uname:"收款时间" json:"collTime" bson:"collTime"`                  //收款时间    出纳录单时的年月日格式
	OrderTime       string  `uname:"订单日期" json:"orderTime" bson:"orderTime"`                //订单日期
	CollOutline     string  `uname:"公司名称+法人" json:"collOutline" bson:"collOutline"`         //公司名称+法人
	CompanyName     string  `uname:"公司名称" json:"companyName" bson:"companyName"`            //公司名称
	LegalPerson     string  `uname:"法人" json:"legalPerson" bson:"legalPerson"`              //法人
	CollType        string  `uname:"收款类型" json:"collType" bson:"collType"`                  //收款类型
	CollVal         float32 `uname:"收款金额" json:"collVal" bson:"collVal"`                    //收款金额
	CollKind        string  `uname:"收款方式" json:"collKind" bson:"collKind"`                  //收款方式      必选选项，选项分为：建行龙支付、纳川对公、润盈对公、氢数对公、好易标对公、润禾对公、企业支付宝、其他账户
	Commission      float32 `uname:"手续费" json:"commission" bson:"commission"`               //手续费 非必填项，手动输入
	Accounter       string  `uname:"报账人" json:"accounter" bson:"accounter"`                 //报账人   必选选项，下拉按部门显示公司所有人员（客户经理、主管会计角色人员显示在上面）
	AgentBeginDate  string  `uname:"代账开始日期" json:"agentBeginDate" bson:"agentBeginDate"`    //代账开始日期
	AgentFeeEndDate string  `uname:"代账费截止日期" json:"agentFeeEndDate" bson:"agentFeeEndDate"` //代账费截止日期
	CollNote        string  `uname:"收款备注" json:"collNote" bson:"collNote"`                  //收款备注
	TransfeTime     string  `uname:"转账时间" json:"transfeTime" bson:"transfeTime"`            //转账时间
	CollStatus      string  `uname:"转账状态" json:"collStatus" bson:"collStatus"`              //转账状态  出纳录入成功显示“出纳已确认”、内账会计角色确认ok则显示“内账会计已确认”
	Accountant      string  `uname:"主管会计" json:"accountant" bson:"accountant"`              //主管会计
	CustomMgr       string  `uname:"客户经理" json:"customMgr" bson:"customMgr"`                //客户经理
}

//付费成本表
type Cost struct {
	CostId       string  `json:"costId" bson:"costId"`             //序号，新增时系统自动生成，更新时代入
	CostFromType string  `json:"costFromType" bson:"costFromType"` //成本来源类型  订单、客户（用来区分是订单收款还是客户收款）
	CostNo       string  `json:"costNo" bson:"costNo"`             //成本编号  订单编号或者代账客户管理里面的客户编号
	CostFrom     string  `json:"costFrom" bson:"costFrom"`         //成本来源 根据订单类别分别显示为：代账客户、注册代账订单、迁入订单、变更订单、注销订单、商标订单、建筑资质订单、其他订单
	CostTime     string  `json:"costTime" bson:"costTime"`         //成本时间 出纳录单时的年月日格式
	OrderTime    string  `json:"orderTime" bson:"orderTime"`       //订单里面的“订单日期”字段信息，如果是代账客户则不显示即可
	CostOutline  string  `json:"costOutline" bson:"costOutline"`   //公司名称+法人的格式，如：“盐城逍遥鹤新材料有限公司-法人【高源】”
	CompanyName  string  `json:"companyName" bson:"companyName"`   //公司名称
	LegalPerson  string  `json:"legalPerson" bson:"legalPerson"`   //法人
	CostType     string  `json:"costType" bson:"costType"`         //成本类型 营销成本、地址成本、本司提成成本、中介成本、刻章成本、政府规费、广告制作成本、其他成本
	CostVal      float32 `json:"costVal" bson:"costVal"`           //成本金额
	Accounter    string  `json:"accounter" bson:"accounter"`       //报账人 必选选项，下拉按部门显示公司所有人员（客户经理、主管会计角色人员显示在上面）
	CostNote     string  `json:"costNote" bson:"costNote"`         //成本备注
	TransfeTime  string  `json:"transfeTime" bson:"transfeTime"`   //费用使用时间
	CostStatus   string  `json:"costStatus" bson:"costStatus"`     //转账状态 出纳录入成功显示“出纳已确认”、内账会计角色确认ok则显示“内账会计已确认”
	Accountant   string  `json:"accountant" bson:"accountant"`     //主管会计
	CustomMgr    string  `json:"customMgr" bson:"customMgr"`       //客户经理
}

// 请求日志
type Log struct {
	User      string `json:"user" bson:"user"` // 任务名字
	AccessUrl string `json:"accessUrl" bson:"accessUrl"`
	Content   string `json:"content" bson:"content"` // 任务地址
	Time      string `json:"time" bson:"time"`       // 任务执行用户
}

type UserFilter struct {
	PhoneNo  string `bson:"phoneNo"`  //手机号码
	PassWord string `bson:"passWord"` // 密码
}

type UserPhoneFilter struct {
	PhoneNo string `bson:"phoneNo"` //手机号码
}
type UserDeptFilter struct {
	UserDept string `bson:"userDept"` //手机号码
}

type ClueIdFilter struct {
	ClueNo string `bson:"clueNo"` //手机号码
}

type OrderNoFilter struct {
	OrderNo string `bson:"orderNo"`
}

type OrderIdFilter struct {
	OrderNo string `bson:"orderNo"` //手机号码
}

type CustomNoFilter struct {
	CustomNo string `bson:"customNo"` //手机号码
}

type CertiFilter struct {
	CustomNo   string `bson:"customNo"` //
	ServiceMon string `bson:"serviceMon"`
}

type CustomTaxFilter struct {
	CustomNo   string `bson:"customNo"` //
	ServiceMon string `bson:"serviceMon"`
}

type CustomOrderNoFilter struct {
	OrderNo string `bson:"orderNo"` //手机号码
}

type OrderEventFilter struct {
	OrderNo string `bson:"orderNo"`
	EventNo string `bson:"eventNo"`
}

type ClueTraceIdFilter struct {
	ID string `bson:"id"` //id
}

type RoleFilter struct {
	RoleName string `bson:"roleName"`
	PartName string `bson:"partName"` //角色名称
	PartOpt  string `bson:"partOpt"`
}

type RoleByNameFilter struct {
	RoleName string `bson:"roleName"`
}

type UserPhoneAdminFilter struct {
	PhoneNo  string `bson:"phoneNo"`  //手机号码
	UserType string `bson:"userType"` //用户类型
}

type UserSaveFilter struct {
	PhoneNo  string `bson:"phoneNo"`  //手机号码
	PassWord string `bson:"passWord"` // 密码
	UserId   string `bson:"userId"`   //id
	UserName string `bson:"userName"` //用户名
}

type SortLogByClueNo struct {
	SortOrder int `bson:"_id"` //按
}

type SortById struct {
	SortOrder int `bson:"_id"` //按
}
type SortEntClue struct {
	SortOrder int `bson:"buildDate"` //按
}

//
func BuildResponse(errno int, msg string, data interface{}) (resp []byte, err error) {
	var (
		response Response
	)

	response.Errno = errno
	response.Msg = msg
	response.Data = data

	//给resp 和 err赋值
	resp, err = json.Marshal(response)

	return
}

func BuildStr2Time(timeStr string) (nexttime time.Time, err error) {

	var (
		timeLayout = "2006-01-02 15:04:05"
		loc, _     = time.LoadLocation("Local")
	)
	nexttime, err = time.ParseInLocation(timeLayout, timeStr, loc)

	return
}

func BuildTimeNowStr() (timeStr string) {

	timeStr = time.Now().Format("2006-01-02 15:04:05")
	return
}
func BuildUUID() (uuid string) {
	uuid = ksuid.New().String()
	return
}

func BuildDateMonStr2() (timeStr string) {

	timeStr = time.Now().Format("2006-01-02")[0:7]
	return
}
func BuildDateTodayStr2() (timeStr string) {

	timeStr = time.Now().Format("2006-01-02")
	return
}

// func BuildDateTodayStr4() (timeStr string) {
// 	timeStr = time.Now().Format("2006/01/02")
// 	fmt.Println(timeStr)
// 	return
// }
// func BuildDateMonStr4() (timeStr string) {

// 	timeStr = time.Now().Format("2006/01")
// 	fmt.Println(timeStr)
// 	return
// }

func BuildDateTodayStr() (timeStr string) {

	timeStr = time.Now().Format("20060102")[2:]
	return
}

func BuildDateTodayStr4() (timeStr string) {

	timeStr = time.Now().Format("20060102")
	return
}

func BuildDateTodayMonStr() (timeStr string) {

	timeStr = time.Now().Format("20060102")[2:6]
	return
}

func BuildDateTodayStr3() (timeStr string) {

	timeStr = time.Now().Format("20060102")[0:4]
	return
}

func BuildDateLast2Monstr() (timeStr string) {

	now := time.Now()
	yesterMon := now.AddDate(0, -2, 0)
	timeStr = yesterMon.Format("2006-01")
	// timeStr = time.Now().Format("20060102")[0:4]
	return
}

func BuildDateLast1Monstr() (timeStr string) {

	now := time.Now()
	yesterMon := now.AddDate(0, -1, 0)
	timeStr = yesterMon.Format("2006-01")
	// timeStr = time.Now().Format("20060102")[0:4]
	return
}

func BuildDateMore1Monstr() (timeStr string) {

	now := time.Now()
	yesterMon := now.AddDate(0, 1, 0)
	timeStr = yesterMon.Format("2006-01")
	// timeStr = time.Now().Format("20060102")[0:4]
	return
}

func BuildDateMonstr() (timeStr string) {

	now := time.Now()
	yesterMon := now.AddDate(0, 0, 0)
	timeStr = yesterMon.Format("2006-01")
	// timeStr = time.Now().Format("20060102")[0:4]
	return
}

// 获取本机网卡IP
func GetLocalIP() (ipv4 string, err error) {
	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP地址
		isIpNet bool
	)
	// 获取所有网卡
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}
	// 取第一个非lo的网卡IP
	for _, addr = range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				return
			}
		}
	}

	err = ERR_NO_LOCAL_IP_FOUND
	return
}

// func main() {

// 	fmt.Printf("已分配制单")
// }
