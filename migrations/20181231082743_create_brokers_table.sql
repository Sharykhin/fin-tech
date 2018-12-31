-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `brokers` (
  `user_id` INT UNSIGNED NOT NULL,
  `position` TINYINT NULL,
  `created_at` TIMESTAMP NOT NULL,
  `deleted_at` TIMESTAMP NULL,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `fk_user_id_users_id`
   FOREIGN KEY (`user_id`)
     REFERENCES `fintech`.`users` (`id`)
     ON DELETE RESTRICT
     ON UPDATE RESTRICT
);



-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS `brokers`;
