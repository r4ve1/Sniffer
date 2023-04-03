# Sniffer

[r4ve1/Sniffer (github.com)](https://github.com/r4ve1/Sniffer)

## 技术栈

抓包: [google/gopacket: Provides packet processing capabilities for Go (github.com)](https://github.com/google/gopacket)

GUI: [The Wails Project | Wails](https://wails.io/)

- 前端：Vue3 + Vite + Naive UI

## Features

- [x] Pause/Resume
- [x] BPF Filter
- [x] Render abstract of packets
- [x] Render details of packets
- [x] HTTP Assembly
- [ ] Flow Tracking

## Build

1. Install [Npcap](https://npcap.com/)

2. Install Golang

3. Install Wails

   ```bash
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```

4. Build

   ```bash
   # production
   wails build
   # develop
   wails build -ldflags="-s -w" -upx
   ```