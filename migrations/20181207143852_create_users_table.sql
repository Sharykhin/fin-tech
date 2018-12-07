-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `users` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(80) NOT NULL,
  `first_name` VARCHAR(45) NULL,
  `last_name` VARCHAR(45) NULL,
  `created_at` TIMESTAMP NOT NULL,
  `deleted_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC));


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE `users`;
