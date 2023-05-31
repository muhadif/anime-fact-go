package anime_quote

import (
	"context"
	"encoding/json"
	"github.com/muhadif/anime-fact/core/entity"
	"github.com/muhadif/anime-fact/core/repository"
	"io/ioutil"
	"net/http"
)

func NewAnimeQuoteRepository() repository.AnimeQuote {
	return &repo{}
}

type repo struct {

}

func (r repo) GetRandomAnimeQuote(ctx context.Context) (*entity.GetAnimeRandomQuoteResponse, error) {
	url := "https://animechan.vercel.app/api/quotes"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var response AnimeRandomQuoteResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return parseToAnimeRandomQuoteResponse(response)
}

func parseToAnimeRandomQuoteResponse(clientResponse AnimeRandomQuoteResponse) (*entity.GetAnimeRandomQuoteResponse, error) {
	var animeQuotes []*entity.AnimeQuote
	for _, anime := range clientResponse {
		animeQuotes = append(animeQuotes, &entity.AnimeQuote{
			Anime:     anime.Anime,
			Character: anime.Character,
			Quote:     anime.Quote,
		})
	}
	return &entity.GetAnimeRandomQuoteResponse{AnimeQuotes: animeQuotes}, nil
}

