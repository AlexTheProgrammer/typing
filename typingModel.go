package main

type question struct{
	text string
}

func (q question) checkCorrect(s string) bool{
	return q.text == s
}

//set questions
func setQuestions() []question{
	question1 := question{text:"This is your first sentence",}
	question2 := question{text:"Here is a second sentence",}
	question3 := question{text:"Try something a bit different",}
	
	return []question{question1, question2, question3}
	}
