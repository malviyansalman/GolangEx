package main

import (
	"concurrency/util"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/gophercises/quiet_hn/hn"
)

func main() {
	// parse flags
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "the number of top stories to display")
	flag.Parse()
	fmt.Println(port, numStories)
	tpl := template.Must(template.ParseFiles("./index.html"))
	http.HandleFunc("/", handler(numStories, tpl))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("Unable to start Server", err)
	}
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	sc := storyCache{
		numStories: numStories,
		duration:   6 * time.Second,
	}

	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for {
			tempCache := storyCache{
				numStories: numStories,
				duration:   6 * time.Second,
			}
			stories, err := tempCache.stories()
			if err == nil {
				sc.mutex.Lock()
				sc.cache = stories
				sc.expiration = tempCache.expiration
				sc.mutex.Unlock()
			}
			<-ticker.C
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		stories, err := sc.stories()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data := util.TemplateData{
			Stories: stories,
			Time:    time.Now().Sub(start),
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
	})
}

var (
	cache          []util.Item
	cacheEpiration time.Time
)

type storyCache struct {
	cache      []util.Item
	duration   time.Duration
	expiration time.Time
	mutex      sync.Mutex
	numStories int
}

func (sc *storyCache) stories() ([]util.Item, error) {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	if time.Now().Sub(sc.expiration) < 0 {
		return sc.cache, nil
	}
	stories, err := getTopStories(sc.numStories)
	if err != nil {
		return nil, err
	}
	sc.expiration = time.Now().Add(sc.duration)
	sc.cache = stories
	return sc.cache, nil
}

func getCachedStories(numStories int) ([]util.Item, error) {
	if time.Now().Sub(cacheEpiration) < 0 {
		return cache, nil
	}
	stories, err := getTopStories(numStories)
	if err != nil {
		return nil, err
	}
	cacheEpiration = time.Now().Add(15 * time.Second)
	cache = stories
	return cache, nil
}
func getTopStories(numStories int) ([]util.Item, error) {
	var client hn.Client
	ids, err := client.TopItems()
	if err != nil {
		return nil, errors.New("failed to Get top stories")
	}
	var stories []util.Item
	from := 0
	for len(stories) < numStories {
		//fmt.Println("loop")
		need := (numStories - len(stories)) * 5 / 4
		res := getStories(ids[from : from+need])
		from += need
		stories = append(stories, res...)
	}
	return stories[:numStories], nil
}

func getStories(ids []int) []util.Item {
	resultChan := make(chan util.Result)
	for i := 0; i < len(ids); i++ {
		go func(idx, id int) {
			var client hn.Client
			hnItem, err := client.GetItem(id)
			if err != nil {
				resultChan <- util.Result{Idx: idx, Err: err}
			} else {
				resultChan <- util.Result{Idx: idx, Item: util.ParseHNItem(hnItem)}
			}
		}(i, ids[i])
	}
	var results []util.Result
	var stories []util.Item
	for i := 0; i < len(ids); i++ {
		res := <-resultChan
		if res.Err != nil {
			continue
		}
		results = append(results, res)
	}
	sort.Slice(stories, func(i, j int) bool {
		return results[i].Idx < results[j].Idx
	})

	for _, res := range results {
		if util.IsStoryLink(res.Item) {
			stories = append(stories, res.Item)
		}
	}
	return stories
}
