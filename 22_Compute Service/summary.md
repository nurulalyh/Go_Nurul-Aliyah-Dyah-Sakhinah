
# (22) Compute Services

Deployment adalah kegiatan yang bertujuan untuk menyebarkan aplikasi/produk yang telah dikerjakan oleh developer. Penyebarannya tergantung jenis aplikasi, aplikasi web dan api ke server, sedangkan app mobile ke playstore/Appstore

Strategi Deployment :
1. Big-Bang atau Replace
    - Mudah di implementasikan
    - Perubahan pada sistem 100% secara instan
    - Namun terlalu beresiko, rata-rata downtime cukup lama
2. Rollout
    - Lebih aman dan less downtime dari Big-Bang
    - Namun bisa membuat bingung karena akan ada 2 aplikasi yang berjalan barengan hingga semua server terdeploy
    - Deployment dan rollback lebih lama dibandingkan Big-Bang
    - Tidak adanya kontrol request
3. Blue/Green
    - Perubahan sangat cepat, sekali switch service langsung berubah 100%
    - Tidak ada issue beda versi pada service seperti yang terjadi pada Rollout
    - Resource yang dibutuhkan lebih banyak
    - Testing harus diprioritaskan sebelum switch karena aplikasi harus dipastikan aman dari request yang tiba-tiba bertambah
4. Canary
    - Cukup aman
    - mudah untuk rollback jika terjadi error tanpa berimbas ke user
    - Untuk mencapai 100% cukup lama dibanding dengan Blue/Green