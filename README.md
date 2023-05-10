# Password Manager Bot

A simple bot for Telegram that helps manage passwords efficiently.

Link to the bot: https://t.me/password_managment_bot

It is deployed on Vercel as serverless function so the code was not parallelized because Vercel does it itself.

## Deploying the bot on Vercel

[Vercel](https://vercel.com) is a platform for deploying serverless functions and static sites. To deploy this bot on Vercel, follow these steps:

1. Sign up for a Vercel.

2. Install the Vercel CLI by running:

   ```
   npm i -g vercel
   ```

3. Clone this repository:

   ```
   git clone https://github.com/yourusername/password-manager-bot.git
   ```
   Push it to your GitHub add it to Vercel projects.
4. Set up database:

   Vercel supports PostgreSQL databases. After adding project to Vercel, it is possible to set up a PostgreSQL database
   in `storage` tab.
   After that, database url would be available in `storage` tab. It should look like this:
   
   ```
   postgres://username:password@host:port/verceldb
   ```
   
   To create database for this bot, run the following query in Vercel console:

   ```sql
   CREATE TABLE IF NOT EXISTS users
   (
       chat_id    INT         NOT NULL,
       state      VARCHAR(20) NOT NULL,
       message_id integer NOT NULL DEFAULT 0,
       PRIMARY KEY (chat_id),
       UNIQUE (chat_id)
   );

   CREATE TABLE IF NOT EXISTS services
   (
       id           SERIAL       NOT NULL,
       name         VARCHAR(255) NOT NULL,
       login        VARCHAR(255) NOT NULL,
       password     VARCHAR(255) NOT NULL,
       user_chat_id INT          NOT NULL,
       PRIMARY KEY (id),
       FOREIGN KEY (user_chat_id) REFERENCES users (chat_id)
   );

   ALTER TABLE services
    ADD CONSTRAINT unique_name_user_chat_id UNIQUE (name, user_chat_id);
   ```


5. Add the necessary environment variables for your project:

   ```
   vercel secrets add telegram_bot_token "your-telegram-bot-api-token"
   vercel secrets add telegram_webhook_url "your-webhook-url"
   vercel secrets add database_url "your-database-url"
   vercel secrets add encrypt_key "your-encription-key"
   ```
   Encryption key should be 32 characters long. It is used to encrypt and decrypt login and passwords to store 
   them encrypted to try to avoid password leaks in case of database theft (it's not perfect but better than nothing).


6. Push your changes to GitHub to deploy automatically or run the following command to deploy manually:

   ```
   vercel --prod
   ```

   Bot should now be running on Vercel!

## Local start

1. Clone the repository to your local machine.

2. In the project root directory, create a `.env` file and add the following environment variables:

   ```
   TELEGRAM_BOT_TOKEN=<your-bot-token>
   ENCRYPT_KEY=<your-encription-key>
   ```

3. Run `docker-compose up --build` to run it in docker.

## Usage

1. Start a chat with the bot on Telegram.

2. Use the inline keyboard to navigate through the bot's commands (Get, Set, Del).

3. Follow the bot's instructions to store, retrieve, or delete credentials.