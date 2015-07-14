//usr/bin/env go run $0 "$@"; exit

package main

import (
    "encoding/xml"
    "fmt"
    "html/template"
    "io/ioutil"
    "net/http"
    "strings"
    "time"
)

const feed_url = "http://thedoomthatcametopuppet.tumblr.com/rss"

func main() {
    // taken from
    // https://siongui.github.io/2015/03/03/go-parse-web-feed-rss-atom/
    type Item struct {
        // Required
        Title       string      `xml:"title"`
        Link        string      `xml:"link"`
        Description template.HTML   `xml:"description"`
        // Optional
        Content     template.HTML   `xml:"encoded"`
        PubDate     string      `xml:"pubDate"`
        Comments    string      `xml:"comments"`
    }
    type Rss2 struct {
        XMLName     xml.Name    `xml:"rss"`
        Version     string      `xml:"version,attr"`
        // Required
        Title       string      `xml:"channel>title"`
        Link        string      `xml:"channel>link"`
        Description string      `xml:"channel>description"`
        // Optional
        PubDate     string      `xml:"channel>pubDate"`
        ItemList    []Item      `xml:"channel>item"`
    }
    feed := Rss2{}

    resp, err := http.Get(feed_url)
    if err != nil {
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return
    }
    err = xml.Unmarshal(body, &feed)
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    //fmt.Printf("XMLName: %#v\n", feed.XMLName)
    fmt.Println(feed.Title)
    fmt.Println(strings.Repeat("=", len(feed.Title)))
    fmt.Println(feed.Description, "\n")
    time.Sleep(time.Duration(len(feed.Description)) * 50 * time.Millisecond)

    for _, item := range feed.ItemList{
        fmt.Println(item.PubDate)
        fmt.Println(strings.Repeat("-", len(item.PubDate)))
        fmt.Println(item.Description, "\n")
        time.Sleep(time.Duration(len(item.Description)) * 25 * time.Millisecond)
    }
}
