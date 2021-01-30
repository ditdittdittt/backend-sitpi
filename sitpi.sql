-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: mysql
-- Generation Time: Jan 27, 2021 at 02:04 AM
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
  `caught_fish_id` int(11) NOT NULL,
  `status_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `auction`
--

INSERT INTO `auction` (`id`, `tpi_id`, `caught_fish_id`, `status_id`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 2, '2021-01-27 08:07:45', '2021-01-27 08:07:45'),
(2, 1, 1, 1, '2021-01-27 08:07:45', '2021-01-27 08:07:45');

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
(1, 'Belum Terjual', '2021-01-23 11:17:57', '2021-01-23 11:17:57'),
(2, 'Sudah Terjual', '2021-01-23 11:18:15', '2021-01-23 11:18:15');

-- --------------------------------------------------------

--
-- Table structure for table `buyer`
--

CREATE TABLE `buyer` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `user_id` int(11) NOT NULL,
  `nik` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `buyer`
--

INSERT INTO `buyer` (`id`, `user_id`, `nik`, `name`, `address`, `created_at`, `updated_at`) VALUES
(1, 1, '123124124213', 'Buyer 1', 'Updated', '2021-01-23 12:09:33', '2021-01-23 12:14:31'),
(2, 1, '123124213123', 'Buyer 2', 'Candrabaga', '2021-01-23 12:11:11', '2021-01-23 12:11:11');

-- --------------------------------------------------------

--
-- Table structure for table `caught_fish`
--

CREATE TABLE `caught_fish` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `user_id` int(11) NOT NULL,
  `tpi_id` int(11) NOT NULL,
  `fisher_id` int(11) NOT NULL,
  `fish_type_id` int(11) NOT NULL,
  `weight_unit_id` int(11) NOT NULL,
  `fishing_gear_id` int(11) NOT NULL,
  `fishing_area_id` int(11) NOT NULL,
  `weight` double(8,2) NOT NULL,
  `trip_day` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `caught_fish`
--

INSERT INTO `caught_fish` (`id`, `user_id`, `tpi_id`, `fisher_id`, `fish_type_id`, `weight_unit_id`, `fishing_gear_id`, `fishing_area_id`, `weight`, `trip_day`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 1, 1, 1, 1, 1, 800.00, 12, '2021-01-27 08:07:45', '2021-01-27 08:14:41');

-- --------------------------------------------------------

--
-- Table structure for table `district`
--

CREATE TABLE `district` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `location` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `fisher`
--

CREATE TABLE `fisher` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `user_id` int(11) NOT NULL,
  `nik` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `address` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ship_type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `abk_total` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `fisher`
--

INSERT INTO `fisher` (`id`, `user_id`, `nik`, `name`, `address`, `ship_type`, `abk_total`, `created_at`, `updated_at`) VALUES
(1, 1, '3214022207990014', 'Buyer 1', 'Candrabaga', '3 motor', '30', '2021-01-23 11:19:41', '2021-01-23 11:48:34');

-- --------------------------------------------------------

--
-- Table structure for table `fishing_area`
--

CREATE TABLE `fishing_area` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `district_id` int(11) NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `south_latitude_degree` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `south_latitude_minute` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `south_latitude_second` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `east_longitude_degree` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `east_longitude_minute` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `east_longitude_second` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `fishing_area`
--

INSERT INTO `fishing_area` (`id`, `district_id`, `name`, `south_latitude_degree`, `south_latitude_minute`, `south_latitude_second`, `east_longitude_degree`, `east_longitude_minute`, `east_longitude_second`, `created_at`, `updated_at`) VALUES
(1, 1, 'Indramayu', '90', 'aaaaaa', '123123', '12312', '1231', '123', '2021-01-23 18:53:05', '2021-01-23 18:55:28');

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
(1, 'Kail', '2021-01-23 12:45:48', '2021-01-23 12:46:27');

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
(1, 'Tenggiri', '2021-01-27 08:08:56', '2021-01-27 08:08:56');

-- --------------------------------------------------------

--
-- Table structure for table `tpi`
--

CREATE TABLE `tpi` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `district_id` int(11) NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `location` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `tpi`
--

INSERT INTO `tpi` (`id`, `district_id`, `name`, `location`, `created_at`, `updated_at`) VALUES
(1, 1, 'TPI Indramayu', 'Indramayu', '2021-01-27 08:08:16', '2021-01-27 08:08:16');

-- --------------------------------------------------------

--
-- Table structure for table `transaction`
--

CREATE TABLE `transaction` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `user_id` int(11) NOT NULL,
  `tpi_id` int(11) NOT NULL,
  `auction_id` int(11) NOT NULL,
  `buyer_id` int(11) NOT NULL,
  `distribution_area` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `price` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `transaction`
--

INSERT INTO `transaction` (`id`, `user_id`, `tpi_id`, `auction_id`, `buyer_id`, `distribution_area`, `price`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 1, 1, 'Surabaya', 150000, '2021-01-27 08:43:13', '2021-01-27 08:48:43');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `name`, `email`, `email_verified_at`, `password`, `remember_token`, `created_at`, `updated_at`) VALUES
(1, 'Petugas 1', 'petugas@mail.com', '2021-01-23 11:21:03', '123456', '123456', '2021-01-23 11:21:03', '2021-01-23 11:21:03');

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
(1, 'Kg', '2021-01-27 08:09:19', '2021-01-27 08:09:19');

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
-- Indexes for table `district`
--
ALTER TABLE `district`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `fisher`
--
ALTER TABLE `fisher`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `fishing_area`
--
ALTER TABLE `fishing_area`
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
-- Indexes for table `tpi`
--
ALTER TABLE `tpi`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transaction`
--
ALTER TABLE `transaction`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `users_email_unique` (`email`);

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
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

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
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `district`
--
ALTER TABLE `district`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `fisher`
--
ALTER TABLE `fisher`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `fishing_area`
--
ALTER TABLE `fishing_area`
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
-- AUTO_INCREMENT for table `tpi`
--
ALTER TABLE `tpi`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `transaction`
--
ALTER TABLE `transaction`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
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
