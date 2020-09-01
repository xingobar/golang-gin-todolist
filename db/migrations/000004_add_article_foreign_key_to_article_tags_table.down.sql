ALTER TABLE `article_tags`
DROP FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`);