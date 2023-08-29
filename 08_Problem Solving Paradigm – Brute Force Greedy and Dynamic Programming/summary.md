# (8) Problem Solving Paradigm – Brute Force Greedy and Dynamic Programming

## Problem Solving Paradigm
Problem Solving Paradigm merupakan prinsip yang digunakan untuk menyelesaikan masalah. Dari banyaknya Paradigm, ada 4 yang sering digunakan yaitu Brute Force, Divide and Conquer, Greedy, dan Dynamic Programming.

## Brute Force
Brute Force merupakan Problem Solving dengan melakukan pengecekan seluruh data. Source yang dibutuhkan Brute Force sangat besar dan menjadi kurang efektif karena Brute Force benar-benar mengecek seluruh data walaupun jumlah datanua sangat besar. Brute Force cocok digunakan ketika tidak ada algoritma lain yang lebih cepat darinya atau lebih cepat dari O(n).

## Divide and Conquer
Divide and Conquer cara kerjanya dengan membagi data secara terus menerus menjadi 2 bagian. Setelah selesai pembagian, data akan digabungkan lagi satu persatu hingga solusi didapatkan. Notasi kompleksitas dari Divide and Conquer adalah O(log n).

## Greedy
Greedy adalah algoritma dengan metode problem solving mencari nilai local optimal. Data yang biasa digunakan berbentuk graph. Local optimal adalah nilai terkecil dari pilihan lainnya. Greedy bisa jadi bukan algoritma yang paling optimal.

## Dynamic Programming
Dynamic Programming adalah algoritma problem solving seoptimal mungkin dengan memecah masalah ke dalam bagian yang lebih kecil dan mudah dikerjakan namun harus memperhatikan fakta optimal dari setiap masalah. Karakteristik Dynamic Programming :
1. Overlapping Subproblems
2. Optimal Substructure Property

Method in Dynamic Programming :
1. Top-Down with Memoization => memecahkan masalah dengan rekursif, memecahkan masalah hingga ke subproblem yang terkecil, menyimpan datanya ke sub data temporary untuk dipanggil saat memecahkan masalah yang lebih besar.
2. Bottom-up with Tabulation => menyelesaikan masalah dari yang terkecil