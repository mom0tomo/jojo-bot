CREATE DATABASE jojobot_staging;

DROP TABLE calendars;

CREATE TABLE calendars (
  id integer,
  date varchar(10),
  text varchar(500),
  created_at timestamp,
  updated_at timestamp
);

SELECT * FROM calendars;

COPY calendars FROM '/Users/mom0tomo/dev/go/src/github.com/mom0tomo/jojo-bot/seeds.csv' WITH CSV;

SELECT * FROM calendars;