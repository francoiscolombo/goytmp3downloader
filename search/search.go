package search

import (
	"os"

	"github.com/francoiscolombo/goytmp3downloader/utube"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/gookit/color.v1"
)

/*
VideoFromYoutube search on youtube the 20 first links that respond to your query
*/
func VideoFromYoutube(youtubeAPIKey, title string) error {
	videosList, err := utube.SearchVideos(title, youtubeAPIKey)
	if err != nil {
		return err
	}
	cid := color.FgYellow.Render
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Date", "Title", "Description"})
	for _, item := range videosList {
		table.Append([]string{
			cid(item.ID),
			item.Date,
			item.Title,
			item.Description,
		})
	}
	table.Render()
	return nil
}
