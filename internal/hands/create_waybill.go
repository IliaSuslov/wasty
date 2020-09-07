package hands

//type FormGroup struct {
//	Label       string
//	Name        string
//	Type        string
//	Desc        string
//	Placeholder string
//}
//type CreateWaybillFormVals struct {
//	Title      string
//	FormGroups []FormGroup
//}
//
//func CreateWaybillForm(t *template.Template) func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		log.Println("CreateWaybillForm")
//		formGroups := []FormGroup{
//			FormGroup{
//				Name:        "Date",
//				Label:       "Дата",
//				Type:        "text",
//				Desc:        "дата в формате 2020-09-12",
//				Placeholder: "2020-09-12",
//			},
//			FormGroup{
//				Name:        "Desc",
//				Label:       "Описание",
//				Type:        "text",
//				Desc:        "",
//				Placeholder: "дата в формате 2020-09-12",
//			},
//		}
//		err := t.ExecuteTemplate(w, "create_waybill.html", CreateWaybillFormVals{
//			Title:      "Новый путевой лист",
//			FormGroups: formGroups,
//		})
//		if err != nil {
//			log.Println(err)
//		}
//	}
//}
//
//func CreateWaybill(DB *mongo.Database, t *template.Template) func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		log.Println("CreateWaybill")
//		//w.WriteHeader(200)
//		err := t.ExecuteTemplate(w, "index.html", struct{ Title string }{"HOME"})
//		if err != nil {
//			log.Println(err)
//		}
//	}
//}
