package main

import "fmt"
import "strings"
import "net/http"
import "os"
import "io"
import "strconv"
import "sync"


func main(){
	var url string
	buffer := make([]byte, 1024*1024)

	fmt.Println("Please provide a url to download from")
	fmt.Scanf("%s", &url)
	resp := Open_url(url, 0, 0)

	if (resp == nil){
		fmt.Println("Failed to open url")
	}
	
	contentType := strings.Split(resp.Header.Get("Content-Type"), ";")[0]
	fileSize := resp.Header.Get("Content-Length")
	fileType := ConvertType(contentType) 

	// Eventually get name from header and replace "test"
	output, _ := os.Create("test" + fileType)
	totalChunks := 0
	dividedChunks := 0
	
	// For future need more limiters for multithreading, server may not allow some things. Could make some funky things happen
	if fileSize != ""{
		fmt.Println("Loading ")
		
		totalChunk, err := strconv.Atoi(fileSize)
		if err != nil{
			fmt.Printf("Failed at string conversion")
			return;
		}
		totalChunks = totalChunk
		dividedChunks = totalChunks / 8
		fmt.Printf("total chunks = %d\ndivided chunks = %d\n", totalChunks, dividedChunks)

		// Downloader for multi-threading
		Spin_threads(totalChunks, dividedChunks, url, output)


	}	else{ // Singlethread
		Download_chunk(resp, totalChunks, output, buffer)
	}

	fmt.Println()
}


func Download_chunk(resp *http.Response, totalChunks int, output *os.File, buffer []byte){
	var chunksRead int = 0 
	completion := 0.0
	tick := 0.01
	loadingBar:= LoadingBar("", 1)

	fmt.Printf("%s 0%%", loadingBar)
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
}


func Multi_chunk(url string, startByte int, endByte int, output *os.File){
	resp := Open_url(url, startByte, endByte)

	if (resp == nil){
		return
	}
	
	chunk, _ := io.ReadAll(resp.Body)
	output.Seek(int64(startByte), 0)
	output.Write(chunk)

	
	resp.Body.Close()
}

func Open_url(url string, startByte int, endByte int) *http.Response{

	var resp *http.Response
	var err error

	if (startByte == 0 && endByte == 0){
		resp, err = http.Get(url)

	} else{
		// Get limited bytes 
		limitReq, _ := http.NewRequest("GET", url, nil)
		limitReq.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", startByte, endByte))
		resp, err = http.DefaultClient.Do(limitReq)
	}

	if err != nil {
		fmt.Printf("Incorrect or invalid http link")
		return nil
	}

	if resp.StatusCode == 200 || resp.StatusCode == 206{
		fmt.Printf("Valid html\n")
	} else{
		fmt.Println("Invalid html")
	}

	return resp
}


func Spin_threads(totalChunks int, dividedChunks int, url string, output *os.File){
	
	startByte := 0
	endByte := dividedChunks
	var waitGroup sync.WaitGroup	
	waitGroup.Add(9)

	for i := 0; i < 8; i++{	
		go func(startByte, endByte int) {
			defer waitGroup.Done()
			Multi_chunk(url, startByte, endByte - 1, output)
		}(startByte, endByte)
		startByte += dividedChunks
		endByte += dividedChunks
	}
	
	go Multithread_loop()
	waitGroup.Wait()
}

	


