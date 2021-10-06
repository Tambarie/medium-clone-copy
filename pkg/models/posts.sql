CREATE TABLE posts(
    id TEXT not null,
    content TEXT,
    createdAt TIMESTAMP,
    title TEXT,
    updatedAt   TIMESTAMP,
    authorId    TEXT,
    Foreign Key(authorId) references users(id)
    );