package main

import (
    "context"
    "net/http"
    "io"
    "encoding/xml"
    "html"
)

type RSSFeed struct{
    Channel struct{
        Title string `xml:"title"`
        Link string `xml:"link"`
        Description string `xml:"description"`
        Item []RSSItem `xml:"item"`
    } `xml:"channel"`
}

type RSSItem struct{
    Title string `xml:"title"`
    Link string `xml:"link"`
    Description string `xml:"description"`
    PubDate string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error){
    req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
    if err != nil{
        return nil, err
    }

    req.Header.Set("User-Agent", "gator")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil{
        return nil, err
    }
    defer resp.Body.Close()

    bytes, err := io.ReadAll(resp.Body)
    if err != nil{
        return nil, err
    }

    var feed RSSFeed

    err = xml.Unmarshal(bytes, &feed)
    if err != nil{
        return nil, err
    }

    feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
    feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
    for i :=  range feed.Channel.Item{
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
    }

    return &feed, nil
}
