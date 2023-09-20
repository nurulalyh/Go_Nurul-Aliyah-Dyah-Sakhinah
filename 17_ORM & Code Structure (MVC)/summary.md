# (15) RESTful API

API (Application Programming Interface) adalah kumpulan function dan prosedur yang memungkinkan sebuah aplikasi yang mengakses fitur atau data suatu sistem operasi, aplikasi, atau layanan lainnya. API itu seperti messenger yang mengambil permintaan dari aplikasi dan meneruskannya ke server dan server akan memberikan apa yang diminta ke API dan API yang mengembalikannya ke aplikasi.

REST (Representational State Transfer) salah satu jenis API yang menggunakan protokol HTTP. Request dan response nya dalam bentuk JSON (Javascript Object Notation), XML, dan SOAP. Terdapat beberapa method yang dapat digunakan dalam RESTful API yaitu Get, Post, Put, Delete, Head, Option, dan Patch.

RESTful API digunakan untuk mengintegrasikan aplikasi berbasis frontend (mobile dan web) dengan yang backend. Keunggulan RESTful API adalah dengan 1 backend, kita dapat menjawab semua request dari berbagai aplikasi frontend. RESTful API juga dapat digunakan untuk Integrasi backend - backend.

Jenis HTTP Response Code :
1. 200 - OK
2. 201 - Created
3. 400 - Bad Request
4. 404 - Not Found
5. 401 - Unauthorized
6. 405 - Method Not Allowed
7. 500 - Internal Server Error

API testing tools :
1. Katalon
2. Apache JMeter
3. Postman => HTTP client untuk melakukan testing web services.
4. SoapUI
5. Insomnia

REST API Design Best Practice :
1. menggunakan kata benda, bukan kata kerja.
2. penamaan menggunakan kata benda jamak
3. menggunakan resource nesting untuk menunjukkan hirarki

Package go yang digunakan :
1. net/http => menyediakan implementasi http client and server.
2. encoding/json => decode json ke object struct, map[string]interface, dan interface[]. Encode object ke json string.
