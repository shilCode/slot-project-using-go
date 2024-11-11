package main

import (
	"fmt"
	"math/rand"
)



func generateSymbolArray(symbols map[string]uint)[]string{
	symbolArr :=[]string{}
	for symbol,count:=range symbols{
		for i:=uint(0);i<count;i++{
			symbolArr = append(symbolArr,symbol)
		}
	}
	return symbolArr;
}

func getRandomNumber(min int, max int)int{
	randomNumber := rand.Intn(max-min+1)+min
	return randomNumber
}
func getSpin(reel[]string,rows int,cols int)[][]string{
	result:=[][]string{}

	for i:=0;i<rows;i++{
		result=append(result, []string{})
	}
	for col:=0;col<cols;col++{
		selected :=map[int]bool{}
		for row:=0;row<rows;row++{
			for true{
				randomIndex:= getRandomNumber(0,len(reel)-1)
				_,exists := selected[randomIndex]
				if !exists{
					selected[randomIndex]=true
					result[row]= append(result[row],reel[randomIndex])
					break
				}
			}
		}
	}
	return result
}

func checkWin(spin[][]string, multipliers map[string]uint)[]uint{
	lines := []uint{}
	for _, row :=range spin{
		win:=true
		checkSymbol:=row[0]
		for _, symbol:=range row[1:]{
			if checkSymbol!=symbol{
				win=false
				break
			}
		}
		if win{
			lines = append(lines, multipliers[checkSymbol])
		}else{
			lines=append(lines, 0)
		}
	}
	return lines
}

func printSpin(spin[][]string){
	for _,row:=range spin{
		for j, symbol:=range row{
			fmt.Printf(symbol,)
			if j!= len(row)-1{
				fmt.Printf("|")
			}

		}
		fmt.Println()
	}
}

func main()  {

	symbols:=map[string]uint{
		"A":4,
		"B":7,
		"C":12,
		"D":20,
	}
	symbolsArr :=generateSymbolArray(symbols)
	multipliers:=map[string]uint{
		"A":20,
		"B":10,
		"C":5,
		"D":2,
		
	}


	balance := uint(200) 
	GetBet(balance)
	GetName()

	for balance>0{
		
		bet:=GetBet(uint(balance))
		if(bet==0){
			break
		}
		balance-=bet
		spin:=getSpin(symbolsArr,3,3)
		printSpin(spin)
		winningSpin:=checkWin(spin,multipliers)
		for i, multi:=range winningSpin{
			win:=multi*bet
			balance +=win
			if multi>0{
				fmt.Printf("Won $%d, (%dx) on line #%d\n ", win,multi,i+1)
			}
		}

		}

		fmt.Printf("You left with, %d\n",balance)

}