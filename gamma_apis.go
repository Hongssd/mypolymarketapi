package mypolymarketapi

type GammaAPIType int

const (
	// Events
	GammaListEventsKeyset GammaAPIType = iota // GET List events (keyset pagination)
	GammaListEvents                           // GET List events
	GammaGetEventByID                         // GET Get event by id
	GammaGetEventBySlug                       // GET Get event by slug
	GammaGetEventTags                         // GET Get event tags

	// Markets
	GammaListMarketsKeyset // GET List markets (keyset pagination)
	GammaListMarkets       // GET List markets
	GammaGetMarketByID     // GET Get market by id
	GammaGetMarketBySlug   // GET Get market by slug
	GammaGetMarketTags     // GET Get market tags by id

	// Profiles
	GammaGetPublicProfile // GET Get public profile by wallet address (query: address)
)

var GammaAPITypeMap = map[GammaAPIType]string{
	// Events
	GammaListEventsKeyset: "/events/keyset",      // GET List events (keyset pagination)
	GammaListEvents:       "/events",             // GET List events
	GammaGetEventByID:     "/events/{id}",        // GET Get event by id
	GammaGetEventBySlug:   "/events/slug/{slug}", // GET Get event by slug
	GammaGetEventTags:     "/events/{id}/tags",   // GET Get event tags

	// Markets
	GammaListMarketsKeyset: "/markets/keyset", // GET List markets (keyset pagination)
	GammaListMarkets:       "/markets",        // GET List markets
	GammaGetMarketByID:     "/markets/{id}",        // GET Get market by id
	GammaGetMarketBySlug:   "/markets/slug/{slug}", // GET Get market by slug
	GammaGetMarketTags:     "/markets/{id}/tags",   // GET Get market tags by id

	// Profiles
	GammaGetPublicProfile: "/public-profile", // GET Get public profile by wallet address (query: address)
}
