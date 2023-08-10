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

INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES ("0","pba","aloou@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.0.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES ("1","papgueye","pape@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.1.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES ("2","boukha","boukha@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.2.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES ("3","fatou","fatou@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.3.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES ("4","modou","modou@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.4.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES ("5","doudou","doudou@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.5.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES ("6","aladji","aladji@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.6.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES ("7","bouba","bouba@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.7.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES ("8","abdou","abdou@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.8.jpeg","2")
INSERT INTO user(id,username,email,password, avatarURL,type,token,tokenExpirationDate) VALUES ("9","ndeye","ndeye@gmail.com","$2a$04$FiWWNzKyobHYajGiKEoyquXSNsWiIYX.zfUTtx.VY6HhFhNBXbgGG","upload/avatar.9.jpeg","2")
