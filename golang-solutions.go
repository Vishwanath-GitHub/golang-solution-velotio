package main

import (
  "fmt"
  "sort"
)

func main() {
	paraString:="( ((X)) (((Y))) )"
  fmt.Println(para(paraString))
  fmt.Println()

  string1:="abcd"
	string2:="dabc"
	fmt.Println(anagram(string1,string2))
  fmt.Println()

	numberOfTimes := 8
	fib(numberOfTimes)
}

//Parenthesis
func para(string1 string) int {
  var current int //ongoing count
  var finalcounter int //max count

  for i:=0; i<len(string1);i++ {
    if string1[i]=='(' {
      current++
      if current>finalcounter {
        finalcounter=current
      }
    } else if string1[i]==')' {
      if current>0 {
        current--
      } else {
        return -1
      }
    }
  }

  if current!=0 {
    return -1
  }

  return finalcounter
}

//Anagram
type RuneThis []rune
//Building functions to support 'sort.Sort()' function
func (runeValue RuneThis) Len() int {
  return len(runeValue)
}

func (runeValue RuneThis) Swap(value1,value2 int) {
  runeValue[value1],runeValue[value2]=runeValue[value2],runeValue[value1]
}

func (runeValue RuneThis) Less(value1,value2 int) bool{
  return runeValue[value1]<runeValue[value2]
}

func ConvertingStringToRune(s string) []rune {
  var ru []rune
  for _,runeValues:=range s {
    ru=append(ru, runeValues)
  }
  return ru
}

func anagram(string1, string2 string) bool {
	var finalResult bool
  var rune1 RuneThis= ConvertingStringToRune(string1)
  var rune2 RuneThis= ConvertingStringToRune(string2)
  sort.Sort(rune1)
  sort.Sort(rune2)

  if string(rune1)==string(rune2) {
    finalResult=true
  } else {
    finalResult=false
  }
  return finalResult
}

 // Fibonacci Series
func fib(numberOfTimes int) {
	value1 := 0
	value2 := 1

	var finalValue int

	fmt.Print(value1, ",", value2, ",")

	for i := 2; i < numberOfTimes; i++ {
		finalValue = value1 + value2
		fmt.Print(finalValue, ",")
		value1 = value2
		value2 = finalValue
	}
	fmt.Println()
}
