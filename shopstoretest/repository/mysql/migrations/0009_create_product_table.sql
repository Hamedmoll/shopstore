-- +migrate Up
CREATE TABLE `products` (
                        `id` INT PRIMARY KEY auto_increment,
                        `name` VARCHAR(255) NOT NULL,
                        `count` INT NOT NULL,
                        `price` INT unsigned NOT NULL,
                        `Description`   VARCHAR(255),
                        `category_id` INT NOT NULL,
                        `created_at`  TIMESTAMP DEFAULT current_timestamp,
                        FOREIGN KEY (`category_id`) REFERENCES `categories`(`id`)
);

-- +migrate Down
DROP TABLE products;
