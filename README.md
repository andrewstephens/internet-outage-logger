# Internet Outage Logger

This is a simple little program I threw together to track when my internet connection was cutting out.

The program has two flags you can pass:
- interval to check (in seconds)
- log file path

```sh
./outage-logger -interval 5 -logfile outages.txt
```

You can also use the short versions of those flags:

```sh
./outage-logger -i 5 -l outages.txt
```