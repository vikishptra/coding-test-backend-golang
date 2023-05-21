package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// soal 1
func readFile(namaFile string) ([]string, error) {
	//membaca isi file
	isiContent, err := ioutil.ReadFile(namaFile)
	if err != nil {
		return nil, err
	}
	//split "\n" dan kirim ke variable num yang menjadi []string
	num := strings.Split(string(isiContent), "\n")
	return num, err
}

func resultSum() {
	result := 0
	//memangil function dari readFile di tampung ke numArr berupa return []string dan error
	numArr, err := readFile("file.txt")
	if err != nil {
		fmt.Println("err: ", err)
	}
	// //loop untuk menghitung nilai dari isi numArr
	for _, v := range numArr {
		//convert string ke int di tampung variabel num
		num, _ := strconv.Atoi(v)
		//menghitung jumlah angka di slice num lalu di tampung di result sampai range daru numArr
		result += num
	}
	//menghitung jumlah dari len resultArr
	fmt.Println("Total Angka pada file", len(numArr))
	//menampilkan hasil dari result
	fmt.Println("Jumlah semua angka", result)

}

// soal 2
func searchPalindrome(s string) {
	s = strings.ToLower(s)
	flag := true
	count := len(s)
	for i := 0; i < count; i++ {
		j := len(s) - 1 - i
		if s[i] != s[j] {
			flag = false
			break
		}
	}
	if flag {
		fmt.Println("Kata/kalimat yang diinput:", s)
		fmt.Println("Merupakan Palindrom")
	} else {
		fmt.Println("Kata/kalimat yang diinput:", s)
		fmt.Println("Bukan Palindrom")
	}
}

// soal 3
func fetchAPI() {
	//fetch api dengan http method get
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println("err", err)
	}
	//membaca seluruh response body dari res
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err", err)
	}
	//buat struct Response
	var result []Response
	//ubah byte ke []result yaitu Response
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		fmt.Println("err", err)
	}
	//dengan menggunakan MarshalIndent data dari result terlihat rapih dengan terformat json
	jsonData, err := json.MarshalIndent(result, "", " ")
	fmt.Println(string(jsonData))
}

// soal 4
func factorial(n []int) {
	//untuk mengatur seed dari generator angka acak
	rand.Seed(time.Now().UnixNano())
	//jika sudah mencapai kurang dari sama 1 maka di return aja untuk memberhentikan function
	if len(n) <= 1 {
		return
	}
	//lastIndex di gunakan untuk petukaran angka acak
	lastIndex := len(n) - 1
	//memangil func factorial untuk memperoleh slice yang lebih pendek dari lastIndex
	factorial(n[:lastIndex])
	//generate nilai acak dari 0 sampai lastIndex
	randIndex := rand.Intn(lastIndex)
	//petukaran nilai antara slice akhir dan elemen slice secara acak
	n[lastIndex], n[randIndex] = n[randIndex], n[lastIndex]
}

func resultFactorial() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	factorial(num)
	for _, v := range num {
		fmt.Println(v)
	}
}

// soal 5
func convertTime() {
	time12 := "03:04:05PM" //contoh format 12 jam
	time24 := "15:04:05"   //contoh format 24 jam
	var input string
	fmt.Print("Masukkan waktu dalam format 12 jam : ")
	fmt.Scanf("%s", &input)
	//ubah input string menjadi objek time.time dari time12
	t, err := time.Parse(time12, input)
	if err != nil {
		fmt.Println("Invalid input format")
		return
	}
	fmt.Println(t.Format(time24))
}

// soal 6
func factorial2(n int) {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	fmt.Println(fact)
}

// soal 7
// Input: nums = [1,2,3,4]
// Output: [1,3,6,10]
// Input: nums = [1,1,1,1,1]
// Output: [1,2,3,4,5]
// Input: nums = [3,1,2,10,1]
// Output: [3,4,6,16,17]
func CaseRunningSum(nums []int) []int {
	var arr []int
	key := nums[0]
	for i, _ := range nums {
		i++
		arr = append(arr, key)
		if i == len(nums) {
			break
		}
		key = key + nums[i]

	}
	return arr

}

// soal 8
func CaseTwo(s string) map[string]int {
	//s = "amnsdbmansdbmansdba" return [a:4 b:3 d:3 m:3 n:3 s:3]
	flag := 0
	results := make(map[string]int)
	for i := 0; i < len(s); i++ {
		flag = 0
		for _, v := range s {
			if string(s[i]) == string(v) {
				flag++
				results[string(s[i])] = flag
			}
		}
	}
	return results
}

// soal 9
func fizzbuzz(n int) {
	fizz := "fizz"
	buzz := "buzz"
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, fizz+buzz)
		} else if i%3 == 0 {
			fmt.Println(i, fizz)
		} else if i%5 == 0 {
			fmt.Println(i, buzz)
		} else {
			fmt.Println(i)
		}
	}

}

// soal 10
func searchNum(n []int) {

	besar := n[0]
	kecil := n[0]
	for i, _ := range n {
		if i == len(n)-1 {
			break
		}
		i += 1
		if besar < n[i] {
			besar = n[i]
		} else if n[i] < kecil {
			kecil = n[i]
		}

	}
	fmt.Println("Terbesar : ", besar, "Terkecil :", kecil)
}

// soal 11
func countNumeric(n string) {
	angka := 0
	huruf := 0

	for _, v := range n {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			huruf++
		} else if v >= '0' && v <= '9' {
			angka++
		}
	}
	fmt.Println("angka result :", angka, "huruf result :", huruf)
}
