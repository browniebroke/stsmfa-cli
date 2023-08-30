from __future__ import annotations

import configparser
from pathlib import Path

import boto3
import typer
from rich import print

app = typer.Typer()


@app.command()
def run(
    token: str = typer.Argument(
        ...,
        help="The 2FA token.",
    ),
    profile: str = typer.Option(
        "default",
        help="The profile to use for obtaining the session token.",
    ),
    mfa_profile: str = typer.Option(
        "",
        help="The profile to write the session data to. Defaults to <profile>-mfa.",
    ),
) -> None:
    """Get a session token using AWS STS and a 2FA token."""
    creds_path = Path.home() / ".aws" / "credentials"
    if not creds_path.exists():
        print(f"[red]Credentials file not found at {creds_path}[/red]")
        raise typer.Exit()

    config = configparser.ConfigParser()
    config.read(creds_path)
    if profile not in config:
        print(
            f"[red]Profile {profile} not found in credentials file {creds_path}[/red]"
        )
        raise typer.Exit()
    if "mfa_serial" not in config[profile]:
        print(f"[red]Profile {profile} does not have an mfa_serial configured[/red]")
        raise typer.Exit()
    mfa_serial = config[profile]["mfa_serial"]

    # get credentials from AWS
    session = boto3.Session(profile_name=profile)
    sts = session.client("sts")
    try:
        response = sts.get_session_token(SerialNumber=mfa_serial, TokenCode=token)
    except Exception as exc:
        print(f"[red]Error getting session token: {exc}[/red]")
        raise typer.Exit() from exc
    credentials = response["Credentials"]

    # write credentials to file
    mfa_profile = mfa_profile or f"{profile}-mfa"
    config[mfa_profile] = {
        "aws_access_key_id": credentials["AccessKeyId"],
        "aws_secret_access_key": credentials["SecretAccessKey"],
        "aws_session_token": credentials["SessionToken"],
    }
    with creds_path.open("w") as cf:
        config.write(cf)

    print(f"[green]All written to {mfa_profile} profile in {creds_path}[/green]")
