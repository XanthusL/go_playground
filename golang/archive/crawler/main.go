package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"
)

var (
	client *http.Client = &http.Client{}
	pipe   chan string

	urlRepo    map[string]struct{}
	urlRepoLck sync.RWMutex

	wg sync.WaitGroup
)

func main() {
	rootUrl := flag.String("url", "https://github.com/trending", "The root URL to fetch")
	depth := flag.Int("d", 2, "The depth")
	flag.Parse()
	urlRepo = make(map[string]struct{})
	pipe = make(chan string, 20)
	Crawl(*rootUrl, *depth, &RealFetcher{})
	for s := range pipe {
		go download(s)

	}
	wg.Wait()

	fmt.Println("Press ENTER to exit")
	fmt.Scanln(new(string))
}

type Fetcher interface {
	Fetch(url string) (urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	wg.Add(1)
	defer wg.Done()
	if depth <= 0 {
		return
	}
	urls, err := fetcher.Fetch(url)
	if err != nil {
		glog.Errorln("Failed to fetch", url)
		return
	}
	glog.Infoln("Fetch", url, "[done]")
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
	return
}

type RealFetcher struct {
}

func (f *RealFetcher) Fetch(url string) ([]string, error) {
	glog.Infoln(time.Now().UTC(), "[Fetching] start with", url)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	page := string(data)
	links := make([]string, 0)
	parts := strings.Split(page, "\"")
	urlRepoLck.Lock()
	defer urlRepoLck.Unlock()
	for _, v := range parts {
		_, ok := urlRepo[v]
		if ok {
			continue
		}
		if strings.HasPrefix(v, "http") {
			urlRepo[v] = struct{}{}
			frags := strings.Split(v, ".")
			l := len(frags)
			if l > 1 {
				switch frags[l-1] {
				case "mp4", "mp3", "jpg", "jpeg",
					"MP4", "MP3", "JPG", "JPEG":
					pipe <- v
				default:
					links = append(links, v)
				}
			} else {
				links = append(links, v)
			}

			//switch true {
			//case strings.HasSuffix(v, ".mp4"),
			//	strings.HasSuffix(v, ".MP4"),
			//	strings.HasSuffix(v, ".mp3"),
			//	strings.HasSuffix(v, ".MP3"),
			//	strings.HasSuffix(v, ".jpg"),
			//	strings.HasSuffix(v, ".JPG"),
			//	strings.HasSuffix(v, ".jpeg"),
			//	strings.HasSuffix(v, ".jpeg"):
			//	pipe <- v
			//default:
			//	links = append(links, v)
			//}
		}
	}
	return links, nil
}

func download(u string) {
	wg.Add(1)
	defer wg.Done()
	resp, err := client.Get(u)
	if err != nil {
		glog.Errorln("Failed to download", u, err)
		return
	}
	defer resp.Body.Close()

	r := bufio.NewReader(resp.Body)
	ss := strings.Split(u, "/")
	name := ss[len(ss)-1]
	f, err := os.Create(name)
	if err != nil {
		glog.Errorln("Failed to download", u, err)
		return
	}
	defer f.Close()

	_, err = r.WriteTo(f)
	if err != nil {
		glog.Errorln("Failed to download", u, err)
		return
	}
	glog.Infoln(time.Now().UTC(), name, "[saved]")
}
