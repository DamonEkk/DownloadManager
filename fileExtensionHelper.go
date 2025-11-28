package main



func ConvertType(contentType string) string {
	fileTypes := map[string]string{
		// applications
		"application/java-archive": ".jar",
		"application/EDI-X12": ".x12",
		"application/EDIFACT": ".edi",
		"application/javascript (obsolete)": ".js",
		"application/octet-stream": ".bin",
		"application/ogg": ".ogg",
		"application/pdf": ".pdf",
		"application/xhtml+xml": ".xhtml",
		"application/x-shockwave-flash": ".swf",
		"application/json": ".json",
		"application/ld+json": ".jsonld",
		"application/xml": ".xml",
		"application/zip": ".zip",
		"application/x-www-form-urlencoded": ".urlencoded",
		"application/x-msdownload": ".exe",

		// audio
		"audio/mpeg": ".mp3",
		"audio/x-ms-wma": ".wma",
		"audio/vnd.rn-realaudio": ".ra",
		"audio/x-wav": ".wav",

		// image
		"image/gif": ".gif",
		"image/jpeg": ".jpg",
		"image/png": ".png",
		"image/tiff": ".tiff",
		"image/vnd.microsoft.icon": ".ico",
		"image/x-icon": ".ico",
		"image/vnd.djvu": ".djvu",
		"image/svg+xml": ".svg",

		// text
		"text/css": ".css",
		"text/csv": ".csv",
		"text/event-stream": ".event",
		"text/html": ".html",
		"text/javascript": ".js",
		"text/plain": ".txt",
		"text/xml": ".xml",

		// video
		"video/mpeg": ".mpeg",
		"video/mp4": ".mp4",
		"video/quicktime": ".mov",
		"video/x-ms-wmv": ".wmv",
		"video/x-msvideo": ".avi",
		"video/x-flv": ".flv",
		"video/webm": ".webm"	}

	if ext, ok:= fileTypes[contentType]; ok{
		return ext
	} else{
		return "Err"
	}	
		
}
