# cutter_go

Go implementation of the MP3 endpoint from the Python CutterKo service.

## Vercel deployment

This repository includes `api/index.go` and `vercel.json` for Vercel's Go
runtime. Set the Vercel project root to `cutter_go`.

From the repository root:

```powershell
cd cutter_go
npm install -g vercel
vercel login
vercel
vercel --prod
```

In the Vercel dashboard, the equivalent setup is:

1. Import the repository.
2. Set **Root Directory** to `cutter_go`.
3. Keep the framework preset as **Other**.
4. Deploy.

After deployment, verify the function:

```powershell
curl https://YOUR_PROJECT.vercel.app/health
curl "https://YOUR_PROJECT.vercel.app/mp3?url=https%3A%2F%2Fwww.youtube.com%2Fwatch%3Fv%3DVIDEO_ID"
```

The function requires executable `yt-dlp` and FFmpeg binaries. Configure their
paths with:

```text
YTDLP_PATH=/var/task/bin/yt-dlp
FFMPEG_PATH=/var/task/bin/ffmpeg
```

Vercel does not install these tools automatically. Add compatible Linux
executables under `bin/` only if their sizes fit your Vercel function limits.
The deployment environment must support executing those binaries.
Windows `.exe` files from local development cannot be used by a Vercel
function. If the deployment fails because the binaries cannot be bundled or
executed, deploy this service to Railway, Render, Fly.io, or a VPS instead;
those platforms are a better fit for long-running media conversion.

## Requirements

- Go
- `yt-dlp`
- FFmpeg

## Run

```powershell
go run ./cmd/cutter-go
```

The local service listens on `http://localhost:8180` by default. Vercel
provides the public HTTPS URL.

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
