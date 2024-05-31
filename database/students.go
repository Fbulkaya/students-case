package database

import (
	"fmt"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name  string
	Email string

	Plans []Plan `gorm:"foreignKey:StudentID"`
}

func CreateStudent(name string, email string) (*Student, error) {
	student := &Student{
		Name:  name,
		Email: email,
	}

	if err := DB.Create(student).Error; err != nil {
		return nil, fmt.Errorf("student create: %w", err)
	}

	return student, nil
}

func GetStudent(id string) (*Student, error) {
	var student Student
	if err := DB.First(&student, id).Error; err != nil {
		return nil, fmt.Errorf("student get: %w", err)
	}
	return &student, nil
}

func GetStudents() ([]Student, error) {
	var students []Student
	if err := DB.Find(&students).Error; err != nil {
		return nil, fmt.Errorf("students get: %w", err)
	}
	return students, nil
}

func UpdateStudent(id string, name string, email string) (*Student, error) {
	student, err := GetStudent(id)
	if err != nil {
		return nil, fmt.Errorf("student update: %w", err)
	}

	student.Name = name
	student.Email = email

	if err := DB.Save(student).Error; err != nil {
		return nil, fmt.Errorf("student update: %w", err)
	}

	return student, nil
}

func DeleteStudent(id string) error {
	student, err := GetStudent(id)
	if err != nil {
		return fmt.Errorf("student delete: %w", err)
	}

	if err := DB.Delete(student).Error; err != nil {
		return fmt.Errorf("student delete: %w", err)
	}

	return nil
}
