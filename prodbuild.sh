cd Client &&
npm run build &&
cd .. &&
GOOS=linux GOARCH=amd64 go build -v ./Backend/main.go