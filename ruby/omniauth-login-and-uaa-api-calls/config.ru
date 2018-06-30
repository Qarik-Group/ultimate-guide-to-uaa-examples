# Run with "bundle exec rackup"

# Setup $UAA_URL/$UAA_CA_CERT
# source <(path/to/uaa-deployment/bin/uaa-deployment env)
#
# Create UAA client:
# uaa create-client omniauth-login-and-uaa-api-calls -s omniauth-login-and-uaa-api-calls \
#   --authorized_grant_types authorization_code,refresh_token \
#   --scope openid \
#   --redirect_uri http://localhost:9292/auth/cloudfoundry/callback

require 'rubygems'
require 'bundler'
require 'sinatra'
require 'omniauth'
require 'omniauth-uaa-oauth2'
require 'securerandom'

class App < Sinatra::Base
  use Rack::Session::Cookie,
      key: 'rack.omniauth-login-and-uaa-api-calls',
      path: '/',
      expire_after: 2592000, # In seconds
      secret: ENV.fetch('SESSION_SECRET') { SecureRandom.hex(64) }

  get '/' do
    erb :index
  end

  get '/logout' do
    session[:user_email] = nil
    session[:access_token] = nil
    redirect('/')
  end

  get '/auth/cloudfoundry/callback' do
    content_type 'application/json'
    auth = request.env['omniauth.auth']
    session[:user_email] = auth.info.email
    session[:access_token] = auth.credentials.token
    # session[:refresh_token] = auth.credentials.refresh_token
    session[:authorized_scopes] = auth.credentials.authorized_scopes
    redirect('/')
  rescue "No Data"
  end

  get '/auth/failure' do
    content_type 'text/plain'
    request.env['omniauth.auth'].to_hash.inspect rescue "No Data"
  end
end

use Rack::Session::Cookie, :secret => ENV['RACK_COOKIE_SECRET']

use OmniAuth::Builder do
  provider :cloudfoundry, 'omniauth-login-and-uaa-api-calls', 'omniauth-login-and-uaa-api-calls', {
    auth_server_url: ENV['UAA_URL'],
    token_server_url: ENV['UAA_URL'],
    scope: %w[openid],
    skip_ssl_validation: ENV['UAA_CA_CERT'] != '',
    redirect_uri: 'http://localhost:9292/auth/cloudfoundry/callback'
  }
end

run App.new

