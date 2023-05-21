# Releaseter

[![Go version](https://img.shields.io/badge/Go-%3E%3D1.19-blue)](https://golang.org/doc/go1.19)
[![License](https://img.shields.io/badge/License-MIT-blue)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/v/release/JZGoopi/releaseter)](https://github.com/JZGoopi/releaseter/releases)
[![Code coverage](https://img.shields.io/codecov/c/github/JZGoopi/releaseter)](https://codecov.io/gh/JZGoopi/releaseter)
[![Code quality](https://img.shields.io/lgtm/grade/go/github/JZGoopi/releaseter)](https://lgtm.com/projects/g/JZGoopi/releaseter/context:go)
[![Last commit](https://img.shields.io/github/last-commit/JZGoopi/releaseter)](https://github.com/JZGoopi/releaseter/commits/main)

Releaseter is a GitHub Action that helps users generate release drafts from the latest release to the latest pull request. This tool is inspired by the configuration file of [release-drafter/release-drafter](https://github.com/release-drafter/release-drafter).

## Usage

To use Releaseter, simply add it as a step in your GitHub Action workflow:

```yaml
name: example
on:
  workflow_dispatch:

jobs:
  test:
    runs-on: ubuntu-latest
    permissions:
      pull-requests: write
      contents: write
    steps:
      - uses: actions/checkout@v3
      - uses: JZGoopi/releaseter@v1
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
```

When using Releaseter, make sure to pass in the `github-token` and grant it `pull-requests` and `contents` write permissions.

## Customization

Releaseter can be customized by creating a configuration file. By default, Releaseter will look for a `.github/releaseter.yml` file for configuration. To specify a different file path, use the `config-path` input:

```yaml
      - uses: JZGoopi/releaseter@v1
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          config-path: 'apple/banana.yml'
```

This will cause Releaseter to look for a configuration file at `.github/apple/banana.yml`.

### Configuration File Template

The following is an example configuration file for Releaseter:

```yaml
name-template: '{{ $TIME_WORKFLOW }}'

tag-template: 'v{{ $VERSION_NEXT_PATCH }}'
tag-preRelease: ""
tag-build: ""

categories:
  - title: '🚀 Features'
    labels:
      - 'feature'
      - 'feat'
  - title: '🐛 Bug Fixes'
    labels:
      - 'fix'
      - 'bugfix'
      - 'bug'

category-other:
  show: true
  title: 'OTHERRR'

category-except-labels:
  - wont_release
  - documentation

clear-history-draft: true

time-format: '06/01/02 15:04:05'
time-location: 'Asia/Taipei'

# If you want to use the except_releases option, the tag of each item is required.
except_releases:
  - tag: 'v1.0.1'           # require
    name: '230510 00:01:51' # option

except_keyword: '#jump'
```
#### name-template && tag-template

The `name-template` is the template for the release title. It supports several placeholders and defaults to `{{ $TIME_WORKFLOW }}`.

The `tag-template` is the template for the tag. It also supports several placeholders and defaults to `v{{ $VERSION_NEXT_PATCH }}`.

Both `name-template` and `tag-template` support the following placeholders:

- `{{ $VERSION_NEXT_MAJOR }}`: the next major version
- `{{ $VERSION_NEXT_MINOR }}`: the next minor version
- `{{ $VERSION_NEXT_PATCH }}`: the next patch version
- `{{ $TIME_WORKFLOW }}`: the time when the workflow is run
- `{{ $EMPTY }}`: an empty string
- The version numbers follow [Semantic Versioning 2.0.0](https://semver.org/spec/v2.0.0.html) as defined by GitHub, and Releaseter will use the latest tag that matches the `tag-template` as a baseline.


#### tag-preRelease && tag-build

`tag-preRelease` and `tag-build` are pre-release and build tags for [Semantic Versioning 2.0.0](https://semver.org/spec/v2.0.0.html), respectively. They default to an empty string.


#### categories

`categories` is an array of categories, each with a title and a list of labels. If a pull request's labels match any of the specified labels, its title will be added to the corresponding category. The format will be:

```
### Category Title
---
- PR1
- PR2
```


#### category-other

`category-other` controls whether to include pull requests that do not match any of the specified labels. It has two keys:

- `show`: A boolean value that determines whether to include unmatched pull requests. The default value is `false`.
- `title`: The title for the section that includes unmatched pull requests. The default value is `'OTHER'`.

#### category-except-labels

`category-except-labels` Specifies an array of labels. Pull requests with any of these labels will be excluded from the release notes.

#### clear-history-draft

`clear-history-draft` controls whether to clear existing drafts. It defaults to `false`.

#### time-format

`time-format` is the format for the time. It uses Golang's placeholder syntax and defaults to `'060102 15:04:05'`. The available placeholders are:

- Year:
  - `2006`: four-digit year
  - `06`: two-digit year
- Month:
  - `1` or `01`: numeric month
  - `Jan`: abbreviated month name
- Day: `2` or `02`
- Hour:
  - `3`: 12-hour clock hour
  - `15`: 24-hour clock hour
- Minute: `4` or `04`
- Second:
  - `5` or `05`
  - `5.000000`: decimal fraction of a second
- Timezone:
  - `MST`: Mountain Standard Time
  - `UTC`: Coordinated Universal Time
  - `EST`: Eastern Standard Time
  - `CST`: Central Standard Time
  - `PST`: Pacific Standard Time
  - `-0700`: numeric timezone offset from UTC (e.g. -7 hours)
  - `-07:00`: numeric timezone offset from UTC with colon separator (e.g. -7:00)
- Weekday:
  - `Mon`: abbreviated weekday name
  - `Monday`: full weekday name
- AM/PM:
  - `PM`: uppercase AM/PM
  - `pm`: lowercase AM/PM

You can use these placeholders to format the time in the way that you prefer. For example, if you want to include the abbreviated weekday name and the timezone offset, you can use the following format:

```yaml
time-format: 'Mon, Jan 02 2006 15:04:05 -0700'
```

This will produce a time string like `Fri, May 07 2023 10:30:00 -0700`, where `Fri` is the abbreviated weekday name, `May 07 2023` is the date, `10:30:00` is the time, and `-0700` is the timezone offset.


#### time-location

`time-location` Controls the timezone used for timestamps. The value must be a valid name from the IANA Time Zone database. The default value is `'Asia/Taipei'`.


#### except_releases

Releaseter is a tool that captures all PR content from the time of the latest release to the current trigger of the Github action. However, sometimes for various reasons, a new release is manually written before the actual version is released, causing the Releaseter to refresh the time and miss some PRs. In this case, you can use except_releases to exclude those releases and obtain the actual latest release time point you want to capture.

You don't need to write this parameter for every release you want to exclude. You only need to write the ones you want to exclude from the latest release up until the desired release. Once Releaseter finds a release that is not excluded, it will recognize that release as the latest one you want to capture.

- `tag`: The tag name bound to the release is a required parameter.
- `name`: ，
  - Optional parameters.
  - The title of the release is provided as a parameter because there is a small possibility that two releases are bound to the same tag, but you only want to delete one of them. In this case, you can use the name parameter to more accurately find the release you want to delete.

#### except_keyword

The functionality of `except_keyword` is similar to `except_releases`. The difference is that `except_keyword` specifies a keyword, and when the title of a release contains that keyword, the release will be skipped. The process continues until a release that is not excluded is found, and that release is used as the reference time point.


## Conclusion

Releaseter is a helpful tool for generating release drafts from the latest release to the latest pull request. With its customizable configuration options, you can tailor it to your specific needs.


## Welcome Feedback

If you have any feedback or suggestions for Releaseter, please feel free to submit an issue on the [GitHub repository](https://github.com/JZGoopi/releaseter/issues). Your feedback helps us improve the tool for everyone!

## License

Releaseter is licensed under the [MIT License](https://opensource.org/licenses/MIT). Feel free to use, modify, and distribute this tool for any purpose.
