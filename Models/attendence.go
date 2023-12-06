package models

type Attendance struct {
	AttendanceID int    `bson:"attendanceId" json:"attendanceId"`
	LessonID     int    `bson:"lessonId" json:"lessonId"`
	CourseID     string `bson:"courseId" json:"courseId"`
	UserID       int    `bson:"userId" json:"userId"`
	IsPresent    bool   `bson:"isPresent" json:"isPresent"`
}
