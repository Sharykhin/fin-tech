-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `companies` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `symbol` VARCHAR(45) NOT NULL,
  `name` VARCHAR(80) NOT NULL,
  `exchange` VARCHAR(80) NOT NULL,
  `website` VARCHAR(255) NOT NULL,
  `description` VARCHAR(1000) NULL,
  `tags` JSON NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `symbol_UNIQUE` (`symbol` ASC));


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE `companies`;