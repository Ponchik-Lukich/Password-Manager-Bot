ALTER TABLE services
    ADD CONSTRAINT unique_name_user_chat_id UNIQUE (name, user_chat_id);