# Run with "bundle exec rackup"

require 'rubygems'
require 'bundler'
require 'sinatra'
require 'omniauth'
require 'omniauth-uaa-oauth2'
require 'securerandom'
require 'uaa'


class App < Sinatra::Base
  airports = JSON.load(File.read('airports.json'))
  uaa_url = ENV['UAA_URL']
  uaa_options = {symbolize_keys: true}
  if ENV['UAA_CA_CERT_FILE'] && File.exists?(ENV['UAA_CA_CERT_FILE'])
    uaa_options[:ssl_ca_file] = ENV['UAA_CA_CERT_FILE']
  else
    uaa_options[:skip_ssl_validation] = !!ENV['UAA_CA_CERT']
  end
  puts uaa_options.to_json

  get '/' do
    content_type 'application/json'
    limit = 10
    if auth_header = request.env['HTTP_AUTHORIZATION']
      begin
        decoder = CF::UAA::TokenCoder.decode(auth_header.split(' ')[1], verify: false)
        puts decoder.to_json

        # Unnecessary; but demonstrates that the access_token can be used to make UAA API calls
        # Here the .whoami invokes UAA API GET /userinfo, which requires openid scope.
        info = CF::UAA::Info.new(uaa_url, uaa_options)
        who = info.whoami(auth_header)
        puts who.to_json

        if decoder["client_id"] == "airports" && decoder["iss"] == "#{ENV['UAA_URL']}/oauth/token"
          limit = 20

          if decoder["scope"].include?("airports.all")
            limit = -1
          elsif decoder["scope"].include?("airports.50")
            limit = 50
          end
        end
      rescue CF::UAA::TargetError => e
        return e.info.to_json
      rescue CF::UAA::InvalidToken => e
        return e.info.to_json
      end
    end
    airports[0...limit].to_json
  end
end

use Rack::Session::Cookie,
  key: 'rack.airports',
  path: '/',
  expire_after: 2592000, # In seconds
  secret: ENV.fetch('SESSION_SECRET') { SecureRandom.hex(64) }

run App.new

