package main

import "fmt"
import "math"

func main(){
	divideNumber(0,0)
	divideNumber(0,4)

	var diameter float64 = 15
	var area, circumference = multiple(diameter)
	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
    fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	var avg = calculate(2,34,4,2,54,3,4)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
    fmt.Println(msg)

	//closure style 1
	var getMinMax = func(n []int) (int, int) {
        var min, max int
        for i, e := range n {
            switch {
            case i == 0:
                max, min = e, e
            case e > max:
                max = e
            case e < min:
                min = e
            }
        }
        return min, max
    }
	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
    var min, max = getMinMax(numbers)
    fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)

	//closure style
	var maxx = 3
    var numberss = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
    var howMany, getNumbers = findMax(numberss, maxx)
    var theNumbers = getNumbers()

    fmt.Println("numbers\t:", numbers)
    fmt.Printf("find \t: %d\n\n", max)

    fmt.Println("found \t:", howMany)
    fmt.Println("value \t:", theNumbers)
}


//void
func divideNumber(m, n int) {
    if n == 0 {
        fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
        return
    }

    var res = m / n
    fmt.Printf("%d / %d = %d\n", m, n, res)
}

//multiple return
func multiple(d float64)(float64,float64){
	var area = math.Pi * math.Pow(d / 2, 2)
	var circumference = math.Pi

	return area, circumference
}

//variadic = bisa menampung berapa saja
func calculate(numbers ...int) float64 {
    var total int = 0
    for _, number := range numbers {
        total += number
    }

    var avg = float64(total) / float64(len(numbers))
    return avg
}

//find max with closure
func findMax(numbers []int, max int) (int, func() []int) {
    var res []int
    for _, e := range numbers {
        if e <= max {
            res = append(res, e)
        }
    }
    return len(res), func() []int {
        return res
    }
}
