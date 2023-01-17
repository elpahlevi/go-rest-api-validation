# Golang REST-API Validation: Part 2

## Deskripsi
Repository ini berisi source code dari artikel bagian 2 yang di rilis pada platform Medium.
<br>Link: https://elpahlevi.medium.com/implementasi-validasi-input-ke-dalam-rest-api-di-golang-e185f26ad546

## Cara Menggunakan
- Buka terminal/command prompt dan jalankan perintah berikut:
<pre>
<code>go run main.go</code>
</pre>
- Server berjalan di port <code>8080</code>

### Endpoint:
1. `http://localhost:8080/api/v1/auth/register`
<br> Method: POST
<br> Request Body (JSON):

| Field    | Tipe Data | Kriteria       | Catatan                        |
|----------|-----------|----------------|--------------------------------|
| name     | string    | required,min=2 | -                              |
| email    | string    | required,email | -                              |
| phone    | string    | required,phone | format nomor telepon Indonesia |
| password | string    | required,min=8 | -                              |

Contoh request:
<pre>
<code>
{
    "name: "John Mail",
    "email:"john@mail.com",
    "phone": "6281234564321",
    "password: "xyzabcd123"
}
</code>
</pre>

## Credit
Copyright (c) 2023-present <a href="https://github.com/elpahlevi">Reza Pahlevi</a>.