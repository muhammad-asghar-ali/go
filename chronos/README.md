# Chronos

Chronos is a bot application that integrates with Slack and provides various functionalities like age calculation, system monitoring, and more. The bot is built using Go and the Slack API, allowing it to interact with Slack users through custom commands.

## Slack Bot Setup

This guide walks you through the setup process for integrating your Slack bot with your application.

## Prerequisites

Before you begin, ensure you have the following:

1. **Slack Workspace**: You need access to a Slack workspace where you can create a bot.
2. **Slack App**: You must create a Slack App and configure it to work with your bot.

## Steps to Set Up

### 1. Create a Slack App

To create a new Slack app, follow these steps:

1. Go to the [Slack API site](https://api.slack.com/).
2. Click on "Create an App."
3. Choose a workspace where you want to install the app and select "From scratch."
4. Give your app a name and click "Create App."

### 2. Generate Bot Tokens

Your bot needs tokens to authenticate and send messages to Slack.

1. In the Slack app dashboard, navigate to **OAuth & Permissions** under the **Features** section.
2. Scroll down to the **Bot Token Scopes** section and add the necessary scopes, such as:
   - `chat:write`: To send messages on behalf of the bot.
   - `commands`: To create custom commands for the bot.
   - `users:read`: If your bot needs to read user information.
   - mat requried more permissions as well
3. Once you have added the necessary scopes, click on **Install App** to generate the tokens.
4. Copy the **Bot User OAuth Token** and the **App Token** for later use in your app.

### 3. Slack Event Subscriptions

For event-driven bots, enable **Event Subscriptions** and **Socket Mode**:

1. Navigate to **Event Subscriptions** under the **Features** section.
2. Turn on the Event Subscriptions toggle.
3. Enter your request URL (this will be the URL where Slack sends events related to the bot, such as messages, mentions, etc.).
4. Subscribe to the events your bot needs, for example:
   - `message.channels`: To get messages in public channels.
   - `app_mention`: To receive events when your bot is mentioned in a conversation.

### 4. Update Your Bot Configuration

Once you have the tokens and event subscriptions set up, you need to configure your application:

1. Update your bot's configuration with the following fields:

   - `SlackBotToken`: The Bot User OAuth Token generated in step 2.
   - `SlackAppToken`: The App Token generated in step 2 (if needed for interactive components).

   Example configuration:

   ```json
   {
     "SlackBotToken": "xoxb-your-bot-user-oauth-token",
     "SlackAppToken": "xapp-your-app-token"
   }
   ```
