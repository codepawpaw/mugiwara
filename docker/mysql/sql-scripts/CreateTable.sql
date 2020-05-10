CREATE TABLE users (
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    age varchar(10) NOT NULL,
    sex varchar(20) NOT NULL
);

CREATE TABLE accounts (
    id int PRIMARY KEY AUTO_INCREMENT,
    username varchar(30) NOT NULL UNIQUE,
    password varchar(30) NOT NULL,
    user_id int NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE incidents (
    id int PRIMARY KEY AUTO_INCREMENT,
    cityName varchar(500),
    province varchar(200) NOT NULL,
    nation varchar(100) NOT NULL,
    description text NOT NULL,
    date datetime NOT NULL,
    lat varchar(300) NOT NULL,
    lang varchar(300) NOT NULL,
    user_id int NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
)

INSERT INTO users(name, age, sex) values('user1', '21', 'male');
INSERT INTO accounts(username, password, user_id) values('username123', 'pidel123', 1);

-- INSERT INTO organizers(name, description) values('Organizer 1', 'Description of organizer 1');
-- INSERT INTO organizers(name, description) values('Organizer 2', 'Description of organizer 2');
-- INSERT INTO organizers(name, description) values('Organizer 3', 'Description of organizer 3');

-- INSERT INTO users(name, organizer_id) values('User 1', 1);
-- INSERT INTO users(name, organizer_id) values('User 2', 2);
-- INSERT INTO users(name, organizer_id) values('User 3', 3);

-- INSERT INTO rundowns(title, subtitle, show_time, end_time, organizer_id) values('Rundowns 1', 'Subtitle', '2019-04-04 12:30:00', '2019-04-04 14:00:00', 1);
-- INSERT INTO rundowns(title, subtitle, show_time, end_time, organizer_id) values('Rundowns 2', 'Subtitle', '2019-04-04 13:30:00', '2019-04-04 14:00:00', 1);
-- INSERT INTO rundowns(title, subtitle, show_time, end_time, organizer_id) values('Rundowns 3', 'Subtitle', '2019-04-04 12:30:00', '2019-04-04 18:00:00', 1);
-- INSERT INTO rundowns(title, subtitle, show_time, end_time, organizer_id) values('Rundowns 4', 'Subtitle', '2019-04-04 15:30:00', '2019-04-04 18:00:00', 1);

-- INSERT INTO rundown_items(title, subtitle, text, rundown_id) values('Rundown item 1', 'subs', 'text', 1);
-- INSERT INTO rundown_items(title, subtitle, text, rundown_id) values('Rundown item 2', 'subs', 'text', 1);
-- INSERT INTO rundown_items(title, subtitle, text, rundown_id) values('Rundown item 3', 'subs', 'text', 1);
-- INSERT INTO rundown_items(title, subtitle, text, rundown_id) values('Rundown item 4', 'subs', 'text', 1);