package dice

import (
	"math/rand"
	"sort"
)

type Dice struct {
	Value int
}

func NewDice() *Dice {
	return &Dice{}
}

func (d *Dice) Roll() {
	d.Value = rand.Intn(6) + 1
}

type Dicer []*Dice

func NewDicer() *Dicer {
	dicer := make(Dicer, 9)
	for i := range dicer {
		dicer[i] = NewDice()
	}
	return &dicer
}

func (d Dicer) Roll() {
	for _, dice := range d {
		dice.Roll()
	}
}

func (d Dicer) Values() []int {
	Values := make([]int, len(d))
	for i, dice := range d {
		Values[i] = dice.Value
	}
	sort.Ints(Values)
	return Values
}

func CalculatePremiumMultiplier(risk string, Values []int) int {
	var premium int

	switch risk {
	case "low":
		premium = lowRisk(Values)
	case "medium":
		premium = mediumRisk(Values)
	case "high":
		premium = highRisk(Values)
	}

	return premium
}

func lowRisk(Values []int) int {
	counts := make(map[int]int)

	// Count occurrences of Values between 1 and 6
	for _, v := range Values {
		if v >= 1 && v <= 6 {
			counts[v]++
		}
	}

	var points int
	// Assign points based on the counts
	for _, count := range counts {
		switch count {
		case 3:
			points += 1
		case 4:
			points += 2
		case 5:
			points += 5
		case 6:
			points += 10
		case 7:
			points += 25
		case 8:
			points += 100
		case 9:
			points += 500
		}
	}
	return points
}

func mediumRisk(Values []int) int {
	counts := make(map[int]int)

	// Count occurrences of Values between 1 and 6
	for _, v := range Values {
		if v >= 1 && v <= 6 {
			counts[v]++
		}
	}

	var points int
	for _, v := range counts {
		switch v {
		case 4:
			points += 2
		case 5:
			points += 5
		case 6:
			points += 10
		case 7:
			points += 50
		case 8:
			points += 250
		case 9:
			points += 1000
		}
	}
	return points
}

func highRisk(Values []int) int {
	counts := make(map[int]int)

	// Count occurrences of Values between 1 and 6
	for _, v := range Values {
		if v >= 1 && v <= 6 {
			counts[v]++
		}
	}

	var points int
	for _, v := range counts {
		switch v {
		case 5:
			points += 5
		case 6:
			points += 10
		case 7:
			points += 50
		case 8:
			points += 500
		case 9:
			points += 5000
		}
	}
	return points
}
