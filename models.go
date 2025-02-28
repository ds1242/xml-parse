package main

import (
	"encoding/xml"
)

type BookPlate struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Excerpt string   `xml:"excerpt,attr"`
	Content string   `xml:"content,attr"`
	Wfw     string   `xml:"wfw,attr"`
	Dc      string   `xml:"dc,attr"`
	Wp      string   `xml:"wp,attr"`
	Channel struct {
		Text        string `xml:",chardata"`
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		PubDate     string `xml:"pubDate"`
		Language    string `xml:"language"`
		WxrVersion  string `xml:"wxr_version"`
		BaseSiteURL string `xml:"base_site_url"`
		BaseBlogURL string `xml:"base_blog_url"`
		Author      []struct {
			Text              string `xml:",chardata"`
			AuthorID          string `xml:"author_id"`
			AuthorLogin       string `xml:"author_login"`
			AuthorEmail       string `xml:"author_email"`
			AuthorDisplayName string `xml:"author_display_name"`
			AuthorFirstName   string `xml:"author_first_name"`
			AuthorLastName    string `xml:"author_last_name"`
		} `xml:"author"`
		Generator string `xml:"generator"`
		Item      []struct {
			Text    string `xml:",chardata"`
			Title   string `xml:"title"`
			Link    string `xml:"link"`
			PubDate string `xml:"pubDate"`
			Creator string `xml:"creator"`
			Guid    struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			Description     string `xml:"description"`
			Encoded         string `xml:"encoded"`
			PostID          string `xml:"post_id"`
			PostDate        string `xml:"post_date"`
			PostDateGmt     string `xml:"post_date_gmt"`
			PostModified    string `xml:"post_modified"`
			PostModifiedGmt string `xml:"post_modified_gmt"`
			CommentStatus   string `xml:"comment_status"`
			PingStatus      string `xml:"ping_status"`
			PostName        string `xml:"post_name"`
			Status          string `xml:"status"`
			PostParent      string `xml:"post_parent"`
			MenuOrder       string `xml:"menu_order"`
			PostType        string `xml:"post_type"`
			PostPassword    string `xml:"post_password"`
			IsSticky        string `xml:"is_sticky"`
			Category        []struct {
				Text     string `xml:",chardata"`
				Domain   string `xml:"domain,attr"`
				Nicename string `xml:"nicename,attr"`
			} `xml:"category"`
			Postmeta []struct {
				Text      string `xml:",chardata"`
				MetaKey   string `xml:"meta_key"`
				MetaValue string `xml:"meta_value"`
			} `xml:"postmeta"`
		} `xml:"item"`
	} `xml:"channel"`
}

type Post struct {
	PostID       string
	HonoreeName  string
	FirstName    string
	LastName     string
	Department   string
	College      string
	BookTitle    string
	Author 		string
	Genre        string
	Isbn         string
	ChosenReason string
	Statement 	string
	CatalogRecord string
}
