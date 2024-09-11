package yandex

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"unicode"
)

type Speller struct {
	Code        int      `json:"code"`
	Pos         int      `json:"pos"`
	Row         int      `json:"row"`
	Col         int      `json:"col"`
	Len         int      `json:"len"`
	Word        string   `json:"word"`
	Suggestions []string `json:"s"`
}

type pair struct {
	index       int
	punctuation string
}

// мясная функция по волидации текста.  Главная проблема - берет 1  преложенное слово спеллером. Нужно подключать ML для корретной работы
func ValidateBody(body string) ([]byte, error) {
	arr := strings.Split(body, " ")
	m := make(map[string]pair, len(arr))

	for index, word := range arr {
		newPair := new(pair)
		newPair.index = index

		newWord := []rune(word)

		punc := checkPunctuation(word)
		if punc != "" {
			newPair.punctuation = punc
			newWord = newWord[:len(newWord)-1]
		}

		m[string(newWord)] = *newPair
	}

	url := "https://speller.yandex.net/services/spellservice.json/checkText?text="
	newArr := []string{}
	for key := range m {
		newArr = append(newArr, key)
	}
	url = url + strings.Join(newArr, "+")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Неверный запрос к яндексу")
		return nil, err
	}
	defer resp.Body.Close()

	requestBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Не смог прочитать тело запроса")
		return nil, err
	}

	var jsonData []Speller
	err = json.Unmarshal(requestBody, &jsonData)
	if err != nil {
		fmt.Println("Ошибка при анмаршалинге:", err)
		return nil, err
	}

	for _, word := range jsonData {
		old := word.Word
		newWord := word.Suggestions[0]

		prevValue := m[old]
		arr[prevValue.index] = newWord + prevValue.punctuation
	}

	return []byte(strings.Join(arr, " ")), nil
}

func checkPunctuation(word string) string {
	for _, value := range word {
		if unicode.IsPunct(value) {
			return string(value)
		}
	}
	return ""
}
