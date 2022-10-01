package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//load file config
	LoadConf()

	// //koneksi ke database
	DbConnect("")

	//defer db close connection
	defer DbClose()

	// fmt.Println("done connect")

	// DeleteSiswa("12345")
	WebServer()
}

func WebServer() {
	// routing func
	Route()

	//serve address
	var address = ConfServerIp + ":" + ConfServerPort
	fmt.Println("ip addr : ", address)
	fmt.Printf("server started at http://%s\n", address)

	server := new(http.Server)
	server.Addr = address
	server.ReadTimeout = time.Second * 10
	server.WriteTimeout = time.Second * 10

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Route() {
	// static html
	html := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", http.StripPrefix("/", html))

	// // get data route
	http.HandleFunc("/get-siswa", GetSiswa)
	http.HandleFunc("/get-guru", GetGuru)
	http.HandleFunc("/get-prestasi", GetPrestasi)
	http.HandleFunc("/get-mapel", GetMapel)
	http.HandleFunc("/get-mapel-diampu", GetMapelDiampu)
	http.HandleFunc("/get-berita", GetBerita)
	http.HandleFunc("/delete-siswa", DeleteSiswa)
	http.HandleFunc("/delete-guru", DeleteGuru)
	http.HandleFunc("/delete-prestasi", DeletePrestasi)
	http.HandleFunc("/delete-mapel", DeleteMapel)
	http.HandleFunc("/delete-berita", DeleteBerita)
	http.HandleFunc("/delete-mapel-diampu", DeleteMapelDiampu)
	http.HandleFunc("/add-siswa", AddSiswa)
	http.HandleFunc("/add-guru", AddGuru)
	http.HandleFunc("/add-prestasi", AddPrestasi)
	http.HandleFunc("/add-mapel", AddMapel)
	http.HandleFunc("/add-mapel-diampu", AddMapelDiampu)
	http.HandleFunc("/add-berita", AddBerita)
	http.HandleFunc("/update-siswa", UpdateSiswa)
	http.HandleFunc("/update-guru", UpdateGuru)
	http.HandleFunc("/update-prestasi", UpdatePrestasi)
	http.HandleFunc("/update-mapel", UpdateMapel)
	http.HandleFunc("/update-mapel-diampu", UpdateMapelDiampu)
	http.HandleFunc("/update-berita", UpdateBerita)
	http.HandleFunc("/check-login", LoginPage)
	http.HandleFunc("/upload-foto", handleUpload)
	http.HandleFunc("/add-foto", handleUpload)

}
