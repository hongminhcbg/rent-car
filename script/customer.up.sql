DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers` (
    id BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(32) NOT NULL,
    email VARCHAR(64),
    phone VARCHAR(16) UNIQUE NOT NULL,
    location VARCHAR(128)
);