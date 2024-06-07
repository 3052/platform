# amazon

https://amazon.com/gp/video/detail/0NRT15S2XG06SG5HBV5NQAW3E3

## web

1. watch with prime, start your 30-day free trial
2. create your Amazon account
3. your name
4. email
   - email.ml pass
   - mailsac.com fail
5. password
6. re-enter password
7. verify email
8. enter OTP
9. create your Amazon account
10. full name
11. address
12. city
13. state
14. zip
15. phone number
16. continue
17. payment method, change
18. name on card
19. car number
20. month
21. year
22. security code
23. add your card
24. use this address
25. continue
26. start your free trial

for this request:

~~~
POST /cdp/catalog/GetPlaybackResources?deviceID=f1114fdb-4541-47c5-acfe-478978a579f4&deviceTypeID=AOAGZA014O5RE&gascEnabled=false&marketplaceID=ATVPDKIKX0DER&uxLocale=en_US&firmware=1&playerType=xp&operatingSystemName=Windows&operatingSystemVersion=10.0&deviceApplicationName=Firefox64&asin=B0CV72X1BL&consumptionType=Streaming&desiredResources=PlaybackUrls%2CCuepointPlaylist&resourceUsage=ImmediateConsumption&videoMaterialType=Feature&clientId=f22dbddb-ef2c-48c5-8876-bed0d47594fd&userWatchSessionId=b6fdb83c-eb1f-4937-aaa4-39d0bf08bdbb&sitePageUrl=https%3A%2F%2Fwww.amazon.com%2Fgp%2Fvideo%2Fdetail%2F0NRT15S2XG06SG5HBV5NQAW3E3%3Fref_%3Dnav_custrec_signin%26returnFromLogin%3D1&displayWidth=1920&displayHeight=1080&supportsVariableAspectRatio=true&supportsEmbeddedTimedTextForVod=false&deviceProtocolOverride=Https&vodStreamSupportOverride=Auxiliary&deviceStreamingTechnologyOverride=DASH&deviceDrmOverride=CENC&deviceAdInsertionTypeOverride=SSAI&deviceHdrFormatsOverride=None&deviceVideoCodecOverride=H264&deviceVideoQualityOverride=HD&deviceBitrateAdaptationsOverride=CVBR%2CCBR&supportsEmbeddedTrickplayForVod=false&audioTrackId=all&languageFeature=MLFv2&liveManifestType=patternTemplate%2Caccumulating%2Clive&supportedDRMKeyScheme=DUAL_KEY&supportsEmbeddedTrickplay=true&daiSupportsEmbeddedTrickplay=true&daiLiveManifestType=patternTemplate%2Caccumulating%2Clive&ssaiSegmentInfoSupport=Base&ssaiStitchType=MultiPeriod&gdprEnabled=false&playerAttributes=%7B%22middlewareName%22%3A%22Firefox64%22%2C%22middlewareVersion%22%3A%22111.0%22%2C%22nativeApplicationName%22%3A%22Firefox64%22%2C%22nativeApplicationVersion%22%3A%22111.0%22%2C%22supportedAudioCodecs%22%3A%22AAC%22%2C%22frameRate%22%3A%22HFR%22%2C%22H264.codecLevel%22%3A%224.2%22%2C%22H265.codecLevel%22%3A%220.0%22%2C%22AV1.codecLevel%22%3A%220.0%22%7D HTTP/1.1
Host: atv-ps.amazon.com
~~~

if you leave as is, you get:

~~~json
"videoQuality":"HD","videoResolution":"1080p"
~~~

but you will only be able to request key for audio, video will fail:

~~~json
{
   "Widevine2License": {
      "downstreamReason": "vmp_validation_required",
      "errorCode": "PRS.Dependency.DRM.Widevine.HdContentNotAllowed.VmpValidationRequired",
      "message": "Cannot complete request.",
      "type": "PRSWidevine2LicenseDeniedException"
   }
}
~~~

if you use `deviceVideoQualityOverride=SD` or remove `operatingSystemName`, you
get:

~~~json
"videoQuality":"SD","videoResolution":"480p"
~~~

## android phone

https://play.google.com/store/apps/details?id=com.amazon.avod.thirdpartyclient

~~~
> play -i com.amazon.avod.thirdpartyclient  -b armeabi-v7a
details[6] = Amazon Mobile LLC
details[8] = 0 USD
details[13][1][4] = 3.0.370.2045
details[13][1][16] = May 28, 2024
details[13][1][17] = APK
details[13][1][82][1][1] = 5.0 and up
downloads = 652.91 million
name = Amazon Prime Video
size = 46.66 megabyte
version code = 370002045
~~~

Create Android 9 device. even without proxy, app fails after password

https://apkmirror.com/apk/amazon-mobile-llc/amazon-prime-video

Create Android 6 device. the `noarch` options just open the Amazon web page

## android tv

https://play.google.com/store/apps/details?id=com.amazon.amazonvideo.livingroom

~~~
> play -i com.amazon.amazonvideo.livingroom -t -b armeabi-v7a
details[6] = Amazon Mobile LLC
details[8] = 0 USD
details[13][1][4] = 6.16.20+v15.1.0.240-armv7a
details[13][1][16] = Apr 30, 2024
details[13][1][17] = APK APK APK APK APK
details[13][1][82][1][1] = 5.0 and up
downloads = 245.72 million
name = Prime Video - Android TV
size = 42.93 megabyte
version code = 606016201
~~~

Create Android 9 device:

~~~
INSTALL_FAILED_NO_MATCHING_ABIS
~~~

https://apkmirror.com/apk/amazon-mobile-llc/prime-video-android-tv-android-tv

arm only
