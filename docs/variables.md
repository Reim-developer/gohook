# Variables
**__What are the variables?__**
* In **GoHook** variables are placeholders used to represent dynamic or reusable values in your configuration. They allow you to write cleaner and more maintainable configurations by avoiding hard-coded values.
* Variables follow the syntax:
```
$VARIABLE_NAME
```

**__Variables list:__**
| Syntax              | Description                              |
| ------------------- | ---------------------------------------- |
| `$USER_HOME`        | Your user home directory.                |
| `$TIME_NOW`         | Time now, format it's: `YYYY:MM:DD`.     |
| `$USER_HOSTNAME`    | Your device host name.                   |
| `$USER_OS`          | Your operating system name.              |
| `$LAST_COMMIT_HASH` | Get your last commit hash, if available. |
* You can find too many  examples [here](#example-usage)

## Example usage:

* Show current time:
```toml
[webhook]
url = "Your webhook URL"

[message]
content = "Current time is: $TIME_NOW"
```

* Show current user home:
```toml
[webhook]
url = "Your webhook URL"

[message]
content = "Current user home is: $USER_HOME"
```

* Show current user hostname:
```toml
[webhook]
url = "Your webhook URL"

[message]
content = "Current user hostname is: $USER_HOSTNAME"
```

* Show current user operating system:
```toml
[webhook]
url = "Your webhook URL"

[message]
content = "Current user operating system is: $USER_OS"
```

* Show the last commit hash in your repo:
```toml
url = "Your webhook URL"

[message]
content = "The last commit hash is: $LAST_COMMIT_HASH"
```