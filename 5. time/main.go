/*
	TIME BASIC
*/
package main

import (
	"fmt"
	"time"
)

func main(){
	//now
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)

	//date
	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
    fmt.Printf("time2 %v\n", time2)
	fmt.Println("year: ",time1.Year(), "month : ",time1.Month())

	//string to time
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())


	//predefined layout
	var dates, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	fmt.Println(dates.String())

	//time to string
	var dateS1 = date.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateS1", dateS1)

	//handle error konvert
	var datess, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")
	if err != nil{
		fmt.Println("Error",err.Error())
		return
	}
	fmt.Println(datess)
}