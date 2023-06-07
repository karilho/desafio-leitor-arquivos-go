package utils

import (
	"fmt"
	"github.com/mozillazg/go-unidecode"
	"strings"
)

func NameFormater(name string) string {
	name = strings.ReplaceAll(name, " ", "")
	//name = strings.ToUpper(name)
	name = unidecode.Unidecode(name)
	name = unidecode.Unidecode(name)
	fmt.Println(name + ".txt")
	return name + ".txt"
}
