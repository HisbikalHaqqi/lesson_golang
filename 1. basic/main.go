package main

import "fmt";

func main(){

	// function pertama ku
	fmt.Println("hello world");

	/* ----------------------- TIPE VARIABLE ---------------------------- */
	
	//tipe satu
	var firstname string = "John";
	var lastname string;
	lastname = "Wick";
	fmt.Println(firstname,lastname);

	//tipe dua tanpa kata var
	fullname := "Hisbikal";
	fmt.Println(fullname);

	//tipe tiga multi variable
	var first, second, third = "satu","dua","tiga";
	// seventh, eight, ninth := "tujuh", "delapan", "sembilan";
	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println(first,second,third);


	//tipe data yg hanya menampung saja
	_ = "Mitra Integrasi Informatika";

	//tipe data yang menggunaakan keyword new
	name := new(string);
	fmt.Println(name);

	/* ----------------------- TIPE VARIABLE ---------------------------- */

	/* ----------------------- TIPE DATA ---------------------------- */

	// number
	var positiveNumber uint8 = 89;
	var negativeNumber = -1243423644
	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	//decimal
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	//boolean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	//string
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)
	var greeting = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(message, greeting)

	//konstanta
	const phi = 3.14;
	fmt.Println("PHI :" , phi);

	//operator
	var value = (((2 + 6) % 3) * 4 - 2) / 3
	fmt.Println(value);
	var values = (((2 + 6) % 3) * 4 - 2) / 3
	var isEqual = (values == 2)
	fmt.Printf("nilai %d (%t) \n", values, isEqual)

	//condition if
	var point = 8
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	//condition switch case
	var score = 6
	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	//condition switch case gaya if else
	var scores = 6
	switch {
	case scores == 8:
		fmt.Println("perfect")
	case (scores < 8) && (scores > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	//perulangan
	for i:=0; i<10; i++{
		fmt.Println(i);
	}
	for i := 1; i <= 10; i++ {
		if i % 2 == 1 {
			continue
		}
	
		if i > 8 {
			break
		}
		fmt.Println("Angka", i)
	}


	/*------------------------- ARRAY -------------------- */

	//style 1
	var names[4]string
	names[0] = "hisbikal"
	names[1] = "haqqi"
	names[2] = "muflihin"
	names[3] = "tamvan"
	fmt.Println(names[0], names[1], names[2], names[3])

	//style 2
	var jobs = [4]string{"programmer","ui/ux","frontend"}
	fmt.Println("Jumlah element \t\t", len(jobs))
	fmt.Println("Isi semua element \t", jobs)

	//inisiliasasi array tanpa jumlah elemen
	var numbers = [...]string{"1","2"}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	//array multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	//contoh pemanggilan array style 1
	for i:=1;i<len(jobs);i++{
		fmt.Println(jobs[i])
	}

	//contoh pemanggilan array style 2
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	//contoh pemanggilan array style 3
	var favoriteFruits = [4]string{"melon","pear","jackfruit","apple"}
	for _,fruit := range favoriteFruits{
		fmt.Printf("nama buah : %s\n", fruit)
	}

	//alokasi element array dengan menggunakan make
	var hobbies = make([]string, 2)
	hobbies[0] = "playing footbal"
	hobbies[1] = "playing badminton"


	/*------------------------- ARRAY -------------------- */

	/*------------------------- SLICE -------------------- */
	var cooks = []string{"fried egg","fried rice"}
	for _,cook := range cooks{
		fmt.Println(cook)
	}
	fmt.Println(cooks[0:1])
	//len = jumlah elemen , cap = menghitung kapasitas maksimum
	fmt.Println(len(cooks[0:1]))
	fmt.Println(cap(cooks[0:2]))
	/*------------------------- SLICE -------------------- */

	/*------------------------- MAP -------------------- */

	//style 1
	var chicken map[string]int
	chicken = map[string]int{}
	chicken["januari"] = 90
	chicken["februari"] = 100
	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei",     chicken["februari"])  

	//style 2
	var food =  map[string]string{
		"januari": "fried chicken",
		"february": "friend egg",
	}
	fmt.Println(len(chicken))
	fmt.Println("januari", food["januari"])
	fmt.Println("february",food["february"])  
	for key, val := range food{
		fmt.Println(key, "  \t:", val)
	}

	//delete map
	delete(chicken,"januari")

	//cek item dengan key tertentu
	var meal = map[string]string{
		"januari":"pizza",
		"february":"hotdog",
	}
	var valuee, isExist = meal["januari"]

	if isExist {
		fmt.Println(valuee)
	} else {
		fmt.Println("item is not exists")
	}
	
	//kombinasi slice dengan map
	var data = []map[string]string{
		map[string]string{"name":"chicken blue","gender":"male"},
		map[string]string{"name":"chicken red","gender":"male"},
		map[string]string{"name":"chicken yellow","gender":"male"},
	}

	for _,item:= range data{
		fmt.Println("name : ",item["name"], "gender : ",item["gender"])
	}

	/*------------------------- MAP -------------------- */

	/*------------------------- POINTER  -------------------- */
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)


	var number = 4
    fmt.Println("before :", number) // 4

    change(&number, 10)
    fmt.Println("after  :", number) // 10
	/*------------------------- POINTER -------------------- */

}

func change(original *int, value int)  {
	*original = value
}