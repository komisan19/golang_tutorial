package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	lines := []string{
		"りんご、Appplce、バラ",
		"みかん、Orange、ミカン",
		"すいか、Watermelon, ウリ",
	}
	csvStr := strings.Join(lines, "\n")

	r := csv.newReader(strings.NewReader(csvStr))
	for {
		record,err := r.Read()
		if err == io.EOF{
			break
		}
		if err != nil {
			// 読み込みエラー
			fmt.Println("Read error: ", err)
			break
		}

		fmt.Printf("日本語[%s], 英語[%s] 化学[%s]\n", record[0], record[1], record[2])
	}
}
