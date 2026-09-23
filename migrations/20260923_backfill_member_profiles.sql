INSERT INTO member_profiles (user_id, real_name, student_id, `group`, join_year)
SELECT u.id, '', '', u.`group`, CAST(NULLIF(u.timejoin, '') AS UNSIGNED)
FROM users u
WHERE (COALESCE(u.`group`,'') <> '' OR COALESCE(u.timejoin,'') <> '')
  AND NOT EXISTS (SELECT 1 FROM member_profiles m WHERE m.user_id = u.id);