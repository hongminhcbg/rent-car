DROP TABLE IF EXISTS `tenants`;
CREATE TABLE `tenants` (
    id BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name_agency VARCHAR(64) NOT NULL,
    email VARCHAR(64),
    phone VARCHAR(16) UNIQUE NOT NULL,
    location VARCHAR(128),
    bank_information VARCHAR(255)
);