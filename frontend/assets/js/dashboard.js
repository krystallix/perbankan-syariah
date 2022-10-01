function limit(text, count, insertDots) {
   return text.slice(0, count) + (((text.length > count) && insertDots) ? "..." : "");
}

(function ($) {
   $.fn.serializefiles = function () {
      var obj = $(this);
      /* ADD FILE TO PARAM AJAX */
      var formData = new FormData();
      $.each($(obj).find("input[type='file']"), function (i, tag) {
         $.each($(tag)[0].files, function (i, file) {
            formData.append(tag.name, file);
         });
      });
      var params = $(obj).serializeArray();
      $.each(params, function (i, val) {
         formData.append(val.name, val.value);
      });
      return formData;
   };
})(jQuery);

$(document).ready(function () {
   
   // CRUD SISWA
   $.ajax({
      type: "post",
      url: "/get-siswa",
      crossDomain: true,
      success: function (response) {
         dataSiswa = ""
         dataSiswa1 = ""
         $.each(response, function (k, v) {
            no = k + 1
            dataSiswa += `<tr>
            <td>
            `+ no + `.
            </td>
            <td>
            `+ v.nis + `
            </td>
            <td>
            `+ v.nama_siswa + `
            </td>
            <td>
            `+ v.JK + `
            </td>
            <td>
            `+ v.kelas + `
            </td>            </tr>
            `
            dataSiswa1 += `<tr>
            <td>
            `+ no + `.
            </td>
            <td>
            `+ v.nis + `
            </td>
            <td>
            `+ v.nama_siswa + `
            </td>
            <td>
            `+ v.JK + `
            </td>
            <td>
            `+ v.kelas + `
            </td><td><i  data-id="`+ v.nis + `" data-siswa="` + v.nama_siswa + `" data-gender="` + v.JK + `" data-kelas="` + v.kelas + `" class="far edit-siswa-btn px-2 fa-lg fa-edit action-icon"></i>
            <i  data-id="`+ v.nis + `" data-siswa="` + v.nama_siswa + `" data-gender="` + v.JK + `" data-kelas="` + v.kelas + `" class="far delete-siswa-btn fa-lg px-3 action-icon-delete fa-trash-alt"></i></td>
            </tr>
            `
         })
         
         $(".data-siswa").html(dataSiswa)
         
         $("#data-siswa-dashboard").html(dataSiswa1)
         $('.table-siswa').DataTable();
      }
   })
   $(document).on("click", ".delete-siswa-btn", function () {
      nis = $(this).attr("data-id")
      siswa = $(this).attr("data-siswa")
      $("#nis-delete").val(nis)
      $("#siswa-delete").text(siswa)
      $("#delete-siswa-modal").modal("show")
   })
   $("#deleteSiswaForm").submit(function (e) {
      e.preventDefault()
      data = $(this).serializeObject()
      $.ajax({
         url: "/delete-siswa",
         method: "post",
         dataType: "JSON",
         data: JSON.stringify(data),
         success: function (resp) {
            if (resp == "sukses") {
               Snackbar.show({
                  text: 'Edit Sukses',
                  backgroundColor: '#fff',
                  textColor: '#24D1BC',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               })
               $("#delete-siswa-modal").modal("hide")
            } else {
               Snackbar.show({
                  text: 'Edit Gagal',
                  // backgroundColor: '#fff',
                  textColor: '#ff69b4',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               });
            }
         }
      })
   })
   $(document).on("click", ".edit-siswa-btn", function () {
      nis = $(this).attr("data-id")
      siswa = $(this).attr("data-siswa")
      gender = $(this).attr("data-gender")
      kelas = $(this).attr("data-kelas")
      // (nis + siswa + gender + kelas)
      $("#oldNis-edit").val(nis)
      $("#nis-edit").val(nis)
      $("#nama-siswa-edit").val(siswa)
      $('#jk-edit option[value=' + gender + ']').attr('selected', 'selected');
      $("#kelas-edit").val(kelas)
      $("#edit-siswa-modal").modal("show")
   })
   $("#editSiswaForm").submit(function (e) {
      e.preventDefault()
      data = $(this).serializeObject()
      $.ajax({
         url: "/update-siswa",
         method: "post",
         dataType: "JSON",
         data: JSON.stringify(data),
         success: function (resp) {
            if (resp == "sukses") {
               Snackbar.show({
                  text: 'Edit Sukses',
                  backgroundColor: '#fff',
                  textColor: '#24D1BC',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               })
               $("#edit-siswa-modal").modal("hide")
            } else {
               Snackbar.show({
                  text: 'Edit Gagal',
                  // backgroundColor: '#fff',
                  textColor: '#ff69b4',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               });
            }
         }
      })
   })
   $("#add-siswa-btn").click(function () {
      $("#add-siswa-modal").modal("show")
   })
   $("#addSiswaForm").submit(function (e) {
      e.preventDefault()
      data = $(this).serializeObject()
      $.ajax({
         url: "/add-siswa",
         method: "post",
         dataType: "JSON",
         data: JSON.stringify(data),
         success: function (resp) {
            if (resp == "sukses") {
               Snackbar.show({
                  text: 'Tambah Sukses',
                  backgroundColor: '#fff',
                  textColor: '#24D1BC',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               })
               $("#add-siswa-modal").modal("hide")
            } else {
               Snackbar.show({
                  text: 'Tambah Gagal',
                  // backgroundColor: '#fff',
                  textColor: '#ff69b4',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               });
            }
         }
      })
   })
   
   // CRUD PRESTASI
   $.ajax({
      type: "post",
      url: "/get-prestasi",
      crossDomain: true,
      success: function (response) {
         dataPrestasi = ""
         dataPrestasi1 = ""
         $.each(response, function (k, v) {
            no = k + 1
            dataPrestasi += `<tr>
            <td>
            `+ no + `.
            </td>
            <td>
            `+ v.nis + `
            </td>
            <td>
            `+ v.nama_siswa + `
            </td>
            <td>
            `+ v.jenis_prestasi + `
            </td>
            <td>
            `+ v.nama_prestasi + `
            </td>
            <td>
            `+ v.tahun_prestasi + `
            </td>
            </tr>`
            dataPrestasi1 += `<tr>
            <td>
            `+ no + `.
            </td>
            <td>
            `+ v.nis + `
            </td>
            <td>
            `+ v.nama_siswa + `
            </td>
            <td>
            `+ v.jenis_prestasi + `
            </td>
            <td>
            `+ v.nama_prestasi + `
            </td>
            <td>
            `+ v.tahun_prestasi + `
            </td>
            <td><i  data-id="`+ v.nis + `" data-siswa="` + v.nama_siswa + `" data-jenis="` + v.jenis_prestasi + `" data-prestasi="` + v.nama_prestasi + `" data-tahun="` + v.tahun_prestasi + `" class="far edit-prestasi-btn px-2 fa-lg fa-edit action-icon"></i>
            <i  data-id="`+ v.nis + `" data-siswa="` + v.nama_siswa + `" data-jenis="` + v.jenis_prestasi + `" data-prestasi="` + v.nama_prestasi + `" data-tahun="` + v.tahun_prestasi + `" class="far delete-prestasi-btn fa-lg px-3 action-icon-delete fa-trash-alt"></i></td>
            </tr>`
         })
         $(".data-prestasi").html(dataPrestasi)
         $(".data-prestasi-dashboard").html(dataPrestasi1)
      }
   })
   $(document).on("click", ".edit-prestasi-btn", function () {
      s = $(this).attr("data-siswa")
      j = $(this).attr("data-jenis")
      p = $(this).attr("data-prestasi")
      t = $(this).attr("data-tahun")
      n = $(this).attr("data-id")
      $("#nis-old").val(n)
      $("#nama-siswa-old").val(s)
      $("#jenis-prestasi-old").val(j)
      $("#nama-prestasi-old").val(p)
      $("#tahun-prestasi-old").val(t)
      $("#nis-edit").val(n)
      $("#nama-siswa-edit").val(s)
      $("#jenis-prestasi-edit").val(j)
      $("#nama-prestasi-edit").val(p)
      $("#tahun-prestasi-edit").val(t)
      $("#edit-prestasi-modal").modal("show")
   })
   $(document).on("click", ".delete-prestasi-btn", function () {
      s = $(this).attr("data-siswa")
      j = $(this).attr("data-jenis")
      p = $(this).attr("data-prestasi")
      t = $(this).attr("data-tahun")
      n = $(this).attr("data-id")
      $("#nis-delete").val(n)
      $("#nama-siswa-delete").val(s)
      $("#jenis-prestasi-delete").val(j)
      $("#nama-prestasi-delete").val(p)
      $("#tahun-prestasi-delete").val(t)
      $("#delete-prestasi-data").text("Ingin menghapus prestasi dengan nama siswa " + s + " dengan prestasi dan tahun prestasi " + p + " " + t)
      $("#delete-prestasi-modal").modal("show")
   })
   $("#deletePrestasiForm").submit(function (e) {
      e.preventDefault()
      data = $(this).serializeObject()
      $.ajax({
         url: "/delete-prestasi",
         method: "post",
         dataType: "JSON",
         data: JSON.stringify(data),
         success: function (resp) {
            if (resp == "sukses") {
               Snackbar.show({
                  text: 'Hapus Sukses',
                  backgroundColor: '#fff',
                  textColor: '#24D1BC',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               })
               $("#delete-prestasi-modal").modal("hide")
            } else {
               Snackbar.show({
                  text: 'Hapus Gagal',
                  // backgroundColor: '#fff',
                  textColor: '#ff69b4',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               });
            }
         }
      })
   })
   $("#editPrestasiForm").submit(function (e) {
      e.preventDefault()
      data = $(this).serializeObject()
      // (data)
      $.ajax({
         url: "/update-prestasi",
         method: "post",
         dataType: "JSON",
         data: JSON.stringify(data),
         success: function (resp) {
            if (resp == "sukses") {
               Snackbar.show({
                  text: 'Edit Sukses',
                  backgroundColor: '#fff',
                  textColor: '#24D1BC',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               })
               $("#edit-prestasi-modal").modal("hide")
            } else {
               Snackbar.show({
                  text: 'Edit Gagal',
                  // backgroundColor: '#fff',
                  textColor: '#ff69b4',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               });
            }
         }
      })
   })
   $("#add-prestasi-btn").click(function () {
      $("#add-prestasi-modal").modal("show")
   })
   $("#addPrestasiForm").submit(function (e) {
      e.preventDefault()
      data = $(this).serializeObject()
      $.ajax({
         url: "/add-prestasi",
         method: "post",
         dataType: "JSON",
         data: JSON.stringify(data),
         success: function (resp) {
            if (resp == "sukses") {
               Snackbar.show({
                  text: 'Tambah Sukses',
                  backgroundColor: '#fff',
                  textColor: '#24D1BC',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               })
               $("#add-prestasi-modal").modal("hide")
            } else {
               Snackbar.show({
                  text: 'Tambah Gagal',
                  // backgroundColor: '#fff',
                  textColor: '#ff69b4',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               });
            }
         }
      })
   })
   
   // CRUD GURU
   $.ajax({
      type: "post",
      url: "/get-guru",
      crossDomain: true,
      success: function (response) {
         dataGuru = ""
         show_guru = ""
         $.each(response, function (k, v) {
            no = k + 1
            dataGuru += `<tr>
            <td>
            `+ v.nip + `
            </td>
            <td>
            `+ v.nama_guru + `
            </td>
            <td>
            `+ v.tgl_lahir + `
            </td>
            <td>
            `+ v.agama + `
            </td>
            <td  class='detail-get-val'data-id="`+ v.alamat + `">
            `+ limit(v.alamat, 20, 1) + `
            </td>
            <td>
            `+ v.no_hp + `
            </td>
            <td  class='detail-get-val'data-id="`+ v.pendidikan + `">
            `+ limit(v.pendidikan, 20, 1) + `
            </td>
            <td>
            `+ v.foto + `
            </td>
            <td><i  data-id="`+ v.nip + `" data-guru="` + v.nama_guru + `" data-born="` + v.tgl_lahir + `" data-agama="` + v.agama + `" data-alamat="` + v.alamat + `" data-hp="` + v.no_hp + `" data-pendidikan="` + v.pendidikan + `" class="far edit-guru-btn px-2 fa-lg fa-edit action-icon"></i>
            <i  data-id="`+ v.nip + `" data-guru="` + v.nama_guru + `" data-born="` + v.tgl_lahir + `" data-agama="` + v.agama + `" data-alamat="` + v.alamat + `" data-hp="` + v.no_hp + `" data-pendidikan="` + v.pendidikan + `"  class="far delete-guru-btn fa-lg px-3 action-icon-delete fa-trash-alt"></i></td>
            </tr>
            `
            
            show_guru += `
            <div class="swiper-slide">
            <div class="testimonial-item px-2 py-2">
            <img src="assets/img/testimonials/testimonials-1.jpg" class="testimonial-img" alt="">
            <h3>`+ v.nama_guru + `</h3>
            <p class="">`+ v.tgl_lahir + `</p>
            <p class="">`+ v.agama + `</p>
            <p class="">`+ v.alamat + `</p>
            <p class="">`+ v.no_hp + `</p>
            <p class="">`+ v.pendidikan + `</p>
            </div>
            </div>
            `
         })
         
         $(".data-guru").html(dataGuru)
         $("#data-guru").html(show_guru)
         
         $('.detail-get-val').each(function () {
            tippy(this, {
               content: function (reference) {
                  return reference.getAttribute("data-id");
               }
            });
         })
      }
   })
   $(document).on("click", ".edit-guru-btn", function () {
      g = $(this).attr("data-guru")
      a = $(this).attr("data-alamat")
      b = $(this).attr("data-born")
      ag = $(this).attr("data-agama")
      h = $(this).attr("data-hp")
      p = $(this).attr("data-pendidikan")
      n = $(this).attr("data-id")
      $("#old_nip").val(n)
      $("#guru-edit").val(g)
      $("#ttl-edit").val(b)
      $("#alamat-edit").val(a)
      $("#agama-edit").val(ag)
      $("#hp-edit").val(h)
      $("#pendidikan-edit").val(p)
      $("#nip-edit").val(n)
      $("#edit-guru-modal").modal("show")
   })
   $("#add-guru-btn").click(function () {
      $("#add-guru-modal").modal("show")
   })
   $("#editGuruForm").submit(function (e) {
      e.preventDefault()
      data = $(this).serializeObject()
      // (data)
      $.ajax({
         url: "/update-guru",
         method: "post",
         dataType: "JSON",
         data: JSON.stringify(data),
         success: function (resp) {
            if (resp == "sukses") {
               Snackbar.show({
                  text: 'Tambah Sukses',
                  backgroundColor: '#fff',
                  textColor: '#24D1BC',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               })
               $("#edit-guru-modal").modal("hide")
            } else {
               Snackbar.show({
                  text: 'Tambah Gagal',
                  // backgroundColor: '#fff',
                  textColor: '#ff69b4',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               });
            }
         }
      })
   })
   $("#addGuruForm").submit(function (e) {
      e.preventDefault()
      data = $(this).serializeObject()
      // (data)
      $.ajax({
         url: "/add-guru",
         method: "post",
         dataType: "JSON",
         data: JSON.stringify(data),
         success: function (resp) {
            if (resp == "sukses") {
               Snackbar.show({
                  text: 'Tambah Sukses',
                  backgroundColor: '#fff',
                  textColor: '#24D1BC',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               })
               $("#add-guru-modal").modal("hide")
            } else {
               Snackbar.show({
                  text: 'Tambah Gagal',
                  // backgroundColor: '#fff',
                  textColor: '#ff69b4',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               });
            }
         }
      })
   })
   $(document).on("click", ".delete-guru-btn", function () {
      nip = $(this).attr("data-id")
      guru = $(this).attr("data-guru")
      $("#nip-delete").val(nip)
      $("#guru-delete").text(guru)
      $("#delete-guru-modal").modal("show")
   })
   $("#deleteGuruForm").submit(function (e) {
      e.preventDefault()
      data = $(this).serializeObject()
      $.ajax({
         url: "/delete-guru",
         method: "post",
         dataType: "JSON",
         data: JSON.stringify(data),
         success: function (resp) {
            if (resp == "sukses") {
               Snackbar.show({
                  text: 'Edit Sukses',
                  backgroundColor: '#fff',
                  textColor: '#24D1BC',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               })
               $("#delete-guru-modal").modal("hide")
            } else {
               Snackbar.show({
                  text: 'Edit Gagal',
                  // backgroundColor: '#fff',
                  textColor: '#ff69b4',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               });
            }
         }
      })
      
   })
   
   // CRUD MAPEL 
   $.ajax({
      type: "post",
      url: "/get-mapel",
      crossDomain: true,
      success: function (response) {
         dataMapel = ""
         dataMapelTable = ""
         $.each(response, function (k, v) {
            dataMapel += `
            <span>`+ v.kode_mapel + `. ` + v.nama_mapel + `</span><br>
            `
            dataMapelTable += `
            <tr><td>`+ v.kode_mapel + `</td><td>` + v.nama_mapel + `</td>
            <td><i  data-id="`+ v.kode_mapel + `" data-nama="` + v.nama_mapel + `" class="far edit-mapel-btn px-2 fa-lg fa-edit action-icon"></i>
            <i  data-id="`+ v.kode_mapel + `" data-nama="` + v.nama_mapel + `" class="far delete-mapel-btn fa-lg px-3 action-icon-delete fa-trash-alt"></i></td>
            </tr>
            `
         })
         $("#data-mapel").html(dataMapelTable)
         $("#mapel").html(dataMapel)
      }
   })
   $("#add-mapel-btn").click(function () {
      $("#add-mapel-modal").modal("show")
   })
   $("#addMapelForm").submit(function (e) {
      e.preventDefault()
      data = $(this).serializeObject()
      $.ajax({
         url: "/add-mapel",
         method: "post",
         dataType: "JSON",
         data: JSON.stringify(data),
         success: function (resp) {
            if (resp == "sukses") {
               Snackbar.show({
                  text: 'Tambah Sukses',
                  backgroundColor: '#fff',
                  textColor: '#24D1BC',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               })
               $("#add-mapel-modal").modal("hide")
            } else {
               Snackbar.show({
                  text: 'Tambah Gagal',
                  // backgroundColor: '#fff',
                  textColor: '#ff69b4',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               });
            }
         }
      })
   })
   $(document).on("click", ".edit-mapel-btn", function () {
      k = $(this).attr("data-id")
      n = $(this).attr("data-nama")
      $("#old_kode_mapel").val(k)
      $("#kode_mapel-edit").val(k)
      $("#nama_mapel-edit").val(n)
      $("#edit-mapel-modal").modal("show")
   })
   $("#editMapelForm").submit(function (e) {
      e.preventDefault()
      data = $(this).serializeObject()
      // (data)
      $.ajax({
         url: "/update-mapel",
         method: "post",
         dataType: "JSON",
         data: JSON.stringify(data),
         success: function (resp) {
            if (resp == "sukses") {
               Snackbar.show({
                  text: 'Edit Sukses',
                  backgroundColor: '#fff',
                  textColor: '#24D1BC',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               })
               $("#edit-mapel-modal").modal("hide")
            } else {
               Snackbar.show({
                  text: 'Edit Gagal',
                  // backgroundColor: '#fff',
                  textColor: '#ff69b4',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               });
            }
         }
      })
   })
   $(document).on("click", ".delete-mapel-btn", function () {
      k = $(this).attr("data-id")
      n = $(this).attr("data-nama")
      $("#kode_mapel-delete").val(k)
      $("#delete-mapel-info").text("Ingin Menghapus Mapel dengan kode " + k + " dan Nama Mapel " + n + " ?")
      $("#delete-mapel-modal").modal("show")
   })
   $("#deleteMapelForm").submit(function (e) {
      e.preventDefault()
      data = $(this).serializeObject()
      // (data)
      $.ajax({
         url: "/delete-mapel",
         method: "post",
         dataType: "JSON",
         data: JSON.stringify(data),
         success: function (resp) {
            if (resp == "sukses") {
               Snackbar.show({
                  text: 'Delete Sukses',
                  backgroundColor: '#fff',
                  textColor: '#24D1BC',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               })
               $("#delete-mapel-modal").modal("hide")
            } else {
               Snackbar.show({
                  text: 'Delete Gagal',
                  // backgroundColor: '#fff',
                  textColor: '#ff69b4',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               });
            }
         }
      })
   })
   
   // CRUD Berita
   $.ajax({
      type: "post",
      url: "/get-berita",
      crossDomain: true,
      success: function (response) {
         console.log(response)
         html_berita = ""
         berita_home = ""
         galeri = ""
         $.each(response, function(k,v){
            html_berita += `
            <div class="card mx-4 my-2" style="height: 450px; width: 20rem; padding: 2px;">
            <img class="card-img-top" src="../../assets/img/files/`+v.foto+`" alt="Card image cap" style="width: 100%; height: 166px; object-fit: cover;">
            <div class="card-body">
            <h5 class="card-title">`+v.judul_berita+`</h5>
            <h6>`+v.tanggal_dibuat+`</h6>
            <p class="card-text">`+limit(v.isi_berita, 250, 1)+`</p>
            <div class='d-flex justify-content-end pt-2'>
                  <i data-id="`+v.id_berita+`"  data-foto="`+v.foto+`"  data-judul="`+v.judul_berita+`"  data-tanggal="`+v.tanggal_dibuat+`"  data-isi="`+v.isi_berita+`" class="show-berita-btn fa-solid action-icon px-2 fa-lg fa-arrow-up-right-from-square"></i>      
                  <i data-id="`+v.id_berita+`"  data-foto="`+v.foto+`"  data-judul="`+v.judul_berita+`" class="fa-regular action-icon-delete px-2 fa-lg fa-trash-can delete-berita-btn"></i>
            </div>
            </div>
            </div>`

            berita_home += `
            <div class="col-lg-4 col-md-6 berita_show_modal" data-foto="`+v.foto+`"  data-judul="`+v.judul_berita+`"  data-tanggal="`+v.tanggal_dibuat+`"  data-isi="`+v.isi_berita+`" >
            <div class="icon-box" data-aos="zoom-in-left">
              <div class="icon">            
              <img class="card-img-top" src="../../assets/img/files/`+v.foto+`" alt="Card image cap" style="width: 80px; height: 200px; object-fit: cover;">
              </div>
              <h4 class="title text-center">`+v.judul_berita+`</h4>
              <h6 class="text-center">`+v.tanggal_dibuat+`</h6>
              <p class="description">`+limit(v.isi_berita, 250, 1)+`</p>
            </div>
          </div>`

          galeri += `
            <img src="../../assets/img/files/`+v.foto+`">
          `
         })
         $("#photos").append(galeri)
         $("#berita-show").html(berita_home)
         $("#content-berita").html(html_berita)
      }
   })
   $(document).on("click", ".berita_show_modal", function(){
      j = $(this).attr("data-judul")
      f = $(this).attr("data-foto")
      is = $(this).attr("data-isi")
      t = $(this).attr("data-tanggal")
      $("#judul-home").html("<h3>"+j+"</h3>")
      $("#tanggal-home").html("<h6>"+t+"</h6>")
      $("#foto-home").attr("src","../../assets/img/files/"+f);
      $("#isi-home").html("<p>"+is+"</p>")
      $("#home-berita-modal").modal("show")
   })
   $(document).on("click", ".show-berita-btn", function(){
      j = $(this).attr("data-judul")
      i = $(this).attr("data-id")
      f = $(this).attr("data-foto")
      is = $(this).attr("data-isi")
      t = $(this).attr("data-tanggal")
      $("#judul-show").html("<h3>"+j+"</h3>")
      $("#tanggal-show").html("<h6>"+t+"</h6>")
      $("#foto-show").attr("src","../../assets/img/files/"+f);
      $("#isi-show").html("<p>"+is+"</p>")
      $("#show-berita-modal").modal("show")
   })
   $(document).on("click", ".delete-berita-btn", function(){
      j = $(this).attr("data-judul")
      $("#delete-berita-info").text("Ingin menghapus berita dengan judul "+j+" ?")
      i = $(this).attr("data-id")
      f = $(this).attr("data-foto")
      $("#foto-delete").val(f)
      $("#id-berita-delete").val(i)
      $("#delete-berita-modal").modal("show")
   }) 
   $("#deleteBeritaForm").submit(function (e) {
      e.preventDefault()
      data = $(this).serializeObject()
      $.ajax({
         url: "/delete-berita",
         method: "post",
         dataType: "JSON",
         data: JSON.stringify(data),
         success: function (resp) {
            if (resp == "sukses") {
               Snackbar.show({
                  text: 'Hapus Sukses',
                  backgroundColor: '#fff',
                  textColor: '#24D1BC',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               })
               $("#delete-berita-modal").modal("hide")
            } else {
               Snackbar.show({
                  text: 'Hapus Gagal',
                  // backgroundColor: '#fff',
                  textColor: '#ff69b4',
                  pos: 'top-right',
                  duration: '2000',
                  showAction: false,
               });
            }
         }
      })
      
   })
   $("#add-berita-btn").click(function () {
      $("#add-berita-modal").modal("show")
   })
   status_upload = ""
   $("#addBeritaForm").on("submit", function (e) {
      e.preventDefault();
      var $self = $(this);
      var files = $("#upload-file")[0].files;
      var formData = new FormData();
      for (var i = 0; i < files.length; i++) {
         formData.append("files", files[i]);
      }
      
      $.ajax({
         url: "/upload-foto",
         type: $self.attr("method"),
         data: formData,
         processData: false,
         contentType: false,
         mimeType: "multipart/form-data",
         success : function (res) {
            res = JSON.parse(res)
            res = JSON.parse(res)
            // (res)
            $("#foto-berita").val(res.NamaFoto)
            status_upload = res.Status
            if( res.Status == "oke"){
               // (res.Status)
               send_data(res.NamaFoto)
            }
         }
      })
      function send_data(NamaFoto){
         // ("NamaFoto send-data = "+ NamaFoto)
         $("#foto-berita").val(NamaFoto)
         data_final = $("#addBeritaForm").serializeObject()
         $.ajax({
            url: "/add-berita",
            type: $self.attr("method"),
            data: JSON.stringify(data_final),
            success: function(resp){
               if (resp == "sukses") {
                  Snackbar.show({
                     text: 'Tambah Sukses',
                     backgroundColor: '#fff',
                     textColor: '#24D1BC',
                     pos: 'top-right',
                     duration: '2000',
                     showAction: false,
                  })
                  $("#add-berita-modal").modal("hide")
               } else {
                  Snackbar.show({
                     text: 'Tambah Gagal',
                     // backgroundColor: '#fff',
                     textColor: '#ff69b4',
                     pos: 'top-right',
                     duration: '2000',
                     showAction: false,
                  });
               }
            }
         });  
      }
   })
})