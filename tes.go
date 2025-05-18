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
	ID int
	Nama string
	Kategori string
	Poin int
	Tanggal string
	Frekuensi int // Menghitung berapa kali aktivitas dilakukan
}

// Struktur data untuk daftar aktivitas yang tersedia
type MasterAktivitas struct {
	Nama string
	Kategori string
	Poin int
	EmisiKarbon float64 // Gram CO2 yang berkurang per aktivitas
}

// Struktur untuk melacak frekuensi aktivitas
type AktivitasFrekuensi struct {
	Nama string
	Frekuensi int
}

// Variabel global
var daftarHabit [MAX]Habit
var jumlah int = 0
var frekuensiAktivitas [JML_AKTIVITAS]AktivitasFrekuensi

// Daftar aktivitas yang tersedia
var master [JML_AKTIVITAS]MasterAktivitas = [JML_AKTIVITAS]MasterAktivitas{
	{"Naik Sepeda", "Transportasi", 10, 250.0},
	{"Gunakan Botol Isi Ulang", "Konsumsi", 7, 80.0},
	{"Matikan Lampu Tak Terpakai", "Energi", 5, 45.0},
	{"Bawa Tas Belanja Sendiri", "Konsumsi", 8, 30.0},
	{"Pisahkan Sampah", "Limbah", 9, 100.0},
	{"Tanam Pohon", "Lingkungan", 15, 500.0},
	{"Hemat Air", "Energi", 6, 40.0},
	{"Gunakan Energi Terbarukan", "Energi", 12, 300.0},
	{"Kurangi Kendaraan Pribadi", "Transportasi", 14, 400.0},
	{"Beli Produk Lokal", "Konsumsi", 10, 150.0},
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
		fmt.Printf("%d. %s | %s | %d poin | %.1f g CO2 berkurang\n", 
			i+1, master[i].Nama, master[i].Kategori, master[i].Poin, master[i].EmisiKarbon)
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

// Binary Search untuk mencari aktivitas master
func cariAktivitas() {
	// Untuk binary search, array harus diurutkan terlebih dahulu
	var tempMaster [JML_AKTIVITAS]MasterAktivitas
	var i int
	for i = 0; i < JML_AKTIVITAS; i++ {
		tempMaster[i] = master[i]
	}
	
	// Bubble sort untuk mengurutkan master berdasarkan nama
	var j int
	for i = 0; i < JML_AKTIVITAS-1; i++ {
		for j = 0; j < JML_AKTIVITAS-i-1; j++ {
			if tempMaster[j].Nama > tempMaster[j+1].Nama {
				var temp MasterAktivitas
				temp = tempMaster[j]
				tempMaster[j] = tempMaster[j+1]
				tempMaster[j+1] = temp
			}
		}
	}
	
	var cari string
	fmt.Print("Masukkan nama aktivitas yang ingin dicari: ")
	fmt.Scan(&cari)
	cari = strings.ToLower(cari)
	
	// Binary Search
	var left int = 0
	var right int = JML_AKTIVITAS-1
	var found bool = false
	var mid int
	
	for left <= right {
		mid = (left + right) / 2
		
		if strings.ToLower(tempMaster[mid].Nama) == cari {
			fmt.Printf("\nAktivitas ditemukan: %s | %s | %d poin | %.1f g CO2 berkurang\n", 
				tempMaster[mid].Nama, tempMaster[mid].Kategori, 
				tempMaster[mid].Poin, tempMaster[mid].EmisiKarbon)
			found = true
			break
		} else if strings.ToLower(tempMaster[mid].Nama) < cari {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	if !found {
		fmt.Println("Aktivitas tidak ditemukan.")
	}
}

// Fungsi untuk mengambil bulan dan tahun dari string tanggal dd-mm-yyyy
func getBulanTahun(tanggal string) (int, int, error) {
	var parts []string
	parts = strings.Split(tanggal, "-")
	if len(parts) != 3 {
		return 0, 0, fmt.Errorf("format tanggal salah")
	}
	
	var bulan int
	var err error
	_, err = fmt.Sscan(parts[1], &bulan)
	if err != nil {
		return 0, 0, err
	}
	
	var tahun int
	_, err = fmt.Sscan(parts[2], &tahun)
	if err != nil {
		return 0, 0, err
	}
	
	return bulan, tahun, nil
}

// Laporan bulanan
func laporanBulanan() {
	if jumlah == 0 {
		fmt.Println("Belum ada data habit.")
		return
	}
	
	var bulanInput int
	var tahunInput int
	
	fmt.Print("Masukkan bulan (1-12): ")
	fmt.Scan(&bulanInput)
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahunInput)
	
	if bulanInput < 1 || bulanInput > 12 {
		fmt.Println("Bulan tidak valid.")
		return
	}
	
	// Statistik
	var totalPoin int = 0
	var totalKarbon float64 = 0
	var aktivitasPerKategori map[string]int
	aktivitasPerKategori = make(map[string]int)
	var aktivitasDilakukan []string
	aktivitasDilakukan = make([]string, 0)
	
	// Menghitung statistik dari data yang ada
	var i int
	var j int
	var bulan int
	var tahun int
	var err error
	for i = 0; i < jumlah; i++ {
		bulan, tahun, err = getBulanTahun(daftarHabit[i].Tanggal)
		if err != nil {
			continue
		}
		
		if bulan == bulanInput && tahun == tahunInput {
			totalPoin = totalPoin + daftarHabit[i].Poin
			
			// Cari emisi karbon yang berkurang untuk aktivitas ini
			for j = 0; j < JML_AKTIVITAS; j++ {
				if daftarHabit[i].Nama == master[j].Nama {
					totalKarbon = totalKarbon + master[j].EmisiKarbon
					break
				}
			}
			
			aktivitasPerKategori[daftarHabit[i].Kategori] = aktivitasPerKategori[daftarHabit[i].Kategori] + 1
			aktivitasDilakukan = append(aktivitasDilakukan, daftarHabit[i].Nama)
		}
	}
	
	// Nama bulan
	var namaBulan []string
	namaBulan = []string{
		"", "Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}
	
	fmt.Printf("\n===== LAPORAN BULANAN GREEN HABIT =====\n")
	fmt.Printf("Periode: %s %d\n\n", namaBulan[bulanInput], tahunInput)
	
	if totalPoin == 0 {
		fmt.Println("Tidak ada aktivitas yang tercatat pada bulan ini.")
		return
	}
	
	fmt.Printf("Total Poin: %d\n", totalPoin)
	fmt.Printf("Total Pengurangan Emisi CO2: %.2f gram (%.2f kg)\n", totalKarbon, totalKarbon/1000)
	
	fmt.Println("\nAktivitas per Kategori:")
	var kategori string
	var jumlahKategori int
	for kategori, jumlahKategori = range aktivitasPerKategori {
		fmt.Printf("- %s: %d aktivitas\n", kategori, jumlahKategori)
	}
	
	// Rekomendasi berdasarkan aktivitas yang belum dilakukan
	fmt.Println("\nRekomendasi untuk Meningkatkan Jejak Lingkungan:")
	
	// Mencari aktivitas yang belum dilakukan
	var aktivitasBelumDilakukan []MasterAktivitas
	aktivitasBelumDilakukan = make([]MasterAktivitas, 0)
	var m MasterAktivitas
	var aktivitas string
	var found bool
	for _, m = range master {
		found = false
		for _, aktivitas = range aktivitasDilakukan {
			if m.Nama == aktivitas {
				found = true
				break
			}
		}
		if !found {
			aktivitasBelumDilakukan = append(aktivitasBelumDilakukan, m)
		}
	}
	
	// Menampilkan rekomendasi
	if len(aktivitasBelumDilakukan) > 0 {
		fmt.Println("Aktivitas yang bisa Anda coba bulan depan:")
		for i = 0; i < len(aktivitasBelumDilakukan) && i < 3; i++ {
			fmt.Printf("- %s (%s): +%d poin, -%.1f g CO2\n", 
				aktivitasBelumDilakukan[i].Nama, aktivitasBelumDilakukan[i].Kategori, 
				aktivitasBelumDilakukan[i].Poin, aktivitasBelumDilakukan[i].EmisiKarbon)
		}
	} else {
		fmt.Println("Selamat! Anda telah mencoba semua jenis aktivitas ramah lingkungan.")
		fmt.Println("Teruslah konsisten dan tingkatkan frekuensinya!")
	}
	
	// Estimasi jejak karbon secara keseluruhan
	fmt.Printf("\nEstimasi Jejak Karbon Bulan Ini: %.2f kg CO2\n", totalKarbon/1000)
	fmt.Println("(Diperkirakan dari aktivitas ramah lingkungan yang Anda lakukan)")
	
	// Tips untuk bulan depan
	fmt.Println("\nTips untuk Bulan Depan:")
	
	// Cek kategori yang paling sedikit dilakukan
	var kategoriMinimal string = ""
	var jumlahMinimal int = 999
	
	for kategori, jumlahKategori = range aktivitasPerKategori {
		if jumlahKategori < jumlahMinimal {
			jumlahMinimal = jumlahKategori
			kategoriMinimal = kategori
		}
	}
	
	if kategoriMinimal != "" {
		fmt.Printf("- Fokus pada aktivitas kategori '%s' yang masih jarang dilakukan\n", kategoriMinimal)
	}
	
	fmt.Println("- Coba konsisten melakukan minimal satu aktivitas ramah lingkungan setiap hari")
	fmt.Println("- Ajak teman atau keluarga untuk bergabung dalam gerakan ramah lingkungan")
}

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
		fmt.Println("9. Cari Aktivitas Berdasarkan Kategori") //Sequential Search
		fmt.Println("10. Cari Aktivitas") //Binary Search
		fmt.Println("11. Laporan Bulanan Jejak Karbon")
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
		case 11:
			laporanBulanan()
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
