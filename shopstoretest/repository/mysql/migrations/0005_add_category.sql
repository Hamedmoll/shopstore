-- +migrate Up
CREATE TABLE categories (
                        `id` INT PRIMARY KEY auto_increment,
                        `name` VARCHAR(255) NOT NULL,
                        `created_at`  TIMESTAMP DEFAULT current_timestamp
);

-- +migrate Down
DROP TABLE categories;