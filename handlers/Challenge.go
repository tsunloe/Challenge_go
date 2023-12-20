package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Challenge1(c *fiber.Ctx) error {

	filePath := "hard.json"
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	var data [][]int
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}

	tmp_data := make([][]int, len(data))

	for i := range tmp_data {
		tmp_data[i] = make([]int, len(data[i]))
		copy(tmp_data[i], data[i])
	}
	for i := 1; i < len(data); i++ {
		for j := range data[i] {
			if j == 0 {
				tmp_data[i][j] += tmp_data[i-1][j]
			} else if j == len(data[i])-1 {
				tmp_data[i][j] += tmp_data[i-1][j-1]
			} else {
				tmp_data[i][j] += findmax(tmp_data[i-1][j-1], tmp_data[i-1][j])
			}
		}
	}

	maxSum := 0
	for _, val := range tmp_data[len(tmp_data)-1] {
		if val > maxSum {
			maxSum = val
		}
	}

	return c.JSON(fiber.Map{"Output": maxSum})
}

func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type ChanlengeRequest2 struct {
	Input string `json:"input" bson:"input"`
}

func Challenge2(c *fiber.Ctx) error {
	nInput := new(ChanlengeRequest2)

	if err := c.BodyParser(nInput); err != nil {
		return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
	}

	if nInput.Input == "" || nInput == nil {
		return c.Status(400).JSON(fiber.Map{"bad input": "Input field cannot be empty"})
	}

	text := strings.ToUpper(nInput.Input)
	pattern := "^[LR=]+$"
	regex := regexp.MustCompile(pattern)

	if !regex.MatchString(text) {
		return c.Status(400).JSON(fiber.Map{"bad input": "Text contains other characters or does not contain only L, R, and ="})
	}

	nums := make([]int, len(text)+1)
	log.Println("output: ", nums)

	tokens := strings.Split(text, "")

	log.Println("tokens: ", tokens)

	for i, token := range tokens {
		if token == "L" {
			if i == 0 {
				nums[i] += 1
			} else {
				log.Println("I: ", i)

				for j := i; j >= 0; j-- {
					log.Println("tokens[j] ", tokens[j])

					if j == 0 || tokens[j] == "R" {
						log.Println("Break ")

						break
					}
					if tokens[j-1] != "R" {

						nums[j-1] += 1
					}
				}
				if nums[i] == 0 {
					nums[i] += 1

				}
			}
		} else if token == "R" {
			nums[i+1] = nums[i] + 1
		} else if token == "=" {
			nums[i+1] = nums[i]
		}

	}

	var str string
	for _, num := range nums {
		str += strconv.Itoa(num)
	}

	return c.JSON(fiber.Map{"Output": str})
}

func Challenge3(c *fiber.Ctx) error {

	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return c.Status(500).SendString("Internal Server Error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return c.Status(500).SendString("Internal Server Error")
	}

	text := string(body)
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, ",", "")

	words := strings.Fields(text)

	beefCount := make(map[string]int)

	for _, word := range words {
		beefCount[word]++
	}

	return c.JSON(fiber.Map{"beef": beefCount})
}
