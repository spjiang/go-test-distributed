package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Nick",
			LastName:  "Carter",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 94,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 82,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Jon",
			LastName:  "Steven",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 81,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 92,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 87,
				},
			},
		},
		{
			ID:        3,
			FirstName: "Li",
			LastName:  "Na",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 96,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 99,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 98,
				},
			},
		},
	}
}
