package main

import (
	"flag"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	dbName := flag.String("dbname", "test", "database name")
	flag.Parse()
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file: ", err)
		return
	}
	if *dbName == "" {
		fmt.Println("Database name not provided")
		return
	}
	testDb := NewDb(*dbName)
	err = testDb.db.AutoMigrate(&Author{})
	if err != nil {
		fmt.Println("Error creating table: ", err)
		return
	}
	err = testDb.db.AutoMigrate(&Blog{})
	if err != nil {
		fmt.Println("Error creating table: ", err)
		return
	}

	// john, err := testDb.CreateAuthor(Author{Name: "John Doe"})
	// if err != nil {
	// 	fmt.Println("Error creating author: ", err)
	// 	return
	// }
	// fmt.Println("Created author: ", john)

	// alice, err := testDb.CreateAuthor(Author{Name: "Alice Smith"})
	// if err != nil {
	// 	fmt.Println("Error creating author: ", err)
	// 	return
	// }
	// fmt.Println("Created author: ", alice)

	// blog1, err := testDb.CreateBlog(Blog{Title: "John's Blog", Content: "This is a blog written by john", AuthorId: int(john.ID)})
	// if err != nil {
	// 	fmt.Println("Error creating blog: ", err)
	// 	return
	// }
	// fmt.Println("Created blog: ", blog1)

	// blog2, err := testDb.CreateBlog(Blog{Title: "Alice's Blog", Content: "This is a blog written by alice", AuthorId: int(alice.ID)})
	// if err != nil {
	// 	fmt.Println("Error creating blog: ", err)
	// 	return
	// }
	// fmt.Println("Created blog: ", blog2)

	// err = testDb.DeleteBlog(1)
	// if err != nil {
	// 	fmt.Println("Error deleting blog: ", err)
	// 	return
	// }

	blogs, err := testDb.GetBlogs()
	if err != nil {
		fmt.Println("Error getting blogs: ", err)
		return
	}
	for _, blog := range blogs {
		fmt.Println(blog.ID, blog.Title, blog.Content, blog.Author.Name)
	}

	blog, err := testDb.UpdateBlog(int(2), Blog{Title: "Alice Smith's Blog"})
	if err != nil {
		fmt.Println("Error updating blog: ", err)
		return
	}
	fmt.Println("Updated blog: ", blog)

	blogs, err = testDb.GetBlogs()
	if err != nil {
		fmt.Println("Error getting blogs: ", err)
		return
	}
	for _, blog := range blogs {
		fmt.Println(blog.ID, blog.Title, blog.Content, blog.Author.Name)
	}

	// blog, err := testDb.GetABlog(int(1))
	// if err != nil {
	// 	fmt.Println("Error getting blog: ", err)
	// 	return
	// }
	// fmt.Println("Got blog: ", blog)

	blogs, err = testDb.SearchBlogs(SearchByAuthor, "Alice")
	if err != nil {
		fmt.Println("Error searching blogs: ", err)
	}
	fmt.Println("Searched blogs: ", blogs)
	blogs, err = testDb.SearchBlogs(SearchByTitle, "Blog")
	if err != nil {
		fmt.Println("Error searching blogs: ", err)
	}
	fmt.Println("Searched blogs: ", blogs)

	// err = testDb.db.Migrator().DropTable(&Blog{})
	// if err != nil {
	// 	fmt.Println("Error dropping table: ", err)
	// }
	// err = testDb.db.Migrator().DropTable(&Author{})
	// if err != nil {
	// 	fmt.Println("Error dropping table: ", err)
	// }
}
