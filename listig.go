package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var sources = [2]string{"DockerPS", "ExampleList"}

type DockerPs struct {
	ContainerId string `json:"CONTAINER ID"`
	Image       string
	Command     string
	Created     string
	Ports       string
	Names       string
}

type ExampleListItem struct {
	Id       string
	Title    string
	Reminder string
	Priority string
}

type GenericList struct {
	Title string
	Items []string
}

func getDockerPsStatusList(w http.ResponseWriter, req *http.Request) {
	var dockerPs []DockerPs
	dat, err := os.ReadFile("./result.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(dat, &dockerPs)

	for i := 0; i < len(dockerPs); i++ {
		fmt.Println(dockerPs[i].Image)
	}

	result, _ := json.Marshal(dockerPs)

	fmt.Fprint(w, string(result))
}

func getExampleList(w http.ResponseWriter, req *http.Request) {
	var exampleList [1]ExampleListItem

	exampleList[0].Title = "Segelfartyg"
	exampleList[0].Id = "x"
	exampleList[0].Reminder = "x"
	exampleList[0].Priority = "x"

	result, _ := json.Marshal(exampleList)

	fmt.Fprint(w, string(result))
}

func main() {

	http.HandleFunc("/status", getGenericList)
	http.HandleFunc("/status/docker", getDockerPsStatusList)
	http.HandleFunc("/status/example", getExampleList)
	http.ListenAndServe(":9000", nil)
}

func getGenericList(w http.ResponseWriter, req *http.Request) {
	var genericList [2]GenericList

	// GETTING DOCKER PS LIST
	var dockerPs []DockerPs
	dat, err := os.ReadFile("./result.json")
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(dat, &dockerPs)
	genericList[0].Title = sources[0]
	for i := 0; i < len(dockerPs); i++ {
		res := append(genericList[0].Items, dockerPs[i].ContainerId+" "+dockerPs[i].Image+" "+dockerPs[i].Ports+" "+dockerPs[i].Names+" "+dockerPs[i].Command+" "+dockerPs[i].Created)
		genericList[0].Items = append(genericList[0].Items, res...)
	}

	genericList[1].Title = sources[1]

	// GETTING EXAMPLE LIST
	var exampleList [1]ExampleListItem

	exampleList[0].Title = "Segelfartyg"
	exampleList[0].Id = "x"
	exampleList[0].Reminder = "x"
	exampleList[0].Priority = "x"

	genericList[1].Title = exampleList[0].Title

	result, _ := json.Marshal(genericList)

	fmt.Fprint(w, string(result))
}
