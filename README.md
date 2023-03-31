# Sniffer

国科大，网络空间安全学院，网络攻防基础课程，实验一

## Requirements

[实验要求](./assets/exercise-1.pdf)

## Features

- [x] Pause/Resume
- [x] BPF Filter
- [x] Render abstract of packets
- [x] Render details of packets
- [x] HTTP Assembly
- [ ] Flow Tracking

## Known Issues

1. 效率差点意思。瓶颈在于 Wails 框架前后端不支持流式传输，导致 packet 在显示的时候前后端开销比较大。除非更换更加 native
   的框架否则无解
