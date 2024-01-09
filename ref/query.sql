
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(250) NOT NULL,
    password VARCHAR(250) NOT NULL
);


CREATE TABLE user_reset_password(
    id SERIAL PRIMARY KEY,
    code VARCHAR(4) NOT NULL,
    user_id SERIAL REFERENCES users,
    expired_at TIMESTAMP,
    is_active BOOLEAN DEFAULT true 
);


INSERT INTO users (username, password) VALUES (
    'maksim',
    'qwerty123'
);


SELECT * from users;


INSERT INTO user_reset_password(code, user_id, expired_at) VALUES (
    '4478',
    1,
    '09.01.2024'
);



UPDATE users SET password = 'new_password' WHERE id = 1;




UPDATE user_reset_password SET is_active = false WHERE user_id = 1 and code = '4478';






select * from users where username = 'maksim'