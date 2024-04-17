-- DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS
    `users` (
        id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
        email TEXT NOT NULL,
        password TEXT NOT NULL
    );