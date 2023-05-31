package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muhadif/anime-fact/core/module"
	"net/http"
	"strconv"
)

type AnimeHandler interface {
	GetAnime(ctx *gin.Context)
	GetAnimeRandomQuote(ctx *gin.Context)
}

func NewAnimeHandler(animeModule module.Anime) AnimeHandler {
	return &animeHandler{animeModule: animeModule}
}

type animeHandler struct {
	animeModule module.Anime
}

func (a animeHandler) GetAnimeRandomQuote(ctx *gin.Context) {
	totalDataStr := ctx.Query("totalData")
	totalData, err := strconv.Atoi(totalDataStr)

	resp, err := a.animeModule.GetAnimeRandomQuote(ctx, int32(totalData))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		ctx.Abort()
	}

	ctx.JSON(http.StatusOK, resp)
}

func (a animeHandler) GetAnime(ctx *gin.Context) {
	resp, err := a.animeModule.GetAnime(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		ctx.Abort()
	}

	ctx.JSON(http.StatusOK, resp)
}
