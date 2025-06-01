package main

import "fmt"

type Sampah struct {
	ID                int
	Jenis_Sampah      string
	Daur_Ulang        bool
	sampah            string
	Metode_Daur_Ulang string
	Jumlah_Sampah     int
}

const NMAX int = 10

type DataSampah [NMAX]Sampah

var Data DataSampah
var nData int = 0
var pilihan int
var IDc int = 1000

func tampilanMenu() {
	fmt.Println()
	fmt.Println("========================================================================")
	fmt.Println("========== Selamat Datang di Aplikasi Pengelolaan Data Sampah ==========")
	fmt.Println("========================================================================")
	fmt.Println("")
	fmt.Println("1. Penjelasan Jenis Sampah ")
	fmt.Println("2. Input Data Sampah")
	fmt.Println("3. Tambah, Edit, dan Hapus Data Sampah")
	fmt.Println("4. Tampilkan Semua Data Sampah atau Cari Data Sampah")
	fmt.Println("5. Urutkan Data Sampah")
	fmt.Println("6. Tampilkan Statistik Data Sampah")
	fmt.Println("0. Keluar")
	fmt.Println(" ")
	fmt.Print("Pilih 1/2/3/4/5/6/0? ")
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
				tambahHapusData()
				fmt.Println()
			}
		case 4:
			if nData == 0 {
				fmt.Println("Masukan Data Sampah Terlebih Dahulu")
			} else {
				BacadanCari()
			}
		case 5:
			if nData == 0 {
				fmt.Println("Masukan Data Sampah Terlebih Dahulu")
			} else {
				pengurutanData()

			}
		case 6:
			if nData == 0 {
				fmt.Println("Masukan Data Sampah Terlebih Dahulu")
			} else {
				StatistikData()
			}
		case 7:
		}
		if pilihan == 0 {
			fmt.Println("Selamat Tinggal dan Terimakasih Telah Menggunakan Aplikasi :)")
			break
		}
	}

}

func BacadanCari() {
	var pilih int
	fmt.Println()
	fmt.Println("==========Selamat Datang di Menu Cetak dan Cari Data Sampah==========")
	fmt.Println("")
	fmt.Println("1. Cari Data Sampah ")
	fmt.Println("2. Cetak semua Data Sampah")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih 1/2/3? ")
	fmt.Scan(&pilih)
	fmt.Println()
	if pilih == 1 {
		cariData()
		fmt.Println()
	} else if pilih == 2 {
		cetakData(Data, nData)
		fmt.Println()
	} else {
		return
	}

}

func cetakData(A DataSampah, n int) {
	fmt.Println("Sampah yang sudah kamu inputkan: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. ID: %d, Sampah: %s, Jumlah: %d kg, Jenis: %s, Daur Ulang: %s\n", i+1, A[i].ID, A[i].Jenis_Sampah, A[i].Jumlah_Sampah, A[i].sampah, A[i].Metode_Daur_Ulang)
	}
}

func tambahHapusData() {
	var x int
	fmt.Println()
	fmt.Println("==========Selamat Datang di Menu Tambah dan Hapus Sampah==========")
	fmt.Println("")
	fmt.Println()
	fmt.Println("1. Tambah Data Sampah")
	fmt.Println("2. Hapus Data Sampah")
	fmt.Println("3. Edit Data Berat")
	fmt.Println("4. Keluar")
	fmt.Print("Pilih 1/2/3/4? ")
	fmt.Scan(&x)
	fmt.Println()
	if x == 1 {
		fmt.Println()
		TambahData(&Data, &nData)
	} else if x == 2 {
		fmt.Println()
		hapusData(&Data, &nData)
	} else if x == 3 {
		fmt.Println()
		EditData(&Data, &nData)
	}
}

func EditData(A *DataSampah, n *int) {
	var i int = 0
	var Jenis string
	var id int
	var ketemu bool = false
	fmt.Println()
	fmt.Println("==========Selamat Datang di Menu Edit Data Sampah==========")
	fmt.Println("")
	cetakData(Data, nData)
	fmt.Println()

	fmt.Print("ID Sampah Apa Yang Ingin Kamu Ubah? ")
	fmt.Scan(&id)
	for i < *n && !ketemu {
		ketemu = id == A[i].ID
		if ketemu {
			fmt.Print("Mau ubah menjadi berapa kg? ")
			fmt.Scan(&A[i].Jumlah_Sampah)
			fmt.Println()
			fmt.Println("Data", Jenis, "telah diubah menjadi", A[i].Jumlah_Sampah, "kg")
		}
		i++
	}
	if ketemu == false {
		fmt.Println("Data Tidak Ditemukan :(")

	}
}

func TambahData(A *DataSampah, n *int) {
	var pilih, x int
	fmt.Println("Silahkan masukan data sampah yang ingin kamu tambahkan:")
	fmt.Println()
	fmt.Println("Data yang kamu miliki (Disarankan tidak memasukan data yang sama): ")
	cetakData(Data, nData)
	fmt.Println()
	if *n >= NMAX {
		*n = NMAX
		fmt.Println("Kapasitas Maximal tercapai = ", *n)
		return
	}
	fmt.Print("Masukan Sampah: ")
	fmt.Scan(&A[*n].Jenis_Sampah)

	fmt.Print("Masukan Jumlah Sampah (kg): ")
	fmt.Scan(&A[*n].Jumlah_Sampah)

	for {
		fmt.Println("Masukan Jenis Sampah : ")
		fmt.Println("1. Organik")
		fmt.Println("2. Anorganik")
		fmt.Println("3. B3")
		fmt.Println("Pilih 1/2/3")
		fmt.Scan(&x)
		if x == 1 {
			A[*n].sampah = "Organik"
			break
		} else if x == 2 {
			A[*n].sampah = "Anorganik"
			break
		} else if x == 3 {
			A[*n].sampah = "B3"
			break
		} else {
			fmt.Println("Pilihan Tidak Valid, Silahkan Pilih Kembali")
		}
	}
	for {
		if A[*n].sampah == "Organik" || A[*n].sampah == "organik" {
			fmt.Println("Pilih Metode Daur Ulang:")
			fmt.Println("1. Daur Ulang Organik (Organic Recycling)")
			fmt.Println("2. Daur Ulang Energi (Energy Recovery)")
			fmt.Println("3. Tidak Akan Di Daur Ulang")
			fmt.Print("Pilih 1/2/3? ")
			fmt.Scan(&pilih)

			if pilih == 1 {
				A[*n].Metode_Daur_Ulang = "Daur ulang Organik"
				A[*n].Daur_Ulang = true
			} else if pilih == 2 {
				A[*n].Metode_Daur_Ulang = "Daur ulang Energi"
				A[*n].Daur_Ulang = true
			} else if pilih == 3 {
				A[*n].Metode_Daur_Ulang = "Tidak Didaur Ulang"
				A[*n].Daur_Ulang = false
			} else {
				fmt.Println("Pilihan tidak valid. Dianggap tidak didaur ulang.")
				A[*n].Metode_Daur_Ulang = "Tidak Didaur Ulang"
				A[*n].Daur_Ulang = false
			}
		} else if A[*n].sampah == "Anorganik" || A[*n].sampah == "anorganik" {
			fmt.Println("Pilih Metode Daur Ulang:")
			fmt.Println("1. Daur Ulang Mekanis (Mechanical Recycling)")
			fmt.Println("2. Daur Ulang Kimia (Chemical Recovery)")
			fmt.Println("3. Daur Ulang Upcycling")
			fmt.Println("4. Tidak Akan Di Daur Ulang")
			fmt.Print("Pilih 1/2/3/4? ")
			fmt.Scan(&pilih)

			if pilih == 1 {
				A[*n].Metode_Daur_Ulang = "Daur ulang Mekanis"
				A[*n].Daur_Ulang = true
			} else if pilih == 2 {
				A[*n].Metode_Daur_Ulang = "Daur ulang Kimia"
				A[*n].Daur_Ulang = true
			} else if pilih == 3 {
				A[*n].Metode_Daur_Ulang = "Daur Ulang Upcycling"
				A[*n].Daur_Ulang = true
			} else if pilih == 4 {
				A[*n].Metode_Daur_Ulang = "Tidak Didaur Ulang"
				A[*n].Daur_Ulang = false
			} else {
				fmt.Println("Pilihan tidak valid. Dianggap tidak didaur ulang.")
				A[*n].Metode_Daur_Ulang = "Tidak Didaur Ulang"
				A[*n].Daur_Ulang = false
			}
		} else if A[*n].sampah == "B3" || A[*n].sampah == "b3" {
			fmt.Println("Pilih Metode Daur Ulang:")
			fmt.Println("1. Daur Ulang Mekanis (Mechanical Recycling)")
			fmt.Println("2. Daur Ulang Kimia (Chemical Recycling)")
			fmt.Println("3. Daur Ulang Energi (Energy Recovery)")
			fmt.Println("4. Tidak Akan Di Daur Ulang")
			fmt.Print("Pilih 1/2/3/4? ")
			fmt.Scan(&pilih)

			if pilih == 1 {
				A[*n].Metode_Daur_Ulang = "Daur ulang Mekanis"
				A[*n].Daur_Ulang = true
			} else if pilih == 2 {
				A[*n].Metode_Daur_Ulang = "Daur Ulang Kimia"
				A[*n].Daur_Ulang = true
			} else if pilih == 3 {
				A[*n].Metode_Daur_Ulang = "Daur ulang Energi"
				A[*n].Daur_Ulang = true
			} else if pilih == 4 {
				A[*n].Metode_Daur_Ulang = "Tidak Didaur Ulang"
				A[*n].Daur_Ulang = false
			} else {
				fmt.Println("Pilihan tidak valid. Dianggap tidak didaur ulang.")
				A[*n].Metode_Daur_Ulang = "Tidak Didaur Ulang"
				A[*n].Daur_Ulang = false
			}
		} else {
			fmt.Println("Masukan data dengan benar!")
		}
		if A[*n].Metode_Daur_Ulang != "" {
			break
		}
	}

	A[*n].ID = generateID()
	fmt.Printf("ID Sampah: %d\n\n", A[*n].ID)
	*n = *n + 1
	fmt.Println()
	fmt.Println("Data Telah Ditambahkan :)")
}

func hapusData(A *DataSampah, n *int) {
	var i, j, id int
	var ketemu bool = false
	fmt.Println("Data yang kamu miliki: ")
	cetakData(Data, nData)
	fmt.Println()

	fmt.Print("ID Sampah apa yang akan dihapus: ")
	fmt.Scan(&id)

	for i = 0; i < *n && ketemu == false; i++ {
		if A[i].ID == id {
			for j = i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			fmt.Println()
			fmt.Println("Data Berhasil Dihapus :)")
			ketemu = true
			*n = *n - 1
			fmt.Println()
		}
	}
	if ketemu == false {
		fmt.Println()
		fmt.Println("Data yang kamu ingin hapus tidak ditemukan :(")
		fmt.Println()

	}
}

func cariData() {
	var pilih int
	fmt.Println()
	fmt.Println("==========Selamat Datang di Menu Cari Sampah==========")
	fmt.Println("")
	fmt.Println("1. Cari Berdasarkan ID")
	fmt.Println("2. Cari Berdasarkan Jenis")
	fmt.Println("3. Keluar")
	fmt.Print("1/2/3? ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		CariDataIDBin(&Data, nData)
	} else if pilih == 2 {
		CariDataJenisSeq(&Data, nData)
	} else {
		return
	}

}

func CariDataJenisSeq(A *DataSampah, n int) {
	var pilih, i, j int
	var x string
	fmt.Println("Mau Cari Data Jenis Apa? ")
	fmt.Println("1. Organik")
	fmt.Println("2. Anorganik")
	fmt.Println("3. B3")
	fmt.Println("4. Keluar")
	fmt.Print("Pilih 1/2/3/4? ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		for i = 0; i < n; i++ {
			if A[i].sampah == "Organik" || A[i].sampah == "organik" {
				j++
				fmt.Printf("%d. ID: %d, Sampah: %s, Jumlah: %d kg, Jenis: %s, Daur Ulang: %s\n", j, A[i].ID, A[i].Jenis_Sampah, A[i].Jumlah_Sampah, A[i].sampah, A[i].Metode_Daur_Ulang)
			}
		}
	} else if pilih == 2 {
		for i = 0; i < n; i++ {
			if A[i].sampah == "Anorganik" || A[i].sampah == "anorganik" {
				j++
				fmt.Printf("%d. ID: %d, Sampah: %s, Jumlah: %d kg, Jenis: %s, Daur Ulang: %s\n", j, A[i].ID, A[i].Jenis_Sampah, A[i].Jumlah_Sampah, A[i].sampah, A[i].Metode_Daur_Ulang)
			}
		}
	} else if pilih == 3 {
		for i = 0; i < n; i++ {
			if A[i].sampah == "B3" || A[i].sampah == "b3" {
				j++
				fmt.Printf("%d. ID: %d, Sampah: %s, Jumlah: %d kg, Jenis: %s, Daur Ulang: %s\n", j, A[i].ID, A[i].Jenis_Sampah, A[i].Jumlah_Sampah, A[i].sampah, A[i].Metode_Daur_Ulang)
			}
		}
	} else {
		return
	}
	fmt.Print("Apakah kamu masih ingin membaca dan mencari data (Yes/No)? ")
	fmt.Scan(&x)
	if x == "Yes" || x == "yes" || x == "Y" || x == "y" {
		BacadanCari()
		fmt.Println()
	}
}

func CariDataIDBin(A *DataSampah, n int) {
	var left, right, mid int
	var idx int
	var dicari int
	var x string
	fmt.Print("Masukan ID Yang Ingin Dicari: ")
	fmt.Scan(&dicari)

	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if dicari < A[mid].ID {
			right = mid - 1
		} else if dicari > A[mid].ID {
			left = mid + 1
		} else {
			idx = mid
		}
		mid = (left + right) / 2
	}
	if idx == -1 {
		fmt.Println("Data yang kamu cari tidak ditemukan :(")
	} else {
		fmt.Printf("ID: %d, Sampah: %s, Jumlah: %d kg, Jenis: %s, Daur Ulang: %s\n", A[idx].ID, A[idx].Jenis_Sampah, A[idx].Jumlah_Sampah, A[idx].sampah, A[idx].Metode_Daur_Ulang)
	}
	fmt.Print("Apakah kamu masih ingin membaca dan mencari data (Yes/No)? ")
	fmt.Scan(&x)
	if x == "Yes" || x == "yes" || x == "Y" || x == "y" {
		BacadanCari()
		fmt.Println()
	}
}

func pengurutanData() {
	var Pilih int
	var x string
	fmt.Println("==========Selamat Datang di Menu Pengurutan Data Sampah==========")
	fmt.Println("")
	fmt.Println("1. Pengurutan berdasarkan banyaknya sampah yang paling berat")
	fmt.Println("2. Pengurutan berdasarkan banyaknya sampah yang paling ringan")
	fmt.Println("3. Pengurutan berdasarkan Jenis Sampah (Alphabet)")
	fmt.Println("4. Keluar")
	fmt.Print("Pilih 1/2/3/4? ")
	fmt.Scan(&Pilih)
	if Pilih == 1 {
		fmt.Println()
		fmt.Println("Pengurutan data Sampah dari yang paling berat: ")
		SortDataMaxSelection(Data, nData)
	} else if Pilih == 2 {
		fmt.Println()
		fmt.Println("Pengurutan data Sampah dari yang paling kecil: ")
		SortDataMinInsert(Data, nData)
	} else if Pilih == 3 {
		fmt.Println()
		fmt.Println("Pengurutan data sampah berdasarkan jenis (Alphabet): ")
		SortDataJenisInset(Data, nData)
	} else {
		return
	}
	fmt.Println()
	fmt.Print("Apakah kamu masih ingin mengurutkan data (Yes/No)? ")
	fmt.Scan(&x)
	if x == "Yes" || x == "yes" || x == "Y" || x == "y" {
		pengurutanData()
		fmt.Println()
	}
}

func PenjelasanSampah() {
	var x int
	var pilih int
	fmt.Println()
	fmt.Println("Penjelasan Tentang Sampah yang perlu kamu ketahui sebelum menginput data :")
	fmt.Println("")
	fmt.Println("Apa yang ingin Kamu Ketahui?")
	fmt.Println("1. Jenis Sampah")
	fmt.Println("2. Metode Daur Ulang")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih 1/2/3? ")
	fmt.Scan(&pilih)
	if pilih == 1 {
		fmt.Println("Jenis Sampah :")
		fmt.Println()
		fmt.Println("1. Sampah Organik: Sampah organik berasal dari sisa-sisa makhluk hidup, baik tumbuhan maupun hewan, yang secara alami dapat terurai oleh mikroorganisme (dekomposer). Proses penguraian ini menghasilkan zat-zat yang dapat menyuburkan tanah.")
		fmt.Println("Contoh Sampah Organik: Sisa Makanan, Daun-daun kering/ranting, Kertas/Kardus bekas, Kotoran Hewan.")
		fmt.Println("")
		fmt.Println("2. Sampah Anorganik: Sampah anorganik adalah jenis sampah yang tidak dapat diuraikan secara alami oleh mikroorganisme dalam waktu yang relatif singkat. Beberapa jenis sampah anorganik dapat didaur ulang menjadi barang yang berguna.")
		fmt.Println("Contoh Sampah Anorganik: Plastik, Kaca, Logam, Karet,Keramik. ")
		fmt.Println("")
		fmt.Println("3. Sampah B3: Sampah B3 mengandung zat-zat yang dapat membahayakan kesehatan manusia, makhluk hidup lain, dan lingkungan. Penanganannya memerlukan perlakuan khusus agar tidak mencemari lingkungan.")
		fmt.Println("Contoh Sampah B3: Baterai, Lampu, Obat-obatan kadaluarsa, Pestisida, Cat, Aki bekas, Limbah Elektronik.")
		fmt.Println("")
		fmt.Print("Ketik angka atau huruf apa saja jika sudah selesai membaca: ")
		fmt.Scan(&x)
		if x == 0 {
			return
		}
	} else if pilih == 2 {
		fmt.Println("Metode Daur Ulang Sampah :")
		fmt.Println()
		fmt.Println("1. Daur Ulang Mekanis (Mechanical Recycling): Metode daur ulang yang paling umum dan melibatkan proses fisik untuk mengubah sampah menjadi bahan baku baru.")
		fmt.Println()
		fmt.Println("2. Daur Ulang Kimia (Chemical Recycling): Metode ini melibatkan proses kimia untuk memecah sampah menjadi monomer atau bahan kimia dasar penyusunnya. Bahan-bahan ini kemudian dapat digunakan untuk membuat produk baru yang kualitasnya seringkali setara dengan produk dari bahan baku primer.")
		fmt.Println()
		fmt.Println("3. Daur Ulang Organik (Organic Recycling): Metode ini khusus untuk limbah organik yang dapat terurai secara hayati.")
		fmt.Println()
		fmt.Println("4. Daur Ulang Energi (Energy Recovery): Metode ini mengubah sampah menjadi energi melalui pembakaran terkontrol.")
		fmt.Println()
		fmt.Println("5. Daur Ulang Upcycling (Upcycling): Ini adalah bentuk daur ulang kreatif di mana barang bekas atau sampah diubah menjadi produk baru dengan nilai yang lebih tinggi atau fungsi yang berbeda, tanpa perlu proses pengolahan yang kompleks.")
		fmt.Println()
		fmt.Print("Ketik angka atau huruf apa saja jika sudah selesai membaca: ")
		fmt.Scan(&x)
		if x == 0 {
			return
		}
	} else {
		return
	}
}

func generateID() int {
	IDc++
	return IDc
}

func InputSampah(A *DataSampah, n *int) {
	var i int = 0
	var pilih, x int
	fmt.Println()
	if *n != 0 {
		fmt.Println("Jika ingin menginput data, masukan di menu tambah data")
	} else {
		fmt.Println("==========Selamat Datang di Menu Input Sampah==========")
		fmt.Println("")
		fmt.Println("Masukan data sampah yang akan diinputkan")
		fmt.Println("Contoh Format Input: ")
		fmt.Println("Masukan Sampah: Plastik")
		fmt.Println("Masukan Jumlah Sampah (kg): 5")
		fmt.Println("Catatan: Isi dengan '#' ketika sudah selesai menginput sampah")
		fmt.Println("")
		fmt.Println("Silahkan Masukan Data Sampah:")
		for {
			if *n >= NMAX {
				*n = NMAX
				fmt.Println("Kapasitas Maximal tercapai =", *n)
				return
			}
			fmt.Print("Masukan Sampah: ")
			fmt.Scan(&A[i].Jenis_Sampah)
			if A[i].Jenis_Sampah == "#" {
				break
			}
			fmt.Print("Masukan Jumlah Sampah (kg): ")
			fmt.Scan(&A[i].Jumlah_Sampah)
			if A[i].Jenis_Sampah != "#" && A[i].Jumlah_Sampah != 0 {
				for {
					fmt.Println("Masukan Jenis Sampah : ")
					fmt.Println("1. Organik")
					fmt.Println("2. Anorganik")
					fmt.Println("3. B3")
					fmt.Println("Pilih 1/2/3")
					fmt.Scan(&x)
					if x == 1 {
						A[i].sampah = "Organik"
						break
					} else if x == 2 {
						A[i].sampah = "Anorganik"
						break
					} else if x == 3 {
						A[i].sampah = "B3"
						break
					} else {
						fmt.Println("Pilihan Tidak Valid, Silahkan Pilih Kembali")
					}
				}
				if A[i].sampah == "Organik" || A[i].sampah == "organik" {
					fmt.Println("Pilih Metode Daur Ulang:")
					fmt.Println("1. Daur Ulang Organik (Organic Recycling)")
					fmt.Println("2. Daur Ulang Energi (Energy Recovery)")
					fmt.Println("3. Tidak Akan Di Daur Ulang")
					fmt.Print("Pilih 1/2/3? ")
					fmt.Scan(&pilih)

					if pilih == 1 {
						A[i].Metode_Daur_Ulang = "Daur ulang Organik"
						A[i].Daur_Ulang = true
					} else if pilih == 2 {
						A[i].Metode_Daur_Ulang = "Daur ulang Energi"
						A[i].Daur_Ulang = true
					} else if pilih == 3 {
						A[i].Metode_Daur_Ulang = "Tidak Didaur Ulang"
						A[i].Daur_Ulang = false
					} else {
						fmt.Println("Pilihan tidak valid. Dianggap tidak didaur ulang.")
						A[i].Metode_Daur_Ulang = "Tidak Didaur Ulang"
						A[i].Daur_Ulang = false
					}
				} else if A[i].sampah == "Anorganik" || A[i].sampah == "anorganik" {
					fmt.Println("Pilih Metode Daur Ulang:")
					fmt.Println("1. Daur Ulang Mekanis (Mechanical Recycling)")
					fmt.Println("2. Daur Ulang Kimia (Chemical Recovery)")
					fmt.Println("3. Daur Ulang Upcycling")
					fmt.Println("4. Tidak Akan Di Daur Ulang")
					fmt.Print("Pilih 1/2/3/4? ")
					fmt.Scan(&pilih)

					if pilih == 1 {
						A[i].Metode_Daur_Ulang = "Daur ulang Mekanis"
						A[i].Daur_Ulang = true
					} else if pilih == 2 {
						A[i].Metode_Daur_Ulang = "Daur ulang Kimia"
						A[i].Daur_Ulang = true
					} else if pilih == 3 {
						A[i].Metode_Daur_Ulang = "Daur Ulang Upcycling"
						A[i].Daur_Ulang = true
					} else if pilih == 4 {
						A[i].Metode_Daur_Ulang = "Tidak Didaur Ulang"
						A[i].Daur_Ulang = false
					} else {
						fmt.Println("Pilihan tidak valid. Dianggap tidak didaur ulang.")
						A[i].Metode_Daur_Ulang = "Tidak Didaur Ulang"
						A[i].Daur_Ulang = false
					}
				} else if A[i].sampah == "B3" || A[i].sampah == "b3" {
					fmt.Println("Pilih Metode Daur Ulang:")
					fmt.Println("1. Daur Ulang Mekanis (Mechanical Recycling)")
					fmt.Println("2. Daur Ulang Kimia (Chemical Recycling)")
					fmt.Println("3. Daur Ulang Energi (Energy Recovery)")
					fmt.Println("4. Tidak Akan Di Daur Ulang")
					fmt.Print("Pilih 1/2/3/4? ")
					fmt.Scan(&pilih)

					if pilih == 1 {
						A[i].Metode_Daur_Ulang = "Daur ulang Mekanis"
						A[i].Daur_Ulang = true
					} else if pilih == 2 {
						A[i].Metode_Daur_Ulang = "Daur Ulang Kimia"
						A[i].Daur_Ulang = true
					} else if pilih == 3 {
						A[i].Metode_Daur_Ulang = "Daur ulang Energi"
						A[i].Daur_Ulang = true
					} else if pilih == 4 {
						A[i].Metode_Daur_Ulang = "Tidak Didaur Ulang"
						A[i].Daur_Ulang = false
					} else {
						fmt.Println("Pilihan tidak valid. Dianggap tidak didaur ulang.")
						A[i].Metode_Daur_Ulang = "Tidak Didaur Ulang"
						A[i].Daur_Ulang = false
					}
				}

				A[i].ID = generateID()
				fmt.Printf("ID Sampah: %d\n\n", A[i].ID)
			}
			nData++
			i++
		}
	}
}

func SortDataMaxSelection(A DataSampah, n int) {
	var i, idx, pass int
	var temp Sampah

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].Jumlah_Sampah > A[idx].Jumlah_Sampah {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp

		pass++
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d. ID: %d, Sampah: %s, Jumlah: %d kg, Jenis: %s, Daur Ulang: %s\n", i+1, A[i].ID, A[i].Jenis_Sampah, A[i].Jumlah_Sampah, A[i].sampah, A[i].Metode_Daur_Ulang)
	}
}

func SortDataMinInsert(A DataSampah, n int) {
	var pass, i int
	var temp Sampah
	pass = 1
	for pass < n {
		i = pass
		temp = A[pass]
		for i > 0 && temp.Jumlah_Sampah < A[i-1].Jumlah_Sampah {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d. ID: %d, Sampah: %s, Jumlah: %d kg, Jenis: %s, Daur Ulang: %s\n", i+1, A[i].ID, A[i].Jenis_Sampah, A[i].Jumlah_Sampah, A[i].sampah, A[i].Metode_Daur_Ulang)
	}
}

func SortDataJenisInset(A DataSampah, n int) {
	var pass, i int
	var temp Sampah
	pass = 1
	for pass < n {
		i = pass
		temp = A[pass]
		for i > 0 && temp.Jenis_Sampah < A[i-1].Jenis_Sampah {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d. ID: %d, Sampah: %s, Jumlah: %d kg, Jenis: %s, Daur Ulang: %s\n", i+1, A[i].ID, A[i].Jenis_Sampah, A[i].Jumlah_Sampah, A[i].sampah, A[i].Metode_Daur_Ulang)
	}
}

func StatistikData() {
	var Pilih int
	var x string
	fmt.Println()
	fmt.Println("==========Selamat Datang di Menu Statistik Data Sampah==========")
	fmt.Println("")
	fmt.Println("Apa yang ingin anda lihat ketahui?")
	fmt.Println("1. Jumlah Total Data Berat Sampah Yang Sudah Dikumpulkan")
	fmt.Println("2. Jumlah Total Data Sampah Yang di Daur Ulang")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih 1/2/3? ")
	fmt.Scan(&Pilih)
	fmt.Println()
	if Pilih == 1 {
		JumlahDataSampah(Data, nData)
	} else if Pilih == 2 {
		DaurUlang(&Data, nData)
	} else {
		return
	}
	fmt.Println()
	fmt.Print("Apakah kamu masih ingin melihat statistik data (Yes/No)? ")
	fmt.Scan(&x)
	if x == "Yes" || x == "yes" || x == "Y" || x == "y" {
		StatistikData()
		fmt.Println()
	}
}

func JumlahDataSampah(A DataSampah, n int) {
	var i, jumlah int
	jumlah = 0
	for i = 0; i < n; i++ {
		jumlah = jumlah + A[i].Jumlah_Sampah
	}
	fmt.Println("Jumlah Sampah Yang Berhasil Dikumpulkan Seberat ", jumlah, " kg")
}

func DaurUlang(A *DataSampah, n int) {
	var i, jumlah int
	var persenan float64
	jumlah = 0
	for i = 0; i < n; i++ {
		if A[i].Daur_Ulang == true {
			jumlah = jumlah + A[i].Jumlah_Sampah
		}
	}
	persenan = hitungPersen(Data, nData, jumlah)
	fmt.Println("Jumlah Sampah Yang Berhasil Di Daur Ulang Seberat ", jumlah, " kg")
	fmt.Printf("Sampah Yang di Daur Ulang %.2f%% Dari Data Yang Berhasil Dikumpulkan", persenan)
}

func hitungPersen(A DataSampah, n int, bagian int) float64 {
	var i int
	var Total int
	var persen float64
	for i = 0; i < n; i++ {
		Total += A[i].Jumlah_Sampah

	}
	persen = (float64(bagian) / float64(Total)) * 100
	return persen
}
