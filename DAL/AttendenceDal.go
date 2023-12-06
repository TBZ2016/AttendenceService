package dal

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	models "example/attendance/Models"
)

var (
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
)

// Initialize initializes the MongoDB client and connects to the database.
func Initialize(connectionString, dbName, collectionName string) error {
	clientOptions := options.Client().ApplyURI(connectionString)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return err
	}

	database = client.Database(dbName)
	collection = database.Collection(collectionName)
	return nil
}

// Close closes the MongoDB client when it's no longer needed.
func Close() {
	if client != nil {
		client.Disconnect(context.Background())
	}
}

// InsertAttendance inserts a new attendance record into the database.
func InsertAttendance(attendance models.Attendance) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, attendance)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAttendance updates an existing attendance record in the database.
func UpdateAttendance(attendance models.Attendance) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": attendance.AttendanceID}
	update := bson.M{"$set": attendance}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAttendanceByID deletes an attendance record by its ID.
func DeleteAttendanceByID(attendanceID int) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": attendanceID}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAttendanceByUserID retrieves attendance records by user ID.
func GetAttendanceByUserID(userID int) ([]models.Attendance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"userId": userID}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var attendanceList []models.Attendance
	for cursor.Next(ctx) {
		var attendance models.Attendance
		if err := cursor.Decode(&attendance); err != nil {
			return nil, err
		}
		attendanceList = append(attendanceList, attendance)
	}

	return attendanceList, nil
}
