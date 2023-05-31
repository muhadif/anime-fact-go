package repository

import (
	"context"
	"github.com/muhadif/anime-fact/core/entity"
)

type AnimeQuote interface {
	GetRandomAnimeQuote(ctx context.Context) (*entity.GetAnimeRandomQuoteResponse, error)
}
