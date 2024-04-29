<div align="center">

# GCP-SA2BEARER

<img src="./assets/jwtbear.png" align="center" width="200px" height="200px"/>

This CLI tool allows you to get a valid Bearer Token from a GCP ServiceAccount.json file.  
You can find the format of an example SA-File in the **/docs** folder.  

## How it works
The tool creates a jwt and exchanges it for a Bearer token via OAuth2 Workflow.  

## How to use
Just use one of the releases or build from source.  
```
sa2bearer -keyfile=./Path/to/keyfile.json

# Compile from source
go build -o sa2bearer
```

</div>