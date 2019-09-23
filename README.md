# gobot

[![Build Badge][1]][2]

## Usage

```
$ bin/gobot

Commands:

  - Place X<Number> Y<Number> Direction<North|East|South|West>
  - Move
  - Left
  - Right

> Place
invalid number of command arguments
> Place 1,2,West
1 2 West
> Move
0 2 West
> Left
0 2 South
> Right
command not implemented
>
```


## Build

```
bin/build
```

## Test

```
bin/test
```

## Run

```
bin/gobot
```

[1]: https://codebuild.us-east-1.amazonaws.com/badges?uuid=eyJlbmNyeXB0ZWREYXRhIjoiMy9EbDM4dHZuWEN6Z3A3bEU2S3ZEcnVXQWdjYmhQSUQxVWs0UHp4WG5XdWY5WThoSkVLa0h4WU84UVFQRDRhdnFVSkFibGVYdTNhT0NIaXhFSFg3K1NJPSIsIml2UGFyYW1ldGVyU3BlYyI6IkFCNHZ5QmJKTUNrL1ZxNjQiLCJtYXRlcmlhbFNldFNlcmlhbCI6MX0%3D&branch=master
[2]: https://console.aws.amazon.com/codesuite/codebuild/projects/github-gobot/history?region=us-east-1
