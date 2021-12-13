地址：180.169.66.202 9922
账号1：xsw  密码 ACD1357!

# 训练服务器
ssh -p 9922 xsw@180.169.66.202
Acd1357!

# nx box
ssh -p 9222 xsw@180.169.66.202
Acd1357!

scp -P 9222 -r xsw@180.169.66.202:/home/xsw/voiceVideo/ /Users/henryleu/dev/codebase/com/henryleu/go-vad/examples/data/gongan

sox file1.mpg -r 44100 file1-enc.mpg

sox jc_test_1.wav -r 8000 -c 1 -b 16 jc_test_1_8k_16bit.wav
sox jc_test_2.wav -r 8000 -c 1 -b 16 jc_test_2_8k_16bit.wav
sox jc_test_3.wav -r 8000 -c 1 -b 16 jc_test_3_8k_16bit.wav
sox jc_test_4.wav -r 8000 -c 1 -b 16 jc_test_4_8k_16bit.wav
sox jc_test_5.wav -r 8000 -c 1 -b 16 jc_test_5_8k_16bit.wav
sox jc_test_6.wav -r 8000 -c 1 -b 16 jc_test_6_8k_16bit.wav
