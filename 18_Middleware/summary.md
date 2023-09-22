# (18) Middleware

Middleware merupakan entitiy yang dipasangkan ke dalam proses server. Middleware memiliki fungsi-fungsi khusus yang akan membantu komunikasi klien dan server. Terdapat 4 layer fungsi middleware, yaitu :
- LoggingMiddleware => mencatat log-log yang keluar masuk server
- SessionMiddleware => mencatat aktif tidaknya user di dalam aplikasi
- AuthMiddleware => mengauntentifikasi user berhak masuk atau tidak
- CustomMiddleware => Jika third party middleware tidak menyediakan fungsi yang diinginkan, maka dapat dibuat sendiri

Contoh Third Party Middleware :
- Negroni
- Echo
- Interpose
- Alice

Middleware di echo dibagi menjadi 2 jenis, yaitu :
- Echo#Pre() => Dieksekusi sebelum route nya masuk ke controller. Contoh : HTTPSRedirect, AddTrailingSlash, dll.
- Echo#Use() => Dieksekusi saat route nya masuk ke controller. Contoh : BodyLimit, Logger, JWTAuth, Gzip, Static, CORS, dll.

Fungsi di Echo Middleware :
- HTTPSRedirect => mengarahkan user ke HTTPS untuk keamanan.
- Logger (Logger Middleware) => mencatat apa yang terjadi di dalam server khususnya request.
- JWT (Auth Middleware) => sebuah proses dalam API untuk mengidentifikasi user dari sisi server dengan cara memvalidasi sebelum mengakses API agar keamanan data terjamin. 

Basic Authentication sebuah teknik autentikasi user yang membutuhkan data username dan password dari user untuk di insert ke request header. Aturan penulisan pada Basic Authentication :
`'Authorization: Basic' + base64encode('username:password')`

Jika Basic Authentication hanya terdiri dari satu token, JWT terdiri dari 3 token yang dipisahkan tanda titik, yaitu header (algoritma dan type), payload (data), dan verify signature. Aturan penulisan JWT :
`'Authorization: Bearer' + token`