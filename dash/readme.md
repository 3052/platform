# DASH

~~~
ISO/IEC 23009-1:2019
ISO/IEC 23009-1:2022
~~~

- <https://wikipedia.org/wiki/Dynamic_Adaptive_Streaming_over_HTTP>
- https://standards.iso.org/ittf/PubliclyAvailableStandards

## isoff-live

amc:

~~~
<MPD profiles="urn:mpeg:dash:profile:isoff-live:2011">
  <Period>
    <AdaptationSet mimeType="audio/mp4" startWithSAP="1" segmentAlignment="true" lang="en">
      <ContentProtection schemeIdUri="urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed" xmlns:cenc="urn:mpeg:cenc:2013" bc:licenseAcquisitionUrl="https://manifest.prod.boltdns.net/license/v1/cenc/widevine/6245817279001/5978efd6-b7f7-4476-bcdb-f56ac700da76/26a942cf-a363-488d-a40d-67b0db4c1b19?fastly_token=NjQ3NDkyZDJfZWY4ZTdiYmI1YWIxNmYwMmNkNmE4OGY0ZmQwMjNjZjc1MjAxYThiOGEzMGNlMGM1MTAxNjc3OTA0NWFjYzI4ZA%3D%3D" xmlns:bc="urn:brightcove:2015">
      </ContentProtection>
      <SegmentTemplate initialization="http://redirector.us-east-1.prod-a.boltdns.net/v1/6245817279001/fbeebd45-8380-476d-a3a6-ddcabf55a237/xe0/$RepresentationID$/init.m4f" media="http://redirector.us-east-1.prod-a.boltdns.net/v1/6245817279001/fbeebd45-8380-476d-a3a6-ddcabf55a237/xe0/$RepresentationID$/segment$Number$.m4f" startNumber="0" timescale="48000">
      </SegmentTemplate>
    </AdaptationSet>
  </Period>
</MPD>
~~~

nbc:

~~~xml
<MPD profiles="urn:mpeg:dash:profile:isoff-live:2011">
  <Period duration="PT1H8M51.427S">
    <AdaptationSet contentType="video" group="1" mimeType="video/mp4" segmentAlignment="true" startWithSAP="1" subsegmentAlignment="true" subsegmentStartsWithSAP="1" maxWidth="1920" maxHeight="1080" frameRate="30000/1001" par="16:9" sar="1:1">
      <ContentProtection schemeIdUri="urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed">
      </ContentProtection>
      <Representation bandwidth="1854858" codecs="avc1.4D401F" height="540" id="video_1854858" scanType="progressive" width="960">
        <SegmentTemplate timescale="30000" initialization="1850k_540_cmaf/_226294604_0.mp4" media="1850k_540_cmaf/_226294604_0_$Number$.mp4" startNumber="0">
        </SegmentTemplate>
      </Representation>
    </AdaptationSet>
  </Period>
</MPD>
~~~

paramount:

~~~xml
<MPD profiles="urn:mpeg:dash:profile:isoff-live:2011">
  <Period id="0">
    <AdaptationSet id="6" contentType="audio" lang="en-US" segmentAlignment="true">
      <ContentProtection schemeIdUri="urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed">
      </ContentProtection>
      <Representation id="6" bandwidth="195492" codecs="ec-3" mimeType="audio/mp4" audioSamplingRate="48000">
        <SegmentTemplate timescale="48000" initialization="CBS_SEAL_TEAM_101_HD_R_en-US_1563332_1_eac3_192/init.m4v" media="CBS_SEAL_TEAM_101_HD_R_en-US_1563332_1_eac3_192/seg_$Number$.m4s" startNumber="1">
        </SegmentTemplate>
      </Representation>
    </AdaptationSet>
  </Period>
</MPD>
~~~

## isoff-main

roku:

~~~xml
<MPD profiles="urn:mpeg:dash:profile:isoff-main:2011">
  <Period start="PT0.000S" id="1" duration="PT2775.560S">
    <AdaptationSet id="1485523442" mimeType="video/mp4" segmentAlignment="true" startWithSAP="1" subsegmentAlignment="true" subsegmentStartsWithSAP="1" bitstreamSwitching="true">
      <Representation id="1" width="1920" height="1080" frameRate="25/1" bandwidth="5161512" codecs="avc1.640028">
        <ContentProtection schemeIdUri="urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed">
        </ContentProtection>
        <SegmentTemplate timescale="25000" media="../29fa2dea19214f098321a12d80a02477/1ed39e9befba477eb2f83a61e3e92310/index_video_1_0_$Number$.mp4" initialization="../29fa2dea19214f098321a12d80a02477/1ed39e9befba477eb2f83a61e3e92310/index_video_1_0_init.mp4" startNumber="1" presentationTimeOffset="0">
        </SegmentTemplate>
      </Representation>
    </AdaptationSet>
  </Period>
</MPD>
~~~

## isoff-on-demand

hulu:

~~~xml
<MPD profiles="urn:mpeg:dash:profile:isoff-on-demand:2011">
  <Period id="content-0" duration="PT2657.908S" start="PT0.000S">
    <AdaptationSet mimeType="video/mp4" maxWidth="1280" maxHeight="720" segmentAlignment="true" bitstreamSwitching="true">
      <ContentProtection schemeIdUri="urn:mpeg:dash:mp4protection:2011" value="cenc" cenc:default_KID="21b82dc2ebb24d5aa9f8631f04726650"></ContentProtection>
      <Representation id="018b2012-8917-2054-3da6-00786e737677" codecs="avc1.64001E" bandwidth="1257926" startWithSAP="1" width="704" height="396" frameRate="24000/1001">
        <SegmentBase indexRange="1530-7957" indexRangeExact="true" presentationTimeOffset="0" timescale="10000000">
        </SegmentBase>
      </Representation>
    </AdaptationSet>
  </Period>
</MPD>
~~~
