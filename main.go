
package main

import (
    "fmt"
    "net/http"
    
	"appengine"
	//"appengine/datastore"
	"appengine/user"
	
	"appengine/blobstore"
)


// Golang BUILD output
/*
func main() {
	fmt.Printf("Boomcase Console App Output\n")
}
*/


// ========== ========== ========== ========== ==========
// GAE config handlers for URIs
func init() {
	http.HandleFunc("/savecasedriver",	handlerSaveCaseDriver)
	http.HandleFunc("/savecase",		handlerSaveCase)
	http.HandleFunc("/savedriver",		handlerSaveDriver)
	
	http.HandleFunc("/serve/",			handlerServe)
	
    http.HandleFunc("/",				handlerRoot)
}
// ========== ========== ========== ========== ==========


// ========== ========== ========== ========== ==========
// GAE handler - handlerSaveCaseDriver
func handlerSaveCaseDriver(w http.ResponseWriter, r *http.Request) {
	output := ""
	
	output += "<h1>handlerSaveCaseDriver()</h1>"
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	ctx := appengine.NewContext(r) // c or ctx
	// Can send to func via: (c context.Context)
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	// ========== ========== ========== ========== ==========
	// [START if_user]
	if u := user.Current(ctx); u != nil {
		//g.Author = u.String()
	}
	// [END if_user]
	// ========== ========== ========== ========== ==========
	
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    
    output += "<h1>START: saveCaseDriverBlobstore()</h1>"
    //blobkey := saveCaseDriverBlobstore(r)
    blobkey := ""
    
    output += "<h1>START: saveCaseDriverDatastore()</h1>"
    output += saveCaseDriverDatastore(r, ctx, blobkey)
    
    fmt.Fprint(w, output)
}
// ========== ========== ========== ========== ==========


// ========== ========== ========== ========== ==========
// GAE handler - handlerSaveCase
func handlerSaveCase(w http.ResponseWriter, r *http.Request) {
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	ctx := appengine.NewContext(r) // c or ctx
	// Send to func via: (c context.Context)
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	// ========== ========== ========== ========== ==========
	// [START if_user]
	if u := user.Current(ctx); u != nil {
		//g.Author = u.String()
	}
	// [END if_user]
	// ========== ========== ========== ========== ==========
	
	/*
	htmlContent := `
		<h1>Boomcase splash screen.</h1>
		<h2><a href="customize">Boomcase Customize</a></h2>
		<h2><a href="boombarrel">Boom Barrel</a></h2>
	`
	*/
    //fmt.Fprint(w, drawPage(r.URL.Path[1:]))
    w.Header().Set("Content-Type", "text/html")
    //fmt.Fprint(w, saveCase(r, ctx))
    fmt.Fprint(w, saveCase(r))
}
// ========== ========== ========== ========== ==========


// ========== ========== ========== ========== ==========
// GAE handler - handlerSaveDriver
func handlerSaveDriver(w http.ResponseWriter, r *http.Request) {
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	ctx := appengine.NewContext(r) // c or ctx
	// Send to func via: (c context.Context)
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	// ========== ========== ========== ========== ==========
	// [START if_user]
	if u := user.Current(ctx); u != nil {
		//g.Author = u.String()
	}
	// [END if_user]
	// ========== ========== ========== ========== ==========
	
	/*
	htmlContent := `
		<h1>Boomcase splash screen.</h1>
		<h2><a href="customize">Boomcase Customize</a></h2>
		<h2><a href="boombarrel">Boom Barrel</a></h2>
	`
	*/
    //fmt.Fprint(w, drawPage(r.URL.Path[1:]))
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, saveDriver(r, ctx))
}
// ========== ========== ========== ========== ==========


// ========== ========== ========== ========== ==========
func handlerServe(w http.ResponseWriter, r *http.Request) {
	blobstore.Send(w, appengine.BlobKey(r.FormValue("blobKey")))
}
// ========== ========== ========== ========== ==========


// ========== ========== ========== ========== ==========
// GAE handler - handlerRoot
func handlerRoot(w http.ResponseWriter, r *http.Request) {
	
	
	// ========== ========== ========== ========== ==========
	// New Context - opaque value used by many functions in the Go App Engine SDK to communicate with the App Engine service
	// [START new_context]
	ctx := appengine.NewContext(r) // c or ctx
	// Send to func via: (c context.Context)
	// [END new_context]
	// ========== ========== ========== ========== ==========
	
	
	// ========== ========== ========== ========== ==========
	// [START if_user]
	if u := user.Current(ctx); u != nil {
		//g.Author = u.String()
	}
	// [END if_user]
	// ========== ========== ========== ========== ==========
	
	
	/*
	htmlContent := `
		<h1>Boomcase splash screen.</h1>
		<h2><a href="customize">Boomcase Customize</a></h2>
		<h2><a href="boombarrel">Boom Barrel</a></h2>
	`
	*/
    //fmt.Fprint(w, drawPage(r.URL.Path[1:]))
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, drawPage(r, ctx))
}
// ========== ========== ========== ========== ==========

