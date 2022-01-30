package main

import (
	"errors"
	"fmt"
	"os"
)

// Pokemon はポケモンの構造体
type Pokemon struct {
	ID    int
	Name  string
	moves []Move
}

// Moves はポケモンの持ってる技を回す関数
func (p Pokemon) Moves(isDynamax bool) ([]Move, error) {
	if isDynamax {
		dynamaxMoves := []Move{}

		for _, move := range p.moves {
			dynMove, err := move.dynamax()
			if err != nil {
				return nil, err
			}
			dynamaxMoves = append(dynamaxMoves, dynMove)
		}
		return dynamaxMoves, nil
	}
	return p.moves, nil
}

// Move は技の構造体
type Move struct {
	Name string
	Type string
}

func (m Move) dynamax() (Move, error) {
	switch m.Type {
	case "くさ":
		return Move{
			Name: "ダイソウゲン",
			Type: m.Type,
		}, nil
	case "ほのお":
		return Move{
			Name: "ダイバーン",
			Type: m.Type,
		}, nil
	case "みず":
		return Move{
			Name: "ダイストリーム",
			Type: m.Type,
		}, nil
	default:
		return Move{}, errors.New("unknown type")
	}
}

var party = []Pokemon{
	{ID: 3, Name: "フシギバナ", moves: []Move{{Name: "つるのむち", Type: "くさ"}}},
	{ID: 6, Name: "リザードン", moves: []Move{{Name: "かえんほうしゃ", Type: "ほのお"}}},
	{ID: 9, Name: "カメックス", moves: []Move{{Name: "みずでっぽう", Type: "みず"}}},
}

func main() {
	venusaur := party[0] //「Venusaur」って英語で「フシギバナ」ですって。
	fmt.Println("ポケモン:", venusaur.Name)

	moves, err := venusaur.Moves(false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("通常わざ", moves[0].Name)

	dynMoves, err := venusaur.Moves(true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("ダイマックスわざ:", dynMoves[0].Name)
}
