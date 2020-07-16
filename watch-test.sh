run="make"

while true
do
    watchman-wait -p "**/*.go" -- .
    eval $run
    echo "-- finished --"
done
