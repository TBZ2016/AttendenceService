package main

import (
	AttendenceController "example/attendance/Controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/attendance", AttendenceController.GetAttendance)
	router.POST("/attendance", AttendenceController.CreateAttendance)
	router.GET("/attendance/{userId}", AttendenceController.GetAttendanceByUserID)
	router.PUT("/attendance/{userId}", AttendenceController.UpdateAttendanceByUserID)
	router.DELETE("/attendance/{userId}", AttendenceController.DeleteAttendanceByUserID)
	router.GET("/attendance/lesson/{lessonId}", AttendenceController.GetAttendanceByLessonID)
	router.GET("/attendance/course/{courseId}", AttendenceController.GetAttendanceByCourseID)

	router.Run("localhost:8080")
}
