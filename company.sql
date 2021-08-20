-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Aug 20, 2021 at 06:52 AM
-- Server version: 10.4.11-MariaDB
-- PHP Version: 7.4.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `company`
--

-- --------------------------------------------------------

--
-- Table structure for table `companies`
--

CREATE TABLE `companies` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '1- Active, 2- Inactive',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `companies`
--

INSERT INTO `companies` (`id`, `name`, `status`, `created_at`, `updated_at`) VALUES
(3, 'werwer', 1, '2021-07-02 05:58:25', '2021-07-02 05:58:25'),
(4, 'werwesdr', 1, '2021-07-02 05:59:10', '2021-07-02 05:59:10'),
(5, 'werwesdr', 1, '2021-07-02 05:59:13', '2021-07-02 05:59:13'),
(6, 'test', 1, '2021-07-02 09:55:31', '2021-07-02 09:55:31'),
(7, 'strssing', 1, '2021-07-02 10:00:50', '2021-07-02 10:00:50'),
(8, 'strssing', 1, '2021-07-02 10:01:31', '2021-07-02 10:01:31'),
(9, 'newuuu', 1, '2021-07-02 10:43:11', '2021-07-02 10:43:11'),
(10, 'werwesdr', 1, '2021-07-02 10:47:10', '2021-07-02 10:47:10'),
(11, 'adsad', 1, '2021-07-02 10:47:41', '2021-07-02 10:47:41'),
(12, 'adsad', 1, '2021-07-02 10:52:26', '2021-07-02 10:52:26'),
(13, 'dasdsad sad dsa das d', 1, '2021-07-02 10:52:34', '2021-07-02 10:52:34'),
(14, 'poiuiurtyretrrete  tbretretret', 1, '2021-07-02 10:53:27', '2021-07-02 10:53:27'),
(15, 'kjhkhkjhkjghjghhjg', 1, '2021-07-02 10:54:01', '2021-07-02 10:54:01'),
(16, 'adsd', 1, '2021-07-02 10:55:25', '2021-07-02 10:55:25'),
(17, 'tee', 1, '2021-07-15 05:52:53', '2021-07-15 05:52:53'),
(18, 'sed', 2, '2021-07-15 05:55:02', '2021-07-15 05:55:02'),
(19, 'sed', 2, '2021-07-15 06:05:12', '2021-07-15 06:05:12'),
(20, 'sd', 1, '2021-07-15 06:25:51', '2021-07-15 06:25:51'),
(21, 'sd', 1, '2021-07-15 06:26:04', '2021-07-15 06:26:04'),
(23, 'jkhqwjkqjklqwjklsakldddddddd', 1, '2021-07-15 06:40:26', '2021-07-15 06:40:48'),
(24, 'ewaee', 1, '2021-07-15 08:33:58', '2021-07-15 08:33:58'),
(25, 'ewaee', 1, '2021-07-15 08:34:00', '2021-07-15 08:34:00'),
(26, 'ewaee', 1, '2021-07-15 08:34:08', '2021-07-15 08:34:08'),
(27, 'tyry', 1, '2021-07-15 08:35:43', '2021-07-15 08:35:43'),
(28, 'sdfdsf', 1, '2021-07-15 08:39:14', '2021-07-15 08:39:14'),
(29, 'ass', 1, '2021-07-15 08:40:09', '2021-07-15 08:40:09'),
(31, 'sddddddddds', 1, '2021-07-15 09:06:42', '2021-07-15 09:06:42'),
(32, 'assa', 1, '2021-07-15 09:08:14', '2021-07-15 09:08:14'),
(33, 'string', 1, '2021-07-15 11:22:31', '2021-07-15 11:22:31'),
(34, 'ds', 0, '2021-08-19 06:17:19', '2021-08-19 06:17:19'),
(35, 'strisng', 0, '2021-08-20 04:39:03', '2021-08-20 04:39:03'),
(36, 'string', 0, '2021-08-20 04:52:04', '2021-08-20 04:52:04');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` tinyint(1) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `email`, `password`, `role`, `created_at`, `updated_at`) VALUES
(1, 'chetan.meniya@bacancy.com', 'e10adc3949ba59abbe56e057f20f883e', 2, '2021-06-21 07:55:29', '2021-06-21 12:22:34');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `companies`
--
ALTER TABLE `companies`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `companies`
--
ALTER TABLE `companies`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=37;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
