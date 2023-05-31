package anime_quote

type AnimeRandomQuoteResponse []struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}
