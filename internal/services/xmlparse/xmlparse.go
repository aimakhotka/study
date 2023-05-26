package xmlparse

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Info struct {
	BookTitle  string `xml:"book-title"`
	Annotation string `xml:"annotation"`
}

type Chapter struct {
	XMLName xml.Name `xml:"chapter"`
	ChTitle string   `xml:"title"`
	Body    string   `xml:"section"`
}

type Book struct {
	XMLName  xml.Name `xml:"book"`
	Chapters []Chapter
}

//Books := make(map[string]int) // list of booktitle

// функция открытие файла
// парсинг в структуры
// запись в бд
// все в го рутинах

func XmlParsing() {
	var file string
	file = "files/Whale-52.fb2"
	xmlFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened" + file)

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var book Book
	xml.Unmarshal(byteValue, &book)

	for i := 0; i < len(book.Chapters); i++ {
		fmt.Println(book.Chapters[i].ChTitle)
	}
}
