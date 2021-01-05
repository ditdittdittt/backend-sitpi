-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: mysql
-- Generation Time: Jan 05, 2021 at 10:07 AM
-- Server version: 5.7.32
-- PHP Version: 7.4.13

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `sitpi`
--

-- --------------------------------------------------------

--
-- Table structure for table `auction`
--

CREATE TABLE `auction` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `tpi_id` int(11) NOT NULL,
  `officer_id` int(11) NOT NULL,
  `caught_fish_id` int(11) NOT NULL,
  `weight_unit_id` int(11) NOT NULL,
  `weight` double(8,2) NOT NULL,
  `status` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `auction`
--

INSERT INTO `auction` (`id`, `tpi_id`, `officer_id`, `caught_fish_id`, `weight_unit_id`, `weight`, `status`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 1, 1, 1200.00, 2, '2021-01-03 12:36:29', '2021-01-03 12:55:56'),
(2, 1, 1, 2, 1, 12345.00, 1, '2021-01-03 15:37:20', '2021-01-03 15:37:20'),
(3, 1, 1, 3, 1, 700.00, 2, '2021-01-03 15:37:56', '2021-01-03 16:13:14');

-- --------------------------------------------------------

--
-- Table structure for table `auction_status`
--

CREATE TABLE `auction_status` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `status` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `auction_status`
--

INSERT INTO `auction_status` (`id`, `status`, `created_at`, `updated_at`) VALUES
(1, 'Belum Terjual', '2021-01-03 12:53:25', '2021-01-03 12:53:25'),
(2, 'Sudah Terjual', '2021-01-03 12:53:49', '2021-01-03 12:53:49');

-- --------------------------------------------------------

--
-- Table structure for table `buyer`
--

CREATE TABLE `buyer` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nik` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `buyer`
--

INSERT INTO `buyer` (`id`, `name`, `nik`, `address`, `created_at`, `updated_at`) VALUES
(1, 'Buyer 1', '123124213123', 'Candrabga', '2021-01-03 12:33:44', '2021-01-03 12:33:44');

-- --------------------------------------------------------

--
-- Table structure for table `caught_fish`
--

CREATE TABLE `caught_fish` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `tpi_id` int(11) NOT NULL,
  `officer_id` int(11) NOT NULL,
  `fisher_id` int(11) NOT NULL,
  `fish_type_id` int(11) NOT NULL,
  `weight_unit_id` int(11) NOT NULL,
  `fishing_gear_id` int(11) NOT NULL,
  `weight` double(8,2) NOT NULL,
  `fishing_area` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `caught_fish`
--

INSERT INTO `caught_fish` (`id`, `tpi_id`, `officer_id`, `fisher_id`, `fish_type_id`, `weight_unit_id`, `fishing_gear_id`, `weight`, `fishing_area`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 1, 1, 1, 1, 800.00, 'Madura', '2021-01-03 12:36:29', '2021-01-03 12:38:46'),
(3, 1, 1, 2, 1, 2, 1, 1200.00, 'Samudera Hindia', '2021-01-03 15:37:56', '2021-01-03 15:37:56'),
(4, 1, 1, 1, 1, 3, 1, 20.00, 'Maluku', '2021-01-03 15:38:19', '2021-01-03 15:46:03');

-- --------------------------------------------------------

--
-- Table structure for table `fisher`
--

CREATE TABLE `fisher` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `nik` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `fisher`
--

INSERT INTO `fisher` (`id`, `name`, `nik`, `address`, `created_at`, `updated_at`) VALUES
(1, 'Fisher 1', '123124124213', 'Bogor', '2021-01-03 12:33:25', '2021-01-03 12:33:25'),
(2, 'Fisher 2', '123444445555', 'Home', '2021-01-03 16:16:10', '2021-01-03 16:16:10');

-- --------------------------------------------------------

--
-- Table structure for table `fishing_gear`
--

CREATE TABLE `fishing_gear` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `fishing_gear`
--

INSERT INTO `fishing_gear` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Sodo', '2021-01-03 12:34:26', '2021-01-03 12:34:26'),
(2, 'Songko', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(3, 'Bubu', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(4, 'Pakaja', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(5, 'Sero Besar', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(6, 'Tugu Ganda', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(7, 'Jernal', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(8, 'Mourami', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(9, 'Jaring Kepiting', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(10, 'Jaring Rajungan', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(11, 'Dogol', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(12, 'Pancing Tonda', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(13, 'Bagan Tancap', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(14, 'Bagan Perahu', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(15, 'Payang Lampu', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(16, 'Payang Rumpon', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(17, 'Soma Dampar', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(18, 'Pukat Tepi', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(19, 'Pukat Harimau', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(20, 'Trawl Dasar', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(21, 'Trawl Udang Ganda', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(22, 'Trawl Udang BED', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(23, 'Jaring Insang Tetap', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(24, 'Jaring Insang Hanyut', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(25, 'Jaring Gondrong', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(26, 'Jaring Insang Lingkar', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(27, 'Pukat Cincin Rumpon', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(28, 'Pukat Cincin Lampu', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(29, 'Soma', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(30, 'Pukat Cincin Cakalang', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(31, 'Pancing Rawai Dasar', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(32, 'Pancing Rawai Tuna', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(33, 'Rumpon Laut Dalam', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(34, 'Huhate', '2021-01-04 00:55:01', '2021-01-04 00:55:01'),
(35, 'Pukat Cincin', '2021-01-04 00:55:01', '2021-01-04 00:55:01');

-- --------------------------------------------------------

--
-- Table structure for table `fish_type`
--

CREATE TABLE `fish_type` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `fish_type`
--

INSERT INTO `fish_type` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Layang', '2021-01-03 12:34:19', '2021-01-03 12:34:19'),
(2, 'Bawal', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(3, 'Kembung', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(4, 'Selar', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(5, 'Tembang', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(6, 'Udang Barong', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(7, 'Udang Windu', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(8, 'Udang Jrebung', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(9, 'Udang Dogol', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(10, 'Udang Lainnya', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(11, 'Teri', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(12, 'Tongkol', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(13, 'Kurisi', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(14, 'Lemuru', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(15, 'Cakalang', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(16, 'Tenggir', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(17, 'Layur', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(18, 'Ikan Terbang', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(19, 'Julung-Julung', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(20, 'Tiga Waja', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(21, 'Ekor Kuning', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(22, 'Ikan Kowe', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(23, 'Petek/Peperek', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(24, 'Manyung', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(25, 'Songot', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(26, 'Cucut', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(27, 'Pari', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(28, 'Kakap', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(29, 'Sunglir', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(30, 'Bambangan', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(31, 'Kerapu', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(32, 'Kurau', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(33, 'Belanak', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(34, 'Tuna', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(35, 'Cumi-Cumi', '2021-01-03 14:40:30', '2021-01-03 14:40:30'),
(36, 'Lainnya', '2021-01-03 14:40:30', '2021-01-03 14:40:30');

-- --------------------------------------------------------

--
-- Table structure for table `transaction`
--

CREATE TABLE `transaction` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `tpi_id` int(11) NOT NULL,
  `officer_id` int(11) NOT NULL,
  `auction_id` int(11) NOT NULL,
  `buyer_id` int(11) NOT NULL,
  `price` int(11) NOT NULL,
  `distribution_area` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `transaction`
--

INSERT INTO `transaction` (`id`, `tpi_id`, `officer_id`, `auction_id`, `buyer_id`, `price`, `distribution_area`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 1, 1, 12000, 'Malang', '2021-01-03 13:00:12', '2021-01-03 13:00:12'),
(3, 1, 1, 3, 1, 150000000, 'Malang', '2021-01-03 20:21:17', '2021-01-03 20:21:17');

-- --------------------------------------------------------

--
-- Table structure for table `weight_unit`
--

CREATE TABLE `weight_unit` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `unit` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `weight_unit`
--

INSERT INTO `weight_unit` (`id`, `unit`, `created_at`, `updated_at`) VALUES
(1, 'Ton', '2021-01-03 12:34:34', '2021-01-03 12:34:34'),
(2, 'Kwintal', '2021-01-03 13:22:24', '2021-01-03 13:22:24'),
(3, 'Kg', '2021-01-03 13:22:48', '2021-01-03 13:22:48');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `auction`
--
ALTER TABLE `auction`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `auction_status`
--
ALTER TABLE `auction_status`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `buyer`
--
ALTER TABLE `buyer`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `caught_fish`
--
ALTER TABLE `caught_fish`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `fisher`
--
ALTER TABLE `fisher`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `fishing_gear`
--
ALTER TABLE `fishing_gear`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `fish_type`
--
ALTER TABLE `fish_type`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transaction`
--
ALTER TABLE `transaction`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `weight_unit`
--
ALTER TABLE `weight_unit`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `auction`
--
ALTER TABLE `auction`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `auction_status`
--
ALTER TABLE `auction_status`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `buyer`
--
ALTER TABLE `buyer`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `caught_fish`
--
ALTER TABLE `caught_fish`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `fisher`
--
ALTER TABLE `fisher`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `fishing_gear`
--
ALTER TABLE `fishing_gear`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=36;

--
-- AUTO_INCREMENT for table `fish_type`
--
ALTER TABLE `fish_type`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=37;

--
-- AUTO_INCREMENT for table `transaction`
--
ALTER TABLE `transaction`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `weight_unit`
--
ALTER TABLE `weight_unit`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
