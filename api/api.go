package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Fbulkaya/Student_scheduler/database"
	"github.com/labstack/echo/v4"
)

type CreateStudentRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Create a student
func CreateStudent(c echo.Context) error {
	var req CreateStudentRequest
	if err := c.Bind(&req); err != nil {
		return fmt.Errorf("binding request payload: %w", err)
	}

	student, err := database.CreateStudent(req.Name, req.Email)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, student)
}

// Get a student
func GetStudent(c echo.Context) error {
	id := c.Param("id")

	student, err := database.GetStudent(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, student)
}

// Get all students
func GetStudents(c echo.Context) error {
	students, err := database.GetStudents()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, students)
}

type UpdateStudentRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Update a student
func UpdateStudent(c echo.Context) error {
	id := c.Param("id")

	var req UpdateStudentRequest
	if err := c.Bind(&req); err != nil {
		return fmt.Errorf("binding request payload: %w", err)
	}

	student, err := database.UpdateStudent(id, req.Name, req.Email)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, student)
}

// Delete a student
func DeleteStudent(c echo.Context) error {
	id := c.Param("id")

	if err := database.DeleteStudent(id); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

type CreatePlanRequest struct {
	StudentID string `param:"student_id"`
	Desc      string `json:"desc"`
	StartDate string `json:"start_date"` // YYYY-MM-DD HH:MM
	EndDate   string `json:"end_date"`   // YYYY-MM-DD HH:MM
}

// Create a plan
func CreatePlan(c echo.Context) error {
	var req CreatePlanRequest
	if err := c.Bind(&req); err != nil {
		return fmt.Errorf("binding request payload: %w", err)
	}

	startDate, err := parseTime(req.StartDate)
	if err != nil {
		return err
	}

	endDate, err := parseTime(req.EndDate)
	if err != nil {
		return err
	}

	plan, err := database.CreatePlan(req.StudentID, *startDate, *endDate, req.Desc)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, plan)
}

type UpdatePlanRequest struct {
	StudentID string `param:"student_id"`
	PlanID    string `param:"plan_id"`
	Desc      string `json:"desc"`
	State     string `json:"state"`      // pending, done, canceled
	StartDate string `json:"start_date"` // YYYY-MM-DD HH:MM
	EndDate   string `json:"end_date"`   // YYYY-MM-DD HH:MM
}

// Update a plan
func UpdatePlan(c echo.Context) error {
	var req UpdatePlanRequest
	if err := c.Bind(&req); err != nil {
		return fmt.Errorf("binding request payload: %w", err)
	}

	startDate, err := parseTime(req.StartDate)
	if err != nil {
		return err
	}

	endDate, err := parseTime(req.EndDate)
	if err != nil {
		return err
	}

	plan, err := database.UpdatePlan(req.StudentID, req.PlanID, req.Desc, *startDate, *endDate, req.State)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, plan)
}

// Delete a plan
func DeletePlan(c echo.Context) error {
	studentID := c.Param("student_id")
	planID := c.Param("plan_id")

	if err := database.DeletePlan(studentID, planID); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

// Get student's plans
func GetStudentPlans(c echo.Context) error {
	studentID := c.Param("student_id")

	plans, err := database.GetPlans(studentID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, plans)
}

// Get a student's plan
func GetStudentPlan(c echo.Context) error {
	studentID := c.Param("student_id")
	planID := c.Param("plan_id")

	plan, err := database.GetPlan(studentID, planID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, plan)
}

func parseTime(param string) (*time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04", param)
	if err != nil {
		return nil, fmt.Errorf("parsing time: %w", err)
	}
	return &t, nil
}
