package main

import "fmt"

type Sampah struct {
	Jenis_Sampah  string
	Jumlah_Sampah int
}

const NMAX int = 100

type DataSampah [NMAX]Sampah

var Data DataSampah
var nData int = 0
var pilihan int

func tampilanMenu() {
	fmt.Println()
	fmt.Println("========== Selamat Datang di Aplikasi Pengelolaan Data Sampah ==========")
	fmt.Println("")
	fmt.Println("1. Penjelasan Jenis Sampah ")
	fmt.Println("2. Input Data Sampah")
	fmt.Println("3. Cari Data Sampah")
	fmt.Println("4. Urutkan Data Sampah")
	fmt.Println("5. Tampilkan Statistik Data Sampah")
	fmt.Println("0. Keluar")
	fmt.Println(" ")
	fmt.Print("Pilih 1/2/3/4/5/0? ")
}

func main() {
	for {
		tampilanMenu()
		fmt.Scan(&pilihan)
		fmt.Println()
		switch pilihan {
		case 1:
			PenjelasanSampah()
		case 2:
			InputSampah(&Data, &nData)
			fmt.Println()
		case 3:
			if nData == 0 {
				fmt.Println("Masukan Data Sampah Terlebih Dahulu")
			} else {
				cariData()
				fmt.Println()
			}
		case 4:
			if nData == 0 {
				fmt.Println("Masukan Data Sampah Terlebih Dahulu")
			} else {
				pengurutanData()
				fmt.Println()
			}
		case 5:
			if nData == 0 {
				fmt.Println("Masukan Data Sampah Terlebih Dahulu")
			} else {
				StatistikData()
			}
		case 6:
		}
		if pilihan == 0 {
			break
		}
	}

}

func cariData() {
	var Jenis string
	fmt.Println()
	fmt.Println("==========Selamat Datang di Menu Cari Sampah==========")
	fmt.Println("")
	fmt.Println("Jenis Sampah apa yang kamu cari? ")
	fmt.Scan(&Jenis)
	fmt.Println()
	nama, jumlah, ditemukan := CariDataSeq(&Data, nData, Jenis)
	if ditemukan {
		fmt.Println("Sampah Jenis", nama, "memiliki jumlah:", jumlah, "kg")
	} else {
		fmt.Println("Data sampah", Jenis, "tidak ditemukan")
	}
}

func pengurutanData() {
	var Pilih int
	fmt.Println("==========Selamat Datang di Menu Pengurutan Data Sampah==========")
	fmt.Println("")
	fmt.Println("1. Pengurutan berdasarkan banyaknya sampah yang paling berat")
	fmt.Println("2. Pengurutan berdasarkan banyaknya sampah yang paling ringan")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih 1/2/3? ")
	fmt.Scan(&Pilih)
	if Pilih == 1 {
		fmt.Println()
		fmt.Println("Pengurutan data Sampah dari yang paling berat: ")
		SortDataMaxSelection(&Data, nData)
		cetakDataSort(Data, nData)
	} else if Pilih == 2 {
		fmt.Println()
		fmt.Println("Pengurutan data Sampah dari yang paling kecil: ")
		SortDataMinInsert(&Data, nData)
		cetakDataSort(Data, nData)
	} else if Pilih == 3 {
		tampilanMenu()
	}
}

func PenjelasanSampah() {
	var x int
	fmt.Println()
	fmt.Println("Penjelasan Tentang Sampah yang perlu kamu ketahui sebelum menginput data :")
	fmt.Println("")
	fmt.Println("Apa yang ingin Kamu Ketahui?")
	fmt.Println("1. Sampah Organik: Sampah organik berasal dari sisa-sisa makhluk hidup, baik tumbuhan maupun hewan, yang secara alami dapat terurai oleh mikroorganisme (dekomposer). Proses penguraian ini menghasilkan zat-zat yang dapat menyuburkan tanah.")
	fmt.Println("Contoh Sampah Organik: Sisa Makanan, Daun-daun kering/ranting, Kertas/Kardus bekas, Kotoran Hewan.")
	fmt.Println("")
	fmt.Println("2. Sampah Anorganik: Sampah anorganik adalah jenis sampah yang tidak dapat diuraikan secara alami oleh mikroorganisme dalam waktu yang relatif singkat. Beberapa jenis sampah anorganik dapat didaur ulang menjadi barang yang berguna.")
	fmt.Println("Contoh Sampah Anorganik: Plastik, Kaca, Logam, Karet,Keramik. ")
	fmt.Println("")
	fmt.Println("3. Sampah B3: Sampah B3 mengandung zat-zat yang dapat membahayakan kesehatan manusia, makhluk hidup lain, dan lingkungan. Penanganannya memerlukan perlakuan khusus agar tidak mencemari lingkungan.")
	fmt.Println("Contoh Sampah B3: Baterai, Lampu, Obat-obatan kadaluarsa, Pestisida, Cat, Aki bekas, Limbah Elektronik.")
	fmt.Println("")
	fmt.Println("4. Sampah Medis: Sampah medis dihasilkan dari kegiatan di fasilitas kesehatan seperti rumah sakit, klinik, puskesmas, dan laboratorium. Jenis sampah ini berpotensi mengandung mikroorganisme patogen yang dapat menularkan penyakit.")
	fmt.Println("Cotoh Sampah Medis: Jarum suntik bekas, perban bekas, infus set bekas, organ tubuh.")
	fmt.Println("")
	fmt.Print("Ketik 0 jika sudah selesai membaca: ")
	fmt.Scan(&x)
	if x == 0 {
		main()
	}
}

func InputSampah(A *DataSampah, n *int) {
	var i int = 0
	fmt.Println()
	fmt.Println("==========Selamat Datang di Menu Input Sampah==========")
	fmt.Println("")
	fmt.Println("Masukan data sampah yang akan diinputkan")
	fmt.Println("Format input: Sampah Berat(Kg)")
	fmt.Println("Contoh: Plastik 5")
	fmt.Println("Catatan: Isi dengan '# 0' ketika sudah selesai menginput sampah")
	fmt.Println("")
	fmt.Println("Silahkan Masukan Data Sampah:")
	for {
		fmt.Print("Masukan Jenis Sampah: ")
		fmt.Scan(&A[i].Jenis_Sampah)
		fmt.Print("Masukan Jumlah Sampah (kg): ")
		fmt.Scan(&A[i].Jumlah_Sampah)
		fmt.Println()
		if A[i].Jenis_Sampah == "#" && A[i].Jumlah_Sampah == 0 {
			break
		}
		nData++
		i++
	}
}

func CariDataSeq(A *DataSampah, n int, x string) (string, int, bool) {
	var i, jumlah int
	var nama string
	var ditemukan bool
	ditemukan = false
	for i = 0; i < n; i++ {
		if A[i].Jenis_Sampah == x {
			nama = A[i].Jenis_Sampah
			jumlah = A[i].Jumlah_Sampah
			ditemukan = true
			return nama, jumlah, ditemukan
		}
	}
	return "", 0, ditemukan
}

func SortDataMaxSelection(A *DataSampah, n int) {
	var i, idx, pass int
	var temp Sampah

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].Jumlah_Sampah >= A[idx].Jumlah_Sampah {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp

		pass++
	}
}

func SortDataMinInsert(A *DataSampah, n int) {
	var pass, i int
	var temp Sampah
	pass = 1
	for pass < n {
		i = pass
		temp = A[pass]
		for i > 0 && temp.Jumlah_Sampah < A[i-1].Jumlah_Sampah {
			A[i].Jumlah_Sampah = A[i-1].Jumlah_Sampah
			i--
		}
		A[i].Jumlah_Sampah = temp.Jumlah_Sampah
		pass++
	}
}

func cetakDataSort(A DataSampah, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Println(A[i].Jenis_Sampah, A[i].Jumlah_Sampah)
	}
}
