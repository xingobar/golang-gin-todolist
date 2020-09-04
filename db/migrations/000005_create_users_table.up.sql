CREATE TABLE `users` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `email` varchar(255) COMMENT '電子郵件',
    `username` varchar(255) COMMENT '用戶姓名',
    `password` varchar(255) COMMENT '密碼',
    `created_at` timestamp NULL,
    `updated_at` timestamp NULL,
    `deleted_at` timestamp NULL
)