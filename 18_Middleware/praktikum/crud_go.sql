-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 25 Sep 2023 pada 06.51
-- Versi server: 10.4.22-MariaDB
-- Versi PHP: 7.3.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `crud_go`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `blogs`
--

CREATE TABLE `blogs` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` int(10) UNSIGNED DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `blogs`
--

INSERT INTO `blogs` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `title`, `content`) VALUES
(1, '2023-09-24 14:04:36', '2023-09-24 14:04:36', NULL, 1, 'Pengenalan Golang', 'Golang adalah bahasa pemrograman yang dikembangkan oleh Google.'),
(2, '2023-09-24 14:05:00', '2023-09-24 14:05:00', NULL, 2, 'Tutorial Echo Framework', 'Echo adalah framework web berbasis Go yang ringan dan cepat.'),
(3, '2023-09-24 14:05:09', '2023-09-24 22:31:50', NULL, 1, 'Menggunakan GORM', 'GORM adalah ORM (Object-Relational Mapping) untuk Go yang powerful.'),
(4, '2023-09-24 14:05:13', '2023-09-24 14:05:13', '2023-09-24 19:38:45', 3, 'Pemrograman Web dengan Go', 'Go adalah bahasa yang sangat cocok untuk pengembangan web modern.'),
(5, '2023-09-24 22:16:32', '2023-09-25 08:07:07', NULL, 1, 'Menggunakan GORM', 'GORM adalah ORM (Object-Relational Mapping) untuk Go yang powerful.');

-- --------------------------------------------------------

--
-- Struktur dari tabel `books`
--

CREATE TABLE `books` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `author` varchar(255) DEFAULT NULL,
  `publisher` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `books`
--

INSERT INTO `books` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `author`, `publisher`) VALUES
(1, '2023-09-24 12:20:17', '2023-09-25 12:43:14', '2023-09-25 12:43:24', 'Pulang-Pergi', 'Tere Liye', 'Republika'),
(2, '2023-09-24 12:20:57', '2023-09-24 12:20:57', NULL, 'Pergi', 'Tere Liye', 'Republika'),
(3, '2023-09-24 12:21:04', '2023-09-24 12:21:04', '2023-09-24 12:23:13', 'Pulang-Pergi', 'Tere Liye', 'Republika'),
(4, '2023-09-25 12:42:52', '2023-09-25 12:42:52', NULL, 'Ananta Prahadi', 'Risa Saraswati', 'Bukune');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `email`, `password`) VALUES
(1, '2023-09-23 23:31:14', '2023-09-24 22:31:50', NULL, 'Lee Mujin', 'morilla_lmj@gmail.com', '28Des2000'),
(2, '2023-09-23 23:31:57', '2023-09-23 23:31:57', NULL, 'Jo Byeonggyu', 'bk_arta@gmail.com', '23Apr1996'),
(3, '2023-09-23 23:32:25', '2023-09-23 23:34:21', NULL, 'Kim Beom', 'k.kbeom@gmail.com', '7Jul1993'),
(4, '2023-09-24 00:13:03', '2023-09-24 12:03:42', NULL, 'Do Kyungsoo', 'd.ohkyungsoo@gmail.com', '12Jan1993'),
(5, '2023-09-24 00:18:11', '2023-09-24 00:18:11', NULL, 'Choi Minho', 'choiminho_1209@gmail.com', '9Des1991'),
(6, '2023-09-25 12:36:14', '2023-09-25 12:44:36', '2023-09-25 12:44:45', 'Kim Kibum', 'key_shinee@gmail.com', '23Sep1991');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `blogs`
--
ALTER TABLE `blogs`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_blogs_deleted_at` (`deleted_at`);

--
-- Indeks untuk tabel `books`
--
ALTER TABLE `books`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_books_deleted_at` (`deleted_at`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `blogs`
--
ALTER TABLE `blogs`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT untuk tabel `books`
--
ALTER TABLE `books`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
