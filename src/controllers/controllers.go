package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"mysql_conn"
)

// 변수 명은 반드시 대문자로 시작해야만 json으로 export 된다.
type mydata struct{
	Id string `json:"id"`
	Name string `json:"name"`
}

func GetSingleRow(c echo.Context) error{
	id := c.QueryParam("id")
	db := mysql_conn.CreateMysqlConn()
	defer db.Close()
	var name string

	err := db.QueryRow("select name from mytable where id = ?", id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, name)
}

func GetMultipleRow(c echo.Context) error{
	result := []mydata{}
	db := mysql_conn.CreateMysqlConn()
	defer db.Close()

	row, err := db.Query("select * from mytable")
	if err != nil{
		log.Fatal(err)
	}

	result_object := mydata{}
	for row.Next(){
		err := row.Scan(&result_object.Id, &result_object.Name)
		if err != nil{
			log.Fatal(err)
		}
		result = append(result, result_object)
	}

	return c.JSON(http.StatusOK, result)
}

func InsertRow(c echo.Context) error{
	name := c.QueryParam("name")
	id := c.QueryParam("id")
	db := mysql_conn.CreateMysqlConn()
	defer db.Close()

	result, err := db.Exec("insert into mytable values(?, ?)", name, id)
	if err != nil {
		log.Fatal(err)
	}
	
	n, err := result.RowsAffected()
	if n == 1{
		return c.JSON(http.StatusOK, true)
	}else{
		return c.JSON(http.StatusOK, false)
	}
}