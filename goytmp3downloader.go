package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/francoiscolombo/goytmp3downloader/fetch"
	"github.com/francoiscolombo/goytmp3downloader/play"
	"github.com/francoiscolombo/goytmp3downloader/search"
	"gopkg.in/gookit/color.v1"
)

const (
	versionNumber = "0.2"
	versionName   = "lightning plasma"
	youtubeAPIKey = "AIzaSyDL_uUF6-4Kxf3NwOPFhOqyBX71SRcFKhE"
)

type parameters struct {
	Search  string
	Fetch   string
	Path    string
	Play    string
	Full    bool
	Version bool
	Help    bool
}

func usage() {
	fmt.Println(`goytmp3downloader: CLI for downloading mp3 from youtube

 1/ Search video list
    $ goytmp3downloader --search <video name>
    this allow you to search for a video list by entering the name of the
    video you are searching.
    this will display a list of video ID that you can use to fetch the mp3
    later. some useful information are also displayed, to help you choose
    the proper video to download.

 2/ Download video or mp3
	$ goytmp3downloader --fetch <video id> [--path <download path>] [--full]
	you have to enter a video id, that you can retrieve with the previous
	option. you can also give a path where to download the video, if you don't
	then the current path is going to be used.
	and if you set the flag 'full' then you will just download the video, and
	don't extract the mp3. otherwise, the video is downloaded and the audio is
	extracted, then the video from which you extract the audio is removed.

 3/ Playback mp3
	$ goytmp3downloader --play <mp3 path>
	you have to enter the path where is located the mp3 you want to play.

  4/ Additional commands
     $ goytmp3downloader --version
     $ goytmp3downloader --help

For the full documentation please refer to:
https://github.com/francoiscolombo/goytmp3downloader`)
	fmt.Println("")
}
func main() {

	fmt.Printf("\n%s on ", color.FgGreen.Render("Welcome"))
	color.S256(15, 208).Print("goytmp3downloader")
	fmt.Printf("\n----------------------------\n\n")

	var params parameters

	flag.StringVar(&params.Search, "search", "???", "search for videos / mp3 to download")
	flag.StringVar(&params.Fetch, "fetch", "???", "id of the video / mp3 to download")
	flag.StringVar(&params.Path, "path", ".", "allow to download to another path instead of the current one")
	flag.StringVar(&params.Play, "play", "???", "allow to play a mp3")
	flag.BoolVar(&params.Full, "full", false, "download the video but don't extract the mp3")
	flag.BoolVar(&params.Version, "version", false, "display the version and exit")
	flag.BoolVar(&params.Help, "help", false, "display help and exit")

	flag.Parse()

	if params.Help {
		usage()
		os.Exit(0)
	}

	if params.Version {
		fmt.Print("version ")
		color.C256(69).Print(versionNumber)
		fmt.Print(" ")
		color.S256(124, 231).Printf("(%s)", versionName)
		fmt.Println()
		os.Exit(0)
	}

	if params.Search != "???" {
		ccmd := color.FgLightBlue.Render
		cprm := color.FgLightCyan.Render
		fmt.Printf("- %s command selected, with the following parameters:\n", ccmd("Search"))
		fmt.Printf("  > Search for video with title : '%s'\n", cprm(params.Search))
		err := search.VideoFromYoutube(youtubeAPIKey, params.Search)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	if params.Fetch != "???" {
		ccmd := color.FgLightBlue.Render
		cprm := color.FgLightCyan.Render
		fmt.Printf("- %s command selected, with the following parameters:\n", ccmd("Fetch"))
		fmt.Printf("  > Download video with id : '%s'\n", cprm(params.Fetch))
		fmt.Printf("  > Download to path : '%s'\n", cprm(params.Path))
		fmt.Printf("  > Don't extract mp3 : '%s'\n", cprm(fmt.Sprintf("%t", params.Full)))
		err := fetch.VideoFromYoutube(params.Fetch, params.Path, params.Full)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	if params.Play != "???" {
		ccmd := color.FgLightBlue.Render
		cprm := color.FgLightCyan.Render
		cctrlc := color.FgLightYellow.Render
		fmt.Printf("- %s command selected, with the following parameters:\n", ccmd("Play"))
		fmt.Printf("  > Mp3 path : '%s'\n", cprm(params.Play))
		fmt.Printf("  Press %s to stop...\n", cctrlc("CTRL+C"))
		err := play.Mp3(params.Play)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	color.Print("<red>Sorry my friend</>, but you didn't give me the good parameters, so I'm not able to help you!\n")
	color.Print("<cyan>Maybe a little help can be what you really need?</> Okay, this should be usefull then...\n\n")
	usage()

}
