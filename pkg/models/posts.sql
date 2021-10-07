CREATE TABLE posts(
    id TEXT PRIMARY KEY not null,
    content TEXT,
    createdAt TIMESTAMP,
    title TEXT,
    updatedAt   TIMESTAMP,
    authorId    TEXT,
    Foreign Key(authorId) references users(id)
    );

    CREATE TABLE comments(
        id TEXT PRIMARY KEY not null,
        content TEXT,
        createdAt TEXT,
        authorId    TEXT,
        postId      TEXT,
        Foreign Key(authorId) references users(id),
        Foreign Key(postId) references posts(id)
        );


// Change the value of a column of a table
        ALTER TABLE posts
        ADD PRIMARY KEY (id);