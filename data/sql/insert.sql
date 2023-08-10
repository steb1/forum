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

