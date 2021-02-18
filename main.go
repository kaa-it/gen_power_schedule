package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func test() {
	ds := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)

	fmt.Println(strconv.Quote(ds.Format(time.RFC3339)))

	dd := time.Date(2021, time.January, 1, 2, 3, 4, 0, time.Local)

	fmt.Println(strconv.Quote(dd.Format(time.RFC3339)))

	tt, err := time.Parse(time.RFC3339, "2020-01-02T00:00:00+03:00")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tt)
	}

	y, m, d := tt.UTC().Date()

	_ = time.Date(y, m, d, 0, 0, 0, 0, time.UTC).Unix()
}

func main() {
	f, err := os.Create("2021.csv")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}

	defer f.Close()

	ds := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)

	end := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()

	hs := 6
	ms := 25

	for ds.Unix() < end {
		ts := time.Date(ds.Year(), ds.Month(), ds.Day(), hs, ms, 0, 0, time.UTC)
		te := ts.Add(time.Hour * 8)

		fmt.Fprintf(f, "%s,%s,%s\n",
			ds.Format("01-02"),
			ts.Format("15:04"),
			te.Format("15:04"),
		)

		ds = ds.AddDate(0, 0, 1)
		ms++
	}

}
