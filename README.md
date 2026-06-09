# tor-relay-info

Monitors Tor relay(s) by nickname and posts their status to a Discord channel as message embeds, refreshing every 10 minutes (scheduled with [robfig/cron/v3](https://github.com/robfig/cron)).

## What it shows

One embed per matching relay:

- Nickname and OR address(es) in the title
- Running status (green when up, red when down)
- Uptime since last restart
- Observed bandwidth and consensus weight
- Last seen timestamp (local time)
- Consensus flags (e.g. Guard, Exit, Fast, Stable)

Relay data is fetched from the [Onionoo API](https://onionoo.torproject.org) — no authentication required.

## Install

```sh
go install tor-relay-info@latest
```

Or build from source:

```sh
git clone https://github.com/yourname/tor-relay-info
cd tor-relay-info
go build -o tor-relay-info .
```

## Usage

The Discord webhook URL is read from the `DISCORD_WEBHOOK_URL` environment variable (never commit it):

```sh
export DISCORD_WEBHOOK_URL="https://discord.com/api/webhooks/.../..."
tor-relay-info <nickname-or-fingerprint>
```

The argument is passed as a search query to Onionoo, so partial nicknames and fingerprints work.

## Docker

```sh
docker build -t tor-relay-info .
docker run --rm -e DISCORD_WEBHOOK_URL="https://discord.com/api/webhooks/.../..." \
  tor-relay-info <nickname-or-fingerprint>
```

## Dependencies

- [robfig/cron/v3](https://github.com/robfig/cron) — scheduling
- [zerolog](https://github.com/rs/zerolog) — structured logging
