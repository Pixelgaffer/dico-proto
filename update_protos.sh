set -e
LATEST_COMMIT_MSG=`git log -1 --pretty=%B`
protoc --go_out=. *.proto
echo "commiting code"
git add *.pb.go
git commit -m "[skip ci] sync: $LATEST_COMMIT_MSG"
