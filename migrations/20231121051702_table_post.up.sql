CREATE TABLE IF NOT EXISTS
  `post` (
    `id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `tiltle` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
    `information` json NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `post_fk_1` (`created_by`),
    CONSTRAINT `post_fk_1` FOREIGN KEY (`created_by`) REFERENCES `account` (`id`)
  ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci