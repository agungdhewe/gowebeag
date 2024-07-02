package handlers

import (
	"fmt"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/fgtago/fgweb/defaulthandlers"
)

func (hdr *Handler) Home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pv := ctx.Value(appsmodel.PageVariableKeyName).(*appsmodel.PageVariable)
	pv.PageName = "home"
	pv.Use(hdr.LoginCheck)

	setup := pv.Setup.(*PageSetup)
	fmt.Println(setup)
	setup.ShowFooterRow4 = true

	defaulthandlers.SimplePageHandler(pv, w, r)
}
