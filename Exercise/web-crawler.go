package main

import (
	"fmt"
	"sync"
)

var listOfURL urlList

func init() {
	listOfURL.site = make(map[string]string)
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}
type urlList struct {
	mu   sync.Mutex
	site map[string]string
}

func (u *urlList) isDone(url string) bool {
	_, ok := u.site[url]
	return ok
}
func (u *urlList) add(url string, fetcher Fetcher) ([]string, error) {
	if u.isDone(url) {
		return nil, nil
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		return nil, err
	}
	fmt.Printf("found: %s %q\n", url, body)
	u.mu.Lock()
	defer u.mu.Unlock()
	u.site[url] = body
	return urls, err
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	urls, err := listOfURL.add(url, fetcher)
	if err != nil {
		fmt.Println(err)
		return
	}
	if urls != nil {
		crawlDone := make(chan bool)
		for _, u := range urls {
			go func(u string) {
				Crawl(u, depth-1, fetcher)
				crawlDone <- true
			}(u)
		}
		for range urls {
			<-crawlDone
		}
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
