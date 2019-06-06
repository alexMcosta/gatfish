## Create the layout for the the Accepted YAML file
### Resources
1. List of AWS resources that take tagging; https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html
2. AWS Tagging best practices: https://aws.amazon.com/answers/account-management/aws-tagging-strategies/
3. Add account info to YAML config

#### YAML config 
The YAML will have product type followed by keys, values, than level of requirement like so:

```
ALL:
  Name: REQUIRED    
```

---

## Requirement settings
#### Levels of requirement from most to least:

**DIRE**: Will immediately remove the resource regardless of anything. Useful if the tag is required for a security or cost reason. These will appear at the top of the logs when Gatfish is ran so the team in charge of handling the cloud accounts will know who to notify.

**REQUIRED**: If the resource can be stopped then it will be stopped. The resource will also get logged below DIRE.

**WANTED**: Will Log that the resource has not been tagged with this. Eventually you can tag this resource with a time stamp so you can give the team / owner a deadline to tag the resource.

**AUTO**: If the tag can be automatically added such as having an owner then AUTO will tag the resource with the Owner tag by gathering the user information from Cloudtrail or some sort of external resource.

---

## KNOWN ISSUES

1. EBSCompliance will not report the volume ID if it has not tags
