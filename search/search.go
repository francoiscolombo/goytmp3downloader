package search

import (
	"fmt"
	"os"

	"github.com/francoiscolombo/goytmp3downloader/utube"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/gookit/color.v1"
)

/*
VideoFromYoutube search on youtube the 20 first links that respond to your query
*/
func VideoFromYoutube(youtubeAPIKey, title string) error {
	var videosList []utube.VideoMetaData
	pageToken := ""
	pageNumber := 0
	cpage := color.FgLightMagenta.Render
	cfind := color.FgLightGreen.Render
	for {
		pageNumber = pageNumber + 1
		fmt.Printf("> Search Page #%s\n", cpage(fmt.Sprintf("%d", pageNumber)))
		videos, nextPageToken, err := utube.SearchVideos(title, youtubeAPIKey, pageToken)
		if err != nil {
			return err
		}
		fmt.Printf(">>> Find %s new videos\n", cfind(fmt.Sprintf("%d", len(videos))))
		for _, video := range videos {
			videosList = append(videosList, video)
			if len(videosList) >= utube.NbLinksToRetrieve {
				break
			}
		}
		if len(videosList) >= utube.NbLinksToRetrieve {
			break
		}
		pageToken = nextPageToken
	}
	cid := color.FgYellow.Render
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", "Id", "Date", "Title", "Description"})
	for i, item := range videosList {
		table.Append([]string{
			fmt.Sprintf("%02d", (i + 1)),
			cid(item.ID),
			item.Date,
			item.Title,
			item.Description,
		})
	}
	table.Render()
	return nil
}
