package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"os"
	"io"
	"strings"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func main() {
	xmlFile, err := os.Open("facultybookplate.xml")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return
	}

	fmt.Println("file opened")
	defer xmlFile.Close()

	utf8Reader := transform.NewReader(xmlFile, unicode.UTF8.NewDecoder())
	decoder := xml.NewDecoder(utf8Reader)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		enc, err := determineEncoding(input)
		if err != nil {
			return nil, err
		}
		return transform.NewReader(input, enc.NewDecoder()), nil
	}

	// Create the CSV File
	csvFile, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Add UTF-8 BOM for Excel compatability
	csvFile.Write([]byte{0xEF, 0xBB, 0xBF})
	// Create the Writer
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write CSV header
	header := []string{"PostID", "Book Plate Year", "HonoreeName", "FirstName", "LastName", "Department", "College", "Book Title", "Author", "Genre", "Isbn", "Notes from Honoree", "Book Choice", "Statement", "Catalog Record", "Book Cover"}
	writer.Write(header)

	var postFeed BookPlate
	err = decoder.Decode(&postFeed)
	if err != nil {
		fmt.Printf("error decoding xml: %v\n", err)
		return
	}

	var PostSlice []Post
	for _, item := range postFeed.Channel.Item {
		post := Post{
			PostID:      item.PostID,
			HonoreeName: item.Title,
		}

		for _, category := range item.Category {
			if category.Domain == "bookplate_year" {
				post.BookPlateYear = category.Nicename
			}
			if category.Domain == "bookplate_genre" {
				post.Genre = append(post.Genre, category.Nicename)
			}
		}

		// Extract metadata
		for _, postmeta := range item.Postmeta {
			// Skip metadata keys starting with underscore (internal WordPress use)
			if strings.HasPrefix(postmeta.MetaKey, "_") {
				continue
			}

			// Map metadata to struct fields
			switch postmeta.MetaKey {
			case "first_name":
				post.FirstName = postmeta.MetaValue
			case "last_name":
				post.LastName = postmeta.MetaValue
			case "notes_from_honoree":
				post.ChosenReason = postmeta.MetaValue
			case "book_title":
				post.BookTitle = postmeta.MetaValue
			case "department":
				post.Department = postmeta.MetaValue
			case "college":
				post.College = postmeta.MetaValue
			case "author":
				post.Author = postmeta.MetaValue
			case "book_choice":
				post.BookChoice = postmeta.MetaValue
			case "statement":
				post.Statement = postmeta.MetaValue
			case "catalog_record":
				post.CatalogRecord = postmeta.MetaValue
			case "isbn":
				post.Isbn = postmeta.MetaValue
			case "book_cover":
				post.BookCover = postmeta.MetaValue
				// Add more cases as needed for additional metadata fields
			}
		}
		PostSlice = append(PostSlice, post)

		// Write to CSV
		record := []string{
			post.PostID,
			post.BookPlateYear,
			post.HonoreeName,
			post.FirstName,
			post.LastName,
			post.Department,
			post.College,
			post.BookTitle,
			post.Author,
			strings.Join(post.Genre, ", "),
			post.Isbn,
			post.ChosenReason,
			post.BookChoice,
			post.Statement,
			post.CatalogRecord,
			post.BookCover,
		}
		writer.Write(record)

	}

	fmt.Printf("Processed %d posts and wrote them to output.csv\n", len(PostSlice))
}




func determineEncoding(r io.Reader) (encoding.Encoding, error) {
	// Read a small chunk to detect encoding
	b, err := io.ReadAll(io.LimitReader(r, 1024))
	if err != nil {
		return unicode.UTF8, err
	}
	
	e, _, _ := charset.DetermineEncoding(b, "")
	return e, nil
}