mysql -uroot -ppassword stalefish -h127.0.0.1 -N -e 'show tables' | while read table; do mysql -uroot -ppassword -h127.0.0.1 -e "truncate table $table" stalefish; done
