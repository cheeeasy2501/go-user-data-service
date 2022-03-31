CREATE TABLE users
(
    id        BIGINT AUTO_INCREMENT PRIMARY KEY,
    email     VARCHAR(100) NOT NULL,
    password  VARCHAR(255),
    firstName VARCHAR(50),
    lastName  VARCHAR(50),
    active    tinyint UNSIGNED,
    CONSTRAINT users_email_uindex
        UNIQUE (email)
)