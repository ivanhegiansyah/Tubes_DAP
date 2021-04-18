package main
import (
	"fmt"
	"strconv"
)
	type transkip struct {
		namaMHS string
		nim string
		kelas string
		matkul [50]string
		nilai string
		
	}
	type MK2 struct {
	namaMK2 string
	sks int
	uts float64
	uas float64
	kuis float64
	nilairata float64
	}
	type mahasiswa struct {
		namaMHS string
		nim string
		kelas string
		matkul [50]MK2
		jummatkul int
		
		ipk float64
		
	}
	type arrMHS [100]mahasiswa 
	type arrTranskip [100]transkip

	var siswa arrMHS
	var jumlahmhs int

	var i, j, c int

	var tampung float64
	
func main() {
	var menu int
	var quit bool
	quit = false

	for !quit {

		fmt.Println("=======================================================")
		fmt.Println("		     Aplikasi Nilai					")
		fmt.Println("			Mahasiswa				")
		fmt.Println("=======================================================")
		fmt.Println()
		fmt.Println("			MENU")
		fmt.Println()
		fmt.Println("1.Mahasiswa dan Mata Kuliah")
		fmt.Println("2.Transkip Nilai")
		fmt.Println("3.Keluar")
		fmt.Println()
		fmt.Println("=======================================================")
		fmt.Print("Silahkan pilih menu: ")
		fmt.Scan(&menu)
		
		 
			if menu == 1 {
				submenu1()
			} else if menu == 2 {
				submenu2()
			} else if menu == 3 {
				quit = true 
			} else {
				fmt.Println("Pilihan Tidak Valid, Siliahkan Pilih Kembali")
			}
		}
	
}
func submenu1() {
	var pil int
	//var kunci string
	fmt.Println("=======================================================")
	fmt.Println()
	fmt.Println("1.Masukkan Data Mahasiswa Beserta Mata Kuliahnya")
	fmt.Println("2.Edit Data Mahasiswa")
	fmt.Println("3.Hapus Data Mahasiswa")
	fmt.Println("4.Edit Data Mata Kuliah")
	fmt.Println("5.Hapus Data Mata Kuliah")
	fmt.Println("6.Cari Mahasiswa yang Mengambil Mata Kuliah Tertentu")
	fmt.Println("7.Cari Mata Kuliah yang Diambil oleh Mahasiswa Tertentu")
	fmt.Println("8.Daftar Mahasiswa Berdasarkan Nilai")
	fmt.Println("9.Daftar Mahasiswa Berdasarkan SKS")
	fmt.Println("10.IPK Tertinggi dari Seluruh Mahasiswa")
	fmt.Println("11.Lihat Seluruh Data Mahasiswa serta Mata Kuliah yang Diambil")
	fmt.Println()
	fmt.Println("=======================================================")
	fmt.Println()
	fmt.Print("Pilihan : ")
	fmt.Scan(&pil)
	
	switch pil {
		case 1 : inMhs()
		case 2 : editMhs(jumlahmhs)
		case 3 : deleteMhs(jumlahmhs)
		case 4 : editMK(jumlahmhs)
		case 5 : deleteMK(jumlahmhs)
		case 6 : cariMhs(jumlahmhs)
		case 7 : cariMk(jumlahmhs)
		case 8 : sortNilai(jumlahmhs)
		case 9 : sortSKS2(jumlahmhs)
		case 10: ipkMax(jumlahmhs)
		case 11: viewMhs(jumlahmhs)
		default : 
		fmt.Println("Pilihan Tidak Valid, Siliahkan Pilih Kembali")
			submenu1()
	}
}

// memunculkan transkip nilai 
func submenu2() { 
	var g, c int
	g = 0
	
	for g < jumlahmhs {
		fmt.Println("Nama Mahasiswa: ", siswa[g].namaMHS)
		fmt.Println("NIM : ", siswa[g].nim)
		fmt.Println("Kelas :", siswa[g].kelas)
		c = 0
		for c < siswa[g].jummatkul {
			fmt.Println("mata kuliah ke-", c+1)
			fmt.Println("mata kuliah yang diambil oleh mahasiswa: ", siswa[g].matkul[c].namaMK2)
			fmt.Println(" SKS :", siswa[g].matkul[c].sks)
			fmt.Printf("Nilai total : %.2f \n", siswa[g].matkul[c].nilairata)
			fmt.Println(hitungGrade(g,c))
			c++
		}
		fmt.Println("Jumlah SKS : ", hitungSKS(g))
		fmt.Printf("IPK : %.2f \n ", hitungIPK(g))
		g++
	}

}

func inMhs() {
	var maha mahasiswa
	
	jumlahmhs = 0
	fmt.Print("Masukkan Jumlah Mahasiswa: ")
	fmt.Scan(&jumlahmhs)
	sampe := jumlahmhs+c
	for c < sampe {
		fmt.Println("=======================================================")
		fmt.Println()
		fmt.Println("Mahasiswa ke -",c+1)
		fmt.Println()	
		fmt.Print("Nama Mahasiswa: ")	
		fmt.Scan(&maha.namaMHS)
		fmt.Print("NIM: ")
		fmt.Scan(&maha.nim)	
		fmt.Print("Kelas: ")
		fmt.Scan(&maha.kelas)
		fmt.Print("Jumlah Mata Kuliah yang Diambil: ")
		fmt.Scan(&maha.jummatkul)
		
		x := 0
		for x < maha.jummatkul {
			fmt.Println("mata kuliah ke-", x+1)
			fmt.Print("Masukkan Mata Kuliah yang Diambil oleh Mahasiswa: ")
			fmt.Scan(&maha.matkul[x].namaMK2)
			fmt.Print("Masukkan SKS: ")
			errhandling1(1,5,"Masukkan SKS: ",&tampung)
			maha.matkul[x].sks = int(tampung)
			fmt.Print("Masukkan Nilai UTS: ")
			errhandling1(0,100,"Masukkan Nilai UTS: ",&maha.matkul[x].uts)
			fmt.Print("Masukkan Nilai UAS: ")
			errhandling1(0,100,"Masukkan Nilai UAS: ",&maha.matkul[x].uas)
			fmt.Print("Masukkan Nilai Kuis: ")
			errhandling1(0,100,"Masukkan Nilai Kuis: ",&maha.matkul[x].kuis)
			
			siswa[c].matkul[x].namaMK2 = maha.matkul[x].namaMK2
			siswa[c].matkul[x].sks = maha.matkul[x].sks
			siswa[c].matkul[x].uts = maha.matkul[x].uts
			siswa[c].matkul[x].uas = maha.matkul[x].uas
			siswa[c].matkul[x].kuis = maha.matkul[x].kuis
			siswa[c].matkul[x].nilairata = hitungNilai(c, x)
			x++
		}
		siswa[c].namaMHS = maha.namaMHS
		siswa[c].nim = maha.nim
		siswa[c].kelas = maha.kelas
		siswa[c].jummatkul = maha.jummatkul
		
		siswa[c].ipk = hitungIPK(c)
		c++
	}
	jumlahmhs = c
}

func viewMhs(jumlahmhs int) {
	var q, c int
	q = 0
	for q < jumlahmhs {
		fmt.Println("Nama Mahasiswa: ", siswa[q].namaMHS)
		fmt.Println("NIM : ", siswa[q].nim)
		fmt.Println("Kelas :", siswa[q].kelas)
		c = 0
		for c < siswa[q].jummatkul {
			fmt.Println("mata kuliah ke-", c+1)
			fmt.Println("mata kuliah yang diambil oleh mahasiswa: ", siswa[q].matkul[c].namaMK2)
			fmt.Println("Jumlah SKS :", siswa[q].matkul[c].sks)
			fmt.Println("Nilai UTS :", siswa[q].matkul[c].uts)
			fmt.Println("Nilai UAS :", siswa[q].matkul[c].uas)
			fmt.Println("Nilai kuis :", siswa[q].matkul[c].kuis)
			fmt.Printf("Nilai total : %.2f \n", siswa[q].matkul[c].nilairata)
			c++
		}
		q++
	}
}

//cari dengan metode sekuensial
func cariMhs(jumlahmhs int) {
	var key string
	
	fmt.Print("masukkan nama mata kuliah yang ingin dicari: ")
	fmt.Scan(&key)
	for z := 0;z < jumlahmhs;z++ {	
		for y := 0;y < siswa[z].jummatkul;y++ {
			if siswa[z].matkul[y].namaMK2 == key {
				fmt.Println("Nama Siswa :")
				fmt.Println(siswa[z].namaMHS)
			} else {
				fmt.Println("Tidak ada mahasiswa yang mengambil mata kuliah tersebut")
			}
		}
	}
} 

//cari dengan metode binary
func cariMk(jumlahmhs int) {
	
	w:=0
	sortMhs(jumlahmhs)
	var key string
	
	fmt.Print("Masukkan Nama Mahasiswa: ")
	fmt.Scan(&key)
	
	var akhir, tengah int
	awal := 0
	akhir = jumlahmhs
	tengah = (awal+akhir) / 2
	for siswa[tengah].namaMHS != key && awal < akhir {
		if key > siswa[tengah].namaMHS {
			awal = tengah + 1
		} else {
			akhir = tengah - 1
		}
		tengah = (awal+akhir) / 2
	}
	if siswa[tengah].namaMHS == key {
		fmt.Println("Mata Kuliah yang diambil : ")
		for w < siswa[tengah].jummatkul {
			fmt.Println(siswa[tengah].matkul[w].namaMK2)
			w++
		}
	} else {
		fmt.Println("Mahasiswa tersebut tidak mengambil Mata Kuliah apapun")
	}

}

// cari NIM 
func cariNIM(jumlahmhs int) int {
	var key string

	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&key)
	for z := 0;z < jumlahmhs;z++ {	
		if siswa[z].nim == key {
			return z
		} 
	}
	return -1
}

//hitung nilai
func hitungNilai(idxMHS, idxMK int) float64 {
	var nilaiUTS, nilaiUAS, nilaiKUIS, total float64
	nilaiUTS = siswa[idxMHS].matkul[idxMK].uts * 0.4
	nilaiUAS = siswa[idxMHS].matkul[idxMK].uas * 0.4
	nilaiKUIS = siswa[idxMHS].matkul[idxMK].kuis * 0.2
	total = (nilaiUTS + nilaiUAS + nilaiKUIS)  / 25
	return total

}

//hitung ipk
func hitungIPK(idxN int) float64 {
	var totalmutu, mutu float64
	var totalsks int
	z := 0
	totalmutu = 0
	totalsks = 0
	
	for z < siswa[idxN].jummatkul {
		mutu = siswa[idxN].matkul[z].nilairata * float64(siswa[idxN].matkul[z].sks)
		totalmutu = totalmutu + mutu
		totalsks = siswa[idxN].matkul[z].sks + totalsks
		z++
		
	}
	return totalmutu / float64(totalsks)
	
}

// hitung total sks
func hitungSKS(idxS int) int {
	var jumlahsks int
	jumlahsks = 0
	for w :=0; w < siswa[idxS].jummatkul; w++ {
		jumlahsks = jumlahsks + siswa[idxS].matkul[w].sks
		
	}
	return jumlahsks
}

// hitung grade
func hitungGrade(idxM, idxMK int) string {
	if siswa[idxM].matkul[idxMK].nilairata > 4.00 {
		return "Tidak termasuk"
	} else if siswa[idxM].matkul[idxMK].nilairata > 3.50 {
		return "Nilai mata kuliah : A"
	} else if siswa[idxM].matkul[idxMK].nilairata > 3.00 {
		return "Nilai mata kuliah : AB"
	} else if siswa[idxM].matkul[idxMK].nilairata > 3.00 {
		return "Nilai mata kuliah : B"
	} else if siswa[idxM].matkul[idxMK].nilairata > 2.50 {
		return "Nilai mata kuliah : BC"
	} else if siswa[idxM].matkul[idxMK].nilairata > 2.00 {
		return "Nilai mata kuliah : C"
	} else if siswa[idxM].matkul[idxMK].nilairata > 1.00 {
		return "Nilai mata kuliah : D"
	} else if siswa[idxM].matkul[idxMK].nilairata <= 0.99 {
		return "Nilai mata kuliah : E"
	} else {
		return "tidak termasuk"
	}
}

//sorting selection berdasarkan ipk
func sortNilai(jumlahmhs int) {
	
	for m:= 0; m < jumlahmhs ; m++ {
		idx := m
		for n := m + 1; n < jumlahmhs; n++ {
			if siswa[n].ipk < siswa[idx].ipk {
				idx = n
			}
		}
		temp := siswa[m]
		siswa[m]= siswa[idx]
		siswa[idx] = temp
	}
	for h :=0; h < jumlahmhs; h++ {
		fmt.Printf("IPK : %.2f \n", siswa[h].ipk)
		fmt.Println("Nama siswa :", siswa[h].namaMHS)
	}
		fmt.Println()
}

//sorting selection nama mahasiswa
func sortMhs(jumlahmhs int) {
	
	for m:= 0; m < jumlahmhs ; m++ {
		idx := m
		for n := m + 1; n < jumlahmhs; n++ {
			if siswa[n].namaMHS < siswa[idx].namaMHS {
				idx = n
			}
		}
		temp := siswa[m]
		siswa[m]= siswa[idx]
		siswa[idx] = temp
	}
}

func sortSKS2(jumlahmhs int) {
	 for i := 0; i < jumlahmhs-1; i++ {
   		j := i+1
    	for j > 0 && hitungSKS(j-1) < hitungSKS(j) {
			temp := siswa[j]
			siswa[j] = siswa[j-1]
			siswa[j-1] = temp
     		j--
    	}
	}

	for i:=0; i < jumlahmhs; i++ {
		fmt.Println("Nama Siswa: ", siswa[i].namaMHS)
		fmt.Println("Total SKS: ", hitungSKS(i))
	}
}
	
  


// menampilkan IPK tertinggi
func ipkMax(jumlahmhs int) {
	
	
	var max float64
	var t, idxmax int
	t = 0
	
	max = siswa[t].ipk

	for t < jumlahmhs {
		if siswa[t].ipk > max {
			max = siswa[t].ipk
			idxmax = t
		}
		t++
	}
	fmt.Printf("IPK : %.2f \n", max)
	fmt.Println("Nama siswa :", siswa[idxmax].namaMHS)
	
	fmt.Println()
}

// fungsi delete mahasiswa
func deleteMhs(jummhs int) {
	idxM := cariNIM(jummhs)
	if idxM != -1 {
		for i :=idxM; i < jummhs; i++ {
			siswa[i] = siswa[i+1]
		}
		jumlahmhs--
		fmt.Println("BERHASIL DIHAPUS")
	} else {
		fmt.Println("Tidak ada")
	}
}

// fungsi delete mata kuliah
func deleteMK(jumlahmhs int) {
	var idxMK, idxM int

	idxM = cariNIM(jumlahmhs)
	if idxM != -1 {
		fmt.Println("matakuliah yang di ambil : ")
		for i:=0; i<siswa[idxM].jummatkul; i++ {
			fmt.Printf("%d. %s\n", i+1, siswa[idxM].matkul[i].namaMK2)
		}
		fmt.Print("Pilihan: ")
		fmt.Scan(&idxMK)

		for i :=idxMK-1; i < siswa[idxM].jummatkul; i++ {
			siswa[idxM].matkul[i] = siswa[idxM].matkul[i+1]
		}
		siswa[idxM].jummatkul--
		fmt.Println("BERHASIL DIHAPUS")
	} else {
		fmt.Println("Tidak ada")
	}
}

// edit mahasiswa
func editMhs(jumlahmhs int) {
	idxMhs := cariNIM(jumlahmhs)
	if idxMhs != -1 {
		fmt.Print("Nama Mahasiswa: ")
		fmt.Scan(&siswa[idxMhs].namaMHS)
		fmt.Print("NIM : ")
		fmt.Scan(&siswa[idxMhs].nim)
		fmt.Print("Kelas: ")
		fmt.Scan(&siswa[idxMhs].kelas)
	} else  {
		fmt.Println("tidak ada")
	}

}

// edit mata kuliah
func editMK(jumlahmhs int) {
	var idxMK int
	idxMhs := cariNIM(jumlahmhs)
	if idxMhs != -1 {
		fmt.Println("matakuliah yang di ambil : ")
		for i:=0; i<siswa[idxMhs].jummatkul; i++ {
			fmt.Printf("%d. %s\n", i+1, siswa[idxMhs].matkul[i].namaMK2)
		}
		fmt.Print("Pilihan: ")
		fmt.Scan(&idxMK)
		fmt.Print("Masukkan Mata Kuliah yang Diambil oleh Mahasiswa: ")
		fmt.Scan(&siswa[idxMhs].matkul[idxMK-1].namaMK2)
		fmt.Print("Masukkan Jumlah SKS: ")
		fmt.Scan(&siswa[idxMhs].matkul[idxMK-1].sks)
		fmt.Print("Masukkan Nilai UTS: ")
		fmt.Scan(&siswa[idxMhs].matkul[idxMK-1].uts)
		fmt.Print("Masukkan Nilai UAS: ")
		fmt.Scan(&siswa[idxMhs].matkul[idxMK-1].uas)
		fmt.Print("Masukkan Nilai Kuis: ")
		fmt.Scan(&siswa[idxMhs].matkul[idxMK-1].kuis)
	} else {
		fmt.Println("tidak ada")
	}
}
	
func errhandling1(bawah, atas float64, print string,  out *float64) {
	var angka string

	fmt.Scan(&angka)
	value, err := strconv.ParseFloat(angka, 64)
	if value > atas || value < bawah || err != nil {
		fmt.Println("Data tersebut tidak bisa dimasukkan")
	} else {
		*out = value
	}
	for err != nil || value > atas || value < bawah {

		fmt.Print(print)
		fmt.Scan(&angka)
		value, err = strconv.ParseFloat(angka, 64)
		if value > atas || value < bawah || err != nil {
			fmt.Println("Data tersebut tidak bisa dimasukkan")
		} else {
			*out = value
		}
	}
	
}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	