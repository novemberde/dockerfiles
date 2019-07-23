
    vegeta attack -rate 5000 -duration 10m 'GET http://ip-172-31-11-128.ap-northeast-2.compute.internal:8080'

    echo "GET http://ip-172-31-11-128.ap-northeast-2.compute.internal:8080" | vegeta attack -rate=1000/s > results.gob