CREATE TABLE `tags`(
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `title` varchar(255) NOT NULL COMMENT '名稱',
    `created_at` timestamp NULL,
    `updated_at` timestamp NULL,
    `deleted_at` timestamp NULL
);