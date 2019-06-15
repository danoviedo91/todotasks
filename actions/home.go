package actions

import (
	"log"
	"time"

	"github.com/danoviedo91/todotasks/models"
	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up a home page.
func HomeHandler(c buffalo.Context) error {

	// Establish database connection
	db := models.DB

	// queryValues gets all GET parameters sent with the URL
	// queryValues := r.URL.Query()

	// Initialize database-query variables

	allRecords := []models.Todo{}
	err := db.All(&allRecords)

	if err != nil {
		log.Fatal(err)
	}

	records := allRecords

	filterStatus := struct {
		Incompleted bool
		Completed   bool
	}{
		false,
		false,
	}

	// If /?completed=true

	// if queryValues.Get("status") == "completed" {
	// 	for _, row := range allRecords {
	// 		if row.Completed == true {
	// 			records = append(records, row)
	// 		}
	// 	}
	// 	incompletedTasks = len(allRecords) - len(records)
	// 	filterStatus.Completed = true
	// 	// If /?completed=false
	// } else if queryValues.Get("status") == "incompleted" {
	// 	for _, row := range allRecords {
	// 		if row.Completed == false {
	// 			records = append(records, row)
	// 		}
	// 	}
	// 	incompletedTasks = len(records)
	// 	filterStatus.Incompleted = true
	// 	// If /
	// } else {
	for _, row := range allRecords {
		if row.Completed == false {
			records = append(records, row)
		}
	}
	pendingTasksNumber := len(records)
	records = allRecords
	//}

	// Catch if there are tasks to show or not...

	defaultMsgFlag := false

	if len(records) == 0 {
		defaultMsgFlag = true
	}

	// Prepare to send data to template

	c.Set("defaultMsgFlag", defaultMsgFlag)
	c.Set("pendingTasksNumber", pendingTasksNumber)
	c.Set("currentDateTime", time.Now())
	c.Set("tasksArray", records)
	c.Set("filterStatus", filterStatus)

	// templateData := struct {
	// 	PendingTasksNumber int
	// 	CurrentDateTime    time.Time
	// 	TasksArray         []models.Todo
	// 	TaskStruct         models.Todo
	// 	DefaultMsgFlag     bool
	// 	FilterStatus       struct {
	// 		Incompleted bool
	// 		Completed   bool
	// 	}
	// }{
	// 	DefaultMsgFlag:     defaultMsgFlag,
	// 	PendingTasksNumber: incompletedTasks,
	// 	CurrentDateTime:    time.Now(),
	// 	TasksArray:         records,
	// 	FilterStatus:       filterStatus,
	// }

	return c.Render(200, r.HTML("index.html"))
}
