package main

import (
	"log"

	"github.com/Fbulkaya/Student_scheduler/api"
	"github.com/Fbulkaya/Student_scheduler/database"

	"github.com/labstack/echo/v4"
)

func main() {
	if err := database.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Debug = true
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Student Scheduler API")
	})
	// Students
	e.GET("/students", api.GetStudents)
	e.GET("/students/:id", api.GetStudent)
	e.POST("/students", api.CreateStudent)
	e.PUT("/students/:id", api.UpdateStudent)
	e.DELETE("/students/:id", api.DeleteStudent)

	// Plans
	e.GET("/students/:student_id/plans", api.GetStudentPlans)
	e.GET("/students/:student_id/plans/:plan_id", api.GetStudentPlan)
	e.POST("/students/:student_id/plans", api.CreatePlan)
	e.PUT("/students/:student_id/plans/:plan_id", api.UpdatePlan)
	e.DELETE("/students/:student_id/plans/:plan_id", api.DeletePlan)

	e.Logger.Fatal(e.Start(":8080"))
}
