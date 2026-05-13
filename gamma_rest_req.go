package mypolymarketapi

type GammaListEventsAPI struct {
	client *GammaRestClient
	req    *GammaListEventsReq
}

type GammaListEventsReq struct {
	Limit           *int64    `json:"limit"`            // 返回的最大结果数
	Offset          *int64    `json:"offset"`           // 分页偏移量
	Order           *string   `json:"order"`            // 用于排序的 JSON 字段名
	Ascending       *bool     `json:"ascending"`        // 排序方向
	ID              *[]int64  `json:"id"`               // 按事件 ID 过滤
	TagID           *int64    `json:"tag_id"`           // 按标签 ID 过滤
	ExcludeTagID    *[]int64  `json:"exclude_tag_id"`   // 排除这些标签 ID
	Slug            *[]string `json:"slug"`             // 按事件 slug 过滤
	TagSlug         *string   `json:"tag_slug"`         // 按标签 slug 过滤
	RelatedTags     *bool     `json:"related_tags"`     // 是否包含相关标签过滤
	Active          *bool     `json:"active"`           // 按是否激活过滤
	Archived        *bool     `json:"archived"`         // 按是否归档过滤
	Featured        *bool     `json:"featured"`         // 按是否精选过滤
	Cyom            *bool     `json:"cyom"`             // 按 cyom 标记过滤
	IncludeChat     *bool     `json:"include_chat"`     // 为 true 时包含 Chats 和 Series.Chats 关联
	IncludeTemplate *bool     `json:"include_template"` // 为 true 时包含 Templates 关联
	Recurrence      *string   `json:"recurrence"`       // 按 recurrence 过滤
	Closed          *bool     `json:"closed"`           // 按是否已关闭过滤
	LiquidityMin    *float64  `json:"liquidity_min"`    // 最小流动性过滤
	LiquidityMax    *float64  `json:"liquidity_max"`    // 最大流动性过滤
	VolumeMin       *float64  `json:"volume_min"`       // 最小成交量过滤
	VolumeMax       *float64  `json:"volume_max"`       // 最大成交量过滤
	StartDateMin    *string   `json:"start_date_min"`   // 开始时间下限过滤
	StartDateMax    *string   `json:"start_date_max"`   // 开始时间上限过滤
	EndDateMin      *string   `json:"end_date_min"`     // 结束时间下限过滤
	EndDateMax      *string   `json:"end_date_max"`     // 结束时间上限过滤
}

// 返回的最大结果数
func (api *GammaListEventsAPI) Limit(limit int64) *GammaListEventsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// 分页偏移量
func (api *GammaListEventsAPI) Offset(offset int64) *GammaListEventsAPI {
	api.req.Offset = GetPointer(offset)
	return api
}

// 用于排序的 JSON 字段名
func (api *GammaListEventsAPI) Order(order string) *GammaListEventsAPI {
	api.req.Order = GetPointer(order)
	return api
}

// 排序方向
func (api *GammaListEventsAPI) Ascending(ascending bool) *GammaListEventsAPI {
	api.req.Ascending = GetPointer(ascending)
	return api
}

// 按事件 ID 过滤
func (api *GammaListEventsAPI) ID(id []int64) *GammaListEventsAPI {
	api.req.ID = GetPointer(id)
	return api
}

// 按标签 ID 过滤
func (api *GammaListEventsAPI) TagID(tagID int64) *GammaListEventsAPI {
	api.req.TagID = GetPointer(tagID)
	return api
}

// 排除这些标签 ID
func (api *GammaListEventsAPI) ExcludeTagID(excludeTagID []int64) *GammaListEventsAPI {
	api.req.ExcludeTagID = GetPointer(excludeTagID)
	return api
}

// 按事件 slug 过滤
func (api *GammaListEventsAPI) Slug(slug []string) *GammaListEventsAPI {
	api.req.Slug = GetPointer(slug)
	return api
}

// 按标签 slug 过滤
func (api *GammaListEventsAPI) TagSlug(tagSlug string) *GammaListEventsAPI {
	api.req.TagSlug = GetPointer(tagSlug)
	return api
}

// 是否包含相关标签过滤
func (api *GammaListEventsAPI) RelatedTags(relatedTags bool) *GammaListEventsAPI {
	api.req.RelatedTags = GetPointer(relatedTags)
	return api
}

// 按是否激活过滤
func (api *GammaListEventsAPI) Active(active bool) *GammaListEventsAPI {
	api.req.Active = GetPointer(active)
	return api
}

// 按是否归档过滤
func (api *GammaListEventsAPI) Archived(archived bool) *GammaListEventsAPI {
	api.req.Archived = GetPointer(archived)
	return api
}

// 按是否精选过滤
func (api *GammaListEventsAPI) Featured(featured bool) *GammaListEventsAPI {
	api.req.Featured = GetPointer(featured)
	return api
}

// 按 cyom 标记过滤
func (api *GammaListEventsAPI) Cyom(cyom bool) *GammaListEventsAPI {
	api.req.Cyom = GetPointer(cyom)
	return api
}

// 为 true 时包含 Chats 和 Series.Chats 关联
func (api *GammaListEventsAPI) IncludeChat(includeChat bool) *GammaListEventsAPI {
	api.req.IncludeChat = GetPointer(includeChat)
	return api
}

// 为 true 时包含 Templates 关联
func (api *GammaListEventsAPI) IncludeTemplate(includeTemplate bool) *GammaListEventsAPI {
	api.req.IncludeTemplate = GetPointer(includeTemplate)
	return api
}

// 按 recurrence 过滤
func (api *GammaListEventsAPI) Recurrence(recurrence string) *GammaListEventsAPI {
	api.req.Recurrence = GetPointer(recurrence)
	return api
}

// 按是否已关闭过滤
func (api *GammaListEventsAPI) Closed(closed bool) *GammaListEventsAPI {
	api.req.Closed = GetPointer(closed)
	return api
}

// 最小流动性过滤
func (api *GammaListEventsAPI) LiquidityMin(liquidityMin float64) *GammaListEventsAPI {
	api.req.LiquidityMin = GetPointer(liquidityMin)
	return api
}

// 最大流动性过滤
func (api *GammaListEventsAPI) LiquidityMax(liquidityMax float64) *GammaListEventsAPI {
	api.req.LiquidityMax = GetPointer(liquidityMax)
	return api
}

// 最小成交量过滤
func (api *GammaListEventsAPI) VolumeMin(volumeMin float64) *GammaListEventsAPI {
	api.req.VolumeMin = GetPointer(volumeMin)
	return api
}

// 最大成交量过滤
func (api *GammaListEventsAPI) VolumeMax(volumeMax float64) *GammaListEventsAPI {
	api.req.VolumeMax = GetPointer(volumeMax)
	return api
}

// 开始时间下限过滤
func (api *GammaListEventsAPI) StartDateMin(startDateMin string) *GammaListEventsAPI {
	api.req.StartDateMin = GetPointer(startDateMin)
	return api
}

// 开始时间上限过滤
func (api *GammaListEventsAPI) StartDateMax(startDateMax string) *GammaListEventsAPI {
	api.req.StartDateMax = GetPointer(startDateMax)
	return api
}

// 结束时间下限过滤
func (api *GammaListEventsAPI) EndDateMin(endDateMin string) *GammaListEventsAPI {
	api.req.EndDateMin = GetPointer(endDateMin)
	return api
}

// 结束时间上限过滤
func (api *GammaListEventsAPI) EndDateMax(endDateMax string) *GammaListEventsAPI {
	api.req.EndDateMax = GetPointer(endDateMax)
	return api
}

type GammaListEventsKeysetAPI struct {
	client *GammaRestClient
	req    *GammaListEventsKeysetReq
}

type GammaListEventsKeysetReq struct {
	Limit            *int64    `json:"limit"`              // 返回的最大结果数，最大 500
	Order            *string   `json:"order"`              // 用于排序的 JSON 字段名，支持逗号分隔多个字段
	Ascending        *bool     `json:"ascending"`          // 排序方向，仅在设置 order 时生效
	AfterCursor      *string   `json:"after_cursor"`       // 上一页响应 next_cursor 返回的游标
	Offset           *int64    `json:"offset"`             // 文档标注为不允许传入，传入会返回 422
	ID               *[]int64  `json:"id"`                 // 按事件 ID 过滤
	Slug             *[]string `json:"slug"`               // 按事件 slug 过滤
	Closed           *bool     `json:"closed"`             // 按是否已关闭过滤
	Live             *bool     `json:"live"`               // 按是否进行中过滤
	Featured         *bool     `json:"featured"`           // 按是否精选过滤
	Cyom             *bool     `json:"cyom"`               // 按 cyom 标记过滤
	TitleSearch      *string   `json:"title_search"`       // 按标题搜索过滤
	LiquidityMin     *float64  `json:"liquidity_min"`      // 最小流动性过滤
	LiquidityMax     *float64  `json:"liquidity_max"`      // 最大流动性过滤
	VolumeMin        *float64  `json:"volume_min"`         // 最小成交量过滤
	VolumeMax        *float64  `json:"volume_max"`         // 最大成交量过滤
	StartDateMin     *string   `json:"start_date_min"`     // 开始时间下限过滤
	StartDateMax     *string   `json:"start_date_max"`     // 开始时间上限过滤
	EndDateMin       *string   `json:"end_date_min"`       // 结束时间下限过滤
	EndDateMax       *string   `json:"end_date_max"`       // 结束时间上限过滤
	StartTimeMin     *string   `json:"start_time_min"`     // 开赛时间下限过滤
	StartTimeMax     *string   `json:"start_time_max"`     // 开赛时间上限过滤
	TagID            *[]int64  `json:"tag_id"`             // 按标签 ID 过滤
	TagSlug          *string   `json:"tag_slug"`           // 按标签 slug 过滤
	ExcludeTagID     *[]int64  `json:"exclude_tag_id"`     // 排除这些标签 ID，且不能与 tag_id 重叠
	RelatedTags      *bool     `json:"related_tags"`       // 是否包含相关标签过滤
	TagMatch         *string   `json:"tag_match"`          // 标签匹配方式
	SeriesID         *[]int64  `json:"series_id"`          // 按系列 ID 过滤
	GameID           *[]int64  `json:"game_id"`            // 按比赛 ID 过滤
	EventDate        *string   `json:"event_date"`         // 按事件日期过滤
	EventWeek        *int64    `json:"event_week"`         // 按事件周过滤
	FeaturedOrder    *bool     `json:"featured_order"`     // 按 featured_order 相关条件过滤
	Recurrence       *string   `json:"recurrence"`         // 按 recurrence 过滤
	CreatedBy        *[]string `json:"created_by"`         // 按创建者过滤
	ParentEventID    *int64    `json:"parent_event_id"`    // 按父事件 ID 过滤
	IncludeChildren  *bool     `json:"include_children"`   // 是否包含子事件
	PartnerSlug      *string   `json:"partner_slug"`       // 传入后匹配事件会包含 external_partners
	IncludeChat      *bool     `json:"include_chat"`       // 为 true 时包含 Chats 和 Series.Chats 关联
	IncludeTemplate  *bool     `json:"include_template"`   // 为 true 时包含 Templates 关联
	IncludeBestLines *bool     `json:"include_best_lines"` // 为 true 时包含 BestLines 关联
	Locale           *string   `json:"locale"`             // 语言区域参数
}

// 返回的最大结果数，最大 500
func (api *GammaListEventsKeysetAPI) Limit(limit int64) *GammaListEventsKeysetAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// 用于排序的 JSON 字段名，支持逗号分隔多个字段
func (api *GammaListEventsKeysetAPI) Order(order string) *GammaListEventsKeysetAPI {
	api.req.Order = GetPointer(order)
	return api
}

// 排序方向，仅在设置 order 时生效
func (api *GammaListEventsKeysetAPI) Ascending(ascending bool) *GammaListEventsKeysetAPI {
	api.req.Ascending = GetPointer(ascending)
	return api
}

// 上一页响应 next_cursor 返回的游标
func (api *GammaListEventsKeysetAPI) AfterCursor(afterCursor string) *GammaListEventsKeysetAPI {
	api.req.AfterCursor = GetPointer(afterCursor)
	return api
}

// 文档标注为不允许传入，传入会返回 422
func (api *GammaListEventsKeysetAPI) Offset(offset int64) *GammaListEventsKeysetAPI {
	api.req.Offset = GetPointer(offset)
	return api
}

// 按事件 ID 过滤
func (api *GammaListEventsKeysetAPI) ID(id []int64) *GammaListEventsKeysetAPI {
	api.req.ID = GetPointer(id)
	return api
}

// 按事件 slug 过滤
func (api *GammaListEventsKeysetAPI) Slug(slug []string) *GammaListEventsKeysetAPI {
	api.req.Slug = GetPointer(slug)
	return api
}

// 按是否已关闭过滤
func (api *GammaListEventsKeysetAPI) Closed(closed bool) *GammaListEventsKeysetAPI {
	api.req.Closed = GetPointer(closed)
	return api
}

// 按是否进行中过滤
func (api *GammaListEventsKeysetAPI) Live(live bool) *GammaListEventsKeysetAPI {
	api.req.Live = GetPointer(live)
	return api
}

// 按是否精选过滤
func (api *GammaListEventsKeysetAPI) Featured(featured bool) *GammaListEventsKeysetAPI {
	api.req.Featured = GetPointer(featured)
	return api
}

// 按 cyom 标记过滤
func (api *GammaListEventsKeysetAPI) Cyom(cyom bool) *GammaListEventsKeysetAPI {
	api.req.Cyom = GetPointer(cyom)
	return api
}

// 按标题搜索过滤
func (api *GammaListEventsKeysetAPI) TitleSearch(titleSearch string) *GammaListEventsKeysetAPI {
	api.req.TitleSearch = GetPointer(titleSearch)
	return api
}

// 最小流动性过滤
func (api *GammaListEventsKeysetAPI) LiquidityMin(liquidityMin float64) *GammaListEventsKeysetAPI {
	api.req.LiquidityMin = GetPointer(liquidityMin)
	return api
}

// 最大流动性过滤
func (api *GammaListEventsKeysetAPI) LiquidityMax(liquidityMax float64) *GammaListEventsKeysetAPI {
	api.req.LiquidityMax = GetPointer(liquidityMax)
	return api
}

// 最小成交量过滤
func (api *GammaListEventsKeysetAPI) VolumeMin(volumeMin float64) *GammaListEventsKeysetAPI {
	api.req.VolumeMin = GetPointer(volumeMin)
	return api
}

// 最大成交量过滤
func (api *GammaListEventsKeysetAPI) VolumeMax(volumeMax float64) *GammaListEventsKeysetAPI {
	api.req.VolumeMax = GetPointer(volumeMax)
	return api
}

// 开始时间下限过滤
func (api *GammaListEventsKeysetAPI) StartDateMin(startDateMin string) *GammaListEventsKeysetAPI {
	api.req.StartDateMin = GetPointer(startDateMin)
	return api
}

// 开始时间上限过滤
func (api *GammaListEventsKeysetAPI) StartDateMax(startDateMax string) *GammaListEventsKeysetAPI {
	api.req.StartDateMax = GetPointer(startDateMax)
	return api
}

// 结束时间下限过滤
func (api *GammaListEventsKeysetAPI) EndDateMin(endDateMin string) *GammaListEventsKeysetAPI {
	api.req.EndDateMin = GetPointer(endDateMin)
	return api
}

// 结束时间上限过滤
func (api *GammaListEventsKeysetAPI) EndDateMax(endDateMax string) *GammaListEventsKeysetAPI {
	api.req.EndDateMax = GetPointer(endDateMax)
	return api
}

// 开赛时间下限过滤
func (api *GammaListEventsKeysetAPI) StartTimeMin(startTimeMin string) *GammaListEventsKeysetAPI {
	api.req.StartTimeMin = GetPointer(startTimeMin)
	return api
}

// 开赛时间上限过滤
func (api *GammaListEventsKeysetAPI) StartTimeMax(startTimeMax string) *GammaListEventsKeysetAPI {
	api.req.StartTimeMax = GetPointer(startTimeMax)
	return api
}

// 按标签 ID 过滤
func (api *GammaListEventsKeysetAPI) TagID(tagID []int64) *GammaListEventsKeysetAPI {
	api.req.TagID = GetPointer(tagID)
	return api
}

// 按标签 slug 过滤
func (api *GammaListEventsKeysetAPI) TagSlug(tagSlug string) *GammaListEventsKeysetAPI {
	api.req.TagSlug = GetPointer(tagSlug)
	return api
}

// 排除这些标签 ID，且不能与 tag_id 重叠
func (api *GammaListEventsKeysetAPI) ExcludeTagID(excludeTagID []int64) *GammaListEventsKeysetAPI {
	api.req.ExcludeTagID = GetPointer(excludeTagID)
	return api
}

// 是否包含相关标签过滤
func (api *GammaListEventsKeysetAPI) RelatedTags(relatedTags bool) *GammaListEventsKeysetAPI {
	api.req.RelatedTags = GetPointer(relatedTags)
	return api
}

// 标签匹配方式
func (api *GammaListEventsKeysetAPI) TagMatch(tagMatch string) *GammaListEventsKeysetAPI {
	api.req.TagMatch = GetPointer(tagMatch)
	return api
}

// 按系列 ID 过滤
func (api *GammaListEventsKeysetAPI) SeriesID(seriesID []int64) *GammaListEventsKeysetAPI {
	api.req.SeriesID = GetPointer(seriesID)
	return api
}

// 按比赛 ID 过滤
func (api *GammaListEventsKeysetAPI) GameID(gameID []int64) *GammaListEventsKeysetAPI {
	api.req.GameID = GetPointer(gameID)
	return api
}

// 按事件日期过滤
func (api *GammaListEventsKeysetAPI) EventDate(eventDate string) *GammaListEventsKeysetAPI {
	api.req.EventDate = GetPointer(eventDate)
	return api
}

// 按事件周过滤
func (api *GammaListEventsKeysetAPI) EventWeek(eventWeek int64) *GammaListEventsKeysetAPI {
	api.req.EventWeek = GetPointer(eventWeek)
	return api
}

// 按 featured_order 相关条件过滤
func (api *GammaListEventsKeysetAPI) FeaturedOrder(featuredOrder bool) *GammaListEventsKeysetAPI {
	api.req.FeaturedOrder = GetPointer(featuredOrder)
	return api
}

// 按 recurrence 过滤
func (api *GammaListEventsKeysetAPI) Recurrence(recurrence string) *GammaListEventsKeysetAPI {
	api.req.Recurrence = GetPointer(recurrence)
	return api
}

// 按创建者过滤
func (api *GammaListEventsKeysetAPI) CreatedBy(createdBy []string) *GammaListEventsKeysetAPI {
	api.req.CreatedBy = GetPointer(createdBy)
	return api
}

// 按父事件 ID 过滤
func (api *GammaListEventsKeysetAPI) ParentEventID(parentEventID int64) *GammaListEventsKeysetAPI {
	api.req.ParentEventID = GetPointer(parentEventID)
	return api
}

// 是否包含子事件
func (api *GammaListEventsKeysetAPI) IncludeChildren(includeChildren bool) *GammaListEventsKeysetAPI {
	api.req.IncludeChildren = GetPointer(includeChildren)
	return api
}

// 传入后匹配事件会包含 external_partners
func (api *GammaListEventsKeysetAPI) PartnerSlug(partnerSlug string) *GammaListEventsKeysetAPI {
	api.req.PartnerSlug = GetPointer(partnerSlug)
	return api
}

// 为 true 时包含 Chats 和 Series.Chats 关联
func (api *GammaListEventsKeysetAPI) IncludeChat(includeChat bool) *GammaListEventsKeysetAPI {
	api.req.IncludeChat = GetPointer(includeChat)
	return api
}

// 为 true 时包含 Templates 关联
func (api *GammaListEventsKeysetAPI) IncludeTemplate(includeTemplate bool) *GammaListEventsKeysetAPI {
	api.req.IncludeTemplate = GetPointer(includeTemplate)
	return api
}

// 为 true 时包含 BestLines 关联
func (api *GammaListEventsKeysetAPI) IncludeBestLines(includeBestLines bool) *GammaListEventsKeysetAPI {
	api.req.IncludeBestLines = GetPointer(includeBestLines)
	return api
}

// 语言区域参数
func (api *GammaListEventsKeysetAPI) Locale(locale string) *GammaListEventsKeysetAPI {
	api.req.Locale = GetPointer(locale)
	return api
}

type GammaListMarketsAPI struct {
	client *GammaRestClient
	req    *GammaListMarketsReq
}

type GammaListMarketsReq struct {
	Limit               *int64    `json:"limit"`                 // integer 每页最大条数；query 可选，≥0，不传由服务端默认
	Offset              *int64    `json:"offset"`                // integer 分页偏移；query 可选，≥0
	Order               *string   `json:"order"`                 // string 排序字段，逗号分隔多个 JSON 字段名；仅传参时参与排序
	Ascending           *bool     `json:"ascending"`             // boolean 与 order 联用的升降序；无 order 时通常无意义
	ID                  *[]int64  `json:"id"`                    // integer[] 按市场 id 列表过滤；序列化为 JSON 数组作 query
	Slug                *[]string `json:"slug"`                  // string[] 按市场 slug 列表过滤
	ClobTokenIds        *[]string `json:"clob_token_ids"`        // string[] 按 CLOB token id 列表过滤
	ConditionIds        *[]string `json:"condition_ids"`         // string[] 按 Polymarket condition id 列表过滤
	MarketMakerAddress  *[]string `json:"market_maker_address"`  // string[] 按做市地址列表过滤
	LiquidityNumMin     *float64  `json:"liquidity_num_min"`     // number 流动性（数值）下限
	LiquidityNumMax     *float64  `json:"liquidity_num_max"`     // number 流动性（数值）上限
	VolumeNumMin        *float64  `json:"volume_num_min"`        // number 成交量（数值）下限
	VolumeNumMax        *float64  `json:"volume_num_max"`        // number 成交量（数值）上限
	StartDateMin        *string   `json:"start_date_min"`        // string(date-time) 市场开始时间下限，ISO 8601 风格字符串
	StartDateMax        *string   `json:"start_date_max"`        // string(date-time) 市场开始时间上限
	EndDateMin          *string   `json:"end_date_min"`          // string(date-time) 市场结束时间下限
	EndDateMax          *string   `json:"end_date_max"`          // string(date-time) 市场结束时间上限
	TagID               *int64    `json:"tag_id"`                // integer 单个标签 id；本 List markets 接口为单值（keyset 为多值数组）
	RelatedTags         *bool     `json:"related_tags"`          // boolean 是否把相关标签纳入过滤语义
	Cyom                *bool     `json:"cyom"`                  // boolean cyom 标记过滤
	UmaResolutionStatus *string   `json:"uma_resolution_status"` // string UMA 结算状态筛选
	GameID              *string   `json:"game_id"`               // string 比赛/对局 id（Gamma OpenAPI 为 string）
	SportsMarketTypes   *[]string `json:"sports_market_types"`   // string[] 体育市场类型列表
	RewardsMinSize      *float64  `json:"rewards_min_size"`      // number 奖励最小规模下限（仅本 offset 分页接口）
	QuestionIds         *[]string `json:"question_ids"`          // string[] 问题 id 列表过滤
	IncludeTag          *bool     `json:"include_tag"`           // boolean 为 true 时在每条 market 上附带 Tags 关联
	Closed              *bool     `json:"closed"`                // boolean 是否包含已关闭市场；OpenAPI 默认 false
}

// int64 每页最大条数，≥0；未传则由服务端默认
func (api *GammaListMarketsAPI) Limit(limit int64) *GammaListMarketsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// int64 分页偏移，≥0
func (api *GammaListMarketsAPI) Offset(offset int64) *GammaListMarketsAPI {
	api.req.Offset = GetPointer(offset)
	return api
}

// string 逗号分隔的排序 JSON 字段名
func (api *GammaListMarketsAPI) Order(order string) *GammaListMarketsAPI {
	api.req.Order = GetPointer(order)
	return api
}

// bool 是否升序，需配合 Order
func (api *GammaListMarketsAPI) Ascending(ascending bool) *GammaListMarketsAPI {
	api.req.Ascending = GetPointer(ascending)
	return api
}

// []int64 市场 id 列表
func (api *GammaListMarketsAPI) ID(id []int64) *GammaListMarketsAPI {
	api.req.ID = GetPointer(id)
	return api
}

// []string 市场 slug 列表
func (api *GammaListMarketsAPI) Slug(slug []string) *GammaListMarketsAPI {
	api.req.Slug = GetPointer(slug)
	return api
}

// []string CLOB token id 列表
func (api *GammaListMarketsAPI) ClobTokenIds(clobTokenIds []string) *GammaListMarketsAPI {
	api.req.ClobTokenIds = GetPointer(clobTokenIds)
	return api
}

// []string condition id 列表
func (api *GammaListMarketsAPI) ConditionIds(conditionIds []string) *GammaListMarketsAPI {
	api.req.ConditionIds = GetPointer(conditionIds)
	return api
}

// []string 做市地址列表
func (api *GammaListMarketsAPI) MarketMakerAddress(marketMakerAddress []string) *GammaListMarketsAPI {
	api.req.MarketMakerAddress = GetPointer(marketMakerAddress)
	return api
}

// float64 流动性（数值）下限
func (api *GammaListMarketsAPI) LiquidityNumMin(liquidityNumMin float64) *GammaListMarketsAPI {
	api.req.LiquidityNumMin = GetPointer(liquidityNumMin)
	return api
}

// float64 流动性（数值）上限
func (api *GammaListMarketsAPI) LiquidityNumMax(liquidityNumMax float64) *GammaListMarketsAPI {
	api.req.LiquidityNumMax = GetPointer(liquidityNumMax)
	return api
}

// float64 成交量（数值）下限
func (api *GammaListMarketsAPI) VolumeNumMin(volumeNumMin float64) *GammaListMarketsAPI {
	api.req.VolumeNumMin = GetPointer(volumeNumMin)
	return api
}

// float64 成交量（数值）上限
func (api *GammaListMarketsAPI) VolumeNumMax(volumeNumMax float64) *GammaListMarketsAPI {
	api.req.VolumeNumMax = GetPointer(volumeNumMax)
	return api
}

// string 开始时间下限，date-time 格式字符串
func (api *GammaListMarketsAPI) StartDateMin(startDateMin string) *GammaListMarketsAPI {
	api.req.StartDateMin = GetPointer(startDateMin)
	return api
}

// string 开始时间上限，date-time 格式字符串
func (api *GammaListMarketsAPI) StartDateMax(startDateMax string) *GammaListMarketsAPI {
	api.req.StartDateMax = GetPointer(startDateMax)
	return api
}

// string 结束时间下限，date-time 格式字符串
func (api *GammaListMarketsAPI) EndDateMin(endDateMin string) *GammaListMarketsAPI {
	api.req.EndDateMin = GetPointer(endDateMin)
	return api
}

// string 结束时间上限，date-time 格式字符串
func (api *GammaListMarketsAPI) EndDateMax(endDateMax string) *GammaListMarketsAPI {
	api.req.EndDateMax = GetPointer(endDateMax)
	return api
}

// int64 单个标签 id
func (api *GammaListMarketsAPI) TagID(tagID int64) *GammaListMarketsAPI {
	api.req.TagID = GetPointer(tagID)
	return api
}

// bool 是否把相关标签纳入过滤语义
func (api *GammaListMarketsAPI) RelatedTags(relatedTags bool) *GammaListMarketsAPI {
	api.req.RelatedTags = GetPointer(relatedTags)
	return api
}

// bool cyom 标记过滤
func (api *GammaListMarketsAPI) Cyom(cyom bool) *GammaListMarketsAPI {
	api.req.Cyom = GetPointer(cyom)
	return api
}

// string UMA 结算状态筛选
func (api *GammaListMarketsAPI) UmaResolutionStatus(umaResolutionStatus string) *GammaListMarketsAPI {
	api.req.UmaResolutionStatus = GetPointer(umaResolutionStatus)
	return api
}

// string 比赛/对局 id（Gamma OpenAPI 为 string）
func (api *GammaListMarketsAPI) GameID(gameID string) *GammaListMarketsAPI {
	api.req.GameID = GetPointer(gameID)
	return api
}

// []string 体育市场类型列表
func (api *GammaListMarketsAPI) SportsMarketTypes(sportsMarketTypes []string) *GammaListMarketsAPI {
	api.req.SportsMarketTypes = GetPointer(sportsMarketTypes)
	return api
}

// float64 奖励最小规模下限
func (api *GammaListMarketsAPI) RewardsMinSize(rewardsMinSize float64) *GammaListMarketsAPI {
	api.req.RewardsMinSize = GetPointer(rewardsMinSize)
	return api
}

// []string 问题 id 列表
func (api *GammaListMarketsAPI) QuestionIds(questionIds []string) *GammaListMarketsAPI {
	api.req.QuestionIds = GetPointer(questionIds)
	return api
}

// bool 为 true 时在每条 market 上附带 Tags 关联
func (api *GammaListMarketsAPI) IncludeTag(includeTag bool) *GammaListMarketsAPI {
	api.req.IncludeTag = GetPointer(includeTag)
	return api
}

// bool 是否包含已关闭市场，默认语义见 OpenAPI
func (api *GammaListMarketsAPI) Closed(closed bool) *GammaListMarketsAPI {
	api.req.Closed = GetPointer(closed)
	return api
}

type GammaListMarketsKeysetAPI struct {
	client *GammaRestClient
	req    *GammaListMarketsKeysetReq
}

type GammaListMarketsKeysetReq struct {
	Limit               *int64    `json:"limit"`                 // integer 每页最大条数；OpenAPI 最大 1000，默认由服务端决定
	Order               *string   `json:"order"`                 // string 排序字段，逗号分隔，如 volume_num,liquidity_num
	Ascending           *bool     `json:"ascending"`             // boolean 与 order 联用；无 order 时通常无意义
	AfterCursor         *string   `json:"after_cursor"`          // string 游标分页；取上一响应 next_cursor 原样传入
	Offset              *int64    `json:"offset"`                // integer 勿传；文档禁止出现在 query，否则 422，应使用 after_cursor
	ID                  *[]int64  `json:"id"`                    // integer[] 按市场 id 列表过滤
	Slug                *[]string `json:"slug"`                  // string[] 按市场 slug 列表过滤
	Closed              *bool     `json:"closed"`                // boolean 是否包含已关闭；OpenAPI 默认 false
	Decimalized         *bool     `json:"decimalized"`           // boolean 是否以小数化等形式返回部分数值展示
	ClobTokenIds        *[]string `json:"clob_token_ids"`        // string[] CLOB token id 列表
	ConditionIds        *[]string `json:"condition_ids"`         // string[] condition id 列表
	QuestionIds         *[]string `json:"question_ids"`          // string[] 问题 id 列表
	MarketMakerAddress  *[]string `json:"market_maker_address"`  // string[] 做市地址列表
	LiquidityNumMin     *float64  `json:"liquidity_num_min"`     // number 流动性数值下限
	LiquidityNumMax     *float64  `json:"liquidity_num_max"`     // number 流动性数值上限
	VolumeNumMin        *float64  `json:"volume_num_min"`        // number 成交量数值下限
	VolumeNumMax        *float64  `json:"volume_num_max"`        // number 成交量数值上限
	StartDateMin        *string   `json:"start_date_min"`        // string(date-time) 开始时间下限
	StartDateMax        *string   `json:"start_date_max"`        // string(date-time) 开始时间上限
	EndDateMin          *string   `json:"end_date_min"`          // string(date-time) 结束时间下限
	EndDateMax          *string   `json:"end_date_max"`          // string(date-time) 结束时间上限
	TagID               *[]int64  `json:"tag_id"`                // integer[] 标签 id 列表（本 keyset 接口为数组；与 offset 版单值 tag_id 不同）
	RelatedTags         *bool     `json:"related_tags"`          // boolean 相关标签是否参与过滤
	TagMatch            *string   `json:"tag_match"`             // string 多标签时的匹配策略字符串（取值以服务端为准）
	Cyom                *bool     `json:"cyom"`                  // boolean cyom 标记
	RfqEnabled          *bool     `json:"rfq_enabled"`           // boolean 是否只保留启用 RFQ 的市场
	UmaResolutionStatus *string   `json:"uma_resolution_status"` // string UMA 结算状态
	GameID              *string   `json:"game_id"`               // string 比赛/对局 id
	SportsMarketTypes   *[]string `json:"sports_market_types"`   // string[] 体育市场类型
	IncludeTag          *bool     `json:"include_tag"`           // boolean 为 true 时展开每条 market 的 Tags
	Locale              *string   `json:"locale"`                // string 区域/语言偏好，如 en、具体格式以服务端为准
}

// int64 每页最大条数，OpenAPI 最大 1000；未传由服务端默认
func (api *GammaListMarketsKeysetAPI) Limit(limit int64) *GammaListMarketsKeysetAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// string 逗号分隔排序字段，如 volume_num,liquidity_num
func (api *GammaListMarketsKeysetAPI) Order(order string) *GammaListMarketsKeysetAPI {
	api.req.Order = GetPointer(order)
	return api
}

// bool 是否升序，需配合 Order
func (api *GammaListMarketsKeysetAPI) Ascending(ascending bool) *GammaListMarketsKeysetAPI {
	api.req.Ascending = GetPointer(ascending)
	return api
}

// string 上一页响应中的 next_cursor，原样传入以翻页
func (api *GammaListMarketsKeysetAPI) AfterCursor(afterCursor string) *GammaListMarketsKeysetAPI {
	api.req.AfterCursor = GetPointer(afterCursor)
	return api
}

// int64 勿传；出现在 query 将 422，请用 AfterCursor
func (api *GammaListMarketsKeysetAPI) Offset(offset int64) *GammaListMarketsKeysetAPI {
	api.req.Offset = GetPointer(offset)
	return api
}

// []int64 市场 id 列表
func (api *GammaListMarketsKeysetAPI) ID(id []int64) *GammaListMarketsKeysetAPI {
	api.req.ID = GetPointer(id)
	return api
}

// []string 市场 slug 列表
func (api *GammaListMarketsKeysetAPI) Slug(slug []string) *GammaListMarketsKeysetAPI {
	api.req.Slug = GetPointer(slug)
	return api
}

// bool 是否包含已关闭；OpenAPI 默认 false
func (api *GammaListMarketsKeysetAPI) Closed(closed bool) *GammaListMarketsKeysetAPI {
	api.req.Closed = GetPointer(closed)
	return api
}

// bool 是否以小数化等形式返回部分数值展示
func (api *GammaListMarketsKeysetAPI) Decimalized(decimalized bool) *GammaListMarketsKeysetAPI {
	api.req.Decimalized = GetPointer(decimalized)
	return api
}

// []string CLOB token id 列表
func (api *GammaListMarketsKeysetAPI) ClobTokenIds(clobTokenIds []string) *GammaListMarketsKeysetAPI {
	api.req.ClobTokenIds = GetPointer(clobTokenIds)
	return api
}

// []string condition id 列表
func (api *GammaListMarketsKeysetAPI) ConditionIds(conditionIds []string) *GammaListMarketsKeysetAPI {
	api.req.ConditionIds = GetPointer(conditionIds)
	return api
}

// []string 问题 id 列表
func (api *GammaListMarketsKeysetAPI) QuestionIds(questionIds []string) *GammaListMarketsKeysetAPI {
	api.req.QuestionIds = GetPointer(questionIds)
	return api
}

// []string 做市地址列表
func (api *GammaListMarketsKeysetAPI) MarketMakerAddress(marketMakerAddress []string) *GammaListMarketsKeysetAPI {
	api.req.MarketMakerAddress = GetPointer(marketMakerAddress)
	return api
}

// float64 流动性数值下限
func (api *GammaListMarketsKeysetAPI) LiquidityNumMin(liquidityNumMin float64) *GammaListMarketsKeysetAPI {
	api.req.LiquidityNumMin = GetPointer(liquidityNumMin)
	return api
}

// float64 流动性数值上限
func (api *GammaListMarketsKeysetAPI) LiquidityNumMax(liquidityNumMax float64) *GammaListMarketsKeysetAPI {
	api.req.LiquidityNumMax = GetPointer(liquidityNumMax)
	return api
}

// float64 成交量数值下限
func (api *GammaListMarketsKeysetAPI) VolumeNumMin(volumeNumMin float64) *GammaListMarketsKeysetAPI {
	api.req.VolumeNumMin = GetPointer(volumeNumMin)
	return api
}

// float64 成交量数值上限
func (api *GammaListMarketsKeysetAPI) VolumeNumMax(volumeNumMax float64) *GammaListMarketsKeysetAPI {
	api.req.VolumeNumMax = GetPointer(volumeNumMax)
	return api
}

// string 开始时间下限，date-time 格式
func (api *GammaListMarketsKeysetAPI) StartDateMin(startDateMin string) *GammaListMarketsKeysetAPI {
	api.req.StartDateMin = GetPointer(startDateMin)
	return api
}

// string 开始时间上限，date-time 格式
func (api *GammaListMarketsKeysetAPI) StartDateMax(startDateMax string) *GammaListMarketsKeysetAPI {
	api.req.StartDateMax = GetPointer(startDateMax)
	return api
}

// string 结束时间下限，date-time 格式
func (api *GammaListMarketsKeysetAPI) EndDateMin(endDateMin string) *GammaListMarketsKeysetAPI {
	api.req.EndDateMin = GetPointer(endDateMin)
	return api
}

// string 结束时间上限，date-time 格式
func (api *GammaListMarketsKeysetAPI) EndDateMax(endDateMax string) *GammaListMarketsKeysetAPI {
	api.req.EndDateMax = GetPointer(endDateMax)
	return api
}

// []int64 标签 id 列表（与 List markets 的单 int tag_id 不同）
func (api *GammaListMarketsKeysetAPI) TagID(tagID []int64) *GammaListMarketsKeysetAPI {
	api.req.TagID = GetPointer(tagID)
	return api
}

// bool 相关标签是否参与过滤
func (api *GammaListMarketsKeysetAPI) RelatedTags(relatedTags bool) *GammaListMarketsKeysetAPI {
	api.req.RelatedTags = GetPointer(relatedTags)
	return api
}

// string 多标签时的匹配策略（取值以服务端为准）
func (api *GammaListMarketsKeysetAPI) TagMatch(tagMatch string) *GammaListMarketsKeysetAPI {
	api.req.TagMatch = GetPointer(tagMatch)
	return api
}

// bool cyom 标记
func (api *GammaListMarketsKeysetAPI) Cyom(cyom bool) *GammaListMarketsKeysetAPI {
	api.req.Cyom = GetPointer(cyom)
	return api
}

// bool 是否只保留启用 RFQ 的市场
func (api *GammaListMarketsKeysetAPI) RfqEnabled(rfqEnabled bool) *GammaListMarketsKeysetAPI {
	api.req.RfqEnabled = GetPointer(rfqEnabled)
	return api
}

// string UMA 结算状态
func (api *GammaListMarketsKeysetAPI) UmaResolutionStatus(umaResolutionStatus string) *GammaListMarketsKeysetAPI {
	api.req.UmaResolutionStatus = GetPointer(umaResolutionStatus)
	return api
}

// string 比赛/对局 id
func (api *GammaListMarketsKeysetAPI) GameID(gameID string) *GammaListMarketsKeysetAPI {
	api.req.GameID = GetPointer(gameID)
	return api
}

// []string 体育市场类型列表
func (api *GammaListMarketsKeysetAPI) SportsMarketTypes(sportsMarketTypes []string) *GammaListMarketsKeysetAPI {
	api.req.SportsMarketTypes = GetPointer(sportsMarketTypes)
	return api
}

// bool 为 true 时在每条 market 上展开 Tags
func (api *GammaListMarketsKeysetAPI) IncludeTag(includeTag bool) *GammaListMarketsKeysetAPI {
	api.req.IncludeTag = GetPointer(includeTag)
	return api
}

// string 区域/语言偏好，具体格式以服务端为准
func (api *GammaListMarketsKeysetAPI) Locale(locale string) *GammaListMarketsKeysetAPI {
	api.req.Locale = GetPointer(locale)
	return api
}

type GammaGetMarketByIDAPI struct {
	client *GammaRestClient
	id     *int64
	req    *GammaGetMarketByIDReq
}

type GammaGetMarketByIDReq struct {
	IncludeTag *bool `json:"include_tag"` // boolean 为 true 时在 market 上展开 Tags
}

// int64 市场 ID（路径参数）
func (api *GammaGetMarketByIDAPI) ID(id int64) *GammaGetMarketByIDAPI {
	api.id = GetPointer(id)
	return api
}

// bool 为 true 时在 market 上展开 Tags
func (api *GammaGetMarketByIDAPI) IncludeTag(includeTag bool) *GammaGetMarketByIDAPI {
	api.req.IncludeTag = GetPointer(includeTag)
	return api
}

type GammaGetMarketBySlugAPI struct {
	client *GammaRestClient
	slug   *string
	req    *GammaGetMarketBySlugReq
}

type GammaGetMarketBySlugReq struct {
	IncludeTag *bool `json:"include_tag"` // boolean 为 true 时在 market 上展开 Tags（与 Get market by id 一致）
}

// string 市场 slug（路径参数 `{slug}`）
func (api *GammaGetMarketBySlugAPI) Slug(slug string) *GammaGetMarketBySlugAPI {
	api.slug = GetPointer(slug)
	return api
}

// bool 为 true 时在 market 上展开 Tags
func (api *GammaGetMarketBySlugAPI) IncludeTag(includeTag bool) *GammaGetMarketBySlugAPI {
	api.req.IncludeTag = GetPointer(includeTag)
	return api
}

type GammaGetMarketTagsAPI struct {
	client *GammaRestClient
	id     *int64
}

// int64 市场 ID（路径参数）
func (api *GammaGetMarketTagsAPI) ID(id int64) *GammaGetMarketTagsAPI {
	api.id = GetPointer(id)
	return api
}

type GammaGetEventByIDAPI struct {
	client *GammaRestClient
	id     *int64
	req    *GammaGetEventByIDReq
}

type GammaGetEventByIDReq struct {
	IncludeChat     *bool `json:"include_chat"`     // 为 true 时包含 Chats 和 Series.Chats 关联
	IncludeTemplate *bool `json:"include_template"` // 为 true 时包含 Templates 关联
}

// 事件 ID（路径参数）
func (api *GammaGetEventByIDAPI) ID(id int64) *GammaGetEventByIDAPI {
	api.id = GetPointer(id)
	return api
}

// 为 true 时包含 Chats 和 Series.Chats 关联
func (api *GammaGetEventByIDAPI) IncludeChat(includeChat bool) *GammaGetEventByIDAPI {
	api.req.IncludeChat = GetPointer(includeChat)
	return api
}

// 为 true 时包含 Templates 关联
func (api *GammaGetEventByIDAPI) IncludeTemplate(includeTemplate bool) *GammaGetEventByIDAPI {
	api.req.IncludeTemplate = GetPointer(includeTemplate)
	return api
}

type GammaGetEventBySlugAPI struct {
	client *GammaRestClient
	slug   *string
	req    *GammaGetEventBySlugReq
}

type GammaGetEventBySlugReq struct {
	IncludeChat     *bool `json:"include_chat"`     // 为 true 时包含 Chats 和 Series.Chats 关联
	IncludeTemplate *bool `json:"include_template"` // 为 true 时包含 Templates 关联
}

// 事件 slug（路径参数）
func (api *GammaGetEventBySlugAPI) Slug(slug string) *GammaGetEventBySlugAPI {
	api.slug = GetPointer(slug)
	return api
}

// 为 true 时包含 Chats 和 Series.Chats 关联
func (api *GammaGetEventBySlugAPI) IncludeChat(includeChat bool) *GammaGetEventBySlugAPI {
	api.req.IncludeChat = GetPointer(includeChat)
	return api
}

// 为 true 时包含 Templates 关联
func (api *GammaGetEventBySlugAPI) IncludeTemplate(includeTemplate bool) *GammaGetEventBySlugAPI {
	api.req.IncludeTemplate = GetPointer(includeTemplate)
	return api
}

type GammaGetEventTagsAPI struct {
	client *GammaRestClient
	id     *int64
}

// 事件 ID（路径参数）
func (api *GammaGetEventTagsAPI) ID(id int64) *GammaGetEventTagsAPI {
	api.id = GetPointer(id)
	return api
}

type GammaGetPublicProfileAPI struct {
	client *GammaRestClient
	req    *GammaGetPublicProfileReq
}

// GammaGetPublicProfileReq 对应 GET /public-profile 的 query（OpenAPI 参数名 `address`）。
type GammaGetPublicProfileReq struct {
	Address *string `json:"address"` // string 钱包地址（proxy wallet 或 user address），须符合 ^0x[a-fA-F0-9]{40}$
}

// string 钱包地址（0x 前缀 + 40 位 hex），作为 query `address` 传入
func (api *GammaGetPublicProfileAPI) Address(address string) *GammaGetPublicProfileAPI {
	api.req.Address = GetPointer(address)
	return api
}
