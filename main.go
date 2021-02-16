package main

import (
	"fmt"
	"os"
	"time"
)

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

		fmt.Printf("%s,%s,%s\n",
			ds.Format("01-02"),
			ts.Format("15:04"),
			te.Format("15:04"),
		)

		ds = ds.AddDate(0, 0, 1)
		ms++
	}
}
