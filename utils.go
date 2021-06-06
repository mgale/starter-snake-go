package main

import (
	"fmt"
	"math/rand"
)

func determineMove(gameRequest GameRequest) MoveResponse {
	// Choose a random direction to move in
	possibleMoves := []string{"up", "down", "left", "right"}
	move := possibleMoves[rand.Intn(len(possibleMoves))]

	fmt.Printf("GS: Move: %v, Head: %v, Body: %v",
		gameRequest.Turn,
		gameRequest.You.Head,
		gameRequest.You.Body,
	)

	if gameRequest.You.Head.X == 0 {
		move = "left"
	}

	if gameRequest.You.Head.Y == 0 {
		move = "up"
	}

	if gameRequest.You.Head.X == gameRequest.Board.Height {
		move = "right"
	}

	if gameRequest.You.Head.Y == gameRequest.Board.Width {
		move = "left"
	}

	response := MoveResponse{
		Move: move,
	}

	return response
}
