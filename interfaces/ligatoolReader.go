package interfaces

/*
func GetData(client *http.Client) {
	params := url.Values{"username": []string{this.forumUser}, "password": []string{this.forumPasswd}, "login": {"anmelden"}}
	resp, err := client.PostForm("http://kickern-hamburg.de/phpBB2/login.php", params)
	if err != nil {
		fmt.Println(err)
	}
	resp.Body.Close()
}

func (this *ForumReader) getHTMLData(client *http.Client) []byte {
	resp, err := client.Get("http://kickern-hamburg.de/phpBB2/viewforum.php?f=15")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return body
}


resp, err := http.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form",
	url.Values{"key": {"Value"}, "id": {"123"}})
...
resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// ...
*/
