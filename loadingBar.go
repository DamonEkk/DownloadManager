package main

import "fmt"

type ProgressBar struct {
	loading string
	maxValue int
	currentValue int 

}

func LoadingBar(loadingBar string, percent float64)string{
	loadingTemplate := "[----------------------------------------------------------------------------------------------------]"
	
	if (loadingBar == ""){
		return loadingTemplate
	}

	runes := []rune(loadingBar)
	if (runes[int(percent * 100)] == '-'){
		runes[int(percent * 100)] = '/'
	}
	return string(runes)
}


func Set_progress(bar *ProgressBar, completedBytes int){
	bar.currentValue += completedBytes 
	return
}

func Display_progress(bar *ProgressBar){
	runes := []rune(bar.loading)

	percent := bar.currentValue / bar.maxValue

	if (runes[int(percent * 100)] == '-'){
		runes[int(percent * 100)] = '/'
	}

	fmt.Printf("\r")
	return
}

// Need to edit to have a cond broadcast to avoid endless spinning
func Multithread_loop(bar *ProgressBar){
	delta := 0.00
	var currentPercent float64

	for {
		currentPercent = float64(bar.currentValue) / float64(bar.maxValue)

		if (float64(bar.currentValue) > delta){
			Display_progress(bar)
			delta = currentPercent + 0.02
		}
	}
	return
}

func Create_progresbar(maxBytes int)*ProgressBar{
	bar := &ProgressBar{}
	bar.maxValue = maxBytes
	bar.loading = "[----------------------------------------------------------------------------------------------------]"
	bar.currentValue = 0
	return bar
}


