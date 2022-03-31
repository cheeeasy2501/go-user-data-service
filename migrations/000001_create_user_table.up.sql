CREATE TABLE users(
                      id BIGINT PRIMARY KEY,
                      email VARCHAR(100),
                      password  VARCHAR(255),
                      firstName VARCHAR(50),
                      lastName VARCHAR(50),
                      active tinyint UNSIGNED
)