DELETE FROM user;
DELETE FROM post;
DELETE FROM report;
DELETE FROM response;
DELETE FROM view;
DELETE FROM comment_like;
DELETE FROM comment;
DELETE FROM post_category;
DELETE FROM category;

-- Insert 5 categories
INSERT INTO "user" (id, username, email, password, avatarURL, role)
VALUES
    ('1', 'yazmin_fisher', 'a@a', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.1.jpeg', '2'),
    ('2', 'katlyn_zboncak', 'devan.turcotte82@hotmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.2.jpeg', '2'),
    ('3', 'rosal_da54', 'alayna52@hotmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.3.jpeg', '2'),
    ('4', 'herta31', 'melany.brown8@gmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.4.jpeg', '2'),
    ('6', 'rodolfo_simonis58', 'cathy25@hotmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.1.jpeg', '2'),
    ('7', 'boris_boyle49', 'amalia_mann@yahoo.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.1.jpeg', '2'),
    ('8', 'geovanny50', 'chris69@hotmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.5.jpeg', '2'),
    ('9', 'kristian_heaney', 'wilfredo.gorczany@yahoo.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.2.jpeg', '2'),
    ('10', 'lavinia_koch', 'kyle.batz@hotmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.4.jpeg', '2'),
    ('11', 'bo.damore', 'nolan_pagac70@gmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.2.jpeg', '2'),
    ('12', 'garland.keeling', 'moises50@gmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.5.jpeg', '2'),
    ('13', 'lura_walker', 'reba.stiedemann@gmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.3.jpeg', '2'),
    ('14', 'anabel58', 'bobbie_haley@gmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.3.jpeg', '2'),
    ('15', 'filomena.jacobi97', 'otilia.oconner64@gmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.4.jpeg', '2'),
    ('16', 'syble.lueilwitz16', 'destin65@yahoo.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.2.jpeg', '2'),
    ('17', 'rosalyn17', 'anabelle57@gmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.1.jpeg', '2'),
    ('18', 'adah.hammes', 'seth_schmeler@gmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.2.jpeg', '2'),
    ('19', 'lempi.ward', 'treva.rogahn@hotmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.4.jpeg', '2'),
    ('20', 'ferne.hartmann', 'michele_kling87@yahoo.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.4.jpeg', '2'),
    ('21', 'einar_schulist', 'danyka.gaylord59@gmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.3.jpeg', '2'),
    ('22', 'anabelle.von', 'bertrand_mclaughlin@gmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.4.jpeg', '2'),
    ('23', 'arnold.padberg', 'cole_sporer@yahoo.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.1.jpeg', '2'),
    ('24', 'edmond20', 'admin@ok.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.5.jpeg', '0'),
    ('5', 'andreane_flatley', 'melyna.beahan7@gmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.5.jpeg', '2');

-- Insert 10 categories
INSERT INTO "category" (id, name, createDate, modifiedDate)
VALUES
    ('1', 'Painting',  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ('2', 'Sculpture',  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ('3', 'Photography',  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ('4', 'Digital Art',  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ('5', 'Drawing',  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ('6', 'Mixed Media',  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ('7', 'Printmaking',  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ('8', 'Ceramics',  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ('9', 'Installation Art',  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ('10', 'Performance Art',  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now')));

-- Insert 10 sample posts for each category
INSERT INTO "post" (id, title, description, imageURL, authorID, isEdited, createDate, modifiedDate)
VALUES
-- Painting Category
    ('1', 'Sunset Landscape', 'A beautiful sunset landscape painting.', '/uploads/banner.4.jpg', '1', false,  "2023-06-10 19:59:44",  "2023-06-10 19:59:44"),
    ('2', 'Abstract Expression', 'An abstract expressionist painting with bold colors.', '/uploads/banner.3.jpg', '2', false,  "2023-04-27 03:39:42",  "2023-04-27 03:39:42"),
-- Sculpture Category
    ('11', 'Bronze Sculpture', 'A classical bronze sculpture of a figure.', '/uploads/banner.2.jpg', '3', false,  "2023-01-26 21:14:26",  "2023-01-26 21:14:26"),
    ('12', 'Modern Sculpture', 'A contemporary abstract sculpture made from metal and wood.', '/uploads/banner.1.jpg', '4', false,  "2023-08-01 04:55:25",  "2023-08-01 04:55:25"),
-- Photography Category
    ('21', 'Cityscape at Night', 'A stunning cityscape photograph captured at night.', '/uploads/banner.4.jpg', '5', false,  "2023-04-23 04:40:53",  "2023-04-23 04:40:53"),
    ('22', 'Nature Close-up', 'A macro photograph of a flower in nature.', '/uploads/banner.3.jpg', '1', false,  "2023-02-16 19:52:49",  "2023-02-16 19:52:49"),
-- Digital Art Category
    ('31', 'Digital Painting', 'A digital painting created using a graphics tablet.', '/uploads/banner.2.jpg', '1', false,  "2023-06-04 15:24:12",  "2023-06-04 15:24:12"),
    ('32', 'Sci-Fi Concept Art', 'A concept art of a futuristic city in a sci-fi world.', '/uploads/banner.1.jpg', '1', false,  "2023-08-02 20:18:47",  "2023-08-02 20:18:47"),
-- Drawing Category
    ('41', 'Charcoal Portrait', 'A realistic charcoal portrait of a person.', '/uploads/banner.4.jpg', '2', false,  "2023-04-17 15:47:47",  "2023-04-17 15:47:47"),
    ('42', 'Ink Sketch', 'An ink sketch of a cityscape.', '/uploads/banner.3.jpg', '2', false,  "2023-06-03 17:01:39",  "2023-06-03 17:01:39"),
-- Mixed Media Category
    ('51', 'Collage Art', 'A collage artwork combining various materials.', '/uploads/banner.2.jpg', '5', false,  "2023-06-04 05:23:54",  "2023-06-04 05:23:54"),
    ('52', 'Assemblage', 'An assemblage art piece created from found objects.', '/uploads/banner.1.jpg', '4', false,  "2023-01-24 22:35:27",  "2023-01-24 22:35:27"),
-- Printmaking Category
    ('61', 'Linocut Print', 'A linocut print of a nature scene.', '/uploads/banner.4.jpg', '4', false,  "2023-03-28 03:03:30",  "2023-03-28 03:03:30"),
    ('62', 'Etching', 'An etching print with intricate details.', '/uploads/banner.3.jpg', '4', false,  "2022-09-20 09:12:45",  "2022-09-20 09:12:45"),
-- Ceramics Category
    ('71', 'Porcelain Vase', 'A delicate porcelain vase with intricate patterns.', '/uploads/banner.2.jpg', '3', false,  "2022-12-01 15:38:42",  "2022-12-01 15:38:42"),
    ('72', 'Stoneware Sculpture', 'A stoneware sculpture of an animal.', '/uploads/banner.1.jpg', '4', false,  "2023-08-06 09:34:15",  "2023-08-06 09:34:15"),
-- Installation Art Category
    ('81', 'Interactive Installation', 'An interactive art installation involving lights and sound.', '/uploads/banner.4.jpg', '5', false,  "2023-01-08 02:16:15",  "2023-01-08 02:16:15"),
    ('82', 'Sculptural Installation', 'A large-scale sculptural installation in a public space.', '/uploads/banner.3.jpg', '4', false,  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
-- Performance Art Category
    ('91', 'Live Art Performance', 'A live art performance involving movement and expression.', '/uploads/banner.2.jpg', '1', false,  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ('92', 'Body Painting Show', 'A body painting performance with intricate designs.', '/uploads/banner.1.jpg', '4', false,  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now')));

INSERT INTO "post_category" (id, categoryID, postID)
VALUES
    ('11', '1', '1'),
    ('12', '1', '2'),
    ('211', '2', '11'),
    ('212', '2', '12'),
    ('321', '3', '21'),
    ('322', '3', '22'),
    ('431', '4', '31'),
    ('432', '4', '32'),
    ('541', '5', '41'),
    ('542', '5', '42'),
    ('651', '6', '51'),
    ('652', '6', '52'),
    ('761', '7', '61'),
    ('762', '7', '62'),
    ('871', '8', '71'),
    ('872', '8', '72'),
    ('981', '9', '81'),
    ('982', '9', '82'),
    ('1091', '10', '91'),
    ('1092', '10', '92');

INSERT INTO comment(id, text, authorID, postID, parentID, createDate, modifiedDate)
VALUES
    ("0", "beautiful", "4", "1", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("01", "beautiful", "5", "1", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("02", "beautiful", "2", "1", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("03", "beautiful", "5", "1", "01",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("04", "beautiful", "5", "11", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("1", "waouh!", "3", "12", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("11", "waouh!", "4", "12", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("12", "waouh!", "5", "21", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("13", "waouh!", "5", "22", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("14", "waouh!", "1", "32", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("2", "nice", "4", "41", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("21", "nickel", "5", "51", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("22", "niceuh", "1", "52", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("3", "respect!", "4", "72", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("31", "respect!", "5", "72", "3",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("4", "beautiful", "4", "72", "3",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("5", "beautiful", "1", "82", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("6", "beautiful", "2", "82", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("7", "beautiful", "2", "91", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now'))),
    ("8", "beautiful", "3", "91", "",  strftime('%Y-%m-%d %H:%M:%S', datetime('now')),  strftime('%Y-%m-%d %H:%M:%S', datetime('now')));

INSERT INTO view(id,isBookmarked,rate,authorID,postID)
VALUES 
    ("0",false,"1","2","1"),
    ("1",false,"2","2","2"),
    ("3",false,"1","1","11"),
    ("4",false,"1","2","41"),
    ("5",true,"1","1","42"),
    ("6",true,"1","1","72"),
    ("7",false,"1","1","21");