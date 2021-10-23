package gaws

//func TestClient_ListObjects(t *testing.T) {
//	bucket := os.Getenv("BUCKET")
//
//	pkg, err := New(&bucket)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	ctx := context.Background()
//
//	files, err := pkg.listObjects(ctx)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//
//	jsonResult, _ := json.Marshal(files)
//	log.Println(string(jsonResult))
//}
