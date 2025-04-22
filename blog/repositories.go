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
	res := testDb.db.Where("title = ?", blog.Title).Find(&existingBlog)
	if res.RowsAffected > 0 {
		return existingBlog, fmt.Errorf("blog with title %s already exists", blog.Title)
	}
	res = testDb.db.Create(&blog)
	if res.Error != nil {
		return Blog{}, res.Error
	}
	return blog, nil
}

func (testDb *TestDb) UpdateBlog(blogId int, blog Blog) (Blog, error) {
	var existingBlog Blog
	res := testDb.db.First(&existingBlog, blogId)
	if res.Error != nil {
		return Blog{}, res.Error
	}
	if res.RowsAffected == 0 {
		return Blog{}, fmt.Errorf("blog id %d doesn't exist", blogId)
	}
	title := blog.Title
	content := blog.Content
	if blog.Title == "" {
		title = existingBlog.Title
	}
	if blog.Content == "" {
		content = existingBlog.Content
	}
	updates := map[string]interface{}{
		"title":   title,
		"content": content,
	}
	res = testDb.db.Model(&Blog{}).Where("id = ?", blogId).Updates(updates)
	if res.Error != nil {
		return Blog{}, res.Error
	}
	var updatedBlog Blog
	err := testDb.db.First(&updatedBlog, blogId).Error
	if err != nil {
		return Blog{}, err
	}
	return updatedBlog, nil
}

func (testDb *TestDb) GetABlog(blogId int) (Blog, error) {
	var blog Blog
	res := testDb.db.Find(&blog, blogId)
	if res.Error != nil {
		return Blog{}, res.Error
	}
	if res.RowsAffected == 0 {
		return Blog{}, fmt.Errorf("blog id %d doesn't exist", blogId)
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
	res := testDb.db.Find(&blog, blogId)
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
		return Author{}, res.Error
	}
	return author, nil
}

func (testDb *TestDb) UpdateAuthor(authorId int, author Author) (Author, error) {
	var existingAuthor Author
	res := testDb.db.First(&existingAuthor, authorId)
	if res.Error != nil {
		return Author{}, res.Error
	}
	if res.RowsAffected == 0 {
		return Author{}, fmt.Errorf("author id %d doesn't exist", authorId)
	}
	name := author.Name
	if author.Name == "" {
		name = existingAuthor.Name
	}
	updates := map[string]interface{}{
		"name": name,
	}
	res = testDb.db.Model(&Author{}).Where("id = ?", authorId).Updates(updates)
	if res.Error != nil {
		return Author{}, res.Error
	}
	var updatedAuthor Author
	err := testDb.db.First(&updatedAuthor, authorId).Error
	if err != nil {
		return Author{}, err
	}

	return updatedAuthor, nil
}

func (testDb *TestDb) GetAuthor(authorId int) (Author, error) {
	var author Author
	res := testDb.db.First(&author, authorId)
	if res.Error != nil {
		return Author{}, res.Error
	}
	if res.RowsAffected == 0 {
		return Author{}, fmt.Errorf("author id %d doesn't exist", authorId)
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
