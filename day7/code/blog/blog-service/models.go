package blogservice

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Blog struct {
	db *gorm.DB
}

type Post struct {
	gorm.Model
	Title    string
	Text     string
	Comments []Comment
}

type Comment struct {
	gorm.Model
	Author string
	Text   string
}

func NewModelBlog(dbfile string) *Blog {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to open database")
	}
	db.AutoMigrate(&Post{})
	return &Blog{db: db}
}

func (b *Blog) Close() {
	b.db.Close()
}

func (b *Blog) CreatePost(title, text string) {
	b.db.Create(&Post{Title: title, Text: text})
}

func (b *Blog) GetPost(id int) Post {
	var post Post
	b.db.First(&post, id)
	return post
}

func (b *Blog) GetAllPosts() []Post {
	var posts []Post
	b.db.Find(&posts)
	log.Printf("Found %d posts", len(posts))
	return posts
}

func (b *Blog) GetCommentsByPostId(postId int) []Comment {
	var comments []Comment
	var post Post
	post.ID = uint(postId)
	b.db.Model(&post).Related(&comments)
	return comments
}
