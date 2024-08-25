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

The log output will look like this:

```txt
2024/08/25 11:04:54 Internet connection restored
2024/08/25 11:20:22 Internet connection lost
2024/08/25 11:20:48 Internet connection restored
2024/08/25 11:31:53 Internet connection restored
2024/08/25 11:35:42 Internet connection lost
2024/08/25 11:36:02 Internet connection restored
2024/08/25 12:03:45 Internet connection lost
2024/08/25 12:04:14 Internet connection restored
2024/08/25 12:52:48 Internet connection lost
2024/08/25 12:53:03 Internet connection restored
```