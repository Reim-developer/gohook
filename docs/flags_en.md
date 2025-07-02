# GoHook CLI Flags

## Command:
`wh-send`
* **Description:** Send content or embed to Discord webhook via TOML configurations.

| Flag            | Type   | Default | Description                                                |
| --------------- | ------ | ------- | ---------------------------------------------------------- |
| `--verbose`     | bool   | false   | Show payload after sending. Optional.                      |
| `--dry-run`     | bool   | false   | Do not send webhook, just print payload.                   |
| `--loop`        | int    | `1`     | Send webhook multiple times.                               |
| `--delay`       | int    | `2`     | Delay between each webhook in loop (seconds).              |
| `--threads`     | int    | `1`     | Future use. Not yet implemented.                           |
| `--explicit`    | bool   | false   | Print full Discord response (message ID, channel ID, etc). |
| `--use-env-url` | string | ""      | Use environtment variable instead of `url` setting         |

## Examples
**Send webhook once:**
```bash
gohook wh-send settings.toml
```

**Dry run, don't send any:**
```bash
gohook wh-send settings.toml --dry-run
```

**Loop 5 times with 3s delay:**
```bash
gohook wh-send settings.toml --loop 5 --delay 3
```

**Send webhook with environment variable:**
```bash
WEBHOOK_URL=https://discord.com/api/webhooks/abc/xyz gohook wh-send settings.toml --use-env-url WEBHOOK_URL
```

**For more settings examples, please visit:**
* https://github.com/Reim-developer/gohook/tree/stable/examples