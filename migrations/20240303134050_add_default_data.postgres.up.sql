-- Setup admin user
insert into users (first_name, last_name, email, password, access_level, created_at, updated_at)
values ('admin', 'admin', 'admin@admin.com', '$2a$12$0ED7LWxMUU0G6l9l17iHy.bmzPQCyDxjkyxWpSLomPz3fUWl8IVc.', 1, '2024-03-03 12:30:37.390053', '2024-03-03 12:30:37.390053');

-- Setup rooms data
insert into rooms (room_name, created_at, updated_at)
values ('General''s Quarters', current_timestamp, current_timestamp);

insert into rooms (room_name, created_at, updated_at)
values ('Major''s Suite', current_timestamp, current_timestamp);

-- Set up initial restrictions
insert into restrictions (restriction_name, created_at, updated_at)
values ('Reservation', current_timestamp, current_timestamp);

insert into restrictions (restriction_name, created_at, updated_at)
values ('Restricted', current_timestamp, current_timestamp);