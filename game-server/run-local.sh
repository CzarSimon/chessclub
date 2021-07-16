echo "Building game-server"
cd cmd
go build
cd ..
mv cmd/cmd game-server

export MIGRATIONS_PATH='./resources/db/sqlite'

export DB_TYPE='sqlite'
export DB_FILENAME=':memory:'

export JAEGER_SERVICE_NAME='game-server'
export JAEGER_SAMPLER_TYPE='const'
export JAEGER_SAMPLER_PARAM=1
export JAEGER_REPORTER_LOG_SPANS='1'

./game-server