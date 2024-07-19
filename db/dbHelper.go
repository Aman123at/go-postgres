package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Aman123at/go-postgres/model"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getPostgresConnString() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Unable to load env file")
		log.Fatal(err.Error())
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return connStr
}

func InitiateDbConnection() *sql.DB {
	conn, connerr := sql.Open("postgres", getPostgresConnString())

	if connerr != nil {
		log.Println("Unable to connect postgres")
		log.Fatal(connerr.Error())
	}

	pingerr := conn.Ping()

	if pingerr != nil {
		log.Println("Unable to connect postgres")
		log.Fatal(pingerr.Error())
	}

	log.Println("Successfully connected to Database")

	return conn
}

func InsertNewPost(post *model.Post) int {
	dbInstance := InitiateDbConnection()

	defer dbInstance.Close()

	sqlQuery := `INSERT INTO posts (title, content, author) VALUES ($1, $2, $3) RETURNING postid`

	var postId int

	execErr := dbInstance.QueryRow(sqlQuery, post.Title, post.Content, post.Author).Scan(&postId)

	if execErr != nil {
		log.Printf("ERROR: Unable to execute query")
		log.Fatal(execErr.Error())
	}

	log.Printf("Successfully inserted into posts table %v", postId)

	return postId
}

func GetAllPostsFromDB() []model.Post {
	var posts []model.Post

	dbInstance := InitiateDbConnection()

	defer dbInstance.Close()

	sqlQuery := `SELECT * FROM posts`

	rows, err := dbInstance.Query(sqlQuery)

	if err != nil {
		log.Printf("ERROR: Unable to execute query")
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var post model.Post
		err = rows.Scan(&post.PostId, &post.Title, &post.Content, &post.Author)
		if err != nil {
			log.Fatalf("Error : Not able to scan rows %v", err)
		}
		posts = append(posts, post)
	}

	return posts
}

func DeleteOnePost(postId int) {
	dbInstance := InitiateDbConnection()

	defer dbInstance.Close()

	queryStr := `DELETE FROM posts WHERE postid=$1`

	result, err := dbInstance.Exec(queryStr, postId)

	if err != nil {
		log.Fatalf("Error: Unable to execute delete query %v", err.Error())
		return
	}

	rowsAffected, rowsErr := result.RowsAffected()

	if rowsErr != nil {
		log.Fatalf("Something went wrong in rowsEffected %v", rowsErr.Error())
		return
	}

	fmt.Printf("%v Rows Affected", rowsAffected)
}

func DeleteAllPosts() {
	dbInstance := InitiateDbConnection()

	defer dbInstance.Close()

	queryStr := `DELETE FROM posts`

	result, err := dbInstance.Exec(queryStr)

	if err != nil {
		log.Fatalf("Error: Unable to execute delete query %v", err.Error())
		return
	}

	rowsAffected, rowsErr := result.RowsAffected()

	if rowsErr != nil {
		log.Fatalf("Something went wrong in rowsEffected %v", rowsErr.Error())
		return
	}

	fmt.Printf("%v Rows Affected", rowsAffected)
}

func GetOnePost(postId int) model.Post {
	var post model.Post

	dbInstance := InitiateDbConnection()

	defer dbInstance.Close()

	queryStr := `SELECT * FROM posts WHERE postid=$1`

	row := dbInstance.QueryRow(queryStr, postId)

	sqlErr := row.Scan(&post.PostId, &post.Title, &post.Content, &post.Author)

	if sqlErr == sql.ErrNoRows {
		return model.Post{}
	}

	if sqlErr != nil {
		log.Fatalf("Error: Unable to scal row data %v", sqlErr)
	}

	return post
}

func UpdatePost(postId int, postData model.Post) {
	dbInstance := InitiateDbConnection()

	defer dbInstance.Close()

	queryStr := `UPDATE posts SET title=$2, content=$3, author=$4 WHERE postid=$1`

	res, execErr := dbInstance.Exec(queryStr, postId, postData.Title, postData.Content, postData.Author)

	if execErr != nil {
		log.Fatalf("Error: Unable to execute update query %v", execErr.Error())
	}

	rowsAffected, rowerr := res.RowsAffected()

	if rowerr != nil {
		log.Fatalf("Error: Unable to find Affected Rows %v", rowerr.Error())
	}

	log.Printf("Total rows affected after updated is : %v", rowsAffected)
}
