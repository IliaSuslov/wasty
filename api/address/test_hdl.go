
package address

import (
	gcontext "github.com/gorilla/context"
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
	"net/http"
)

func TestHDL(Col *mongo.Collection,
	IsAllow func(User interface{}) bool,
	Denied func(w http.ResponseWriter),
	None func(w http.ResponseWriter),
	OnError func(http.ResponseWriter, error) bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := gcontext.Get(r, "user")
		if !IsAllow(user) {
			Denied(w)
			return
		}

		t, err:= template.New("html").
			Parse(CreateFormTemplate)

		if OnError(w, err) {
			return
		}
		w.Header().Set("ContentType", "text/html; charset=utf-8")
		err = t.Execute(w, nil)
		if OnError(w, err) {
			return
		}
	}
}


var CreateFormTemplate = `
<html>
<head>
</head>
<body>
	<h4>Test Create</h4>
	<p>test desc</p>	
	<textarea id="json">{"name":"test"}</textarea>
	<button onclick="sendJSON()">Send JSON</button> 
<script>
function sendJSON(){
	let json = document.querySelector('#json');
	fetch("/api/v1/addrs", {
		method: 'POST',
        headers: {
		  'Content-Type': 'application/json'
		  // 'Content-Type': 'application/x-www-form-urlencoded',
		},
		body: json.value,
	})
}

</script> 
</body>
</html>
`
