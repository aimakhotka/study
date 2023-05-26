-- Создание сводной таблицей со всеми фанфиками и информацией по ним
DROP TABLE IF EXISTS Fanfics;

CREATE TABLE Fanfics (
    id  int generated always as identity primary key,
    name    text,
    fandom  text,
    genre   text,
    rating    text,
    status  text,
    "desc"    text,
    link    text
);

DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO public;