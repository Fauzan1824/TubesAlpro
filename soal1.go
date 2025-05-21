package main

import (
	"fmt"
	"strings"
)

// Konstanta untuk program
const MAX = 100
const JML_AKTIVITAS = 10

// Struktur data untuk menyimpan informasi habit
type Habit struct {
	ID        int
	Nama      string
	Kategori  string
	Poin      int
	Tanggal   string
	Frekuensi int // Menghitung berapa kali aktivitas dilakukan
}

// Struktur data untuk daftar aktivitas yang tersedia
type MasterAktivitas struct {
	Nama     string
	Kategori string
	Poin     int
}

// Struktur untuk melacak frekuensi aktivitas
type AktivitasFrekuensi struct {
	Nama      string
	Frekuensi int
}

// Variabel global
var daftarHabit [MAX]Habit
var jumlah int = 0
var frekuensiAktivitas [JML_AKTIVITAS]AktivitasFrekuensi

// Daftar aktivitas yang tersedia
var master [JML_AKTIVITAS]MasterAktivitas = [JML_AKTIVITAS]MasterAktivitas{
	{"Naik Sepeda", "Transportasi", 10},
	{"Gunakan Botol Isi Ulang", "Konsumsi", 7},
	{"Matikan Lampu Tak Terpakai", "Energi", 5},
	{"Bawa Tas Belanja Sendiri", "Konsumsi", 8},
	{"Pisahkan Sampah", "Limbah", 9},
	{"Tanam Pohon", "Lingkungan", 15},
	{"Hemat Air", "Energi", 6},
	{"Gunakan Energi Terbarukan", "Energi", 12},
	{"Kurangi Kendaraan Pribadi", "Transportasi", 14},
	{"Beli Produk Lokal", "Konsumsi", 10},
}

// Inisialisasi nama aktivitas di struktur frekuensi
func init() {
	var i int
	for i = 0; i < JML_AKTIVITAS; i++ {
		frekuensiAktivitas[i].Nama = master[i].Nama
		frekuensiAktivitas[i].Frekuensi = 0
	}
}

// Menampilkan daftar aktivitas yang tersedia
func tampilMaster() {
	fmt.Println("\n--- Daftar Aktivitas Green Habit ---")
	var i int
	for i = 0; i < JML_AKTIVITAS; i++ {
		fmt.Printf("%d. %s | %s | %d poin\n",
			i+1, master[i].Nama, master[i].Kategori, master[i].Poin)
	}
}

// Menambahkan aktivitas ramah lingkungan baru
func tambahHabit() {
	if jumlah >= MAX {
		fmt.Println("Data penuh!")
		return
	}

	tampilMaster()
	var pilihan int
	fmt.Print("Pilih aktivitas (1-", JML_AKTIVITAS, "): ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > JML_AKTIVITAS {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	var tanggal string
	fmt.Print("Masukkan tanggal (dd-mm-yyyy): ")
	fmt.Scan(&tanggal)

	var h Habit
	h.ID = jumlah + 1
	h.Nama = master[pilihan-1].Nama
	h.Kategori = master[pilihan-1].Kategori
	h.Poin = master[pilihan-1].Poin
	h.Tanggal = tanggal
	h.Frekuensi = 1

	daftarHabit[jumlah] = h
	jumlah = jumlah + 1

	// Update frekuensi aktivitas
	frekuensiAktivitas[pilihan-1].Frekuensi = frekuensiAktivitas[pilihan-1].Frekuensi + 1

	fmt.Println("Habit berhasil dicatat!")
}

// Mengubah aktivitas yang sudah ada
func ubahHabit() {
	if jumlah == 0 {
		fmt.Println("Belum ada data habit.")
		return
	}

	tampilHabit()
	var id int
	fmt.Print("Masukkan ID habit yang ingin diubah: ")
	fmt.Scan(&id)

	// Cari indeks dari habit dengan ID yang dimaksud
	var idx int = -1
	var i int
	for i = 0; i < jumlah; i++ {
		if daftarHabit[i].ID == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println("ID habit tidak ditemukan!")
		return
	}

	// Mengurangi frekuensi aktivitas lama
	for i = 0; i < JML_AKTIVITAS; i++ {
		if frekuensiAktivitas[i].Nama == daftarHabit[idx].Nama {
			frekuensiAktivitas[i].Frekuensi = frekuensiAktivitas[i].Frekuensi - 1
			break
		}
	}

	// Pilih aktivitas baru
	tampilMaster()
	var pilihan int
	fmt.Print("Pilih aktivitas baru (1-", JML_AKTIVITAS, "): ")
	fmt.Scan(&pilihan)

	if pilihan < 1 || pilihan > JML_AKTIVITAS {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	var tanggal string
	fmt.Print("Masukkan tanggal baru (dd-mm-yyyy): ")
	fmt.Scan(&tanggal)

	// Update data habit
	daftarHabit[idx].Nama = master[pilihan-1].Nama
	daftarHabit[idx].Kategori = master[pilihan-1].Kategori
	daftarHabit[idx].Poin = master[pilihan-1].Poin
	daftarHabit[idx].Tanggal = tanggal

	// Update frekuensi untuk aktivitas baru
	frekuensiAktivitas[pilihan-1].Frekuensi = frekuensiAktivitas[pilihan-1].Frekuensi + 1

	fmt.Println("Habit berhasil diubah!")
}

// Menghapus aktivitas yang sudah ada
func hapusHabit() {
	if jumlah == 0 {
		fmt.Println("Belum ada data habit.")
		return
	}

	tampilHabit()
	var id int
	fmt.Print("Masukkan ID habit yang ingin dihapus: ")
	fmt.Scan(&id)

	// Cari indeks dari habit dengan ID yang dimaksud
	var idx int = -1
	var i int
	for i = 0; i < jumlah; i++ {
		if daftarHabit[i].ID == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		fmt.Println("ID habit tidak ditemukan!")
		return
	}

	// Mengurangi frekuensi aktivitas
	for i = 0; i < JML_AKTIVITAS; i++ {
		if frekuensiAktivitas[i].Nama == daftarHabit[idx].Nama {
			frekuensiAktivitas[i].Frekuensi = frekuensiAktivitas[i].Frekuensi - 1
			break
		}
	}

	// Hapus data dengan menggeser seluruh array
	for i = idx; i < jumlah-1; i++ {
		daftarHabit[i] = daftarHabit[i+1]
	}
	jumlah = jumlah - 1

	// Perbarui ID
	for i = idx; i < jumlah; i++ {
		daftarHabit[i].ID = i + 1
	}

	fmt.Println("Habit berhasil dihapus!")
}

// Menampilkan riwayat aktivitas
func tampilHabit() {
	if jumlah == 0 {
		fmt.Println("Belum ada data habit.")
		return
	}
	fmt.Println("\n--- Riwayat Green Habit ---")
	var i int
	for i = 0; i < jumlah; i++ {
		fmt.Printf("[%d] %s | %s | %d poin | %s\n",
			daftarHabit[i].ID, daftarHabit[i].Nama, daftarHabit[i].Kategori,
			daftarHabit[i].Poin, daftarHabit[i].Tanggal)
	}
}

// Menghitung total poin
func totalPoin() {
	var total int = 0
	var i int
	for i = 0; i < jumlah; i++ {
		total = total + daftarHabit[i].Poin
	}
	fmt.Println("Total poin ramah lingkungan Anda:", total)
}

// Mengurutkan aktivitas berdasarkan poin menggunakan Selection Sort
func urutPoin(naik bool) {
	var i int
	var j int
	var idx int
	for i = 0; i < jumlah-1; i++ {
		idx = i
		for j = i + 1; j < jumlah; j++ {
			if (naik && daftarHabit[j].Poin < daftarHabit[idx].Poin) ||
				(!naik && daftarHabit[j].Poin > daftarHabit[idx].Poin) {
				idx = j
			}
		}
		var temp Habit
		temp = daftarHabit[i]
		daftarHabit[i] = daftarHabit[idx]
		daftarHabit[idx] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan poin.")
}

// Mengurutkan aktivitas berdasarkan tanggal menggunakan Insertion Sort
func urutTanggal(naik bool) {
	var i int
	var j int
	for i = 1; i < jumlah; i++ {
		var temp Habit
		temp = daftarHabit[i]
		j = i - 1
		for j >= 0 && ((naik && daftarHabit[j].Tanggal > temp.Tanggal) ||
			(!naik && daftarHabit[j].Tanggal < temp.Tanggal)) {
			daftarHabit[j+1] = daftarHabit[j]
			j = j - 1
		}
		daftarHabit[j+1] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan tanggal.")
}

// Mengurutkan aktivitas berdasarkan frekuensi menggunakan Selection Sort
func urutFrekuensi(naik bool) {
	// Selection Sort untuk frekuensi aktivitas
	var temp [JML_AKTIVITAS]AktivitasFrekuensi
	var i int
	for i = 0; i < JML_AKTIVITAS; i++ {
		temp[i] = frekuensiAktivitas[i]
	}

	var j int
	var idx int
	for i = 0; i < JML_AKTIVITAS-1; i++ {
		idx = i
		for j = i + 1; j < JML_AKTIVITAS; j++ {
			if (naik && temp[j].Frekuensi < temp[idx].Frekuensi) ||
				(!naik && temp[j].Frekuensi > temp[idx].Frekuensi) {
				idx = j
			}
		}
		var tempAkt AktivitasFrekuensi
		tempAkt = temp[i]
		temp[i] = temp[idx]
		temp[idx] = tempAkt
	}

	fmt.Println("\n--- Aktivitas Berdasarkan Frekuensi ---")
	for i = 0; i < JML_AKTIVITAS; i++ {
		if temp[i].Frekuensi > 0 {
			fmt.Printf("%s: %d kali\n", temp[i].Nama, temp[i].Frekuensi)
		}
	}
}

// Sequential Search untuk mencari habit berdasarkan kategori
func cariKategori() {
	var kategori string
	fmt.Print("Masukkan kategori yang ingin dicari: ")
	fmt.Scan(&kategori)

	var found bool = false
	fmt.Printf("\n--- Hasil Pencarian untuk Kategori '%s' ---\n", kategori)

	var i int
	for i = 0; i < jumlah; i++ {
		if strings.ToLower(daftarHabit[i].Kategori) == strings.ToLower(kategori) {
			fmt.Printf("[%d] %s | %s | %d poin | %s\n",
				daftarHabit[i].ID, daftarHabit[i].Nama, daftarHabit[i].Kategori,
				daftarHabit[i].Poin, daftarHabit[i].Tanggal)
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan habit dengan kategori tersebut.")
	}
}

// Binary Search untuk mencari aktivitas 
func cariAktivitas() {
	
// Menu utama aplikasi
func menu() {
	var pilih int = 0
	for pilih != 99 {
		fmt.Println("\n===== MENU GREEN HABIT =====")
		fmt.Println("1. Tambah Aktivitas")
		fmt.Println("2. Ubah Aktivitas")
		fmt.Println("3. Hapus Aktivitas")
		fmt.Println("4. Tampilkan Riwayat")
		fmt.Println("5. Lihat Total Poin")
		fmt.Println("6. Urutkan Berdasarkan Poin")
		fmt.Println("7. Urutkan Berdasarkan Tanggal")
		fmt.Println("8. Urutkan Berdasarkan Frekuensi")
		fmt.Println("9. Cari Aktivitas Berdasarkan Kategori") // Sequential Search
		fmt.Println("10. Cari Aktivitas")                     // Binary Search
		fmt.Println("99. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahHabit()
		case 2:
			ubahHabit()
		case 3:
			hapusHabit()
		case 4:
			tampilHabit()
		case 5:
			totalPoin()
		case 6:
			var opsi int
			fmt.Print("Naik (1) atau Turun (0): ")
			fmt.Scan(&opsi)
			urutPoin(opsi == 1)
		case 7:
			var opsi int
			fmt.Print("Naik (1) atau Turun (0): ")
			fmt.Scan(&opsi)
			urutTanggal(opsi == 1)
		case 8:
			var opsi int
			fmt.Print("Naik (1) atau Turun (0): ")
			fmt.Scan(&opsi)
			urutFrekuensi(opsi == 1)
		case 9:
			cariKategori()
		case 10:
			cariAktivitas()
		case 99:
			fmt.Println("Terima kasih telah membangun kebiasaan hijau bersama Green Habit!")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func main() {

	// Menampilkan tanggal hari ini
	fmt.Println("Mari mulai menambahkan aktivitas ramah lingkungan Anda!")

	menu()
}
