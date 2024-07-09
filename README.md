# HashWatcher 🔍
Monitors, detects, and notifies changes in files (using SHA-256)

Scenarios where HashWatcher can be useful:
- Security Monitoring: Detect unauthorized changes to critical system or configuration files.
- Configuration Management: Track changes in configuration files to ensure they are intentional and documented.
- Software Development: Monitor source code or project files for unexpected modifications

## Requirements 🔗
<b>Supported OS</b>:
- Linux
- Windows
- macOS (née OS X, aka Darwin)
- OpenBSD
- DragonFly BSD
- FreeBSD
- NetBSD
- Solaris

<b>To build</b>:
- Go compiler

## Quick usage guide 📚
#### Build:
```
go build -o hash-watcher cmd/main.go
```

<br>

> <b>NOTE</b>: Generate in the root directory of the project

---

#### Configuring targets:
Currently, you must create a JSON configuration file for each directory you wish to monitor. Put the file in a non-intrusive directory and name it as you wish. The JSON must follow the structure below:
```json
{
  "directoryPath": "/full/path/to/directory/",
  "fileNames": [],
  "checkFrequencyInSeconds": 30
}

```

You can also filter by specific file:
```json
"fileNames": [
  "file-1.txt",
  "file-2.txt"
],
```

---

#### Usage example:
```
./hash-watcher /full/path/to/configuration/file/settings.json
```

The  `file-1.txt` was initially empty. Upon adding the letter 'X', the change was detected (after 60 seconds), and the hash comparison triggered a notification:

```
2024/07/08 19:59:20 => Watching...

2024/07/08 20:00:50 => file-1.txt was modified! 
  - Before: 01ba4719c80b6fe911b091a7c05124b64eeece964e09c058ef8f9805daca546b 
  - Now: 7058299627365fc7a3dd7840fd3d56f29306cd30c0f2c13cb500fe79617290ff
```

## Future plans 📌
- Check sub-directories
- Create customized schedules
- Write unit and functional tests
- Implement support for notifications in Slack channels

## Do you want help me? 👥
If you have any ideas or wish to contribute to the project, contact me on X (<a href="https://x.com/ohtoaki" target="_blank">@ohtoaki</a>) or send me a pull request :)

## License 🏳️
```
MIT License

Copyright (c) 2024 Vitu Ohto

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
