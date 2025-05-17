package main

import "fmt"

type Sampah struct {
	Jenis_Sampah  string
	Jumlah_Sampah int
}

const NMAX int = 100

type DataSampah [NMAX]Sampah

var Data DataSampah
var nData, pilihan int

func tampilanMenu() {
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
	nData = 0
	for {
		tampilanMenu()
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			PenjelasanSampah()
		case 2:
			InputSampah(&Data, &nData)
			fmt.Println()
			fmt.Println()
		case 3:
			if nData == 0 {
				fmt.Println("Masukan Data Sampah Terlebih Dahulu")
			} else {
				var Jenis string
				fmt.Println("==========Selamat Datang di Menu Cari Sampah==========")
				fmt.Println("")
				fmt.Println("Jenis Sampah apa yang kamu cari? ")
				fmt.Scan(&Jenis)
				nama, jumlah, ditemukan := CariDataSeq(&Data, nData, Jenis)
				if ditemukan {
					fmt.Println("Sampah", nama, "memiliki jumlah:", jumlah, "kg")
				} else {
					fmt.Println("Data sampah", Jenis, "tidak ditemukan")
				}
			}
			fmt.Println()
			fmt.Println()
		case 4:
			if nData == 0 {
				fmt.Println("Masukan Data Sampah Terlebih Dahulu")
			} else {
				UrutkanData()
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

func PenjelasanSampah() {
	var x int
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
	fmt.Println("==========Selamat Datang di Menu Input Sampah==========")
	fmt.Println("")
	fmt.Println("Masukan data sampah yang akan diinputkan")
	fmt.Println("Format input: Tipe_Sampah Sampah Berat(Kg)")
	fmt.Println("Contoh: Anorganik Plastik 5")
	fmt.Println("Catatan: Isi dengan '# 0' ketika sudah selesai menginput sampah")
	fmt.Println("")
	fmt.Println("Silahkan Masukan Data Sampah:")
	for {
		fmt.Scan(&A[i].Jenis_Sampah, &A[i].Jumlah_Sampah)
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
