## summary

TBD

## how to use

GitHub Actions
```bash
TBD
```

Go
```bash
go run cmd/exporter/main.go -folder tech -dt 20210601
```

Docker
```bash
TBD
```

### options

| key | value | required | note |
| --- | --- | --- | --- |
| folder | string | yes | |
| dist | string | optional | default: `./` |
| dt | string `yyyyMMdd` | optional | default: `19700101` |

## outputs (directory)

```
target_directory
├── folder_name
│   └── title1.md
│   └── title2.md
│   └── title3.md
├── folder_name
│   └── title1.md
│   └── title2.md
└── dt.txt // executed timestamp
```

## outputs (markdown file)

This action inserts frontmatter into head of exported markdown file.

| key | value | required |
| --- | --- | --- |
| permalink | string | yes |
| title | string | yes |
| tags | list of string | optional |


Example.

```markdown
---
permalink: "parmalink-of-page"
title: "Title of Page"
tags:
  - golang
  - docker
  - github
---

# First element of Page

hogefuga...
```

## contribute

```
make prepare_env
make up // then runs batch in local
```
