from pathlib import Path

import pytest
from click.exceptions import Exit

from stsmfa.cli import run

HOME_PATH = Path.home()
CREDS_PATH = Path.home() / ".aws" / "credentials"


class TestRun:
    @pytest.fixture()
    def rich_print(self, mocker):
        return mocker.patch("stsmfa.cli.print")

    def test_no_file(self, fs, rich_print):
        with pytest.raises(Exit):
            run(token="123456", profile="dummy", mfa_profile="")
        rich_print.assert_called_once_with(
            f"[red]Credentials file not found at {CREDS_PATH}[/red]"
        )

    def test_profile_missing(self, fs, rich_print):
        fs.create_file(CREDS_PATH)
        with pytest.raises(Exit):
            run(token="123456", profile="dummy", mfa_profile="")
        rich_print.assert_called_once_with(
            "[red]Profile dummy not found in credentials file " f"{CREDS_PATH}[/red]"
        )

    def test_mfa_device_missing(self, fs, rich_print):
        fs.create_file(CREDS_PATH, contents="[dummy]\n")
        with pytest.raises(Exit):
            run(token="123456", profile="dummy", mfa_profile="")
        rich_print.assert_called_once_with(
            "[red]Profile dummy does not have an mfa_serial configured[/red]"
        )

    def test_get_sts_error(self, fs, rich_print, mocker):
        fs.create_file(
            CREDS_PATH,
            contents="[dummy]\nmfa_serial = arn:aws:iam::123456789012:mfa/dummy\n",
        )

        boto3 = mocker.patch("stsmfa.cli.boto3")
        sts = boto3.Session.return_value.client.return_value
        sts.get_session_token.side_effect = Exception("SOME ERROR")

        with pytest.raises(Exit):
            run(token="123456", profile="dummy", mfa_profile="")

        rich_print.assert_called_once_with(
            "[red]Error getting session token: SOME ERROR[/red]"
        )

    def test_get_sts_ok(self, fs, rich_print, mocker):
        fs.create_file(
            CREDS_PATH,
            contents="[dummy]\nmfa_serial = arn:aws:iam::123456789012:mfa/dummy\n",
        )

        boto3 = mocker.patch("stsmfa.cli.boto3")
        sts = boto3.Session.return_value.client.return_value
        sts.get_session_token.return_value = {
            "Credentials": {
                "AccessKeyId": "AKSOMEACCESSKEY",
                "SecretAccessKey": "SOMESECRET",
                "SessionToken": "SOMETOKEN",
            }
        }

        run(token="123456", profile="dummy", mfa_profile="")

        rich_print.assert_called_once_with(
            f"[green]All written to dummy-mfa profile in {CREDS_PATH}[/green]"
        )
        assert Path(CREDS_PATH).read_text() == (
            "[dummy]\n"
            "mfa_serial = arn:aws:iam::123456789012:mfa/dummy\n"
            "\n"
            "[dummy-mfa]\n"
            "aws_access_key_id = AKSOMEACCESSKEY\n"
            "aws_secret_access_key = SOMESECRET\n"
            "aws_session_token = SOMETOKEN\n"
            "\n"
        )
