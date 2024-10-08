package main

import (
	"net/http"

	"github.com/fgtago/fgweb"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
	"github.com/fgtago/fgweb/midware"
	"github.com/go-chi/chi/v5"
	"github.com/transfashion/gowebeag/apis"
	"github.com/transfashion/gowebeag/handlers"
)

func Router(mux *chi.Mux) error {

	mux.Use(PageSetup)

	// Default handler
	fgweb.Get(mux, "/favicon.ico", defaulthandlers.FaviconHandler)
	fgweb.Get(mux, "/asset/*", defaulthandlers.AssetHandler)
	fgweb.Get(mux, "/template/*", defaulthandlers.TemplateHandler)

	hnd := handlers.New(appsmodel.GetWebservice())
	fgweb.Get(mux, "/manifest.json", hnd.Manifest)
	fgweb.Get(mux, "/", hnd.Home)

	fgweb.Get(mux, "/catalog/{brand}/{gender}/{category}", hnd.Catalog)
	fgweb.Get(mux, "/catalog/{brand}/{gender}", hnd.Catalog)
	fgweb.Get(mux, "/catalog/{brand}", hnd.Catalog)
	//fgweb.Get(mux, "/catalog", hnd.Catalog)

	fgweb.Get(mux, "/content/index", hnd.ContentIndex)
	fgweb.Get(mux, "/content/{id}", hnd.Content)
	fgweb.Get(mux, "/collection/{tag}", hnd.Collection)

	api := apis.New(appsmodel.GetWebservice())
	fgweb.Post(mux, "/api/promoter", api.Promoter)
	fgweb.Post(mux, "/api/search", api.Search)

	return nil
}

func PageSetup(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if midware.IsAsset(r.URL.Path) || midware.IsTemplate(r.URL.Path) {
			next.ServeHTTP(w, r)
		} else {
			ctx := r.Context()
			pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)

			var base_href string
			if r.Header.Get("Base_href") != "" {
				base_href = r.Header.Get("Base_href")
			} else {
				base_href = "/"
			}

			pv.Setup = &handlers.PageSetup{
				BaseUrl:        base_href,
				ShowHeader:     true,
				ShowFooter:     true,
				ShowFooterRow3: false,
			}
			next.ServeHTTP(w, r)
		}

	})
}
