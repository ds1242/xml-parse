package main

import (
	"encoding/xml"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	xmlFile, err := os.Open("facultybookplate.xml")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return
	}

	fmt.Println("file opened")
	defer xmlFile.Close()

	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Printf("error reading file content: %v\n", err)
		return
	}

	// Create the CSV File
	csvFile, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()
	// Create the Writer
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write CSV header
	header := []string{"PostID", "BookTitle", "HonoreeName", "FirstName", "LastName", "Title"}
	writer.Write(header)

	var postFeed BookPlate
	err = xml.Unmarshal(byteValue, &postFeed)
	if err != nil {
		fmt.Printf("error unmarshalling xml: %v\n", err)
		return
	}

	var PostSlice []Post
	for _, item := range postFeed.Channel.Item {
		post := Post{
			PostID:      item.PostID,
			HonoreeName: item.Title,
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
			case "statement":
				post.Statement = postmeta.MetaValue
			case "catalog_record":
				post.CatalogRecord = postmeta.MetaValue
			// Add more cases as needed for additional metadata fields
			}
		}
		PostSlice = append(PostSlice, post)

		// Write to CSV
		record := []string{
			post.PostID,
			post.BookTitle,
			post.HonoreeName,
			post.FirstName,
			post.LastName,
			post.BookTitle,
		}
		writer.Write(record)

	}

	fmt.Printf("Processed %d posts and wrote them to output.csv\n", len(PostSlice))
}