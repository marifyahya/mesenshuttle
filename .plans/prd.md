# Product Requirements Document (PRD): Aplikasi Pemesanan Bis & Shuttle

## 1. Executive Summary
Aplikasi berbasis web untuk pemesanan tiket bis dan shuttle yang memungkinkan pengguna memesan tiket dengan cepat tanpa perlu melakukan registrasi akun (guest checkout). Sistem akan menggunakan alamat email sebagai penanda unik pengguna. Aplikasi ini dilengkapi dengan fitur pemilihan kursi interaktif, integrasi pembayaran otomatis, dan sistem *concurrency* yang tangguh untuk mencegah pemesanan ganda (double-booking). Terdapat juga portal Admin untuk mengelola operasional seperti rute dan jadwal.

## 2. Tech Stack
*   **Backend:** Golang
*   **Frontend:** Vue.js
*   **Database Persistent:** MySQL
*   **Database In-Memory:** Redis
*   **Payment Gateway:** Xendit
*   **Notifikasi:** Email (via SMTP)

## 3. User Personas
1.  **Customer (Penumpang):** Ingin memesan tiket bis/shuttle dengan cepat, memilih kursi, dan membayar tanpa hambatan registrasi.
2.  **Admin:** Staf operasional yang mengatur rute (trayek), armada bis, dan jadwal keberangkatan.

## 4. Alur Pengguna (User Flows)

### 4.1. Alur Customer (Tanpa Registrasi)
1.  **Pencarian (Search):** Pengguna mencari jadwal keberangkatan berdasarkan Pool Asal dan Pool Tujuan (dari daftar rute yang telah ditentukan oleh Admin), beserta tanggal keberangkatan.
2.  **Pemilihan Jadwal:** Sistem menampilkan daftar jadwal beserta tipe bis (**Premium** atau **Standart**). Pengguna memilih jadwal yang sesuai.
3.  **Pemilihan Kursi (Seat Map):** Pengguna melihat peta/layout kursi yang tersedia dan memilih kursi yang diinginkan.
4.  **Input Data Pemesan:** Pengguna memasukkan alamat Email (sebagai ID) dan data diri dasar (Nama, No HP).
5.  **Penguncian Kursi (Lock Seat):** Sistem melakukan *lock* sementara pada kursi yang dipilih menggunakan **Redis** agar tidak bisa dipilih oleh orang lain.
6.  **Pembayaran:** Sistem men-generate invoice via **Xendit** dan mengarahkan pengguna ke halaman pembayaran.
7.  **Konfirmasi:** Setelah Xendit mengirimkan *webhook* pembayaran berhasil, backend akan mengubah status tiket menjadi `PAID`.
8.  **Notifikasi:** E-Ticket dan bukti pembayaran dikirimkan ke email pengguna.

### 4.2. Alur Admin
1.  **Manajemen Trayek (Rute):** Admin dapat membuat, mengubah, dan menghapus rute perjalanan (Contoh: Jakarta - Bandung).
2.  **Manajemen Armada & Tipe Bis:** Admin mendefinisikan tipe armada beserta kapasitasnya. Layout kursi bersifat statis: **Standart (konfigurasi 1-2)** dan **Premium (konfigurasi 1-1)**.
3.  **Manajemen Jadwal:** Admin membuat jadwal keberangkatan dengan memasangkan Trayek, Armada, waktu keberangkatan, dan harga tiket.

## 5. Requirement Teknis & Concurrency (Race Condition)
*   **Redis Pool / Lua Scripting:** Saat beberapa pengguna mencoba memilih kursi yang sama di waktu yang bersamaan (*race condition*), backend Golang akan menggunakan Redis Lua Script. Operasi pengecekan ketersediaan kursi dan penguncian (*set with expiry*) akan dilakukan secara **atomik**. Ini memastikan secara absolut tidak ada dua pengguna yang mendapatkan kursi yang sama.
*   **Lock Expiration:** Jika pengguna tidak menyelesaikan pembayaran Xendit dalam batas waktu **30 menit**, Redis key akan otomatis kedaluwarsa (*expired*), dan kursi akan kembali berstatus `Available`.
*   **Webhook Handling:** Endpoint Golang untuk Xendit akan dibuat seaman mungkin (verifikasi token/signature) dan dapat memproses status *PAID* atau *EXPIRED* dengan cepat.
