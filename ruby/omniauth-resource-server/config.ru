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
        decoder = CF::UAA::TokenCoder.decode(auth_header.split(' ')[1], verify: false)
        puts decoder.to_json
        limit = 20
        info = CF::UAA::Info.new(uaa_url, uaa_options)
        who = info.whoami(auth_header)
        puts who.to_json
        limit = 30

        if decoder["scope"].include?("airports.all")
          limit = -1
        elsif decoder["scope"].include?("airports.50")
          limit = 50
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

