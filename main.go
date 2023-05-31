package main

import (
	"fmt"
	"github.com/muhadif/anime-fact/core/module"
	"github.com/muhadif/anime-fact/handler"
	anime_quote "github.com/muhadif/anime-fact/repository/anime-quote"
	"github.com/muhadif/anime-fact/router"
)

func main()  {
	animeRepo := anime_quote.NewAnimeQuoteRepository()
	animeModule := module.NewAnime(animeRepo)
	animeHandler := handler.NewAnimeHandler(animeModule)

	routes := router.NewRouter(animeHandler)

	err := routes.Run()
	if err != nil {
		fmt.Println("error run : ", err.Error())
	}
}