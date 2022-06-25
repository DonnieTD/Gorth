#    --no-cache 
docker build -t testgoenv .     
docker run -it --rm --name my-running-app testgoenv 
chmod 777 ./build.sh