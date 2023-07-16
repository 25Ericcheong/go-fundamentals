package files_test

import (
	"errors"
	"files"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts, err := files.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}

// typically if we were doing something with err - we would write a test case for it but since we are not and there is only one method. This test is trivial but doesn't hurt to be pragmatic either way
type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestFailingFromFS(t *testing.T) {
	_, err := files.NewPostsFromFS(StubFailingFS{})

	if err == nil {
		t.Errorf("expected error to occur but none was found")
	}
}
