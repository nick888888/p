package pt

import (
	"testing"
	"fmt"
)

func TestStringA(t *testing.T) {
	StringA()
}

func TestStringB(t *testing.T) {
	StringB()
}

func TestStringC(t *testing.T) {
	StringC()
}

func TestStringToUpper(t *testing.T) {
	a, b := StringToUpper("abc")
	fmt.Println("111--" + a)
	fmt.Println("222--" + b)
}

func TestStringD(t *testing.T) {
	StringD()
}

func TestStringE(t *testing.T) {
	StringRuneE()
}

func TestStringRuneF(t *testing.T) {
	StringRuneF()
}

func TestStringRuneG(t *testing.T) {
	StringRuneG()
}