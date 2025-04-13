package main

import "fmt"

const (
	SearchByTitle  = 0
	SearchByAuthor = 1
)

type Blogger interface {
	CreateBlog(blog Blog) (Blog, error)
	UpdateBlog(blogId int, blog Blog) (Blog, error)
	GetABlog(blogId int) (Blog, error)
	GetBlogs() ([]Blog, error)
	DeleteBlog(blogId int) error
	SearchBlogs(search string) ([]Blog, error)

	CreateAuthor(author Author) (Author, error)
	UpdateAuthor(authorId int, author Author) (Author, error)
	GetAuthor(author Author) (Author, error)
	DeleteAuthor(authorId int) error
}

func (testDb *TestDb) CreateBlog(blog Blog) (Blog, error) {
	var existingBlog Blog
	res := testDb.db.Where("title = ?", blog.Title).First(&existingBlog)
	if res.RowsAffected > 0 {
		return existingBlog, fmt.Errorf("blog with title %s already exists", blog.Title)
	}
	res = testDb.db.Create(&blog)
	if res.Error != nil {
		return blog, res.Error
	}
	return blog, nil
}

func (testDb *TestDb) UpdateBlog(blogId int, blog Blog) (Blog, error) {
	var existingBlog Blog
	res := testDb.db.First(&existingBlog, blogId)
	if res.Error != nil {
		return blog, res.Error
	}
	if res.RowsAffected == 0 {
		return existingBlog, fmt.Errorf("blog id %d doesn't exist", blogId)
	}
	if blog.Title != "" {
		existingBlog.Title = blog.Title
	}
	if blog.Content != "" {
		existingBlog.Content = blog.Content
	}
	if blog.AuthorId != 0 {
		var author Author
		res := testDb.db.First(&author, blog.AuthorId)
		if res.Error != nil {
			return blog, res.Error
		}
		existingBlog.AuthorId = blog.AuthorId
	}
	res = testDb.db.Save(&existingBlog)
	if res.Error != nil {
		return existingBlog, res.Error
	}
	return existingBlog, nil
}

func (testDb *TestDb) GetABlog(blogId int) (Blog, error) {
	var blog Blog
	res := testDb.db.First(&blog, blogId)
	if res.Error != nil {
		return blog, res.Error
	}
	if res.RowsAffected == 0 {
		return blog, fmt.Errorf("blog id doesn't exist", blogId)
	}
	return blog, nil
}

func (testDb *TestDb) GetBlogs() ([]Blog, error) {
	var blogs []Blog
	res := testDb.db.Find(&blogs)
	if res.Error != nil {
		return nil, res.Error
	}
	return blogs, nil
}

func (testDb *TestDb) DeleteBlog(blogId int) error {
	var blog Blog
	res := testDb.db.First(&blog, blogId)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("blog id %d doesn't exist", blogId)
	}
	res = testDb.db.Delete(&blog)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (testDb *TestDb) SearchBlogs(searchBy int, name string) ([]Blog, error) {
	var blogs []Blog
	query := testDb.db.Preload("Author")
	switch searchBy {
	case SearchByTitle:
		query = query.Where("title ILIKE ?", "%"+name+"%")
	case SearchByAuthor:
		query = query.Joins("JOIN authors ON authors.id = blogs.author_id").
			Where("authors.name ILIKE ?", "%"+name+"%")
	default:
		return nil, fmt.Errorf("invalid search criteria")
	}
	if res := query.Find(&blogs); res.Error != nil {
		return nil, res.Error
	}
	return blogs, nil
}

func (testDb *TestDb) CreateAuthor(author Author) (Author, error) {
	res := testDb.db.Create(&author)
	if res.Error != nil {
		return author, res.Error
	}
	return author, nil
}

func (testDb *TestDb) UpdateAuthor(authorId int, author Author) (Author, error) {
	var existingAuthor Author
	res := testDb.db.First(&existingAuthor, authorId)
	if res.Error != nil {
		return author, res.Error
	}
	if res.RowsAffected == 0 {
		return existingAuthor, fmt.Errorf("author id %d doesn't exist", authorId)
	}
	if author.Name != "" {
		existingAuthor.Name = author.Name
	}
	res = testDb.db.Save(&existingAuthor)
	if res.Error != nil {
		return existingAuthor, res.Error
	}
	return existingAuthor, nil
}

func (testDb *TestDb) GetAuthor(authorId int) (Author, error) {
	var author Author
	res := testDb.db.First(&author, authorId)
	if res.Error != nil {
		return author, res.Error
	}
	if res.RowsAffected == 0 {
		return author, fmt.Errorf("author id %d doesn't exist", authorId)
	}
	return author, nil
}

func (testDb *TestDb) DeleteAuthor(authorId int) error {
	var author Author
	res := testDb.db.First(&author, authorId)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("author id %d doesn't exist", authorId)
	}
	res = testDb.db.Delete(&author)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
