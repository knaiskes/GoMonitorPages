## A Go website monitor

**How it works**
You add the urls of the webpage(s) that want to monitor in a list (slice) and the program constanly checks their status code,
 if any of the webpage in the list has status code other than 200 an e-mail will be sent to you with the url of the webpage and its status code, after this the url of this webpage will be removed from the temporary from the list for 20 minutes.
 This means that if a webpage for example has status code for an hour it will be sent to you three e-mails.

**Requirements**
Go 1.8 or greater
A Gmail address

**Before running it**
Add in the **email.go** file your Gmail address and password and the address that you want to be emailed.

```
// enter your credentials
receiver := ""
sender := ""
passw := ""
```

Add the webpage's urls that you want to monitor in the slice called **urls** which can be found in **watchSites.go** file
