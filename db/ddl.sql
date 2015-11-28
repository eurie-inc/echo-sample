DROP TABLE IF EXISTS member;
CREATE TABLE IF NOT EXISTS member (
  number       INT PRIMARY KEY,
  name         VARCHAR(255) NOT NULL,
  date_created BIGINT       NOT NULL
);

