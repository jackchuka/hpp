# hpp

CLI tool for searching Japanese restaurants using the [HotPepper Gourmet API](https://webservice.recruit.co.jp/doc/hotpepper/reference.html).

## Install

### Homebrew

```bash
brew install jackchuka/tap/hpp
```

### Go

```bash
go install github.com/jackchuka/hpp@latest
```

### Agent Skill

This repository includes a [Claude Code](https://docs.anthropic.com/en/docs/claude-code) skill. Install the skill to search restaurants with natural language:

```bash
npx skills add jackchuka/hpp
```

## Setup

Get a free API key from [Recruit Web Service](https://webservice.recruit.co.jp/register/) and set it:

```bash
export HOTPEPPER_API_KEY=your_key_here
```

## Usage

### Claude Code (natural language)

Use `/restaurant-search` or just ask in conversation:

```
> find an izakaya near 浜松町 for 6 people with private rooms
> ramen spots near Shibuya under 1500 yen
> dinner for 10 in 新宿 with all-you-can-drink
```

### Search restaurants

```bash
# Keyword search
hpp search --keyword "ramen" --area Z011

# Location search (lat/lng + range)
hpp search --lat 35.6812 --lng 139.7671 --range 3

# Filter by features
hpp search --keyword "izakaya" --wifi --private-room --english --non-smoking

# All-you-can-drink spots with lunch in Shinjuku
hpp search --keyword "新宿" --free-drink --lunch

# JSON output
hpp search --keyword "sushi" --format json

# Pagination
hpp search --keyword "ramen" --count 20 --start 1
```

### Search by shop name or phone

```bash
hpp shop --keyword "居酒屋"
hpp shop --tel 0312345678
```

### Browse genres

```bash
hpp genre
hpp genre --keyword ramen
```

### Browse areas

```bash
# Large areas (prefectures/regions)
hpp area large
hpp area large --keyword tokyo

# Middle areas within a large area
hpp area middle --large-area Z011

# Small areas within a middle area
hpp area small --middle-area Y005
```

### Browse service areas

```bash
hpp service-area large
hpp service-area list
```

### Browse budgets

```bash
hpp budget
```

### Browse credit card types

```bash
hpp creditcard
```

### Browse specials/features

```bash
hpp special list
hpp special list --category SPC0
hpp special category
```

### Version

```bash
hpp version
```

### Global flags

| Flag | Description | Default |
|------|-------------|---------|
| `--format` | Output format: `table` or `json` | `json` |

## Search flags

| Flag | Description |
|------|-------------|
| `--keyword` | Free text search |
| `--name` | Shop name (partial match) |
| `--lat`, `--lng`, `--range` | Location search (range: 1=300m, 2=500m, 3=1km, 4=2km, 5=3km) |
| `--area` | Large area codes |
| `--middle-area` | Middle area codes |
| `--genre` | Genre codes |
| `--budget` | Budget codes |
| `--wifi` | Has WiFi |
| `--lunch` | Lunch service |
| `--english` | English menu |
| `--private-room` | Has private rooms |
| `--non-smoking` | Non-smoking seats |
| `--parking` | Has parking |
| `--pet` | Pet allowed |
| `--free-drink` | All-you-can-drink |
| `--free-food` | All-you-can-eat |
| `--midnight` | Open after 11pm |
| `--card` | Accepts cards |
| `--count` | Results per page (max 100) |
| `--order` | Sort: 1=name, 2=genre, 3=area, 4=recommended |

Run `hpp search --help` for the full list of 50+ flags.

## API Coverage

All 12 HotPepper API endpoints are supported:

| Endpoint | Command |
|----------|---------|
| Gourmet Search | `hpp search` |
| Shop Name Search | `hpp shop` |
| Genre Master | `hpp genre` |
| Budget Master | `hpp budget` |
| Large Area Master | `hpp area large` |
| Middle Area Master | `hpp area middle` |
| Small Area Master | `hpp area small` |
| Large Service Area | `hpp service-area large` |
| Service Area | `hpp service-area list` |
| Credit Card Master | `hpp creditcard` |
| Special Master | `hpp special list` |
| Special Category | `hpp special category` |

## License

MIT
