use aws_config::BehaviorVersion;
use aws_sdk_sts::Client as StsClient;
use clap::Parser;
use std::process;

#[derive(Parser)]
#[command(author, version, about)]
struct Cli {
    #[arg(short, long)]
    token: String,
    
    #[arg(short, long)]
    serial_number: String,
    
    #[arg(short, long, default_value = "3600")]
    duration: i32,
}

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let cli = Cli::parse();
    
    let config = aws_config::defaults(BehaviorVersion::latest())
        .load()
        .await;
    
    let client = StsClient::new(&config);
    
    let resp = client
        .get_session_token()
        .token_code(&cli.token)
        .serial_number(&cli.serial_number)
        .duration_seconds(cli.duration)
        .send()
        .await?;
    
    if let Some(credentials) = resp.credentials {
        println!("export AWS_ACCESS_KEY_ID={}", credentials.access_key_id.unwrap_or_default());
        println!("export AWS_SECRET_ACCESS_KEY={}", credentials.secret_access_key.unwrap_or_default());
        println!("export AWS_SESSION_TOKEN={}", credentials.session_token.unwrap_or_default());
        Ok(())
    } else {
        eprintln!("Failed to get credentials");
        process::exit(1);
    }
} 