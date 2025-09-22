package main

import (
	"database/sql"

	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	id       int
	content  string
	author   string
	comments []Comment
}

type Comment struct {
	id      int
	content string
	author  string
	post    *Post
}

var Db *sql.DB

func init() {

	var err error

	Db, err = sql.Open("postgres", "user=ak dbname=users_db password=4455@mint sslmode=disable")

	if err != nil {
		panic(err)
	}
}

func (c *Comment) Create() (err error) {
	if c.post == nil {
		err = errors.New("Post not found")
		return
	}

	err = Db.QueryRow("insert into comments (content, author, post_id) values ($1,$2,$3) returning id", c.content, c.author, c.post.id).Scan(&c.id)
	return
}

func Getpost(id int) (post Post, err error) {
	post = Post{}
	post.comments = []Comment{}
	err = Db.QueryRow("select id,content,author from posts where id=$1", id).Scan(&post.id, &post.content, &post.author)

	rows, err := Db.Query("select id, content,author from comments")
	if err != nil {
		return
	}

	for rows.Next() {
		comment := Comment{post: &post}
		err = rows.Scan(&comment.id, &comment.content, &comment.author)
		if err != nil {
			return
		}
		post.comments = append(post.comments, comment)
	}
	rows.Close()

	return
}

func (p *Post) Create() (err error) {

	err = Db.QueryRow("insert into posts (content, author) values($1, $2) returning id", p.content, p.author).Scan(&p.id)
	return
}

func main() {
	post := Post{content: "new post", author: "ak"}
	post.Create()

	cmm := Comment{content: "niceeeeee", author: "joes", post: &post}
	cmderr := cmm.Create()
	fmt.Println("cmd err:", cmderr)

	re, _ := Getpost(post.id)

	fmt.Println(re)
	fmt.Println("cmds:", re.comments)

}
