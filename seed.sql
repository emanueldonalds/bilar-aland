CREATE DATABASE bilkoll;
CREATE USER `bilkoll`@`%` IDENTIFIED BY 'abc123';
GRANT ALL PRIVILEGES ON `bilkoll`.* TO `bilkoll`@'%';
USE bilkoll;

CREATE TABLE `listing` (
  `url` varchar(255),
  `title` text,
  `normalized_title` text,
  `price` text,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`url`),
  FULLTEXT(normalized_title)
) ENGINE=InnoDB AUTO_INCREMENT=392 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `scrape_event` (
  `date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`date`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4;

INSERT INTO `scrape_event` VALUES('2025-02-25 12:21:37');

INSERT INTO `listing` VALUES 
(
    'https://test.com/some-car',
    'Ford modell T, nybesiktigad',
    'Ford modell T, nybesiktigad',
    '45 000',
    '2025-02-25 10:21:37'
),
(
    'https://test.com/some-other-car',
    'Mazda RX-7 -99',
    'Mazda RX7 99',
    '32 222',
    '2025-02-25 10:21:37'
),
(
    'https://test.com/focus-fint-skick',
    'Ford focus finfint skick',
    'Ford focus finfint skick',
    '32 222',
    '2025-02-25 10:21:37'
)
