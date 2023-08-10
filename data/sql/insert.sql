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
    ('1', 'yazmin_fisher', 'audie_legros97@yahoo.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.1.jpeg', '2'),
    ('2', 'katlyn_zboncak', 'devan.turcotte82@hotmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.2.jpeg', '2'),
    ('3', 'rosal_da54', 'alayna52@hotmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.3.jpeg', '2'),
    ('4', 'herta31', 'melany.brown8@gmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.4.jpeg', '2'),
    ('5', 'andreane_flatley', 'melyna.beahan7@gmail.com', '$2a$04$pFAHXsdDLPp5banDftsQrOz/XZ0AVveq8b2mEH2lHzEnzvStZBpeq', '/uploads/avatar.5.jpeg', '2');

-- Insert 10 categories
INSERT INTO "category" (id, name, createDate, modifiedDate)
VALUES
    ('1', 'Painting', DATETIME('now'), DATETIME('now')),
    ('2', 'Sculpture', DATETIME('now'), DATETIME('now')),
    ('3', 'Photography', DATETIME('now'), DATETIME('now')),
    ('4', 'Digital Art', DATETIME('now'), DATETIME('now')),
    ('5', 'Drawing', DATETIME('now'), DATETIME('now')),
    ('6', 'Mixed Media', DATETIME('now'), DATETIME('now')),
    ('7', 'Printmaking', DATETIME('now'), DATETIME('now')),
    ('8', 'Ceramics', DATETIME('now'), DATETIME('now')),
    ('9', 'Installation Art', DATETIME('now'), DATETIME('now')),
    ('10', 'Performance Art', DATETIME('now'), DATETIME('now'));

-- Insert 10 sample posts for each category
INSERT INTO "post" (id, title, description, imageURL, authorID, isEdited, createDate, modifiedDate)
VALUES
-- Painting Category
    ('1', 'Sunset Landscape', 'A beautiful sunset landscape painting.', '/uploads/banner.4.jpg', '1', false, DATETIME('now'), DATETIME('now')),
    ('2', 'Abstract Expression', 'An abstract expressionist painting with bold colors.', '/uploads/banner.3.jpg', '2', false, DATETIME('now'), DATETIME('now')),
-- Sculpture Category
    ('11', 'Bronze Sculpture', 'A classical bronze sculpture of a figure.', '/uploads/banner.2.jpg', '3', false, DATETIME('now'), DATETIME('now')),
    ('12', 'Modern Sculpture', 'A contemporary abstract sculpture made from metal and wood.', '/uploads/banner.1.jpg', '4', false, DATETIME('now'), DATETIME('now')),
-- Photography Category
    ('21', 'Cityscape at Night', 'A stunning cityscape photograph captured at night.', '/uploads/banner.4.jpg', '5', false, DATETIME('now'), DATETIME('now')),
    ('22', 'Nature Close-up', 'A macro photograph of a flower in nature.', '/uploads/banner.3.jpg', '1', false, DATETIME('now'), DATETIME('now')),
-- Digital Art Category
    ('31', 'Digital Painting', 'A digital painting created using a graphics tablet.', '/uploads/banner.2.jpg', '1', false, DATETIME('now'), DATETIME('now')),
    ('32', 'Sci-Fi Concept Art', 'A concept art of a futuristic city in a sci-fi world.', '/uploads/banner.1.jpg', '1', false, DATETIME('now'), DATETIME('now')),
-- Drawing Category
    ('41', 'Charcoal Portrait', 'A realistic charcoal portrait of a person.', '/uploads/banner.4.jpg', '2', false, DATETIME('now'), DATETIME('now')),
    ('42', 'Ink Sketch', 'An ink sketch of a cityscape.', '/uploads/banner.3.jpg', '2', false, DATETIME('now'), DATETIME('now')),
-- Mixed Media Category
    ('51', 'Collage Art', 'A collage artwork combining various materials.', '/uploads/banner.2.jpg', '5', false, DATETIME('now'), DATETIME('now')),
    ('52', 'Assemblage', 'An assemblage art piece created from found objects.', '/uploads/banner.1.jpg', '4', false, DATETIME('now'), DATETIME('now')),
-- Printmaking Category
    ('61', 'Linocut Print', 'A linocut print of a nature scene.', '/uploads/banner.4.jpg', '4', false, DATETIME('now'), DATETIME('now')),
    ('62', 'Etching', 'An etching print with intricate details.', '/uploads/banner.3.jpg', '4', false, DATETIME('now'), DATETIME('now')),
-- Ceramics Category
    ('71', 'Porcelain Vase', 'A delicate porcelain vase with intricate patterns.', '/uploads/banner.2.jpg', '3', false, DATETIME('now'), DATETIME('now')),
    ('72', 'Stoneware Sculpture', 'A stoneware sculpture of an animal.', '/uploads/banner.1.jpg', '4', false, DATETIME('now'), DATETIME('now')),
-- Installation Art Category
    ('81', 'Interactive Installation', 'An interactive art installation involving lights and sound.', '/uploads/banner.4.jpg', '5', false, DATETIME('now'), DATETIME('now')),
    ('82', 'Sculptural Installation', 'A large-scale sculptural installation in a public space.', '/uploads/banner.3.jpg', '4', false, DATETIME('now'), DATETIME('now')),
-- Performance Art Category
    ('91', 'Live Art Performance', 'A live art performance involving movement and expression.', '/uploads/banner.2.jpg', '1', false, DATETIME('now'), DATETIME('now')),
    ('92', 'Body Painting Show', 'A body painting performance with intricate designs.', '/uploads/banner.1.jpg', '4', false, DATETIME('now'), DATETIME('now'));

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
    ("0", "beautiful", "4", "1", "", DATETIME('now'), DATETIME('now')),
    ("01", "beautiful", "5", "1", "", DATETIME('now'), DATETIME('now')),
    ("02", "beautiful", "2", "1", "", DATETIME('now'), DATETIME('now')),
    ("03", "beautiful", "5", "1", "01", DATETIME('now'), DATETIME('now')),
    ("04", "beautiful", "5", "11", "", DATETIME('now'), DATETIME('now')),
    ("1", "waouh!", "3", "12", "", DATETIME('now'), DATETIME('now')),
    ("11", "waouh!", "4", "12", "", DATETIME('now'), DATETIME('now')),
    ("12", "waouh!", "5", "21", "", DATETIME('now'), DATETIME('now')),
    ("13", "waouh!", "5", "22", "", DATETIME('now'), DATETIME('now')),
    ("14", "waouh!", "1", "32", "", DATETIME('now'), DATETIME('now')),
    ("2", "nice", "4", "41", "", DATETIME('now'), DATETIME('now')),
    ("21", "nickel", "5", "51", "", DATETIME('now'), DATETIME('now')),
    ("22", "niceuh", "1", "52", "", DATETIME('now'), DATETIME('now')),
    ("3", "respect!", "4", "72", "", DATETIME('now'), DATETIME('now')),
    ("31", "respect!", "5", "72", "3", DATETIME('now'), DATETIME('now')),
    ("4", "beautiful", "4", "72", "3", DATETIME('now'), DATETIME('now')),
    ("5", "beautiful", "1", "82", "", DATETIME('now'), DATETIME('now')),
    ("6", "beautiful", "2", "82", "", DATETIME('now'), DATETIME('now')),
    ("7", "beautiful", "2", "91", "", DATETIME('now'), DATETIME('now')),
    ("8", "beautiful", "3", "91", "", DATETIME('now'), DATETIME('now'));

INSERT INTO view(id,isBookmarked,rate,authorID,postID)
VALUES 
("0",false,"1","2","1"),
("1",false,"2","2","2"),
("2",false,"1","2","1"),
("3",false,"1","2","3"),
("4",false,"1","2","4"),
("5",false,"1","2","5"),
("6",false,"1","2","5"),
("7",false,"1","2","3"),
("8",false,"1","2","1"),
("9",false,"1","2","4"),
("10",false,"1","2","1"),
("11",false,"1","2","3"),
("13",false,"1","2","2");
