nilai =float(input("masukan nilai Akhir:"))

if nilai >= 80:
    print("Lulus")
    print("Predikat: Sangat Baik")
elif 71 <= nilai <= 85:
    print("Lulus")
    print("Predikat: Baik")
elif 60 <= nilai <= 70:
    print("Lulus")
    print("Predikat: Memuaskan")
else:
    print("Tidak Lulus")
