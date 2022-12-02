package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Symbol string

const (
	Rock     Symbol = "rock"
	Paper    Symbol = "paper"
	Scissors Symbol = "scissors"
)

type GameResult string

const (
	IWin         GameResult = "I win!"
	OpponentWins GameResult = "Opponent wins :("
	Draw         GameResult = "Draw"
)

var loseTo = map[Symbol]Symbol{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

var winsWith = map[Symbol]Symbol{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

func whoWins(me Symbol, opponent Symbol) GameResult {
	if me == opponent {
		return Draw
	} else if loseTo[me] == opponent {
		return IWin
	} else {
		return OpponentWins
	}
}

var opponentsPlaysDict = map[string]Symbol{
	"A": Rock, "B": Paper, "C": Scissors,
}

var scorePerSymbol = map[Symbol]int{
	Rock: 1, Paper: 2, Scissors: 3,
}

var scorePerResult = map[GameResult]int{
	IWin: 6, Draw: 3, OpponentWins: 0,
}

func CalculateScoreWithGuessedStrategy(input *os.File) (string, error) {
	var guessedStrategyDict = map[string]Symbol{
		"X": Rock, "Y": Paper, "Z": Scissors,
	}
	return calculateScoreWithStrategy(input, func(code string, opponentsSymbol Symbol) Symbol {
		return guessedStrategyDict[code]
	})
}

func CalculateScoreWithCorrectStrategy(input *os.File) (string, error) {
	var suggestedResult = map[string]GameResult{
		"X": OpponentWins, "Y": Draw, "Z": IWin,
	}
	return calculateScoreWithStrategy(input, func(code string, opponentsSymbol Symbol) Symbol {
		suggestedResult := suggestedResult[code]
		if suggestedResult == OpponentWins {
			return loseTo[opponentsSymbol]
		} else if suggestedResult == Draw {
			return opponentsSymbol
		} else {
			return winsWith[opponentsSymbol]
		}
	})
}

func calculateScoreWithStrategy(input *os.File, strategy func(code string, opponentsSymbol Symbol) Symbol) (string, error) {
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	totalScore := 0

	for fileScanner.Scan() {
		round := strings.Split(fileScanner.Text(), " ")
		opponentsSymbol := opponentsPlaysDict[round[0]]
		mySymbol := strategy(round[1], opponentsSymbol)
		gameResult := whoWins(mySymbol, opponentsSymbol)
		myScoreThisRound := scorePerSymbol[mySymbol] + scorePerResult[gameResult]
		//fmt.Printf("Me: %s, Opponent: %s, Result: %s, Points: %d\n", mySymbol, opponentsSymbol, gameResult, myScoreThisRound)
		totalScore += myScoreThisRound
	}

	return strconv.Itoa(totalScore), nil
}
