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
  uaa_options = {skip_ssl_validation: ENV['UAA_CA_CERT'] != '', symbolize_keys: true}

  get '/' do
    content_type 'application/json'
    limit = 10
    if auth_header = request.env['HTTP_AUTHORIZATION']
      begin
        info = CF::UAA::Info.new(uaa_url, uaa_options)
        who = info.whoami(auth_header)
        puts who.to_json
        limit = 20
        who
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

