-- up
create table user_table
(
    id           int,
    quest_number int,
    last_message int unique,
    right_answer int,
    finished     boolean not null default false,
    quests       TEXT,
    answers      TEXT,

    primary key (id)
);

-- down
drop table user_table;