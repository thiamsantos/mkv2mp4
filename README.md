# mkv2mp4

CLI to convert mkv files to mp4 using [ffmpeg](https://ffmpeg.org/).

### Installation

You need `go` installed and `GOBIN` in your `PATH`. Once that is done, run the
command:

```shell
$ go get -u github.com/thiamsantos/mkv2mp4
```

## Usage

```sh
$ mkv2mp4
```

It will find all the mkv in the current working directory, convert them to mp4,
storing the name of the completed files in `mkv2mp4.state` file.
