# mysql-backup-restore  
A tool set leveraging Docker to backup and restore the mysql dumps to and from s3.It aims to ease up the backup and restore process of databases.

##Pre-requisites  
1.Docker installed on the VM that is to run this job.  
2.IAM user with rights to push to s3 and its keys.  


##Getting the Image  
`docker pull kaddiya/mysql-backup-restore`


##Environment variables required  
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
`mode` : the mode to run the tool chain in.either BACKUP or RESTORE.


##Usage for BACKUP  
###Case 1: To backup a DB which is running outside of a docker container.  
Suppose there is a DB named `sample` at host `sample.db.com` on port `3306`.The user is `user` and the password is `pswd`.The access key is `access_key` and the secret key is `secret_key`.The dump has to be uploaded to `sample-db-backups` bucket in the `us-west-1` region.The path in the bucket is `/data`.  
Then the usage is as follows:  

```
docker run \
--env dumper_db_host=sample.db.com --env dumper_db_port=3306 --env dumper_db_user=user \
--env dumper_db_password="pswd" --env dumper_db_name="sample" --env s3_access_key=access_key \
--env s3_secret_key="secret_key" --env s3_bucket_name="sample-db-backups" --env mode=BACKUP \
--dumper_s3_region="us-west-1" --env path_in_bucket="/data" kaddiya/mysql-backup-restore
```  

###Case 2: To backup a DB which is running inside of a docker container.   
Suppose there is a DB named `sample` running inside a docker container named  `container.db.com` on port `3306`.The user is `user` and the password is `pswd`.The access key is `access_key` and the secret key is `secret_key`.The dump has to be uploaded to `sample-db-backups` bucket in the `us-west-1` region.The path in the bucket is `/data`.  
Then the usage is as follows:  

```
docker run  \
--env dumper_db_host=container.db.com --env dumper_db_port=3306 --env dumper_db_user=user \
--env dumper_db_password="pswd" --env dumper_db_name="sample" --env s3_access_key=access_key \
--env s3_secret_key="secret_key" --env s3_bucket_name="sample-db-backups" \
--dumper_s3_region="us-west-1" --env path_in_bucket="/data"  --env mode=BACKUP \
--link container.db.com:container.db.com kaddiya/mysql-backup-restore
```

##Usage for RESTORE    
###Case 1: To restore a DB  which is running outside of a docker container with the latest backup.  
Suppose there is a DB named `sample` at host `sample.standby.db.com` on port `3306`.The user is `user` and the password is `pswd`.The access key is `access_key` and the secret key is `secret_key`.The db has to be restored from the dump uploaded to `/data/latest` in the `sample-db-backups` bucket in the `us-west-1` region.  
Then the usage is as follows:  

```
docker run \
--env dumper_db_host=sample.standby.db.com --env dumper_db_port=3306 --env dumper_db_user=user \
--env dumper_db_password="pswd" --env dumper_db_name="sample" --env s3_access_key=access_key \
--env s3_secret_key="secret_key" --env s3_bucket_name="sample-db-backups" --env mode=RESTORE \
--dumper_s3_region="us-west-1" --env path_in_bucket="/data" kaddiya/mysql-backup-restore
```  
###Case 2: To restore a DB  which is running inside of a docker container with the latest backup.
Suppose there is a DB named `sample` running inside a docker container named  `container.standby.db.com` on port `3306`.The user is `user` and the password is `pswd`.The access key is `access_key` and the secret key is `secret_key`.The db has to be restored from the dump uploaded to `/data/latest` in the `sample-db-backups` bucket in the `us-west-1` region.  
Then the usage is as follows:  

```
docker run --link container.standby.db.com:container.standby.db.com \
--env dumper_db_host=container.db.com --env dumper_db_port=3306 --env dumper_db_user=user \
--env dumper_db_password="pswd" --env dumper_db_name="sample" --env s3_access_key=access_key \
--env s3_secret_key="secret_key" --env s3_bucket_name="sample-db-backups" \
--dumper_s3_region="us-west-1" --env path_in_bucket="/data"  --env mode=RESTORE \
 kaddiya/mysql-backup-restore
```

##Note on the directory structure
###Directory Structure of the backups on the host

```
/backups
  /latest
    sample-latest-backup.sql
    error.log
  /archived
    26-December-2016-15:50.sql
```


###Directory Structure of the backups on s3

```
/sample-db-backups
  /data
    /latest
      sample-latest-backup.sql
      error.log
    /archived
      sample-26-December-2016-15:50.sql
```
