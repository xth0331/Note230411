package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() { // 连接到数据库
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=cw2AEiTZBn8QtYP- host=192.168.1.243 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int) (post Post, err error) { // 获取指定的帖子
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) create() (err error) { //创建一篇帖子
	statement := "insert into posts(content,author) values ($1,$2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}
func (post *Post) update() (err error) { // 更新指定的帖子
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}
func (post *Post) delete() (err error) { // 删除指定帖子
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}
