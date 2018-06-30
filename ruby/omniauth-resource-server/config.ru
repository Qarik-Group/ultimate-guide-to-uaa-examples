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

  get '/' do
    content_type 'application/json'
    limit = 10
    airports[0...limit].to_json
  end

end

use Rack::Session::Cookie,
  key: 'rack.airports',
  path: '/',
  expire_after: 2592000, # In seconds
  secret: ENV.fetch('SESSION_SECRET') { SecureRandom.hex(64) }

run App.new

