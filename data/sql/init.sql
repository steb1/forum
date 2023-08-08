-- SQL schema for user-class-diagram

-- Table for 'user'
CREATE TABLE user (
    id VARCHAR PRIMARY KEY,
    username VARCHAR,
    email VARCHAR,
    password TEXT,
    avatarURL VARCHAR,
    role VARCHAR,
    token TEXT,
    tokenExpirationDate DATE
);

CREATE TABLE session (
    sessionID  VARCHAR PRIMARY KEY,
    ExpirationDate Date
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
