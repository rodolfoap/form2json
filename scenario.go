package main
import("html/template"; "net/http"; "github.com/gorilla/mux";)

func scenario(w http.ResponseWriter, r *http.Request) {
	formtype:="scenario";
	form:=&Form{tmpl: template.Must(template.ParseFiles("forms/"+formtype+".html"))};
	id:=mux.Vars(r)["id"]
	formData:=Scenario{Id: id}
	switch{
		case r.Method == http.MethodGet: {
			formData.Edit=true
			readUnmarshal("defns/"+formtype+"/"+id+".json", &formData)
		}
		case r.Method == http.MethodPost: {
			formData.Edit=false
			formData.Constraint=r.FormValue("constraint")
			marshalWrite(&formData, "defns/"+formtype+"/"+id+".json")
		}
	}
	form.tmpl.Execute(w, formData);
}
