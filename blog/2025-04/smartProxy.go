package smartProxy

type test struct {
   email    string
   response string
}

const suspicious = `In case of suspicious activity being detected, customers
are required to undergo an ID verification process, which are then reviewed by
a third-party ID verification tool.`

var Tests = []test{
   {
      email:    "367@tuta.io",
      response: suspicious,
   },
   {
      email:    "10317@proton.me",
      response: suspicious,
   },
   {
      email:    "srpen7@proton.me",
      response: suspicious,
   },
}
