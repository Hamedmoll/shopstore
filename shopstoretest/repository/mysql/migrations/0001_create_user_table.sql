-- +migrate Up
CREATE TABLE `users` (
                        `id` INT PRIMARY KEY auto_increment,
                        `name` VARCHAR(255) NOT NULL,
                        `hashed_password` VARCHAR(255) NOT NULL,
                        `credit` INT DEFAULT 0,
                        `phone_number` VARCHAR(255) NOT NULL UNIQUE,
                        `created_at`  TIMESTAMP DEFAULT current_timestamp
);

-- +migrate Down
DROP TABLE users;
