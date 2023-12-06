package AttendenceController

import (
	dal "example/attendance/DAL"
	models "example/attendance/Models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAttendance(c *gin.Context) {
	attendanceRecords, err := dal.GetAllAttendance()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendanceRecords)
}

func CreateAttendance(c *gin.Context) {
	var newRecords []models.Attendance
	if err := c.ShouldBindJSON(&newRecords); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := dal.InsertAttendance(newRecords)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "Attendance recorded")
}

func GetAttendanceByUserID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userId"})
		return
	}

	userAttendance, err := dal.GetAttendanceByUserID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(userAttendance) == 0 {
		c.JSON(http.StatusNotFound, "Attendance records not found")
		return
	}

	c.JSON(http.StatusOK, userAttendance)
}

func UpdateAttendanceByUserID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userId"})
		return
	}

	var updatedRecord models.Attendance
	if err := c.ShouldBindJSON(&updatedRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = dal.UpdateAttendanceByUserID(userId, updatedRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Attendance record updated")
}

func DeleteAttendanceByUserID(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userId"})
		return
	}

	err = dal.DeleteAttendanceByUserID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, "Attendance record deleted")
}

func GetAttendanceByLessonID(c *gin.Context) {
	lessonId, err := strconv.Atoi(c.Param("lessonId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lessonId"})
		return
	}

	lessonAttendance, err := dal.GetAttendanceByLessonID(lessonId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(lessonAttendance) == 0 {
		c.JSON(http.StatusNotFound, "Attendance records not found")
		return
	}

	c.JSON(http.StatusOK, lessonAttendance)
}

// Handler to retrieve attendance records by course ID
func GetAttendanceByCourseID(c *gin.Context) {
	courseId := c.Param("courseId")

	courseAttendance, err := dal.GetAttendanceByCourseID(courseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(courseAttendance) == 0 {
		c.JSON(http.StatusNotFound, "Attendance records not found")
		return
	}

	c.JSON(http.StatusOK, courseAttendance)
}
