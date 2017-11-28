# wpaclient [![Travis-CI](https://travis-ci.org/brlbil/wpaclient.svg)](https://travis-ci.org/brlbil/wpaclient) [![GoDoc](https://godoc.org/github.com/brlbil/wpaclient?status.svg)](http://godoc.org/github.com/brlbil/wpaclient) [![Go Report Card](https://goreportcard.com/badge/github.com/brlbil/wpaclient)](https://goreportcard.com/report/github.com/brlbil/wpaclient)

Package wpaclient provides a high level wap_supplicant client.

## Install

```bash
go get github.com/brlbil/wpaclient
```

## Usage

### Execute

Client can execute commands, all of the available commands are exported.

```go
// Initialize n new client
client, err := New("wlan0")
if err != nil {
    return err
}
// Close connection
defer client.Close()

// Execute a command, all commands are exported
buf, err := client.Execute(CmdPing)

fmt.Println(buf)
$ PONG
```

### Scan access-points

Scan is a helper function for SCAN and SCAN_RESULTS commands.

```go
aps, err := client.Scan()

for _, ap := range aps {
    fmt.Printf("ssid: %s mac: %s freq: %d signal strength: %d\n",
    ap.SSID, ap.BSSID, ap.Frequency, ap.SignalStrength)
}
```

### List networks

List is a helper function for LIST_NETWORKS command.

```go
nets, err := client.Networks()

for _, nt := range aps {
    fmt.Printf("id: %d, ssid: %s\n", nt.ID, nt.SSID)
}
```

### Get event notifications

Client get event notification by opening a second socket connection.

```go
// Get all events
ch, err := client.Notify()

// Just get connection and disconnection events.
ch, err := client.Notify(WpaEventConnected, WpaEventDisconnected)

ev := <- ch
```

## Credits

 * [Birol Bilgin](https://github.com/brlbil)

## License

The MIT License (MIT) - see [`LICENSE.md`](https://github.com/fatih/color/blob/master/LICENSE.md) for more details