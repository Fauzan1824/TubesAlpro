package main

import "fmt"

const MAX = 100
const JML_AKTIVITAS = 10

type Habit struct {
    ID       int
    Nama     string
    Kategori string
    Poin     int
    Tanggal  string
}

type MasterAktivitas struct {
    Nama     string
    Kategori string
    Poin     int
}

var daftarHabit [MAX]Habit
var jumlah int = 0

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

func tampilMaster() {
    fmt.Println("\n--- Daftar Aktivitas Green Habit ---")
    for i := 0; i < JML_AKTIVITAS; i++ {
        fmt.Printf("%d. %s | %s | %d poin\n", i+1, master[i].Nama, master[i].Kategori, master[i].Poin)
    }
}

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

    h := Habit{
        ID:       jumlah + 1,
        Nama:     master[pilihan-1].Nama,
        Kategori: master[pilihan-1].Kategori,
        Poin:     master[pilihan-1].Poin,
        Tanggal:  tanggal,
    }

    daftarHabit[jumlah] = h
    jumlah++

    fmt.Println("Habit berhasil dicatat!")
}

func tampilHabit() {
    if jumlah == 0 {
        fmt.Println("Belum ada data habit.")
        return
    }
    fmt.Println("\n--- Riwayat Green Habit ---")
    for i := 0; i < jumlah; i++ {
        fmt.Printf("[%d] %s | %s | %d poin | %s\n",
            daftarHabit[i].ID, daftarHabit[i].Nama, daftarHabit[i].Kategori,
            daftarHabit[i].Poin, daftarHabit[i].Tanggal)
    }
}

func totalPoin() {
    var total int
    for i := 0; i < jumlah; i++ {
        total += daftarHabit[i].Poin
    }
    fmt.Println("Total poin ramah lingkungan Anda:", total)
}

func urutPoin(naik bool) {
    for i := 0; i < jumlah-1; i++ {
        for j := i + 1; j < jumlah; j++ {
            if (naik && daftarHabit[j].Poin < daftarHabit[i].Poin) ||
                (!naik && daftarHabit[j].Poin > daftarHabit[i].Poin) {
                daftarHabit[i], daftarHabit[j] = daftarHabit[j], daftarHabit[i]
            }
        }
    }
    fmt.Println("Data berhasil diurutkan berdasarkan poin.")
}

func urutTanggal(naik bool) {
    for i := 1; i < jumlah; i++ {
        temp := daftarHabit[i]
        j := i - 1
        for j >= 0 && ((naik && daftarHabit[j].Tanggal > temp.Tanggal) ||
            (!naik && daftarHabit[j].Tanggal < temp.Tanggal)) {
            daftarHabit[j+1] = daftarHabit[j]
            j--
        }
        daftarHabit[j+1] = temp
    }
    fmt.Println("Data berhasil diurutkan berdasarkan tanggal.")
}

func menu() {
    var pilih int
    for pilih != 99 {
        fmt.Println("\n===== MENU GREEN HABIT =====")
        fmt.Println("1. Tambah Aktivitas")
        fmt.Println("2. Tampilkan Riwayat")
        fmt.Println("3. Lihat Total Poin")
        fmt.Println("4. Urutkan Berdasarkan Poin")
        fmt.Println("5. Urutkan Berdasarkan Tanggal")
        fmt.Println("99. Keluar")
        fmt.Print("Pilih menu: ")
        fmt.Scan(&pilih)

        switch pilih {
        case 1:
            tambahHabit()
        case 2:
            tampilHabit()
        case 3:
            totalPoin()
        case 4:
            var opsi int
            fmt.Print("Naik (1) atau Turun (0): ")
            fmt.Scan(&opsi)
            urutPoin(opsi == 1)
        case 5:
            var opsi int
            fmt.Print("Naik (1) atau Turun (0): ")
            fmt.Scan(&opsi)
            urutTanggal(opsi == 1)
        case 99:
            fmt.Println("Terima kasih telah membangun kebiasaan hijau bersama Green Habit!")
        default:
            fmt.Println("Pilihan tidak valid.")
        }
    }
}

func main() {
    menu()
}
