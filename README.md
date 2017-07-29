# gif-wall 

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

## Description
The program is decomposed in three parts.

#### 1 - The GO Twitter Stream Crawler Thingy Backend
This part like the others one is very simple. Basically it ask twitter for a personnalied twitter stream and watch if there is some gif that passed throw. As soon as it see one it save the url of the gif to Firebase.

#### 2 - Firebase
Firebase is a company acquired by Google that allow us to store and retrieve item in real time using a super simple API and because this API is available in numerous langage, it was my choice. Once again dead simple.

### 3 - The Client Thingy 
This part is a simple html page that communicate with Firebase to know if a new gif has been found and show it in a video player that is in fullscreen, autoplay, loop mode. You can check it here : http://fabienfoerster.github.io/gif-wall/ !
( and yes a video player because gif in twitter are converted to mp4 video but you know why not)



## First step

```bash
# get the dependencies
go get 
# build the binary
go install
# run the binary
$GOPATH/bin/gif-wall
```
## Environment variables
The app use the twitter API so you'll need to set 5 environment variables
```
heroku config:set TWITTER_API_KEY=XXX
heroku config:set TWITTER_API_SECRET=XXX
heroku config:set TWITTER_ACCESS_TOKEN=XXX
heroku config:set TWITTER_ACCESS_TOKEN_SECRET=XXX
heroku config:set TWITTER_STREAM_FOLLOW=personid1,personid2,...
```

## To deploy to heroku

First you need to check if you have godep : https://github.com/tools/godep

```
go get github.com/tools/godep
```
Then make the dependencies know to heroku
```
godep save -r
```

An finally the standard heroku command 
```
heroku create -b https://github.com/kr/heroku-buildpack-go.git
git push heroku master
heroku open
```

And if you don't want to do that you can just click on the deploy to heroku button <3
