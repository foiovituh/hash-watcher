# HashWatcher 🔍
![GitHub License](https://img.shields.io/github/license/foiovituh/hash-watcher)

![banner](https://github.com/foiovituh/hash-watcher/assets/68431603/466efc10-64cf-4ac7-84ca-02c9af63770b)

Monitors, detects, and notifies changes in files (using SHA-256)

Scenarios where HashWatcher can be useful:
- Security Monitoring: Detect unauthorized changes to critical system or configuration files.
- Configuration Management: Track changes in configuration files to ensure they are intentional and documented.
- Software Development: Monitor source code or project files for unexpected modifications

## Summary 📝
- [HashWatcher 🔍](#hashwatcher-)
  - [Summary 📝](#summary-)
  - [How does it work? 💡](#how-does-it-work-)
  - [Requirements 🔗](#requirements-)
  - [Quick usage guide 📚](#quick-usage-guide-)
      - [Build:](#build)
      - [Configuring targets:](#configuring-targets)
      - [Usage example:](#usage-example)
  - [Slack app 💬](#slack-app-)
      - [Main steps:](#main-steps)
  - [Future plans 📌](#future-plans-)
  - [Do you want help me? 👥](#do-you-want-help-me-)
  - [License 🏳️](#license-️)

## How does it work? 💡
SHA-256 is a hash function that works like a fingerprint for files. It converts any text into a 256-bit code. So if anything in the file changes, even a single character, the hash generated will be different. This is very useful (among other things) for checking whether a file has been modified by comparing the current hash with a previous one.

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
  "checkFrequencyInSeconds": 60
}

```

If you prefer not to monitor all files, you can filter specific files to include in the monitoring process:
```json
"fileNames": [
  "file-1.txt",
  "file-2.txt"
],
```

Optionally, you can also send notifications to Slack channels using an app. Endpoint refers to the identifier of the channel where the app is located. Token refers to the "Bot User OAuth Token" for your Workspace. To achieve this, configure the JSON as follows:
```json
{
  "directoryPath": "/full/path/to/directory/",
  "fileNames": [
    "file-1.txt",
    "file-2.txt"
  ],
  "checkFrequencyInSeconds": 60,
  "notification": {
    "endpoint": "CXXXXXXXXXX",
    "token": "xoxb-11111111111-2222222222222-abcdefghijklmnopqrstuvwx"
  }
}

```

---

#### Usage example:
```
./hash-watcher /full/path/to/configuration/file/settings.json
```

The  `file-1.txt` was initially empty. Upon adding the letter 'X', the change was detected (after 30 seconds), and the hash comparison triggered a notification:

```
2024/07/09 11:44:43 => Watching...
2024/07/09 11:45:43 => file-1.txt was modified!
  - Before: e6de32585e70330a8de848b7b7859911e1e108e00dd6527391533853dd7c9409
  - Now: ea6fcfe57703205da4d1b74ec99a8c67f721b2ab2e9c31d2222da066606d5d44
```

## Slack app 💬
Quick guide to create an app using a manifest. For more information, see https://api.slack.com/reference/manifests

#### Main steps:
1- Go to https://api.slack.com/apps, and click on "Create New App".<br>
2- Select "From an app manifest":

![create_an_app](https://github.com/foiovituh/hash-watcher/assets/68431603/b0c85247-301c-47f4-9fb9-758b6216f228)

3- Choose a workspace:

![pick_a_workspace](https://github.com/foiovituh/hash-watcher/assets/68431603/67156b6a-03d2-496f-8374-39d600942065)

4- Copy the YAML from `hash-watcher/doc/slack/slack_app_manifest.yml` and paste it in:

![enter_app_manifest_bellow](https://github.com/foiovituh/hash-watcher/assets/68431603/7e52d43a-78fe-4764-b1f9-ef65f518c5a0)

5- Check that the settings are correct and proceed:

![review_summary_and_create_your_app](https://github.com/foiovituh/hash-watcher/assets/68431603/d1fbf572-ad09-4d85-a1a1-98919fe2419c)

6- Once the app has been created, go to `Settings -> Basic Information` and set up an icon (e.g. the official one in `hash-watcher/doc/slack/logo.png`):

![display_information](https://github.com/foiovituh/hash-watcher/assets/68431603/f9dd1f85-783c-446b-9898-2a580c790ceb)

7- Go to `Features -> OAuth & Permissions -> OAuth Tokens for Your Workspace` and copy the token generated when creating:

![oauth_tokens_for_your_workspace](https://github.com/foiovituh/hash-watcher/assets/68431603/dd64df61-746b-400b-8e6d-4b91420cd791)

8- Finally, choose or create a new channel and add the HashWatcher app to the channel: `Channel details -> Integrations -> Add apps`

## Future plans 📌
- Check sub-directories
- Create customized schedules
- Write unit and functional tests

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

