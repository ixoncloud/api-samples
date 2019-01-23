# Agent Event Downloader

This is a simple CLI application which allows you to export agent event data to CSV format.

Using this data you can for instance analyze agent downtime or other events that you are interested in.

## Usage

**Windows**

In powershell (or `cmd`):

Set password environment variable:

`$env:IXON_IXPLATFORM_PASSWORD = "foobar"`

Execute the tool:

```
.\agent-event-dl_vX.X.X.exe
    --application <application_id>
    --user <email>
    --company <company_id>
    --agent <agent_id>
    --after <ISO8601 timestamp (start time)>
    --before <ISO8601 timestamp (end time)>
    --output <output csv filename>
```

Example:

```
.\agent-event-dl_v1.0.0.exe
    --application PvFddOGlMxFh
    --user john.doe@gmail.com
    --company 1111-2222-3333-4444-5502
    --agent 4jV2914dPDVL
    --after 2019-01-23T08:30:18+0000
    --before 2019-01-30T08:30:18+0000
    --output output.csv
```

**Linux/MacOS**

In terminal:

Set password environment variable:

`export IXON_IXPLATFORM_PASSWORD=foobar`

```
./agent-event-dl_vX.X.X
    --application <application_id>
    --user <email>
    --company <company_id>
    --agent <agent_id>
    --after <ISO8601 timestamp (start time)>
    --before <ISO8601 timestamp (end time)>
    --output <output csv filename>
```

Example:

```bash
./agent-event-dl_v1.0.0
    --application PvFddOGlMxFh
    --user john.doe@gmail.com
    --company 1111-2222-3333-4444-5502
    --agent 4jV2914dPDVL
    --after 2019-01-23T08:30:18+0000
    --before 2019-01-30T08:30:18+0000
    --output output.csv
```

**Two factor authentication**

If you have 2FA enabled on your Ixon Cloud account, you'll need to specify your token with `--otp <otp token>`
