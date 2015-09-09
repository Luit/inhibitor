package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/godbus/dbus"
)

func inhibit(o dbus.BusObject, appId, reason string, flags uint32) (cookie uint32) {
	c := o.Call("org.gnome.SessionManager.Inhibit", 0, appId, uint32(0), reason, flags)
	err := c.Store(&cookie)
	if err != nil {
		log.Print(err)
		return 0
	}
	return
}

func uninhibit(o dbus.BusObject, cookie uint32) {
	c := o.Go("org.gnome.SessionManager.Uninhibit", dbus.FlagNoReplyExpected, nil, cookie)
	if c.Err != nil {
		log.Print(c.Err)
	}
}

const (
	inhibitLogout     = 2 ^ iota // 1: Inhibit logging out
	inhibitUserSwitch            // 2: Inhibit user switching
	inhibitSuspend               // 4: Inhibit suspending the session or computer
	inhibitIdle                  // 8: Inhibit the session being marked as idle
)

var (
	pid = flag.Int("p", 0, "PID to wait for")
	d   = flag.Duration("d", time.Second, "sleep duration for polling PID")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
	conn, err := dbus.SessionBus()
	if err != nil {
		log.Fatal(err)
	}
	o := conn.Object("org.gnome.SessionManager", "/org/gnome/SessionManager")
	inhibit(o, "inhibitor", "inhibiting", inhibitIdle)
	if *pid == 0 {
		fmt.Fprint(os.Stderr, "Inhibiting until this process is killed\n")
		select {}
	}
	fmt.Fprintf(os.Stderr, "Inhibiting until this process or PID %d is killed\n", *pid)
	for {
		_, err := os.Stat(fmt.Sprintf("/proc/%d", *pid))
		if os.IsNotExist(err) {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(*d)
	}
}
