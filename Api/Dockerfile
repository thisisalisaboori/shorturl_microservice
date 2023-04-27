from mcr.microsoft.com/dotnet/sdk:7.0 as build
COPY ./ /app/
#RUN apt-get update && apt-get install -y clang zlib1g-dev


WORKDIR  /app/
RUN dotnet restore
RUN dotnet publish --os linux --arch x64 -p:PublishProfile=true -c Release  -o out
#-r linux-arm64 -p:PublishTrimmed=true 

FROM mcr.microsoft.com/dotnet/aspnet:7.0
#from ubuntu:20.04
COPY --from=build /app/out /app
WORKDIR /app
RUN apt-get update && apt-get install -y  apt-utils libgdiplus libc6-dev fontconfig   tzdata

ENV TZ Asia/Tehran

ENTRYPOINT ["./Api"] 
