package database

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	Description string
	StudentID   int64
	StartDate   time.Time
	EndDate     time.Time
	State       string // pending, done, canceled

	Student *Student `json:"student" gorm:"foreignKey:StudentID"`
}

func CreatePlan(studentID string, startTime time.Time, endTime time.Time, desc string) (*Plan, error) {
	student, err := GetStudent(studentID)
	if err != nil {
		return nil, err
	}

	// check for overlapping plans
	var count int64
	if err := DB.Model(&Plan{}).
		Where("student_id = ? AND ((start_date BETWEEN ? AND ?) OR (end_date BETWEEN ? AND ?))", studentID, startTime, endTime, startTime, endTime).
		Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, fmt.Errorf("overlapping plans")
	}

	plan := &Plan{
		Description: desc,
		StartDate:   startTime,
		EndDate:     endTime,
		State:       "pending",
	}

	if err := DB.Model(student).Association("Plans").Append(plan); err != nil {
		return nil, err
	}

	return plan, nil
}

func UpdatePlan(studentID string, planID string, desc string, startTime time.Time, endTime time.Time, state string) (*Plan, error) {
	// find student's plan
	var plan Plan
	if err := DB.Where("student_id = ? AND id = ?", studentID, planID).First(&plan).Error; err != nil {
		return nil, err
	}

	// check for overlapping plans
	var count int64
	if err := DB.Model(&Plan{}).
		Where("student_id = ? AND id != ? AND ((start_date BETWEEN ? AND ?) OR (end_date BETWEEN ? AND ?))", studentID, planID, startTime, endTime, startTime, endTime).
		Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, fmt.Errorf("overlapping plans")
	}

	plan.Description = desc
	plan.StartDate = startTime
	plan.EndDate = endTime
	plan.State = state

	if err := DB.Save(&plan).Error; err != nil {
		return nil, err
	}

	return &plan, nil

}

// DeletePlan deletes a plan
func DeletePlan(studentID string, planID string) error {
	if err := DB.Where("student_id = ? AND id = ?", studentID, planID).Delete(&Plan{}).Error; err != nil {
		return err
	}

	return nil
}

// GetPlans returns all plans of a student
func GetPlans(studentID string) ([]Plan, error) {
	var plans []Plan
	if err := DB.Where("student_id = ?", studentID).Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}

// GetPlan returns a plan
func GetPlan(studentID string, planID string) (*Plan, error) {
	var plan Plan
	if err := DB.Where("student_id = ? AND id = ?", studentID, planID).First(&plan).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}
