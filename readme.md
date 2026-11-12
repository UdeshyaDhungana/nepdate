# nepdate

A command-line tool for working with the Nepali (Bikram Sambat) calendar. Built with [Cobra](https://github.com/spf13/cobra).

## Quick Install

```bash
curl -sSL https://raw.githubusercontent.com/UdeshyaDhungana/nepdate/main/install.sh | bash
```

## Features

- Display today's Nepali date
- Display this month's Nepali calendar
- Convert dates between B.S. and A.D.

## Installation

Build from source:

```
git clone <repo-url>
cd nepdate
go build -o nepdate .
```

## Usage

```
Usage: nepdate [subcommand]

Displays today's Nepali date when run without a subcommand.

Subcommands:
  conv          Convert dates between A.D. and B.S.
  cal           Display Nepali calendar

Flags:

conv
  --ad yyyy/mm/dd    Convert to A.D.
  --bs yyyy/mm/dd    Convert to B.S.
```

## Examples

Show today's Nepali date:

```
nepdate
```

Show this month's calendar:

```
nepdate cal
```

Show this month as a larger grid, with the A.D. date in subscript next to each B.S. day (cannot be combined with `-y`):

```
nepdate cal -e
```

Convert a B.S. date to A.D.:

```
nepdate conv --bs 2083/04/05
```

Convert an A.D. date to B.S.:

```
nepdate conv --ad 2026/07/21
```
