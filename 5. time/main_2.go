/*
    TIME DURATION

*/

package main

import "fmt"
import "time"

func main () {

	//type 1
    fmt.Println("start")
    time.Sleep(time.Second * 4)
    fmt.Println("after 4 seconds")

	//type 2
	start := time.Now()
    time.Sleep(5 * time.Second)
    duration := time.Since(start)
    fmt.Println("time elapsed in seconds:", duration.Seconds())
    fmt.Println("time elapsed in minutes:", duration.Minutes())
    fmt.Println("time elapsed in hours:", duration.Hours())

    //type 3
    t1 := time.Now()
    time.Sleep(5 * time.Second)
    t2 := time.Now()

    var dur = t2.Sub(t1)

    fmt.Println("time elapsed in seconds:", dur.Seconds())
    fmt.Println("time elapsed in minutes:", dur.Minutes())
    fmt.Println("time elapsed in hours:", dur.Hours())


    //tpye 4
    var minute = 12 * time.Minute             
    var hour = 65 * time.Hour                
    var microsecond = 45 * time.Microsecond  
    var nanosecond = 233 * time.Nanosecond         

    fmt.Println(hour,"-",minute,"-",microsecond,"-",nanosecond)
}