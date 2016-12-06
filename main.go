package main

import (
	"fmt"

	"github.com/ivancduran/kaltura/config"

	"github.com/ivancduran/kaltura/session"
)

func main() {

	c := config.New()

	session.SetUrl(c.Url)

	s := session.New(c.Token, c.Email, "KalturaSessionType::ADMIN", c.Id, c.Expire, "", c.Format)
	fmt.Println(s)

}
