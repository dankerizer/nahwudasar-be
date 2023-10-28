package helper

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"strings"
	"testing"
)

func PrintJSONTest(title string, t *testing.T, obj interface{}) {
	empJSON, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		//fmt.Println(err.Error())
		t.Errorf(err.Error())
	}
	//fmt.Printf("%s\n%s\n", title, string(empJSON))
	t.Log(fmt.Sprintf("%s\n%s\n", title, string(empJSON)))
}

func PrintJSON(title string, obj interface{}) {
	empJSON, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s\n%s\n", title, string(empJSON))
}

func OnBadRequest(err error) error {
	if strings.HasPrefix(err.Error(), "badrequest") {
		return fiber.ErrBadRequest
	}
	return fiber.ErrInternalServerError
}

func OnError(c *fiber.Ctx, result interface{}, err error) error {
	errMessage := struct {
		Message string
	}{Message: err.Error()}
	var isZero = reflect.ValueOf(result).IsZero()
	if isZero {
		return c.Status(200).JSON(&result)
	}
	return c.Status(500).JSON(&errMessage)
}
