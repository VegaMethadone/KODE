package yandex_test

import (
	"fmt"
	"kode/internal/yandex"
	"testing"
)

func TestSpeller(t *testing.T) {
	str := "Масква. Вайна. Свабода. Физика. Матиматика"
	result, err := yandex.ValidateBody(str)
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(str)
	fmt.Println(string(result))

}
