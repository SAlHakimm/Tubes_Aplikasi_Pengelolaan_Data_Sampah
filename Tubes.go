package main

import "fmt"

type dataSampah struct {
	Jenis_Sampah  string
	Jumlah_Sampah int
}

const NMAX int = 100

type DataSampah [NMAX]dataSampah

func main() {
	var Data dataSampah
	var nData int
	fmt.Scan(&nData)
}

func InputSampah(A *dataSampah, n int) {
	var i int
	for i = 0; i <= n-1; i++ {
		fmt.Scan(&A[i].Jenis_Sampah, &A[i].Jumlah_Sampah)
	}
}
