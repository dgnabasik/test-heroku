echo dgnabasik
git pull
echo -n "push?"
read
git add --all :/
git commit -am "Release 1.0.0"
git push -u # herouk master
