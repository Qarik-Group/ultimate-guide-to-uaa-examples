#!/bin/bash

set -eu

ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
cd $ROOT

: ${UAA_URL:?required}
if [[ "${UAA_CA_CERT:-X}" == "X" ]]; then
  check_custom_ca=$(curl -s --insecure -v $UAA_URL 2>&1 | grep "unable to get local issuer certificate")
  if [[ "${check_custom_ca:-X}" != "X" ]]; then
    >&2 echo "ERROR: UAA ${UAA_URL} has a custom certificate. Require: \$UAA_CA_CERT"
    exit 1
  fi
else
  tmp=$(mktemp -d) # temp fix for https://github.com/cloudfoundry/cf-uaac/issues/60
  export UAA_CA_CERT_FILE="${tmp}/ca.pem"
  echo "${UAA_CA_CERT}" > "${UAA_CA_CERT_FILE}"
fi

bundle exec rackup --host 0.0.0.0 --port ${PORT:-9393}