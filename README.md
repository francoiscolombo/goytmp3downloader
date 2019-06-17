# goytmp3downloader
cli for video and mp3 download written in golang

## Prerequisite
please note that to extract mp3 from the video, you will need to install ffmpeg in your system path.
see [FFMpeg download](https://ffmpeg.org/download.html) to download the proper version for your system.

## Search for videos or mp3

Simply use the following command:

    λ go run goytmp3downloader.go --search "Led Zeppelin - Stairway To Heaven (Lyrics)"

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

λ go run goytmp3downloader.go --fetch Nnu1E5Kslig --path D:\Data\

    Welcome on goytmp3downloader
    ----------------------------
    
    - Fetch command selected, with the following parameters:
      > Download video with id : 'Nnu1E5Kslig'
      > Download to path : 'D:\Data\'
      > Don't extract mp3 : 'false'
     100% |████████████████████████████████████████| (547.1 kB/s) [48s:0s]
    Extracting audio ...
    ffmpeg version N-94057-g1c3ed11893 Copyright (c) 2000-2019 the FFmpeg developers
      built with gcc 8.3.1 (GCC) 20190414
      configuration: --enable-gpl --enable-version3 --enable-sdl2 --enable-fontconfig --enable-gnutls --enable-iconv --enable-libass --enable-libdav1d --enable-libbluray --enable-libfreetype --enable-libmp3lame --enable-libopencore-amrnb --enable-libopencore-amrwb --enable-libopenjpeg --enable-libopus --enable-libshine --enable-libsnappy --enable-libsoxr --enable-libtheora --enable-libtwolame --enable-libvpx --enable-libwavpack --enable-libwebp --enable-libx264 --enable-libx265 --enable-libxml2 --enable-libzimg --enable-lzma --enable-zlib --enable-gmp --enable-libvidstab --enable-libvorbis --enable-libvo-amrwbenc --enable-libmysofa --enable-libspeex --enable-libxvid --enable-libaom --enable-libmfx --enable-amf --enable-ffnvcodec --enable-cuvid --enable-d3d11va --enable-nvenc --enable-nvdec --enable-dxva2 --enable-avisynth --enable-libopenmpt
      libavutil      56. 29.100 / 56. 29.100
      libavcodec     58. 53.100 / 58. 53.100
      libavformat    58. 27.103 / 58. 27.103
      libavdevice    58.  7.100 / 58.  7.100
      libavfilter     7. 55.100 /  7. 55.100
      libswscale      5.  4.101 /  5.  4.101
      libswresample   3.  4.100 /  3.  4.100
      libpostproc    55.  4.100 / 55.  4.100
    Input #0, mov,mp4,m4a,3gp,3g2,mj2, from 'D:\Data\/led-zeppelin-stairway-to-heaven-lyrics.mp4':
      Metadata:
        major_brand     : mp42
        minor_version   : 0
        compatible_brands: isommp42
        creation_time   : 2018-11-09T08:57:53.000000Z
      Duration: 00:08:01.40, start: 0.000000, bitrate: 379 kb/s
        Stream #0:0(und): Video: h264 (Constrained Baseline) (avc1 / 0x31637661), yuv420p(tv, bt709), 640x360 [SAR 1:1 DAR 16:9], 281 kb/s, 29.97 fps, 29.97 tbr, 30k tbn, 59.94 tbc (default)
        Metadata:
          creation_time   : 2018-11-09T08:57:53.000000Z
          handler_name    : ISO Media file produced by Google Inc. Created on: 11/09/2018.
        Stream #0:1(und): Audio: aac (LC) (mp4a / 0x6134706D), 44100 Hz, stereo, fltp, 95 kb/s (default)
        Metadata:
          creation_time   : 2018-11-09T08:57:53.000000Z
          handler_name    : ISO Media file produced by Google Inc. Created on: 11/09/2018.
    Stream mapping:
      Stream #0:1 -> #0:0 (aac (native) -> mp3 (libmp3lame))
    Press [q] to stop, [?] for help
    Output #0, mp3, to 'D:\Data\/led-zeppelin-stairway-to-heaven-lyrics.mp3':
      Metadata:
        major_brand     : mp42
        minor_version   : 0
        compatible_brands: isommp42
        TSSE            : Lavf58.27.103
        Stream #0:0(und): Audio: mp3 (libmp3lame), 44100 Hz, stereo, fltp (default)
        Metadata:
          creation_time   : 2018-11-09T08:57:53.000000Z
          handler_name    : ISO Media file produced by Google Inc. Created on: 11/09/2018.
          encoder         : Lavc58.53.100 libmp3lame
    size=    7523kB time=00:08:01.41 bitrate= 128.0kbits/s speed=  87x
    video:0kB audio:7522kB subtitle:0kB other streams:0kB global headers:0kB muxing overhead: 0.004492%
    
    Extracted audio:D:\Data\/led-zeppelin-stairway-to-heaven-lyrics.mp3

## Simply download the video, without extracting mp3

    λ go run goytmp3downloader.go --fetch Nnu1E5Kslig --path D:\Data\ --full
    
    Welcome on goytmp3downloader
    ----------------------------
    
    - Fetch command selected, with the following parameters:
      > Download video with id : 'Nnu1E5Kslig'
      > Download to path : 'D:\Data\'
      > Don't extract mp3 : 'true'
     100% |████████████████████████████████████████| (440.4 kB/s) [55s:0s]

