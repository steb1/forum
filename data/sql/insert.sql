DELETE FROM user;
DELETE FROM post;
DELETE FROM report;
DELETE FROM response;
DELETE FROM view;
DELETE FROM comment_like;
DELETE FROM comment;
DELETE FROM post_category;
DELETE FROM category;

INSERT INTO category(id,name,createDate,modifiedDate) VALUES ("0","Culture",now(),now())
INSERT INTO category(id,name,createDate,modifiedDate) VALUES ("1","Sport",now(),now())
INSERT INTO category(id,name,createDate,modifiedDate) VALUES ("2","Education",now(),now())
INSERT INTO category(id,name,createDate,modifiedDate) VALUES ("3","Movie",now(),now())
INSERT INTO category(id,name,createDate,modifiedDate) VALUES ("4","Game",now(),now())
INSERT INTO category(id,name,createDate,modifiedDate) VALUES ("5","Dance",now(),now())
INSERT INTO category(id,name,createDate,modifiedDate) VALUES ("6","Musique",now(),now())
INSERT INTO category(id,name,createDate,modifiedDate) VALUES ("7","Art",now(),now())
INSERT INTO category(id,name,createDate,modifiedDate) VALUES ("8","Information",now(),now())
INSERT INTO category(id,name,createDate,modifiedDate) VALUES ("9","Actuality",now(),now())

INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES("0","pba","aloou@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.0.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES("1","papgueye","pape@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.1.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES("2","boukha","boukha@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.2.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES("3","fatou","fatou@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.3.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES("4","modou","modou@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.4.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES("5","doudou","doudou@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.5.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES("6","aladji","aladji@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.6.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES("7","bouba","bouba@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.7.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES("8","abdou","abdou@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.8.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES("9","ndeye","ndeye@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.9.jpeg","2")

INSERT INTO post(id,title,description,imageURL,authorID,isEdited,createDate,modifiedDate) VALUES("0","bon","good vibe","upload/avatar.0.jpeg","0","false",now(),now())
INSERT INTO post(id,title,description,imageURL,authorID,isEdited,createDate,modifiedDate) VALUES("1","bon","good vibe","upload/avatar.1.jpeg","1","false",now(),now())
INSERT INTO post(id,title,description,imageURL,authorID,isEdited,createDate,modifiedDate) VALUES("2","bon","good vibe","upload/avatar.0.jpeg","0","false",now(),now())
INSERT INTO post(id,title,description,imageURL,authorID,isEdited,createDate,modifiedDate) VALUES("3","bon","good vibe","upload/avatar.2.jpeg","1","false",now(),now())
INSERT INTO post(id,title,description,imageURL,authorID,isEdited,createDate,modifiedDate) VALUES("4","bon","good vibe","upload/avatar.6.jpeg","4","false",now(),now())
INSERT INTO post(id,title,description,imageURL,authorID,isEdited,createDate,modifiedDate) VALUES("5","bon","good vibe","upload/avatar.7.jpeg","5","false",now(),now())
INSERT INTO post(id,title,description,imageURL,authorID,isEdited,createDate,modifiedDate) VALUES("6","bon","good vibe","upload/avatar.8.jpeg","6","false",now(),now())
INSERT INTO post(id,title,description,imageURL,authorID,isEdited,createDate,modifiedDate) VALUES("7","bon","good vibe","upload/avatar.3.jpeg","3","false",now(),now())
INSERT INTO post(id,title,description,imageURL,authorID,isEdited,createDate,modifiedDate) VALUES("8","bon","good vibe","upload/avatar.4.jpeg","4","false",now(),now())
INSERT INTO post(id,title,description,imageURL,authorID,isEdited,createDate,modifiedDate) VALUES("9","bon","good vibe","upload/avatar.5.jpeg","9","false",now(),now())

INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("0","beautiful","4","8","4",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("01","beautiful","5","8","4",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("02","beautiful","6","8","4",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("03","beautiful","7","8","4",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("04","beautiful","8","8","4",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("1","waouh!","3","4","4",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("11","waouh!","4","4","4",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("12","waouh!","5","4","4",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("13","waouh!","6","4","4",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("14","waouh!","7","4","4",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("2","nice","4","6","6",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("21","nickel","5","6","6",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("22","niceuh","6","6","6",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("3","respect!","4","9","9",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("31","respect!","5","9","9",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("4","beautiful","4","1","1",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("5","beautiful","8","1","1",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("6","beautiful","9","1","1",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("7","beautiful","2","1","1",now(),now())
INSERT INTO comment(id,text,authorID,postID,parentID,createDate,modifiedDate) VALUES("8","beautiful","1","1","1",now(),now())
