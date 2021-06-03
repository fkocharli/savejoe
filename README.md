<h1 align='center'>Save Joe - Simple Audio Streaming App</h1>

## Usage

- Create new folder in `static` folder with name of your file. 
- Convert your audio file to HLS with below command and put output to folder created on previous step.

`ffmpeg -i audio_file_name.mp3 -c:a libmp3lame -b:a 128k -map 0:0 -f segment -segment_time 10 -segment_list outputlist.m3u8 -segment_format mpegts output%03d.ts`

## Build and Run

- `docker build -t savejoe .`
- `docker run --name savejoe -p 8080:8080 -d savejoe`

---

To test it use any client or player which supports HLS (for ex. VLC) and send request to below URL:

```http://localhost:8080/v1/stream/YourFolderName```

As a sample I've already created "coldplay" folder and icluded static files in repository. So, you can get played audio from below link.

```http://localhost:8080/v1/stream/coldplay/```
