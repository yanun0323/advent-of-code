# Advent Of Code (2022)
https://adventofcode.com

## Requirement
*go version 1.19 or above*

## Usage

#### Setup
```bash
$ cp ./config.yaml.sample ./config.yaml
```
- Then add your browser session into `config.yaml`


#### Run Code

```bash
$ DAY={day} make run 
# example:
# DAY=1 make run
# 
# program will get the data from url "`url.prefix` + 1 + `url.suffix`""
# and try to run func Day1A() and Day1B()
```
