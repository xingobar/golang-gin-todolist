ALTER TABLE `article_tags`
ADD FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`);