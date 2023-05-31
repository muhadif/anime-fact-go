package router

import (
	"github.com/gin-gonic/gin"
	"github.com/muhadif/anime-fact/handler"
)

func NewRouter(animeHandler handler.AnimeHandler) *gin.Engine {
	routes := gin.Default()

	routes.Handle("GET", "/", animeHandler.GetAnime)
	routes.Handle("GET", "/quote", animeHandler.GetAnimeRandomQuote)

	return routes
}
