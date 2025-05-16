import pandas as pd

# Data awal
data_dict = {
    "jenis": [
        "Rumput Rottboellia", "Rumput cretica", "Rumput Aur-aur", "Rumput Gajah Mini", 
        "Rumput Teki", "Rumput Dellis", "Rumput Ceker Ayam", "Rumput Carex",
        "Rumput Lili Paris", "Rumput Briza Minor", "Rumput Oplismenus", 
        "Rumput Leersia Virginica", "Rumput Panicum Repens", "Rumput Elymus", "Rumput Vasey"
    ],
    "panjang": [34.1, 14.5, 4.5, 9.6, 23.5, 23.5, 12.8, 44.5, 19, 29.6, 10.8, 43, 50, 25, 34],
    "pendek": [21, 6, 2, 5, 10.2, 11, 11.7, 27, 5.8, 16.3, 5.5, 27.8, 25.5, 23.5, 29.5]
}

# Buat DataFrame
data = pd.DataFrame(data_dict)

# Tambah 200 kolom baru
for i in range(1, 201):
    data[f'kolom{i}'] = None

# Simpan ke Excel
data.to_excel('tugas_besar_dengan_200_kolom.xlsx', index=False)
print("File berhasil dibuat sebagai tugas_besar_dengan_200_kolom.xlsx")
