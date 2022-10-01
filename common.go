package main
import("net/http"; "encoding/json"; "io/ioutil"; "github.com/gorilla/mux";)

func readUnmarshal(filename string, formData interface{}){
	formJson, _ := ioutil.ReadFile(filename)
	_ = json.Unmarshal([]byte(formJson), formData) // Fills only saved fields, Id and Edit are not saved, so, not modified
}

func marshalWrite(formData interface{}, filename string){
	formJson, _:=json.MarshalIndent(formData, "", " ")
	_ = ioutil.WriteFile(filename, formJson, 0644)
}

func main(){
	router:=mux.NewRouter()
	router.HandleFunc("/instance/{id}", instance)
	router.HandleFunc("/scenario/{id}", scenario)
	http.ListenAndServe(":8888", router)
}
