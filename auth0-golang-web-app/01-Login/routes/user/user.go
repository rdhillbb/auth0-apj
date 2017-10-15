package user

import (
	"fmt"
	"github.com/rdhillbb/auth0-apj/auth0-golang-web-app/01-Login/app"
	templates "github.com/rdhillbb/auth0-apj/auth0-golang-web-app/01-Login/routes"
	"github.com/rdhillbb/apj911"
	"github.com/rdhillbb/createTxT2Spch"
	"github.com/rdhillbb/buildWatsonNumber"
	"time"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Method == "POST" {

		r.ParseForm()
		fmt.Println("Post-------------->")
		// logic part of log in
		fmt.Println("Company Name:", r.Form["companyName"])
		fmt.Println("Case Number:", r.Form["caseNumber"])
		fmt.Println("Case Priority:", r.Form["casePriority"])
		companyName := r.Form["companyName"][0]
		caseNumber := r.Form["caseNumber"][0]
		message := buildWatsonNumber.CrMsg4wat(companyName, caseNumber)
		start := time.Now()
		createTxT2Spch.CreateWatAudio(apj911.ServiceConfig.Video_dir, apj911.ServiceConfig.Audio_file, message, apj911.AuthTokens.WatsonToken, apj911.AuthTokens.WatsonPass)
		fmt.Println("Broadcast:000 Wattson: ", time.Since(start))
		apj911.Call(w, r, "Alert P1 "+companyName+" Case Number: "+caseNumber+" Please contact the TAC Duty Manager.")
	}

	fmt.Println(session.Values["profile"])

	templates.RenderTemplate(w, "user", session.Values["profile"])
}
