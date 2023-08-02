-- SQL schema for user-class-diagram

-- Table for 'user'
CREATE TABLE user (
    id VARCHAR PRIMARY KEY,
    username VARCHAR,
    email VARCHAR,
    password TEXT,
    avatarURL VARCHAR,
    type VARCHAR,
    token TEXT,
    tokenExpirationDate DATE
);

-- Table for 'post'
CREATE TABLE post (
    id VARCHAR PRIMARY KEY,
    title VARCHAR,
    description VARCHAR,
    imageURL VARCHAR,
    authorID VARCHAR,
    isEdited BOOLEAN,
    createDate DATE,
    modifiedDate DATE,
    FOREIGN KEY (authorID) REFERENCES user(id)
);

-- Table for 'report'
CREATE TABLE report (
    id VARCHAR PRIMARY KEY,
    authorID VARCHAR,
    reportedID VARCHAR,
    cause VARCHAR,
    type VARCHAR,
    createDate DATE,
    modifiedDate DATE,
    FOREIGN KEY (authorID) REFERENCES user(id),
    FOREIGN KEY (reportedID) REFERENCES user(id)
);

-- Table for 'response'
CREATE TABLE response (
    id VARCHAR PRIMARY KEY,
    authorID VARCHAR,
    reportID VARCHAR,
    content TEXT,
    createDate DATE,
    modifiedDate DATE,
    FOREIGN KEY (authorID) REFERENCES user(id),
    FOREIGN KEY (reportID) REFERENCES report(id)
);

-- Table for 'view'
CREATE TABLE view (
    id VARCHAR PRIMARY KEY,
    isBookmarked BOOLEAN,
    rate VARCHAR,
    authorID VARCHAR,
    postID VARCHAR,
    FOREIGN KEY (authorID) REFERENCES user(id),
    FOREIGN KEY (postID) REFERENCES post(id)
);

-- Table for 'comment_like'
CREATE TABLE comment_like (
    id VARCHAR PRIMARY KEY,
    authorID VARCHAR,
    commentID VARCHAR,
    FOREIGN KEY (authorID) REFERENCES user(id),
    FOREIGN KEY (commentID) REFERENCES comment(id)
);

-- Table for 'comment'
CREATE TABLE comment (
    id VARCHAR PRIMARY KEY,
    text VARCHAR,
    authorID VARCHAR,
    postID VARCHAR,
    parentID VARCHAR,
    createDate DATE,
    modifiedDate DATE,
    FOREIGN KEY (authorID) REFERENCES user(id),
    FOREIGN KEY (postID) REFERENCES post(id),
    FOREIGN KEY (parentID) REFERENCES comment(id)
);

-- Table for 'post_category'
CREATE TABLE post_category (
    id VARCHAR PRIMARY KEY,
    categoryID VARCHAR,
    postID VARCHAR,
    FOREIGN KEY (categoryID) REFERENCES category(id),
    FOREIGN KEY (postID) REFERENCES post(id)
);

-- Table for 'category'
CREATE TABLE category (
    id VARCHAR PRIMARY KEY,
    name VARCHAR,
    createDate DATE,
    modifiedDate DATE
);

-- Enum for 'RATE'
CREATE TABLE RATE (
    id INTEGER PRIMARY KEY,
    value VARCHAR
);

-- Enum values for 'RATE'
INSERT INTO RATE (id, value) VALUES
    (1, 'NONE'),
    (2, 'LIKE'),
    (3, 'DISLIKE');

-- Enum for 'TYPE'
CREATE TABLE TYPE (
    id INTEGER PRIMARY KEY,
    value VARCHAR
);

-- Enum values for 'TYPE'
INSERT INTO TYPE (id, value) VALUES
    (1, 'COMMENT'),
    (2, 'POST');

-- Enum for 'ROLE'
CREATE TABLE ROLE (
    id INTEGER PRIMARY KEY,
    value VARCHAR
);

-- Enum values for 'ROLE'
INSERT INTO ROLE (id, value) VALUES
    (1, 'ADMIN'),
    (2, 'MODERATOR'),
    (3, 'USER');

-- Enum for 'CAUSE'
CREATE TABLE CAUSE (
    id INTEGER PRIMARY KEY,
    value VARCHAR
);

-- Enum values for 'CAUSE'
INSERT INTO CAUSE (id, value) VALUES
    (1, 'IRRELEVANT'),
    (2, 'OBSCENE'),
    (3, 'ILLEGAL'),
    (4, 'INSULTING');
