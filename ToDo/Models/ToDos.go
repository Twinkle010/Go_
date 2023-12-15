package Models


// write DB queries
import (
	"fmt"
	"ToDo/Config"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/georgysavva/scany/sqlscan"
	_ "gorm.io/gorm"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"database/sql"
)


func CreateATodo(todo *ToDo) (err error){
	db, err := sql.Open("postgres", Config.DBUrl(Config.BuildDBConfig()))
	query := fmt.Sprintf("insert into ToDo(id, title, description) values($1,$2,$3)")
	_, err = db.Query(query, todo.ID, todo.Title, todo.Description)
	if err != nil {
		return err
	}
	return nil
}

func GetAllTodoObjects(todo *[]ToDo) (todos []ToDo, err error) {
	db, err := sql.Open("postgres", Config.DBUrl(Config.BuildDBConfig()))
	query := fmt.Sprintf("select id, title, description from ToDo;")
	rows, err := db.Query(query)
	for rows.Next() {
		var todo ToDo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
            return todos, err
        }
        todos = append(todos, todo)
	}
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func GetATodo(todo *ToDo, id string) (todos []ToDo, err error) {
	db, err := sql.Open("postgres", Config.DBUrl(Config.BuildDBConfig()))
	query := fmt.Sprintf("select id, title, description from ToDo where id=$1;")
	rows, err := db.Query(query, id)
	for rows.Next() {
		var todo ToDo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
            return todos, err
        }
        todos = append(todos, todo)
	}
	fmt.Println(todo)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func UpdateTodo(todo *ToDo, id string) (rows *sql.Rows, err error) {
	db, err := sql.Open("postgres", Config.DBUrl(Config.BuildDBConfig()))
	query := fmt.Sprintf("update ToDo set title=$2, description=$3 where id=$1;")
	rows, err = db.Query(query, todo.ID, todo.Title, todo.Description)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func DeleteTodo(todo *ToDo, id string) (flag string, err error){
	db, err := sql.Open("postgres", Config.DBUrl(Config.BuildDBConfig()))
	query := fmt.Sprintf("delete from ToDo where id=$1;")
	_, err = db.Query(query, id)
	flag = "True"
	if err != nil {
		flag = "False"
		return flag, err
	}
	return flag, nil
}

