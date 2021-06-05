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
	err := os.MkdirAll(path.Join(baseFolder, "original"), 0755)
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(baseFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".mkv" {
			continue
		}

		partialFile := strings.ReplaceAll(file.Name(), ".mkv", ".part.mp4")
		finalFile := strings.ReplaceAll(file.Name(), ".mkv", ".mp4")

		input := path.Join(baseFolder, file.Name())
		partial := path.Join(baseFolder, partialFile)
		final := path.Join(baseFolder, finalFile)
		original := path.Join(baseFolder, "original", file.Name())

		if containsFile(files, finalFile) {
			err = os.Rename(input, original)
			if err != nil {
				log.Fatal(err)
			}

			continue
		}

		if containsFile(files, partialFile) {
			err := os.Remove(partial)
			if err != nil {
				log.Fatal(err)
			}
		}

		cmd := exec.Command("ffmpeg", "-i", input, "-vcodec", "h264", "-acodec", "ac3", partial)
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

		err = os.Rename(input, original)
		if err != nil {
			log.Fatal(err)
		}
	}

}
