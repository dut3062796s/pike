# Pike

HTTP cache server such as varnish.

## script

### dev

You shuld install go and nodejs, then run the scripts:
```bash
# use etcd for config's storage
go run main.go --config etcd://127.0.0.1:2379/pike --init

# use file for config's storage
go run main.go --config /tmp --init
```

```bash
cd web && yarn start
```

then open `http://127.0.0.1:3015/` in the browser.
