package blogposts

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		fmt.Printf("This is the post: %v", post)
		if err != nil {
			return nil, err //todo: needs clarification, should
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)
	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}
	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagsSeparator), ", ")
	scanner.Scan() // ignore a line
	var b strings.Builder
	for scanner.Scan() {
		fmt.Fprintln(&b, scanner.Text())
	}
	body := strings.TrimSuffix(b.String(), "\n")
	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}, nil
}
