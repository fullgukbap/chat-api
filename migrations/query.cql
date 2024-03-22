-- CREATE KEYSPACE chat_api with replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
CREATE TABLE user (
    user_id varchar,
    name varchar,
	password varchar,

    PRIMARY KEY (user_id)
);

-- INSERT INTO user (user_id, name, password) VALUES (1, "JunBUm", "1234");
-- SELECT * FROM users WHERE user_id;
-- UPDATE user SET name = "sexy" WHERE id = 1;
-- DELETE * FROM user WHERE id = 1;

-- SELECT * FROM user WHERE id = 1 ORDER BY user_id ASC;
-- SELECT * FROM table_name WHERE column_name1 = value ORDER BY cloumn_name2 ASC;


