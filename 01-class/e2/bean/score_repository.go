package bean

import "fmt"

type ScoreRepository struct {
}

func (u *ScoreRepository) FindAll() []Score {
	return []Score{
		{1, 1, Student{1, "Manty"}, 100},
		{2, 2, Student{1, "Manty"}, 99},
		{3, 3, Student{1, "Manty"}, 98},
	}
}

func NewScoreRepository() *ScoreRepository {
	fmt.Println("init UserRepository")
	return new(ScoreRepository)
}

type Score struct {
	id       int64
	semester int
	student  Student
	score    int
}

type Student struct {
	id   int64
	name string
}
