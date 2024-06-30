# TalkToShout
**Made for Habbo: Origins ONLY**

Converts talking packets to shouting packets and the otherway around using [GoEarth](https://github.com/xabbo/goearth)

## Features
- Converts Outgoing CHAT packet to SHOUT
- Converts Outgoing SHOUT packet to CHAT
- Dance from chat using `:dance`
- Wave by saying `o/` (anywhere in the text)
- Hold shift to talk
- **NEW** `:pickall` to pick up all the furni from your room

### Known Issues
- ~~[Other extensions can't block chat packets](https://github.com/Edaurd/TalkToShout/issues/1)~~ Fixed in 1.1

### why?
Switching rooms makes you always select Say instead of Shout by default and I want to switch the functionality of these 2 so that talk will be shout and shout will be talk, 

### Buildcommand used:
`go build -ldflags -H=windowsgui`

`$env:GOOS="darwin"; go build -o bin/${Name}-mac .`

`$env:GOOS="linux"; go build -o bin/${Name}-linux .`

I also included a build.ps1 script in case anyone wants to build it easily themselves

### Credits:
- me for this
- [b7](https://github.com/b7c) for [GoEarth](https://github.com/xabbo/goearth)
