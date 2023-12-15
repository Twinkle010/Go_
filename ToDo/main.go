package main 
import (
	_ "fmt"
	"ToDo/Config"
	_ "ToDo/Controllers"
	_ "ToDo/Models"
	"ToDo/Routes"
	"database/sql"
	_ "github.com/lib/pq"
)

// cnf := Config.GetConfig()
// db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
func main() {
	// Config.DB, err = gorm.Open(Config.DBUrl(Config.BuildDBConfig()))
	db, err := sql.Open("postgres", Config.DBUrl(Config.BuildDBConfig()))

	defer db.Close()
	if err != nil {
		panic(err)
	}
	r := Routes.SetUpRouter()
	r.Run()
}