CREATE TABLE `articles` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `user_id` BIGINT NOT NULL COMMENT '會員編號',
    `title` varchar(255) NOT NULL COMMENT '標題',
    `content` TEXT NOT NULL COMMENT '內容',
    `created_at` timestamp null,
    `updated_at` timestamp null,
    `deleted_at` timestamp null
)