package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

type PushEvent struct {
	EventName string            `json:"event_name"`
	Project   map[string]string `json:"project"`
}

func Webhook(c *fiber.Ctx) error {
	p := new(PushEvent)
	c.BodyParser(p)
	fmt.Println(p.EventName, p.Project)

	if p.EventName == "push" {
		fmt.Print(p.Project["name"])
		jsonData := GetJsonData()
		out, err := exec.Command(jsonData[p.Project["name"]].(string)).Output()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(out))
	}

	return c.SendString("done")
}

func GetJsonData() map[string]interface{} {
	jsonFile, err := os.Open("./api/projects.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}
