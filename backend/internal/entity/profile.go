package entity

type Profile struct {
	ID               int
	Username         string
	Email            string
	Status           string
	ImageUrl         string
	TasksCompleted   int
	ActiveCourses    []int
	CompletedCourses []int
}
