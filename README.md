# docker-mysql-backup-restore  
A tool set to backup and restore the mysql dumps to and from s3.  

##Pre-requisites  
1.Docker installed.  

##Environment variables required  
`dump_path`: the local path to dump to  
`dumper_db_host`: the host name of the mysql db  
`dumper_db_port`: the port where mysql server runs  
`dumper_db_user`: the user name of the db  
`dumper_db_password`: the password of the db  
`dumper_db_name`: the db name for which the dump has to be done  
`s3_access_key`:  the access key  
`s3_secret_key`:  the secret key  
`s3_bucket_name`:  the name of the bucket to upload the backup to  
`dumper_s3_region`: the region in which the bucket with `s3_bucket_name` is created  
`path_in_bucket`:  the path in the bucket to upload to  


##Usage
Suppose there is a DB named `sample` at host `sample.db.com` on port `3306`.The user is `backuper` and the password is `pswd`.The access key is `access_key` and the secret key is `secret_key`.The dump has to be uploaded to `sample-db-backups` bucket in the `us-west-1` region.The path in the bucket is `/data`.The path on the local is `/home/user/db-backups`
the usage is as follows:  

```
docker run  kaddiya/mysql-backup-restore -v /home/user/backups:/backup 
--env dumper_db_host=sample.db.com --env dumper_db_port=3306 --env dumper_db_user=backuper \
--env dumper_db_password="pswd" --env dumper_db_name="sample" --env s3_access_key=access_key \
--env s3_secret_key="secret_key" --env s3_bucket_name="sample-db-backups" \
--dumper_s3_region="us-west-1" --env path_in_bucket="/data"
```
