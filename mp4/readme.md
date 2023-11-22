# MP4

https://w3.org/TR/eme-stream-mp4

pass:

~~~
> packager-win-x64 --enable_raw_key_decryption `
>> --keys key_id=21b82dc2ebb24d5aa9f8631f04726650:key=602a9289bfb9b1995b75ac63f123fc86 `
>> stream=video,in=enc.mp4,output=dec.mp4
[1121/225104:INFO:demuxer.cc(89)] Demuxer::Run() on file 'enc.mp4'.
[1121/225104:INFO:demuxer.cc(155)] Initialize Demuxer for file 'enc.mp4'.
[1121/225104:WARNING:track_run_iterator.cc(699)] Seeing non-zero composition offset 834167. An EditList is probably missing.
[1121/225104:WARNING:track_run_iterator.cc(703)] Adjusting timestamps by -834167. Please file a bug to https://github.com/google/shaka-packager/issues if you do not think it is right or if you are seeing any problems.
~~~

https://github.com/shaka-project/shaka-packager/releases

fail:

~~~
.\mp4ff-decrypt -k 602a9289bfb9b1995b75ac63f123fc86 enc.mp4 dec.mp4
~~~

fail:

~~~
.\mp4ff-decrypt -k 602a9289bfb9b1995b75ac63f123fc86 -init enc.mp4 `
enc.mp4 dec.mp4
~~~

result:

~~~
2023/11/21 22:48:36 decryptSegment: no senc box in traf
~~~

- https://github.com/Eyevinn/mp4ff/issues/298
- https://github.com/Eyevinn/mp4ff/tree/master/cmd/mp4ff-decrypt
