package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application)home(w http.ResponseWriter,r *http.Request){
	if r.URL.Path!="/"{
		app.notFound(w)
		return
	}

	files:=[]string{"./ui/html/base.tmpl","./ui/html/partials/nav.tmpl","./ui/html/pages/home.tmpl"}


	ts,err:=template.ParseFiles(files...)
	if err!=nil{
		app.serverError(w,err)
		http.Error(w,"Internal Server Error",500)
		return
	}

	err=ts.ExecuteTemplate(w,"base",nil)

	if err!=nil{
		app.serverError(w,err)
		http.Error(w,"Internal Server Error",500)
	}

	// w.Write([]byte("Hello from Server"))
}

func (app *application)snippetCreate(w http.ResponseWriter, r *http.Request){
	if r.Method !=http.MethodPost{
		w.Header().Set("Allow","POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		//using http.Error()
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		app.clientError(w,http.StatusMethodNotAllowed)
		return
	}
	//deleting default response header
	// w.Header()["Date"]=nil
	w.Write([]byte("Create a new Snippet"))
}

func (app *application)snippetView(w http.ResponseWriter, r *http.Request){
	id,err:=strconv.Atoi(r.URL.Query().Get("id"))
	if err!=nil || id<1{
		app.notFound(w)
		return
	}
	// w.Write([]byte("Display a specific Snippet"))
	fmt.Fprintf(w,"Display a specific snippet with ID %d...",id)
}
