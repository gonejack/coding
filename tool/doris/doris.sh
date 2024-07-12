curl --location-trusted -u root:hadoop \
    -H "Expect:100-continue" \
    -H "column_separator:," \
    -H "columns:user_id,name,age" \
    -T test.csv \
    -XPUT http://192.168.11.34:8030/api/research/test/_stream_load
