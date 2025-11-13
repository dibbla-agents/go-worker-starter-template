# Credentials Folder

This folder is for storing credentials and sensitive configuration information during development.

## Purpose

Use this folder to store:
- Database connection strings and credentials
- API keys and tokens
- Service account information
- Any other sensitive configuration data

## Usage

Create simple text or markdown files with your credentials. For example:

**postgres.md**
```
Host: localhost
Port: 5432
Database: myapp_db
Username: myuser
Password: mypassword123

Connection String:
postgresql://myuser:mypassword123@localhost:5432/myapp_db
```

**api_keys.md**
```
OpenAI API Key: sk-...
AWS Access Key: AKIA...
AWS Secret: ...
```

## Security

⚠️ **Important**: This entire `docs/credentials/` folder is git-ignored, so files you create here will NOT be committed to version control.

## Tips

- Keep credentials organized in separate files by service
- Use descriptive filenames (e.g., `postgres.md`, `production_api_keys.md`)
- Add comments or notes about where each credential is used
- For production deployments, use environment variables or secret management services instead

