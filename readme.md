# Releaster

Releaster is a GitHub Action that helps users generate release drafts from the latest release to the latest pull request. This tool is inspired by the configuration file of [release-drafter/release-drafter](https://github.com/release-drafter/release-drafter).

## Usage

To use Releaster, simply add it as a step in your GitHub Action workflow:

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

When using Releaster, make sure to pass in the `github-token` and grant it `pull-requests` and `contents` write permissions.

## Customization

Releaster can be customized by creating a configuration file. By default, Releaster will look for a `.github/releaster.yml` file for configuration. To specify a different file path, use the `config-path` input:

```yaml
      - uses: JZGoopi/releaseter@v1
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          config-path: 'apple/banana.yml'
```

This will cause Releaster to look for a configuration file at `.github/apple/banana.yml`.

### Configuration File Template

The following is an example configuration file for Releaster:

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

clear-history-draft: true
```

The configuration file supports the following options:

- `name-template`: The template for the release title. It supports the following placeholders:
  - `{{ $VERSION_NEXT_MAJOR }}`: The next major version.
  - `{{ $VERSION_NEXT_MINOR }}`: The next minor version.
  - `{{ $VERSION_NEXT_PATCH }}`: The next patch version.
  - `{{ $TIME_WORKFLOW }}`: The time the workflow is run.
  - `{{ $EMPTY }}`: An empty string.
  - The version numbers follow [Semantic Versioning 2.0.0](https://semver.org/spec/v2.0.0.html) and are based on the `tag-template`.
- `tag-template`: The template for the tag name. It supports the same placeholders as `name-template`. The default value is `'v{{ $VERSION_NEXT_PATCH }}'`.
- `tag-preRelease`: The pre-release version, following [Semantic Versioning 2.0.0](https://semver.org/spec/v2.0.0.html). The default value is an empty string.
- `tag-build`: The build metadata, following [Semantic Versioning 2.0.0](https://semver.org/spec/v2.0.0.html). The default value is an empty string.
- `categories`: An array of categories, each with a title and one or more labels. Pull requests with matching labels will be grouped under the corresponding category title in the release draft.
- `category-other`: An object that controls whether to include pull requests that do not match any of the specified categories. It has two keys:
  - `show`: A boolean value that determines whether to include unmatched pull requests. The default value is `true`.
  - `title`: The title for the section that includes unmatched pull requests. The default value is `'OTHER'`.
- `clear-history-draft`: A boolean value that determines whether to clear existing draft releases. The default value is `false`.

## Conclusion

Releaster is a helpful tool for generating release drafts from the latest release to the latest pull request. With its customizable configuration options, you can tailor it to your specific needs.


## Welcome Feedback

If you have any feedback or suggestions for Releaster, please feel free to submit an issue on the [GitHub repository](https://github.com/JZGoopi/releaseter/issues). Your feedback helps us improve the tool for everyone!

## License

Releaster is licensed under the [MIT License](https://opensource.org/licenses/MIT). Feel free to use, modify, and distribute this tool for any purpose.
