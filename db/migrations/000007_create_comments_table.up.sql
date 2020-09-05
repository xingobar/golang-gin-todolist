CREATE TABLE `comments` (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `parent_id` BIGINT NULL COMMENT '回覆哪條訊息',
    `user_id` BIGINT NOT NULL COMMENT '留言的使用者',
    `content` TEXT NOT NULL COMMENT '留言內容',
    `created_at` timestamp NULL COMMENT '新增時間',
    `updated_at` timestamp NULL COMMENT '更新時間',
    `deleted_at` timestamp NULL COMMENT '刪除時間',
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);