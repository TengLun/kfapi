# Library for the Fraud Identification API

This is a privately created library to aid in using the Fraud Identification API.

Note that this library is not officially endorsed. 
This is simply my library to help me interact with their API, described here:

https://support.kochava.com/analytics-reports-api/fraud-api

# Retrieve Data

Effort has been made to make this library simple and readable.

First, initiaze the client using your API key, and account id. The system will automatically detect if you are using the library
as a marketer or network (which changes the calls).

```golang
client, err := kfapi.GetAccount("my_api_key", "my_accound_id")
```

Once the account has been initialized, data can be returned using the "list" endpoint, or the "data" endpoint. 

```golang
response, err := client.List.Apps("fraud type",start time, end time)
```

Example:
```golang
response, err := client.List.Apps(kfapi.AnonymousInstall, time.Unix(1510000000,0), time.Now())
```

# Constants

When gathering information,
it must be accessed for one particular fraud type at a time. For convenience, constants have been added and exported from the library.

The constants are directly accessible from the library:
```golang
kfapi.AdStacking
kfapi.AnonymousInstall
kfapi.DeviceHighClickVolume
...
etc.
```

# Improvements that will happen

Currently, functionality is being worked on that will allow the return to be ported directly back to the API to add suspicious actors to the 
blacklist, after some threshold has been applied.

Stay tuned for more developments.


