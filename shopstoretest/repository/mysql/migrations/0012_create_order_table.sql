-- +migrate Up
CREATE TABLE `order_list` (
                                `id` INT PRIMARY KEY auto_increment,
                                `product_id` INT NOT NULL,
                                `user_id` INT NOT NULL,
                                `price` INT unsigned NOT NULL,
                                `count` INT unsigned NOT NULL,
                                `total_price` INT unsigned NOT NULL,
                                `created_at`  TIMESTAMP DEFAULT current_timestamp
);
-- +migrate Down
DROP TABLE order_list;