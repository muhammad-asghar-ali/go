# Slack File Upload Bot

A simple Slack bot that uploads files to a Slack channel using the Slack API. The bot uses an OAuth token and file paths to upload files to specified channels in Slack.

## Features

- Uploads files to Slack channels.
- Supports multiple channels for file upload.
- Configurable bot token and channels via `.env` file.

## Prerequisites

Before running the bot, make sure you have the following:

1. A Slack workspace where you have permission to install Slack apps.
2. A Slack bot created through the [Slack API](https://api.slack.com/apps) with the appropriate permissions.
   - `files:write`: To upload files.
   - `channels:read`: To read public channels.
   - `groups:read`: To read private groups (if applicable).
3. Go installed on your machine (version 1.18 or higher).
