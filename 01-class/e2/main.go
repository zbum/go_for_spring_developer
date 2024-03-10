package main

import (
	"go_for_spring_developer/01-class/e2/bean"
	"log"
)

func main() {
	scoreRepository := bean.NewScoreRepository()
	scoreService := bean.NewScoreService(scoreRepository)
	log.Printf("%v \n", scoreService.GetScores())
}
