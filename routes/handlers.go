package routes

import (
    "net/http"
)


func IndexHandler(w http.ResponseWriter, r *http.Request) error {
    page := PublicPageData{
        Page: "index.html",
        Title: "ChiChi",
        Content: "Sample Content",
        CSS: CSS_URL,
    }

    page.RenderHTMLTemplate(w, page)
    return nil
}