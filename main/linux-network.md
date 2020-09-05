# linux网络

## http通讯流程

1. HTTP `GET google.com http1.1`
2. DNS `domain name: google.com; IP 39.156.69.79`
3. TCP `src port: 5678; dst port: 80`
4. IP `src IP: 120.244.152.147; dst IP 39.156.69.79`
5. MAC `src MAC: dc:a9:04:8f:98:aa; src GateWay MAC: 32:35:2f:dc:e4:8a`
6. Router `dst {router1 ip}, {router2 ip}, {router3 ip}, {router4 ip} ...;`

## ref

[理解TCP](https://www.jianshu.com/p/ca64764e4a26)
