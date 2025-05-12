-- Таблица пользователей
CREATE TABLE
    users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        username VARCHAR(255) UNIQUE NOT NULL,
        password_hash VARCHAR(255) NOT NULL
    );

-- Таблица списков (список ссылается на одного пользователя, связь - 1:N)
CREATE TABLE
    notes_lists (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        user_id INT NOT NULL,
        CONSTRAINT fk_notes_lists_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );

-- Таблица заметок ()
CREATE TABLE
    notes (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        content TEXT,
        pinned BOOLEAN DEFAULT FALSE,
        created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        list_id INT NOT NULL,
        CONSTRAINT fk_notes_list FOREIGN KEY (list_id) REFERENCES notes_lists (id) ON DELETE CASCADE
    );