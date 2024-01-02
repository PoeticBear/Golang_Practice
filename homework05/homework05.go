package homework05

type Student struct {
	Name         string
	Age          int
	Gender       string
	Courses      []Course
	ClassMonitor ClassMonitor
}

type Course struct {
	Name  string
	Score float64
}

type ClassMonitor struct {
	Name   string
	Duty   string
	Access bool
}

func addStudent(student Student) {
	students := make([]Student, 0)
	students = append(students, student)
}

func FindStudentByName(name string) *Student {
	students := make([]Student, 0)
	students = append(students, Student{
		Name: "张三",
		Age:  18,
	})
	for i, student := range students {
		if student.Name == name {
			return &students[i]
		}
	}
	return nil
}

func SetClassMonitor(student *Student, classMonitor ClassMonitor) {
	student.ClassMonitor = classMonitor
}

func (s *Student) CalculateAverageScore() float64 {
	totalScore := 0.0
	for _, course := range s.Courses {
		totalScore += course.Score
	}
	return float64(totalScore) / float64(len(s.Courses))
}
