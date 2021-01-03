-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: mysql
-- Generation Time: Jan 03, 2021 at 06:03 AM
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
(1, 1, 1, 1, 1, 1200.00, 2, '2021-01-03 12:36:29', '2021-01-03 12:55:56');

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
(1, 1, 1, 1, 1, 1, 1, 800.00, 'Madura', '2021-01-03 12:36:29', '2021-01-03 12:38:46');

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
(1, 'Fisher 1', '123124124213', 'Bogor', '2021-01-03 12:33:25', '2021-01-03 12:33:25');

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
(1, 'Pancingan', '2021-01-03 12:34:26', '2021-01-03 12:34:26');

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
(1, 'Ikan Asin', '2021-01-03 12:34:19', '2021-01-03 12:34:19');

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
(1, 1, 1, 1, 1, 12000, 'Malang', '2021-01-03 13:00:12', '2021-01-03 13:00:12');

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
(1, 'Ton', '2021-01-03 12:34:34', '2021-01-03 12:34:34');

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
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

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
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `fisher`
--
ALTER TABLE `fisher`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `fishing_gear`
--
ALTER TABLE `fishing_gear`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `fish_type`
--
ALTER TABLE `fish_type`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `transaction`
--
ALTER TABLE `transaction`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `weight_unit`
--
ALTER TABLE `weight_unit`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
