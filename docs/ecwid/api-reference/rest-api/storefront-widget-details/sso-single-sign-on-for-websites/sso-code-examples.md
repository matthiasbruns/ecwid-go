# SSO code examples

### PHP

```html
<html><body>
<script>
<?php
if (!$_REQUEST['logoff']) {
        $profile = array(
            // Example values used. Replace with your customer and app details

                'appClientId' => "my-cool-app", 
                'userId' => "234",
                'profile' => array(
                        'email' => "test@example.com",
                        'billingPerson' => array(
                                'name' => "Tester",
                                'companyName' => "Company Name",
                                'street' => "Street",
                                'city' => "City",
                                'countryCode' => "US",
                                'postalCode' => "10001",
                                'stateOrProvinceCode' => "NY"
                        )
                )
        );
        $client_secret = "A1Lu7ANIhKD6A1Lu7ANIhKD6ADsaSdsa";    // this is an example client_secret value
        $message = json_encode($profile);
        $message = base64_encode($message);
        $timestamp = time();
        $hmac = hash_hmac('sha256', "$message $timestamp", $client_secret);   
        echo "var ecwid_sso_profile='$message $hmac $timestamp'";
} else {
        echo "var ecwid_sso_profile=''";
}
?>
</script>
<script src="http://app.ecwid.com/script.js?1003"></script>
<script>
xProductBrowser();
function logout() {
        window.Ecwid.OnAPILoaded.add(function() {
                window.Ecwid.setSsoProfile('');
        });
}
</script>
<a href="javascript: logout()">Log out</a>
</body></html>
```

### VB.Net

Find an example here: <https://github.com/balajiselcom/Ecwid> (thanks to Balaji Sridharan)

### Wordpress

Ecwid official Wordpress plugins uses SSO to sync Wordpress site users with customers in an Ecwid store. You can find the code here: <https://github.com/Ecwid/ecwid-wordpress-plugin>
