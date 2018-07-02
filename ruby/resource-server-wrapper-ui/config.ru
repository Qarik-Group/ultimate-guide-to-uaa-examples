# Run with "bundle exec rackup"

# Setup $UAA_URL/$UAA_CA_CERT
#   source <(path/to/uaa-deployment/bin/uaa-deployment env)
#
# Create UAA client:
#   uaa create-client airports-map -s airports-map \
#     --authorized_grant_types authorization_code,refresh_token \
#     --scope openid,airports-50,airports-all \
#     --redirect_uri http://localhost:9393/auth/cloudfoundry/callback
#
# Run as :9393 (assuming backend airports app on :9292)
#   bundle exec shotgun -p 9393

require 'rubygems'
require 'bundler'
require 'sinatra'
require 'omniauth'
require 'omniauth-uaa-oauth2'
require 'securerandom'
require 'uaa'

UAA_URL = ENV['UAA_URL']
UAA_CLIENT = 'airports-map'
UAA_CLIENT_SECRET = 'airports-map'
UAA_OPTIONS = {symbolize_keys: true}
if ENV['UAA_CA_CERT_FILE'] && File.exists?(ENV['UAA_CA_CERT_FILE'])
  UAA_OPTIONS[:ssl_ca_file] = ENV['UAA_CA_CERT_FILE']
else
  UAA_OPTIONS[:skip_ssl_validation] = !!ENV['UAA_CA_CERT']
end
puts UAA_OPTIONS.to_json

class App < Sinatra::Base

  get '/' do
    if session[:user_email]
      erb :index
    else
      erb :guest
    end
  end

  get '/logout' do
    session[:user_email] = nil
    session[:access_token] = nil
    session[:authorized_scopes] = nil
    redirect('/')
  end

  get '/auth/cloudfoundry/callback' do
    puts "SESSION on callback:"
    p session
    content_type 'application/json'
    auth = request.env['omniauth.auth']
    session[:user_email] = auth.info.email
    # session[:access_token] = auth.credentials.token
    session[:refresh_token] = auth.credentials.refresh_token
    session[:authorized_scopes] = auth.credentials.authorized_scopes
    redirect('/')
  rescue "No Data"
  end

  get '/auth/failure' do
    content_type 'text/plain'
    request.env['omniauth.auth'].to_hash.inspect rescue "No Data"
  end
end

use Rack::Session::Cookie,
  key: 'rack.airports-map',
  path: '/',
  expire_after: 2592000, # In seconds
  secret: ENV.fetch('SESSION_SECRET') { SecureRandom.hex(64) }

use OmniAuth::Builder do
  options = {
    auth_server_url: UAA_URL,
    token_server_url: UAA_URL,
    redirect_uri: 'http://localhost:9393/auth/cloudfoundry/callback'
  }
  if ENV['UAA_CA_CERT_FILE'] && File.exists?(ENV['UAA_CA_CERT_FILE']) && !File.read(ENV['UAA_CA_CERT_FILE']).strip.empty?
    options[:ssl_ca_file] = ENV['UAA_CA_CERT_FILE']
  else
    options[:skip_ssl_validation] = !!ENV['UAA_CA_CERT']
  end

  provider :cloudfoundry, UAA_CLIENT, UAA_CLIENT_SECRET, options
end

run App.new

