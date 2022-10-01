package main
import("html/template"; "net/http"; "github.com/gorilla/mux";)

func instance(w http.ResponseWriter, r *http.Request) {
	formtype:="instance";
	form:=&Form{tmpl: template.Must(template.ParseFiles("forms/"+formtype+".html"))};
	id:=mux.Vars(r)["id"]
	formData:=Instance{Id: id}
	switch{
		case r.Method == http.MethodGet: {
			formData.Edit=true
			readUnmarshal("defns/"+formtype+"/"+id+".json", &formData)
		}
		case r.Method == http.MethodPost: {
			formData.Edit=false
			formData.Scope=		r.FormValue("scope")
			formData.Type=		r.FormValue("type")
			formData.Definition=	r.FormValue("definition")
			formData.Campaign=	r.FormValue("campaign")
			formData.Notes=		r.FormValue("notes")
			marshalWrite(&formData, "defns/"+formtype+"/"+id+".json")
		}
	}
	form.tmpl.Execute(w, formData);
}

