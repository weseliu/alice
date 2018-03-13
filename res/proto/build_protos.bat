cd /d %~dp0

protoc --go_out=..\src\proto\logic  logicproto.proto
protoc-3.0.0-beta-4-windows-x86_64.exe --csharp_out F:\XunTengGame\Client\GameHall\Assets\Scripts\Protocol\ logicproto.proto

protoc --go_out=..\src\proto\command  commandproto.proto
protoc-3.0.0-beta-4-windows-x86_64.exe --csharp_out F:\XunTengGame\Client\GameHall\Assets\Scripts\Protocol\ commandproto.proto