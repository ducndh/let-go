-- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO snippets (title, content, created, expires) VALUES (
'An old silent pond',
'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
CURRENT_TIMESTAMP,
CURRENT_TIMESTAMP + INTERVAL '365 days'
);
INSERT INTO snippets (title, content, created, expires) VALUES (
'Over the wintry forest',
'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
CURRENT_TIMESTAMP,
CURRENT_TIMESTAMP + INTERVAL '365 days'
);
INSERT INTO snippets (title, content, created, expires) VALUES (
'First autumn morning',
'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
CURRENT_TIMESTAMP,
CURRENT_TIMESTAMP + INTERVAL '7 days'
);