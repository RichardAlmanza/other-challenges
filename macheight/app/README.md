# Content

- [Content](#content)
  - [Environment](#environment)
  - [Solution](#solution)
    - [Usage](#usage)
    - [Build](#build)
    - [Running the program](#running-the-program)
      - [Reading from a file](#reading-from-a-file)

## Environment

OS: Arch Linux x86_64

Shell: GNU bash, version 5.1.16(1)-release (x86_64-pc-linux-gnu) and zsh 5.9 (x86_64-pc-linux-gnu)

Go: go version go1.20.4 linux/amd64

Docker: Docker version 23.0.6, build ef23cbc431(client) 9dbdbd4b6d(server)

## Solution

### Usage

The program receives two parameters

```bash
./app <numbers separated by comma> <target sum>
```

Assumptions

- There are no repeated numbers
- All number are integers
- The numbers are comma separated without spaces

The program can receive repeated numbers, but, this can lead to an output with repeated pairs.

### Build

We can build the program using

```bash
go build .
```

Or, using docker

```bash
docker run --rm -v "$PWD":/app -w /app golang:1.20 go build .
```

Another option to get the binary without downloading this repository directly

```bash
go install -v github.com/RichardAlmanza/other-challenges/macheight/app@latest
```

Using Docker

```bash
docker run --rm -v "$PWD":/go/bin golang:1.20 go install github.com/RichardAlmanza/other-challenges/macheight/app@latest
```

### Running the program

Also, Go can Build and run a temporal binary

```bash
go run . 1,9,5,0,20,-4,12,16,7 12
```

Or, using docker

```bash
docker run --rm -v "$PWD":/app -w /app golang:1.20 go run . 1,9,5,0,20,-4,12,16,7 12
```

Using the binary compiled from the build stage

```bash
./app 1,9,5,0,20,-4,12,16,7 12
```

#### Reading from a file

Given a file, where each line has the two parameters separated by space

e.g., `cases.lst`

```text
-34,-22,34,37,14,-11,-29,53,23,29,46,-33,-28 -67
-32,17,30,-16,78,-26,71,25,0,31,42,61,33,10,13,-28 31
-34,23,29,46,-32,17,30,-16,73,3,-15,-36,-8,48,76,15,-13 8
1,9,5,0,20,-4,12,16,7 12
1,9,5 6
9,3,20,28,65,29,-40,-8,55,80,70,38,64,-13,8,51,30,67,60,59 0

```

We can use bash to read the file and pass the parameters to the program

```bash
while read -r numbers target; do ./app $numbers $target; echo; done < cases.lst
```
