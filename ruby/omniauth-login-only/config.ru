# Run with "bundle exec rackup"

# Setup $UAA_URL/$UAA_CA_CERT
# source <(path/to/uaa-deployment/bin/uaa-deployment env)
#
# Create UAA client:
# uaa create-client omniauth-login-only -s omniauth-login-only \
#   --authorized_grant_types authorization_code,refresh_token \
#   --scope openid \
#   --redirect_uri http://localhost:9292/auth/cloudfoundry/callback

require 'rubygems'
require 'bundler'
require 'sinatra'
require 'omniauth'
require 'omniauth-uaa-oauth2'

class App < Sinatra::Base
  get '/' do
    <<-HTML
    <ul>
      <li><a href='/auth/cloudfoundry'>Sign in with Cloud Foundry</a></li>
    </ul>
    HTML
  end

  get '/auth/cloudfoundry/callback' do
    content_type 'application/json'
    request.env['omniauth.auth'].to_hash.to_json rescue "No Data"
  end

  get '/auth/failure' do
    content_type 'text/plain'
    request.env['omniauth.auth'].to_hash.inspect rescue "No Data"
  end
end

use Rack::Session::Cookie,
  key: 'rack.omniauth-login-only',
  path: '/',
  expire_after: 2592000, # In seconds
  secret: ENV.fetch('SESSION_SECRET') { SecureRandom.hex(64) }

use OmniAuth::Builder do
  options = {
    auth_server_url: ENV['UAA_URL'],
    token_server_url: ENV['UAA_URL'],
    redirect_uri: 'http://localhost:9292/auth/cloudfoundry/callback'
  }
  if ENV['UAA_CA_CERT_FILE'] && File.exists?(ENV['UAA_CA_CERT_FILE']) && !File.read(ENV['UAA_CA_CERT_FILE']).strip.empty?
    options[:ssl_ca_file] = ENV['UAA_CA_CERT_FILE']
  else
    options[:skip_ssl_validation] = !!ENV['UAA_CA_CERT']
  end

  provider :cloudfoundry, "omniauth-login-only", "omniauth-login-only", options
end

run App.new

