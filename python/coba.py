from time import sleep
import sys

def print_lirik():
    lines = [
        ("Aku yang jatuh cinta", 0.15),
        ("Haruskah kau kuberi kesempatan", 0.07),
        ("Ingin aku jadi kekasih yang baik", 0.07),
        ("Berikan aku kesempatan oh", 0.08),
        ("Tahukah dirimu?, Tahukah hatimu?", 0.06),
        ("Berulang kuketuk, aku mencintaimu", 0.08),
        ("Tapi dirimu tak pernah sadari", 0.05),
        ("Aku yang jatuh cinta", 0.10)
    ]
    
    delays={7.2, 3, 2.5, 7.5, 3.5, 4, 3.5, 3.5}
    
    for line, char_delay in lines:
        for char in line:
            print(char, end='')
            sys.stdout.flush()  # Untuk memaksa output tampil per karakter
            sleep(char_delay)
        print()  # Ganti baris setelah setiap kalimat

# Jalankan fungsi
print_lirik()
