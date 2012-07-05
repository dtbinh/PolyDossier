package adapters


func (*DefaultAdapter) Run (r *http.Request) {
	resp, err := goResponse(r)
	parser = make.HTMLParser(resp)
	
	// TODO: put here what we exactly need
	the_first_value := parser.getFeild("MAP OF FEILDS HERE")
	the_second_value := parser.getFeild("MAP OF FEILDS HERE")
}