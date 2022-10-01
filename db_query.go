package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var ReturnData ReturnLogin
		r.ParseForm()
		StatusLogin := "Username / Password Salah"
		var Username string
		var Password string
		_ = Password

		Username = r.FormValue("username")
		Username = strings.ToLower(Username)
		Password = r.FormValue("password")

		fmt.Println("user : " + Username)
		fmt.Println("pass : " + Password)

		query := "SELECT password FROM tb_user_login WHERE username='" + Username + "';"
		rows := DbQuery(query)

		var err error

		fmt.Println()
		// read hasil query
		for rows.Next() {
			var PassEncryp string
			rows.Scan(&PassEncryp)
			fmt.Println("origin db : " + PassEncryp)
			//cek password
			if CheckPassword(Password, PassEncryp) == true {
				StatusLogin = "Sukses"
				break
			}
		}

		if err = rows.Err(); err != nil {
			fmt.Println(err.Error())
			return
		}
		if StatusLogin == "Sukses" {
			ReturnData.StatusLogin = StatusLogin
			RHash, _ := randomHex(20)
			ReturnData.Hash = RHash
		} else {
			ReturnData.StatusLogin = StatusLogin
		}

		ReturnLogin, _ := json.Marshal(ReturnData)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(string(ReturnLogin))
	}
}
func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GetSiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var RtSiswaList []JsonSiswa
		var dum JsonSiswa
		//perintah sql
		query := "SELECT * FROM tb_siswa;"

		//eksekusi perintah sql
		ResultQuery := DbQuery(query)

		//baca hasil return data
		for ResultQuery.Next() {
			err := ResultQuery.Scan(&dum.Nis, &dum.Nama_siswa, &dum.JK, &dum.Kelas)
			if err != nil {
				ResultQuery.Close()
				break
			} else {
				RtSiswaList = append(RtSiswaList, dum)
			}
		}

		//tutup koneksi
		ResultQuery.Close()

		//return detail token
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(RtSiswaList)
	}
}

func GetGuru(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var RtList []JsonGuru
		var dum JsonGuru
		//perintah sql
		query := "SELECT * FROM tb_guru;"

		//eksekusi perintah sql
		ResultQuery := DbQuery(query)

		//baca hasil return data
		for ResultQuery.Next() {
			err := ResultQuery.Scan(&dum.Nip, &dum.Nama_guru, &dum.Tgl_lahir, &dum.Agama, &dum.Alamat, &dum.No_hp, &dum.Pendidikan, &dum.Foto)
			if err != nil {
				ResultQuery.Close()
				break
			} else {
				RtList = append(RtList, dum)
			}
		}

		//tutup koneksi
		ResultQuery.Close()

		//return detail token
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(RtList)
	}
}

func GetPrestasi(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var RtList []JsonPrestasi
		var dum JsonPrestasi
		//perintah sql
		query := "SELECT * FROM tb_prestasi;"

		//eksekusi perintah sql
		ResultQuery := DbQuery(query)

		//baca hasil return data
		for ResultQuery.Next() {
			err := ResultQuery.Scan(&dum.Nis, &dum.Nama_siswa, &dum.Jenis_prestasi, &dum.Nama_prestasi, &dum.Tahun_prestasi)
			if err != nil {
				ResultQuery.Close()
				break
			} else {
				RtList = append(RtList, dum)
			}
		}

		//tutup koneksi
		ResultQuery.Close()

		//return detail token
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(RtList)
	}
}

func GetMapel(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var RtList []JsonMapel
		var dum JsonMapel
		//perintah sql
		query := "SELECT * FROM tb_mapel;"

		//eksekusi perintah sql
		ResultQuery := DbQuery(query)

		//baca hasil return data
		for ResultQuery.Next() {
			err := ResultQuery.Scan(&dum.Kode_mapel, &dum.Nama_mapel)
			if err != nil {
				ResultQuery.Close()
				break
			} else {
				RtList = append(RtList, dum)
			}
		}

		//tutup koneksi
		ResultQuery.Close()

		//return detail token
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(RtList)
	}
}

func GetMapelDiampu(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var RtList []JsonMapelDiampu
		var dum JsonMapelDiampu
		//perintah sql
		query := "SELECT * FROM tb_mapel;"

		//eksekusi perintah sql
		ResultQuery := DbQuery(query)

		//baca hasil return data
		for ResultQuery.Next() {
			err := ResultQuery.Scan(&dum.Nip, &dum.Kode_mapel)
			if err != nil {
				ResultQuery.Close()
				break
			} else {
				RtList = append(RtList, dum)
			}
		}

		//tutup koneksi
		ResultQuery.Close()

		//return detail token
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(RtList)
	}
}

func GetBerita(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var RtList []JsonBerita
		var dum JsonBerita
		//perintah sql
		query := "SELECT * FROM tb_berita;"

		//eksekusi perintah sql
		ResultQuery := DbQuery(query)

		//baca hasil return data
		for ResultQuery.Next() {
			err := ResultQuery.Scan(&dum.Id_berita, &dum.Judul_berita, &dum.Isi_berita, &dum.Tanggal_dibuat, &dum.Foto)
			if err != nil {
				ResultQuery.Close()
				break
			} else {
				RtList = append(RtList, dum)
			}
		}

		//tutup koneksi
		ResultQuery.Close()

		//return detail token
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(RtList)
	}
}

///delete

func DeleteSiswa(w http.ResponseWriter, r *http.Request) {

	var data SDeleteSiswa
	var nis string
	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nis = data.Nis
	query := "DELETE FROM tb_siswa WHERE NIS = '" + nis + "';"

	ReturnStatus := DbExec(query)
	fmt.Println(ReturnStatus)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func DeleteGuru(w http.ResponseWriter, r *http.Request) {

	var data SDeleteGuru
	var nip string
	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nip = data.Nip
	query := "DELETE FROM tb_guru WHERE NIP = '" + nip + "';"

	ReturnStatus := DbExec(query)
	fmt.Println(ReturnStatus)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func DeletePrestasi(w http.ResponseWriter, r *http.Request) {

	var data SDeletePrestasi
	var nis string
	var nama_siswa string
	var jenis_prestasi string
	var nama_prestasi string
	var tahun_prestasi string
	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nis = data.Nis
	nama_siswa = data.Nama_siswa
	jenis_prestasi = data.Jenis_prestasi
	nama_prestasi = data.Nama_prestasi
	tahun_prestasi = data.Tahun_prestasi

	query := "DELETE FROM tb_prestasi WHERE nis = '" + nis + "' AND nama_siswa = '" + nama_siswa + "' AND jenis_prestasi = '" + jenis_prestasi + "' AND nama_prestasi = '" + nama_prestasi + "' AND tahun_prestasi = '" + tahun_prestasi + "';"

	ReturnStatus := DbExec(query)
	fmt.Println(ReturnStatus)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func DeleteMapel(w http.ResponseWriter, r *http.Request) {

	var data SDeleteMapel
	var kode_mapel string
	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	kode_mapel = data.Kode_mapel
	query := "DELETE FROM tb_mapel WHERE kode_mapel = '" + kode_mapel + "';"

	ReturnStatus := DbExec(query)
	fmt.Println(ReturnStatus)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func DeleteMapelDiampu(w http.ResponseWriter, r *http.Request) {

	var data SDeleteMapelDiampu
	var nip string
	var kode_mapel string
	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nip = data.Nip
	kode_mapel = data.Kode_mapel
	query := "DELETE FROM tb_mapel_diampu WHERE nip = '" + nip + "' AND kode_mapel = '" + kode_mapel + "';"

	ReturnStatus := DbExec(query)
	fmt.Println(ReturnStatus)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func DeleteBerita(w http.ResponseWriter, r *http.Request) {

	var data SDeleteBerita
	var id_berita string
	var Foto string
	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Foto = data.Foto
	id_berita = data.Id_berita
	query := "DELETE FROM tb_berita WHERE id_berita = '" + id_berita + "';"
	e := os.Remove("frontend/assets/img/files/" + Foto)
	if e != nil {
		log.Fatal(e)
	}
	ReturnStatus := DbExec(query)
	fmt.Println(ReturnStatus)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

//add / create
func AddSiswa(w http.ResponseWriter, r *http.Request) {
	var data JsonSiswa
	var Nis string
	var Nama_siswa string
	var JK string
	var Kelas string

	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Nis = data.Nis
	Nama_siswa = data.Nama_siswa
	JK = data.JK
	Kelas = data.Kelas
	var query = "INSERT INTO tb_siswa(`nis`, `nama_siswa`, `JK`, `kelas`) VALUES ('" + Nis + "', '" + Nama_siswa + "', '" + JK + "', '" + Kelas + "');"
	ReturnStatus := DbExec(query)
	fmt.Println(ReturnStatus)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func AddGuru(w http.ResponseWriter, r *http.Request) {

	var data JsonGuru
	var Nip string
	var Nama_guru string
	var Tgl_lahir string
	var Agama string
	var Alamat string
	var No_hp string
	var Pendidikan string
	var Foto string

	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Nip = data.Nip
	Nama_guru = data.Nama_guru
	Tgl_lahir = data.Tgl_lahir
	Agama = data.Agama
	Alamat = data.Alamat
	No_hp = data.No_hp
	Pendidikan = data.Pendidikan
	Foto = data.Foto

	var query = "INSERT INTO tb_guru(`nip`, `nama_guru`, `tgl_lahir`, `agama`, `alamat`, `no_hp`, `pendidikan`, `foto`) VALUES ('" + Nip + "', '" + Nama_guru + "', '" + Tgl_lahir + "', '" + Agama + "', '" + No_hp + "',  '" + Alamat + "', '" + Pendidikan + "', '" + Foto + "' );"
	ReturnStatus := DbExec(query)
	fmt.Println(ReturnStatus)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func AddPrestasi(w http.ResponseWriter, r *http.Request) {

	var data JsonPrestasi
	var Nis string
	var Nama_siswa string
	var Jenis_prestasi string
	var Nama_prestasi string
	var Tahun_prestasi string

	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Nis = data.Nis
	Nama_siswa = data.Nama_siswa
	Jenis_prestasi = data.Jenis_prestasi
	Nama_prestasi = data.Nama_prestasi
	Tahun_prestasi = data.Tahun_prestasi
	var query = "INSERT INTO tb_prestasi(`nis`, `nama_siswa`, `jenis_prestasi`, `nama_prestasi`, `tahun_prestasi`) VALUES ('" + Nis + "', '" + Nama_siswa + "', '" + Jenis_prestasi + "', '" + Nama_prestasi + "', '" + Tahun_prestasi + "');"
	ReturnStatus := DbExec(query)
	fmt.Println(ReturnStatus)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func AddMapel(w http.ResponseWriter, r *http.Request) {

	var data JsonMapel
	var Kode_mapel string
	var Nama_mapel string

	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Kode_mapel = data.Kode_mapel
	Nama_mapel = data.Nama_mapel
	var query = "INSERT INTO tb_mapel(`kode_mapel`, `nama_mapel`) VALUES ('" + Kode_mapel + "', '" + Nama_mapel + "');"
	ReturnStatus := DbExec(query)
	fmt.Println(ReturnStatus)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func AddMapelDiampu(w http.ResponseWriter, r *http.Request) {

	var data JsonMapelDiampu
	var Nip string
	var Kode_mapel string

	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Nip = data.Nip
	Kode_mapel = data.Kode_mapel
	var query = "INSERT INTO tb_siswa(`nip`, `kode_mapel`) VALUES ('" + Nip + "', '" + Kode_mapel + "');"
	ReturnStatus := DbExec(query)
	fmt.Println(ReturnStatus)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func AddBerita(w http.ResponseWriter, r *http.Request) {

	var data JsonBerita
	var Judul_berita string
	var Isi_berita string
	var Tanggal_dibuat string
	var Foto string

	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Judul_berita = data.Judul_berita
	Isi_berita = data.Isi_berita
	Tanggal_dibuat = data.Tanggal_dibuat
	Foto = data.Foto
	fmt.Println(Foto)
	var query = "INSERT INTO tb_berita(`id_berita`, `judul_berita`, `isi_berita`, `tanggal_dibuat`, `foto`) VALUES (NULL, '" + Judul_berita + "', '" + Isi_berita + "', '" + Tanggal_dibuat + "', '" + Foto + "');"
	ReturnStatus := DbExec(query)
	fmt.Println(query)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only accept POST request", http.StatusBadRequest)
		return
	}

	var respData UploadFoto
	var Foto string

	basePath, _ := os.Getwd()

	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		fileLocation := filepath.Join(basePath, "frontend/assets/img/files", part.FileName())
		dst, err := os.Create(fileLocation)
		if dst != nil {
			defer dst.Close()
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(part.FileName())
		Foto = part.FileName()
	}

	respData.Status = "oke"
	respData.NamaFoto = Foto
	UploadFoto, _ := json.Marshal(respData)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(string(UploadFoto))
}

func UpdateSiswa(w http.ResponseWriter, r *http.Request) {
	var data JsonSiswa
	var Nis string
	var Nama_siswa string
	var JK string
	var Kelas string
	var OldNis string
	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	OldNis = data.Old_nis
	Nis = data.Nis
	Nama_siswa = data.Nama_siswa
	JK = data.JK
	Kelas = data.Kelas
	var query = "UPDATE tb_siswa SET `nis`='" + Nis + "',  `Nama_siswa`='" + Nama_siswa + "',  `JK`='" + JK + "',  `Kelas`='" + Kelas + "' WHERE Nis = '" + OldNis + "';"
	fmt.Println(query)
	ReturnStatus := DbExec(query)
	fmt.Println(ReturnStatus)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func UpdateGuru(w http.ResponseWriter, r *http.Request) {
	var data JsonGuru
	var Nip string
	var Nama_guru string
	var Tgl_lahir string
	var Agama string
	var Alamat string
	var No_hp string
	var Pendidikan string
	var Foto string
	var OldNip string

	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Nip = data.Nip
	Nama_guru = data.Nama_guru
	Tgl_lahir = data.Tgl_lahir
	Agama = data.Agama
	Alamat = data.Alamat
	No_hp = data.No_hp
	Pendidikan = data.Pendidikan
	Foto = data.Foto
	OldNip = data.OldNip
	var query = "UPDATE tb_guru SET `nip`='" + Nip + "',  `nama_guru`='" + Nama_guru + "',  `tgl_lahir`='" + Tgl_lahir + "', `alamat`='" + Alamat + "',  `agama`='" + Agama + "', `no_hp`='" + No_hp + "', `pendidikan`='" + Pendidikan + "', `foto`='" + Foto + "' WHERE nip = '" + OldNip + "';"
	ReturnStatus := DbExec(query)
	fmt.Println(ReturnStatus)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func UpdatePrestasi(w http.ResponseWriter, r *http.Request) {
	var data JsonPrestasi
	var Nis string
	var Nama_siswa string
	var Jenis_prestasi string
	var Nama_prestasi string
	var Tahun_prestasi string
	var Nis_old string
	var Nama_siswa_old string
	var Jenis_prestasi_old string
	var Nama_prestasi_old string
	var Tahun_prestasi_old string

	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Nis_old = data.Nis_old
	Nama_siswa_old = data.Nama_siswa_old
	Jenis_prestasi_old = data.Jenis_prestasi_old
	Nama_prestasi_old = data.Nama_prestasi_old
	Tahun_prestasi_old = data.Tahun_prestasi_old
	Nis = data.Nis
	Nama_siswa = data.Nama_siswa
	Jenis_prestasi = data.Jenis_prestasi
	Nama_prestasi = data.Nama_prestasi
	Tahun_prestasi = data.Tahun_prestasi
	var query = "UPDATE tb_prestasi SET `nis`='" + Nis + "',  `nama_siswa`='" + Nama_siswa + "',  `jenis_prestasi`='" + Jenis_prestasi + "',  `nama_prestasi`='" + Nama_prestasi + "', `tahun_prestasi`='" + Tahun_prestasi + "' WHERE nis = '" + Nis_old + "' AND nama_siswa = '" + Nama_siswa_old + "' AND jenis_prestasi = '" + Jenis_prestasi_old + "' AND nama_prestasi = '" + Nama_prestasi_old + "' AND tahun_prestasi = '" + Tahun_prestasi_old + "';"
	ReturnStatus := DbExec(query)
	fmt.Println(query)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func UpdateMapel(w http.ResponseWriter, r *http.Request) {
	var data JsonMapel
	var Kode_mapel string
	var Kode_mapel_old string
	var Nama_mapel string

	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Kode_mapel = data.Kode_mapel
	Kode_mapel_old = data.Kode_mapel_old
	Nama_mapel = data.Nama_mapel
	var query = "UPDATE tb_mapel SET `kode_mapel`='" + Kode_mapel + "',  `nama_mapel`='" + Nama_mapel + "' WHERE Kode_mapel = '" + Kode_mapel_old + "';"
	ReturnStatus := DbExec(query)
	fmt.Println(query)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func UpdateMapelDiampu(w http.ResponseWriter, r *http.Request) {
	var data JsonMapelDiampu
	var Nip string
	var Kode_mapel string

	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Nip = data.Nip
	Kode_mapel = data.Kode_mapel
	var query = "UPDATE tb_mapeldiampu SET `nip`='" + Nip + "',  `Kode_mapel`='" + Kode_mapel + "' WHERE Nip = '" + Nip + "' AND Kode_mapel = '" + Kode_mapel + "';"
	ReturnStatus := DbExec(query)
	fmt.Println(ReturnStatus)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

func UpdateBerita(w http.ResponseWriter, r *http.Request) {
	var data JsonBerita
	var Id_berita string
	var Judul_berita string
	var Isi_berita string
	var Tanggal_dibuat string
	var Foto string

	var err error

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Id_berita = data.Id_berita
	Judul_berita = data.Judul_berita
	Isi_berita = data.Isi_berita
	Tanggal_dibuat = data.Tanggal_dibuat
	Foto = data.Foto
	var query = "UPDATE tb_berita SET `id_berita`='" + Id_berita + "',  `Judul_berita`='" + Judul_berita + "',  `Isi_berita`='" + Isi_berita + "',  `Tanggal_dibuat`='" + Tanggal_dibuat + "', `Foto`='" + Foto + "'  WHERE Id_berita = '" + Id_berita + "';"
	ReturnStatus := DbExec(query)
	fmt.Println(ReturnStatus)

	if ReturnStatus == "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("sukses")
	} else {
		fmt.Println(err.Error())
		return
	}
}

type SDeleteSiswa struct {
	Nis string `json:"nis"`
}

type SDeleteGuru struct {
	Nip string `json:"nip"`
}

type SDeletePrestasi struct {
	Nis            string `json:"nis"`
	Nama_siswa     string `json:"nama_siswa"`
	Jenis_prestasi string `json:"jenis_prestasi"`
	Nama_prestasi  string `json:"nama_prestasi"`
	Tahun_prestasi string `json:"tahun_prestasi"`
}

type SDeleteMapel struct {
	Kode_mapel string `json:"kode_mapel"`
}

type SDeleteMapelDiampu struct {
	Nip        string `json:"nip"`
	Kode_mapel string `json:"kode_mapel"`
}

type SDeleteBerita struct {
	Id_berita string `json:"id_berita"`
	Foto      string `json:"foto"`
}

type JsonSiswa struct {
	Old_nis    string `json:"old_nis"`
	Nis        string `json:"nis"`
	Nama_siswa string `json:"nama_siswa"`
	JK         string `json:"JK"`
	Kelas      string `json:"kelas"`
}

type JsonGuru struct {
	OldNip     string `json:"old_nip"`
	Nip        string `json:"nip"`
	Nama_guru  string `json:"nama_guru"`
	Tgl_lahir  string `json:"tgl_lahir"`
	Agama      string `json:"agama"`
	Alamat     string `json:"alamat"`
	No_hp      string `json:"no_hp"`
	Pendidikan string `json:"pendidikan"`
	Foto       string `json:"foto"`
}

type JsonPrestasi struct {
	Nis_old            string `json:"nis_old"`
	Nama_siswa_old     string `json:"nama_siswa_old"`
	Jenis_prestasi_old string `json:"jenis_prestasi_old"`
	Nama_prestasi_old  string `json:"nama_prestasi_old"`
	Tahun_prestasi_old string `json:"tahun_prestasi_old"`
	Nis                string `json:"nis"`
	Nama_siswa         string `json:"nama_siswa"`
	Jenis_prestasi     string `json:"jenis_prestasi"`
	Nama_prestasi      string `json:"nama_prestasi"`
	Tahun_prestasi     string `json:"tahun_prestasi"`
}

type JsonMapel struct {
	Kode_mapel_old string `json:"kode_mapel_old"`
	Kode_mapel     string `json:"kode_mapel"`
	Nama_mapel     string `json:"nama_mapel"`
}

type JsonMapelDiampu struct {
	Nip        string `json:"nip"`
	Kode_mapel string `json:"kode_mapel"`
}

type JsonBerita struct {
	Id_berita      string `json:"id_berita"`
	Judul_berita   string `json:"judul_berita"`
	Isi_berita     string `json:"isi_berita"`
	Tanggal_dibuat string `json:"tanggal_dibuat"`
	Foto           string `json:"foto"`
}
type ReturnLogin struct {
	StatusLogin string
	Hash        string
}

type UploadFoto struct {
	Status   string
	NamaFoto string
}

// func LogFailCallApi(NamaToken string, Market string, ApiResult string, ApiMarket string) {
// 	query := "INSERT INTO `fail_api_log` (`namaToken`, `market`, `time_detect`,  `api_result`, `apiMarket`, `no_dummy`) VALUES ("
// 	// '', '', '', '', '', '', NULL);

// 	ApiResultReplace := strings.ReplaceAll(ApiResult, "'", "\\'")

// 	query = query + "'" + NamaToken + "',"
// 	query = query + "'" + Market + "',"
// 	query = query + "'" + time.Now().String() + "',"
// 	query = query + "'" + ApiResultReplace + "',"
// 	query = query + "'" + ApiMarket + "',"
// 	query = query + "null);"
// 	_ = query
// 	DbExec(query)
// 	// fmt.Println("\n\n\n" + query + "\n\n\n\n")
// 	// os.Exit(1)
// }

// func UpdStTokenProblem(NamaToken string, StToken string, ScannerIdentifierCabang string) {
// 	Query := "CALL UpdSttokenProblem('" + NamaToken + "','" + StToken + "','" + ScannerIdentifierCabang + "');"
// 	fmt.Println("UpdStTokenProblem : ", DbFunc(Query))
// }

// func ResetLeaser() {
// 	Query := "UPDATE token SET leaser='-';"
// 	DbExec(Query)
// }

// func UpdStTokenNonArbitable(NamaToken string, StToken string, ScannerIdentifierCabang string) {
// 	TokenPriorList.SyMutex.Lock()
// 	var dum []TokenPrior
// 	for i := 0; i < len(TokenPriorList.List); i++ {
// 		if NamaToken == TokenPriorList.List[i].NamaToken {
// 			if TokenPriorList.List[i].PriorScan == CnfPriorCount {
// 				DeleteHasilScan(NamaToken)
// 			}
// 			TokenPriorList.List[i].PriorScan--
// 			if TokenPriorList.List[i].PriorScan < 1 {
// 				continue
// 			} else {
// 				TokenPriorList.List[i].Last_scan = time.Now()
// 				TokenPriorList.List[i].Leased = false
// 			}
// 		}
// 		dum = append(dum, TokenPriorList.List[i])
// 	}
// 	TokenPriorList.List = dum
// 	TokenPriorList.SyMutex.Unlock()
// 	// DbUpdStTokenNonArbitable(NamaToken, StToken, ScannerIdentifierCabang)
// }

// func DeleteHasilScan(NamaToken string) {
// 	Query := "CALL DeleteHasilScan('" + NamaToken + "');"
// 	DbFunc(Query)
// }

// func DbUpdStTokenNonArbitable(NamaToken string, StToken string, ScannerIdentifierCabang string) {
// 	Query := "CALL UpdStTokenNonArbitable('" + NamaToken + "','" + StToken + "','" + ScannerIdentifierCabang + "');"
// 	// fmt.Println("DbUpdStTokenNonArbitable : ", DbFunc(Query))
// 	_ = DbFunc(Query)
// }

// func GetDetailTokenFront(NamaToken string) DetailTokenFront {
// 	var DTokenFront DetailTokenFront

// 	query := "SELECT kepanjanganToken, logo_token_thumb,  " +
// 		"platform, explorer_1, explorer_2,  " +
// 		"contract_address, linkMarketCap, comment, status_verif " +
// 		"FROM token " +
// 		"WHERE namaToken='" + NamaToken + "'"

// 	ResultQuery := DbQuery(query)
// 	for ResultQuery.Next() {
// 		err := ResultQuery.Scan(&DTokenFront.Kepanjangan, &DTokenFront.LogoToken, &DTokenFront.Platform, &DTokenFront.Explorer1, &DTokenFront.Explorer2, &DTokenFront.ContractAddr, &DTokenFront.LinkMarketCap, &DTokenFront.Comment, &DTokenFront.StatusVerif)
// 		if err != nil {
// 			ResultQuery.Close()
// 			break
// 		}
// 	}
// 	ResultQuery.Close()
// 	return DTokenFront
// }

// func SaveHasilScan(NamaToken string, ScannerIdentifierCabang string, DTokenFront DetailTokenFront, RsComparizon []ReturnComparizonCore, ToComp StructStDetailToken) {
// 	var Query string
// 	var JsonPair string
// 	var JsonDetailPair JsonDetailPair
// 	var UniqKey string
// 	var ComparizonKey string
// 	_ = UniqKey
// 	_ = Query
// 	_ = JsonPair
// 	_ = ComparizonKey

// 	var BatchKey string
// 	BatchKey = RandStringRunes(5)
// 	for i := 0; i < len(RsComparizon); i++ {
// 		UniqKey = ""
// 		ComparizonKey = ""
// 		JsonDetailPair.Bids = RsComparizon[i].DtlBuy
// 		JsonDetailPair.Asks = RsComparizon[i].DtlSell

// 		//set ComparizonKey
// 		ComparizonKey += NamaToken + "|"
// 		ComparizonKey += RsComparizon[i].SellTo + "_"
// 		for k := 0; k < len(RsComparizon[i].DtlBuy); k++ {
// 			ComparizonKey += RsComparizon[i].DtlBuy[k].Pair

// 			if k != (len(RsComparizon[i].DtlBuy) - 1) {
// 				ComparizonKey += "_"
// 			}
// 		}
// 		ComparizonKey += "|"
// 		ComparizonKey += RsComparizon[i].BuyFrom + "_"
// 		for k := 0; k < len(RsComparizon[i].DtlSell); k++ {
// 			ComparizonKey += RsComparizon[i].DtlSell[k].Pair

// 			if k != (len(RsComparizon[i].DtlSell) - 1) {
// 				ComparizonKey += "_"
// 			}
// 		}
// 		//end set ComparizonKey

// 		//set UniqKey
// 		UniqKey += NamaToken + "-"
// 		UniqKey += strings.ReplaceAll(RsComparizon[i].SellTo, " ", "_") + "-"
// 		UniqKey += strings.ReplaceAll(RsComparizon[i].BuyFrom, " ", "_")
// 		//end set UniqKey

// 		jsonbyte, err := json.Marshal(JsonDetailPair)
// 		if err != nil {
// 			continue
// 		}
// 		JsonPair = string(jsonbyte)

// 		Query = "INSERT INTO hasilscan(`namaToken`, `statusVerif`, `platform`, " +
// 			"`st_read`, `SelisihPersen`, `SelisihEth`, " +
// 			"`sell_to`, `buy_from`, `detail_pair`, `UniqKey`, `batch_key`, `ComparizonKey`) " +
// 			"VALUES  ('" + NamaToken + "', " +
// 			" '" + DTokenFront.StatusVerif + "', " +
// 			" '" + DTokenFront.Platform + "', " +
// 			" '0', " +
// 			" '" + strconv.FormatFloat(RsComparizon[i].SelisihPersen, 'f', 0, 64) + "', " +
// 			" '" + strconv.FormatFloat(RsComparizon[i].SelisihEth, 'f', 2, 64) + "', " +
// 			" '" + RsComparizon[i].SellTo + "', " +
// 			" '" + RsComparizon[i].BuyFrom + "', " +
// 			" '" + JsonPair + "', " +
// 			" '" + UniqKey + "', " +
// 			" '" + BatchKey + "', " +
// 			" '" + ComparizonKey + "') " +
// 			"ON DUPLICATE KEY UPDATE " +
// 			"`statusVerif`='" + DTokenFront.StatusVerif + "', " +
// 			"`platform`='" + DTokenFront.Platform + "', " +
// 			"`st_read`='0', " +
// 			"`SelisihPersen`='" + strconv.FormatFloat(RsComparizon[i].SelisihPersen, 'f', 0, 64) + "', " +
// 			"`SelisihEth`='" + strconv.FormatFloat(RsComparizon[i].SelisihEth, 'f', 3, 64) + "', " +
// 			"`sell_to`='" + RsComparizon[i].SellTo + "', " +
// 			"`buy_from`='" + RsComparizon[i].BuyFrom + "', " +
// 			"`detail_pair`='" + JsonPair + "', " +
// 			"`batch_key`='" + BatchKey + "', " +
// 			"`ComparizonKey`='" + ComparizonKey + "';"
// 		DbExec(Query)
// 	}
// 	Query = "DELETE FROM hasilscan WHERE (namaToken='" + NamaToken + "') AND (batch_key!='" + BatchKey + "');"
// 	DbExec(Query)

// 	TokenPriorList.SyMutex.Lock()
// 	var dum TokenPrior
// 	AlreadyInList := false
// 	for i := 0; i < len(TokenPriorList.List); i++ {
// 		if NamaToken == TokenPriorList.List[i].NamaToken {
// 			AlreadyInList = true
// 			TokenPriorList.List[i].Last_scan = time.Now()
// 			TokenPriorList.List[i].PriorScan = 6
// 			TokenPriorList.List[i].Leased = false
// 			break
// 		}
// 	}
// 	if AlreadyInList == false {
// 		dum.NamaToken = NamaToken
// 		dum.Last_scan = time.Now()
// 		dum.PriorScan = CnfPriorCount
// 		TokenPriorList.List = append(TokenPriorList.List, dum)
// 	}
// 	TokenPriorList.SyMutex.Unlock()
// }

// func SaveDetailHasilScan(NamaToken string, ScannerIdentifierCabang string, DTokenFront DetailTokenFront, RsComparizon []ReturnComparizonCore, ToComp StructStDetailToken) {
// 	// var DetailTokenFront DetailTokenFront
// 	DTokenFront.NamaToken = NamaToken
// 	DTokenFront.ScannerIdentifierCabang = ScannerIdentifierCabang

// 	var DMarket DetailMarket
// 	//detail market bo
// 	for i := 0; i < len(ToComp.DToken); i++ {
// 		DMarket.Market = ToComp.DToken[i].Market
// 		DMarket.LinkMarket = ToComp.DToken[i].LinkMarket
// 		DMarket.Pair = ToComp.DToken[i].Pair
// 		DMarket.Price = ToComp.DToken[i].BestBid
// 		DTokenFront.DetailMarketBo = append(DTokenFront.DetailMarketBo, DMarket)
// 	}
// 	//end detail market bo

// 	//detail market bo
// 	for i := 0; i < len(ToComp.DToken); i++ {
// 		DMarket.Market = ToComp.DToken[i].Market
// 		DMarket.LinkMarket = ToComp.DToken[i].LinkMarket
// 		DMarket.Pair = ToComp.DToken[i].Pair
// 		DMarket.Price = ToComp.DToken[i].BestAsk
// 		DTokenFront.DetailMarketSo = append(DTokenFront.DetailMarketSo, DMarket)
// 	}
// 	//end detail market bo

// 	//sorting detail market bo by price
// 	DTokenFront.DetailMarketBo = SortDescDTokenFront(DTokenFront.DetailMarketBo)
// 	//end sorting

// 	//sorting detail market so by price
// 	DTokenFront.DetailMarketSo = SortAscDTokenFront(DTokenFront.DetailMarketSo)
// 	//end sorting

// 	//sorting detail arbitable
// 	RsComparizon = SortDetailArbitable(RsComparizon)
// 	DTokenFront.DetailArbitable = RsComparizon
// 	//end sorting

// 	//marshalling detail token
// 	var JsonDetailToken []byte
// 	var err error
// 	for {
// 		JsonDetailToken, err = json.Marshal(DTokenFront)
// 		if err == nil {
// 			break
// 		}
// 	}
// 	//end marshal
// 	Query := "CALL SvDetailHasilScan('" + NamaToken + "', '" + ScannerIdentifierCabang + "', '" + string(JsonDetailToken) + "', '" + strconv.Itoa(CnfPriorScanNum) + "')"
// 	DbFunc(Query)
// }

// func PrecissionAfterComma(precission int, val float64) float64 {
// 	var X float64
// 	X = 1
// 	for i := 0; i < precission; i++ {
// 		X = X * 10
// 	}
// 	return float64(int(val*X)) / X
// }

// func SortDetailArbitable(DetailArbitable []ReturnComparizonCore) []ReturnComparizonCore {
// 	var Temp ReturnComparizonCore
// 	for i := 0; i < len(DetailArbitable); i++ {
// 		for k := i + 1; k < len(DetailArbitable); k++ {
// 			if DetailArbitable[i].SelisihPersen < DetailArbitable[k].SelisihPersen {
// 				Temp = DetailArbitable[i]
// 				DetailArbitable[i] = DetailArbitable[k]
// 				DetailArbitable[k] = Temp
// 			}
// 		}
// 	}
// 	return DetailArbitable
// }

// func SortAscDTokenFront(DMarket []DetailMarket) []DetailMarket {
// 	var Temp DetailMarket
// 	for i := 0; i < len(DMarket); i++ {
// 		for k := i + 1; k < len(DMarket); k++ {
// 			if DMarket[i].Price > DMarket[k].Price {
// 				Temp = DMarket[i]
// 				DMarket[i] = DMarket[k]
// 				DMarket[k] = Temp
// 			}
// 		}
// 	}
// 	return DMarket
// }

// func SortDescDTokenFront(DMarket []DetailMarket) []DetailMarket {
// 	var Temp DetailMarket
// 	for i := 0; i < len(DMarket); i++ {
// 		for k := i + 1; k < len(DMarket); k++ {
// 			if DMarket[i].Price < DMarket[k].Price {
// 				Temp = DMarket[i]
// 				DMarket[i] = DMarket[k]
// 				DMarket[k] = Temp
// 			}
// 		}
// 	}
// 	return DMarket
// }

// type DetailTokenFront struct {
// 	NamaToken               string
// 	ScannerIdentifierCabang string
// 	Kepanjangan             string
// 	LogoToken               string
// 	Platform                string
// 	Explorer1               string
// 	Explorer2               string
// 	ContractAddr            string
// 	LinkMarketCap           string
// 	Comment                 string
// 	StatusVerif             string
// 	DetailArbitable         []ReturnComparizonCore
// 	DetailMarketBo          []DetailMarket
// 	DetailMarketSo          []DetailMarket
// }

// type DetailMarket struct {
// 	Market     string
// 	Pair       string
// 	Price      float64
// 	LinkMarket string
// }

// type JsonDetailPair struct {
// 	Bids []DetailArbitable
// 	Asks []DetailArbitable
// }
