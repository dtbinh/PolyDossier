package pages

type Credentials struct {
  Id		string
	Pass	string
	Birth	string
}

// func goHandler(w http.ResponseWriter, r *http.Request) {
  // log.Println("Lolsaure")
  // //file.Write([]byte("lolsaure\n"))

	// file, err := os.Create(fmt.Sprintf("[%s] - %s", kLogFile, time.Now().Unix(), ".log"))
	// defer file.Close()
	// if err != nil { log.Fatal(err) }
	
	// t := time.Now()
	// resp, err := goResponse(r)

	// if err != nil {
		// fmt.Println(r.Method, err)
		// return
	// }

  // z := html.NewTokenizer(resp.Body)

	// for {
		// tt := z.Next()
		// if tt == html.ErrorToken {
			// fmt.Println("Error token")
			// break
		// } else {
		// fmt.Println(z.Token())
		// }
		
	// }
	// file.Write([]byte("lolsaure"))
	
	// //decoder := xml.NewDecoder(resp.Body)
	// //token, tErr := xml.Token(nil), error(nil)
	// //token, tErr = decoder.Token()
	// //for i := 0; tErr == nil; i++ {
		// // if i == 1 {
		// //	decoder.Skip()
		// //} else {*/
		// //	fmt.Println(token, tErr)
	// //		token, tErr = decoder.Token()
		// //}
	// //}

	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
		// fmt.Println(r.Method, err)
		// return
	// }

	// w.Write(data)

	// resp.Body.Close()

	// elapsed := time.Since(t)

	// if elapsed < 200*time.Millisecond {
		// fmt.Println(time.Since(t).String() + "\t" + r.URL.Path)
	// } else {
		// fmt.Println(time.Since(t).String() + "\t" + r.URL.Path + "\t[Trop Long!!!]")
	// }
// }

// func goResponse(r *http.Request) (*http.Response, error) {
	// defaultClient := &http.Client{}

	// // Si on arrive sur la page.
	// if r.URL.Path == "/" {
		// return defaultClient.Get(kPolyHost + "/poly/poly.html")
	// }

	// switch r.Method {
	// case "GET":
		// return defaultClient.Get(kPolyHost + r.URL.Path)
	// case "POST":
		// r.ParseForm()
		// return defaultClient.PostForm(kPolyHost+r.URL.Path, r.Form)
	// }
	// return nil, nil
// }