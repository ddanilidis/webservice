package main

import (
	"fmt"
	"github.com/ddanilidis/webservice/model"
)

func main() {
	u := model.User{
		ID:        2,
		FirstName: "Vasya",
		LastName:  "Pupkin",
	}
	fmt.Println(u)
}
