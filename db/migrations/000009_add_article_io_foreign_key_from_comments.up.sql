ALTER TABLE `comments`
ADD FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`);