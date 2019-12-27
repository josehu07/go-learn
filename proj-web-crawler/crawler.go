package main


import (
	"fmt"
	"sync"
)


// A set (map of bool) protected by a mutex.
type SafeSet struct {
	has map[string]bool
	mux sync.Mutex
}


// `Crawl` uses fetcher to recursively crawl pages starting with url, to
// a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, tried *SafeSet,
	       wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Maximum depth reached.
	if depth <= 0 {
		return
	}
	
	// Check if this page is tried before.
	tried.mux.Lock()
	if tried.has[url] {
		defer tried.mux.Unlock()
		return
	}
	tried.has[url] = true
	tried.mux.Unlock()
	
	// If not, fetch this page and crawl deeper from it.
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, tried, wg)
	}
	wg.Add(len(urls))
}


func main() {
	tried := SafeSet{has: make(map[string]bool)}
	var wg sync.WaitGroup
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, &tried, &wg)
	wg.Wait()
}


// The `fakeFetcher` is Fetcher that returns canned results.
type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLs found on it.
	Fetch(url string) (body string, urls []string, err error)
}

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


// This `fetcher` is a populated fakeFetcher. It simulates a subset of
// Golang's official website.
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
