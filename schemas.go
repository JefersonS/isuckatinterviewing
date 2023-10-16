package main

type Question struct {
    ID     int64
    Question  string
    Answer string
    ItSucks  int32
	YouSuck int32
}

var questions = []Question{
	{
		Question: "This is a question",
		Answer:   "And this is the answer",
	},
	{
		Question: "This is another question",
		Answer:   "And this is another answer",
	},
}

func GetQuestion() Question {
	return questions[0]
}
