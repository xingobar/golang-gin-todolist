CREATE TABLE `article_tags` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `article_id` BIGINT NOT NULL COMMENT '文章編號',
    `tag_id` BIGINT NOT NULL COMMENT '標籤編號',
    FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`)
)