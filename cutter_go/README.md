# cutter_go

Go implementation of the MP3 endpoint from the Python CutterKo service.

## Requirements

- Go
- `yt-dlp`
- FFmpeg

## Run

```powershell
go run .
```

The service listens on `http://localhost:8180` by default.

## Endpoints

```text
GET /health
GET /mp3?url=https://www.youtube.com/watch?v=VIDEO_ID
```

The MP3 endpoint:

1. Validates the URL is a supported HTTPS YouTube URL.
2. Uses yt-dlp to download the best audio stream.
3. Uses FFmpeg to convert it to 192 kbps MP3.
4. Streams the MP3 to the caller.
5. Deletes the per-request temporary directory after the response finishes.

Only download content you own or are authorized to download.
