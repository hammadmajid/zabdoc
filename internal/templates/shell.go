package templates

import (
	"bytes"
	"html/template"
	"os"
)

var shell = *template.Must(template.New("header").Parse(`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>zabdoc</title>
    <link rel="stylesheet" href="/css/terminal.css"/>
    <meta name="viewport" content="width=device-width, initial-scale=1"/>
    <link rel="icon" href="/static/favicon.svg"/>
</head>

<body class="container-fluid">
<header class="container">
    <div class="terminal-nav">
        <div class="terminal-logo">
            <a href="/" class="terminal-prompt">zabdoc</a>
        </div>
        <nav class="terminal-menu">
            <ul>
                <li><a href="/about">About</a></li>
            </ul>
        </nav>
    </div>
</header>
<main class="container">
    {{ .Content }}
</main>

<footer class="container footer">
    <p>
        <a href="https://github.com/hammadmajid/zabdoc" target="_blank" rel="noopener noreferrer">GitHub</a> |
        Licensed under <a href="https://www.gnu.org/licenses/gpl-3.0.html" target="_blank" rel="noopener noreferrer">GPL
        v3</a>
    </p>
</footer>
</body>
</html>
`))

func renderWithShell(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = shell.Execute(&buf, struct{ Content template.HTML }{
		Content: template.HTML(content),
	})
	if err != nil {
		panic(err)
	}

	return buf.String()
}
