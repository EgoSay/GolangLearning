package main

import (
	"Learning/interface/mock"
	"Learning/interface/real"
	"fmt"
	"time"
)

const url = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}

func Download(r Retriever) string {
	return r.Get(url)
}
func Post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "cjw",
		"course": "Go",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func Session(session RetrieverPoster) string {
	session.Post(url, map[string]string{
		"contents": "www.cjwdream.top",
	})
	return session.Get(url)
}
func main() {
	var r Retriever
	mockRetriever := mock.Retriever{Contents: "this is me"}
	r = mockRetriever
	inspect(r)

	r = real.Retriever{
		UserAgent: "Mozelle/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("r is not a mock Retriever")
	}
	fmt.Println(Session(mockRetriever))
}
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > Type:%T Value:%v\n", r, r)
	fmt.Print(" > Type switch: ")
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
	fmt.Println()
}
