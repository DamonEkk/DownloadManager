package main


func LoadingBar(loadingBar string, percent float64)string{
	loadingTemplate := "[----------------------------------------------------------------------------------------------------]"
	
	if (loadingBar == ""){
		return loadingTemplate
	}

	runes := []rune(loadingBar)
	if (runes[int(percent)] == '-'){
		runes[int(percent)] = '/'
	}

	

	return string(runes)
}


