package entity

type GetAnimeResponse struct {
	Anime []*Anime
}

type Anime struct {
	ID		string
	Title	string
}

type GetAnimeRandomQuoteResponse struct {
	AnimeQuotes []*AnimeQuote
}

type AnimeQuote struct {
	Anime	string
	Character string
	Quote	string
}