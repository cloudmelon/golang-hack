package loader

// BooksLiteral is a slice literal of Bookdata struct pointers, containing a subset of the real book data
// Once you have created your CSV loader, you will not need this file.
import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func DataReader() *[]*BookData {
	file, err := os.Open("./assets/books.csv")
	r := csv.NewReader(file)
	ret := make([]*BookData, 0, 0)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	for {
		rec, err := r.Read()

		if err == io.EOF {
			log.Println("End of File")
			break
		} else if err != nil {
			log.Println(err)
			break
		}

		AverageRating, _ := strconv.ParseFloat(rec[3], 64)
		NumPages, _ := strconv.Atoi(rec[7])
		Ratings, _ := strconv.Atoi(rec[8])
		Reviews, _ := strconv.Atoi(rec[9])

		book := &BookData{
			BookID:        rec[0],
			Title:         rec[1],
			Authors:       rec[2],
			AverageRating: AverageRating,
			ISBN:          rec[4],
			ISBN13:        rec[5],
			LanguageCode:  rec[6],
			NumPages:      NumPages,
			Ratings:       Ratings,
			Reviews:       Reviews,
		}

		fmt.Println("Our book is ==>", book)
		ret = append(ret, book)
		fmt.Println(rec)
	}

	//fmt.Println(r)
	fmt.Println(err)
	return &ret

}
