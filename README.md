<pre>
 ____                        _                            _
/ ___|  __ _ _ __ ___  _ __ | | ___   ___  ___ _ ____   _(_) ___ ___
\___ \ / _` | '_ ` _ \| '_ \| |/ _ \ / __|/ _ \ '__\ \ / / |/ __/ _ \
 ___) | (_| | | | | | | |_) | |  __/ \__ \  __/ |   \ V /| | (_|  __/
|____/ \__,_|_| |_| |_| .__/|_|\___| |___/\___|_|    \_/ |_|\___\___|
                      |_|
  __               ____         _   _ _       _     _ _       _     _
 / _| ___  _ __   / ___| ___   | | | (_) __ _| |__ | (_) __ _| |__ | |_ ___
| |_ / _ \| '__| | |  _ / _ \  | |_| | |/ _` | '_ \| | |/ _` | '_ \| __/ __|
|  _| (_) | |    | |_| | (_) | |  _  | | (_| | | | | | | (_| | | | | |_\__ \
|_|  \___/|_|     \____|\___/  |_| |_|_|\__, |_| |_|_|_|\__, |_| |_|\__|___/
                                        |___/           |___/
                                _        _   _
 _ __  _ __ ___  ___  ___ _ __ | |_ __ _| |_(_) ___  _ __
| '_ \| '__/ _ \/ __|/ _ \ '_ \| __/ _` | __| |/ _ \| '_ \
| |_) | | |  __/\__ \  __/ | | | || (_| | |_| | (_) | | | |
| .__/|_|  \___||___/\___|_| |_|\__\__,_|\__|_|\___/|_| |_|
|_|

</pre>

# How to run

```bash
go run cmd/theservice/main.go
```

## Using webapp

### Index HTML:

```bash
firefox http://localhost:8080/?name=damian
```

### Github Zen quotation:

```bash
curl http://localhost:8080/github/zen
```

### Echo server:

Echo request headers:

```bash
curl http://localhost:8080/echo/headers  -v
```

Echo request body:

```bash
curl http://localhost:8080/echo/body -H 'application/json' -d '{ "field1": "testValue" }'
```

## Load test

1. Pre-requisites:
   - **NodeJS** installed: `brew install node`
   - **Loadtest** tool installed globally: `sudo npm install -g loadtest`

1. Open Activity Monitor and open **main** app details.

1. Run:
   
   ```bash
   loadtest -n 10000 --rps 300 -T 'application/json' -P '{"field1": "abcdefghijkolmnopqrstuvqxyz", "field2": "abcdefghijkolmnopqrstuvqxyz", "field3": "abcdefghijkolmnopqrstuvqxyz"}' http://localhost:8080/echo/body
   ```

   or:

   ```bash
   loadtest -n 10000 --rps 300 http://localhost:8080
   ```

# How to run tests

```bash
go test ./...
```
