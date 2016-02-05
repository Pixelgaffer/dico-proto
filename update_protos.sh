set -e
git checkout master
LATEST_COMMIT_MSG=`git log -1 --pretty=%B`
TMPDIR=`mktemp -d -t protosXXXXXX`
echo "TMPDIR: $TMPDIR"
cp ./*.proto $TMPDIR
git checkout go
protoc --go_out=. --proto_path=$TMPDIR $TMPDIR/*.proto
echo "commiting code"
git add *.pb.go
git commit -m "sync: $LATEST_COMMIT_MSG"
git checkout master
echo "deleting $TMPDIR"
rm -rf $TMPDIR
