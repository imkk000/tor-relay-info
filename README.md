# tor-relay-info

A terminal dashboard that monitors Tor relay(s) by nickname, refreshing every 10 minutes.

## What it shows

For each matching relay:
- OR address(es) and running status
- Uptime since last restart
- Last seen timestamp (local time)
- Consensus flags (e.g. Guard, Exit, Fast, Stable)
- Observed bandwidth and consensus weight

Data is fetched from the [Onionoo API](https://onionoo.torproject.org) — no authentication required.

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

```sh
tor-relay-info <nickname-or-fingerprint>
```

The argument is passed as a search query to Onionoo, so partial nicknames and fingerprints work.

## Example

```
185.220.101.1:9001 [running]
312 hours 4 min 22 sec
Sat, 19 Apr 2026 14:32:00 CEST
Exit, Fast, Guard, HSDir, Running, Stable, V2Dir, Valid
1.23 GB - 9812
```

## Dependencies

- [zerolog](https://github.com/rs/zerolog) — structured logging
