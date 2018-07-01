#!/usr/bin/env ruby

require 'json'
require 'uaa'
require 'highline'
require 'terminal-table'

uaa_url = ENV['UAA_URL']
uaa_client = "airports"
uaa_client_secret = "airports"
uaa_options = {symbolize_keys: false}
if ENV['UAA_CA_CERT_FILE'] && File.exists?(ENV['UAA_CA_CERT_FILE'])
  uaa_options[:ssl_ca_file] = ENV['UAA_CA_CERT_FILE']
else
  uaa_options[:skip_ssl_validation] = !!ENV['UAA_CA_CERT']
end

airports_url = ENV['AIRPORTS_URL'] || "http://localhost:9292/"
user_config_path = ENV['AIRPORTS_USER_CONFIG_PATH'] || ".user.json"

begin
  unless File.exist?(user_config_path) && JSON.parse(File.read(user_config_path))["refresh_token"]
    token_issuer = CF::UAA::TokenIssuer.new(uaa_url, uaa_client, uaa_client_secret, uaa_options)

    cli = HighLine.new
    creds = token_issuer.prompts.inject({}) do |prompts, prompt|
      label, meta = prompt
      prompt_type, _ = meta
      if prompt_type == "text"
        prompts[label] = cli.ask("#{label} > ")
      else
        prompts[label] = cli.ask("#{label} > ") { |q| q.echo = "*" }
      end
      prompts
    end
    token = token_issuer.owner_password_grant(creds["username"], creds["password"])
    File.open(user_config_path, "w") {|f| f << token.info.to_json}
  end

  user_config = JSON.parse(File.read(user_config_path))
  token_info = CF::UAA::TokenInfo.new(user_config)

  httpclient = HTTPClient.new
  httpclient.default_header = {"Authorization": token_info.auth_header}
  airports = JSON.parse(httpclient.get(airports_url).body)

  rows = airports.map {|airport| [airport["Name"], airport["ICAO"]]}
  table = Terminal::Table.new :headings => ['Name', 'ICAO'], :rows => rows
  puts table
rescue CF::UAA::TargetError => e
  $stderr.puts e.info
end