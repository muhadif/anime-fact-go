package module

import (
	"context"
	"github.com/muhadif/anime-fact/core/entity"
	"github.com/muhadif/anime-fact/core/repository"
)

func NewAnime(animeRepo repository.AnimeQuote) Anime {
	return &anime{animeRepo : animeRepo}
}

type Anime interface {
	GetAnime(ctx context.Context) (*entity.GetAnimeResponse, error)
	GetAnimeRandomQuote(ctx context.Context, totalData int32) (*entity.GetAnimeRandomQuoteResponse, error)
}

type anime struct {
	animeRepo repository.AnimeQuote
}

func (a anime) GetAnimeRandomQuote(ctx context.Context, totalData int32) (*entity.GetAnimeRandomQuoteResponse, error) {
	var filteredAnimeQuote []*entity.AnimeQuote
	for {
		resp, err := a.animeRepo.GetRandomAnimeQuote(ctx)
		if err != nil {
			return nil, err
		}

		for _, animeQuote := range resp.AnimeQuotes {
			if len(animeQuote.Character) > 10 || len(filteredAnimeQuote) >= int(totalData) {
				continue
			}

			filteredAnimeQuote = append(filteredAnimeQuote, animeQuote)
		}

		if len(filteredAnimeQuote) >= int(totalData) {
			break
		}
	}

	return &entity.GetAnimeRandomQuoteResponse{AnimeQuotes: filteredAnimeQuote}, nil
}

func (a anime) GetAnime(ctx context.Context) (*entity.GetAnimeResponse, error) {
	return &entity.GetAnimeResponse{
		Anime: []*entity.Anime{
			{
				ID:    "ANIME=1",
				Title: "Naruto",
			},
		},
	}, nil
}



