echo dgnabasik
git pull https://github.com/dgnabasik/test-heroku 
echo -n "push?"
read
git add --all :/
git commit -am "Release 1.0.1"
git push https://github.com/dgnabasik/test-heroku master
