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

- change the function `Run()` in `./internal/service/svc.go` to execute different day functions. 

```bash
$ DAY={day} make run
```
{day} will combine `url.prefix` and `url.suffix` into an url to get the input from http://adventofcode.com

#### example
```bash
$ DAY=1 make run
```