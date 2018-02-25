cd Client &&
npm install &&
npm run buildstaging &&
cd .. &&
GOOS=linux GOARCH=amd64 go build -v ./Backend/main.go