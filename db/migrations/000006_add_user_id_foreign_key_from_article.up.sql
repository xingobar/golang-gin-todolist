ALTER TABLE `articles`
ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);