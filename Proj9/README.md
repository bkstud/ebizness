# Project 9 - auto deploy backend and fronend to cloud

## For this project I created 2 repositiories:
### Backend:
Link: https://github.com/bkstud/azure-test-backend <br>
Simple python flask rest server that is automatically build as docker and deployed to azure cloud. <br>
It has single endpoint / that returns "Hello World!". <br>
Application runs on port 80. <br>
It was deployed on azure to following link: <br>
https://ebizproj9-backend.azurewebsites.net/
### Frontend:
Link: https://github.com/bkstud/azure-test-frontend <br>
Single component single hook react gui app that performs GET request on backend server
and prints 'Backend_says: {BACKEND_RESPONSE}'. <br>
Also built as docker and pushed to docker hub and then deployed to azure web app per repo push. <br>
It was deployed on azure to following link: <br>
http://ebizproj9-frontend.azurewebsites.net/ <br>
It loads quite long as the free 1gb resource plan is a bit too slow for react, but it starts after some time.