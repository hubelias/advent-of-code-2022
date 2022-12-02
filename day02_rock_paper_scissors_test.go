package main

import (
	"fmt"
	"testing"
)

func Test_whoWins(t *testing.T) {
	tests := []struct {
		me       Symbol
		opponent Symbol
		result   GameResult
	}{
		{me: Scissors, opponent: Scissors, result: Draw},
		{me: Scissors, opponent: Rock, result: OpponentWins},
		{me: Scissors, opponent: Paper, result: IWin},
		{me: Rock, opponent: Scissors, result: IWin},
		{me: Rock, opponent: Paper, result: OpponentWins},
		{me: Rock, opponent: Rock, result: Draw},
		{me: Paper, opponent: Scissors, result: OpponentWins},
		{me: Paper, opponent: Paper, result: Draw},
		{me: Paper, opponent: Rock, result: IWin},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			if got := whoWins(tt.me, tt.opponent); got != tt.result {
				t.Errorf("whoWins() = %v, want %v", got, tt.result)
			}
		})
	}
}
