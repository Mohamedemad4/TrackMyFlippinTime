# TrackMyFlippinTime
a web application (with the company of a cli script),to track your time accross **devices**(still no support for native mobile)
you can run it as a service anywhere heroku,your PC,A pi, the nighboars unpatched router the possiblites are endless!

## Service API

the service API has 4 endpoints 
####newstatement
```
GET /newstatement/<statement>/<statement_encoded>
```

where Statement is the full statement denoted by statment_encoded in other requests

##### transaltestatement
```
GET /transaltestatement/<statement_encoded>/
```
return the actual statement from an encoded statement (Remeber Statements is how you discibe shunks of your time)
##### deposit
```
GET /deposit/<fromtimestamp>/<totimestamp>/<statement_encode>
```

where fromtimestamp and totimestamp are Unix Timestamps of what he where you doing between the 1st and the 2nd respectivly
and where statement_encode is the encoded statement do denot the activity

##### withdraw
```
GET /withdraw/<fromtimestamp>/<totimestamp>/
```

where fromtimestamp and totimestamp are Unix Timestamps 

and it returns JSON history deposits in the time between fromtimestamp and totimestamp
