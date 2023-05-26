-- Создание таблицы для каждого фанфика из главной таблицы
DO $$
    DECLARE t RECORD;
    BEGIN
        FOR t IN (SELECT name FROM Fanfics)
            LOOP
                RAISE NOTICE 'Creating table %...', t.name;
                EXECUTE format('CREATE TABLE %I '
                    '(
                       f_id  int references fanfics(id),
                       ch_num  int,
                       ch_name text,
                       content text
                    )', t);
            END LOOP;
    END
$$ LANGUAGE plpgsql;

ANALYSE VERBOSE;