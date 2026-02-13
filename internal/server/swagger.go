package server

import (
	"net/http"
	"titiktopup-core/constant"
)

func serveSwaggerJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	http.ServeFile(w, r, constant.DefaultSwaggerJSON)
}

func serveSwaggerUI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	const html = `<!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="utf-8" />
            <title>TitikTopup API Docs</title>
            <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
        </head>
        <body>
            <div id="swagger-ui"></div>
            <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
            <script>
                window.onload = () => {
                    window.ui = SwaggerUIBundle({
                        url: '/swagger.json',
                        dom_id: '#swagger-ui',
                    });
                };
            </script>
        </body>
        </html>`
	_, _ = w.Write([]byte(html))
}
