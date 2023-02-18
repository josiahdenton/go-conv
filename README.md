# Go Conv

Converts different measurements, such as 
- temperature
- weight
- length
as a cli tool.

## Usage
  - `-from / -f`
    - Sets what we are converting from. Possible values are: `m ,ft,kg,lb,c,f`
- `-mode / -m`
    - Sets the conversion mode, expected values are: `t, l, w`
- `-to / -t`
    - Sets what we are converting to. Has the same possible values as from. 

#### Examples
convert from celsius to fahrentheit
```bash
conv -m=t -f=c -t=f 30
# outputs: 30
```
erroneously convert celsius to kilograms
```bash
conv -m=t -f=c -t=kg 30
#outputs: <timestamp> unsupported conversion!
```

### Changelog

- 0.0.1
    - basic functionality

### Roadmap

- [x] converts between temperature, weight, and length
- [ ] add unit tests 
- [ ] add an install script for bin
    - [ ] macos
    - [ ] linux
    - [ ] windows