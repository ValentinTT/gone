# Gone

## Date Utility CLI

### Overview

**gone** is a command-line utility written in Go for comparing dates and calculating date differences.

### Features

- **Date Comparison**: Calculate differences in years, months, weeks, days, and hours between two dates.
- **Flexible Input**: Supports multiple date formats (`dd-mm-yyyy`, `mm-dd-yyyy`, `yyyy-mm-dd`) and various separators (`' ', '.', '-', '/'`).
- **Today Flag**: Compare a date with the current date using the `--today` flag.

### Usage

#### Commands

- **`diff`**: Calculate date differences.

  ```sh
  $ gone diff "5 10 2020" "17-6-2024" -f "dmy"

  • start date: Mon Oct 5 2020 (05/10/2020)
  • end date: Mon Jun 17 2024 (17/06/2024)

  Difference:
  • 3 years 8 months 13 days
  • 193 weeks 4 days
  • 1351 days
  • 32424 hours
  ```

  ##### Flags

  - `-l, --layout`: Specify date format (`dmy`, `mdy`, `ymd`). Both dates must be in the same format.

    - `l="dmy"`: `dd-mm-yyyy` (default)
    - `l="mdy"`: `mm-dd-yyyy`
    - `l="ymd"`: `yyyy-mm-dd`


> [!NOTE]
>
> If you ommit the year in the `dmy` layout it takes the current year.

> [!IMPORTANT]
>
> If you use `ymd` or `mdy` the year is required. Otherwise it fails.

  - `-t, --today`: Compare with today's date.
> [!IMPORTANT]
>
> If you pass more than one parameter it fails.

### Installation

1. Clone the repository:

```sh
  git clone https://github.com/ValentinTT/Gone.git
```

2. Build and install the project: make sure to have go added to your local path.

```go
  go build
  go install
```

1. Run the CLI:

```sh
  gone diff "5 10 2020" "17-6-2024" -f "dmy"
```

## Why I did this project?

The reason why I create this project is to learn about go. I think this cli project is small enough to focus in the basic of the language before diving into backend development.

## Author

- [ValentinTT](https://github.com/ValentinTT/)
