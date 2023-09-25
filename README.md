# Backup Manager

![Issues](https://img.shields.io/github/issues/Yendric/backup-manager)

Single binary backup manager written in Golang.

## Installation

You can install backup-manager by downloading the latest release from the [releases page](https://github.com/Yendric/backup-manager/releases).

### Building from source

> You need to have Golang installed in order to build backup-manager from source.

1. Clone: `git clone git@github.com:Yendric/backup-manager.git`
2. Build: `go build`
3. Configure: `cp config.yaml.example config.yaml`
4. Run: `./backup-manager`

## Backup Manager

Backup Manager is an open source single binary backup manager written in Golang. It allows you to manage your backups in a simple way and sends notifications when doing so.

Backups and notifications are defined in a `config.yaml` file. This file also allows you to configure custom actions you can run on backups (you can do more stuff than just create backups). These custom actions (eg. create, restore, delete, ...) are commands that are executed by the backup manager: this can be a bash script, a binary or anything else that can be executed by the system.

## Configuration

Configuration is done in a `config.yaml` file. You can find an example configuration in `config.yaml.example`.

```yaml
notification:
  email:
    enabled: false
    fromAddress: ""
    smtpHost: "smtp.example.com"
    smtpPort: 587
    smtpUsername: ""
    smtpPassword: ""
    to: []
  webhook:
    enabled: false
    contentField: "content"
    url: "https://hooks.example.com/services/"
backups:
  - name: "backup1"
    createScript: "/path/to/backup1/backup.sh"
    backupDir: "backups"
    actions:
      - name: "restore"
        script: "/path/to/backup1/restore.sh"
      - name: "delete"
        script: "/path/to/backup1/delete.sh"
```

## Example scripts

### Usage with restic

Personally I use backup-manager to manage my restic backups on remote servers. Below I left an example script which you can use as a template to manage your restic backups.

#### create.sh

```bash
# This, besides opening an ssh connection, also creates a tunnel from the remote computer to the local backupserver, which runs restic on port 8000
ssh -R 8000:localhost:8000 username@example.com <<'EOL'
    # Restic repo password
    echo "RESTIC_PASSWORD" > ~/.resticrepo
    # Run restic backup
    restic -r rest:http://localhost:8000/reponame --password-file="/home/username/.resticrepo" --verbose backup /
    # Remove restic repo password
    rm ~/.resticrepo
EOL
```

## Contribute

You're always welcome to create an issue/PR if you have suggestions or find mistakes.
