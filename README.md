# Slomo
Simple command-line utilty to slow down console output. Simply pipe the output of any command through `slomo` to print out the output at a slower rate.

## Purpose
Sometimes I am too lazy to scroll up through console output and I don't have a search term. So slowing down output makes things more convenient.

## Installation
You can download the release manually or use this command
```bash
# macOS
$ curl https://github.com/ravernkoh/slomo/releases/download/1.0.0/slomo_macos_1.0.0.tar.gz -O -L

# Linux
$ curl https://github.com/ravernkoh/slomo/releases/download/1.0.0/slomo_linux64_1.0.0.tar.gz -O -L
```
Then untar it and move it into somewhere in `PATH`
```bash
# macOS
$ tar xvzf slomo_macos_1.0.0.tar.gz
$ mv slomo /usr/local/bin/slomo
```

## Usage
```bash
# Normal usage
$ ls | slomo

# Change the rate to 50ms
$ cat hello.txt | slomo -rate 50

# Change the buffer size
$ tree | slomo -buf 100
```

## Limitations
- [ ] Since Unix pipes are used, only stdout of previous command is slomo-ed.
- [ ] Installation via curl or manual download.
