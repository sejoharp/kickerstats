package interfaces

import (
	"bitbucket.org/joscha/hpfeed/helper"
	"io/ioutil"
	"net/http"
	"net/url"
)

func DownloadSeason(id string) []byte {
	params := url.Values{"filter_saison_id": []string{id}, "task": []string{"veranstaltungen"}}
	resp, err := http.PostForm("http://www.kickern-hamburg.de/liga-tool/mannschaftswettbewerbe", params)
	helper.HandleFatalError("error getting season: ", err)
	body, err := ioutil.ReadAll(resp.Body)
	helper.HandleFatalError("error reading season data: ", err)
	resp.Body.Close()
	return body
}

func DownloadHTMLData(url string) []byte {
	resp, err := http.Get(url)
	helper.HandleFatalError("error getting html: ", err)
	body, err := ioutil.ReadAll(resp.Body)
	helper.HandleFatalError("error reading html data: ", err)
	resp.Body.Close()
	return body
}
