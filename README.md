## High Performance Fizz Buzz 

Trying to get Fizz Buzz to go fast
Current fastest is
(Testing is done with `./main | pv > /dev/null`)
~ 60 MiB 

Using `strconv.FormatInt(int64(n), 10)` instead of `fmt.Sprintf` improves performance to 
` 288MiB 0:00:03 [ 103MiB/s]`

Using channels to create sequential multi-core results in:
` 577MiB 0:00:03 [ 195MiB/s]` (sometimes 200 MiB. Inconsistent)

IF YOU RuN INTO AN ISSUE YOU HARDCODED THE BYTES MY DUDE. CHECK THE CHUNK BYTE

## How 2 use pprof to analyze
```bash
kevin@localhost ~/p/g/s/g/K/fizzbuzz> go build  -ldflags "-s -w" cmd/main.go                                                                                                    main!?
kevin@localhost ~/p/g/s/g/K/fizzbuzz> ./main 2&>/dev/null > cpu.pprof   
```

Have to uncomment these lines
```golang
pprof.StartCPUProfile(os.Stdout)
defer pprof.StopCPUProfile()
```