package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	
	//set vars 
	intro := "Welcome to a touch typing game. Try to type out the following sentences without looking at your keyboard. As you move along it will get progressively harder"
	outro := "Congratulations, you have won the game. best of luck with your typing endevours"
	
	questions:= setQuestions() //setQuestions() defined in typingModel.go
	level := 0;
	finalLevel := 3	;
	tries := 0;
	//intro 
	fmt.Println(intro)
	fmt.Println(questions[level].text)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tries++
		if (questions[level].checkCorrect(scanner.Text())){
			fmt.Println("Correct")
			fmt.Println("level complete!")
			level++
			if(level == finalLevel){
				if (tries == level){
					fmt.Println("bang on! you didn't get it wrong once")
				}else{
					fmt.Println("You had to retry ", tries - level, " times" )
				}
				fmt.Println(outro)
				break
			}
			fmt.Println("Your next sentence is\n \t", questions[level].text)

		} else {
			fmt.Println("try again, the sentence you want to type is\n \t", questions[level].text)

		}
			
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

