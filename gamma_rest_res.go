package mypolymarketapi

// GammaListEventsRes 是 List events 接口的响应体。
type GammaListEventsRes []GammaEvent

// GammaListEventsKeysetRes 是 List events（keyset pagination）接口的响应体。
type GammaListEventsKeysetRes struct {
	Events     []GammaEvent `json:"events"`      // 返回的事件列表；未命中时为空数组
	NextCursor string       `json:"next_cursor"` // 下一页游标；存在时可用于继续翻页
}

// GammaListMarketsRes 是 List markets 接口的响应体。
type GammaListMarketsRes []GammaMarket

// GammaListMarketsKeysetRes 是 List markets（keyset pagination）接口的响应体。
type GammaListMarketsKeysetRes struct {
	Markets    []GammaMarket `json:"markets"`     // 返回的市场列表；未命中时为空数组
	NextCursor string        `json:"next_cursor"` // 下一页游标；存在时可用于继续翻页
}

// GammaGetEventByIDRes 是 Get event by id 接口的响应体。
type GammaGetEventByIDRes GammaEvent

// GammaGetMarketByIDRes 是 Get market by id 接口的响应体。
type GammaGetMarketByIDRes GammaMarket

// GammaGetMarketBySlugRes 是 Get market by slug 接口的响应体（Gamma `GET /markets/slug/{slug}`）。
type GammaGetMarketBySlugRes GammaMarket

// GammaGetMarketTagsRes 是 Get market tags by id 接口的响应体（Gamma `GET /markets/{id}/tags`）。
type GammaGetMarketTagsRes []GammaTag

// GammaGetEventBySlugRes 是 Get event by slug 接口的响应体。
type GammaGetEventBySlugRes GammaEvent

// GammaGetEventTagsRes 是 Get event tags 接口的响应体。
type GammaGetEventTagsRes []GammaTag

// GammaEvent 表示事件对象。
type GammaEvent struct {
	ID                           string                  `json:"id"`                           // 事件 ID
	Ticker                       string                  `json:"ticker"`                       // 事件代码/简称
	Slug                         string                  `json:"slug"`                         // URL 友好的事件标识
	Title                        string                  `json:"title"`                        // 事件标题
	Subtitle                     string                  `json:"subtitle"`                     // 事件副标题
	Description                  string                  `json:"description"`                  // 事件描述
	ResolutionSource             string                  `json:"resolutionSource"`             // 事件的结算来源
	StartDate                    string                  `json:"startDate"`                    // 事件开始时间
	CreationDate                 string                  `json:"creationDate"`                 // 事件创建时间
	EndDate                      string                  `json:"endDate"`                      // 事件结束时间
	Image                        string                  `json:"image"`                        // 事件图片地址
	Icon                         string                  `json:"icon"`                         // 事件图标地址
	Active                       bool                   `json:"active"`                       // 是否处于激活状态
	Closed                       bool                   `json:"closed"`                       // 是否已关闭
	Archived                     bool                   `json:"archived"`                     // 是否已归档
	New                          bool                   `json:"new"`                          // 是否为新事件
	Featured                     bool                   `json:"featured"`                     // 是否为精选事件
	Restricted                   bool                   `json:"restricted"`                   // 是否受限
	Liquidity                    float64                `json:"liquidity"`                    // 事件总流动性
	Volume                       float64                `json:"volume"`                       // 事件总成交量
	OpenInterest                 float64                `json:"openInterest"`                 // 未平仓量
	SortBy                       string                  `json:"sortBy"`                       // 排序字段
	Category                     string                  `json:"category"`                     // 主分类
	Subcategory                  string                  `json:"subcategory"`                  // 子分类
	IsTemplate                   bool                   `json:"isTemplate"`                   // 是否为模板事件
	TemplateVariables            string                  `json:"templateVariables"`            // 模板变量
	PublishedAt                  string                  `json:"published_at"`                 // 发布时间
	CreatedBy                    string                  `json:"createdBy"`                    // 创建者
	UpdatedBy                    string                  `json:"updatedBy"`                    // 更新者
	CreatedAt                    string                  `json:"createdAt"`                    // 记录创建时间
	UpdatedAt                    string                  `json:"updatedAt"`                    // 记录更新时间
	CommentsEnabled              bool                   `json:"commentsEnabled"`              // 是否允许评论
	Competitive                  float64                `json:"competitive"`                  // 竞争度相关数值
	Volume24hr                   float64                `json:"volume24hr"`                   // 24 小时成交量
	Volume1wk                    float64                `json:"volume1wk"`                    // 1 周成交量
	Volume1mo                    float64                `json:"volume1mo"`                    // 1 月成交量
	Volume1yr                    float64                `json:"volume1yr"`                    // 1 年成交量
	FeaturedImage                string                  `json:"featuredImage"`                // 精选展示图地址
	DisqusThread                 string                  `json:"disqusThread"`                 // Disqus 讨论串 ID
	ParentEvent                  string                  `json:"parentEvent"`                  // 父事件 ID
	EnableOrderBook              bool                   `json:"enableOrderBook"`              // 是否启用订单簿
	LiquidityAmm                 float64                `json:"liquidityAmm"`                 // AMM 流动性
	LiquidityClob                float64                `json:"liquidityClob"`                // CLOB 流动性
	NegRisk                      bool                   `json:"negRisk"`                      // 是否启用 neg risk
	NegRiskMarketID              string                  `json:"negRiskMarketID"`              // neg risk 关联市场 ID
	NegRiskFeeBips               int64                  `json:"negRiskFeeBips"`               // neg risk 费用基点
	CommentCount                 int64                  `json:"commentCount"`                 // 评论数
	ImageOptimized               GammaImageOptimization `json:"imageOptimized"`               // 图片优化信息
	IconOptimized                GammaImageOptimization `json:"iconOptimized"`                // 图标优化信息
	FeaturedImageOptimized       GammaImageOptimization `json:"featuredImageOptimized"`       // 精选图优化信息
	SubEvents                    []string                `json:"subEvents"`                    // 子事件 ID 列表
	Markets                      []GammaMarket           `json:"markets"`                      // 事件下的市场列表
	Series                       []GammaSeries           `json:"series"`                       // 关联的系列列表
	Categories                   []GammaCategory         `json:"categories"`                   // 关联的分类列表
	Collections                  []GammaCollection       `json:"collections"`                  // 所属合集列表
	Tags                         []GammaTag              `json:"tags"`                         // 关联标签列表
	Cyom                         bool                   `json:"cyom"`                         // 文档未进一步说明的布尔标记
	ClosedTime                   string                  `json:"closedTime"`                   // 关闭时间
	ShowAllOutcomes              bool                   `json:"showAllOutcomes"`              // 是否展示全部结果
	ShowMarketImages             bool                   `json:"showMarketImages"`             // 是否展示市场图片
	AutomaticallyResolved        bool                   `json:"automaticallyResolved"`        // 是否自动结算
	EnableNegRisk                bool                   `json:"enableNegRisk"`                // 是否启用 neg risk 功能
	AutomaticallyActive          bool                   `json:"automaticallyActive"`          // 是否自动激活
	EventDate                    string                  `json:"eventDate"`                    // 事件日期
	StartTime                    string                  `json:"startTime"`                    // 开始时间
	EventWeek                    int64                  `json:"eventWeek"`                    // 事件所在周序号
	SeriesSlug                   string                  `json:"seriesSlug"`                   // 关联系列 slug
	Score                        string                  `json:"score"`                        // 比分/评分信息
	Elapsed                      string                  `json:"elapsed"`                      // 已进行时长
	Period                       string                  `json:"period"`                       // 当前阶段/周期
	Live                         bool                   `json:"live"`                         // 是否进行中
	Ended                        bool                   `json:"ended"`                        // 是否已结束
	FinishedTimestamp            string                  `json:"finishedTimestamp"`            // 完成时间戳
	GmpChartMode                 string                  `json:"gmpChartMode"`                 // GMP 图表模式
	EventCreators                []GammaEventCreator     `json:"eventCreators"`                // 事件创建者信息列表
	TweetCount                   int64                  `json:"tweetCount"`                   // 推文数量
	Chats                        []GammaChat             `json:"chats"`                        // 关联聊天频道列表
	FeaturedOrder                int64                  `json:"featuredOrder"`                // 精选排序值
	EstimateValue                bool                   `json:"estimateValue"`                // 是否启用估值
	CantEstimate                 bool                   `json:"cantEstimate"`                 // 是否无法估值
	EstimatedValue               string                  `json:"estimatedValue"`               // 估算值
	Templates                    []GammaTemplate         `json:"templates"`                    // 模板列表
	SpreadsMainLine              float64                `json:"spreadsMainLine"`              // 让分主线值
	TotalsMainLine               float64                `json:"totalsMainLine"`               // 大小分主线值
	CarouselMap                  string                  `json:"carouselMap"`                  // 轮播映射信息
	PendingDeployment            bool                   `json:"pendingDeployment"`            // 是否等待部署
	Deploying                    bool                   `json:"deploying"`                    // 是否正在部署
	DeployingTimestamp           string                  `json:"deployingTimestamp"`           // 部署开始时间戳
	ScheduledDeploymentTimestamp string                  `json:"scheduledDeploymentTimestamp"` // 计划部署时间戳
	GameStatus                   string                  `json:"gameStatus"`                   // 比赛状态
}

// GammaImageOptimization 表示图片优化记录。
type GammaImageOptimization struct {
	ID                        string   `json:"id"`                        // 图片优化记录 ID
	ImageURLSource            string   `json:"imageUrlSource"`            // 原始图片 URL
	ImageURLOptimized         string   `json:"imageUrlOptimized"`         // 优化后图片 URL
	ImageSizeKbSource         float64 `json:"imageSizeKbSource"`         // 原始图片大小（KB）
	ImageSizeKbOptimized      float64 `json:"imageSizeKbOptimized"`      // 优化后图片大小（KB）
	ImageOptimizedComplete    bool    `json:"imageOptimizedComplete"`    // 图片优化是否完成
	ImageOptimizedLastUpdated string   `json:"imageOptimizedLastUpdated"` // 图片优化最后更新时间
	RelID                     int64   `json:"relID"`                     // 关联对象 ID
	Field                     string   `json:"field"`                     // 关联字段名
	Relname                   string   `json:"relname"`                   // 关联对象名称
}

// GammaMarket 表示事件下的市场对象。
type GammaMarket struct {
	ID                           string                  `json:"id"`                           // 市场 ID
	Question                     string                  `json:"question"`                     // 市场问题/标题
	ConditionID                  string                  `json:"conditionId"`                  // 条件 ID
	Slug                         string                  `json:"slug"`                         // URL 友好的市场标识
	TwitterCardImage             string                  `json:"twitterCardImage"`             // Twitter 卡片图地址
	ResolutionSource             string                  `json:"resolutionSource"`             // 市场结算来源
	EndDate                      string                  `json:"endDate"`                      // 结束时间
	Category                     string                  `json:"category"`                     // 市场分类
	AmmType                      string                  `json:"ammType"`                      // AMM 类型
	Liquidity                    string                  `json:"liquidity"`                    // 流动性（字符串形式）
	SponsorName                  string                  `json:"sponsorName"`                  // 赞助方名称
	SponsorImage                 string                  `json:"sponsorImage"`                 // 赞助方图片地址
	StartDate                    string                  `json:"startDate"`                    // 开始时间
	XAxisValue                   string                  `json:"xAxisValue"`                   // X 轴值
	YAxisValue                   string                  `json:"yAxisValue"`                   // Y 轴值
	DenominationToken            string                  `json:"denominationToken"`            // 计价代币
	Fee                          string                  `json:"fee"`                          // 费率
	Image                        string                  `json:"image"`                        // 市场图片地址
	Icon                         string                  `json:"icon"`                         // 市场图标地址
	LowerBound                   string                  `json:"lowerBound"`                   // 下界值
	UpperBound                   string                  `json:"upperBound"`                   // 上界值
	Description                  string                  `json:"description"`                  // 市场描述
	Outcomes                     string                  `json:"outcomes"`                     // 结果项数据
	OutcomePrices                string                  `json:"outcomePrices"`                // 各结果价格数据
	Volume                       string                  `json:"volume"`                       // 成交量（字符串形式）
	Active                       bool                   `json:"active"`                       // 是否激活
	MarketType                   string                  `json:"marketType"`                   // 市场类型
	FormatType                   string                  `json:"formatType"`                   // 展示/格式类型
	LowerBoundDate               string                  `json:"lowerBoundDate"`               // 下界日期
	UpperBoundDate               string                  `json:"upperBoundDate"`               // 上界日期
	Closed                       bool                   `json:"closed"`                       // 是否关闭
	MarketMakerAddress           string                  `json:"marketMakerAddress"`           // 做市地址
	CreatedBy                    int64                  `json:"createdBy"`                    // 创建者 ID
	UpdatedBy                    int64                  `json:"updatedBy"`                    // 更新者 ID
	CreatedAt                    string                  `json:"createdAt"`                    // 记录创建时间
	UpdatedAt                    string                  `json:"updatedAt"`                    // 记录更新时间
	ClosedTime                   string                  `json:"closedTime"`                   // 关闭时间
	WideFormat                   bool                   `json:"wideFormat"`                   // 是否使用宽布局
	New                          bool                   `json:"new"`                          // 是否为新市场
	MailchimpTag                 string                  `json:"mailchimpTag"`                 // Mailchimp 标签
	Featured                     bool                   `json:"featured"`                     // 是否精选
	Archived                     bool                   `json:"archived"`                     // 是否归档
	ResolvedBy                   string                  `json:"resolvedBy"`                   // 结算执行者
	Restricted                   bool                   `json:"restricted"`                   // 是否受限
	MarketGroup                  int64                  `json:"marketGroup"`                  // 市场分组 ID
	GroupItemTitle               string                  `json:"groupItemTitle"`               // 分组项标题
	GroupItemThreshold           string                  `json:"groupItemThreshold"`           // 分组项阈值
	QuestionID                   string                  `json:"questionID"`                   // 问题 ID
	UmaEndDate                   string                  `json:"umaEndDate"`                   // UMA 截止时间
	EnableOrderBook              bool                   `json:"enableOrderBook"`              // 是否启用订单簿
	OrderPriceMinTickSize        float64                `json:"orderPriceMinTickSize"`        // 订单价格最小变动单位
	OrderMinSize                 float64                `json:"orderMinSize"`                 // 最小订单数量
	UmaResolutionStatus          string                  `json:"umaResolutionStatus"`          // UMA 结算状态
	CurationOrder                int64                  `json:"curationOrder"`                // 内容排序值
	VolumeNum                    float64                `json:"volumeNum"`                    // 成交量（数值形式）
	LiquidityNum                 float64                `json:"liquidityNum"`                 // 流动性（数值形式）
	EndDateIso                   string                  `json:"endDateIso"`                   // ISO 格式结束时间
	StartDateIso                 string                  `json:"startDateIso"`                 // ISO 格式开始时间
	UmaEndDateIso                string                  `json:"umaEndDateIso"`                // ISO 格式 UMA 截止时间
	HasReviewedDates             bool                   `json:"hasReviewedDates"`             // 日期是否已审核
	ReadyForCron                 bool                   `json:"readyForCron"`                 // 是否可用于定时任务处理
	CommentsEnabled              bool                   `json:"commentsEnabled"`              // 是否允许评论
	Volume24hr                   float64                `json:"volume24hr"`                   // 24 小时成交量
	Volume1wk                    float64                `json:"volume1wk"`                    // 1 周成交量
	Volume1mo                    float64                `json:"volume1mo"`                    // 1 月成交量
	Volume1yr                    float64                `json:"volume1yr"`                    // 1 年成交量
	GameStartTime                string                  `json:"gameStartTime"`                // 比赛开始时间
	SecondsDelay                 int64                  `json:"secondsDelay"`                 // 延迟秒数
	ClobTokenIds                 string                  `json:"clobTokenIds"`                 // CLOB token ID 列表
	DisqusThread                 string                  `json:"disqusThread"`                 // Disqus 讨论串 ID
	ShortOutcomes                string                  `json:"shortOutcomes"`                // 简短结果项数据
	TeamAID                      string                  `json:"teamAID"`                      // 队伍 A ID
	TeamBID                      string                  `json:"teamBID"`                      // 队伍 B ID
	UmaBond                      string                  `json:"umaBond"`                      // UMA bond 配置值
	UmaReward                    string                  `json:"umaReward"`                    // UMA 奖励值
	FpmmLive                     bool                   `json:"fpmmLive"`                     // FPMM 是否在线
	Volume24hrAmm                float64                `json:"volume24hrAmm"`                // 24 小时 AMM 成交量
	Volume1wkAmm                 float64                `json:"volume1wkAmm"`                 // 1 周 AMM 成交量
	Volume1moAmm                 float64                `json:"volume1moAmm"`                 // 1 月 AMM 成交量
	Volume1yrAmm                 float64                `json:"volume1yrAmm"`                 // 1 年 AMM 成交量
	Volume24hrClob               float64                `json:"volume24hrClob"`               // 24 小时 CLOB 成交量
	Volume1wkClob                float64                `json:"volume1wkClob"`                // 1 周 CLOB 成交量
	Volume1moClob                float64                `json:"volume1moClob"`                // 1 月 CLOB 成交量
	Volume1yrClob                float64                `json:"volume1yrClob"`                // 1 年 CLOB 成交量
	VolumeAmm                    float64                `json:"volumeAmm"`                    // AMM 总成交量
	VolumeClob                   float64                `json:"volumeClob"`                   // CLOB 总成交量
	LiquidityAmm                 float64                `json:"liquidityAmm"`                 // AMM 流动性
	LiquidityClob                float64                `json:"liquidityClob"`                // CLOB 流动性
	MakerBaseFee                 int64                  `json:"makerBaseFee"`                 // Maker 基础费率
	TakerBaseFee                 int64                  `json:"takerBaseFee"`                 // Taker 基础费率
	CustomLiveness               int64                  `json:"customLiveness"`               // 自定义存活时间
	AcceptingOrders              bool                   `json:"acceptingOrders"`              // 是否接受下单
	NotificationsEnabled         bool                   `json:"notificationsEnabled"`         // 是否开启通知
	Score                        int64                  `json:"score"`                        // 评分/比分值
	ImageOptimized               GammaImageOptimization `json:"imageOptimized"`               // 图片优化信息
	IconOptimized                GammaImageOptimization `json:"iconOptimized"`                // 图标优化信息
	Events                       []GammaEvent            `json:"events"`                       // 关联事件列表
	Categories                   []GammaCategory         `json:"categories"`                   // 分类列表
	Tags                         []GammaTag              `json:"tags"`                         // 标签列表
	Creator                      string                  `json:"creator"`                      // 创建者标识
	Ready                        bool                   `json:"ready"`                        // 是否已就绪
	Funded                       bool                   `json:"funded"`                       // 是否已注资
	PastSlugs                    string                  `json:"pastSlugs"`                    // 历史 slug 记录
	ReadyTimestamp               string                  `json:"readyTimestamp"`               // 就绪时间戳
	FundedTimestamp              string                  `json:"fundedTimestamp"`              // 注资时间戳
	AcceptingOrdersTimestamp     string                  `json:"acceptingOrdersTimestamp"`     // 开始接受订单时间戳
	Competitive                  float64                `json:"competitive"`                  // 竞争度相关数值
	RewardsMinSize               float64                `json:"rewardsMinSize"`               // 奖励最小规模
	RewardsMaxSpread             float64                `json:"rewardsMaxSpread"`             // 奖励允许的最大价差
	Spread                       float64                `json:"spread"`                       // 当前价差
	AutomaticallyResolved        bool                   `json:"automaticallyResolved"`        // 是否自动结算
	OneDayPriceChange            float64                `json:"oneDayPriceChange"`            // 1 日价格变化
	OneHourPriceChange           float64                `json:"oneHourPriceChange"`           // 1 小时价格变化
	OneWeekPriceChange           float64                `json:"oneWeekPriceChange"`           // 1 周价格变化
	OneMonthPriceChange          float64                `json:"oneMonthPriceChange"`          // 1 月价格变化
	OneYearPriceChange           float64                `json:"oneYearPriceChange"`           // 1 年价格变化
	LastTradePrice               float64                `json:"lastTradePrice"`               // 最近成交价
	BestBid                      float64                `json:"bestBid"`                      // 最优买价
	BestAsk                      float64                `json:"bestAsk"`                      // 最优卖价
	AutomaticallyActive          bool                   `json:"automaticallyActive"`          // 是否自动激活
	ClearBookOnStart             bool                   `json:"clearBookOnStart"`             // 开始时是否清空订单簿
	ChartColor                   string                  `json:"chartColor"`                   // 图表颜色
	SeriesColor                  string                  `json:"seriesColor"`                  // 系列颜色
	ShowGmpSeries                bool                   `json:"showGmpSeries"`                // 是否展示 GMP 系列
	ShowGmpOutcome               bool                   `json:"showGmpOutcome"`               // 是否展示 GMP 结果
	ManualActivation             bool                   `json:"manualActivation"`             // 是否手动激活
	NegRiskOther                 bool                   `json:"negRiskOther"`                 // neg risk 相关附加标记
	GameID                       string                  `json:"gameId"`                       // 比赛 ID
	GroupItemRange               string                  `json:"groupItemRange"`               // 分组项范围
	SportsMarketType             string                  `json:"sportsMarketType"`             // 体育市场类型
	Line                         float64                `json:"line"`                         // 盘口线值
	UmaResolutionStatuses        string                  `json:"umaResolutionStatuses"`        // UMA 结算状态集合
	PendingDeployment            bool                   `json:"pendingDeployment"`            // 是否等待部署
	Deploying                    bool                   `json:"deploying"`                    // 是否正在部署
	DeployingTimestamp           string                  `json:"deployingTimestamp"`           // 部署开始时间戳
	ScheduledDeploymentTimestamp string                  `json:"scheduledDeploymentTimestamp"` // 计划部署时间戳
	RfqEnabled                   bool                   `json:"rfqEnabled"`                   // 是否启用 RFQ
	EventStartTime               string                  `json:"eventStartTime"`               // 事件开始时间
	FeesEnabled                  bool                   `json:"feesEnabled"`                  // 是否启用手续费
	FeeSchedule                  GammaFeeSchedule       `json:"feeSchedule"`                  // 手续费计划配置
}

// GammaSeries 表示系列对象。
type GammaSeries struct {
	ID                string            `json:"id"`                // 系列 ID
	Ticker            string            `json:"ticker"`            // 系列代码/简称
	Slug              string            `json:"slug"`              // URL 友好的系列标识
	Title             string            `json:"title"`             // 系列标题
	Subtitle          string            `json:"subtitle"`          // 系列副标题
	SeriesType        string            `json:"seriesType"`        // 系列类型
	Recurrence        string            `json:"recurrence"`        // 重复/周期规则
	Description       string            `json:"description"`       // 系列描述
	Image             string            `json:"image"`             // 系列图片地址
	Icon              string            `json:"icon"`              // 系列图标地址
	Layout            string            `json:"layout"`            // 展示布局
	Active            bool             `json:"active"`            // 是否激活
	Closed            bool             `json:"closed"`            // 是否关闭
	Archived          bool             `json:"archived"`          // 是否归档
	New               bool             `json:"new"`               // 是否为新系列
	Featured          bool             `json:"featured"`          // 是否精选
	Restricted        bool             `json:"restricted"`        // 是否受限
	IsTemplate        bool             `json:"isTemplate"`        // 是否为模板系列
	TemplateVariables bool             `json:"templateVariables"` // 模板变量信息
	PublishedAt       string            `json:"publishedAt"`       // 发布时间
	CreatedBy         string            `json:"createdBy"`         // 创建者
	UpdatedBy         string            `json:"updatedBy"`         // 更新者
	CreatedAt         string            `json:"createdAt"`         // 记录创建时间
	UpdatedAt         string            `json:"updatedAt"`         // 记录更新时间
	CommentsEnabled   bool             `json:"commentsEnabled"`   // 是否允许评论
	Competitive       string            `json:"competitive"`       // 竞争度相关值
	Volume24hr        float64          `json:"volume24hr"`        // 24 小时成交量
	Volume            float64          `json:"volume"`            // 总成交量
	Liquidity         float64          `json:"liquidity"`         // 总流动性
	StartDate         string            `json:"startDate"`         // 开始时间
	PythTokenID       string            `json:"pythTokenID"`       // Pyth token ID
	CgAssetName       string            `json:"cgAssetName"`       // CoinGecko 资产名称
	Score             int64            `json:"score"`             // 评分/比分值
	Events            []GammaEvent      `json:"events"`            // 关联事件列表
	Collections       []GammaCollection `json:"collections"`       // 所属合集列表
	Categories        []GammaCategory   `json:"categories"`        // 分类列表
	Tags              []GammaTag        `json:"tags"`              // 标签列表
	CommentCount      int64            `json:"commentCount"`      // 评论数
	Chats             []GammaChat       `json:"chats"`             // 关联聊天频道列表
}

// GammaCategory 表示分类对象。
type GammaCategory struct {
	ID             string `json:"id"`             // 分类 ID
	Label          string `json:"label"`          // 分类展示名称
	ParentCategory string `json:"parentCategory"` // 父分类标识
	Slug           string `json:"slug"`           // URL 友好的分类标识
	PublishedAt    string `json:"publishedAt"`    // 发布时间
	CreatedBy      string `json:"createdBy"`      // 创建者
	UpdatedBy      string `json:"updatedBy"`      // 更新者
	CreatedAt      string `json:"createdAt"`      // 记录创建时间
	UpdatedAt      string `json:"updatedAt"`      // 记录更新时间
}

// GammaCollection 表示合集对象。
type GammaCollection struct {
	ID                   string                  `json:"id"`                   // 合集 ID
	Ticker               string                  `json:"ticker"`               // 合集代码/简称
	Slug                 string                  `json:"slug"`                 // URL 友好的合集标识
	Title                string                  `json:"title"`                // 合集标题
	Subtitle             string                  `json:"subtitle"`             // 合集副标题
	CollectionType       string                  `json:"collectionType"`       // 合集类型
	Description          string                  `json:"description"`          // 合集描述
	Tags                 string                  `json:"tags"`                 // 字符串形式的标签数据
	Image                string                  `json:"image"`                // 合集图片地址
	Icon                 string                  `json:"icon"`                 // 合集图标地址
	HeaderImage          string                  `json:"headerImage"`          // 头图地址
	Layout               string                  `json:"layout"`               // 展示布局
	Active               bool                   `json:"active"`               // 是否激活
	Closed               bool                   `json:"closed"`               // 是否关闭
	Archived             bool                   `json:"archived"`             // 是否归档
	New                  bool                   `json:"new"`                  // 是否为新合集
	Featured             bool                   `json:"featured"`             // 是否精选
	Restricted           bool                   `json:"restricted"`           // 是否受限
	IsTemplate           bool                   `json:"isTemplate"`           // 是否为模板合集
	TemplateVariables    string                  `json:"templateVariables"`    // 模板变量
	PublishedAt          string                  `json:"publishedAt"`          // 发布时间
	CreatedBy            string                  `json:"createdBy"`            // 创建者
	UpdatedBy            string                  `json:"updatedBy"`            // 更新者
	CreatedAt            string                  `json:"createdAt"`            // 记录创建时间
	UpdatedAt            string                  `json:"updatedAt"`            // 记录更新时间
	CommentsEnabled      bool                   `json:"commentsEnabled"`      // 是否允许评论
	ImageOptimized       GammaImageOptimization `json:"imageOptimized"`       // 图片优化信息
	IconOptimized        GammaImageOptimization `json:"iconOptimized"`        // 图标优化信息
	HeaderImageOptimized GammaImageOptimization `json:"headerImageOptimized"` // 头图优化信息
}

// GammaTag 表示标签对象。
type GammaTag struct {
	ID          string `json:"id"`          // 标签 ID
	Label       string `json:"label"`       // 标签展示名称
	Slug        string `json:"slug"`        // URL 友好的标签标识
	ForceShow   bool  `json:"forceShow"`   // 是否强制展示
	PublishedAt string `json:"publishedAt"` // 发布时间
	CreatedBy   int64 `json:"createdBy"`   // 创建者 ID
	UpdatedBy   int64 `json:"updatedBy"`   // 更新者 ID
	CreatedAt   string `json:"createdAt"`   // 记录创建时间
	UpdatedAt   string `json:"updatedAt"`   // 记录更新时间
	ForceHide   bool  `json:"forceHide"`   // 是否强制隐藏
	IsCarousel  bool  `json:"isCarousel"`  // 是否用于轮播展示
}

// GammaEventCreator 表示事件创建者信息。
type GammaEventCreator struct {
	ID            string `json:"id"`            // 创建者记录 ID
	CreatorName   string `json:"creatorName"`   // 创建者显示名
	CreatorHandle string `json:"creatorHandle"` // 创建者用户名/句柄
	CreatorURL    string `json:"creatorUrl"`    // 创建者资料或外部链接
	CreatorImage  string `json:"creatorImage"`  // 创建者头像地址
	CreatedAt     string `json:"createdAt"`     // 记录创建时间
	UpdatedAt     string `json:"updatedAt"`     // 记录更新时间
}

// GammaChat 表示聊天频道信息。
type GammaChat struct {
	ID           string `json:"id"`           // 聊天记录 ID
	ChannelID    string `json:"channelId"`    // 频道 ID
	ChannelName  string `json:"channelName"`  // 频道名称
	ChannelImage string `json:"channelImage"` // 频道图片地址
	Live         bool  `json:"live"`         // 是否处于直播/活跃状态
	StartTime    string `json:"startTime"`    // 开始时间
	EndTime      string `json:"endTime"`      // 结束时间
}

// GammaTemplate 表示模板对象。
type GammaTemplate struct {
	ID               string `json:"id"`               // 模板 ID
	EventTitle       string `json:"eventTitle"`       // 模板中的事件标题
	EventSlug        string `json:"eventSlug"`        // 模板中的事件 slug
	EventImage       string `json:"eventImage"`       // 模板中的事件图片地址
	MarketTitle      string `json:"marketTitle"`      // 模板中的市场标题
	Description      string `json:"description"`      // 模板描述
	ResolutionSource string `json:"resolutionSource"` // 模板使用的结算来源
	NegRisk          bool  `json:"negRisk"`          // 是否启用 neg risk
	SortBy           string `json:"sortBy"`           // 排序字段
	ShowMarketImages bool  `json:"showMarketImages"` // 是否展示市场图片
	SeriesSlug       string `json:"seriesSlug"`       // 关联系列 slug
	Outcomes         string `json:"outcomes"`         // 模板中的结果项数据
}

// GammaFeeSchedule 表示市场手续费计划。
type GammaFeeSchedule struct {
	Exponent   float64 `json:"exponent"`   // 手续费曲线指数
	Rate       float64 `json:"rate"`       // 手续费率
	TakerOnly  bool    `json:"takerOnly"`  // 是否仅对 taker 收费
	RebateRate float64 `json:"rebateRate"` // 返佣/回扣费率
}

// GammaGetPublicProfileRes 是 Get public profile by wallet address 接口的响应体（Gamma `GET /public-profile`）。
type GammaGetPublicProfileRes struct {
	CreatedAt             string                  `json:"createdAt,omitempty"`             // string(date-time) 资料创建时间（ISO 8601），可空
	ProxyWallet           string                  `json:"proxyWallet,omitempty"`           // string proxy 钱包地址，可空
	ProfileImage          string                  `json:"profileImage,omitempty"`          // string 头像 URL，可空
	DisplayUsernamePublic bool                    `json:"displayUsernamePublic,omitempty"` // boolean 是否公开展示用户名，可空
	Bio                   string                  `json:"bio,omitempty"`                   // string 个人简介，可空
	Pseudonym             string                  `json:"pseudonym,omitempty"`             // string 系统自动生成的别名，可空
	Name                  string                  `json:"name,omitempty"`                  // string 用户自选显示名，可空
	Users                 []GammaPublicProfileUser `json:"users,omitempty"`                 // 关联用户对象列表；JSON null 解码为 nil 切片
	XUsername             string                  `json:"xUsername,omitempty"`             // string X（Twitter）用户名，可空
	VerifiedBadge         bool                    `json:"verifiedBadge,omitempty"`         // boolean 是否带认证徽标，可空
}

// GammaPublicProfileUser 表示公开资料关联的用户条目（`PublicProfileResponse.users` 元素）。
type GammaPublicProfileUser struct {
	ID      string `json:"id"`      // string 用户 ID
	Creator bool   `json:"creator"` // boolean 是否为 creator
	Mod     bool   `json:"mod"`     // boolean 是否为 moderator
}
