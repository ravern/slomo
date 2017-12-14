# Slomo
Simple command-line utilty to slow down console output. Simply pipe the output of any command through `slomo` to print out the output at a slower rate.

## Purpose
Sometimes I am too lazy to scroll up through console output and I don't have a search term. So slowing down output makes things more convenient.

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
