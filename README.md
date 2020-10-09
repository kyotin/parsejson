# parsejson

nohup ./parse -inJson=/home/ubuntu/Apollo.io_none_July_2018-PART-1.json -out=./Apollo.io_none_July_2018-PART-1-have-email.json -workers=1000 -buffLines=10000 -filter=have_email > log_par1_have_email.log &  
nohup ./parse -inJson=/home/ubuntu/Apollo.io_none_July_2018-PART-2.json -out=./Apollo.io_none_July_2018-PART-2-have-email.json -workers=1000 -buffLines=10000 -filter=have_email > log_par2_have_email.log &  

nohup ./parse -inJson=/home/ubuntu/Apollo.io_none_July_2018-PART-1.json -out=./Apollo.io_none_July_2018-PART-1-have-phone.json -workers=1000 -buffLines=10000 -filter=have_phone > log_par1_have_phone.log &  
nohup ./parse -inJson=/home/ubuntu/Apollo.io_none_July_2018-PART-2.json -out=./Apollo.io_none_July_2018-PART-2-have-phone.json -workers=1000 -buffLines=10000 -filter=have_phone > log_par2_have_phone.log &  

nohup ./parse -inJson=/home/ubuntu/Apollo.io_none_July_2018-PART-1.json -out=./Apollo.io_none_July_2018-PART-1-have-email-phone.json -workers=1000 -buffLines=10000 -filter=have_email,have_phone > log_par1_have_both.log &  
nohup ./parse -inJson=/home/ubuntu/Apollo.io_none_July_2018-PART-2.json -out=./Apollo.io_none_July_2018-PART-2-have-email-phone.json -workers=1000 -buffLines=10000 -filter=have_email,have_phone > log_par2_have_both.log &  

nohup ./parse -inJson=./Apollo.io_none_July_2018-PART-1-have-phone.json -out=./Apollo.io_none_July_2018-PART-1-have-phone-33.json -workers=1000 -buffLines=10000 -filter=have_phone_33 > log_par1_have_phone_33.log &  
nohup ./parse -inJson=./Apollo.io_none_July_2018-PART-2-have-phone.json -out=./Apollo.io_none_July_2018-PART-2-have-phone-33.json -workers=1000 -buffLines=10000 -filter=have_phone_33 > log_par2_have_phone_33.log &  

nohup ./parse -inJson=./Apollo.io_none_July_2018-PART-1-have-phone.json -out=./Apollo.io_none_July_2018-PART-1-have-phone-336.json -workers=1000 -buffLines=10000 -filter=have_phone_336 > log_par1_have_phone_336.log &  
nohup ./parse -inJson=./Apollo.io_none_July_2018-PART-2-have-phone.json -out=./Apollo.io_none_July_2018-PART-2-have-phone-336.json -workers=1000 -buffLines=10000 -filter=have_phone_336 > log_par2_have_phone_336.log &  

nohup ./parse -inJson=./Apollo.io_none_July_2018-PART-1-have-phone.json -out=./Apollo.io_none_July_2018-PART-1-have-phone-337.json -workers=1000 -buffLines=10000 -filter=have_phone_337 > log_par1_have_phone_33.log &  
nohup ./parse -inJson=./Apollo.io_none_July_2018-PART-2-have-phone.json -out=./Apollo.io_none_July_2018-PART-2-have-phone-337.json -workers=1000 -buffLines=10000 -filter=have_phone_337 > log_par2_have_phone_33.log &  

nohup ./parse -inJson=/home/ubuntu/Apollo.io_none_July_2018-PART-1.json -out=./Apollo.io_none_July_2018-PART-1-company-ibm.json -workers=1000 -buffLines=10000 -filter=company_ibm &  
nohup ./parse -inJson=/home/ubuntu/Apollo.io_none_July_2018-PART-2.json -out=./Apollo.io_none_July_2018-PART-2-company-ibm.json -workers=1000 -buffLines=10000 -filter=company_ibm &  

nohup ./parse -inJson=/home/ubuntu/Apollo.io_none_July_2018-PART-1.json -out=./Apollo.io_none_July_2018-PART-1-company-facebook.json -workers=1000 -buffLines=10000 -filter=company_facebook &  
nohup ./parse -inJson=/home/ubuntu/Apollo.io_none_July_2018-PART-2.json -out=./Apollo.io_none_July_2018-PART-2-company-facebook.json -workers=1000 -buffLines=10000 -filter=company_facebook &  

nohup ./parse -inJson=/home/ubuntu/Apollo.io_none_July_2018-PART-1.json -out=./Apollo.io_none_July_2018-PART-1-company-google.json -workers=1000 -buffLines=10000 -filter=company_google &  
nohup ./parse -inJson=/home/ubuntu/Apollo.io_none_July_2018-PART-2.json -out=./Apollo.io_none_July_2018-PART-2-company-google.json -workers=1000 -buffLines=10000 -filter=company_google &  

nohup ./randompick -inJson=./Apollo.io_none_July_2018-PART-2.json -out=./randompick-part-2.json -numOfRows=10000 -maxRowsInFile=84750000 &  
nohup ./randompick -inJson=./Apollo.io_none_July_2018-PART-1.json -out=./randompick-part-1.json -numOfRows=10000 -maxRowsInFile=1350000 &  

nohup ./distinct -inJson=./Apollo.io_none_July_2018-PART-1.json -out=./Apollo.io_none_July_2018-PART-1-distinct.json -maxRowsInFile=1350000 > distinct-part-1.log &   
nohup ./distinct -inJson=./Apollo.io_none_July_2018-PART-2.json -out=./Apollo.io_none_July_2018-PART-2-distinct.json -maxRowsInFile=84750000 > distinct-part-2.log &  

### Push to s3
nohup aws s3 cp Apollo.io_none_July_2018-PART-1-have-phone.json s3://thedirector/processed/ &  
nohup aws s3 cp Apollo.io_none_July_2018-PART-1-have-email-phone.json s3://thedirector/processed/ &  
nohup aws s3 cp Apollo.io_none_July_2018-PART-1-have-email.json s3://thedirector/processed/  &  
nohup aws s3 cp Apollo.io_none_July_2018-PART-2-have-email-phone.json s3://thedirector/processed/ &  
nohup aws s3 cp Apollo.io_none_July_2018-PART-2-have-phone.json s3://thedirector/processed/ &  
nohup aws s3 cp Apollo.io_none_July_2018-PART-2-have-email.json s3://thedirector/processed/ &  

nohup aws s3 cp Apollo.io_none_July_2018-PART-1-have-phone-337.json s3://thedirector/processed/ &  
nohup aws s3 cp Apollo.io_none_July_2018-PART-1-have-phone-336.json s3://thedirector/processed/ &  
nohup aws s3 cp Apollo.io_none_July_2018-PART-1-have-phone-33.json s3://thedirector/processed/ &  
nohup aws s3 cp Apollo.io_none_July_2018-PART-2-have-phone-336.json s3://thedirector/processed/ & 
nohup aws s3 cp Apollo.io_none_July_2018-PART-2-have-phone-33.json s3://thedirector/processed/ &  
nohup aws s3 cp Apollo.io_none_July_2018-PART-2-have-phone-337.json s3://thedirector/processed/ &  

### Summarize
count of total row: 86100000 rows
* 1350000 Apollo.io_none_July_2018-PART-1.json
* 84750000 Apollo.io_none_July_2018-PART-2.json

count of row with email: 54259302 rows
* 781250 Apollo.io_none_July_2018-PART-1-have-email.json
* 53478052 Apollo.io_none_July_2018-PART-2-have-email.json

count of row with phone: 32958 rows
* 174 Apollo.io_none_July_2018-PART-1-have-phone.json
* 32784 Apollo.io_none_July_2018-PART-2-have-phone.json

count of row with phone + email: 31502 rows
* 162 Apollo.io_none_July_2018-PART-1-have-email-phone.json
* 31340 Apollo.io_none_July_2018-PART-2-have-email-phone.json

count of row where phone begin by +33: 75 rows
* 1 Apollo.io_none_July_2018-PART-1-have-phone-33.json
* 74 Apollo.io_none_July_2018-PART-2-have-phone-33.json

count of row where phone begin by +33 6 or +33 7: 15 rows
* 14 Apollo.io_none_July_2018-PART-2-have-phone-336.json
* 0 Apollo.io_none_July_2018-PART-1-have-phone-336.json
* 0 Apollo.io_none_July_2018-PART-1-have-phone-337.json

count of unique row (by whole line): 
* 1349794 Apollo.io_none_July_2018-PART-1-distinct.json (Duplicated: 206)
* 83920250 Apollo.io_none_July_2018-PART-2-distinct.json (Duplicated: 829750)