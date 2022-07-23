1. To run a container use:
	dokcer run -itd --name test fosfordevops/csvgenerator:latest

2. Check if container is running or not
	docker ps
   if there is nothing then container stopped, so check the logs
	docker log test
   test here is container name
   This will give the details why container stopped by viewing the logs.
   It says some file not attached.

2.1. Remove the created container using 
	docker rm test

3. Build the file in your system at any location ( suppose root)
	mkdir /root/Desktop/t/csvserver/inputdata

4. Now run container using by attaching this file to it
	docker run -itd -v /root/Desktop/t/hackathon/appserver/solution/inputFile:/csvserver/inputdata --name test fosfordevops/csvgenerator:latest

5. Create a gencsv.sh file to generate the given format of csv:
	0, 234
	1, 98
	2, 34
  its alrady present in this folder. Run it using
	sh gencsv.sh > inputFile
  It will give inputFile as output with 10 random values.
  To run with argument give command
	sh gencsv.sh 15 > inputFile
  It will give 15 random value in the file.

6. Delete previous container using
	docker rm -f test
	-f: forcefully because container is still running
7. Now run container with this generated file:
	docker run -itd -v /root/Desktop/t/hackathon/appserver/solution/inputFile:/csvserver/inputdata --name test fosfordevops/csvgenerator:latest

8. Get the shell using
	docker attach test

9. Now publish the port and run again using
docker run -itd -v /root/Desktop/t/hackathon/appserver/solution/inputFile:/csvserver/inputdata -p 9393:9300 --name test fosfordevops/csvgenerator:latest

10. Delete this container and run again this command to add environment variable
docker run -itd -v /root/Desktop/t/hackathon/appserver/solution/inputFile:/csvserver/inputdata -p 9393:9300 -e CSVSERVER_BORDER=Orange --name test fosfordevops/csvgenerator:latest

