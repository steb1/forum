SELECT u.id AS user_id,
       u.username AS user_username,
       COUNT(v.id) AS number_of_likes
FROM "user" u
LEFT JOIN "post" p ON u.id = p.authorID
LEFT JOIN "view" v ON p.id = v.postID
GROUP BY u.id, u.username
ORDER BY (number_of_likes) DESC;

SELECT u.id AS user_id,
       u.username AS user_username,
       COUNT(DISTINCT c.id) + COUNT(DISTINCT v.id) AS number_reaction
FROM "user" u
LEFT JOIN "post" p ON u.id = p.authorID
LEFT JOIN "comment" c ON p.id = c.postID
LEFT JOIN "view" v ON p.id = v.postID
GROUP BY u.id, u.username
ORDER BY number_reaction DESC;

SELECT
    c.postID AS post_id,
    GROUP_CONCAT(DISTINCT u1.avatarURL) AS avatar1,
    GROUP_CONCAT(DISTINCT u2.avatarURL) AS avatar2,
    GROUP_CONCAT(DISTINCT u3.avatarURL) AS avatar3
FROM "comment" c
LEFT JOIN "user" u1 ON c.authorID = u1.id
LEFT JOIN "user" u2 ON c.authorID = u2.id
LEFT JOIN "user" u3 ON c.authorID = u3.id
WHERE c.id IN (
    SELECT id
    FROM "comment"
    WHERE postID = c.postID
    ORDER BY createDate DESC
    LIMIT 3
)
GROUP BY c.postID;

SELECT
    p.id AS post_id,
    p.title AS post_title,
    p.imageURL AS post_image,
    COALESCE(commentators.avatar1, '') AS top_commentator_avatar1,
    COALESCE(commentators.avatar2, '') AS top_commentator_avatar2,
    COALESCE(commentators.avatar3, '') AS top_commentator_avatar3
FROM "post" p
LEFT JOIN (
    SELECT
        c.postID AS post_id,
        GROUP_CONCAT(DISTINCT u1.avatarURL) AS avatar1,
        GROUP_CONCAT(DISTINCT u2.avatarURL) AS avatar2,
        GROUP_CONCAT(DISTINCT u3.avatarURL) AS avatar3
    FROM "comment" c
    LEFT JOIN "user" u1 ON c.authorID = u1.id
    LEFT JOIN "user" u2 ON c.authorID = u2.id
    LEFT JOIN "user" u3 ON c.authorID = u3.id
    WHERE c.id IN (
        SELECT id
        FROM "comment"
        WHERE postID = c.postID
        ORDER BY createDate DESC
        LIMIT 3
    )
    GROUP BY c.postID
) commentators ON p.id = commentators.post_id;

SELECT u.id AS user_id,
       u.username AS user_username,
       p.title AS title,
       c.text AS comment,
       v.rate AS rate
FROM "user" u
LEFT JOIN "post" p ON u.id = p.authorID
LEFT JOIN "comment" c ON p.id = c.postID
LEFT JOIN "view" v ON p.id = v.postID;

-- SELECT u.id AS user_id,
--        u.username AS user_username,
--        COUNT(c.id) AS number_of_comments
-- FROM "user" u
-- LEFT JOIN "post" p ON u.id = p.authorID
-- LEFT JOIN "comment" c ON p.id = c.postID
-- GROUP BY u.id
-- ORDER BY (number_of_comments) DESC
-- LIMIT 3;
-- SELECT u.id AS user_id,
--        u.username AS user_username,
--        COUNT(c.id) AS number_of_comments
-- FROM "user" u
-- LEFT JOIN "post" p ON u.id = p.authorID
-- LEFT JOIN "comment" c ON p.id = c.postID
-- GROUP BY u.id
-- ORDER BY (number_of_comments) DESC
-- LIMIT 3;