<div align="center">

# hpp

**Search Japanese restaurants from your terminal.**

A small CLI for the HotPepper Gourmet API that covers every endpoint,<br>
with JSON output for scripts and agents, and tables for humans.

[![Test](https://github.com/jackchuka/hpp/actions/workflows/test.yml/badge.svg)](https://github.com/jackchuka/hpp/actions/workflows/test.yml)
[![Release](https://img.shields.io/github/v/release/jackchuka/hpp?sort=semver)](https://github.com/jackchuka/hpp/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

<img src="assets/demo.gif" alt="hpp looking up an area and genre code, then searching ramen and izakaya spots in Shibuya" width="800">

[Installation](#installation) · [Quick Start](#quick-start) · [Commands](#commands) · [Search Flags](#search-flags) · [Agent Skill](#agent-skill)

</div>

---

The [HotPepper Gourmet API](https://webservice.recruit.co.jp/doc/hotpepper/reference.html) has rich filters (private rooms, all-you-can-drink, party capacity, English menus), but using it means hand-building query strings and looking up area and genre codes. **hpp** turns each endpoint into a command with typed flags, so you can find a place for ten people near the station in one line, or pipe the JSON into `jq` or an AI agent.

## Features

- **Full API coverage**: all 12 HotPepper Gourmet endpoints, from restaurant search to every master list
- **60 search filters**: keyword, location radius, area, genre, budget, and 32 yes/no filters such as `--private-room`, `--free-drink`, `--english`, and `--non-smoking`
- **Code lookup**: browse area, genre, budget, credit card, and special codes to feed back into `search`
- **Two output formats**: indented JSON by default for scripting, and `--format table` for reading
- **Agent skill**: ships a Claude Code skill that turns requests like "izakaya near 浜松町 for 6 people" into `hpp` searches
- **Single binary**: Linux, macOS, and Windows on amd64 and arm64

## Installation

### Homebrew

```bash
brew install jackchuka/tap/hpp
```

### Go

```bash
go install github.com/jackchuka/hpp@latest
```

### Binary

Download a prebuilt binary from the [releases page](https://github.com/jackchuka/hpp/releases).

## Quick Start

1. **Get an API key** for free from [Recruit Web Service](https://webservice.recruit.co.jp/register/) and export it:

   ```bash
   export HOTPEPPER_API_KEY=your_key_here
   ```

2. **Look up codes** for the area and genre you want:

   ```bash
   hpp area middle --keyword 渋谷 --format table   # Y030
   hpp genre --format table                       # G001 = 居酒屋
   ```

3. **Search**:

   ```bash
   hpp search --middle-area Y030 --genre G001 --private-room --free-drink --format table
   ```

## Usage

```bash
# Keyword search within a large area (Z011 = Tokyo)
hpp search --keyword ramen --area Z011

# Within 1 km of a point
hpp search --lat 35.6812 --lng 139.7671 --range 3

# Filter by features
hpp search --keyword izakaya --wifi --private-room --english --non-smoking

# Groups: banquet capacity of at least 10, with all-you-can-drink
hpp search --keyword 新宿 --party-capacity 10 --free-drink

# Pagination
hpp search --keyword ramen --count 20 --start 21

# Pipe JSON into jq
hpp search --keyword sushi --count 5 | jq -r '.results.shop[] | .name'

# Find a shop by name or phone number
hpp shop --keyword 九月堂
hpp shop --tel 0312345678

# Specials in a category (SPG4 = 個室・貸切・設備で探す)
hpp special category --format table
hpp special list --category SPG4 --format table
```

## Commands

| Command                   | API endpoint       | Flags                                                        |
| ------------------------- | ------------------ | ------------------------------------------------------------ |
| `hpp search`              | Gourmet Search     | See [Search Flags](#search-flags)                            |
| `hpp shop`                | Shop Name Search   | `--keyword`, `--tel`, `--start`, `--count` (max 30)          |
| `hpp genre`               | Genre Master       | `--code`, `--keyword`                                        |
| `hpp budget`              | Budget Master      |                                                              |
| `hpp area large`          | Large Area Master  | `--code`, `--keyword`                                        |
| `hpp area middle`         | Middle Area Master | `--code`, `--large-area`, `--keyword`, `--start`, `--count`  |
| `hpp area small`          | Small Area Master  | `--code`, `--middle-area`, `--keyword`, `--start`, `--count` |
| `hpp service-area large`  | Large Service Area |                                                              |
| `hpp service-area list`   | Service Area       |                                                              |
| `hpp creditcard`          | Credit Card Master |                                                              |
| `hpp special list`        | Special Master     | `--code`, `--category`                                       |
| `hpp special category`    | Special Category   | `--code`                                                     |
| `hpp version`             |                    |                                                              |

Flags that take codes accept a comma-separated list, for example `--code G001,G002`.

### Global flags

| Flag       | Description                      | Default |
| ---------- | -------------------------------- | ------- |
| `--format` | Output format: `json` or `table` | `json`  |

With `--format table`, columns stay aligned for Japanese (full-width) text, and `search` and `shop` also print the total match count to stderr.

Errors go to stderr as a single `Error: ...` line with a non-zero exit code. Your API key is never included in error messages.

## Search Flags

Every flag is optional. Run `hpp search --help` for the full list.

| Group    | Flags                                                                                                                                                                                         |
| -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Text     | `--keyword`, `--name` (partial match), `--name-kana`, `--name-any`, `--tel` (digits only), `--address` (partial match)                                                                        |
| Location | `--lat`, `--lng`, `--range` (1=300m, 2=500m, 3=1km, 4=2km, 5=3km), `--datum` (`world` or `tokyo`)                                                                                            |
| Area     | `--large-service-area`, `--service-area`, `--area` (large area), `--middle-area`, `--small-area`                                                                                             |
| Category | `--genre`, `--budget`, `--credit-card`, `--special` / `--special-or`, `--special-category` / `--special-category-or` (AND / OR)                                                              |
| Capacity | `--party-capacity` (minimum banquet capacity)                                                                                                                                                 |
| Features | `--wifi`, `--lunch`, `--midnight`, `--midnight-meal`, `--english`, `--private-room`, `--non-smoking`, `--card`, `--parking`, `--pet`, `--child`, `--barrier-free`, `--free-drink`, `--free-food`, `--course` |
| Seating  | `--horigotatsu`, `--tatami`, `--open-air`, `--night-view`, `--charter`                                                                                                                       |
| Drinks   | `--cocktail`, `--shochu`, `--sake`, `--wine`, `--sommelier`                                                                                                                                   |
| Other    | `--wedding`, `--show`, `--equipment`, `--karaoke`, `--band`, `--tv`, `--ktai`, `--ktai-coupon` (0 or 1)                                                                                       |
| Output   | `--type` (`lite`, `credit_card`, `special`), `--order` (1=name, 2=genre, 3=area, 4=recommended), `--start`, `--count` (max 100)                                                             |

## Agent Skill

This repository includes a [Claude Code](https://docs.anthropic.com/en/docs/claude-code) skill, [`restaurant-search`](.claude/skills/restaurant-search/SKILL.md), that turns plain requests into `hpp` searches. Install it with:

```bash
npx skills add jackchuka/hpp
```

Then use `/restaurant-search`, or just ask:

```
> find an izakaya near 浜松町 for 6 people with private rooms
> ramen spots near Shibuya under 1500 yen
> dinner for 10 in 新宿 with all-you-can-drink
```

The skill needs `hpp` on your `PATH` and `HOTPEPPER_API_KEY` set.

## License

[MIT](LICENSE)
