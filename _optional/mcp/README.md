# MCP server starter for Dibbla

A small Go project that becomes your own MCP server when you deploy it to
Dibbla. You write functions; the platform serves them as MCP tools at an
address of their own, requires a Dibbla login, and lets in the people your
app's access list allows.

```bash
dibbla create mcp my-mcp
cd my-mcp
dibbla deploy
dibbla mcp server my-mcp
```

The last command prints what to paste into Claude Code, Cursor, Codex, Gemini
CLI or opencode.

## What is in the project

| File | What it is |
| --- | --- |
| `tools/greeting.go` | The example tool. Copy it to add your own. |
| `tools/tools.go` | The list of tools the server registers. |
| `main.go` | Starts the server. You rarely need to change it. |
| `dibbla.yaml` | The manifest: publishes the server as an MCP and sets who may use it. |
| `Dockerfile` | How the platform builds the app. |
| `env.example` | Settings for running the server on your own machine. |

This starter lives in
[go-worker-starter-template](https://github.com/dibbla-agents/go-worker-starter-template)
under `_optional/mcp`; `dibbla create mcp` copies that directory alone.

There is no MCP code in the project. The server is a Dibbla tool server built
with [sdk-go](https://github.com/dibbla-agents/sdk-go): it connects out to the
platform and registers its functions. The platform speaks MCP to the clients.

## From create to a tool in your client

### 1. Create

`dibbla create mcp <name>` copies this template and writes `<name>` into
`dibbla.yaml`. The name is the last part of the MCP address, so it can hold
letters, digits, `.`, `_` and `-`, at most 64 characters. It must be a name no
other app on your Dibbla installation uses for a tool server.

You need to be signed in (`dibbla login`), so that the project points at your
Dibbla installation.

### 2. Deploy

```bash
dibbla deploy
```

You need the developer role in your organization. No administrator has to
approve or expose anything.

The deploy builds the image, starts the app, and publishes the tool server.
This line in `dibbla.yaml` is what publishes it:

```yaml
services:
  mcp:
    mcp: my-mcp
```

The deploy output names the published server. If another app already owns the
name, the app is still deployed but the server is not published, and the
output says so. Pick another name, change it in both places in `dibbla.yaml`
(`mcp:` and `SERVER_NAME`), and deploy again.

### 3. Connect a client

```bash
dibbla mcp server my-mcp
```

prints the configuration for each client. The address is

```
https://mcp.<your-dibbla-domain>/platform/servers/my-mcp
```

For Claude Code: add the server with the printed `claude mcp add` command,
start a session, and sign in with your Dibbla account when the client asks.
The `greeting` tool is then in the tool list.

To check the whole chain from your terminal:

```bash
dibbla mcp server my-mcp --login    # once per machine
dibbla mcp server my-mcp --check    # lists the tools you get
```

If the address is not found (HTTP 404), per-server MCP addresses are switched
off on your Dibbla installation. Ask the person who operates it.

## Who may use the MCP: `access_policy`

The MCP follows the access list of the app, the same list that decides who may
open the app. It is set in `dibbla.yaml`:

```yaml
    auth:
      require_login: true
      access_policy: all_members
```

| `access_policy` | Who gets the tools |
| --- | --- |
| `all_members` | Every member of your organization. This is the default. |
| `invite_only` | The people you add under **Access & users** on the app in the console, and your organization's owners and admins. |

To limit the MCP to named people, change the value to `invite_only`, deploy,
and add the people in the console. You are on the list yourself from the
start.

Good to know:

- Adding or removing a person applies to their next call. Changing the policy
  can take up to 30 seconds.
- A person who is not let in sees an empty tool list. The address does not
  tell them whether the server exists.
- An MCP always requires a Dibbla login and an account in your organization.
  Removing `require_login` opens the status page of the app, not the MCP.
- `invite_only` limits the MCP address. The functions are also part of your
  organization's function list, which members who build workflows can use.
- Every call is logged with the caller, and the platform's rate limit and your
  organization's quota apply.

## Add a tool

1. Copy `tools/greeting.go` to a new file and rename the types and the
   function.
2. Give it a name and a description. The client shows both to the model, so
   say what the tool does and what it returns.
3. Add it to `Register` in `tools/tools.go`.
4. `dibbla deploy`.

The `json` tags on the input struct are the argument names. A handler that
needs to know who is calling reads it from the context:

```go
caller, ok := sdk.CallerFromContext(ctx)
```

The identity comes from the caller's Dibbla login, not from the arguments, so
you can rely on it when deciding what a person may read or change.

## Run it on your machine

```bash
cp env.example .env     # then fill in SERVER_API_TOKEN
go run .
go test ./...
```

Use a different `SERVER_NAME` locally than the deployed one. A name belongs to
the app that runs it, and the platform refuses a second server under the same
name. A server you run yourself is not published as an MCP. Call its functions
with `dibbla functions invoke` while you develop.

## Take it down

Remove the `mcp:` line and deploy to stop publishing the server while the app
keeps running. Delete the app (`dibbla apps delete <alias>`) to remove it and
free the name.
