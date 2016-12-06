package session

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Session struct {
	Secret     string `json:"secret"`
	UserId     string `json:"userId"`
	Type       string `json:"type"`
	PartnerId  string `json:"partnerId"`
	Expiry     string `json:"expiry"`
	Privileges string `json:"privileges"`
	Format     string `json:"format"`
}

func New(sec, usr, typ, par, exp, priv, form string) string {

	s := &Session{
		Secret:     sec,
		UserId:     usr,
		Type:       typ,
		PartnerId:  par,
		Expiry:     exp,
		Privileges: priv,
		Format:     form,
	}

	fmt.Println(ServiceUrl)

	if ServiceUrl == "" {
		ServiceUrl = SetUrl("")
	}

	sj, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	r, err := http.NewRequest("POST", ServiceUrl, bytes.NewBuffer(sj))
	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

	return string(body)
}

var (
	ServiceUrl string
)

func SetUrl(url string) string {
	u := "http://api.kaltura.com/api_v3/?service=session&action=start"
	if url != "" {
		u = url + "/api_v3/?service=session&action=start"
	}

	ServiceUrl = u

	return u
}
