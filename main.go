package main

import "fmt"
import "strings"
import "net/http"
import "os"
import "io"
import "strconv"


func main(){
	var url string
	buffer := make([]byte, 1024*1024)

	fmt.Println("Please provide a url to download from")
	fmt.Scanf("%s", &url)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Incorrect or invalid http link")
		return
	}

	if resp.StatusCode == 200 || resp.StatusCode == 206{
		fmt.Printf("Valid html\n")
	} else{
		fmt.Println("Invalid html")
		defer resp.Body.Close()
	}
	
	contentType := strings.Split(resp.Header.Get("Content-Type"), ";")[0]
	fileSize := resp.Header.Get("Content-Length")
	fileType := ConvertType(contentType) 

	fmt.Printf("file size: %s\n", fileSize)
	fmt.Printf("Test%s\n", fileType)
	output, _ := os.Create("test" + fileType)
	totalChunks := 0

	if fileSize != ""{
		fmt.Printf("Loading ")
		
		totalChunk, err := strconv.Atoi(fileSize)
		if err != nil{
			fmt.Printf("Failed at string conversion")
			return;
		}
		totalChunks = totalChunk
	}

	var chunksRead int = 0
	completion := 0.0
	tick := 0.01
	loadingBar:= LoadingBar("", 0)
	fmt.Printf("%s 0%", loadingBar)
	for {
		if totalChunks != 0{
			completion = float64(chunksRead) / float64(totalChunks)

		}

		chunk, err := resp.Body.Read(buffer)

		// Check for valid chunks
		if chunk > 0{
			output.Write(buffer[:chunk])
		}

		if err != nil{
			if err == io.EOF{
				break
			}

			fmt.Println("Something failed in the download")
			return
		}

		// Every 5% prints / 
		if totalChunks != 0{
			if completion > tick{
				loadingBar = LoadingBar(loadingBar, completion)
				fmt.Printf("\r%s %d%%", loadingBar, int(completion * 100))

				
			}
		}

		chunksRead += chunk
	}
	fmt.Println()
}
