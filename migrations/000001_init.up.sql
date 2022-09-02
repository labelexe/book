-- Категории спорта
CREATE TABLE categories
(
    id      serial                                          primary key,
    name_ru varchar(255)                                    not null,
    name_en varchar(255)                                    not null,
    api_id  int                                             not null,
    api_src varchar(5)                                      not null
);

CREATE INDEX api_categories ON categories(api_id,api_src);

-- Страны
CREATE TABLE countries
(
    id            serial                                          primary key,
    name_ru       varchar(255)                                    not null,
    name_en       varchar(255)                                    not null
);

-- Команды
CREATE TABLE teams
(
    id            serial                                                  primary key,
    category_id   int constraint fk_categories references categories (id) on delete cascade not null,
    name_ru       varchar(255)                                            not null,
    name_en       varchar(255)                                            not null,
    image         bytea                                                       
);

-- Игры
CREATE TABLE games
(
    id                 serial                                            primary key,
    home_team_id       int constraint fk_home_team references teams (id) on delete cascade not null,
    away_team_id       int constraint fk_away_team references teams (id) on delete cascade not null,
    name_ru            varchar(255)                                      not null,
    name_en            varchar(255)                                      not null,
    start_date         timestamp                                         not null,
    current_event_time time                                              not null,
    time_events        time[]                                            not null                                              
);

-- Линия
CREATE TABLE lines
(
    id            serial                                                  primary key,
    name_ru       varchar(255)                                            not null,
    name_en       varchar(255)                                            not null,
    category_id   int constraint fk_categories references categories (id) on delete cascade not null,
    country       varchar(100)                                            on delete cascade not null,
    tourney       varchar(100)                                            not null,
    type          varchar(10) CONSTRAINT one_of_type CHECK (type = 'live' OR type = 'prematch'),
    game_id       int constraint fk_games references games (id)           on delete cascade not null,
    api_id        int                                                     not null,
    api_src       varchar(5)                                              not null
);

CREATE INDEX api_lines ON lines(api_id,api_src);

-- События
CREATE TABLE events
(
    id            serial                                          primary key,
    name_ru       varchar(255)                                    not null,
    name_en       varchar(255)                                    not null
);

-- События и линия
CREATE TABLE events_lines
(
    id            serial                                          primary key,
    line_id       int constraint fk_lines references lines (id)   on delete cascade not null,
    event_id      int constraint fk_events references events (id) on delete cascade not null
);

-- Исходы
CREATE TABLE bets
(
    id            serial                                          primary key,
    name_ru       varchar(255)                                    not null,
    name_en       varchar(255)                                    not null,
    event_id      int constraint fk_events references events (id) on delete cascade not null,
    coefficient   int                                             not null
);