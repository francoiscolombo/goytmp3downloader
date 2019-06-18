# goytmp3downloader
cli for video and mp3 download written in golang

_*Please note that this is still a work in progress...*_

## Prerequisite

please note that to extract mp3 from the video, you will need to install ffmpeg in your system path.
see [FFMpeg download](https://ffmpeg.org/download.html) to download the proper version for your system.

If you want to use the source code, you must have a proper go workspace already set up.

Then after that, it's as easy that:

    $ git clone https://github.com/francoiscolombo/goytmp3downloader.git
    $ cd goytmp3downloader
    $ go get
    $ go install goytmp3downloader.go

This should produce an executable in your ``${GOPATH}/bin`` directory.

## Search for videos or mp3

Simply use the following command:

    λ goytmp3downloader --search "Led Zeppelin - Stairway To Heaven (Lyrics)"

Then you should have this output:

    Welcome on goytmp3downloader
    ----------------------------
    
    - Search command selected, with the following parameters:
      > Search for video with title : 'Led Zeppelin - Stairway To Heaven (Lyrics)'
     100% |████████████████████████████████████████|  [1m8s:0s]
    +-------------+----------------------+--------------------------------+-----------------------------------------+
    |     ID      |         DATE         |             TITLE              |               DESCRIPTION               |
    +-------------+----------------------+--------------------------------+-----------------------------------------+
    | Nnu1E5Kslig | Sunday, 18-Mar-18    | Led Zeppelin - Stairway To     | Led Zeppelin - Stairway To              |
    |             |                      | Heaven (Lyrics)                | Heaven with lyrics There's              |
    |             |                      |                                | a lady who's sure All that              |
    |             |                      |                                | glitters is gold And she's              |
    |             |                      |                                | buying a stairway to heaven             |
    |             |                      |                                | When she gets there she ...             |
    | 9Bbnkw6gIN8 | Sunday, 14-Oct-18    | Stairway to Heaven | Led       | I have no claim over the song           |
    |             |                      | Zeppelin (LYRICS / MUSIC VIDEO | or the source footage for this          |
    |             |                      | / ORIGINAL)                    | video. All is owned by Led              |
    |             |                      |                                | Zeppelin and the production             |
    |             |                      |                                | behind the band and this                |
    |             |                      |                                | song/performance.                       |
    | D9ioyEvdggk | Sunday, 13-Aug-17    | Led Zeppelin -  Stairway To    | Stairway To Heaven There's              |
    |             |                      | Heaven ᴴᴰ (Legendado/Tradução  | a lady who's sure all that              |
    |             |                      | PTBR)                          | glitters is gold And she's              |
    |             |                      |                                | buying a stairway to heaven             |
    |             |                      |                                | When she gets there she knows,          |
    |             |                      |                                | if the stores are all ...               |
    | DDo4CA13LbY | Friday, 10-Oct-14    | Jimmy Page: How Stairway to    | Subscribe to BBC News                   |
    |             |                      | Heaven was written - BBC News  | www.youtube.com/bbcnews                 |
    |             |                      |                                | Stairway to Heaven was one of           |
    |             |                      |                                | the biggest rock songs of the           |
    |             |                      |                                | 1970s - loved, imitated and             |
    |             |                      |                                | sometimes ...                           |
    +-------------+----------------------+--------------------------------+-----------------------------------------+

## Download a mp3

Once you have a list of video corresponding to your search, you can download a video from this list.

Simply use the video *ID* with this command:

    λ goytmp3downloader --fetch Nnu1E5Kslig --path D:\Data\

    Welcome on goytmp3downloader
    ----------------------------
    
    - Fetch command selected, with the following parameters:
      > Download video with id : 'Nnu1E5Kslig'
      > Download to path : 'D:\Data\'
      > Don't extract mp3 : 'false'
     100% |████████████████████████████████████████| (547.1 kB/s) [48s:0s]
    Extracting audio ...
    size=    7523kB time=00:08:01.41 bitrate= 128.0kbits/s speed=  87x

    Extracted audio:D:\Data\/led-zeppelin-stairway-to-heaven-lyrics.mp3

## Simply download the video, without extracting mp3

    λ goytmp3downloader --fetch Nnu1E5Kslig --path D:\Data\ --full
    
    Welcome on goytmp3downloader
    ----------------------------
    
    - Fetch command selected, with the following parameters:
      > Download video with id : 'Nnu1E5Kslig'
      > Download to path : 'D:\Data\'
      > Don't extract mp3 : 'true'
     100% |████████████████████████████████████████| (440.4 kB/s) [55s:0s]

## Play a mp3, directly from the command line

    λ goytmp3downloader --play \Data\led-zeppelin-stairway-to-heaven-lyrics.mp3
    
    Welcome on goytmp3downloader
    ----------------------------
    
    - Play command selected, with the following parameters:
      > Mp3 path : '\Data\led-zeppelin-stairway-to-heaven-lyrics.mp3'
      Press CTRL+C to stop...
