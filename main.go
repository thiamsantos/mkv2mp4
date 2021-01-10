package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

const baseFolder = "."

func containsFile(arr []os.FileInfo, str string) bool {
	for _, a := range arr {
		if a.Name() == str {
			return true
		}
	}
	return false
}

func main() {
	files, err := ioutil.ReadDir(baseFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".mkv" {
			continue
		}

		partialFile := strings.ReplaceAll(file.Name(), ".mkv", ".mp4.part")
		finalFile := strings.ReplaceAll(file.Name(), ".mkv", ".mp4")

		input := path.Join(baseFolder, file.Name())
		partial := path.Join(baseFolder, partialFile)
		final := path.Join(baseFolder, finalFile)

		if containsFile(files, finalFile) {
			continue
		}

		if containsFile(files, partialFile) {
			err := os.Remove(partial)
			if err != nil {
				log.Fatal(err)
			}
		}

		cmd := exec.Command("ffmpeg", "-i", input, "-vcodec", "h264", "-acodec", "mp3", partial)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		err = cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}

		err = os.Rename(partial, final)
		if err != nil {
			log.Fatal(err)
		}
	}

}
