package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	// open the file
	file, err := os.OpenFile(".bashrc", os.O_APPEND|os.O_RDWR, os.ModeAppend)

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	var EnhancementNeeded bool = true

	// read line by line
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		if fileScanner.Text() == "Begin project specific Enhancements" {
			EnhancementNeeded = false
		}
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	if EnhancementNeeded {
		d := []string{"Begin project specific Enhancements", "tset", "It is easy to learn Go."}

		// writing to file
		for _, v := range d {
			fmt.Println(v)
			l, err := fmt.Fprintln(file, v)
			if err != nil {
				log.Fatalf("Error when writing to file: %s", err)
			}
			fmt.Println(l, "bytes written successfully")
		}
		fmt.Println("File written successfully")
	}

	err = file.Close()
	if err != nil {
		log.Fatalf("Error when closing file: %s", err)
	}

	// open the file
	file2, err := os.Open(".bashrc")

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner2 := bufio.NewScanner(file2)
	for fileScanner2.Scan() {
		fmt.Println(fileScanner2.Text())
	}

	err = file2.Close()
	if err != nil {
		log.Fatalf("Error when closing file: %s", err)
	}

	log.Print(runtime.GOOS)

	cmd := exec.Command("bash", "-c", "chmod 700 *")
	log.Printf("Running command and waiting for it to finish...")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	if errr := cmd.Run(); errr != nil {
		fmt.Println("Command finished with error:", errr)
	}

	searchDir := "/mnt/c/your/workspace"

	fileList := []string{}
	filepath.Walk(searchDir, func(path string, f os.FileInfo, errrr error) error {
		log.Print(path)
		fileList = append(fileList, path)
		if path == ".bashrc" {
			stats, err := os.Stat(path)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Permission File Before: %s\n", stats.Mode())
			err = os.Chmod(path, 0700)
			if err != nil {
				log.Fatal(err)
			}

			stats, err = os.Stat(path)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Permission File After: %s\n", stats.Mode())
		}
		return nil
	})

	for _, file := range fileList {
		fmt.Println(file)
	}
}
