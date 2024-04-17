-- DROP TABLE IF EXISTS `stores`;
CREATE TABLE IF NOT EXISTS
    `stores` (
        id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
        title TEXT NOT NULL,
        description TEXT NOT NULL,
        userId BIGINT NOT NULL
    );