use assert_cmd::Command;

#[test]
fn test_cli_without_args() {
    let mut cmd = Command::cargo_bin("stsmfa").unwrap();
    cmd.assert()
        .failure()
        .stderr(predicates::str::contains("error: the following required arguments were not provided:"));
}

#[test]
fn test_cli_with_args() {
    let mut cmd = Command::cargo_bin("stsmfa").unwrap();
    cmd.arg("--token")
        .arg("123456")
        .arg("--serial-number")
        .arg("arn:aws:iam::123456789012:mfa/user")
        .assert()
        .failure(); // This will fail without proper AWS credentials
} 