-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: db
-- Generation Time: Dec 30, 2020 at 10:16 AM
-- Server version: 5.7.31
-- PHP Version: 7.4.8

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
CREATE DATABASE IF NOT EXISTS `sitpi` DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;
USE `sitpi`;

-- --------------------------------------------------------

--
-- Table structure for table `auction`
--

DROP TABLE IF EXISTS `auction`;
CREATE TABLE `auction` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `tpi_id` int(11) NOT NULL,
  `officer_id` int(11) NOT NULL,
  `caught_fish_id` int(11) NOT NULL,
  `weight` double(8,2) NOT NULL,
  `weight_unit` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `auction`
--

INSERT INTO `auction` (`id`, `tpi_id`, `officer_id`, `caught_fish_id`, `weight`, `weight_unit`, `status`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 1, 12.00, 'kg', 1, '2020-12-28 19:36:04', '2020-12-28 19:36:04'),
(4, 2, 2, 2, 33.00, 'Kwintal', 0, '2020-12-28 19:12:39', '2020-12-28 19:47:54'),
(5, 2, 2, 0, 20.00, 'Kwintal', 1, '2020-12-28 19:16:08', '2020-12-28 19:16:08'),
(6, 0, 0, 0, 12.00, 'kg', 1, '2020-12-28 21:22:44', '2020-12-28 21:22:44'),
(7, 1, 1, 5, 600.00, 'Ton', 2, '2020-12-28 22:08:27', '2020-12-28 22:08:27'),
(8, 1, 1, 6, 70.00, 'Kwintal', 0, '2020-12-28 22:21:34', '2020-12-28 22:22:04'),
(9, 1, 1, 7, 89.00, 'Kg', 0, '2020-12-28 22:27:27', '2020-12-28 22:28:39'),
(10, 1, 1, 8, 70.00, 'Kg', 1, '2020-12-29 09:01:27', '2020-12-29 09:01:47');

-- --------------------------------------------------------

--
-- Table structure for table `auction_status`
--

DROP TABLE IF EXISTS `auction_status`;
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
(1, 'Belum Terjual', '2020-12-28 18:14:51', '2020-12-28 18:14:51'),
(2, 'Sudah Terjual', '2020-12-28 18:15:17', '2020-12-28 18:15:17');

-- --------------------------------------------------------

--
-- Table structure for table `buyer`
--

DROP TABLE IF EXISTS `buyer`;
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
(1, 'Yudit', '765', 'Bogor', '2020-12-28 19:11:50', '2020-12-28 19:11:50');

-- --------------------------------------------------------

--
-- Table structure for table `caught_fish`
--

DROP TABLE IF EXISTS `caught_fish`;
CREATE TABLE `caught_fish` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `tpi_id` int(11) NOT NULL,
  `officer_id` int(11) NOT NULL,
  `fisher_id` int(11) NOT NULL,
  `fish_type_id` int(11) NOT NULL,
  `weight` double(8,2) NOT NULL,
  `weight_unit` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `fishing_gear` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `fishing_area` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `caught_fish`
--

INSERT INTO `caught_fish` (`id`, `tpi_id`, `officer_id`, `fisher_id`, `fish_type_id`, `weight`, `weight_unit`, `fishing_gear`, `fishing_area`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 1, 1, 120.00, 'kg', 'pancingan', 'laut', '2020-12-28 18:05:00', '2020-12-28 18:05:00'),
(2, 2, 2, 1, 1, 555.00, 'Kwintal', 'Pancingan', 'Laut', '2020-12-28 19:12:39', '2020-12-28 19:49:04'),
(3, 2, 2, 1, 1, 120.00, 'Kwintal', 'Pancingan', 'Sumatra', '2020-12-28 19:16:08', '2020-12-28 19:16:08'),
(4, 1, 1, 1, 1, 120.00, 'kg', 'Pancingan', 'Kep. Seribu', '2020-12-28 21:07:59', '2020-12-28 21:07:59'),
(5, 1, 1, 1, 1, 1200.00, 'Ton', 'Kail', 'Samudera Hindia', '2020-12-28 22:08:27', '2020-12-28 22:08:27'),
(7, 1, 1, 1, 1, 899.00, 'Kg', 'Pancingan', 'Laut', '2020-12-28 22:27:27', '2020-12-28 22:29:48'),
(9, 1, 1, 1, 1, 899.00, 'Kwintal', 'Pancingan', 'Aceh', '2020-12-29 09:03:05', '2020-12-29 09:03:05');

-- --------------------------------------------------------

--
-- Table structure for table `fisher`
--

DROP TABLE IF EXISTS `fisher`;
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
(1, 'Yudit', '3214022207990014', 'Candrabaga', '2020-12-28 18:04:45', '2020-12-28 18:04:45');

-- --------------------------------------------------------

--
-- Table structure for table `fish_type`
--

DROP TABLE IF EXISTS `fish_type`;
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
(1, 'Tenggiri', '2020-12-28 18:14:08', '2020-12-28 18:14:08');

-- --------------------------------------------------------

--
-- Table structure for table `transaction`
--

DROP TABLE IF EXISTS `transaction`;
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
(1, 2, 2, 1, 1, 150000, 'Surabaya', '2020-12-28 19:37:09', '2020-12-28 19:43:26'),
(3, 1, 1, 1, 1, 12000, 'malang', '2020-12-28 21:23:21', '2020-12-28 21:23:21'),
(4, 1, 1, 7, 1, 12000, 'malang', '2020-12-28 22:09:50', '2020-12-28 22:09:50');

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
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `auction`
--
ALTER TABLE `auction`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `auction_status`
--
ALTER TABLE `auction_status`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `buyer`
--
ALTER TABLE `buyer`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `caught_fish`
--
ALTER TABLE `caught_fish`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `fisher`
--
ALTER TABLE `fisher`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `fish_type`
--
ALTER TABLE `fish_type`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `transaction`
--
ALTER TABLE `transaction`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
