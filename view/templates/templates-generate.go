package templates

import "fmt"

func GenerateHTMLFilesList(files []string) string {
	html := `
	<html>
	<head>
	<title>File Serve HTTP</title>
	<title>Files list</title>
	</head>
	<body>
	<ul>`

	for _, file := range files {
		html += fmt.Sprintf("<li><a href=\"/files/download?filename=%s\">%s</a></li>\n", file, file)
	}

	html += `
	</ul>
	<form action="/files/download" method="get">
	<label for="filenameForDownload">Input filename:</label>
	<input type="text" id="filenameForDownload" name="filename">
	<button type="submit">Download</button>
	</form>
	</body>
	</html>`

	return html
}
