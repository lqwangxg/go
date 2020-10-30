git config --global user.email "lqwangxg@gmail.com"
git config --global user.name "lqwangxg"

git add . 
git commit -a -m "commit"

if [ -z $1 ];then
    echo "dest branch is empty."
else
    git push -u origin $1
fi
