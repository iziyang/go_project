#!/usr/bin/ bash

# Build web UI
cd ./web/
pwd
go install
cp ~/bin/web.exe ~/bin/video_server_web_ui/web.exe
pwd
cp -R ~/src/video_server/template ~/bin/video_server_web_ui/
