# Inhibitor

My personal tool to inhibit my Gnome3 session being marked as idle, locking my
screen. Things like giving a presentation doesn't mark the session as non-idle
the way watching a YouTube video does. Additionally, watching Twitch doesn't
even do this properly.

For added security, the inhibitor can disable itself as soon as a watched
process no longer exists. This can be done by passing a process ID to the -p
option. Optionally, the polling interval of this process detection can be
configured.

## Installation

I use [gb](https://getgb.io/) because I like it, so you can clone and `gb
build` if you like it too.

Another way is to simply `go get`, hoping the up-to-date libraries don't break
anything (not very likely with this simple program).

    go get github.com/Luit/inhibitor/src/inhibitor
