
package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}
//var seen = make(map[string]bool)
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	seen := make(map[string]bool)
	ch := make(chan []string)
	var n int
	seen[url] = true
	n++
	go Extract(url, depth, fetcher, ch)
	for ;n >0; n-- {
		list := <- ch
	 	for _, u := range list {
			if !seen[u] {
				seen[u] = true
				n++
				go Extract(u, depth-1, fetcher, ch)
			}
		}
	}
	return
}

func Extract(u string, depth int, fetcher Fetcher, ch chan []string){
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(u)
	if err != nil {
		fmt.Println(err)
		//return ----> why did it cause deadlock ??
	} else {
	fmt.Printf("found: %s %q\n", u, body)
	}

	ch <-urls
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
