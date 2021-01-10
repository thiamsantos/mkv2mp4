package main

import (
	"bufio"
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
func containsString(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func readLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func main() {
	stateFilePath := path.Join(baseFolder, "mkv2mp4.state")
	stateFile, err := os.OpenFile(stateFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer stateFile.Close()

	content, err := ioutil.ReadFile(stateFilePath)
	if err != nil {
		log.Fatal(err)
	}
	completedFiles := deleteEmpty(strings.Split(string(content), "\n"))

	files, err := ioutil.ReadDir(baseFolder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".mkv" {
			continue
		}

		finalFile := strings.ReplaceAll(file.Name(), ".mkv", ".mp4")

		if containsFile(files, finalFile) && containsString(completedFiles, finalFile) {
			continue
		}

		input := path.Join(baseFolder, file.Name())
		output := path.Join(baseFolder, finalFile)

		if containsFile(files, finalFile) {
			err := os.Remove(output)
			if err != nil {
				log.Fatal(err)
			}
		}

		cmd := exec.Command("ffmpeg", "-i", input, "-vcodec", "h264", "-acodec", "mp3", output)
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

		if _, err = stateFile.WriteString(finalFile + "\n"); err != nil {
			log.Fatal(err)
		}
	}

}
