package mypolymarketapi

import (
	"errors"
	"strconv"
	"strings"
)

// GET List events
func (c *GammaRestClient) NewGammaListEvents() *GammaListEventsAPI {
	return &GammaListEventsAPI{
		client: c,
		req:    &GammaListEventsReq{},
	}
}

func (api *GammaListEventsAPI) Do() (*PolyMarketRestRes[GammaListEventsRes], error) {
	url := pmHandlerRequestAPIWithPathQueryParam(REST, GAMMA_REST, api.req, GammaAPITypeMap[GammaListEvents])
	return pmCallAPI[GammaListEventsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET List events (keyset pagination)
func (c *GammaRestClient) NewGammaListEventsKeyset() *GammaListEventsKeysetAPI {
	return &GammaListEventsKeysetAPI{
		client: c,
		req:    &GammaListEventsKeysetReq{},
	}
}

func (api *GammaListEventsKeysetAPI) Do() (*PolyMarketRestRes[GammaListEventsKeysetRes], error) {
	url := pmHandlerRequestAPIWithPathQueryParam(REST, GAMMA_REST, api.req, GammaAPITypeMap[GammaListEventsKeyset])
	return pmCallAPI[GammaListEventsKeysetRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get event by id
func (c *GammaRestClient) NewGammaGetEventByID() *GammaGetEventByIDAPI {
	return &GammaGetEventByIDAPI{
		client: c,
		req:    &GammaGetEventByIDReq{},
	}
}

func (api *GammaGetEventByIDAPI) Do() (*PolyMarketRestRes[GammaGetEventByIDRes], error) {
	if api.id == nil {
		return nil, errors.New("id is required")
	}
	path := strings.Replace(GammaAPITypeMap[GammaGetEventByID], "{id}", strconv.FormatInt(*api.id, BIT_BASE_10), 1)
	url := pmHandlerRequestAPIWithPathQueryParam(REST, GAMMA_REST, api.req, path)
	return pmCallAPI[GammaGetEventByIDRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get event by slug
func (c *GammaRestClient) NewGammaGetEventBySlug() *GammaGetEventBySlugAPI {
	return &GammaGetEventBySlugAPI{
		client: c,
		req:    &GammaGetEventBySlugReq{},
	}
}

func (api *GammaGetEventBySlugAPI) Do() (*PolyMarketRestRes[GammaGetEventBySlugRes], error) {
	if api.slug == nil {
		return nil, errors.New("slug is required")
	}
	path := strings.Replace(GammaAPITypeMap[GammaGetEventBySlug], "{slug}", *api.slug, 1)
	url := pmHandlerRequestAPIWithPathQueryParam(REST, GAMMA_REST, api.req, path)
	return pmCallAPI[GammaGetEventBySlugRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get event tags
func (c *GammaRestClient) NewGammaGetEventTags() *GammaGetEventTagsAPI {
	return &GammaGetEventTagsAPI{
		client: c,
	}
}

func (api *GammaGetEventTagsAPI) Do() (*PolyMarketRestRes[GammaGetEventTagsRes], error) {
	if api.id == nil {
		return nil, errors.New("id is required")
	}
	path := strings.Replace(GammaAPITypeMap[GammaGetEventTags], "{id}", strconv.FormatInt(*api.id, BIT_BASE_10), 1)
	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, GAMMA_REST, path)
	return pmCallAPI[GammaGetEventTagsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET List markets
func (c *GammaRestClient) NewGammaListMarkets() *GammaListMarketsAPI {
	return &GammaListMarketsAPI{
		client: c,
		req:    &GammaListMarketsReq{},
	}
}

func (api *GammaListMarketsAPI) Do() (*PolyMarketRestRes[GammaListMarketsRes], error) {
	url := pmHandlerRequestAPIWithPathQueryParam(REST, GAMMA_REST, api.req, GammaAPITypeMap[GammaListMarkets])
	return pmCallAPI[GammaListMarketsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET List markets (keyset pagination)
func (c *GammaRestClient) NewGammaListMarketsKeyset() *GammaListMarketsKeysetAPI {
	return &GammaListMarketsKeysetAPI{
		client: c,
		req:    &GammaListMarketsKeysetReq{},
	}
}

func (api *GammaListMarketsKeysetAPI) Do() (*PolyMarketRestRes[GammaListMarketsKeysetRes], error) {
	url := pmHandlerRequestAPIWithPathQueryParam(REST, GAMMA_REST, api.req, GammaAPITypeMap[GammaListMarketsKeyset])
	return pmCallAPI[GammaListMarketsKeysetRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get market by id
func (c *GammaRestClient) NewGammaGetMarketByID() *GammaGetMarketByIDAPI {
	return &GammaGetMarketByIDAPI{
		client: c,
		req:    &GammaGetMarketByIDReq{},
	}
}

func (api *GammaGetMarketByIDAPI) Do() (*PolyMarketRestRes[GammaGetMarketByIDRes], error) {
	if api.id == nil {
		return nil, errors.New("id is required")
	}
	path := strings.Replace(GammaAPITypeMap[GammaGetMarketByID], "{id}", strconv.FormatInt(*api.id, BIT_BASE_10), 1)
	url := pmHandlerRequestAPIWithPathQueryParam(REST, GAMMA_REST, api.req, path)
	return pmCallAPI[GammaGetMarketByIDRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get market by slug（Gamma：`GET /markets/slug/{slug}`；可选 query `include_tag`）
func (c *GammaRestClient) NewGammaGetMarketBySlug() *GammaGetMarketBySlugAPI {
	return &GammaGetMarketBySlugAPI{
		client: c,
		req:    &GammaGetMarketBySlugReq{},
	}
}

func (api *GammaGetMarketBySlugAPI) Do() (*PolyMarketRestRes[GammaGetMarketBySlugRes], error) {
	if api.slug == nil {
		return nil, errors.New("slug is required")
	}
	path := strings.Replace(GammaAPITypeMap[GammaGetMarketBySlug], "{slug}", *api.slug, 1)
	url := pmHandlerRequestAPIWithPathQueryParam(REST, GAMMA_REST, api.req, path)
	return pmCallAPI[GammaGetMarketBySlugRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get market tags by id（Gamma：`GET /markets/{id}/tags`）
func (c *GammaRestClient) NewGammaGetMarketTags() *GammaGetMarketTagsAPI {
	return &GammaGetMarketTagsAPI{
		client: c,
	}
}

func (api *GammaGetMarketTagsAPI) Do() (*PolyMarketRestRes[GammaGetMarketTagsRes], error) {
	if api.id == nil {
		return nil, errors.New("id is required")
	}
	path := strings.Replace(GammaAPITypeMap[GammaGetMarketTags], "{id}", strconv.FormatInt(*api.id, BIT_BASE_10), 1)
	url := pmHandlerRequestAPIWithoutPathQueryParam(REST, GAMMA_REST, path)
	return pmCallAPI[GammaGetMarketTagsRes](api.client.c, url, NIL_REQBODY, GET)
}

// GET Get public profile by wallet address（Gamma：`GET /public-profile`；必填 query `address`，为 proxy wallet 或用户链上地址，须符合 `^0x[a-fA-F0-9]{40}$`）
func (c *GammaRestClient) NewGammaGetPublicProfile() *GammaGetPublicProfileAPI {
	return &GammaGetPublicProfileAPI{
		client: c,
		req:    &GammaGetPublicProfileReq{},
	}
}

func (api *GammaGetPublicProfileAPI) Do() (*PolyMarketRestRes[GammaGetPublicProfileRes], error) {
	if api.req.Address == nil {
		return nil, errors.New("address is required")
	}
	url := pmHandlerRequestAPIWithPathQueryParam(REST, GAMMA_REST, api.req, GammaAPITypeMap[GammaGetPublicProfile])
	return pmCallAPI[GammaGetPublicProfileRes](api.client.c, url, NIL_REQBODY, GET)
}
