# Samo: Light-weight streaming media server
Samo is a simple lightweight streaming media server utilizes HTML 5 video
functionality. It's easy to deploy and streams media fast.


## Why another streaming media server?
Emby would often get stuggerish when trying to watch media at it's original
quality via the Web UI. I tried to figure out why, but couldn't find the
root cause. So then I tried to use HTML 5 video element and play the HQ
MP4 directly, this to my astonishment didn't get stuggerish at all. So then
I decided to create a simple Streaming Media server that plays video
files(mp4 and mkv) by directly serving the files to an HTML 5 player.

## Implementation details
Samo uses Golang to partially serve the video files as requested by the HTML 5
video player in the browser. The browser requests partial content uses HTTP
range requests, then the server responds with partial content responses. The
HTML 5 video player in the browser makes sure enough content is buffered such
that the video plays smoothly, but at the same makes sure not too many content
is requested.


## Current functionality
- List MP4 and MKV files in directory
- Plays MP4 and MKV files fluently using HTML 5 video player

## Roadmap
- Beautiful stand-alone UI in polymer which consumes REST API
- Add SQLlite or MongoDB database such that user can add additional directories
- Add HTTP auth as simple authentication to prevent public access
