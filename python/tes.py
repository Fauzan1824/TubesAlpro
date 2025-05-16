def atm_simulation():
    correct_pin = "123987"
    attempts = 0
    max_attempts = 3

    while attempts < max_attempts:
        pin_input = input("Masukkan PIN Anda: ")
        
        if pin_input == correct_pin:
            nominal = input("Masukkan nominal uang yang akan diambil: ")
            print(f"Anda berhasil mengambil uang sebesar: {nominal}")
            return
        else:
            attempts += 1
            print(f"PIN salah. Anda telah mencoba {attempts} kali.")
    
    print("PIN terblokir, silahkan kunjungi kantor cabang terdekat!")

atm_simulation()