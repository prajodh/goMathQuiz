package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)



const quizFile string = "./quiz.csv"
func main(){
	//get the command line args if present
	timeRequired := flag.Int("time", 10, "enter the time you need to complete the quiz")
	filepath := flag.String("csv", quizFile, "enter the path of ur csv file to load the quiz")
	flag.Parse()

	//intialize the time object
	timerObj := time.NewTimer(time.Second*time.Duration(*timeRequired))
	

	//load the csv file and read it line by line
	loadFile, err := os.Open(*filepath)
	if err != nil{
		log.Fatal(err)
	}
	fileReader := csv.NewReader(loadFile)
	var noOfquestions int = 0
	var noOfCorrectQuestions int = 0
	for {
		//read the file to load questions and answers
		line, err := fileReader.Read()
		if err == io.EOF{
			fmt.Println("")
			fmt.Println("******** Game Over *********")
			fmt.Println("No of Correct answers given", noOfCorrectQuestions)
			fmt.Println("Total no of questions asked", noOfquestions)
			return
		}else if err != nil{
			log.Fatal(err)
		}
		question := line[0]
		answer := line[1]
		fmt.Printf("%s :",question)
		noOfquestions+=1
		answerC := make(chan string)

		//make the user input non bloaking
		go func(){
		//get input from user for the answer
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		reply := scanner.Text()
		answerC<-reply
		}()

		//select between the two channels
		select {
		case <- timerObj.C:
			fmt.Println(" ")
			fmt.Println("******** Game Over *********")
			fmt.Println("You Ran Out Of Time")
			fmt.Println("No of Correct answers given", noOfCorrectQuestions)
			fmt.Println("Total no of questions asked", noOfquestions)
			return
		case ans := <-answerC:
			if answer == ans{
			noOfCorrectQuestions+=1
		}

		}
		
	}

}